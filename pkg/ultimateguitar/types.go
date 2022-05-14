package ultimateguitar

type Difficulty int32

const (
	DifficultyAbsoluteBeginner Difficulty = 1
	DifficultyBeginner         Difficulty = 2
	DifficultyIntermediate     Difficulty = 3
	DifficultyAdvanced         Difficulty = 4
	// DifficultyNovice       Difficulty = iota
	// DifficultyIntermediate Difficulty = 1
	// DifficultyAdvance      Difficulty = 2
	// DifficultyNone         Difficulty = 3
)

type TabType int32

// Here are the rough known tab types
const (
	TabTypeVideo    TabType = 100 // ??
	TabTypeTabs     TabType = 200
	TabTypeChords   TabType = 300
	TabTypeBass     TabType = 400
	TabTypePro      TabType = 500
	TabTypePower    TabType = 600 // ??
	TabTypeDrums    TabType = 700
	TabTypeUkulele  TabType = 800
	TabTypeOfficial TabType = 900
	TabTypeAll      TabType = 1000
)

type ChordKey string

const (
	ChordKeyA         ChordKey = "A"
	ChordKeyAsus4     ChordKey = "Asus4"
	ChordKeyA6        ChordKey = "A6"
	ChordKeyA7        ChordKey = "A7"
	ChordKeyAmaj7     ChordKey = "Amaj7"
	ChordKeyA7sus4    ChordKey = "A7sus4"
	ChordKeyA_CSharp  ChordKey = "A/C#"
	ChordKeyAm        ChordKey = "Am"
	ChordKeyAm7       ChordKey = "Am7"
	ChordKeyAm_C      ChordKey = "Am/C"
	ChordKeyAm_D      ChordKey = "Am/D"
	ChordKeyAm_E      ChordKey = "Am/E"
	ChordKeyAm_F      ChordKey = "Am/F"
	ChordKeyAm_FSharp ChordKey = "Am/F#"
	ChordKeyAm_G      ChordKey = "Am/G"
	ChordKeyAm_GSharp ChordKey = "Am/G#"
	ChordKeyASharp    ChordKey = "A#"
	ChordKeyB         ChordKey = "B"
	ChordKeyB7        ChordKey = "B7"
	ChordKeyBsus2     ChordKey = "Bsus2"
	ChordKeyBm        ChordKey = "Bm"
	ChordKeyBm9       ChordKey = "Bm9"
	ChordKeyBm7       ChordKey = "Bm7"
	ChordKeyBb        ChordKey = "Bb"
	ChordKeyBb7       ChordKey = "Bb7"
	ChordKeyBbm       ChordKey = "Bbm"
	ChordKeyBbdim     ChordKey = "Bbdim"
	ChordKeyBbdim7    ChordKey = "Bbdim7"
	ChordKeyBbmaj7    ChordKey = "Bbmaj7"
	ChordKeyC         ChordKey = "C"
	ChordKeyCadd9     ChordKey = "Cadd9"
	ChordKeyCsus2     ChordKey = "Csus2"
	ChordKeyCmaj7     ChordKey = "Cmaj7"
	ChordKeyC7        ChordKey = "C7"
	ChordKeyCm        ChordKey = "Cm"
	ChordKeyC_B       ChordKey = "C/B"
	ChordKeyC_D       ChordKey = "C/D"
	ChordKeyC_E       ChordKey = "C/E"
	ChordKeyC_G       ChordKey = "C/G"
	ChordKeyCSharp7   ChordKey = "C#7"
	ChordKeyCSharp    ChordKey = "C#"
	ChordKeyCSharpm   ChordKey = "C#m"
	ChordKeyD         ChordKey = "D"
	ChordKeyD5        ChordKey = "D5"
	ChordKeyD7        ChordKey = "D7"
	ChordKeyD7sus4    ChordKey = "D7sus4"
	ChordKeyD7_A      ChordKey = "D7/A"
	ChordKeyD7_FSharp ChordKey = "D7/F#"
	ChordKeyDadd9     ChordKey = "Dadd9"
	ChordKeyDM7       ChordKey = "DM7"
	ChordKeyDmaj7     ChordKey = "Dmaj7"
	ChordKeyDsus2     ChordKey = "Dsus2"
	ChordKeyDsus4     ChordKey = "Dsus4"
	ChordKeyDdim_F    ChordKey = "Ddim/F"
	ChordKeyDm        ChordKey = "Dm"
	ChordKeyDm_C      ChordKey = "Dm/C"
	ChordKeyDm7       ChordKey = "Dm7"
	ChordKeyD_FSharp  ChordKey = "D/F#"
	ChordKeyDSharp    ChordKey = "D#"
	ChordKeyDSharp5   ChordKey = "D#5"
	ChordKeyDSharp7   ChordKey = "D#7"
	ChordKeyDSharpdim ChordKey = "D#dim"
	ChordKeyEb        ChordKey = "Eb"
	ChordKeyEb5       ChordKey = "Eb5"
	ChordKeyE         ChordKey = "E"
	ChordKeyE5        ChordKey = "E5"
	ChordKeyE7        ChordKey = "E7"
	ChordKeyEsus2     ChordKey = "Esus2"
	ChordKeyEsus4     ChordKey = "Esus4"
	ChordKeyE7_GSharp ChordKey = "E7/G#"
	ChordKeyEm        ChordKey = "Em"
	ChordKeyEm7       ChordKey = "Em7"
	ChordKeyEm_B      ChordKey = "Em/B"
	ChordKeyEm_G      ChordKey = "Em/G"
	ChordKeyE_GSharp  ChordKey = "E/G#"
	ChordKeyF         ChordKey = "F"
	ChordKeyF6        ChordKey = "F6"
	ChordKeyF7        ChordKey = "F7"
	ChordKeyF_A       ChordKey = "F/A"
	ChordKeyF_C       ChordKey = "F/C"
	ChordKeyF_G       ChordKey = "F/G"
	ChordKeyFmaj7     ChordKey = "Fmaj7"
	ChordKeyFm        ChordKey = "Fm"
	ChordKeyFSharp    ChordKey = "F#"
	ChordKeyFSharpm   ChordKey = "F#m"
	ChordKeyFSharpm7  ChordKey = "F#m7"
	ChordKeyG         ChordKey = "G"
	ChordKeyG6        ChordKey = "G6"
	ChordKeyG7        ChordKey = "G7"
	ChordKeyGmaj7     ChordKey = "Gmaj7"
	ChordKeyGsus4     ChordKey = "Gsus4"
	ChordKeyGm        ChordKey = "Gm"
	ChordKeyGm7       ChordKey = "Gm7"
	ChordKeyG_B       ChordKey = "G/B"
	ChordKeyG_Ab      ChordKey = "G/Ab"
	ChordKeyG_D       ChordKey = "G/D"
	ChordKeyGSharp    ChordKey = "G#"
	ChordKeyGSharpm   ChordKey = "G#m"
	ChordKeyAb        ChordKey = "Ab"
)

// TabResult struct - this is what we get when we fetch a tab by id.
type TabResult struct {
	ID                 int     `json:"id"`
	SongName           string  `json:"song_name"`
	ArtistName         string  `json:"artist_name"`
	Type               string  `json:"type"`
	Part               string  `json:"part"`
	Version            int     `json:"version"`
	Votes              int     `json:"votes"`
	Rating             float64 `json:"rating"`
	Date               string  `json:"date"`
	Status             string  `json:"status"`
	PresetID           int     `json:"preset_id"`
	TabAccessType      string  `json:"tab_access_type"`
	TpVersion          int     `json:"tp_version"`
	TonalityName       string  `json:"tonality_name"`
	VersionDescription *string `json:"version_description"`
	Verified           int     `json:"verified"`
	Recording          struct {
		IsAcoustic       int               `json:"is_acoustic"`
		TonalityName     string            `json:"tonality_name"`
		Performance      Performance       `json:"performance"`
		RecordingArtists []RecordingArtist `json:"recording_artists"`
	} `json:"recording"`
	Versions []struct {
		ID                 int     `json:"id"`
		SongName           string  `json:"song_name"`
		ArtistName         string  `json:"artist_name"`
		Type               string  `json:"type"`
		Part               string  `json:"part"`
		Version            int     `json:"version"`
		Votes              int     `json:"votes"`
		Rating             float64 `json:"rating"`
		Date               string  `json:"date"`
		Status             string  `json:"status"`
		PresetID           int     `json:"preset_id"`
		TabAccessType      string  `json:"tab_access_type"`
		TpVersion          int     `json:"tp_version"`
		TonalityName       string  `json:"tonality_name"`
		VersionDescription string  `json:"version_description"`
		Verified           int     `json:"verified"`
		Recording          struct {
			IsAcoustic       int               `json:"is_acoustic"`
			TonalityName     string            `json:"tonality_name"`
			Performance      Performance       `json:"performance"`
			RecordingArtists []RecordingArtist `json:"recording_artists"`
		} `json:"recording"`
	} `json:"versions"`
	Recommended []struct {
		ID                 int     `json:"id"`
		SongName           string  `json:"song_name"`
		ArtistName         string  `json:"artist_name"`
		Type               string  `json:"type"`
		Part               string  `json:"part"`
		Version            int     `json:"version"`
		Votes              int     `json:"votes"`
		Rating             float64 `json:"rating"`
		Date               string  `json:"date"`
		Status             string  `json:"status"`
		PresetID           int     `json:"preset_id"`
		TabAccessType      string  `json:"tab_access_type"`
		TpVersion          int     `json:"tp_version"`
		TonalityName       string  `json:"tonality_name"`
		VersionDescription *string `json:"version_description"`
		Verified           int     `json:"verified"`
		Recording          struct {
			IsAcoustic       int               `json:"is_acoustic"`
			TonalityName     string            `json:"tonality_name"`
			Performance      Performance       `json:"performance"`
			RecordingArtists []RecordingArtist `json:"recording_artists"`
		} `json:"recording"`
	} `json:"recommended"`
	UserRating int    `json:"userRating"`
	Difficulty string `json:"difficulty"`
	Tuning     string `json:"tuning"`
	Capo       int    `json:"capo"`
	URLWeb     string `json:"urlWeb"`
	// Strumming   []interface{} `json:"strumming"`
	VideosCount int `json:"videosCount"`
	// ProBrother  interface{} `json:"pro_brother"`
	Contributor struct {
		UserID   int    `json:"user_id"`
		Username string `json:"username"`
	} `json:"contributor"`
	Applicature []Applicature `json:"applicature"`
	Content     string        `json:"content"`
}

type ListCapo struct {
	Fret        int64 `json:"fret"`
	StartString int64 `json:"startString"`
	LastString  int64 `json:"lastString"`
	Finger      int64 `json:"finger"`
}

type Applicature struct {
	Chord      string `json:"chord"`
	Variations []struct {
		ID        string     `json:"id"`
		ListCapos []ListCapo `json:"listCapos"`
		NoteIndex int        `json:"noteIndex"`
		Notes     []int      `json:"notes"`
		Frets     []int      `json:"frets"`
		Fingers   []int      `json:"fingers"`
		Fret      int        `json:"fret"`
	} `json:"variations"`
}

type SearchResult struct {
	Tabs    []Tab    `json:"tabs"`
	Artists []string `json:"artists"`
}

type ArtistName string

type Tab struct {
	ID                 int64         `json:"id"`
	SongID             int64         `json:"song_id"`
	SongName           string        `json:"song_name"`
	ArtistID           int64         `json:"artist_id"`
	ArtistName         ArtistName    `json:"artist_name"`
	Type               Type          `json:"type"`
	Part               string        `json:"part"`
	Version            int64         `json:"version"`
	Votes              int64         `json:"votes"`
	Rating             float64       `json:"rating"`
	Date               string        `json:"date"`
	Status             Status        `json:"status"`
	PresetID           int64         `json:"preset_id"`
	TabAccessType      TabAccessType `json:"tab_access_type"`
	TpVersion          int64         `json:"tp_version"`
	TonalityName       TonalityName  `json:"tonality_name"`
	VersionDescription *string       `json:"version_description"`
	Verified           int64         `json:"verified"`
	Recording          Recording     `json:"recording"`
}

type Performance struct {
	Name string `json:"name"`
	// Serie     interface{}   `json:"serie"`
	// Venue     interface{}   `json:"venue"`
	DateStart int64  `json:"date_start"`
	DateEnd   int64  `json:"date_end"`
	Cancelled int64  `json:"cancelled"`
	Type      string `json:"type"`
	Comment   string `json:"comment"`
	// VideoUrls []interface{} `json:"video_urls"`
}

type TonalityName string

type Recording struct {
	IsAcoustic       int64             `json:"is_acoustic"`
	TonalityName     string            `json:"tonality_name"`
	Performance      interface{}       `json:"performance"`
	RecordingArtists []RecordingArtist `json:"recording_artists"`
}

type RecordingArtist struct {
	JoinField string `json:"join_field"`
	Artist    Artist `json:"artist"`
}

type Artist struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Status string

const (
	Approved Status = "approved"
)

type TabAccessType string

const (
	Public TabAccessType = "public"
)

type Type string

const (
	Power Type = "Power"
)
