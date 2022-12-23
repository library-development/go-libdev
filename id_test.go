package libdev_test

import (
	"testing"

	"lib.dev/libdev"
)

var inputs []libdev.ID = []libdev.ID{
	{
		Path: []libdev.Name{
			{
				Words: []string{"string"},
			},
		},
	},
	{
		Path: []libdev.Name{
			{
				Words: []string{"schema", "cafe"},
			},
			{
				Words: []string{"types"},
			},
			{
				Words: []string{"schema"},
			},
		},
	},
	{
		Path: []libdev.Name{
			{
				Words: []string{"lib", "dev"},
			},
			{
				Words: []string{"id"},
			},
		},
	},
}

func TestWriteGolang(t *testing.T) {

}
