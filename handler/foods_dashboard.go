package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/newtri-science/synonym-tool/services"
	food_table "github.com/newtri-science/synonym-tool/views/foods"
)

type FoodEntryHandler struct {
	repo   *services.FoodEntryService
	logger *zap.SugaredLogger
}

func NewFoodEntryHandler(
	repo *services.FoodEntryService,
	logger *zap.SugaredLogger,
) FoodEntryHandler {
	return FoodEntryHandler{repo: repo, logger: logger}
}

func (h FoodEntryHandler) ListFoodEntries(c echo.Context) error {
	foodEntries, err := h.repo.GetAllFoodEntries()
	if err != nil {
		fmt.Println("error when looking for all foodEntries:" + err.Error())
	}
	return Render(c, food_table.Index(foodEntries), http.StatusOK)
}

// TODO: Add, Delete and Update FoodEntry