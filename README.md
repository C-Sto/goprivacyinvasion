# goprivacyinvasion

We can use linkname to get access to unexported data in packages. See Example.go for running example, but the gist of it:

secretive.go
```go
type ExportedStruct struct {
	unexportedF1 string
	unexportedF2 string
}
func NewExportedStruct() ExportedStruct {
	return ExportedStruct{
		"hello, this is exported field 1",
		"hello, this is exported field 2",
	}
}
func (e ExportedStruct) unexportedGetField2() string {
	return e.unexportedF2
}
```
used in another package:

example.go
```go
package main
import (
	"log"
	_ "unsafe"
	"secretive"
)
//go:linkname PrivateMethod secretive.ExportedStruct.unexportedGetField2
func PrivateMethod(s secretive.ExportedStruct) string

func main() {
	a := secretive.NewExportedStruct()
	log.Println(PrivateMethod(a))
}
```

output: `hello, this is exported field 2`
