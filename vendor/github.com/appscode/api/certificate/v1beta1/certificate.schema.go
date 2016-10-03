package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var certificateImportRequestSchema *gojsonschema.Schema
var certificateRenewRequestSchema *gojsonschema.Schema
var certificateListRequestSchema *gojsonschema.Schema
var certificateDeleteRequestSchema *gojsonschema.Schema
var certificateRevokeRequestSchema *gojsonschema.Schema
var certificateDescribeRequestSchema *gojsonschema.Schema
var certificateCreateRequestSchema *gojsonschema.Schema
var certificateDeployRequestSchema *gojsonschema.Schema

func init() {
	var err error
	certificateImportRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cert_data": {
      "type": "string"
    },
    "key_data": {
      "type": "string"
    },
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
	certificateRenewRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "account_phid": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
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
	certificateListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
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
	certificateRevokeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "account_phid": {
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
	certificateCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "v1beta1SubjectInfo": {
      "description": "A Name contains the SubjectInfo fields.",
      "properties": {
        "C": {
          "type": "string"
        },
        "L": {
          "type": "string"
        },
        "O": {
          "type": "string"
        },
        "OU": {
          "type": "string"
        },
        "ST": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "account_phid": {
      "type": "string"
    },
    "bundle": {
      "type": "boolean"
    },
    "common_name": {
      "type": "string"
    },
    "issue_private_cert": {
      "type": "boolean"
    },
    "key_data": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "san": {
      "items": {
        "type": "string"
      },
      "type": "array"
    },
    "subject_info": {
      "$ref": "#/definitions/v1beta1SubjectInfo"
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
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "namespace": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "secret_name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
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

func (m *CertificateImportRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateImportRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateImportRequest) IsRequest() {}

func (m *CertificateRenewRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateRenewRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateRenewRequest) IsRequest() {}

func (m *CertificateListRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateListRequest) IsRequest() {}

func (m *CertificateDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDeleteRequest) IsRequest() {}

func (m *CertificateRevokeRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateRevokeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateRevokeRequest) IsRequest() {}

func (m *CertificateDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDescribeRequest) IsRequest() {}

func (m *CertificateCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateCreateRequest) IsRequest() {}

func (m *CertificateDeployRequest) IsValid() (*gojsonschema.Result, error) {
	return certificateDeployRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CertificateDeployRequest) IsRequest() {}

func (m *CertificateListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *CertificateDescribeResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *CertificateRenewResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *CertificateCreateResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
