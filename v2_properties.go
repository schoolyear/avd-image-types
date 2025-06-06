//go:generate go install github.com/atombender/go-jsonschema@v0.20.0
//go:generate go run github.com/atombender/go-jsonschema@v0.20.0 --schema-package=v2_layer_properties=github.com/schoolyear/avdimagetypes --schema-output=v2_layer_properties=v2_properties_types.go v2_properties.json

package avdimagetypes

import (
	_ "embed"
	"github.com/xeipuuv/gojsonschema"
)

//go:embed v2_properties.json
var propertiesSchema []byte
var schemaLoader = gojsonschema.NewBytesLoader(propertiesSchema)

func ValidateV2Properties(json []byte) (*gojsonschema.Result, error) {
	return gojsonschema.Validate(schemaLoader, gojsonschema.NewBytesLoader(json))
}
