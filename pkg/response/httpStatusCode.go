package response

const (
	ErrCodeSuccess     = 2001
	ErrCodePramInvalid = 2002
)

var msg = map[int]string{
	ErrCodeSuccess:     "success",
	ErrCodePramInvalid: "Email is invalid",
}
