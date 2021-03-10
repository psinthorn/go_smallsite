package configs

import "html/template"

// AppConfig is store application configuration data
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
