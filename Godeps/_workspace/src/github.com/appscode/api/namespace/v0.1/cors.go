package namespace

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Namespace_Create_0,
		pattern_Namespace_Get_0,
		pattern_Namespace_IsAvailable_0,
	}
}
