package l

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

const (
	// LESS Variables
	less_variables = `(@[\w-]+)`
)

var LESS = internal.Register(MustNewLexer(

	&Config{
		Name:      "LESS",
		Aliases:   []string{"less"},
		Filenames: []string{"*.less"},
		MimeTypes: []string{"text/x-less-css"},
	},

	Rules{
		"root": {
			Include("basics"),
		},
		"basics": {
			{`\s+`, Text, nil},
			{less_variables, NameVariable, nil},
		},
	},
))
