package main

import (
	"os"

	"lib.dev/libdev"
)

var exampleLibrary = libdev.Library{
	Name:    libdev.Name{Words: []string{"example"}},
	Version: libdev.Version{Major: 1, Minor: 0, Patch: 0},
	Doc:     "This is an example library.",
	Discussion: []libdev.Comment{
		{
			Timestamp: 1671144675,
			AuthorID:  "example",
			Text:      "This is an example comment.",
			Replies: []libdev.Comment{
				{
					Timestamp: 1671144675,
					AuthorID:  "example",
					Text:      "This is an example reply.",
				},
			},
		},
	},
	Constants: []libdev.Constant{
		{
			Name: libdev.Name{Words: []string{"example"}},
			Doc:  "This is an example constant.",
		},
	},
	Types: []libdev.Type{
		{
			Name: libdev.Name{Words: []string{"example"}},
			Doc:  "This is an example type.",
		},
	},
	Functions: []libdev.Function{
		{
			Name: libdev.Name{Words: []string{"example"}},
			Doc:  "This is an example function.",
		},
	},
}

func main() {
	exampleLibrary.WriteHTML(os.Stdout)
}
