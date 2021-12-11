package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sangianpatrick/devoria-article-service/middleware"
	"github.com/sangianpatrick/devoria-article-service/response"
)

type AccountHTTPHandler struct {
	Validate *validator.Validate
	Usecase  AccountUsecase
}

func NewAccountHTTPHandler(
	router *mux.Router,
	basicAuthMiddleware middleware.RouteMiddleware,
	validate *validator.Validate,
	usecase AccountUsecase,
	bearerAuthMiddleware middleware.RouteMiddleware,
) {
	handler := &AccountHTTPHandler{
		Validate: validate,
		Usecase:  usecase,
	}

	router.HandleFunc("/v1/accounts/registration", basicAuthMiddleware.Verify(handler.Register)).Methods(http.MethodPost)
	router.HandleFunc("/v1/accounts/login", basicAuthMiddleware.Verify(handler.Login)).Methods(http.MethodPost)
	router.HandleFunc("/v1/accounts/profile", bearerAuthMiddleware.Verify(handler.GetProfile)).Methods(http.MethodGet)
	router.HandleFunc("/v1/accounts/profile", bearerAuthMiddleware.Verify(handler.UpdateProfile)).Methods(http.MethodPut)
}

func (handler *AccountHTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var params AccountRegistrationRequest
	var ctx = r.Context()

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

	resp = handler.Usecase.Register(ctx, params)
	resp.JSON(w)
}

func (handler *AccountHTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var params AccountAuthenticationRequest
	var ctx = r.Context()
	
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

	resp = handler.Usecase.Login(ctx, params)
	resp.JSON(w)
}

func (handler *AccountHTTPHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var ctx = r.Context()
	var email = r.Header.Get("userEmail")

	resp = handler.Usecase.GetProfile(ctx, email)
	resp.JSON(w)

}

func (handler *AccountHTTPHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	var ctx = r.Context()
	var params AccountUpdateRequest
	ids, _ := r.URL.Query()["id"]
	i, err := strconv.ParseInt(ids[0], 10, 64)
	params.ID = int64(i)
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
	
	resp = handler.Usecase.UpdateProfile(ctx, params)
	resp.JSON(w)

}