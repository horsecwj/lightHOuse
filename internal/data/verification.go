package data

func VerificationLabel(label string) (*Label, error) {
	row := &Label{}
	err := db.Where("word = ?", label).First(&row).Error
	return row, err
}

func VerificationChain(chain string) (*Chain, error) {
	row := &Chain{}
	err := db.Where("name = ?", chain).First(&row).Error
	return row, err
}

func VerificationClass(class string) (*Class, error) {
	row := &Class{}
	err := db.Where("class = ?", class).First(&row).Error
	return row, err
}

func VerificationTitle(id int64) (*Article, error) {
	row := &Article{}
	err := db.Where("id = ?", id).First(&row).Error
	return row, err
}

func VerificationArticle(Title string) error {
	row := &article{}
	err := db.Where("title = ?", Title).First(&row).Error
	return err
}

func VerificationArticleHot(id int64) (*article, error) {
	row := &article{}
	err := db.Where("id = ?", id).First(&row).Error
	return row, err
}

func VerificationHot(hot int64) error {
	row := &article{}
	err := db.Where("hot = ?", hot).First(&row).Error
	return err
}

func VerificationGames(name string) error {
	row := &Game{}
	err := db.Where("game_name = ?", name).First(&row).Error
	return err
}

func VerificationGameParameters(name string) error {
	row := &GameParameter{}
	err := db.Where("game_fi = ?", name).First(&row).Error
	return err
}
