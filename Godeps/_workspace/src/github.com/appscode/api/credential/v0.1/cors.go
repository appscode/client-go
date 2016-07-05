package credential

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_CloudCredential_List_0)
	patterns = append(patterns, pattern_CloudCredential_Create_0)
	patterns = append(patterns, pattern_CloudCredential_Update_0)
	patterns = append(patterns, pattern_CloudCredential_Delete_0)
	return patterns
}
