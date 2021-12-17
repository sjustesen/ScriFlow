package core

import (
	"github.com/probonopd/go-scribus"
	config "github.com/sjustesen/scriflow/core/configuration"
)

func Init() {
	doc, err := scribus.NewScribusDocumentFromFile(config.GetProjectPath() + "Document-1.sla")
	if err != nil {

	}
}
