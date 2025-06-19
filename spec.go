//go:generate go tool oapi-codegen -config oapi-codegen.yml spec.json
//go:generate go run generate/generate_schema_names.go -spec spec.json -package avdimagetypes -output schema.gen.go

package avdimagetypes

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi-validator/config"
	valerr "github.com/pb33f/libopenapi-validator/errors"
	"github.com/pb33f/libopenapi-validator/schema_validation"
	"github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/datamodel/high/v3"
	"sync"
)

//go:embed spec.json
var spec []byte

var specDoc = sync.OnceValues(func() (libopenapi.Document, error) {
	return libopenapi.NewDocument(spec)
})

var specModel = sync.OnceValues(func() (*libopenapi.DocumentModel[v3.Document], error) {
	doc, err := specDoc()
	if err != nil {
		return nil, fmt.Errorf("failed to load spec document: %w", err)
	}

	model, modelErrors := doc.BuildV3Model()
	return model, errors.Join(modelErrors...)
})

func GetSchema(schemaName string) (*base.Schema, error) {
	model, err := specModel()
	if err != nil {
		return nil, fmt.Errorf("failed to load spec model: %w", err)
	}

	schemaProxy, exists := model.Model.Components.Schemas.Get(schemaName)
	if !exists {
		return nil, fmt.Errorf("schema %s not found", schemaName)
	}

	return schemaProxy.Schema(), nil
}

func ValidateBytes(schema *base.Schema, payload []byte) (bool, []*valerr.ValidationError) {
	validator := schema_validation.NewSchemaValidator(config.WithContentAssertions(), config.WithFormatAssertions())
	return validator.ValidateSchemaBytes(schema, payload)
}

func ValidateObject(schema *base.Schema, payload any) (bool, []*valerr.ValidationError) {
	validator := schema_validation.NewSchemaValidator(config.WithContentAssertions(), config.WithFormatAssertions())
	return validator.ValidateSchemaObject(schema, payload)
}
