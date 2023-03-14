package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrMsg struct {
	Message interface{} `json:"message"`
}
type ResponseList struct {
	Pages Pages       `json:"pages"`
	List  interface{} `json:"list"`
}

func OK(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, data)
	c.Abort()
}

func Error(msg interface{}, code int, c *gin.Context) {
	c.JSON(code, ErrorMsg(msg))
	c.Abort()
}

func ErrorMsg(msg interface{}) *ErrMsg {
	return &ErrMsg{
		Message: msg,
	}
}
