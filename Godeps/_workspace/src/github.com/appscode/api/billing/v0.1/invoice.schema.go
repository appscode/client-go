package billing

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var invoiceCreateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	invoiceCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "time_unit": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *InvoiceCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return invoiceCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *InvoiceCreateRequest) IsRequest() {}

func (m *InvoiceCreateResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
