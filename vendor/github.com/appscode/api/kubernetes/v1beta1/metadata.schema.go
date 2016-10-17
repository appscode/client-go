package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var listZonesRequestSchema *gojsonschema.Schema
var listRegionsRequestSchema *gojsonschema.Schema

func init() {
	var err error
	listZonesRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cloud_credential": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "provider": {
      "type": "string"
    },
    "region": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	listRegionsRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cloud_credential": {
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

func (m *ListZonesRequest) IsValid() (*gojsonschema.Result, error) {
	return listZonesRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ListZonesRequest) IsRequest() {}

func (m *ListRegionsRequest) IsValid() (*gojsonschema.Result, error) {
	return listRegionsRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ListRegionsRequest) IsRequest() {}

func (m *ListRegionsResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *ListZonesResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
