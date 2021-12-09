package article

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

type ArticleRepository interface {
	Save(ctx context.Context, article Article) (ID int64, err error)
	// Update(ctx context.Context, ID int64, updatedArticle Article) (err error)
	// FindByID(ctx context.Context, ID int64) (article Article, err error)
	// FindMany(ctx context.Context) (bunchOfArticles []Article, err error)
	// FindManySpecificProfile(ctx context.Context, articleID int64) (bunchOfArticles []Article, err error)
}

type articleRepositoryImpl struct {
	db        *sql.DB
	tableName string
}

func NewArticleRepository(db *sql.DB, tableName string) ArticleRepository {
	return &articleRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}


func (r *articleRepositoryImpl) Save(ctx context.Context, article Article) (ID int64, err error) {
	command := fmt.Sprintf("INSERT INTO %s (title, subtitle, content, status, createdAt, publishedAt, authorId) VALUES (?, ?, ?, ?, ?, ?, ?)", r.tableName)
	stmt, err := r.db.PrepareContext(ctx, command)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()
  
	result, err := stmt.ExecContext(
		ctx,
		article.Title,
		article.Subtitle,
		article.Content,
		article.Status,
		article.CreatedAt,
		article.PublishedAt,
		article.Author.ID,

	)

	if err != nil {
		log.Println(err)
		return
	}

	ID, _ = result.LastInsertId()

	return
}