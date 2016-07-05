package glusterfs

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Clusters_List_0)
	patterns = append(patterns, pattern_Clusters_Describe_0)
	patterns = append(patterns, pattern_Clusters_Create_0)
	patterns = append(patterns, pattern_Clusters_Delete_0)

	patterns = append(patterns, pattern_Volumes_List_0)
	patterns = append(patterns, pattern_Volumes_Create_0)
	patterns = append(patterns, pattern_Volumes_Delete_0)
	return patterns
}
