package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Teams_Create_0,
		pattern_Teams_Get_0,
		pattern_Teams_IsAvailable_0,
	}
}
