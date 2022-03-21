package service

import "help_center/spiderbycolly/database"

type ImportParam struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type AddressService interface {
	GetAddress() (string, error)
	ImportBTCAddress(params ImportParam) error
	GetTxBy(hash string) string
}

type addressService struct {
}

func NewAddressService() AddressService {

	return &addressService{}
}

// 根据公链类型获取一个新地址
func (service *addressService) GetAddress() (addr string, err error) {

	res, err := database.DB().GetBybitArt()
	if err != nil {
		return "", err
	}
	return res[0].Link, nil
}

func (service *addressService) GetTxBy(hash string) string {

	return hash
}

// 导入BTC地址
func (service *addressService) ImportBTCAddress(params ImportParam) (err error) {

	return nil
}
