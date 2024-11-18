package logic

import (
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

func MustBuildErrData(code int, msg string) string {
	data, err := json.Marshal(errors.CodeMsg{
		Code: code,
		Msg:  msg,
	})
	logx.Must(err)
	return string(data)
}
