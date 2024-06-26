package exception

import "net/http"

type Exception struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e Exception) Error() string {
	return e.Message
}

func newException(status int, message string, code int) *Exception {
	return &Exception{
		Status:  status,
		Message: message,
		Code:    code,
	}
}

var (
	ServerError            = newException(http.StatusInternalServerError, "エラーが発生しました", 0)
	ValidationError        = newException(http.StatusBadRequest, "バリデーションエラーが発生しました", 0)
	AuthError              = newException(http.StatusUnauthorized, "認証エラーが発生しました", 0)
	InvalidRequestError    = newException(http.StatusBadRequest, "リクエストが不正です", 0)
	FailedToSigninError    = newException(http.StatusBadRequest, "サインインに失敗しました", 0)
	RecordNotFoundError    = newException(http.StatusNotFound, "レコードが見つかりませんでした", 0)
	UserAlreadyExistsError = newException(http.StatusBadRequest, "すでに使われているユーザー名です", 0)
)
