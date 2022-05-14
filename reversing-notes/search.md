# Search notes

```
:method: GET
:path: /api/v1/tab/search?title=johnny%20flynn&page=1&type[]=300&type[]=800&official[]=0&official[]=1&
:authority: api.ultimate-guitar.com
:scheme: https
accept-charset: utf-8
accept: application/json
user-agent: UGT_ANDROID/5.8.1 (Nexus 6; Android 9)
x-ug-client-id: 4d0aefcf9016df5d
x-ug-api-key: f2cd24ffe84493deadbeef3d307c4ae3
x-ug-client-backend: 1
x-ug-client-build: 3000652
x-ug-ab-variation-id: 8513
accept-encoding: gzip

```


Basic search: `https://api.ultimate-guitar.com/api/v1/tab/search?title=johnny%20flynn&page=1&type[]=0&official[]=0&official[]=1&`

Ukulele: `https://api.ultimate-guitar.com/api/v1/tab/search?title=johnny%20flynn&page=1&type[]=800&official[]=0&official[]=1&`

Ukulele + Chords: `https://api.ultimate-guitar.com/api/v1/tab/search?title=johnny%20flynn&page=1&type[]=300&type[]=800&official[]=0&official[]=1&`


### Search Suggestions:  

Suggestions have a limit of 5 characters: https://api.ultimate-guitar.com/api/v1/tab/suggestion?q=joh

### Filtering:  

Basic search for "johnny f": `https://api.ultimate-guitar.com/api/v1/tab/search?title=johnny%20f&page=1&type[]=0&&display_songs=1&filter=1`  

..and with "official" and "chords" types: `https://api.ultimate-guitar.com/api/v1/tab/search?title=johnny%20f&page=1&type[]=900&type[]=300&&display_songs=1&filter=1`  

with all tunings: `https://api.ultimate-guitar.com/api/v1/tab/search?title=johnny%20f&page=1&&tuning[]=standard&tuning[]=half-step%20down&tuning[]=b%20tuning&tuning[]=c%20tuning&tuning[]=d%20tuning&tuning[]=drop%20a&tuning[]=drop%20a%23&tuning[]=drop%20b&tuning[]=drop%20c&tuning[]=drop%20c%23&tuning[]=drop%20d&tuning[]=open%20c&tuning[]=open%20d&tuning[]=open%20e&tuning[]=open%20g&display_songs=1&filter=1`  


---  

### Search by known chords   

I've also seen references to: `tab/search/chords`  


```smali
Function<>14990(1 params, 15 registers, 0 symbols):
	LoadConstUndefined  	Reg8:3
	LoadConstUndefined  	Reg8:2
	GetArgumentsLength  	Reg8:0, Reg8:2
	LoadConstZero       	Reg8:1
	JNotGreater         	Addr8:12, Reg8:0, Reg8:1
	GetArgumentsPropByVal	Reg8:0, Reg8:1, Reg8:2
	JStrictNotEqual     	Addr8:16, Reg8:3, Reg8:0
	NewObjectWithBuffer 	Reg8:0, UInt16:3, UInt16:3, UInt16:17085, UInt16:32135
	Jmp                 	Addr8:6
	GetArgumentsPropByVal	Reg8:0, Reg8:1, Reg8:2
	GetById             	Reg8:4, Reg8:0, UInt8:1, UInt16:18099
	; Oper[3]: String(18099) 'page'

	GetById             	Reg8:5, Reg8:0, UInt8:2, UInt16:16456
	; Oper[3]: String(16456) 'chords'

	GetById             	Reg8:0, Reg8:0, UInt8:3, UInt16:24956
	; Oper[3]: String(24956) 'only_this_chords'

	GetEnvironment      	Reg8:1, UInt8:0
	LoadFromEnvironment 	Reg8:2, Reg8:1, UInt8:24
	NewObject           	Reg8:1
	PutNewOwnById       	Reg8:1, Reg8:5, UInt16:16456
	; Oper[2]: String(16456) 'chords'

	PutNewOwnById       	Reg8:1, Reg8:4, UInt16:18099
	; Oper[2]: String(18099) 'page'

	PutNewOwnById       	Reg8:1, Reg8:0, UInt16:24956
	; Oper[2]: String(24956) 'only_this_chords'

	LoadConstString     	Reg8:0, UInt16:10815
	; Oper[1]: String(10815) 'tab/search/chords'

	Call3               	Reg8:2, Reg8:2, Reg8:3, Reg8:0, Reg8:1
	GetByIdShort        	Reg8:1, Reg8:2, UInt8:4, UInt8:188
	; Oper[3]: String(188) 'then'

	CreateEnvironment   	Reg8:0
	CreateClosure       	Reg8:0, Reg8:0, UInt16:14991
	Call2               	Reg8:0, Reg8:1, Reg8:2, Reg8:0
	Ret                 	Reg8:0
EndFunction
```




---  


### Response Struct 

```go
type SearchResult struct {
	Tabs []struct {
		ID                 int     `json:"id"`
		SongID             int     `json:"song_id"`
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
			IsAcoustic       int         `json:"is_acoustic"`
			TonalityName     string      `json:"tonality_name"`
			Performance      Performance `json:"performance"`
			RecordingArtists []RecordingArtist `json:"recording_artists"`
		} `json:"recording"`
	} `json:"tabs"`
	Artists []string `json:"artists"`
}
```


Ultimate-Guitar's android app uses React Native, which compiles the JS to bytecode with `Hermes`. Here's one of the decompiled functions that kicks off the search:  

```
Function<>14986(3 params, 22 registers, 1 symbols):
	CreateEnvironment   	Reg8:0
	LoadParam           	Reg8:12, UInt8:1
	LoadConstUndefined  	Reg8:4
	LoadConstUndefined  	Reg8:5
	GetArgumentsLength  	Reg8:1, Reg8:5
	LoadConstUInt8      	Reg8:2, UInt8:2
	JNotGreater         	Addr8:12, Reg8:1, Reg8:2
	GetArgumentsPropByVal	Reg8:1, Reg8:2, Reg8:5
	JStrictNotEqual     	Addr8:16, Reg8:4, Reg8:1
	NewObject           	Reg8:1
	LoadConstUInt8      	Reg8:3, UInt8:1
	PutNewOwnById       	Reg8:1, Reg8:3, UInt16:18099
	; Oper[2]: String(18099) 'page'

	Jmp                 	Addr8:6
	GetArgumentsPropByVal	Reg8:1, Reg8:2, Reg8:5
	GetArgumentsLength  	Reg8:2, Reg8:5
	LoadConstUInt8      	Reg8:3, UInt8:3
	Greater             	Reg8:6, Reg8:2, Reg8:3
	JmpFalse            	Addr8:11, Reg8:6
	GetArgumentsPropByVal	Reg8:2, Reg8:3, Reg8:5
	StrictNeq           	Reg8:6, Reg8:4, Reg8:2
	Not                 	Reg8:2, Reg8:6
	JmpFalse            	Addr8:7, Reg8:6
	GetArgumentsPropByVal	Reg8:2, Reg8:3, Reg8:5
	StoreToEnvironment  	Reg8:0, UInt8:0, Reg8:2
	GetById             	Reg8:10, Reg8:1, UInt8:1, UInt16:18099
	; Oper[3]: String(18099) 'page'

	GetById             	Reg8:8, Reg8:1, UInt8:2, UInt16:21811
	; Oper[3]: String(21811) 'official'

	GetByIdShort        	Reg8:9, Reg8:1, UInt8:3, UInt8:201
	; Oper[3]: String(201) 'type'

	GetById             	Reg8:7, Reg8:1, UInt8:4, UInt16:21046
	; Oper[3]: String(21046) 'tuning'

	GetById             	Reg8:6, Reg8:1, UInt8:5, UInt16:25009
	; Oper[3]: String(25009) 'display_songs'

	GetByIdShort        	Reg8:5, Reg8:1, UInt8:6, UInt8:69
	; Oper[3]: String(69) 'artist_name'

	GetByIdShort        	Reg8:1, Reg8:1, UInt8:7, UInt8:125
	; Oper[3]: String(125) 'filter'

	GetEnvironment      	Reg8:2, UInt8:0
	LoadFromEnvironment 	Reg8:3, Reg8:2, UInt8:24
	NewObject           	Reg8:2
	GetById             	Reg8:11, Reg8:12, UInt8:8, UInt16:20240
	; Oper[3]: String(20240) 'toLowerCase'

	Call1               	Reg8:11, Reg8:11, Reg8:12
	PutNewOwnByIdShort  	Reg8:2, Reg8:11, UInt8:231
	; Oper[2]: String(231) 'title'

	PutNewOwnById       	Reg8:2, Reg8:10, UInt16:18099
	; Oper[2]: String(18099) 'page'

	PutNewOwnByIdShort  	Reg8:2, Reg8:9, UInt8:201
	; Oper[2]: String(201) 'type'

	PutNewOwnById       	Reg8:2, Reg8:8, UInt16:21811
	; Oper[2]: String(21811) 'official'

	PutNewOwnById       	Reg8:2, Reg8:7, UInt16:21046
	; Oper[2]: String(21046) 'tuning'

	PutNewOwnById       	Reg8:2, Reg8:6, UInt16:25009
	; Oper[2]: String(25009) 'display_songs'

	PutNewOwnByIdShort  	Reg8:2, Reg8:5, UInt8:69
	; Oper[2]: String(69) 'artist_name'

	PutNewOwnByIdShort  	Reg8:2, Reg8:1, UInt8:125
	; Oper[2]: String(125) 'filter'

	LoadConstString     	Reg8:1, UInt16:10814
	; Oper[1]: String(10814) 'tab/search'

	Call3               	Reg8:2, Reg8:3, Reg8:4, Reg8:1, Reg8:2
	GetByIdShort        	Reg8:1, Reg8:2, UInt8:9, UInt8:188
	; Oper[3]: String(188) 'then'

	CreateClosure       	Reg8:0, Reg8:0, UInt16:14987
	Call2               	Reg8:0, Reg8:1, Reg8:2, Reg8:0
	Ret                 	Reg8:0
EndFunction
```