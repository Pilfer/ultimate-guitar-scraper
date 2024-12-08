package cmd

import (
    "bufio"
    "errors"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "regexp"
    "strings"
    "syscall"

    "github.com/Pilfer/ultimate-guitar-scraper/pkg/ultimateguitar"
    "github.com/urfave/cli"
    "golang.org/x/term"
)

var GetAll = cli.Command{
    Name:        "get_all",
    Usage:       "Fetches all saved tabs/songs for Ultimate Guitar. Requires you to login.",
    Description: "Fetches all saved tabs/songs for Ultimate Guitar. Requires you to login.",
    Aliases:     []string{"a"},
    Flags: []cli.Flag{
        cli.StringFlag{
            Name:  "user",
            Usage: "--user {your_email}",
        },
        cli.StringFlag{
            Name:  "output",
            Usage: "--output {output path}. Default './out'",
        },
        cli.BoolFlag{
            Name:  "debug",
            Usage: "Enable debug logging",
        },
    },
    Action: GetAllTabs,
}

func GetAllTabs(c *cli.Context) {
    if c.Bool("debug") {
        log.SetFlags(log.LstdFlags | log.Lshortfile)
    }

    var user, password string
    var err error

    if c.IsSet("user") {
        user = c.String("user")
    } else {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Username: ")
        user, err = reader.ReadString('\n')
        user = strings.TrimSpace(user)
        if err != nil {
            log.Fatalf("Error reading username: %v", err)
        }
    }

    fmt.Print("Password: ")
    bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        log.Fatalf("Error reading password: %v", err)
    }
    fmt.Println() // Move to the next line after password input
    password = strings.TrimSpace(string(bytePassword))

    tabs, err := fetchAllTabs(user, password)
    if err != nil {
        log.Fatalf("Error fetching tabs: %v", err)
    }

    path := "./out/"
    if c.IsSet("output") {
        path = c.String("output")
    }

    path, err = filepath.Abs(path)
    if err != nil {
        log.Fatalf("Error resolving output path: %v", err)
    }

    err = writeTabs(path, tabs)
    if err != nil {
        log.Fatalf("Error writing tabs: %v", err)
    }
    fmt.Printf("Wrote %d tabs to %s\n", len(tabs), path)
}

func fetchAllTabs(user string, password string) ([]ultimateguitar.TabResult, error) {
    var tabResults []ultimateguitar.TabResult
    s := ultimateguitar.New()
    res, err := s.Login(user, password)
    if err != nil {
        return tabResults, fmt.Errorf("login error: %w", err)
    }
    if res == "Failed to login" {
        return tabResults, errors.New("login failed: invalid username or password")
    }
    tabResults, err = s.GetAll()
    if err != nil {
        return tabResults, fmt.Errorf("error fetching tabs: %w", err)
    }
    return tabResults, nil
}

func writeTabs(path string, tabs []ultimateguitar.TabResult) error {
    if path == "" {
        return errors.New("writeTabs: requires path")
    }
    if _, err := os.Stat(path); os.IsNotExist(err) {
        err := os.MkdirAll(path, 0775)
        if err != nil {
            return fmt.Errorf("error creating output directory: %w", err)
        }
    }

    fmt.Println("Output directory:", path)

    for _, tab := range tabs {
        artist := tab.ArtistName
        songName := tab.SongName
        capo := tab.Capo
        content := tab.Content
        content = fmt.Sprintf("{artist: %s}\n{title: %s}\n{capo: %d}\n%s", artist, songName, capo, content)

        regex := regexp.MustCompile(`\[(/?tab|/?ch)\]`)
        content = regex.ReplaceAllString(content, "")

        filename := fmt.Sprintf("%s-%s.crd", sanitizeFilename(artist), sanitizeFilename(songName))
        filePath := filepath.Join(path, filename)
        err := os.WriteFile(filePath, []byte(content), 0644)
        if err != nil {
            log.Printf("Error writing file %s: %v", filename, err)
            continue
        }
    }
    return nil
}

// sanitizeFilename removes or replaces characters that are invalid in filenames
func sanitizeFilename(name string) string {
    // Remove any characters that are not letters, numbers, spaces, hyphens, or underscores
    regex := regexp.MustCompile(`[<>:"/\\|?*]`)
    return regex.ReplaceAllString(name, "")
}
