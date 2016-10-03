package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var alertDeleteRequestSchema *gojsonschema.Schema
var alertUpdateRequestSchema *gojsonschema.Schema
var alertListRequestSchema *gojsonschema.Schema
var alertCreateRequestSchema *gojsonschema.Schema
var alertDescribeRequestSchema *gojsonschema.Schema
var alertSyncRequestSchema *gojsonschema.Schema

func init() {
	var err error
	alertDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "phid": {
      "type": "string"
    }
  },
  "title": "Next Id: 2",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	alertUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "v1beta1IcingaParam": {
      "properties": {
        "alert_interval_sec": {
          "type": "integer"
        },
        "check_interval_sec": {
          "type": "integer"
        }
      },
      "type": "object"
    },
    "v1beta1NotifierParam": {
      "properties": {
        "method": {
          "type": "string"
        },
        "state": {
          "type": "string"
        },
        "user_uid": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "icinga_param": {
      "$ref": "#/definitions/v1beta1IcingaParam"
    },
    "notifier_param": {
      "items": {
        "$ref": "#/definitions/v1beta1NotifierParam"
      },
      "type": "array"
    },
    "phid": {
      "type": "string"
    },
    "vars": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    }
  },
  "title": "Next Id: 6",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	alertListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "kubernetes_cluster": {
      "type": "string"
    },
    "kubernetes_namespace": {
      "type": "string"
    },
    "kubernetes_objectName": {
      "type": "string"
    },
    "kubernetes_objectType": {
      "type": "string"
    }
  },
  "title": "Next Id: 5",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	alertCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "v1beta1IcingaParam": {
      "properties": {
        "alert_interval_sec": {
          "type": "integer"
        },
        "check_interval_sec": {
          "type": "integer"
        }
      },
      "type": "object"
    },
    "v1beta1NotifierParam": {
      "properties": {
        "method": {
          "type": "string"
        },
        "state": {
          "type": "string"
        },
        "user_uid": {
          "type": "string"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "check_command": {
      "type": "string"
    },
    "icinga_param": {
      "$ref": "#/definitions/v1beta1IcingaParam"
    },
    "icinga_service": {
      "type": "string"
    },
    "kubernetes_cluster": {
      "type": "string"
    },
    "kubernetes_namespace": {
      "type": "string"
    },
    "kubernetes_objectName": {
      "type": "string"
    },
    "kubernetes_objectType": {
      "type": "string"
    },
    "notifier_param": {
      "items": {
        "$ref": "#/definitions/v1beta1NotifierParam"
      },
      "type": "array"
    },
    "vars": {
      "additionalProperties": {
        "type": "string"
      },
      "type": "object"
    }
  },
  "title": "Next Id: 12",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	alertDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "phid": {
      "type": "string"
    }
  },
  "title": "Next Id: 2",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	alertSyncRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "AlertSyncRequestPodAncestor": {
      "properties": {
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      },
      "title": "Next Id: 3",
      "type": "object"
    }
  },
  "properties": {
    "kubernetes_cluster": {
      "type": "string"
    },
    "kubernetes_namespace": {
      "type": "string"
    },
    "kubernetes_objectName": {
      "type": "string"
    },
    "kubernetes_objectType": {
      "type": "string"
    },
    "pod_ancestor": {
      "items": {
        "$ref": "#/definitions/AlertSyncRequestPodAncestor"
      },
      "type": "array"
    }
  },
  "title": "Next Id: 6",
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *AlertDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return alertDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AlertDeleteRequest) IsRequest() {}

func (m *AlertUpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return alertUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AlertUpdateRequest) IsRequest() {}

func (m *AlertListRequest) IsValid() (*gojsonschema.Result, error) {
	return alertListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AlertListRequest) IsRequest() {}

func (m *AlertCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return alertCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AlertCreateRequest) IsRequest() {}

func (m *AlertDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return alertDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AlertDescribeRequest) IsRequest() {}

func (m *AlertSyncRequest) IsValid() (*gojsonschema.Result, error) {
	return alertSyncRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AlertSyncRequest) IsRequest() {}

func (m *AlertDescribeResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *AlertListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
