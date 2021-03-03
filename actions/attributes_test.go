package actions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPluginBaseArgument(t *testing.T) {
	assert.False(t, IsCommonAttribute("_"))
	assert.True(t, IsCommonAttribute(ArgCount))
}
