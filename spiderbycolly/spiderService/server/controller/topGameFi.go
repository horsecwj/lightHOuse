package controller

import (
	"help_center/spiderbycolly/spiderService/server/service"

	"github.com/kataras/iris/v12/mvc"
)

type AddressController struct {
	Service service.AddressService
}

func (c *AddressController) GetBy() *mvc.Response {

	var err error
	var address string

	address, err = c.Service.GetAddress()
	if err != nil {

		return MessageResponse(false, err.Error())
	}

	return DataResponse(true, address)
}

func (c *AddressController) GetOne() *mvc.Response {

	var err error
	var address string
	address, err = c.Service.GetAddress()
	if err != nil {
		return MessageResponse(false, err.Error())
	}
	return DataResponse(true, address)
}
