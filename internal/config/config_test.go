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

func TestMapConfig(t *testing.T) {
	validConfig := `{
    "username": "someuser",
    "remotes": {
        "origin": {
            "url": "someurl"
        }
    },
    "branches": {
	"main": {"active": true}
    }
}`
	c, err := mapConfig([]byte(validConfig))
	assert.Nil(t, err)
	assert.Equal(t, "someuser", c.Username)
	assert.Equal(t, "someurl", c.Remotes["origin"].Url)
	assert.Equal(t, true, c.Branches["main"].Active)
}
