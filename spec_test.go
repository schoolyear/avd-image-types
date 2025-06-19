package avdimagetypes_test

import (
	avdimagetypes "github.com/schoolyear/avd-image-types"
	"github.com/stretchr/testify/require"
	"testing"
)

func BenchmarkValidateBytes(b *testing.B) {
	payload := []byte(`{
"version": "v2",
"name": "hello-world",
"author": {
"name": "Your mom"
},
"platform_version": "2"
}`)

	schema, err := avdimagetypes.GetSchema(avdimagetypes.V2LayerSchemaName)
	require.NoError(b, err)

	for b.Loop() {
		avdimagetypes.ValidateBytes(schema, payload)
	}
}
