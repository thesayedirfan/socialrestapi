package http_utils

var ErrorMessages struct {
	DbError             string `json:"database error"`
	UserIDError         string `json:"user_id not given"`
	BodyParseError      string `json:"unable to parse body"`
	UserDoesNotExist    string `json:"user does not exist"`
	PasswordParseError  string `json:"password parse error"`
	UnableToLogin       string `json:"unable to login"`
	SomethingWentWrong  string `json:"something went wrong"`
	PostIDError         string `json:"post_id not given"`
	PostDoesNotExist    string `json:"post does not exist"`
	CommentIDError      string `json:"comment id not given"`
	CommentDoesNotError string `json:"comment dors not exist"`
}
