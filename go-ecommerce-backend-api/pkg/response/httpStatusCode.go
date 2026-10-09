package response

const (
	ErrCodeSuccess = 200
	ErrCodeParamInvalid = 400
)

var msg = map[int]string{
	ErrCodeSuccess: "Success",
	ErrCodeParamInvalid: "Email is Invalid",
}