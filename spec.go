//go:generate sh -c "cat spec.json | docker run -i --rm swaggest/json-cli json-cli gen-go - --package-name avdimagetypes --ptr-in-schema \"#/definitions/V2BundleProperties\" --show-const-properties > spec.gen.go"
//go:generate go run generate/generate_schema_names.go -spec spec.json -package avdimagetypes -output schema.gen.go

package avdimagetypes

import (
	_ "embed"
	"sync"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed spec.json
var spec []byte
var specLoader = gojsonschema.NewBytesLoader(spec)

var schemaLoader = sync.OnceValue(func() *gojsonschema.SchemaLoader {
	loader := gojsonschema.NewSchemaLoader()
	if err := loader.AddSchemas(specLoader); err != nil {
		panic(err)
	}
	return loader
})

var compiledDefinitionSchemas = sync.Map{}

func ValidateDefinition(definitionName string, payload []byte) (*gojsonschema.Result, error) {
	schemaGetter, _ := compiledDefinitionSchemas.LoadOrStore(definitionName, sync.OnceValues(func() (*gojsonschema.Schema, error) {
		loader := schemaLoader()
		return loader.Compile(gojsonschema.NewReferenceLoader("/v2_layer_properties#/definitions/V2LayerProperties"))
	}))

	schema, err := schemaGetter.(func() (*gojsonschema.Schema, error))()
	if err != nil {
		return nil, err
	}

	return schema.Validate(gojsonschema.NewBytesLoader(payload))
}
