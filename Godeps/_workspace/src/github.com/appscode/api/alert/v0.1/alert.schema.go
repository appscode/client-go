package alert

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var deleteRequestSchema *gojsonschema.Schema
var createRequestSchema *gojsonschema.Schema
var alertListRequestSchema *gojsonschema.Schema
var updateRequestSchema *gojsonschema.Schema

func init() {
	var err error
	deleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	createRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "alertCheckInfluxData": {
      "properties": {
        "critical_condition": {
          "type": "string"
        },
        "kubernetes_namespace": {
          "type": "string"
        },
        "object_name": {
          "type": "string"
        },
        "query": {
          "additionalProperties": {
            "type": "string"
          },
          "type": "object"
        },
        "r": {
          "type": "string"
        },
        "warning_condition": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "alertCheckKubernetes": {
      "properties": {
        "count": {
          "type": "integer"
        },
        "selector": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "alertCommandParam": {
      "properties": {
        "check_influx_data": {
          "$ref": "#/definitions/alertCheckInfluxData"
        },
        "check_kubernetes": {
          "$ref": "#/definitions/alertCheckKubernetes"
        }
      },
      "type": "object"
    },
    "alertIcingaParam": {
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
    "alertNotifierParam": {
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
    "cluster": {
      "type": "string"
    },
    "command_param": {
      "$ref": "#/definitions/alertCommandParam"
    },
    "icinga_param": {
      "$ref": "#/definitions/alertIcingaParam"
    },
    "icinga_service": {
      "type": "string"
    },
    "notifier_param": {
      "items": {
        "$ref": "#/definitions/alertNotifierParam"
      },
      "type": "array"
    },
    "plugin": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	alertListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "kubernetes_namespace": {
      "type": "string"
    },
    "object_name": {
      "type": "string"
    },
    "plugin": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	updateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "alertCheckInfluxData": {
      "properties": {
        "critical_condition": {
          "type": "string"
        },
        "kubernetes_namespace": {
          "type": "string"
        },
        "object_name": {
          "type": "string"
        },
        "query": {
          "additionalProperties": {
            "type": "string"
          },
          "type": "object"
        },
        "r": {
          "type": "string"
        },
        "warning_condition": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "alertCheckKubernetes": {
      "properties": {
        "count": {
          "type": "integer"
        },
        "selector": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      },
      "type": "object"
    },
    "alertCommandParam": {
      "properties": {
        "check_influx_data": {
          "$ref": "#/definitions/alertCheckInfluxData"
        },
        "check_kubernetes": {
          "$ref": "#/definitions/alertCheckKubernetes"
        }
      },
      "type": "object"
    },
    "alertIcingaParam": {
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
    "alertNotifierParam": {
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
    "command_param": {
      "$ref": "#/definitions/alertCommandParam"
    },
    "icinga_param": {
      "$ref": "#/definitions/alertIcingaParam"
    },
    "notifier_param": {
      "items": {
        "$ref": "#/definitions/alertNotifierParam"
      },
      "type": "array"
    },
    "phid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *DeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return deleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DeleteRequest) IsRequest() {}

func (m *CreateRequest) IsValid() (*gojsonschema.Result, error) {
	return createRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CreateRequest) IsRequest() {}

func (m *AlertListRequest) IsValid() (*gojsonschema.Result, error) {
	return alertListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AlertListRequest) IsRequest() {}

func (m *UpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return updateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *UpdateRequest) IsRequest() {}

func (m *AlertListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
