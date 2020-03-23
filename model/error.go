package model

//Error 错误模型
type Error struct {
	Msg string
}

func (err *Error) Error() string {
	return err.Msg
}
