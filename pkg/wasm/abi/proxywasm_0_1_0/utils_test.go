package proxywasm_0_1_0

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeMap(t *testing.T) {
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	b := EncodeMap(m)
	assert.NotNil(t, b)

	mm := DecodeMap(b)
	assert.NotNil(t, mm)

	assert.True(t, reflect.DeepEqual(m, mm))

	var emptyMap map[string]string
	assert.Nil(t, EncodeMap(emptyMap))
}
