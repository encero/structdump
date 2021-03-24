package structdump_test

import (
	"reflect"

	"github.com/encero/structdump"
)

func ExampleStructDump_Dump() {
	type SimpleType struct {
		AnInt         int
		AnString      string
		AnStringSlice []string
	}

	dumper := structdump.StructDump{}

	dumper.Dump(reflect.TypeOf(SimpleType{}))

	// Output:
	// SimpleType.AnInt int
	// SimpleType.AnString string
	// SimpleType.AnStringSlice[] string
}
