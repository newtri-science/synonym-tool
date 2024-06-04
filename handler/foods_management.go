package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/newtri-science/synonym-tool/model"
	"github.com/newtri-science/synonym-tool/scripts"
	"github.com/newtri-science/synonym-tool/services"
	"github.com/newtri-science/synonym-tool/utils"
	"github.com/newtri-science/synonym-tool/views/pages"
)

type FoodManagementHandler struct {
	s   *services.FoodEntryService
	logger *zap.SugaredLogger
}

func NewFoodManagementHandler(
	repo *services.FoodEntryService,
	logger *zap.SugaredLogger,
) FoodManagementHandler {
	return FoodManagementHandler{s: repo, logger: logger}
}

func (h FoodManagementHandler) RenderFoodManagementPage(c echo.Context) error {
	au := c.(model.AuthenticatedContext).User
	foodEntries, err := h.s.GetAllFoodEntries()

	if err != nil {
		utils.Warning("Could not retrieve food entries")
	}
	
	return Render(c, pages.FoodManagementPage(au, GetTheme(c), foodEntries))
}

func (h FoodManagementHandler) RenderFoodManagementView(c echo.Context) error {
	au := c.(model.AuthenticatedContext).User
	foodEntries, err := h.s.GetAllFoodEntries()
	
	if err != nil {
		return utils.Warning("Could not retrieve food entries")
	}
	return Render(c, pages.FoodManagementView(au, foodEntries))
}

func (h FoodManagementHandler) ListFoodEntries(c echo.Context) error {
	param := c.QueryParam("name")
	var foodEntries []*model.Food
    var err error

	h.logger.Infof("Looking for foodEntries with name: %s", param)

	if param == "" || param == "all" || param == "*" || param == " " {
		foodEntries, err = h.s.GetAllFoodEntries()
	} else {
		var foodEntry *model.Food
		foodEntry, err = h.s.GetByName(param)
		foodEntries = []*model.Food{foodEntry}
	}

	if err != nil {
		fmt.Println("error when looking for all foodEntries:" + err.Error())
	}
	return Render(c, pages.FoodTable(foodEntries))
}

func (h FoodManagementHandler) UploadFoodEntries(c echo.Context) error {
	file, err := c.FormFile("food-entries-file")
	if err != nil {
        fmt.Println("no file provided for food entries upload: " + err.Error())
        return err
    }

	foodEntries, err := scripts.GenerateFoodEntries(file)

	if err != nil {
		return err
	}
	
	if _, err := h.s.AddFoodEntries(foodEntries); err != nil {
		fmt.Println("error when adding food entries: " + err.Error())
		return err
	}
	
	allFoodEntries, err := h.s.GetAllFoodEntries()
	if err != nil {
		fmt.Println("error when looking for all foodEntries:" + err.Error())
	}
	return Render(c, pages.FoodTable(allFoodEntries))
}

// TODO: Add DELETE Entry