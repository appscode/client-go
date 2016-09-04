package operation

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var dataBucketDeleteRequestSchema *gojsonschema.Schema
var describeRequestSchema *gojsonschema.Schema
var namespaceAdminTaskRequestSchema *gojsonschema.Schema

func init() {
	var err error
	dataBucketDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "data_type": {
      "type": "string"
    },
    "namespace": {
      "type": "string"
    },
    "prefix": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	describeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "phid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	namespaceAdminTaskRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "namespace": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *DataBucketDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return dataBucketDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DataBucketDeleteRequest) IsRequest() {}

func (m *DescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return describeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DescribeRequest) IsRequest() {}

func (m *NamespaceAdminTaskRequest) IsValid() (*gojsonschema.Result, error) {
	return namespaceAdminTaskRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *NamespaceAdminTaskRequest) IsRequest() {}

func (m *DescribeResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
