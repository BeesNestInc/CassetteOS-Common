package modmanagement_test

import (
	"os"
	"testing"

	modmanagement "github.com/BeesNestInc/CassetteOS-Common/pkg/mod_management"
	"github.com/stretchr/testify/assert"
)

func TestInstallableModules(t *testing.T) {
	// skip in GitHub Actions
	if os.Getenv("CI") != "" {
		t.Skip()
	}
	client, err := modmanagement.NewClient(modmanagement.ModManagementClientOpts{})
	assert.NoError(t, err)
	modules, err := client.InstallableModules()
	assert.NoError(t, err)

	t.Log(modules)
}

func TestInstallModule(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip()
	}
	err := modmanagement.RequireModule("doconverter", "/var/run/cassetteos")
	assert.NoError(t, err)
}

func TestInstallNoExistModule(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip()
	}
	err := modmanagement.RequireModule("abc", "/var/run/cassetteos")
	assert.ErrorIs(t, err, modmanagement.ErrModuleNoInStore)
}
