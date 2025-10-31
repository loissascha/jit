package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFilepath(t *testing.T) {
	path := GetMainConfigFilePath("/my/base/path")
	assert.Equal(t, "/my/base/path/.jit/config", path)
}

func TestIsJitProject(t *testing.T) {
	is := IsJitProject("/my/base/path")
	assert.False(t, is)
}
