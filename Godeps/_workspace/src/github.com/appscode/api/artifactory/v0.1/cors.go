package artifactory

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)

	patterns = append(patterns, pattern_Artifacts_Search_0)
	patterns = append(patterns, pattern_Artifacts_List_0)
	patterns = append(patterns, pattern_Versions_List_0)
	patterns = append(patterns, pattern_Versions_Describe_0)

	return patterns
}
