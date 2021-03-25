/*
StructDump takes go struct and recursively dumps all fields to standard output in "path" format.

The output format is suitable for pasting to excel sheets, documentation, or generally to any text files.

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
*/
package structdump

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

// StructFieldName returns provided StructField name
func StructFieldName(f reflect.StructField) string {
	return f.Name
}

// JsonTagName returns json tag name of provided struct field or empty string. "-" tags are skipped ( JsonTagName returns empty string )
func JsonTagName(f reflect.StructField) string {
	jsonTag := f.Tag.Get("json")

	tagSplit := strings.Split(jsonTag, ",")

	name := tagSplit[0]

	if name == "" {
		name = f.Name
	}

	if name == "-" {
		return ""
	}

	return name
}

// Dump is convenience function which calls StructDump.Dump(t) with default configuration values
func Dump(t reflect.Type) {
	StructDump{}.Dump(t)
}

// NameFunc is used by StructDump to determine displayed name of struct field, if NameFunc returns empty string, the field is skipped.
type NameFunc func(reflect.StructField) string

// StructDump holds configuration for struct dumping
//
type StructDump struct {
	StopTypes []string  // List of type names which should be displayed as is without decomposition to smaller parts eg. Time
	Output    io.Writer // Output for dumping, defaults to os.Stdout
	NameFunc  NameFunc  // Function to get struct field name, default is StructFieldName
	prefix    string
	depth     int
}

func (sd StructDump) Dump(t reflect.Type) {
	if sd.depth == 0 {
		sd = sd.appendToPrefix(t.Name())
	}

	sd.depth += 1

	if sd.depth > 100 {
		panic("max depth of 100 reached")
	}

	sd.doDump(t)
}

func (sd StructDump) doDump(t reflect.Type) {
	if sd.StopTypes != nil {
		name := t.Name()

		for _, stopType := range sd.StopTypes {
			if stopType == name {
				sd.println(" " + name)
				return
			}
		}
	}

	switch t.Kind() {
	case reflect.Slice:
		sd.appendToPrefix("[]").doDump(t.Elem())
	case reflect.Interface:
		sd.println(" interface{}")
	case reflect.Ptr:
		sd.appendToPrefix("*").doDump(t.Elem())
	case reflect.Struct:
		nameFunc := sd.NameFunc
		if nameFunc == nil {
			nameFunc = StructFieldName
		}

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			name := nameFunc(field)

			if name == "" {
				continue
			}

			sd.appendToPrefix("." + name).Dump(field.Type)
		}
	default:
		sd.println(" " + t.Name())
	}
}

func (sd StructDump) println(in ...interface{}) {
	out := sd.Output

	if out == nil {
		out = os.Stdout
	}

	fmt.Fprint(out, sd.prefix)
	fmt.Fprintln(out, in...)
}

func (sd StructDump) appendToPrefix(prefix string) StructDump {
	sd.prefix += prefix

	return sd
}
