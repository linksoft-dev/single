package i18n

import (
	"embed"
	"golang.org/x/text/language"
	"io/fs"
	"strings"
)

var langs map[language.Tag]fs.File

func AddLanguagesByFiles(files embed.FS) {

}

func GetMessage(lang language.Tag, identification string, template interface{}) string {
	return ""
}

func detectLanguage(fileName string) language.Tag {
	// Supomos que os arquivos sigam o padrão "i18n_Lang_Region.toml"
	parts := strings.Split(fileName, "_")
	if len(parts) >= 3 {
		lang := language.Make(parts[1])
		return lang
	}
	return language.Und
}
