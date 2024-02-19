package tipper_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/tipper"
)

func TestUnwrapPtr(t *testing.T) {
	assert := assert.New(t)
	st := struct{ foo string }{}
	pst := &st
	ppst := &pst

	tt := []any{
		st,
		pst,
		ppst,
	}

	for _, v := range tt {
		typ := reflect.TypeOf(v)
		assert.Equal(reflect.TypeOf(st), tipper.UnwrapPtr(typ))
	}
}

func TestUnwrapPtrNil(t *testing.T) {
	assert := assert.New(t)
	assert.Nil(tipper.UnwrapPtr(nil))
}

type foo struct {
	Home string `env:"HOME"`
	Port int    `env:"PORT" envDefault:"3000"`
	Bar  *bar   `envPrefix:"BAR_"`
}

type bar struct {
	Password     string `env:"PASSWORD,unset"`
	IsProduction bool   `env:"PRODUCTION"`
	Zoo          zoo
	*baz         `envPrefix:"BAZ_"`
}

type zoo struct {
	Duration time.Duration `env:"DURATION"`
	Hosts    []string      `env:"HOSTS" envSeparator:":"`
}

type baz struct {
	TempFolder string         `env:"TEMP_FOLDER,expand" envDefault:"${HOME}/tmp"`
	StringInts map[string]int `env:"MAP_STRING_INT"`
}

func TestDump0(t *testing.T) {
	assert := assert.New(t)

	var f foo
	var pf *foo

	tt := []any{
		f,
		pf,
		foo{},
		&foo{},
	}

	for _, v := range tt {
		ss := tipper.Structs{}
		typ := reflect.TypeOf(v)
		tipper.Dump0(tipper.UnwrapPtr(typ), &ss)
		assert.Equal(tipper.Structs{
			{
				Name: "tipper_test.zoo",
				Fields: []tipper.Field{
					{
						Name: "Duration",
						Type: "time.Duration",
						Tags: []tipper.Tag{
							{Key: "env", Name: "DURATION", Options: nil},
						},
					},
					{
						Name: "Hosts",
						Type: "[]string",
						Tags: []tipper.Tag{
							{Key: "env", Name: "HOSTS", Options: nil},
							{Key: "envSeparator", Name: ":", Options: nil},
						},
					},
				},
			},
			{
				Name: "tipper_test.baz",
				Fields: []tipper.Field{
					{
						Name: "TempFolder",
						Type: "string",
						Tags: []tipper.Tag{
							{Key: "env", Name: "TEMP_FOLDER", Options: []string{"expand"}},
							{Key: "envDefault", Name: "${HOME}/tmp", Options: nil},
						},
					},
					{
						Name: "StringInts",
						Type: "map[string]int",
						Tags: []tipper.Tag{
							{Key: "env", Name: "MAP_STRING_INT", Options: nil},
						},
					},
				},
			},
			{
				Name: "tipper_test.bar",
				Fields: []tipper.Field{
					{
						Name: "Password",
						Type: "string",
						Tags: []tipper.Tag{
							{Key: "env", Name: "PASSWORD", Options: []string{"unset"}},
						},
					},
					{
						Name: "IsProduction",
						Type: "bool",
						Tags: []tipper.Tag{
							{Key: "env", Name: "PRODUCTION", Options: nil},
						},
					},
					{
						Name: "Zoo",
						Type: "tipper_test.zoo",
						Tags: nil,
					},
					{
						Name: "baz",
						Type: "*tipper_test.baz",
						Tags: []tipper.Tag{
							{Key: "envPrefix", Name: "BAZ_", Options: nil},
						},
					},
				},
			},
			{
				Name: "tipper_test.foo",
				Fields: []tipper.Field{
					{
						Name: "Home",
						Type: "string",
						Tags: []tipper.Tag{
							{Key: "env", Name: "HOME", Options: nil},
						},
					},
					{
						Name: "Port",
						Type: "int",
						Tags: []tipper.Tag{
							{Key: "env", Name: "PORT", Options: nil},
							{Key: "envDefault", Name: "3000", Options: nil},
						},
					},
					{
						Name: "Bar",
						Type: "*tipper_test.bar",
						Tags: []tipper.Tag{
							{Key: "envPrefix", Name: "BAR_", Options: nil},
						},
					},
				},
			},
		}, ss)
	}
}

func TestDump(t *testing.T) {
	assert := assert.New(t)

	var f foo
	var pf *foo

	tt := []any{
		f,
		pf,
		foo{},
		&foo{},
	}

	for _, v := range tt {
		ss := tipper.Dump(v)
		assert.Equal(`[
  {
    "name": "tipper_test.zoo",
    "fields": [
      {
        "name": "Duration",
        "type": "time.Duration",
        "tags": [
          {
            "key": "env",
            "name": "DURATION",
            "options": null
          }
        ]
      },
      {
        "name": "Hosts",
        "type": "[]string",
        "tags": [
          {
            "key": "env",
            "name": "HOSTS",
            "options": null
          },
          {
            "key": "envSeparator",
            "name": ":",
            "options": null
          }
        ]
      }
    ]
  },
  {
    "name": "tipper_test.baz",
    "fields": [
      {
        "name": "TempFolder",
        "type": "string",
        "tags": [
          {
            "key": "env",
            "name": "TEMP_FOLDER",
            "options": [
              "expand"
            ]
          },
          {
            "key": "envDefault",
            "name": "${HOME}/tmp",
            "options": null
          }
        ]
      },
      {
        "name": "StringInts",
        "type": "map[string]int",
        "tags": [
          {
            "key": "env",
            "name": "MAP_STRING_INT",
            "options": null
          }
        ]
      }
    ]
  },
  {
    "name": "tipper_test.bar",
    "fields": [
      {
        "name": "Password",
        "type": "string",
        "tags": [
          {
            "key": "env",
            "name": "PASSWORD",
            "options": [
              "unset"
            ]
          }
        ]
      },
      {
        "name": "IsProduction",
        "type": "bool",
        "tags": [
          {
            "key": "env",
            "name": "PRODUCTION",
            "options": null
          }
        ]
      },
      {
        "name": "Zoo",
        "type": "tipper_test.zoo",
        "tags": null
      },
      {
        "name": "baz",
        "type": "*tipper_test.baz",
        "tags": [
          {
            "key": "envPrefix",
            "name": "BAZ_",
            "options": null
          }
        ]
      }
    ]
  },
  {
    "name": "tipper_test.foo",
    "fields": [
      {
        "name": "Home",
        "type": "string",
        "tags": [
          {
            "key": "env",
            "name": "HOME",
            "options": null
          }
        ]
      },
      {
        "name": "Port",
        "type": "int",
        "tags": [
          {
            "key": "env",
            "name": "PORT",
            "options": null
          },
          {
            "key": "envDefault",
            "name": "3000",
            "options": null
          }
        ]
      },
      {
        "name": "Bar",
        "type": "*tipper_test.bar",
        "tags": [
          {
            "key": "envPrefix",
            "name": "BAR_",
            "options": null
          }
        ]
      }
    ]
  }
]`, ss.String())
	}
}

func TestDumpT(t *testing.T) {
	assert := assert.New(t)

	expected := tipper.Structs{
		{
			Name: "tipper_test.zoo",
			Fields: []tipper.Field{
				{
					Name: "Duration",
					Type: "time.Duration",
					Tags: []tipper.Tag{
						{Key: "env", Name: "DURATION", Options: nil},
					},
				},
				{
					Name: "Hosts",
					Type: "[]string",
					Tags: []tipper.Tag{
						{Key: "env", Name: "HOSTS", Options: nil},
						{Key: "envSeparator", Name: ":", Options: nil},
					},
				},
			},
		},
		{
			Name: "tipper_test.baz",
			Fields: []tipper.Field{
				{
					Name: "TempFolder",
					Type: "string",
					Tags: []tipper.Tag{
						{Key: "env", Name: "TEMP_FOLDER", Options: []string{"expand"}},
						{Key: "envDefault", Name: "${HOME}/tmp", Options: nil},
					},
				},
				{
					Name: "StringInts",
					Type: "map[string]int",
					Tags: []tipper.Tag{
						{Key: "env", Name: "MAP_STRING_INT", Options: nil},
					},
				},
			},
		},
		{
			Name: "tipper_test.bar",
			Fields: []tipper.Field{
				{
					Name: "Password",
					Type: "string",
					Tags: []tipper.Tag{
						{Key: "env", Name: "PASSWORD", Options: []string{"unset"}},
					},
				},
				{
					Name: "IsProduction",
					Type: "bool",
					Tags: []tipper.Tag{
						{Key: "env", Name: "PRODUCTION", Options: nil},
					},
				},
				{
					Name: "Zoo",
					Type: "tipper_test.zoo",
					Tags: nil,
				},
				{
					Name: "baz",
					Type: "*tipper_test.baz",
					Tags: []tipper.Tag{
						{Key: "envPrefix", Name: "BAZ_", Options: nil},
					},
				},
			},
		},
		{
			Name: "tipper_test.foo",
			Fields: []tipper.Field{
				{
					Name: "Home",
					Type: "string",
					Tags: []tipper.Tag{
						{Key: "env", Name: "HOME", Options: nil},
					},
				},
				{
					Name: "Port",
					Type: "int",
					Tags: []tipper.Tag{
						{Key: "env", Name: "PORT", Options: nil},
						{Key: "envDefault", Name: "3000", Options: nil},
					},
				},
				{
					Name: "Bar",
					Type: "*tipper_test.bar",
					Tags: []tipper.Tag{
						{Key: "envPrefix", Name: "BAR_", Options: nil},
					},
				},
			},
		},
	}

	assert.Equal(expected, tipper.DumpT[foo]())
	assert.Equal(expected[0], tipper.DumpT[zoo]()[0])
	assert.Equal(expected[1], tipper.DumpT[baz]()[0])
}
