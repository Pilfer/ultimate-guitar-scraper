package scraper

// Here are the rough known tab types - I didn't verify too closely.
const (
	// TabTypeVideo type
	TabTypeVideo = 100
	// TabTypeTabs type
	TabTypeTabs = 200
	// TabTypeChords type
	TabTypeChords = 300
	// TabTypeBass type
	TabTypeBass = 400
	// TabTypePro type
	TabTypePro = 500
	// TabTypePower type
	TabTypePower = 600
	// TabTypeDrums type
	TabTypeDrums = 700
	// TabTypeUkulele type
	TabTypeUkulele = 800
	// TabTypeAll type
	TabTypeAll = 1000
)

// TabResult struct - this is what we get when we fetch a tab by id.
type TabResult struct {
	ID                 int         `json:"id"`
	SongName           string      `json:"song_name"`
	ArtistName         string      `json:"artist_name"`
	Type               string      `json:"type"`
	Part               string      `json:"part"`
	Version            int         `json:"version"`
	Votes              int         `json:"votes"`
	Rating             float64     `json:"rating"`
	Date               string      `json:"date"`
	Status             string      `json:"status"`
	PresetID           int         `json:"preset_id"`
	TabAccessType      string      `json:"tab_access_type"`
	TpVersion          int         `json:"tp_version"`
	TonalityName       string      `json:"tonality_name"`
	VersionDescription interface{} `json:"version_description"`
	Verified           int         `json:"verified"`
	Recording          struct {
		IsAcoustic       int           `json:"is_acoustic"`
		TonalityName     string        `json:"tonality_name"`
		Performance      interface{}   `json:"performance"`
		RecordingArtists []interface{} `json:"recording_artists"`
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
			IsAcoustic       int           `json:"is_acoustic"`
			TonalityName     string        `json:"tonality_name"`
			Performance      interface{}   `json:"performance"`
			RecordingArtists []interface{} `json:"recording_artists"`
		} `json:"recording"`
	} `json:"versions"`
	Recommended []struct {
		ID                 int         `json:"id"`
		SongName           string      `json:"song_name"`
		ArtistName         string      `json:"artist_name"`
		Type               string      `json:"type"`
		Part               string      `json:"part"`
		Version            int         `json:"version"`
		Votes              int         `json:"votes"`
		Rating             float64     `json:"rating"`
		Date               string      `json:"date"`
		Status             string      `json:"status"`
		PresetID           int         `json:"preset_id"`
		TabAccessType      string      `json:"tab_access_type"`
		TpVersion          int         `json:"tp_version"`
		TonalityName       string      `json:"tonality_name"`
		VersionDescription interface{} `json:"version_description"`
		Verified           int         `json:"verified"`
		Recording          struct {
			IsAcoustic       int           `json:"is_acoustic"`
			TonalityName     string        `json:"tonality_name"`
			Performance      interface{}   `json:"performance"`
			RecordingArtists []interface{} `json:"recording_artists"`
		} `json:"recording"`
	} `json:"recommended"`
	UserRating  int           `json:"userRating"`
	Difficulty  string        `json:"difficulty"`
	Tuning      string        `json:"tuning"`
	Capo        int           `json:"capo"`
	URLWeb      string        `json:"urlWeb"`
	Strumming   []interface{} `json:"strumming"`
	VideosCount int           `json:"videosCount"`
	ProBrother  interface{}   `json:"pro_brother"`
	Contributor struct {
		UserID   int    `json:"user_id"`
		Username string `json:"username"`
	} `json:"contributor"`
	Applicature []struct {
		Chord      string `json:"chord"`
		Variations []struct {
			ID        string        `json:"id"`
			ListCapos []interface{} `json:"listCapos"`
			NoteIndex int           `json:"noteIndex"`
			Notes     []int         `json:"notes"`
			Frets     []int         `json:"frets"`
			Fingers   []int         `json:"fingers"`
			Fret      int           `json:"fret"`
		} `json:"variations"`
	} `json:"applicature"`
	Content string `json:"content"`
}
