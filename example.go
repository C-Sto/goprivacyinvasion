package main

import (
	"log"
	_ "unsafe"

	"github.com/C-Sto/goprivacyinvasion/pkg/secretive"
)

//go:linkname PrivateMethod github.com/C-Sto/goprivacyinvasion/pkg/secretive.ExportedStruct.unexportedGetField2
func PrivateMethod(s secretive.ExportedStruct) string

func main() {
	a := secretive.NewExportedStruct()
	log.Println(PrivateMethod(a))

}
