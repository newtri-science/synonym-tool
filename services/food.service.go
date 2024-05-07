package services

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	"github.com/newtri-science/synonym-tool/model"
)

type FoodEntryService struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewFoodEntryService(db *sql.DB, logger *zap.SugaredLogger) *FoodEntryService {
	return &FoodEntryService{
		db:     db,
		logger: logger,
	}
}

// TODO: get by custom ID

func (s *FoodEntryService) GetByName(name string) (*model.Food, error) {
	row := s.db.QueryRow("SELECT * FROM foods WHERE LOWER(foods.name) = LOWER($1)", name)

	var food model.Food
		err :=  row.Scan(
			&food.ID,
			&food.CreatedAt,
			&food.UpdatedAt,
			&food.Name,
			&food.RetentionCategory,
			&food.GeneralCategory,
			&food.IndexCategory,
			&food.Kilocalories,
			&food.Kilojoules,
			&food.Water,
			&food.Protein,
			&food.Fat,
			&food.Carbohydrates,
			&food.DietaryFiber,
			&food.Minerals,
			&food.OrganicAcids,
			&food.Alcohol,
			&food.RetinolActivityEquivalent,
			&food.RetinolEquivalent,
			&food.Retinol,
			&food.BetaCaroteneEquivalent,
			&food.BetaCarotene,
			&food.Calciferols,
			&food.AlphaTocopherolEquivalent,
			&food.AlphaTocopherol,
			&food.Phylloquinone,
			&food.Thiamine,
			&food.Riboflavin,
			&food.Niacin,
			&food.NiacinEquivalent,
			&food.PantothenicAcid,
			&food.Pyridoxine,
			&food.Biotin,
			&food.FolicAcid,
			&food.Cobalamin,
			&food.AscorbicAcid,
			&food.Sodium,
			&food.Potassium,
			&food.Calcium,
			&food.Magnesium,
			&food.Phosphorus,
			&food.Sulfur,
			&food.Chloride,
			&food.Iron,
			&food.Zinc,
			&food.Copper,
			&food.Manganese,
			&food.Fluoride,
			&food.Iodide,
			&food.Selenium,
			&food.Mannitol,
			&food.Sorbitol,
			&food.Xylitol,
			&food.SugarAlcohols,
			&food.Glucose,
			&food.Fructose,
			&food.Galactose,
			&food.Monosaccharides,
			&food.Sucrose,
			&food.Maltose,
			&food.Lactose,
			&food.Disaccharides,
			&food.TotalSugar,
			&food.ResorbableOligosaccharides,
			&food.NonResorbableOligosaccharides,
			&food.Glycogen,
			&food.Starch,
			&food.Polysaccharides,
			&food.Polypentoses,
			&food.Polyhexoses,
			&food.PolyuronicAcid,
			&food.Cellulose,
			&food.Lignin,
			&food.WaterSolubleDietaryFiber,
			&food.WaterInsolubleDietaryFiber,
			&food.Isoleucine,
			&food.Leucine,
			&food.Lysine,
			&food.Methionine,
			&food.Cysteine,
			&food.Phenylalanine,
			&food.Tyrosine,
			&food.Threonine,
			&food.Tryptophan,
			&food.Valine,
			&food.Arginine,
			&food.Histidine,
			&food.EssentialAminoAcids,
			&food.Alanine,
			&food.AsparticAcid,
			&food.GlutamicAcid,
			&food.Glycine,
			&food.Proline,
			&food.Serine,
			&food.NonEssentialAminoAcids,
			&food.UricAcid,
			&food.Purine,
			&food.ButyricAcid,
			&food.HexanoicAcid,
			&food.OctanoicAcid,
			&food.DecanoicAcid,
			&food.DodecanoicAcid,
			&food.TetradecanoicAcid,
			&food.PentadecanoicAcid,
			&food.HexadecanoicAcid,
			&food.HeptadecanoicAcid,
			&food.OctadecanoicAcid,
			&food.EicosanoicAcid,
			&food.DecosanoicAcid,
			&food.TetracosanoicAcid,
			&food.SaturatedFattyAcids,
			&food.TetradecenoicAcid,
			&food.PentadecenoicAcid,
			&food.HexadecenoicAcid,
			&food.HeptadecenoicAcid,
			&food.OctadecenoicAcid,
			&food.EicosenoicAcid,
			&food.DecosenoicAcid,
			&food.TetracosenoicAcid,
			&food.MonounsaturatedFattyAcids,
			&food.HexadecadienoicAcid,
			&food.HexadecatetraenoicAcid,
			&food.OctadecadienoicAcid,
			&food.OctadecatrienoicAcid,
			&food.OctadecatetraenoicAcid,
			&food.NonadecatrienoicAcid,
			&food.EicosadienoicAcid,
			&food.EicosatrienoicAcid,
			&food.EicosatetraenoicAcid,
			&food.EicosapentaenoicAcid,
			&food.DocosadienoicAcid,
			&food.DocosatrienoicAcid,
			&food.DocosatetraenoicAcid,
			&food.DocosapentaenoicAcid,
			&food.DocosahexaenoicAcid,
			&food.PolyunsaturatedFattyAcids,
			&food.ShortChainFattyAcids,
			&food.MediumChainFattyAcids,
			&food.LongChainFattyAcids,
			&food.Omega3FattyAcids,
			&food.Omega6FattyAcids,
			&food.GlycerolAndLipids,
			&food.Cholesterol,
			&food.Salt,
		)
	
	if err != nil {
		return nil, err
	}
	
	return &food, nil
}

func (s *FoodEntryService) GetAllFoodEntries() ([]*model.Food, error) {
	rows, err := s.db.Query("SELECT * FROM foods")
	if err != nil {
		return nil, fmt.Errorf("error while trying to execute query: %s", err)
	}

	var foods []*model.Food
	for rows.Next() {
		var food model.Food
		err :=  rows.Scan(
			&food.ID,
			&food.CreatedAt,
			&food.UpdatedAt,
			&food.Name,
			&food.RetentionCategory,
			&food.GeneralCategory,
			&food.IndexCategory,
			&food.Kilocalories,
			&food.Kilojoules,
			&food.Water,
			&food.Protein,
			&food.Fat,
			&food.Carbohydrates,
			&food.DietaryFiber,
			&food.Minerals,
			&food.OrganicAcids,
			&food.Alcohol,
			&food.RetinolActivityEquivalent,
			&food.RetinolEquivalent,
			&food.Retinol,
			&food.BetaCaroteneEquivalent,
			&food.BetaCarotene,
			&food.Calciferols,
			&food.AlphaTocopherolEquivalent,
			&food.AlphaTocopherol,
			&food.Phylloquinone,
			&food.Thiamine,
			&food.Riboflavin,
			&food.Niacin,
			&food.NiacinEquivalent,
			&food.PantothenicAcid,
			&food.Pyridoxine,
			&food.Biotin,
			&food.FolicAcid,
			&food.Cobalamin,
			&food.AscorbicAcid,
			&food.Sodium,
			&food.Potassium,
			&food.Calcium,
			&food.Magnesium,
			&food.Phosphorus,
			&food.Sulfur,
			&food.Chloride,
			&food.Iron,
			&food.Zinc,
			&food.Copper,
			&food.Manganese,
			&food.Fluoride,
			&food.Iodide,
			&food.Selenium,
			&food.Mannitol,
			&food.Sorbitol,
			&food.Xylitol,
			&food.SugarAlcohols,
			&food.Glucose,
			&food.Fructose,
			&food.Galactose,
			&food.Monosaccharides,
			&food.Sucrose,
			&food.Maltose,
			&food.Lactose,
			&food.Disaccharides,
			&food.TotalSugar,
			&food.ResorbableOligosaccharides,
			&food.NonResorbableOligosaccharides,
			&food.Glycogen,
			&food.Starch,
			&food.Polysaccharides,
			&food.Polypentoses,
			&food.Polyhexoses,
			&food.PolyuronicAcid,
			&food.Cellulose,
			&food.Lignin,
			&food.WaterSolubleDietaryFiber,
			&food.WaterInsolubleDietaryFiber,
			&food.Isoleucine,
			&food.Leucine,
			&food.Lysine,
			&food.Methionine,
			&food.Cysteine,
			&food.Phenylalanine,
			&food.Tyrosine,
			&food.Threonine,
			&food.Tryptophan,
			&food.Valine,
			&food.Arginine,
			&food.Histidine,
			&food.EssentialAminoAcids,
			&food.Alanine,
			&food.AsparticAcid,
			&food.GlutamicAcid,
			&food.Glycine,
			&food.Proline,
			&food.Serine,
			&food.NonEssentialAminoAcids,
			&food.UricAcid,
			&food.Purine,
			&food.ButyricAcid,
			&food.HexanoicAcid,
			&food.OctanoicAcid,
			&food.DecanoicAcid,
			&food.DodecanoicAcid,
			&food.TetradecanoicAcid,
			&food.PentadecanoicAcid,
			&food.HexadecanoicAcid,
			&food.HeptadecanoicAcid,
			&food.OctadecanoicAcid,
			&food.EicosanoicAcid,
			&food.DecosanoicAcid,
			&food.TetracosanoicAcid,
			&food.SaturatedFattyAcids,
			&food.TetradecenoicAcid,
			&food.PentadecenoicAcid,
			&food.HexadecenoicAcid,
			&food.HeptadecenoicAcid,
			&food.OctadecenoicAcid,
			&food.EicosenoicAcid,
			&food.DecosenoicAcid,
			&food.TetracosenoicAcid,
			&food.MonounsaturatedFattyAcids,
			&food.HexadecadienoicAcid,
			&food.HexadecatetraenoicAcid,
			&food.OctadecadienoicAcid,
			&food.OctadecatrienoicAcid,
			&food.OctadecatetraenoicAcid,
			&food.NonadecatrienoicAcid,
			&food.EicosadienoicAcid,
			&food.EicosatrienoicAcid,
			&food.EicosatetraenoicAcid,
			&food.EicosapentaenoicAcid,
			&food.DocosadienoicAcid,
			&food.DocosatrienoicAcid,
			&food.DocosatetraenoicAcid,
			&food.DocosapentaenoicAcid,
			&food.DocosahexaenoicAcid,
			&food.PolyunsaturatedFattyAcids,
			&food.ShortChainFattyAcids,
			&food.MediumChainFattyAcids,
			&food.LongChainFattyAcids,
			&food.Omega3FattyAcids,
			&food.Omega6FattyAcids,
			&food.GlycerolAndLipids,
			&food.Cholesterol,
			&food.Salt,
		)

		if err != nil {
			return nil, fmt.Errorf("error while trying to execute query: %s", err)
		}
		
		foods = append(foods, &food)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while trying to execute query: %s", err)
	}

	defer rows.Close()
	return foods, nil
}

// TODO: Add, Delete and Update food entries

/**
 * Count returns the number of food entries in the database
 */
func (s *FoodEntryService) Count() (int, error) {
	row := s.db.QueryRow("SELECT count(*) FROM foods")
	var count int
	err := row.Scan(&count)
	if err != nil {
		return -1, fmt.Errorf("error while trying to execute query: %s", err)
	}
	return count, nil
}
