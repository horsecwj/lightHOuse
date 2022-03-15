package service

import (
	"help_center/internal/biz"
	"help_center/internal/data"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// useGetGame doc
// @Tags UseApi
// @Summary 查询游戏简介
// @Param body query data.GameQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Game} "返回数据"
// @Router /api/get_game [GET]
func useGetGame(c echo.Context) error {
	d := new(data.GameQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetGame(d, false)
	biz.Goip(c.Request())
	return c.JSON(http.StatusOK, &msg)
}

// useGetArticle doc
// @Tags UseApi
// @Summary 查询文章
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Article} "返回数据"
// @Router /api/get_Article [GET]
func useGerArticle(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	d.Status = 0
	msg := biz.GetArticle(d, false)
	biz.Goip(c.Request())
	return c.JSON(http.StatusOK, &msg)
}

// useGetBanner doc
// @Tags Banner-查看横幅
// @Summary 查看横幅
// @Param token header string true "token"
// @Success 200 {object} biz.BaseJson{data=[]data.Banner} "返回数据"
// @Router /adm/get_banner [GET]
func useGetBanner(c echo.Context) error {
	msg := biz.GetBanner()
	return c.JSON(http.StatusOK, &msg)
}

// useGetCategory doc
// @Tags UseApi
// @Summary 查询分类
// @Param body query data.CategoryQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Category} "返回数据"
// @Router /api/get_category [GET]
func useGetCategory(c echo.Context) error {
	d := new(data.CategoryQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetCategory(d, false)
	return c.JSON(http.StatusOK, &msg)
}

// useGetTopGainers doc
// @Tags GetTopGainers-查询头号玩家
// @Summary 查询头号玩家
// @Success 200 {object} biz.JsonFormat{data=[]data.Cmk} "返回数据"
// @Router /api/get_gamecmk [GET]
func useGetTopGainers(c echo.Context) error {
	d := new(data.GameQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetGameCmk(d, false)
	return c.JSON(http.StatusOK, &msg)
}

// useGetTopLosers doc
// @Tags GetTopLosers-查询头号失败者
// @Summary 查询头号失败者
// @Success 200 {object} biz.JsonFormat{data=[]data.Cmk} "返回数据"
// @Router /api/get_top_losers [GET]
func useGetTopLosers(c echo.Context) error {
	d := new(data.GameQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetGameCmk(d, true)
	return c.JSON(http.StatusOK, &msg)
}

// useGetLikeArticle doc
// @Tags UseApi
// @Summary 查询相关文章
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Article} "返回数据"
// @Router /api/get_like_article [GET]
func useGetLikeArticle(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetLikeArticle(d)
	return c.JSON(http.StatusOK, &msg)
}

// useGetLikeGame doc
// @Tags UseApi
// @Summary 查询相关游戏
// @Param body query data.GameQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Game} "返回数据"
// @Router /api/get_like_gamee [GET]
func useGetLikeGame(c echo.Context) error {
	d := new(data.GameQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetLikeGame(d)
	return c.JSON(http.StatusOK, &msg)
}

// useGetClass doc
// @Tags Class-类型
// @Summary 查询类型
// @Success 200 {object} biz.BaseJson{data=[]data.Class} "返回数据"
// @Router /use/get_class [GET]
func useGetClass(c echo.Context) error {
	msg := biz.GetClass(false)
	return c.JSON(http.StatusOK, &msg)
}

// useGetChain doc
// @Tags Chain-链
// @Summary 查询链
// @Param token header string true "token"
// @Success 200 {object} biz.BaseJson{data=[]data.Chain} "返回数据"
// @Router /use/get_chain [GET]
func useGetChain(c echo.Context) error {
	msg := biz.GetChain(false)
	return c.JSON(http.StatusOK, &msg)
}

// useGetGameParameter doc
// @Tags GameParameter-游戏参数
// @Summary 查询游戏参数
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.GameParameter} "返回数据"
// @Router /api/get_game_parameter [GET]
func useGetGameParameter(c echo.Context) error {
	d := new(data.GameQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetGameParameter(d, false)
	return c.JSON(http.StatusOK, &msg)
}

// useGetGameValue doc
// @Tags GameValue-游戏价值
// @Summary 查询游戏价值
// @Param body query data.GameQuery true "请求数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.GameValue} "返回数据"
// @Router /api/get_game_value [GET]
func useGetGameValue(c echo.Context) error {
	d := new(data.GameQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetGameValue(d)
	biz.Goip(c.Request())
	return c.JSON(http.StatusOK, &msg)
}

// useGetVideoCourse doc
// @Tags VideoCourse-视频教程
// @Summary 查询视频教程
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.CourseBanner} "返回数据"
// @Router /api/get_game_value [GET]
func useGetVideoCourse(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetCourse(d, true, false)
	biz.Goip(c.Request())
	return c.JSON(http.StatusOK, &msg)
}

// useGetImageCourse doc
// @Tags GetImageCourse-图文教程
// @Summary 查询图文教程
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.CourseBanner} "返回数据"
// @Router /api/get_game_value [GET]
func useGetImageCourse(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetCourse(d, false, true)
	biz.Goip(c.Request())
	return c.JSON(http.StatusOK, &msg)
}

// useGetCourse doc
// @Tags GetCourse-教程
// @Summary 查询教程
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.CourseBanner} "返回数据"
// @Router /api/get_value [GET]
func useGetCourse(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetCourse(d, true, true)
	return c.JSON(http.StatusOK, &msg)
}
