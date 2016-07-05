package mailinglist

import "github.com/gengo/grpc-gateway/runtime"

func Patterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_MailingList_SendEmail_0,
		pattern_MailingList_Subscribe_0,
	}
}
