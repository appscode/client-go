package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Servers_Create_0)
	patterns = append(patterns, pattern_Servers_Delete_0)
	patterns = append(patterns, pattern_Clients_Reconfigure_0)
	return patterns
}
