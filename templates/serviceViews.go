package templates

import (
	"text/template"
)

var Tmpl = template.Must(template.ParseGlob("templates/views/*"))
