package backup

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Server_Create_0)
	patterns = append(patterns, pattern_Server_Delete_0)
	patterns = append(patterns, pattern_Client_ReConfigure_0)
	return patterns
}
