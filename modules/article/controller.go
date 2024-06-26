package article

import (
	"net/http"
	"strconv"
	"time"

	"github.com/OctavianoRyan25/be-agriculture/base"
	"github.com/OctavianoRyan25/be-agriculture/constants"
	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	useCase UseCase
}

func NewArticleController(useCase UseCase) *ArticleController {
	return &ArticleController{
		useCase: useCase,
	}
}

func (c *ArticleController) StoreArticle(e echo.Context) error {
	role := e.Get("role").(string)
	if role != "admin" {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Forbidden access",
			Code:    http.StatusForbidden,
		}
		return e.JSON(http.StatusForbidden, errRes)
	}

	title := e.FormValue("title")
	content := e.FormValue("content")
	image, err := e.FormFile("image")
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return e.JSON(http.StatusBadRequest, errRes)
	}

	file, err := image.Open()
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return e.JSON(http.StatusBadRequest, errRes)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	imagePath := "public/" + timestamp

	url, err := UploadToCloudinary(file, imagePath)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return e.JSON(http.StatusInternalServerError, errRes)
	}

	article := &Article{
		Title:   title,
		Content: content,
		Image:   url,
	}
	article, err = c.useCase.StoreArticle(article)
	if err != nil {
		errResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    constants.ErrCodeBadRequest,
		}
		return e.JSON(constants.ErrCodeBadRequest, errResponse)
	}
	res := &base.SuccessResponse{
		Status:  "success",
		Message: "Article created",
		Data:    ArticleTOResponse(article),
	}
	return e.JSON(constants.Created, res)
}

func (c *ArticleController) GetArticle(e echo.Context) error {
	id, _ := strconv.Atoi(e.Param("id"))
	article, err := c.useCase.GetArticle(id)
	if err != nil {
		errResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return e.JSON(http.StatusNotFound, errResponse)
	}
	res := &base.SuccessResponse{
		Status:  "success",
		Message: "Success get article",
		Data:    ArticleTOResponse(article),
	}
	return e.JSON(http.StatusOK, res)
}

func (c *ArticleController) GetAllArticles(e echo.Context) error {
	articles, err := c.useCase.GetAllArticles()
	if err != nil {
		errResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return e.JSON(http.StatusNotFound, errResponse)
	}
	var res []ResponseArticle
	for _, article := range articles {
		res = append(res, *ArticleTOResponse(&article))
	}
	resp := &base.SuccessResponse{
		Status:  "success",
		Message: "All articles",
		Data:    res,
	}
	return e.JSON(http.StatusOK, resp)
}

func (c *ArticleController) UpdateArticle(e echo.Context) error {
	role := e.Get("role").(string)
	if role != "admin" {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Forbidden access",
			Code:    http.StatusForbidden,
		}
		return e.JSON(http.StatusForbidden, errRes)
	}

	id, _ := strconv.Atoi(e.Param("id"))
	getArticle, err := c.useCase.GetArticle(id)
	if err != nil {
		errResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return e.JSON(http.StatusNotFound, errResponse)
	}
	err = DeleteFromCloudinary(getArticle.Image)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return e.JSON(http.StatusBadRequest, errRes)
	}

	title := e.FormValue("title")
	content := e.FormValue("content")
	image, err := e.FormFile("image")
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return e.JSON(http.StatusBadRequest, errRes)
	}

	file, err := image.Open()
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return e.JSON(http.StatusBadRequest, errRes)
	}

	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	imagePath := "public/" + timestamp

	url, err := UploadToCloudinary(file, imagePath)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return e.JSON(http.StatusInternalServerError, errRes)
	}

	article := &Article{
		Title:   title,
		Content: content,
		Image:   url,
	}
	article, err = c.useCase.UpdateArticle(article, id)
	if err != nil {
		errResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    constants.ErrCodeBadRequest,
		}
		return e.JSON(constants.ErrCodeBadRequest, errResponse)
	}
	res := &base.SuccessResponse{
		Status:  "success",
		Message: "Article updated",
		Data:    ArticleTOResponse(article),
	}
	return e.JSON(http.StatusOK, res)
}

func (c *ArticleController) DeleteArticle(e echo.Context) error {
	role := e.Get("role").(string)
	if role != "admin" {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: "Forbidden access",
			Code:    http.StatusForbidden,
		}
		return e.JSON(http.StatusForbidden, errRes)
	}

	id, _ := strconv.Atoi(e.Param("id"))
	getArticle, err := c.useCase.GetArticle(id)
	if err != nil {
		errResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return e.JSON(http.StatusNotFound, errResponse)
	}
	err = DeleteFromCloudinary(getArticle.Image)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return e.JSON(http.StatusBadRequest, errRes)
	}

	err = c.useCase.DeleteArticle(id)
	if err != nil {
		errResponse := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    constants.ErrCodeBadRequest,
		}
		return e.JSON(constants.ErrCodeBadRequest, errResponse)
	}
	res := &base.SuccessResponse{
		Status:  "success",
		Message: "Article deleted",
	}

	return e.JSON(http.StatusOK, res)
}
