package i18n

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
	"strings"
	"text/template"
)

var langs map[language.Tag]map[string]string

func AddLanguagesByFiles(fs embed.FS) {

	// Listar arquivos no filesystem embedido
	fileInfos, err := fs.ReadDir(".")
	if err != nil {
		//log.Fatalf("Erro ao listar arquivos no filesystem embedido: %v", err)
	}

	// Iterar sobre os arquivos
	for _, fileInfo := range fileInfos {
		// Ignorar diretórios
		if fileInfo.IsDir() {
			continue
		}

		// Ler o conteúdo do arquivo
		fileData, err := fs.ReadFile(fileInfo.Name())
		if err != nil {
			//log.Printf("Erro ao ler o arquivo %s: %v", fileInfo.Name(), err)
			continue
		}

		// Detectar a linguagem do arquivo
		lang := detectLanguage(fileInfo.Name())

		// Carregar as traduções do arquivo TOML
		config, err := toml.Load(string(fileData))
		if err != nil {
			//log.Printf("Erro ao carregar as traduções do arquivo %s: %v", fileInfo.Name(), err)
			continue
		}

		// Adicionar as traduções ao mapa de linguagens
		if langs[lang] == nil {
			langs = map[language.Tag]map[string]string{
				lang: {},
			}
		}

		for key, value := range config.ToMap() {
			langs[lang][key] = value.(string)
		}
	}
}

func AddLanguage(lang language.Tag, m map[string]string) {
	if langs[lang] == nil {
		langs[lang] = m
		return
	}

	for k, v := range m {
		langs[lang][k] = v
	}
}

func GetMessage(lang language.Tag, identification string, data interface{}) string {
	tmplString := langs[lang][identification]

	tmpl, err := template.New("message").Parse(tmplString)
	if err != nil {
		logrus.Errorf("erro while parse template")
		return ""
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, data)
	if err != nil {
		logrus.Errorf("erro while execute template")
		return ""
	}

	return tpl.String()
}

func detectLanguage(fileName string) language.Tag {

	// Supomos que os arquivos sigam o padrão "i18n_Lang_Region.toml"
	fileName = strings.ReplaceAll(strings.ToLower(fileName), "-", "_")
	fileName = strings.ReplaceAll(fileName, ".toml", "")
	parts := strings.Split(fileName, "_")
	if len(parts) >= 3 {
		langStr := strings.ToLower(fmt.Sprintf("%s-%s", parts[1], parts[2]))
		langStr = strings.ReplaceAll(langStr, ".toml", "")
		switch langStr {
		case "en-us":
			return language.AmericanEnglish
		default:
			return language.BrazilianPortuguese
		}
	}
	return language.BrazilianPortuguese
}
