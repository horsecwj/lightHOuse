package service

import (
	"help_center/internal/biz"
	"help_center/internal/data"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	KeyUserId = "_USER_ID"
)

func userID(c echo.Context) uint64 {
	return c.Get(KeyUserId).(uint64)
}

func authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		ctx := request.Context()
		authHeader := request.Header.Get("Authorization")
		authScheme := "Bearer "

		if !strings.HasPrefix(authHeader, authScheme) {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid authorization scheme")
		}

		claims, err := data.ParseToken(ctx, strings.TrimPrefix(authHeader, authScheme))
		if err != nil {
			return err
		}

		UserId, err := strconv.ParseUint(claims.Subject, 10, 64)
		if err != nil {
			return err
		}
		c.Set(KeyUserId, UserId)
		return next(c)
	}
}

// apiInvite doc
// @Tags Invite-邀请码
// @Summary 邀请链接
// @Param body query biz.Code true "请求数据"
// @Router /api/invite [post]
func Invite(c echo.Context) error {
	return nil
}

// apiLogin doc
// @Tags auth-谷歌登陆认证
// @Summary 谷歌登陆
// @Param body body biz.ReqGoogleLogin true "请求数据"
// @Success 200 {object} biz.BaseJson{data=string} "返回数据"
// @Router /api/login [post]
func UserLogin(c echo.Context) error {
	d := new(biz.ReqGoogleLogin)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	code := c.QueryParam("code")
	msg := biz.ParseGoogleToken(d.TokenId, code)
	return c.JSON(http.StatusOK, &msg)
}

// apiIntroduce doc
// @Tags 用户信息
// @Summary 用户信息
// @Param token header string true "token"
// @Success 200 {object} biz.BaseJson{data=data.UserData} "返回数据"
// @Router /api/introduce [get]
func UserIntroduce(c echo.Context) error {
	r := c.Request()
	userId := userID(c)
	msg := biz.GetUser(int64(userId), r)
	return c.JSON(http.StatusOK, &msg)
}

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
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}

// useGetArticle doc
// @Tags UseApi
// @Summary 查询文章
// @Param body query data.ArticleQuery true "请求数据"
// @Success 200 {object} biz.BaseJson{data=[]data.Article} "返回数据"
// @Router /api/get_article [GET]
func useGerArticle(c echo.Context) error {
	d := new(data.ArticleQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	d.Status = 0
	msg := biz.GetArticle(d, false)
	biz.Goip(c.Request())
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}

// useGetBanner doc
// @Tags Banner-查看横幅
// @Summary 查看横幅
// @Success 200 {object} biz.BaseJson{data=[]data.Banner} "返回数据"
// @Router /api/get_banner [GET]
func useGetBanner(c echo.Context) error {
	msg := biz.GetBanner()
	go StoreCache(c.Request().URL.Path, &msg)
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
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}

// useGetTopGainers doc
// @Tags GetTopGainers-查询头号玩家
// @Summary 查询头号玩家
// @Success 200 {object} biz.JsonFormat{data=[]data.Cmk} "返回数据"
// @Router /api/get_top_gainers [GET]
func useGetTopGainers(c echo.Context) error {
	d := new(data.GameQuery)
	err := c.Bind(d)
	if err != nil {
		log.Println(err.Error())
	}
	msg := biz.GetGameCmk(d, false)
	go StoreCache(c.Request().URL.Path, &msg)
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
	go StoreCache(c.Request().URL.Path, &msg)
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
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}

// useMatchArticle
// @Tags UseApi
// @Summary title关键字查询文章(至多返回30条数据)
// @Param sub_str query string true "匹配数据"
// @Success 200 {object} biz.JsonFormat{data=[]data.Article} "返回数据"
// @Router /api/match_article [GET]
func useMatchArticle(c echo.Context) error {
	subStr := c.QueryParam("sub_str")
	msg := biz.MatchArticle(subStr, true)
	go StoreCache(c.Request().URL.Path, &msg)
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
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}

// useGetClass doc
// @Tags Class-类型
// @Summary 查询类型
// @Success 200 {object} biz.BaseJson{data=[]data.Class} "返回数据"
// @Router /api/get_class [GET]
func useGetClass(c echo.Context) error {
	msg := biz.GetClass(false)
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}

// useGetChain doc
// @Tags Chain-链
// @Summary 查询链
// @Param token header string true "token"
// @Success 200 {object} biz.BaseJson{data=[]data.Chain} "返回数据"
// @Router /api/get_chain [GET]
func useGetChain(c echo.Context) error {
	msg := biz.GetChain(false)
	go StoreCache(c.Request().URL.Path, &msg)
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
	go StoreCache(c.Request().URL.Path, &msg)
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
	go StoreCache(c.Request().URL.Path, &msg)
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
	go StoreCache(c.Request().URL.Path, &msg)
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
	go StoreCache(c.Request().URL.Path, &msg)
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
	go StoreCache(c.Request().URL.Path, &msg)
	return c.JSON(http.StatusOK, &msg)
}
