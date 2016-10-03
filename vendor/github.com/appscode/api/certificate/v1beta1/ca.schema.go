package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var cACreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	cACreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "csr": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    }
  },
  "title": "Use specific requests for protos",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *CACreateRequest) IsValid() (*gojsonschema.Result, error) {
	return cACreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CACreateRequest) IsRequest() {}

func (m *CACreateResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
