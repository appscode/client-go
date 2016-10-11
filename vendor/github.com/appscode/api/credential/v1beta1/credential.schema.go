package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var credentialDeleteRequestSchema *gojsonschema.Schema
var credentialListRequestSchema *gojsonschema.Schema
var credentialUpdateRequestSchema *gojsonschema.Schema
var credentialCreateRequestSchema *gojsonschema.Schema
var credentialIsAuthorizedRequestSchema *gojsonschema.Schema

func init() {
	var err error
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
	credentialIsAuthorizedRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *CredentialDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return credentialDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CredentialDeleteRequest) IsRequest() {}

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

func (m *CredentialIsAuthorizedRequest) IsValid() (*gojsonschema.Result, error) {
	return credentialIsAuthorizedRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CredentialIsAuthorizedRequest) IsRequest() {}

func (m *CredentialIsAuthorizedResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *CredentialListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
