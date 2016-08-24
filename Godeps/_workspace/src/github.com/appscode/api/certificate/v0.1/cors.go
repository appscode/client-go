package certificate

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	patterns := make([]runtime.Pattern, 0)
	patterns = append(patterns, pattern_Certificates_Create_0)
	patterns = append(patterns, pattern_Certificates_Delete_0)
	patterns = append(patterns, pattern_Certificates_Deploy_0)
	patterns = append(patterns, pattern_Certificates_Describe_0)
	patterns = append(patterns, pattern_Certificates_Import_0)
	patterns = append(patterns, pattern_Certificates_Renew_0)
	patterns = append(patterns, pattern_Certificates_Revoke_0)
	patterns = append(patterns, pattern_Certificates_List_0)
	patterns = append(patterns, pattern_Certificates_Import_0)
	patterns = append(patterns, pattern_Certificates_Renew_0)
	patterns = append(patterns, pattern_Certificates_Revoke_0)
	return patterns
}
