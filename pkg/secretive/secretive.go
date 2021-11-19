package secretive

//Example of unexported func
func notExported() string {
	return "hello from an unexported func"
}

//Example of Exported struct with unexported methods and fields
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
