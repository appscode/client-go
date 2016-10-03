package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_LoadBalancers_Create_0)
	patterns = append(patterns, pattern_LoadBalancers_Delete_0)
	patterns = append(patterns, pattern_LoadBalancers_Describe_0)
	patterns = append(patterns, pattern_LoadBalancers_List_0)
	patterns = append(patterns, pattern_LoadBalancers_Update_0)
	return patterns
}
