package services_test

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/stretchr/testify/assert"

	"github.com/newtri-science/synonym-tool/model"
	"github.com/newtri-science/synonym-tool/services"
	"github.com/newtri-science/synonym-tool/test_utils"
)

var DB *sql.DB
var f = model.Food{
	Name: "TestFood",
	GeneralCategory: "TestCategory",
	RetentionCategory:"TestRetentionCategory",
	IndexCategory: "TestIndexCategory",
	Kilocalories: 1.0,
	Kilojoules: 1.0,
	Water: 1.0,
	Protein: 1.0,
	Fat: 1.0,
	Carbohydrates: 1.0,
	DietaryFiber: 1.0,
	Minerals: 1.0,
	OrganicAcids: 1.0,
	Alcohol: 1.0,
	RetinolActivityEquivalent: 1.0,
	RetinolEquivalent: 1.0,
	Retinol: 1.0,
	BetaCaroteneEquivalent: 1.0,
	BetaCarotene: 1.0,
	Calciferols: 1.0,
	AlphaTocopherolEquivalent: 1.0,
	AlphaTocopherol: 1.0,
	Phylloquinone: 1.0,
	Thiamine: 1.0,
	Riboflavin: 1.0,
	Niacin: 1.0,
	NiacinEquivalent: 1.0,
	PantothenicAcid: 1.0,
	Pyridoxine: 1.0,
	Biotin: 1.0,
	FolicAcid: 1.0,
	Cobalamin: 1.0,
	AscorbicAcid: 1.0,
	Sodium: 1.0,
	Potassium: 1.0,
	Calcium: 1.0,
	Magnesium: 1.0,
	Phosphorus: 1.0,
	Sulfur: 1.0,
	Chloride: 1.0,
	Iron: 1.0,
	Zinc: 1.0,
	Copper: 1.0,
	Manganese: 1.0,
	Fluoride: 1.0,
	Iodide: 1.0,
	Selenium: 1.0,
	Mannitol: 1.0,
	Sorbitol: 1.0,
	Xylitol: 1.0,
	SugarAlcohols: 1.0,
	Glucose: 1.0,
	Fructose: 1.0,
	Galactose: 1.0,
	Monosaccharides: 1.0,
	Sucrose: 1.0,
	Maltose: 1.0,
	Lactose: 1.0,
	Disaccharides: 1.0,
	TotalSugar: 1.0,
	ResorbableOligosaccharides: 1.0,
	NonResorbableOligosaccharides: 1.0,
	Glycogen: 1.0,
	Starch: 1.0,
	Polysaccharides: 1.0,
	Polypentoses: 1.0,
	Polyhexoses: 1.0,
	PolyuronicAcid: 1.0,
	Cellulose: 1.0,
	Lignin: 1.0,
	WaterSolubleDietaryFiber: 1.0,
	WaterInsolubleDietaryFiber: 1.0,
	Isoleucine: 1.0,
	Leucine: 1.0,
	Lysine: 1.0,
	Methionine: 1.0,
	Cysteine: 1.0,
	Phenylalanine: 1.0,
	Tyrosine: 1.0,
	Threonine: 1.0,
	Tryptophan: 1.0,
	Valine: 1.0,
	Arginine: 1.0,
	Histidine: 1.0,
	EssentialAminoAcids: 1.0,
	Alanine: 1.0,
	AsparticAcid: 1.0,
	GlutamicAcid: 1.0,
	Glycine: 1.0,
	Proline: 1.0,
	Serine: 1.0,
	NonEssentialAminoAcids: 1.0,
	UricAcid: 1.0,
	Purine: 1.0,
	ButyricAcid: 1.0,
	HexanoicAcid: 1.0,
	OctanoicAcid: 1.0,
	DecanoicAcid: 1.0,
	DodecanoicAcid: 1.0,
	TetradecanoicAcid: 1.0,
	PentadecanoicAcid: 1.0,
	HexadecanoicAcid: 1.0,
	HeptadecanoicAcid: 1.0,
	OctadecanoicAcid: 1.0,
	EicosanoicAcid: 1.0,
	DecosanoicAcid: 1.0,
	TetracosanoicAcid: 1.0,
	SaturatedFattyAcids: 1.0,
	TetradecenoicAcid: 1.0,
	PentadecenoicAcid: 1.0,
	HexadecenoicAcid: 1.0,
	HeptadecenoicAcid: 1.0,
	OctadecenoicAcid: 1.0,
	EicosenoicAcid: 1.0,
	DecosenoicAcid: 1.0,
	TetracosenoicAcid: 1.0,
	MonounsaturatedFattyAcids: 1.0,
	HexadecadienoicAcid: 1.0,
	HexadecatetraenoicAcid: 1.0,
	OctadecadienoicAcid: 1.0,
	OctadecatrienoicAcid: 1.0,
	OctadecatetraenoicAcid: 1.0,
	NonadecatrienoicAcid: 1.0,
	EicosadienoicAcid: 1.0,
	EicosatrienoicAcid: 1.0,
	EicosatetraenoicAcid: 1.0,
	EicosapentaenoicAcid: 1.0,
	DocosadienoicAcid: 1.0,
	DocosatrienoicAcid: 1.0,
	DocosatetraenoicAcid: 1.0,
	DocosapentaenoicAcid: 1.0,
	DocosahexaenoicAcid: 1.0,
	PolyunsaturatedFattyAcids: 1.0,
	ShortChainFattyAcids: 1.0,
	MediumChainFattyAcids: 1.0,
	LongChainFattyAcids: 1.0,
	Omega3FattyAcids: 1.0,
	Omega6FattyAcids: 1.0,
	GlycerolAndLipids: 1.0,
	Cholesterol: 1.0,
	Salt: 1.0,
}

func TestMain(m *testing.M) {
	// Setup test environment
	ctx := context.Background()
	testDb := test_utils.CreateTestContainer(ctx)
	container := testDb.Container

	DB = testDb.Db

	// Run the actual tests
	exitCode := m.Run()

	// Perform tear down
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	// Exit with the exit code from the tests
	os.Exit(exitCode)

}

func TestGetByName(t *testing.T) {
	repo := services.NewFoodEntryService(DB, nil)
	food, err := repo.GetByName("Erdbeere")
	assert.NoError(t, err)
	assert.NotNil(t, food)
}

func TestUserWithNameNotFound(t *testing.T) {
	repo := services.NewFoodEntryService(DB, nil)
	food, err := repo.GetByName("foo")
	assert.Nil(t, food)
	assert.NotNil(t, err)
}


func TestGetAllFoods(t *testing.T) {
	repo := services.NewFoodEntryService(DB, nil)
	expectedSize, err := repo.Count()
	if err != nil {
		t.Errorf("Error while trying to count foods: %s", err)
	}

	foods, err := repo.GetAllFoodEntries()
	if err != nil {
		t.Errorf("Error while trying to get all foods: %s", err)

	}

	actualSize := len(foods)
	if actualSize != expectedSize {
		t.Errorf("actual size %v is not expectedSize %v", actualSize, expectedSize)
	}
}

func TestAddFoodEntries(t *testing.T) {
	repo := services.NewFoodEntryService(DB, nil)
	beforeSize, err := repo.Count()
	assert.NoError(t, err)

	food, err := repo.AddFoodEntries([]model.Food{f})
	assert.NoError(t, err)
	assert.NotNil(t, food)

	afterSize, err := repo.Count()
	assert.NoError(t, err)
	assert.Equal(t, beforeSize+1, afterSize)
}

func TestAddFoodEntry(t *testing.T) {
	repo := services.NewFoodEntryService(DB, nil)
	beforeSize, err := repo.Count()
	assert.NoError(t, err)

	
	food, err := repo.AddFoodEntry(f)
	assert.NoError(t, err)
	assert.NotNil(t, food)

	afterSize, err := repo.Count()
	assert.NoError(t, err)
	assert.Equal(t, beforeSize+1, afterSize)
}



func TestCountFoods(t *testing.T) {
	repo := services.NewFoodEntryService(DB, nil)
	count, err := repo.Count()
	if err != nil {
		t.Errorf("Error while trying to count foods: %s", err)
	}
	if count == 0 {
		t.Errorf("No foods found")
	}
}
