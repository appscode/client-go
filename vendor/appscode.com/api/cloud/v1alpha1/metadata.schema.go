package v1alpha1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var zoneListRequestSchema *gojsonschema.Schema
var bucketListRequestSchema *gojsonschema.Schema
var regionListRequestSchema *gojsonschema.Schema

func init() {
	var err error
	zoneListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cloud_credential": {
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
	bucketListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cloud_credential": {
      "type": "string"
    },
    "cluster_uid": {
      "type": "string"
    },
    "gce_project": {
      "type": "string"
    },
    "provider": {
      "type": "string"
    },
    "secret_name": {
      "type": "string"
    },
    "secret_namespace": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	regionListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cloud_credential": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *ZoneListRequest) Valid() (*gojsonschema.Result, error) {
	return zoneListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ZoneListRequest) IsRequest() {}

func (m *BucketListRequest) Valid() (*gojsonschema.Result, error) {
	return bucketListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *BucketListRequest) IsRequest() {}

func (m *RegionListRequest) Valid() (*gojsonschema.Result, error) {
	return regionListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *RegionListRequest) IsRequest() {}

