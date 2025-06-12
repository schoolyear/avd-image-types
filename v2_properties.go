//go:generate sh -c "cat v2_properties.json | docker run -i --rm swaggest/json-cli json-cli gen-go - --package-name avdimagetypes --ptr-in-schema \"#/definitions/v2_layer_properties\" --show-const-properties > v2_properties_types.go"

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
