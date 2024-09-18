package audit_test

import (
	"testing"

	"github.com/pantsbuild/example-golang/internal/audit"
	"github.com/pantsbuild/example-golang/pkg/event"
	"github.com/stretchr/testify/require"
)

func TestAudit(t *testing.T) {
	ev := event.NewAuditEvent("joao")
	require.NotNil(t, ev)

	e := audit.PrintAuditEvent(*ev)

	require.NotNil(t, e)
	require.Equal(t, "joao", e.PersonName)
}
