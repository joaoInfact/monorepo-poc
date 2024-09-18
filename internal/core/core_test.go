package core_test

import (
	"testing"

	"github.com/pantsbuild/example-golang/internal/core"
	"github.com/stretchr/testify/require"
)

func TestCore(t *testing.T) {
	new := core.GenerateAuditEvent("joao")
	require.NotNil(t, new)
	require.Equal(t, "joao", new.PersonName)
}
