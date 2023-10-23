package pb

import (
	"fmt"
	"golang.org/x/text/language"
	"regexp"
	"strconv"
	"strings"
)

const (
	codStrBetween = iota
	codEmbeddedV
)

var messages = map[*language.Tag]map[int]string{
	&language.BrazilianPortuguese: {
		codStrBetween: "valor precisa ser entre %d e %d",
	},
	// english is default language
	nil: {
		codStrBetween: "value length must be between %d and %d runes, inclusive",
	},
}
var directMessage = map[string]map[language.Tag]string{
	"embedded message failed validation": {
		language.BrazilianPortuguese: "parametros invalidos",
	},
}

func StrBetween(min, max int, lang *language.Tag) string {
	return fmt.Sprintf(messages[lang][codStrBetween], min, max)
}

func I18n(msg string, lang language.Tag) string {
	if v := directMessage[msg]; v != nil {
		return v[lang]
	}
	if r := I18nStrBetween(msg, &lang); r != "" {
		return r
	}
	return msg
}

type ListError []error

func (m ListError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

type Err interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
}

func ProcessValidationError(err error, tag language.Tag) error {
	if valError, ok := err.(ListError); ok {
		for _, e := range valError {
			if e2, ok2 := e.(Err); ok2 {
				e2.Field()
			}
			fmt.Print(e)
		}
	}
	return nil
}

func I18nStrBetween(msg string, lang *language.Tag) string {
	regex := regexp.MustCompile(`between (\d+) and (\d+)`)
	match := regex.FindStringSubmatch(msg)
	if len(match) == 3 {
		min, _ := strconv.Atoi(match[1])
		max, _ := strconv.Atoi(match[2])
		return StrBetween(min, max, lang)
	}
	return ""
}
