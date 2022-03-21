package schedule

import "github.com/robfig/cron"

// 创建定时执行任务
func create() (c *cron.Cron, err error) {

	c = cron.New()

	// 添加article任务执行
	err = createArtSpider(c)
	if err != nil {

		return
	}

	// 添加topGameFi任务同步
	err = createTopGameFi(c)
	if err != nil {

		return
	}

	return
}

func Run() (err error) {

	c, err := create()

	if err != nil {

		return
	}

	c.Run()
	return
}
