package controller

import (
	"bytes"
	"fmt"
	"help_center/spiderbycolly/common"
	"help_center/spiderbycolly/spiderService/server/service"
	"io/ioutil"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	//"help_center/spiderbycolly/help_center/spiderbycollyService/server/service"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

// 数据返回值
func DataResponse(success bool, data interface{}) *mvc.Response {

	resp := &Response{
		Success: success,
		Data:    data,
	}

	return &mvc.Response{
		Object: resp,
	}
}

// 消息返回值
func MessageResponse(success bool, message string) *mvc.Response {

	resp := &Response{
		Success: success,
		Message: message,
	}

	return &mvc.Response{

		Object: resp,
	}
}

func Run(addr string, release bool) error {

	app := iris.New()

	// 添加日志记录
	app.Use(loggerHandler)
	app.OnAnyErrorCode(loggerHandler)

	//v := validator.New()
	//app.Validator = v

	// 配置路由
	mvc.Configure(app.Party("/address"), address)

	return app.Run(
		iris.Addr(addr),
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

// Address路由
func address(app *mvc.Application) {

	addrService := service.NewAddressService()
	app.Register(addrService)
	app.Handle(new(AddressController))
}

// 日志记录
func loggerHandler(ctx iris.Context) {
	//all except latency to string
	var status, ip, method, path string
	var latency time.Duration
	var startTime, endTime time.Time
	startTime = time.Now()

	//no time.Since in order to format it well after
	endTime = time.Now()
	latency = endTime.Sub(startTime)

	status = strconv.Itoa(ctx.GetStatusCode())
	ip = ctx.RemoteAddr()
	method = ctx.Method()
	path = ctx.Path()

	// no new line, the framework's logger is responsible how to render each log.
	line := fmt.Sprintf("%v %4v %s %s %s", status, latency, ip, method, path)

	// 如果是POST/PUT请求，并且内容类型为JSON，则读取内容体
	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		body, err := ioutil.ReadAll(ctx.Request().Body)
		if err == nil {
			defer func() {
				_ = ctx.Request().Body.Close()
			}()

			ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))

			content := string(body)
			line = fmt.Sprintf("%s \n %s", line, content)
		}
	}

	ctx.Next()
	common.Logger.Info(line)
}
