package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPluginBaseArgument(t *testing.T) {
	assert.False(t, IsPluginBaseArgument("_"))
	assert.True(t, IsPluginBaseArgument(ArgCount))
}
