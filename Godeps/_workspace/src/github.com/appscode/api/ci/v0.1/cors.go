package ci

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Agents_List_0)
	patterns = append(patterns, pattern_Agents_Create_0)
	patterns = append(patterns, pattern_Agents_Delete_0)
	patterns = append(patterns, pattern_Agents_Describe_0)
	patterns = append(patterns, pattern_Agents_Restart_0)
	return patterns
}
