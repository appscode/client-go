package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Databases_List_0)
	patterns = append(patterns, pattern_Databases_Create_0)
	patterns = append(patterns, pattern_Databases_Scale_0)
	patterns = append(patterns, pattern_Databases_Update_0)
	patterns = append(patterns, pattern_Databases_Describe_0)
	patterns = append(patterns, pattern_Databases_Delete_0)
	patterns = append(patterns, pattern_Snapshots_List_0)
	patterns = append(patterns, pattern_Snapshots_BackupSchedule_0)
	patterns = append(patterns, pattern_Snapshots_BackupUnschedule_0)
	patterns = append(patterns, pattern_Snapshots_Restore_0)
	return patterns
}
