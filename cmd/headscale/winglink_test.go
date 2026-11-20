package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/juanfont/headscale/hscontrol/types"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The deploy config carries the names machines see (docs/naming.md in
// internal-docs): the DERP region in netcheck, and the base domain that
// MagicDNS names end in. The base domain needs no public DNS.
func TestWinglinkDeployConfigNames(t *testing.T) {
	path, err := os.Getwd()
	require.NoError(t, err)

	cfgFile := filepath.Join(t.TempDir(), "config.yaml")
	require.NoError(t, os.Symlink(filepath.Clean(path+"/../../deploy/headscale.yaml"), cfgFile))
	require.NoError(t, types.LoadConfig(cfgFile, true))

	assert.Equal(t, "winglink.flockadoo.internal", viper.GetString("dns.base_domain"))
	assert.Equal(t, "flockadoo", viper.GetString("derp.server.region_code"))
	assert.Equal(t, "Flockadoo", viper.GetString("derp.server.region_name"))
}
