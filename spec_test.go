package avdimagetypes

import (
	"testing"
)

func BenchmarkValidate(b *testing.B) {
	for b.Loop() {
		ValidateDefinition("V2BundleProperties", []byte(`{}`))
	}
}
