package data

import (
	"time"
)

type Banner struct {
	Id    int64  `json:"id"`
	Chain int64  `json:"chain"`
	Cover string `json:"cover"`
}

type Article struct {
	Id       int64   `json:"id"`
	Lang     string  `json:"lang"`
	Status   int     `json:"status"`
	CateId   int64   `json:"cate_id"`
	Title    string  `json:"title"`
	Summary  string  `json:"summary"`
	Cover    string  `josn:"cover"`
	Markdown string  `json:"markdown"`
	RichText string  `json:"rich_text"`
	Hot      int64   `json:"hot"`
	Label    []Label `json:"label" gorm:"many2many:article_label"`
	GameId   int64   `json:"game_id"`
}

type article struct {
	Id       int64   `json:"id"`
	Lang     string  `json:"lang"`
	Status   int     `json:"status"`
	CateId   int64   `json:"cate_id"`
	Title    string  `json:"title"`
	Summary  string  `json:"summary"`
	Cover    string  `json:"cover"`
	Markdown string  `json:"markdown"`
	RichText string  `json:"rich_text"`
	Hot      int64   `json:"hot"`
	Label    []Label `json:"label" gorm:"many2many:article_label"`
	Created  string  `json:"created"`
	Updated  string  `json:"updated"`
	GameId   int64   `json:"game_id"`
}

type Game struct {
	Id              int64      `json:"id"`
	GameName        string     `json:"game_name"`
	Cover           string     `json:"cover"`
	Summary         string     `json:"summary"`
	Status          int        `json:"status"`
	Lang            string     `json:"lang"`
	Currency        []currency `json:"currency" gorm:"many2many:game_currency"`
	Chain           []Chain    `json:"chain" gorm:"many2many:game_chain"`
	Label           []Label    `json:"label" gorm:"many2many:game_label"`
	Class           []Class    `json:"class" gorm:"many2many:game_class"`
	Telegram        string     `json:"telegram"`
	Facebook        string     `json:"facebook"`
	Twitter         string     `json:"twitter"`
	Youtube         string     `json:"youtube"`
	GameUrl         string     `json:"game_url"`
	Guide           string     `json:"guide"`
	Release         int64      `json:"release"`
	AboutGames      string     `json:"about_games"`
	New             []Article  `json:"new" gorm:"many2many:game_article"`
	Stragegy        string     `json:"stragegy"`
	RevenueAnalysis string     `json:"revenue_analysis"`
	Created         time.Time  `json:"created"`
}

type GameParameter struct {
	Id        int64  `json:"id"`
	Coin      string `json:"coin"`
	GameFi    string `json:"game_fi"`
	Price     string `json:"price"`
	OneDay    string `json:"one_day"`
	OneWeek   string `json:"one_week"`
	DayVolume string `json:"day_volume"`
	MktCap    string `json:"mkt_cap"`
	Status    int    `json:"status"`
}

type GameValue struct {
	Id            int64         `json:"id"`
	GameName      string        `json:"title"`
	Status        string        `json:"status"`
	Chain         []Chain       `json:"chain" gorm:"many2many:game_chain"`
	Class         []Class       `json:"class" gorm:"many2many:game_class"`
	GameParameter GameParameter `json:"game_parameter" gorm:"foreignkey:game_fi;references:game_name"`
}

type Value struct {
	Id        int64   `json:"id"`
	GameName  string  `json:"title"`
	Status    string  `json:"status"`
	Class     []Class `json:"class" gorm:"many2many:game_class"`
	Coin      string  `json:"coin"`
	Price     string  `json:"price"`
	OneDay    string  `json:"one_day"`
	OneWeek   string  `json:"one_week"`
	DayVolume string  `json:"day_volume"`
	MktCap    string  `json:"mkt_cap"`
}

type Currency struct {
	Id           int64     `json:"id"`
	Logo         string    `json:"logo"`
	CurrencyName string    `json:"currency_name"`
	MaxAmount    uint64    `json:"max_amount"`
	Value        uint64    `json:"value"`
	FlowAmount   uint64    `json:"flow_amount"`
	IssueAt      time.Time `json:"issue_at"`
	Address      string    `json:"address"`
}

type currency struct {
	Id           int64  `json:"id"`
	Logo         string `json:"logo"`
	CurrencyName string `json:"currency_name"`
	MaxAmount    uint64 `json:"max_amount"`
	Value        uint64 `json:"value"`
	FlowAmount   uint64 `json:"flow_amount"`
	//IssueAt      time.Time `json:"issue_at"`
	Address string `json:"address"`
}

type GameCurrency struct {
	GameId     int64 `json:"game_id"`
	CurrencyId int64 `json:"currency_id"`
}

type OutArticle struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	OverView    string `json:"over_view"`
	Article     string `json:"article"`
	Articletext string `json:"articletext"`
	Timestamp   int64  `json:"timestamp"`
}

type Category struct {
	Id       int64  `json:"id"`
	Lang     string `json:"lang"`
	Name     string `json:"name"`
	Intro    string `json:"intro"`
	ParentId int64  `json:"parent_id"`
}

type Label struct {
	Id   int64  `json:"id"`
	Lang string `json:"lang"`
	Word string `json:"word"`
	Game []Game `json:"game" gorm:"many2many:game_label"`
}

type ArticleLabel struct {
	ArticleId int64 `json:"article_id"`
	LabelID   int64 `json:"label_id"`
}

type GameLabel struct {
	GameId  int64 `json:"game_id"`
	LabelId int64 `json:"label_id"`
}

type Class struct {
	Id    int64  `json:"id"`
	Class string `json:"class"`
	//Game  []Game `json:"game" gorm:"many2many:game_class"`
}

type GameClass struct {
	GameId  int64 `json:"game_id"`
	ClassID int64 `json:"class_id"`
}

type GameArticle struct {
	GameId    int64 `json:"game_id"`
	ArticleId int64 `json:"article_id"`
}

type Chain struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	ICON string `json:"icon"`
	//Game []Game `json:"game" gorm:"many2many:game_chain"`
}

type GameChain struct {
	GameId  int64 `json:"game_id"`
	ChainId int64 `json:"chain_id"`
}

type CategoryRelation struct {
	Id       int64 `json:"id"`
	ParentId int64 `json:"parent_id"`
}

type ArticleQuery struct {
	Id       int64    `json:"id" query:"id"`
	Status   int      `json:"status" query:"status"`
	CateId   int64    `json:"cate_id" query:"cate_id"`
	Page     int      `json:"page" query:"page"`
	PageSize int      `json:"page_size" query:"page_size"`
	Word     []string `json:"word" query:"word"`
	Hot      int      `json:"hot" query:"hot"`
}
type CategoryQuery struct {
	Id       int64 `json:"id" query:"id"`
	ParentId int64 `json:"parent_id" query:"parent_id"`
}

type DelQuery struct {
	Id int64 `json:"id" query:"id"`
}

type GameQuery struct {
	Id       int64  `json:"id" query:"id"`
	Status   int    `json:"status" query:"status"`
	ChainId  int64  `json:"chain_id" query:"chain_id"`
	ClassId  int64  `json:"class_id" query:"class_id"`
	LabelId  int64  `json:"label_id" query:"label_id"`
	GameFi   string `json:"game_fi" query:"game_fi"`
	Page     int    `json:"page" query:"page"`
	PageSize int    `json:"page_size" query:"page_size"`
}

type Cmk struct {
	Id        int64  `json:"id"`
	GameFi    string `json:"game_fi"`
	Coin      string `json:"coin"`
	Token     string `json:"token"`
	Price     string `json:"price"`
	OneDay    string `json:"one_day"`
	DayVolume string `json:"day_volume"`
}

type likeArticle struct {
	Id      int64     `json:"id"`
	Cover   string    `json:"cover"`
	Title   string    `json:"title"`
	Summary string    `json:"summary"`
	Created time.Time `json:"created"`
}

type game struct {
	Id       int64   `json:"id"`
	GameName string  `json:"game_name"`
	Status   int     `json:"status"`
	Chain    []Chain `json:"chain" gorm:"many2many:game_chain"`
	Class    []Class `json:"class" gorm:"many2many:game_class"`
}
