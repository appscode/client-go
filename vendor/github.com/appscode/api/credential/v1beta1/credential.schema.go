package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var credentialListRequestSchema *gojsonschema.Schema
var credentialUpdateRequestSchema *gojsonschema.Schema
var credentialCreateRequestSchema *gojsonschema.Schema
var credentialDeleteRequestSchema *gojsonschema.Schema

func init() {
	var err error
	credentialListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	credentialUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "data": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    },
    "name": {
      "type": "string"
    },
    "provider": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	credentialCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "data": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    },
    "name": {
      "type": "string"
    },
    "provider": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	credentialDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *CredentialListRequest) IsValid() (*gojsonschema.Result, error) {
	return credentialListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CredentialListRequest) IsRequest() {}

func (m *CredentialUpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return credentialUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CredentialUpdateRequest) IsRequest() {}

func (m *CredentialCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return credentialCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CredentialCreateRequest) IsRequest() {}

func (m *CredentialDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return credentialDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CredentialDeleteRequest) IsRequest() {}

func (m *CredentialListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
