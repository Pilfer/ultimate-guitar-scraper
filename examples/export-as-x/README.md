# Export as JSON, XML, YAML, MsgPack etc..  

Sometimes you'll want to dump out tabs into different formats. 

The Go ecosystem has a ton of different format serializers/deserializers.  


```go
import (
	"github.com/go-yaml/yaml"
	"github.com/vmihailenco/msgpack/v5"
)
```

then just do something like:  


```go
// We assube you've got a tab in a TabResult object already....
var tab ultimateguitar.TabResult
```

Then you can just dump to JSON  

```go
asJSON, err := json.Marshal(tab)
if err != nil {
  log.Fatal(err)
}
fmt.Println(string(asJSON))
```

...or dump to XML  

```go
asXML, err := xml.Marshal(tab)
if err != nil {
  log.Fatal(err)
}
fmt.Println(string(asXML))
```

...or dump to YAML  

```go
asYAML, err := yaml.Marshal(tab)
if err != nil {
  log.Fatal(err)
}
fmt.Println(string(asYAML))
```

...or dump to MsgPack  

```go
asMsgPack, err := msgpack.Marshal(tab)
if err != nil {
  log.Fatal(err)
}
fmt.Println(string(asMsgPack))
```