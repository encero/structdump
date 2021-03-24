package structdump

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

func StructFieldName(f reflect.StructField) string {
	return f.Name
}

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

type NameFunc func(reflect.StructField) string

type StructDump struct {
	StopTypes []string
	Output    io.Writer
	NameFunc  NameFunc
	prefix    string
	depth     int
}

func (sd StructDump) Dump(t reflect.Type) {
	if sd.depth == 0 {
		sd = sd.appendToPrefix(t.Name())
	}

	sd.depth += 1

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
