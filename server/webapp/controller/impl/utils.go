package impl

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"server/webapp/global/errors"
)

type ErrorResult struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Result struct {
	Result interface{} `json:"results"`
}

func httpResponse(ctx *gin.Context, data interface{}, err errors.IErrCode) {
	if err == nil {
		if reflect.Slice == reflect.TypeOf(data).Kind() {
			ctx.JSON(http.StatusOK, Result{Result: data})
		} else {
			ctx.JSON(http.StatusOK, data)
		}
	} else {
		ctx.JSON(http.StatusExpectationFailed, ErrorResult{Code: err.Code(), Error: err.Error()})
	}
}
