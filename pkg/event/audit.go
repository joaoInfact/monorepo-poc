// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package event

import (
	"encoding/json"

	"github.com/pantsbuild/example-golang/pkg/uuid"
)

type Event struct {
	ID         string `json:"id"`
	PersonName string `json:"person_name"`
}

func NewAuditEvent(personName string) *Event {
	return &Event{
		ID:         uuid.Generate(),
		PersonName: personName,
	}
}

func FromJSONStringToAuditEvent(personName string) *Event {
	parsedEvent := &Event{}
	err := json.Unmarshal([]byte(personName), parsedEvent)
	if err != nil {
		return nil
	}

	return parsedEvent
}
