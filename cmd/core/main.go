// Copyright 2021 Pants project contributors.
// Licensed under the Apache License, Version 2.0 (see LICENSE).

package main

import (
	"fmt"

	"github.com/pantsbuild/example-golang/internal/core"
)

func main() {
	core.GenerateAuditEvent("joao")
	fmt.Println("this is core svc and should call audit under the hood")
}
