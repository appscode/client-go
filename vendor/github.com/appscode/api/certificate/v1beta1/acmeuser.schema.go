package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var acmeUserRegisterRequestSchema *gojsonschema.Schema

func init() {
	var err error
	acmeUserRegisterRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "email": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *AcmeUserRegisterRequest) IsValid() (*gojsonschema.Result, error) {
	return acmeUserRegisterRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AcmeUserRegisterRequest) IsRequest() {}

func (m *AcmeUserRegisterResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
