package tipper

import (
	"encoding/json"
	"reflect"

	"github.com/fatih/structtag"
)

type Tag struct {
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Options []string `json:"options"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Tags []Tag  `json:"tags"`
}

type Struct struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Structs []Struct

func (ss Structs) String() string {
	j, _ := json.MarshalIndent(ss, "", "  ")
	return string(j)
}

func Dump(v any) Structs {
	t := reflect.TypeOf(v)
	t = unwrapPtr(t)
	ss := Structs{}
	dump0(t, &ss)
	return ss
}

func DumpT[T any]() Structs {
	var v T
	t := reflect.TypeOf(v)
	t = unwrapPtr(t)
	ss := Structs{}
	dump0(t, &ss)
	return ss
}

func dump0(t reflect.Type, acc *Structs) {
	if t == nil {
		return
	}

	if t.Kind() != reflect.Struct {
		return
	}

	st := Struct{}
	st.Name = t.String()

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		stf := Field{
			Name: f.Name,
			Type: f.Type.String(),
		}

		if tags, err := structtag.Parse(string(f.Tag)); err == nil {
			for _, tag := range tags.Tags() {
				stt := Tag{
					Key:     tag.Key,
					Name:    tag.Name,
					Options: tag.Options,
				}

				stf.Tags = append(stf.Tags, stt)
			}
		}

		st.Fields = append(st.Fields, stf)

		if r := unwrapPtr(f.Type); r != nil && r.Kind() == reflect.Struct {
			dump0(r, acc)
		}
	}

	*acc = append(*acc, st)
}

func unwrapPtr(t reflect.Type) reflect.Type {
	for {
		if t == nil {
			return nil
		}

		if t.Kind() != reflect.Ptr {
			return t
		}

		t = t.Elem()
	}
}
