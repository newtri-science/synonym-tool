package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/newtri-science/synonym-tool/model"
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

func (h FoodEntryHandler) ListFoodPage(c echo.Context) error {
	foodEntries, err := h.repo.GetAllFoodEntries()

	if err != nil {
		fmt.Println("error when looking for all foodEntries:" + err.Error())
	}
	return Render(c, food_table.Index(foodEntries), http.StatusOK)
}

func (h FoodEntryHandler) ListFoodEntries(c echo.Context) error {
	param := c.QueryParam("name")
	var foodEntries []*model.Food
    var err error

	h.logger.Infof("Looking for foodEntries with name: %s", param)

	if param == "" || param == "all" || param == "*" || param == " " {
		foodEntries, err = h.repo.GetAllFoodEntries()
	} else {
		var foodEntry *model.Food
		foodEntry, err = h.repo.GetByName(param)
		foodEntries = []*model.Food{foodEntry}
	}

	if err != nil {
		fmt.Println("error when looking for all foodEntries:" + err.Error())
	}
	return Render(c, food_table.FoodTable(foodEntries), http.StatusOK)
}

// TODO: Add, Delete and Update FoodEntry