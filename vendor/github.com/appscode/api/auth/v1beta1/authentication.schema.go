package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var loginRequestSchema *gojsonschema.Schema
var logoutRequestSchema *gojsonschema.Schema

func init() {
	var err error
	loginRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "issue_token": {
      "type": "boolean"
    },
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "password": {
      "type": "string"
    },
    "token": {
      "type": "string"
    },
    "username": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	logoutRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "token": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *LoginRequest) IsValid() (*gojsonschema.Result, error) {
	return loginRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *LoginRequest) IsRequest() {}

func (m *LogoutRequest) IsValid() (*gojsonschema.Result, error) {
	return logoutRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *LogoutRequest) IsRequest() {}

func (m *CSRFTokenResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *LogoutResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *LoginResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
