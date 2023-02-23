package attr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAttributes(t *testing.T) {
	attrs := New()

	attrs.Set(Label("test"), Shape("box"))

	assert.Equal(t, "test", attrs["label"])
	assert.Equal(t, "box", attrs["shape"])
}
