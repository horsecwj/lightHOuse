package service

import (
	"github.com/labstack/echo/v4"
)

func UserRouter(user *echo.Group) {

	user.POST("/login", UserLogin)
	user.GET("/introduce", UserIntroduce, authenticate)
	user.POST("/invite", Invite)

	user.GET("/get_game", useGetGame)
	user.GET("/get_article", useGerArticle)
	user.GET("/match_article", useMatchArticle)

	user.GET("/get_banner", useGetBanner)
	user.GET("/get_game_parameter", useGetGameParameter)
	user.GET("/get_top_gainers", useGetTopGainers)
	user.GET("/get_top_losers", useGetTopLosers)
	user.GET("/get_like_article", useGetLikeArticle)
	user.GET("/get_like_game", useGetLikeGame)
	user.GET("/get_game_value", useGetGameValue)
	user.GET("/get_category", useGetCategory)

	user.GET("/get_video_course", useGetVideoCourse)
	user.GET("/get_image_course", useGetImageCourse)
	user.GET("/get_course", useGetCourse)

	user.GET("/get_class", useGetClass)
	user.GET("/get_chain", useGetChain)
}

func AdmRouter(adm *echo.Group) {

	adm.POST("/add_banner", admAddBanner)
	adm.POST("/del_banner", admDelBanner)
	adm.POST("/mod_banner", admModBanner)
	adm.GET("/get_banner", admGetBanner)

	adm.POST("/add_article", admAddArticle)
	adm.POST("/del_article", admDelArticle)
	adm.POST("/mod_article", admModArticle)
	adm.GET("/get_article", admGetArticle)
	adm.GET("/match_article", admMatchArticle)
	adm.POST("/more_mod_article", admMoreModArticle)

	adm.POST("/add_category", admAddCategory)
	adm.POST("/del_category", admDelCategory)
	adm.POST("/mod_category", admModCategory)
	adm.GET("/get_category", admGetCategory)

	adm.POST("/upload_img", admUploadImg)
	adm.POST("/fit_img", admFitImg)

	adm.POST("/add_game", admAddGame)
	adm.POST("/del_game", admDelGame)
	adm.POST("/mod_game", admModGame)
	adm.GET("/get_game", admGetGame)
	adm.GET("/match_game", admMatchGame)

	adm.POST("/add_currency", admAddCurrency)
	adm.POST("/del_currency", admDelCurrency)
	adm.POST("/mod_currency", admModCurrency)
	adm.GET("/get_currency", admGetCurrency)

	adm.POST("/add_class", admAddClass)
	adm.POST("/del_class", admDelClass)
	adm.POST("/mod_class", admModClass)
	adm.GET("/get_class", admGetClass)

	adm.POST("/add_chain", admAddChain)
	adm.POST("/del_chain", admDelChain)
	adm.POST("/mod_chain", admModChain)
	adm.GET("/get_chain", admGetChain)

	adm.POST("/add_label", admAddLabel)
	adm.POST("/del_label", admDelLabel)
	adm.POST("/mod_label", admModLabel)
	adm.GET("/get_label", admGetLabel)

	adm.GET("/get_data", admGetData)
}
func AuthRouter(user *echo.Group) {
	user.POST("/login", admLogin)
}
