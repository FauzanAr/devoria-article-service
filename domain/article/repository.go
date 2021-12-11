package article

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sangianpatrick/devoria-article-service/exception"
)

type ArticleRepository interface {
	Save(ctx context.Context, article Article) (ID int64, err error)
	// Update(ctx context.Context, ID int64, updatedArticle Article) (err error)
	FindByID(ctx context.Context, ID int64) (article Article, err error)
	FindMany(ctx context.Context) (bunchOfArticles []Article, err error)
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

func (r *articleRepositoryImpl) FindByID(ctx context.Context, ID int64) (article Article, err error) {
	query := fmt.Sprintf(`SELECT ar.id, ar.title, ar.subtitle, ar.content, ar.status, ar.publishedAt, ac.id, ac.firstName, ac.lastName, ac.email FROM %s ar 
	LEFT JOIN Account ac
	ON ac.id = ar.authorId
	WHERE ar.id = ?`, r.tableName)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, ID)

	err = row.Scan(
		&article.ID,
		&article.Title,
		&article.Subtitle,
		&article.Content,
		&article.Status,
		&article.PublishedAt,
		&article.Author.ID,
		&article.Author.FirstName,
		&article.Author.LastName,
		&article.Author.Email,
	)

	if err != nil {
		log.Println(err)
		err = exception.ErrNotFound
		return
	}

	return
}

func (r *articleRepositoryImpl) FindMany(ctx context.Context) (bunchOfArticles []Article, err error) {
	query := fmt.Sprintf(`SELECT ar.id, ar.title, ar.subtitle, ar.content, ar.status, ar.publishedAt, ac.id, ac.firstName, ac.lastName, ac.email FROM %s ar 
	LEFT JOIN Account ac
	ON ac.id = ar.authorId`, r.tableName)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}
	defer stmt.Close()

	rows, _ := stmt.QueryContext(ctx)
	
	for rows.Next() {
		article := Article{}

		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Subtitle,
			&article.Content,
			&article.Status,
			&article.PublishedAt,
			&article.Author.ID,
			&article.Author.FirstName,
			&article.Author.LastName,
			&article.Author.Email,
		)
		if err != nil {
			log.Println(err)
			err = exception.ErrNotFound
			return
		} 
		fmt.Println(article)
		bunchOfArticles = append(bunchOfArticles, article)
	}

	fmt.Println("check out")

	return
}