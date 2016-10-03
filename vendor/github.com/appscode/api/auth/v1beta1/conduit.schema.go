package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var conduitRequestSchema *gojsonschema.Schema

func init() {
	var err error
	conduitRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *ConduitRequest) IsValid() (*gojsonschema.Result, error) {
	return conduitRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ConduitRequest) IsRequest() {}

func (m *ConduitUsersResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *ConduitWhoAmIResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
