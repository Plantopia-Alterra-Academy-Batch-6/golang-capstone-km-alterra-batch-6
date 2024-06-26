package search

import (
	"net/http"

	"github.com/OctavianoRyan25/be-agriculture/base"
	"github.com/labstack/echo/v4"
)

type SearchController struct {
	searchUsecase Usecase
}

func NewSearchController(searchUsecase Usecase) *SearchController {
	return &SearchController{
		searchUsecase: searchUsecase,
	}
}

func (c *SearchController) Search(ctx echo.Context) error {
	var params PlantSearchParams
	err := ctx.Bind(&params)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}
		return ctx.JSON(http.StatusBadRequest, errRes)
	}

	plants, err := c.searchUsecase.Search(params)
	if err != nil {
		errRes := base.ErrorResponse{
			Status:  "error",
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return ctx.JSON(http.StatusInternalServerError, errRes)
	}
	if plants == nil {
		res := base.SuccessResponse{
			Status:  "success",
			Message: "No data found",
			Data:    nil,
		}
		return ctx.JSON(http.StatusOK, res)
	}
	res := base.SuccessResponse{
		Status:  "success",
		Message: "Data found",
		Data:    plants,
	}
	return ctx.JSON(http.StatusOK, res)
}
