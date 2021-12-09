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