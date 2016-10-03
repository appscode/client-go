package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var cloudCredentialUpdateRequestSchema *gojsonschema.Schema
var cloudCredentialDeleteRequestSchema *gojsonschema.Schema
var cloudCredentialListRequestSchema *gojsonschema.Schema
var cloudCredentialCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	cloudCredentialUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "data": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
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
	cloudCredentialDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	cloudCredentialListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	cloudCredentialCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "data": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
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
}

func (m *CloudCredentialUpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return cloudCredentialUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CloudCredentialUpdateRequest) IsRequest() {}

func (m *CloudCredentialDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return cloudCredentialDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CloudCredentialDeleteRequest) IsRequest() {}

func (m *CloudCredentialListRequest) IsValid() (*gojsonschema.Result, error) {
	return cloudCredentialListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CloudCredentialListRequest) IsRequest() {}

func (m *CloudCredentialCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return cloudCredentialCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CloudCredentialCreateRequest) IsRequest() {}

func (m *CloudCredentialListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
