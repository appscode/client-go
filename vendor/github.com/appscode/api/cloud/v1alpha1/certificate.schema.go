package v1alpha1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var certificateDeleteRequestSchema *gojsonschema.Schema
var certificateDescribeRequestSchema *gojsonschema.Schema
var certificateListRequestSchema *gojsonschema.Schema
var certificateLoadRequestSchema *gojsonschema.Schema
var certificateDeployRequestSchema *gojsonschema.Schema

func init() {
	var err error
	certificateDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	certificateDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	certificateListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	certificateLoadRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cert_data": {
      "type": "string"
    },
    "key_data": {
      "type": "string"
    },
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	certificateDeployRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    },
    "secret_name": {
      "type": "string"
    },
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *CertificateDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDeleteRequest) IsRequest() {}

func (m *CertificateDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDescribeRequest) IsRequest() {}

func (m *CertificateListRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateListRequest) IsRequest() {}

func (m *CertificateLoadRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateLoadRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateLoadRequest) IsRequest() {}

func (m *CertificateDeployRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateDeployRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDeployRequest) IsRequest() {}

