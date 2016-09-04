package alert

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Alerts_List_0)
	patterns = append(patterns, pattern_Alerts_Create_0)
	patterns = append(patterns, pattern_Alerts_Describe_0)
	patterns = append(patterns, pattern_Alerts_Update_0)
	patterns = append(patterns, pattern_Alerts_Sync_0)
	patterns = append(patterns, pattern_Alerts_Delete_0)

	patterns = append(patterns, pattern_Incidents_List_0)
	patterns = append(patterns, pattern_Incidents_Describe_0)
	patterns = append(patterns, pattern_Incidents_Notify_0)
	patterns = append(patterns, pattern_Incidents_CreateEvent_0)
	return patterns
}
