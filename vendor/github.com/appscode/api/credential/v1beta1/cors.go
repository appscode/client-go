package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_CloudCredentials_List_0)
	patterns = append(patterns, pattern_CloudCredentials_Create_0)
	patterns = append(patterns, pattern_CloudCredentials_Update_0)
	patterns = append(patterns, pattern_CloudCredentials_Delete_0)
	return patterns
}
