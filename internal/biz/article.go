package biz

import (
	"help_center/internal/data"
	"log"

	"gorm.io/gorm"
)

// GetArticle 获取文章列表
func GetArticle(d *data.ArticleQuery, isAdm bool) *JsonFormat {
	if isAdm {
		err := data.OutArticleAdd()
		if err != nil {
			log.Println(err)
		}
	}
	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 {
		d.PageSize = 10
	}
	num := d.ArticleCount()
	if num > 0 {
		list := d.ArticleSearch(isAdm)
		return &JsonFormat{Code: 1, Page: d.Page, PageSize: d.PageSize, PageNum: num/d.PageSize + 1, ArticleNum: num, Data: list}
	}
	return &JsonFormat{Code: 0, Page: d.Page, PageSize: d.PageSize, PageNum: 0, ArticleNum: num, Data: nil}
}

// MatchArticle 搜索文章列表
func MatchArticle(subStr string, user bool) *JsonFormat {
	if subStr != "" {
		list, lenList := data.ArticleMatch(subStr, user)
		return &JsonFormat{Code: 1, Page: 1, PageSize: lenList, PageNum: 1, ArticleNum: lenList, Data: list}
	}
	return &JsonFormat{Code: 0, Page: 0, PageSize: 0, PageNum: 0, ArticleNum: 0, Data: "当前参数获取到的文章数量为0"}
}

// AddArticle 添加文章
func AddArticle(d *data.Article) *BaseJson {
	tx := data.GetDbCli().Session(&gorm.Session{})
	if d.Hot != 0 {
		if tx.Where("status = ?", d.Status).Error == nil {
			return &BaseJson{Code: 0, Data: "热门文章已存在"}
		}
	}
	err := data.ArticleAdd(d)
	if err != nil {
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功添加文章"}
	}
}

// DelArticle 删除文章
func DelArticle(d *data.DelQuery) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	err := data.ArticleDelete(d.Id)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功成功删除文章"}
	}
}

// ModArticle 修改文章
func ModArticle(d *data.Article) *BaseJson {
	if d.Id == 0 {
		return &BaseJson{Code: 0, Data: "参数 id 值不应为0"}
	}
	if d.Hot != 0 {
		row, _ := data.VerificationArticleHot(d.Id)
		if row.Hot != d.Hot {
			err := data.VerificationHot(d.Hot)
			if err == nil {
				return &BaseJson{Code: 0, Data: "热门文章已存在"}
			}
		}
	}
	err := data.ArticleUpdate(d)
	if err != nil {
		log.Println(err.Error())
		return &BaseJson{Code: 0, Data: err.Error()}
	} else {
		return &BaseJson{Code: 1, Data: "成功修改文章"}
	}
}

// 获取相关的文章
func GetLikeArticle(d *data.ArticleQuery) *JsonFormat {
	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 {
		d.PageSize = 4
	}
	num := d.ArticleCount()
	if num > 0 {
		list := d.LikeArticle()
		return &JsonFormat{Code: 1, Page: d.Page, PageSize: d.PageSize - 1, PageNum: num/d.PageSize + 1, ArticleNum: num, Data: list}
	}
	return &JsonFormat{Code: 0, Page: d.Page, PageSize: d.PageSize, PageNum: 0, ArticleNum: num, Data: nil}
}

//获取视频教程及图片教程
func GetCourse(d *data.ArticleQuery, video bool, image bool) *JsonFormat {
	if d.Page == 0 {
		d.Page = 1
	}
	if d.PageSize == 0 {
		d.PageSize = 12
	}
	num := d.CourseCount(video, image)
	if num > 0 {
		list := d.Course(video, image)
		return &JsonFormat{Code: 1, Page: d.Page, PageSize: d.PageSize, PageNum: num/d.PageSize + 1, ArticleNum: num, Data: list}
	}
	return &JsonFormat{Code: 0, Page: d.Page, PageSize: d.PageSize, PageNum: 0, ArticleNum: num, Data: nil}
}
