# stdjson

This package is for `encoding/json` support. `encoding/json` has no configuration.

## Usage

```go
import (
    sio "github.com/kaiterra/socket.io-go"
    jsonparser "github.com/kaiterra/socket.io-go/parser/json"
    "github.com/kaiterra/socket.io-go/parser/json/serializer/stdjson"
)

func main() {
    io := sio.NewServer(&sio.ServerConfig{
        ParserCreator: jsonparser.NewCreator(0, stdjson.New()),
    })

    io.Run()
}
```
