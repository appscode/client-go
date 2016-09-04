package glusterfs

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var volumeListRequestSchema *gojsonschema.Schema
var volumeCreateRequestSchema *gojsonschema.Schema
var volumeDeleteRequestSchema *gojsonschema.Schema

func init() {
	var err error
	volumeListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "glusterfs_cluster": {
      "type": "string"
    },
    "kube_cluster": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	volumeCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "enable_bacula": {
      "type": "integer"
    },
    "glusterfs_cluster": {
      "type": "string"
    },
    "kube_cluster": {
      "type": "string"
    },
    "volume": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	volumeDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "glusterfs_cluster": {
      "type": "string"
    },
    "kube_cluster": {
      "type": "string"
    },
    "volume": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *VolumeListRequest) IsValid() (*gojsonschema.Result, error) {
	return volumeListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *VolumeListRequest) IsRequest() {}

func (m *VolumeCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return volumeCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *VolumeCreateRequest) IsRequest() {}

func (m *VolumeDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return volumeDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *VolumeDeleteRequest) IsRequest() {}

func (m *VolumeListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
