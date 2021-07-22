package database

import "goarticle/app/models"

func (d *DB) CreateArticle(p *models.Article) error {
	res, err := d.db.Exec(insertArticleScheme, p.Title, p.Content, p.Author)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) GetArticle() ([]*models.Article, error) {
	var posts []*models.Article
	err := d.db.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		return posts, err
	}

	return posts, nil
}