package response

const (
	ErrCodeSuccess      = 20001 // success
	ErrCodeParamInvalid = 20003 // email invalid
	ErrInvalidToken     = 30001 // token is invalid

	//Register Code
	ErrCodeUserHasExisted = 50001 // user has existed

)

// message
var msg = map[int]string{
	ErrCodeSuccess:        "success",
	ErrCodeParamInvalid:   "email invalid",
	ErrInvalidToken:       "token is invalid",
	ErrCodeUserHasExisted: "user has existed",
}
