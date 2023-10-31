package validation

import (
	"embed"
	"github.com/linksoft-dev/single/comps/go/i18n"
)

//go:embed *.toml
var files embed.FS

// i18nMessages represent the common data used in the translations
type i18nMessages struct {
	Name           string
	Value          any //value passed by validation
	CurrentValue   any
	Value1, Value2 any
}

func init() {
	// add all supported languages by toml files
	i18n.AddLanguagesByFiles(files)
}
