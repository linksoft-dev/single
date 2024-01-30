package genservice

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/linksoft-dev/single/comps/go/grpc/plugins/service/pb"
	"github.com/linksoft-dev/single/comps/go/str"
	"github.com/linksoft-dev/single/comps/go/tpl"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	log "github.com/sirupsen/logrus"
	"strings"
	"text/template"
)

//go:embed "template_service.go.tmpl"
var templates embed.FS

// ReportModule creates a report of all the target messages generated by the
// protoc run, writing the file into the /tmp directory.
type module struct {
	*pgs.ModuleBase
	pgsgo.Context
	templatePath string
}

// New configures the module with an instance of ModuleBase
func NewModule() pgs.Module {
	return &module{ModuleBase: &pgs.ModuleBase{}}
}

// Name is the identifier used to identify the module. This value is
// automatically attached to the BuildContext associated with the ModuleBase.
func (m *module) Name() string { return "reporter" }

// Execute is passed the target files as well as its dependencies in the pkgs
// map. The implementation should return a slice of Artifacts that represent
// the files to be generated. In this case, "/tmp/report.txt" will be created
// outside of the normal protoc flow.
func (m *module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	buf := &bytes.Buffer{}
	m.templatePath = m.Parameters().StrDefault("path", "./")

	for _, f := range targets {
		m.Push(f.Name().String()).Debug("reporting")
		for i, msg := range f.AllMessages() {
			m.generateCrud(msg, f)
			fmt.Fprintf(buf, "%03d. %v\n", i, msg.Name())
		}
		m.Pop()
	}

	return m.Artifacts()
}

var globalContext pgs.BuildContext

func (m *module) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.Context = pgsgo.InitContext(c.Parameters())
	globalContext = c
}

// generateCrud generate the crud file if proto message has crud option set to true
func (m *module) generateCrud(msg pgs.Message, f pgs.File) {
	var v bool
	_, err := msg.Extension(pb.E_Crud, &v)
	if err != nil {
		log.WithError(err).Warnf("error when try to check if message has Crud flag")
	}
	// this means that crud option is not set
	if v == false {
		return
	}

	var tableName string
	_, err = msg.Extension(pb.E_TableName, &tableName)
	if err != nil {
		log.WithError(err).Warnf("error when try to check if message has tableName flag")
	}

	fileName := m.Context.OutputPath(f).SetExt(".service.go").String()
	type fieldSettings struct {
		Field    pgs.Field
		Settings pb.Field
	}

	data := struct {
		MessageName string
		TableName   string
		Fields      []fieldSettings
	}{
		MessageName: msg.Name().String(),
		TableName:   tableName,
	}
	log.Infof("generating service %s in path '%s'", msg.Name().String(), fileName)

	// check all fields settings
	for _, field := range msg.Fields() {
		// perform the parse into Field object, this is the way to check the options present in each proto field about Field settings
		fs := pb.Field{}
		ok, err := field.Extension(pb.E_Field, &fs)
		if err != nil {
			log.WithError(err).Warnf("error when try to check if proto field has field settings")
		}

		// define default values
		if ok {
			data.Fields = append(data.Fields, fieldSettings{Field: field, Settings: fs})
		}
	}
	funcMap := &template.FuncMap{
		"toCamel": str.ToCamel,
		"toLower": strings.ToLower,
	}
	templateName := "template_service.go.tmpl"

	r, err := tpl.RenderTemplate(templates, templateName, data, funcMap)
	if err != nil {
		log.WithError(err).Fatalf("failed when render template %s", templateName)
		return
	}

	m.AddGeneratorFile(fileName, r)
}
