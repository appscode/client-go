package backup

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var clientReConfigureRequestSchema *gojsonschema.Schema

func init() {
	var err error
	clientReConfigureRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster_name": {
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

func (m *ClientReConfigureRequest) IsValid() (*gojsonschema.Result, error) {
	return clientReConfigureRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClientReConfigureRequest) IsRequest() {}

