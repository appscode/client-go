package db

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var deleteRequestSchema *gojsonschema.Schema
var describeRequestSchema *gojsonschema.Schema
var addStandbyRequestSchema *gojsonschema.Schema
var listRequestSchema *gojsonschema.Schema
var backupScheduleRequestSchema *gojsonschema.Schema
var updateRequestSchema *gojsonschema.Schema
var backupUnscheduleRequestSchema *gojsonschema.Schema
var createRequestSchema *gojsonschema.Schema
var snapshotListRequestSchema *gojsonschema.Schema
var restoreRequestSchema *gojsonschema.Schema

func init() {
	var err error
	deleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
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
	describeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
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
	addStandbyRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "config": {
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
	listRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "status": {
      "items": {
        "type": "string"
      },
      "title": "List of status to get the agent filterd on the status\nvalues in\n  PENDING\n  FAILED\n  READY\n  DELETED",
      "type": "array"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	backupScheduleRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "auth_secret_name": {
      "oneOf": [
        {
          "maxLength": 0
        },
        {
          "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$"
        }
      ],
      "type": "string"
    },
    "bucket_name": {
      "oneOf": [
        {
          "maxLength": 0
        },
        {
          "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$"
        }
      ],
      "type": "string"
    },
    "cluster": {
      "type": "string"
    },
    "credential": {
      "type": "string"
    },
    "dump": {
      "type": "boolean"
    },
    "force": {
      "type": "boolean"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "region": {
      "type": "string"
    },
    "schedule_cron_expr": {
      "type": "string"
    },
    "snapshot_name": {
      "oneOf": [
        {
          "maxLength": 0
        },
        {
          "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$"
        }
      ],
      "type": "string"
    },
    "type": {
      "type": "string"
    },
    "wal": {
      "type": "boolean"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	updateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
      "type": "string"
    },
    "do_not_delete": {
      "type": "boolean"
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
	backupUnscheduleRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
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
	createRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "auth_secret_name": {
      "oneOf": [
        {
          "maxLength": 0
        },
        {
          "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$"
        }
      ],
      "type": "string"
    },
    "bucket_name": {
      "oneOf": [
        {
          "maxLength": 0
        },
        {
          "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$"
        }
      ],
      "type": "string"
    },
    "cluster": {
      "type": "string"
    },
    "credential": {
      "type": "string"
    },
    "do_not_delete": {
      "type": "boolean"
    },
    "node": {
      "type": "string"
    },
    "pv": {
      "type": "string"
    },
    "pv_size": {
      "type": "string"
    },
    "region": {
      "type": "string"
    },
    "service_name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "sku": {
      "type": "string"
    },
    "type": {
      "type": "string"
    },
    "version": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	snapshotListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "cluster": {
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
	restoreRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "auth_secret_name": {
      "oneOf": [
        {
          "maxLength": 0
        },
        {
          "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$"
        }
      ],
      "type": "string"
    },
    "bucket_name": {
      "oneOf": [
        {
          "maxLength": 0
        },
        {
          "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$"
        }
      ],
      "type": "string"
    },
    "cluster": {
      "type": "string"
    },
    "credential": {
      "type": "string"
    },
    "do_not_delete": {
      "type": "boolean"
    },
    "dump": {
      "type": "boolean"
    },
    "force": {
      "type": "boolean"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "node": {
      "type": "string"
    },
    "pv": {
      "type": "string"
    },
    "pv_size": {
      "type": "string"
    },
    "region": {
      "type": "string"
    },
    "service_name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "sku": {
      "type": "string"
    },
    "snapshot_phid": {
      "type": "string"
    },
    "type": {
      "type": "string"
    },
    "version": {
      "type": "string"
    },
    "wal": {
      "type": "boolean"
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

func (m *DescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return describeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DescribeRequest) IsRequest() {}

func (m *AddStandbyRequest) IsValid() (*gojsonschema.Result, error) {
	return addStandbyRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *AddStandbyRequest) IsRequest() {}

func (m *ListRequest) IsValid() (*gojsonschema.Result, error) {
	return listRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ListRequest) IsRequest() {}

func (m *BackupScheduleRequest) IsValid() (*gojsonschema.Result, error) {
	return backupScheduleRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *BackupScheduleRequest) IsRequest() {}

func (m *UpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return updateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *UpdateRequest) IsRequest() {}

func (m *BackupUnscheduleRequest) IsValid() (*gojsonschema.Result, error) {
	return backupUnscheduleRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *BackupUnscheduleRequest) IsRequest() {}

func (m *CreateRequest) IsValid() (*gojsonschema.Result, error) {
	return createRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *CreateRequest) IsRequest() {}

func (m *SnapshotListRequest) IsValid() (*gojsonschema.Result, error) {
	return snapshotListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *SnapshotListRequest) IsRequest() {}

func (m *RestoreRequest) IsValid() (*gojsonschema.Result, error) {
	return restoreRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *RestoreRequest) IsRequest() {}

func (m *ListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *SnapshotListResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *DescribeResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
