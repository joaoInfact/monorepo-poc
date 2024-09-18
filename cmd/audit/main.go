package main

import (
	"github.com/pantsbuild/example-golang/internal/audit"
	"github.com/pantsbuild/example-golang/pkg/event"
)

func main() {

	msgReceivedFromOutside := `{
	 "person_name": "joao"
	}`

	event := event.FromJSONStringToAuditEvent(msgReceivedFromOutside)
	if event == nil {
		return
	}

	audit.PrintAuditEvent(*event)
}
