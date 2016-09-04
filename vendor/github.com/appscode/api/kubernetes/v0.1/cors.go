package kubernetes

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)

	patterns = append(patterns, pattern_Clusters_Create_0)
	patterns = append(patterns, pattern_Clusters_Delete_0)
	patterns = append(patterns, pattern_Clusters_List_0)
	patterns = append(patterns, pattern_Clusters_Scale_0)
	patterns = append(patterns, pattern_Clusters_Describe_0)

	patterns = append(patterns, pattern_Clients_Nodes_0)
	patterns = append(patterns, pattern_Clients_Secrets_0)
	patterns = append(patterns, pattern_Clients_Jobs_0)
	patterns = append(patterns, pattern_Clients_Namespaces_0)
	patterns = append(patterns, pattern_Clients_Pods_0)
	patterns = append(patterns, pattern_Clients_Services_0)
	patterns = append(patterns, pattern_Clients_ReplicationControllers_0)
	patterns = append(patterns, pattern_Clients_Apps_0)
	patterns = append(patterns, pattern_Clients_App_0)
	patterns = append(patterns, pattern_Clients_AppPods_0)
	patterns = append(patterns, pattern_Clients_ConfigMaps_0)
	patterns = append(patterns, pattern_Clients_ConfigMap_0)
	patterns = append(patterns, pattern_Clients_Secret_0)

	return patterns
}
