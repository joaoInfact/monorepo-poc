package audit

import (
	"fmt"

	"github.com/pantsbuild/example-golang/pkg/event"
)

func PrintAuditEvent(event event.Event) event.Event {
	fmt.Println("event created with ID:", event.ID, "and person name:", event.PersonName)
	return event
}
