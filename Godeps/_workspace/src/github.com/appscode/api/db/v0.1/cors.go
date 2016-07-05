package db

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Databases_List_0)
	patterns = append(patterns, pattern_Databases_Create_0)
	patterns = append(patterns, pattern_Databases_AddStandby_0)
	patterns = append(patterns, pattern_Databases_Describe_0)
	patterns = append(patterns, pattern_Databases_Delete_0)
	patterns = append(patterns, pattern_Databases_BackupSchedule_0)
	patterns = append(patterns, pattern_Databases_BackupUnschedule_0)
	patterns = append(patterns, pattern_Databases_Restore_0)
	patterns = append(patterns, pattern_Databases_SnapshotList_0)
	return patterns
}
