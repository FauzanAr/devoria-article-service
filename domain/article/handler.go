package article

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"

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
	router.HandleFunc("/v1/articles", bearerAuthMiddleware.Verify(handler.GetArticle)).Methods(http.MethodGet)
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

	fmt.Println(ID)
	resp = handler.Usecase.GetArticleByID(ctx, ID)
	resp.JSON(w)

}