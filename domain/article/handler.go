package article

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sangianpatrick/devoria-article-service/middleware"
	"github.com/sangianpatrick/devoria-article-service/response"
)

type ArticleHTTPHandler struct {
	Validate *validator.Validate
	Usecase  ArticleUsecase
}

func NewArticleHTTPHandler(
	router *mux.Router,
	validate *validator.Validate,
	usecase ArticleUsecase,
	bearerAuthMiddleware middleware.RouteMiddleware,
) {
	handler := &ArticleHTTPHandler{
		Validate: validate,
		Usecase:  usecase,
	}

	router.HandleFunc("/v1/articles", bearerAuthMiddleware.Verify(handler.CreateArticle)).Methods(http.MethodPost)
	router.HandleFunc("/v1/articles/id", bearerAuthMiddleware.Verify(handler.GetArticle)).Methods(http.MethodGet)
	router.HandleFunc("/v1/articles", bearerAuthMiddleware.Verify(handler.GetArticles)).Methods(http.MethodGet)
	router.HandleFunc("/v1/articles/user", bearerAuthMiddleware.Verify(handler.GetArticlesByUserID)).Methods(http.MethodGet)
	router.HandleFunc("/v1/articles", bearerAuthMiddleware.Verify(handler.UpdateArticle)).Methods(http.MethodPut)
}

func (handler *ArticleHTTPHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {

	var resp response.Response
	var params CreateArticleRequest
	var ctx = r.Context()
  params.Email = r.Header.Get("userEmail")

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp = response.Error(response.StatusUnprocessabelEntity, nil, err)
		resp.JSON(w)
		return
		
	}


	err = handler.Validate.StructCtx(ctx, params)
	if err != nil {
		resp = response.Error(response.StatusInvalidPayload, nil, err)
		resp.JSON(w)
		return
	}

	resp = handler.Usecase.CreateArticle(ctx, params)
	resp.JSON(w)
}

func (handler *ArticleHTTPHandler) GetArticle(w http.ResponseWriter, r *http.Request) {

	var resp response.Response
	var ctx = r.Context()
	ids, _ := r.URL.Query()["id"]

	i, _ := strconv.ParseInt(ids[0], 10, 64)
	ID := int64(i)

	resp = handler.Usecase.GetArticleByID(ctx, ID)
	resp.JSON(w)

}

func (handler *ArticleHTTPHandler) GetArticles(w http.ResponseWriter, r *http.Request) {

	var resp response.Response
	var ctx = r.Context()

	resp = handler.Usecase.GetArticles(ctx)
	resp.JSON(w)

}

func (handler *ArticleHTTPHandler) GetArticlesByUserID(w http.ResponseWriter, r *http.Request) {

	var resp response.Response
	var ctx = r.Context()

	userIDs, _ := r.URL.Query()["userId"]

	i, _ := strconv.ParseInt(userIDs[0], 10, 64)
	userID := int64(i)

	resp = handler.Usecase.GetArticlesByUserID(ctx, userID)
	resp.JSON(w)

}

func (handler *ArticleHTTPHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {

	var resp response.Response
	var ctx = r.Context()
	var params EditArticleRequest

	ids, _ := r.URL.Query()["id"]
	ID, err := strconv.ParseInt(ids[0], 10, 64)
	

	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		resp = response.Error(response.StatusUnprocessabelEntity, nil, err)
		resp.JSON(w)
		return
	}

	err = handler.Validate.StructCtx(ctx, params)
	if err != nil {
		resp = response.Error(response.StatusInvalidPayload, nil, err)
		resp.JSON(w)
		return
	}

	resp = handler.Usecase.UpdateArticle(ctx, ID, params)
	resp.JSON(w)



}


