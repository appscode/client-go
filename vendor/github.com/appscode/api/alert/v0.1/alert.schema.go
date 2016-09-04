package alert

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
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	alertUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "alertCheckInfluxData": {
      "properties": {
        "critical_condition": {
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
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	alertCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "alertCheckInfluxData": {
      "properties": {
        "critical_condition": {
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
    "command_param": {
      "$ref": "#/definitions/alertCommandParam"
    },
    "icinga_param": {
      "$ref": "#/definitions/alertIcingaParam"
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
	alertDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	alertSyncRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
