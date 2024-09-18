// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package core

import (
	"github.com/pantsbuild/example-golang/pkg/event"
)

func GenerateAuditEvent(personName string) (ev *event.Event) {
	return event.NewAuditEvent(personName)
}
