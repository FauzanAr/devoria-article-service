package article

import (
	"context"
	"time"

	"github.com/sangianpatrick/devoria-article-service/exception"
	"github.com/sangianpatrick/devoria-article-service/response"
	"github.com/sangianpatrick/devoria-article-service/domain/account"
)

type ArticleUsecase interface {
	CreateArticle(ctx context.Context, params CreateArticleRequest) (resp response.Response)
	GetArticleByID(ctx context.Context, ID int64) (resp response.Response)
	GetArticles(ctx context.Context) (resp response.Response)
	GetArticlesByUserID(ctx context.Context, userID int64) (resp response.Response) 
	UpdateArticle(ctx context.Context, ID int64, params EditArticleRequest) (resp response.Response)
	
}

type articleUsecaseImpl struct {
	location					*time.Location
	repository   			ArticleRepository
	accountRepository account.AccountRepository
}

func NewArticleUsecase(
	location *time.Location,
	repository ArticleRepository,
	accountRepository account.AccountRepository,
) ArticleUsecase {
	return &articleUsecaseImpl{
		location:     			location,
		repository:   			repository,
		accountRepository:	accountRepository,		
	}
}

func (u *articleUsecaseImpl) CreateArticle(ctx context.Context, params CreateArticleRequest) (resp response.Response) {
	account, err := u.accountRepository.FindByEmail(ctx, params.Email)
	if err != nil {
		return response.Error(response.StatusConflicted, nil, exception.ErrConflicted)
	}
	newArticle := Article{}
	newArticle.Title = params.Email
	newArticle.Subtitle = params.Subtitle
	newArticle.Content = params.Content
	newArticle.Status = params.Status
	newArticle.CreatedAt = time.Now().In(u.location)
	newArticle.PublishedAt = time.Now().In(u.location)
	newArticle.Author = account

	u.repository.Save(ctx, newArticle)
	if err != nil {
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}
	createArticleResponse := CreateArticleResponse{}
	createArticleResponse.Ok = true
	createArticleResponse.Message = "Article Created"
	

	return response.Success(response.StatusCreated, createArticleResponse)
}

func (u *articleUsecaseImpl) GetArticleByID(ctx context.Context, ID int64) (resp response.Response) {

	article, err := u.repository.FindByID(ctx, ID)
	if err != nil {
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}

	return response.Success(response.StatusOK, article)
	
}

func (u *articleUsecaseImpl) GetArticles(ctx context.Context) (resp response.Response) {

	articles, err := u.repository.FindMany(ctx)
	if err != nil {
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}
	
	return response.Success(response.StatusOK, articles)
	
}

func (u *articleUsecaseImpl) GetArticlesByUserID(ctx context.Context, userID int64) (resp response.Response) {

	articles, err := u.repository.FindManySpecificProfile(ctx, userID)
	if err != nil {
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}
	
	return response.Success(response.StatusOK, articles)
	
}

func (u *articleUsecaseImpl) UpdateArticle(ctx context.Context, ID int64, params EditArticleRequest) (resp response.Response) {
	updatedArticle := Article{}
	updatedArticle.Title = params.Title
	updatedArticle.Subtitle = params.Subtitle
	updatedArticle.Content = params.Content

	err := u.repository.Update(ctx, ID, updatedArticle)
	if err != nil {
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}

	editArticleResponse := EditArticleResponse{}
	editArticleResponse.Ok = true
	editArticleResponse.Message = "Article Updated"
	
	return response.Success(response.StatusOK, editArticleResponse)
	
}

