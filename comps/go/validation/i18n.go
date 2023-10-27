package validation

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func Init() {
	bundle = i18n.NewBundle(language.English)
	bundle.LoadMessageFile("i18n_pt-BR.toml")
}
