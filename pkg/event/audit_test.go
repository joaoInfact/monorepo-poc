package event

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAudit(t *testing.T) {
	event := NewAuditEvent("test")
	require.Equal(t, "test", event.PersonName)
}
