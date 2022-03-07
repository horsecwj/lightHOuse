package service

import (
	"help_center/internal/biz"
	"help_center/internal/data"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// admUploadImg doc
// @Tags img-图片
// @Summary 添加图片
// @Param token header string true "token"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/upload_img [POST]
func admUploadImg(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusOK, biz.BaseJson{Code: 0, Data: err.Error()})
	}
	contentType := file.Header.Get("Content-Type")
	if !strings.Contains(contentType, "video") && !strings.Contains(contentType, "image") {
		return c.JSON(http.StatusOK, biz.BaseJson{Code: 0, Data: "请选择图片或视频文件"})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusOK, biz.BaseJson{Code: 0, Data: "请选择图片或视频文件"})
	}
	msg := biz.UploadImage(src, file.Filename)
	return c.JSON(http.StatusOK, &msg)
}

func admFitImg(c echo.Context) error {
	request := c.Request()
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusOK, biz.BaseJson{Code: 0, Data: err.Error()})
	}
	contentType := file.Header.Get("Content-Type")
	if !strings.Contains(contentType, "image") {
		return c.JSON(http.StatusOK, biz.BaseJson{Code: 0, Data: "请选择图片文件"})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusOK, biz.BaseJson{Code: 0, Data: "请选择图片文件"})
	}
	msg := biz.FitImage(src, file.Filename, request)
	return c.JSON(http.StatusOK, &msg)
}

// admAddBanner doc
// @Tags Banner-横幅
// @Summary 添加横幅
// @Param token header string true "token"
/// @Param body body data.Banner true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_banner [POST]
func admAddBanner(c echo.Context) error {
	d := new(data.Banner)
	request := c.Request()
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.AddBanner(d, request)
	return c.JSON(http.StatusOK, &msg)
}

// admDelBanner doc
// @Tags Banner-横幅
// @Summary 删除横幅
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_banner [POST]
func admDelBanner(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelBanner(d)
	return c.JSON(http.StatusOK, &msg)
}

// admModBanner doc
// @Tags Banner-文章
// @Summary 修改横幅
// @Param token header string true "token"
// @Param body body data.Banner true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/mod_banner [POST]
func admModBanner(c echo.Context) error {
	d := new(data.Banner)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.ModBanner(d)
	return c.JSON(http.StatusOK, &msg)
}

// admGetBanner doc
// @Tags Banner-查看横幅
// @Summary 查看横幅
// @Param token header string true "token"
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Banner} "返回数据"
// @Router /adm/get_banner [GET]
func admGetBanner(c echo.Context) error {
	msg := biz.GetBanner()
	return c.JSON(http.StatusOK, &msg)
}

// admAddArticle doc
// @Tags Article-文章
// @Summary 添加文章
// @Param token header string true "token"
// @Param body body data.Article true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_article [POST]
func admAddArticle(c echo.Context) error {
	d := new(data.Article)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	//d.Status = 1
	msg := biz.AddArticle(d)
	return c.JSON(http.StatusOK, &msg)
}

// admDelArticle doc
// @Tags Article-文章
// @Summary 删除文章
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_article [POST]
func admDelArticle(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelArticle(d)
	return c.JSON(http.StatusOK, &msg)
}

// admModArticle doc
// @Tags Article-文章
// @Summary 修改文章
// @Param token header string true "token"
// @Param body body data.Category true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/mod_article [POST]
func admModArticle(c echo.Context) error {
	d := new(data.Article)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.ModArticle(d)
	return c.JSON(http.StatusOK, &msg)
}

// admGetArticle doc
// @Tags Article-发布文章
// @Summary 查询发布文章
// @Param token header string true "token"
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.Article} "返回数据"
// @Router /adm/get_article [GET]
func admGetArticle(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetArticle(d, true)
	return c.JSON(http.StatusOK, &msg)
}

// admMatchArticle
// @Tags UseApi
// @Summary title关键字查询文章(至多返回30条数据)
// @Param sub_str query string true "匹配数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.Article} "返回数据"
// @Router /api/match_article [GET]
func admMatchArticle(c echo.Context) error {
	subStr := c.QueryParam("sub_str")
	msg := biz.MatchArticle(subStr)
	return c.JSON(http.StatusOK, &msg)
}

// admGetGame doc
// @Tags Game-游戏信息
// @Summary 查询游戏信息
// @Param token header string true "token"
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.Game} "返回数据"
// @Router /adm/get_game [GET]
func admGetGame(c echo.Context) error {
	d := new(data.GameQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetGame(d, true)
	return c.JSON(http.StatusOK, &msg)
}

// admAddGame doc
// @Tags Game-游戏信息
// @Summary 添加游戏信息
// @Param token header string true "token"
// @Param body body data.Game true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_game [POST]
func admAddGame(c echo.Context) error {
	d := new(data.Game)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	//d.Status = 1
	msg := biz.AddGame(d)
	return c.JSON(http.StatusOK, &msg)
}

// admDelGame doc
// @Tags Game-游戏信息
// @Summary 删除游戏信息
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_game [POST]
func admDelGame(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelGame(d)
	return c.JSON(http.StatusOK, &msg)
}

// admModGame doc
// @Tags Game-游戏信息
// @Summary 修改游戏信息
// @Param token header string true "token"
// @Param body body data.Game true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/mod_game [POST]
func admModGame(c echo.Context) error {
	d := new(data.Game)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.ModGame(d)
	return c.JSON(http.StatusOK, &msg)
}

// useMatchGame
// @Tags UseApi
// @Summary GameName关键字查询游戏名(至多返回30条数据)
// @Param sub_str query string true "匹配数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.game} "返回数据"
// @Router /api/match_game [GET]
func admMatchGame(c echo.Context) error {
	subStr := c.QueryParam("sub_str")
	msg := biz.MatchGame(subStr)
	return c.JSON(http.StatusOK, &msg)
}

// admGetCurrency doc
// @Tags Currency-代币信息
// @Summary 查询代币信息
// @Param token header string true "token"
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.Currency} "返回数据"
// @Router /adm/get_currency [GET]
func admGetCurrency(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	d.Status = 0
	msg := biz.GetCurrency(d)
	return c.JSON(http.StatusOK, &msg)
}

// admAddCurrency doc
// @Tags Currency-代币信息
// @Summary 添加代币信息
// @Param token header string true "token"
// @Param body body data.Currency true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_currency [POST]
func admAddCurrency(c echo.Context) error {
	d := new(data.Currency)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	//d.Status = 1
	msg := biz.AddCurrency(d)
	return c.JSON(http.StatusOK, &msg)
}

// admDelCurrency doc
// @Tags Currency-代币信息
// @Summary 删除代币信息
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_currency [POST]
func admDelCurrency(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelCurrency(d)
	return c.JSON(http.StatusOK, &msg)
}

// admModCurrency doc
// @Tags Currency-代币信息
// @Summary 修改代币信息
// @Param token header string true "token"
// @Param body body data.currency true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/mod_currency [POST]
func admModCurrency(c echo.Context) error {
	d := new(data.Currency)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.ModCurrency(d)
	return c.JSON(http.StatusOK, &msg)
}

// admAddCategory doc
// @Tags Category-分类
// @Summary 增加分类
// @Param token header string true "token"
// @Param body body data.Category true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_category [POST]
func admAddCategory(c echo.Context) error {
	d := new(data.Category)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.AddCategory(d)
	return c.JSON(http.StatusOK, &msg)
}

// admDelCategory doc
// @Tags Category-分类
// @Summary 删除分类
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_category [POST]
func admDelCategory(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelCategory(d)
	return c.JSON(http.StatusOK, &msg)
}

// admModCategory doc
// @Tags Category-分类
// @Summary 修改分类
// @Param token header string true "token"
// @Param body body data.Category true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/mod_category [post]
func admModCategory(c echo.Context) error {
	d := new(data.Category)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.ModCategory(d)
	return c.JSON(http.StatusOK, &msg)
}

// admGetCategory doc
// @Tags UseApi
// @Summary 查询分类
// @Param body query data.CategoryQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Category} "返回数据"
// @Router /api/get_category [GET]
func admGetCategory(c echo.Context) error {
	d := new(data.CategoryQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetCategory(d)
	return c.JSON(http.StatusOK, &msg)
}

// admLogin doc
// @Tags auth-登陆认证
// @Summary 登陆
// @Param body body biz.LoginData true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /auth/login [post]
func admLogin(c echo.Context) error {
	loginData := new(biz.LoginData)
	err := c.Bind(loginData)
	if err != nil {
		log.Println(err.Error())
	}
	code, msg := biz.AdminLogin(loginData)
	return c.JSON(code, &msg)
}

// admAddLabel doc
// @Tags Label-标签
// @Summary 增加标签
// @Param token header string true "token"
// @Param body body data.Label true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_label [POST]
func admAddLabel(c echo.Context) error {
	d := new(data.Label)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.AddLabel(d)
	return c.JSON(http.StatusOK, &msg)
}

// admDelLabel doc
// @Tags Label-标签
// @Summary 删除标签
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_label [POST]
func admDelLabel(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelLabel(d)
	return c.JSON(http.StatusOK, &msg)
}

// admModLabel doc
// @Tags Label-标签
// @Summary 修改标签
// @Param token header string true "token"
// @Param body body data.Label true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/mod_label [POST]
func admModLabel(c echo.Context) error {
	d := new(data.Label)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.ModLabel(d)
	return c.JSON(http.StatusOK, &msg)
}

// admGetLabel doc
// @Tags Label-标签
// @Summary 查询标签
// @Success 200 {object} biz.BaseJson{data=[]data.Category} "返回数据"
// @Router /api/get_category [GET]
func admGetLabel(c echo.Context) error {
	msg := biz.GetLabel()
	return c.JSON(http.StatusOK, &msg)
}

// admAddClass doc
// @Tags Class-类型
// @Summary 增加类型
// @Param token header string true "token"
// @Param body body data.Class true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_class[POST]
func admAddClass(c echo.Context) error {
	d := new(data.Class)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.AddClass(d)
	return c.JSON(http.StatusOK, &msg)
}

// admDelClass doc
// @Tags Class-类型
// @Summary 删除类型
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_class [POST]
func admDelClass(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelClass(d)
	return c.JSON(http.StatusOK, &msg)
}

// admGetClass doc
// @Tags Class-类型
// @Summary 查询类型
// @Success 200 {object} biz.BaseJson{data=[]data.Class} "返回数据"
// @Router /api/get_class [GET]
func admGetClass(c echo.Context) error {
	msg := biz.GetClass()
	return c.JSON(http.StatusOK, &msg)
}

// admAddChain doc
// @Tags Chain-链
// @Summary 增加链
// @Param token header string true "token"
// @Param body body data.Chain true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/add_chain[POST]
func admAddChain(c echo.Context) error {
	d := new(data.Chain)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.AddChain(d)
	return c.JSON(http.StatusOK, &msg)
}

// admDelChain doc
// @Tags Chain-链
// @Summary 删除链
// @Param token header string true "token"
// @Param body body data.DelQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/del_chain [POST]
func admDelChain(c echo.Context) error {
	d := new(data.DelQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.DelChain(d)
	return c.JSON(http.StatusOK, &msg)
}

// admModChain doc
// @Tags Chain-链
// @Summary 修改链
// @Param token header string true "token"
// @Param body body data.Chain true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /adm/mod_chain [POST]
func admModChain(c echo.Context) error {
	d := new(data.Chain)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.ModChain(d)
	return c.JSON(http.StatusOK, &msg)
}

// admGetChain doc
// @Tags Chain-链
// @Summary 查询链
// @Param token header string true "token"
// @Success 200 {object} biz.BaseJson{data=[]data.Chain} "返回数据"
// @Router /adm/get_chain [GET]
func admGetChain(c echo.Context) error {
	msg := biz.GetChain()
	return c.JSON(http.StatusOK, &msg)
}
