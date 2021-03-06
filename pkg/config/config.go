package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the application configurations
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	PortNumber    string
	InProduction  bool
}
