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
			&food.GeneralCategory,
			&food.RetentionCategory,
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
			&food.GeneralCategory,
			&food.RetentionCategory,
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

func (s *FoodEntryService) AddFoodEntries(food []model.Food) ([]model.Food, error) {
	for _, f := range food {
		_, err := s.AddFoodEntry(f)
		if err != nil {
			return nil, err
		}
	}

	return food, nil
}

func (s *FoodEntryService) AddFoodEntry(food model.Food) (*model.Food, error) {
	_, err := s.db.Exec(
		"INSERT INTO foods (name, general_category, retention_category, index_category, kilocalories, kilojoules, water, protein, fat, carbohydrates, dietary_fiber, minerals, organic_acids, alcohol, retinol_activity_equivalent, retinol_equivalent, retinol, beta_carotene_equivalent, beta_carotene, calciferols, alpha_tocopherol_equivalent, alpha_tocopherol, phylloquinone, thiamine, riboflavin, niacin, niacin_equivalent, pantothenic_acid, pyridoxine, biotin, folic_acid, cobalamin, ascorbic_acid, sodium, potassium, calcium, magnesium, phosphorus, sulfur, chloride, iron, zinc, copper, manganese, fluoride, iodide, selenium, mannitol, sorbitol, xylitol, sugar_alcohols, glucose, fructose, galactose, monosaccharides, sucrose, maltose, lactose, disaccharides, total_sugar, resorbable_oligosaccharides, non_resorbable_oligosaccharides, glycogen, starch, polysaccharides, polypentoses, polyhexoses, polyuronic_acid, cellulose, lignin, water_soluble_dietary_fiber, water_insoluble_dietary_fiber, isoleucine, leucine, lysine, methionine, cysteine, phenylalanine, tyrosine, threonine, tryptophan, valine, arginine, histidine, essential_amino_acids, alanine, aspartic_acid, glutamic_acid, glycine, proline, serine, non_essential_amino_acids, uric_acid, purine, butyric_acid, hexanoic_acid, octanoic_acid, decanoic_acid, dodecanoic_acid, tetradecanoic_acid, pentadecanoic_acid, hexadecanoic_acid, heptadecanoic_acid, octadecanoic_acid, eicosanoic_acid, decosanoic_acid, tetracosanoic_acid, saturated_fatty_acids, tetradecenoic_acid, pentadecenoic_acid, hexadecenoic_acid, heptadecenoic_acid, octadecenoic_acid, eicosenoic_acid, decosenoic_acid, tetracosenoic_acid, monounsaturated_fatty_acids, hexadecadienoic_acid, hexadecatetraenoic_acid, octadecadienoic_acid, octadecatrienoic_acid, octadecatetraenoic_acid, nonadecatrienoic_acid, eicosadienoic_acid, eicosatrienoic_acid, eicosatetraenoic_acid, eicosapentaenoic_acid, docosadienoic_acid, docosatrienoic_acid, docosatetraenoic_acid, docosapentaenoic_acid, docosahexaenoic_acid, polyunsaturated_fatty_acids, short_chain_fatty_acids, medium_chain_fatty_acids, long_chain_fatty_acids, omega_3_fatty_acids, omega_6_fatty_acids, glycerol_and_lipids, cholesterol, salt) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63, $64, $65, $66, $67, $68, $69, $70, $71, $72, $73, $74, $75, $76, $77, $78, $79, $80, $81, $82, $83, $84, $85, $86, $87, $88, $89, $90, $91, $92, $93, $94, $95, $96, $97, $98, $99, $100, $101, $102, $103, $104, $105, $106, $107, $108, $109, $110, $111, $112, $113, $114, $115, $116, $117, $118, $119, $120, $121, $122, $123, $124, $125, $126, $127, $128, $129, $130, $131, $132, $133, $134, $135, $136, $137, $138, $139, $140, $141) " +
		"ON CONFLICT (name) " +
		"DO UPDATE SET " +
		"name = $1," +
		"updated_at = CURRENT_TIMESTAMP, " +
		"general_category = $2, " +
		"retention_category = $3, " +
		"index_category = $4, " +
		"kilocalories = $5, " +
		"kilojoules = $6, " +
		"water = $7, " +
		"protein = $8, " +
		"fat = $9, " +
		"carbohydrates = $10, " +
		"dietary_fiber = $11, " +
		"minerals = $12, " +
		"organic_acids = $13, " +
		"alcohol = $14, " +
		"retinol_activity_equivalent = $15, " +
		"retinol_equivalent = $16, " +
		"retinol = $17, " +
		"beta_carotene_equivalent = $18, " +
		"beta_carotene = $19, " +
		"calciferols = $20, " +
		"alpha_tocopherol_equivalent = $21, " +
		"alpha_tocopherol = $22, " +
		"phylloquinone = $23, " +
		"thiamine = $24, " +
		"riboflavin = $25, " +
		"niacin = $26, " +
		"niacin_equivalent = $27, " +
		"pantothenic_acid = $28, " +
		"pyridoxine = $29, " +
		"biotin = $30, " +
		"folic_acid = $31, " +
		"cobalamin = $32, " +
		"ascorbic_acid = $33, " +
		"sodium = $34, " +
		"potassium = $35, " +
		"calcium = $36, " +
		"magnesium = $37, " +
		"phosphorus = $38, " +
		"sulfur = $39, " +
		"chloride = $40, " +
		"iron = $41, " +
		"zinc = $42, " +
		"copper = $43, " + 
		"manganese = $44, " + 
		"fluoride = $45, " + 
		"iodide = $46, " + 
		"selenium = $47, " + 
		"mannitol = $48, " + 
		"sorbitol = $49, " + 
		"xylitol = $50, " + 
		"sugar_alcohols = $51, " + 
		"glucose = $52, " + 
		"fructose = $53, " + 
		"galactose = $54, " + 
		"monosaccharides = $55, " + 
		"sucrose = $56, " + 
		"maltose = $57, " + 
		"lactose = $58, " + 
		"disaccharides = $59, " + 
		"total_sugar = $60, " + 
		"resorbable_oligosaccharides = $61, " + 
		"non_resorbable_oligosaccharides = $62, " + 
		"glycogen = $63, " + 
		"starch = $64, " + 
		"polysaccharides = $65, " + 
		"polypentoses = $66, " + 
		"polyhexoses = $67, " + 
		"polyuronic_acid = $68, " + 
		"cellulose = $69, " + 
		"lignin = $70, " + 
		"water_soluble_dietary_fiber = $71, " + 
		"water_insoluble_dietary_fiber = $72, " + 
		"isoleucine = $73, " + 
		"leucine = $74, " + 
		"lysine = $75, " + 
		"methionine = $76, " + 
		"cysteine = $77, " + 
		"phenylalanine = $78, " + 
		"tyrosine = $79, " + 
		"threonine = $80, " + 
		"tryptophan = $81, " + 
		"valine = $82, " + 
		"arginine = $83, " + 
		"histidine = $84, " + 
		"essential_amino_acids = $85, " + 
		"alanine = $86, " + 
		"aspartic_acid = $87, " + 
		"glutamic_acid = $88, " + 
		"glycine = $89, " + 
		"proline = $90, " + 
		"serine = $91, " + 
		"non_essential_amino_acids = $92, " + 
		"uric_acid = $93, " + 
		"purine = $94, " + 
		"butyric_acid = $95, " + 
		"hexanoic_acid = $96, " + 
		"octanoic_acid = $97, " + 
		"decanoic_acid = $98, " + 
		"dodecanoic_acid = $99, " + 
		"tetradecanoic_acid = $100, " + 
		"pentadecanoic_acid = $101, " + 
		"hexadecanoic_acid = $102, " + 
		"heptadecanoic_acid = $103, " + 
		"octadecanoic_acid = $104, " + 
		"eicosanoic_acid = $105, " + 
		"decosanoic_acid = $106, " + 
		"tetracosanoic_acid = $107, " + 
		"saturated_fatty_acids = $108, " + 
		"tetradecenoic_acid = $109, " + 
		"pentadecenoic_acid = $110, " + 
		"hexadecenoic_acid = $111, " + 
		"heptadecenoic_acid = $112, " + 
		"octadecenoic_acid = $113, " + 
		"eicosenoic_acid = $114, " + 
		"decosenoic_acid = $115, " + 
		"tetracosenoic_acid = $116, " + 
		"monounsaturated_fatty_acids = $117, " + 
		"hexadecadienoic_acid = $118, " + 
		"hexadecatetraenoic_acid = $119, " + 
		"octadecadienoic_acid = $120, " + 
		"octadecatrienoic_acid = $121, " + 
		"octadecatetraenoic_acid = $122, " + 
		"nonadecatrienoic_acid = $123, " + 
		"eicosadienoic_acid = $124, " + 
		"eicosatrienoic_acid = $125, " + 
		"eicosatetraenoic_acid = $126, " + 
		"eicosapentaenoic_acid = $127, " + 
		"docosadienoic_acid = $128, " + 
		"docosatrienoic_acid = $129, " + 
		"docosatetraenoic_acid = $130, " + 
		"docosapentaenoic_acid = $131, " + 
		"docosahexaenoic_acid = $132, " + 
		"polyunsaturated_fatty_acids = $133, " + 
		"short_chain_fatty_acids = $134, " + 
		"medium_chain_fatty_acids = $135, " + 
		"long_chain_fatty_acids = $136, " + 
		"omega_3_fatty_acids = $137, " + 
		"omega_6_fatty_acids = $138, " + 
		"glycerol_and_lipids = $139, " + 
		"cholesterol = $140, " + 
		"salt = $141 ",
		food.Name,
		food.GeneralCategory,
		food.RetentionCategory,
		food.IndexCategory,
		food.Kilocalories,
		food.Kilojoules,
		food.Water,
		food.Protein,
		food.Fat,
		food.Carbohydrates,
		food.DietaryFiber,
		food.Minerals,
		food.OrganicAcids,
		food.Alcohol,
		food.RetinolActivityEquivalent,
		food.RetinolEquivalent,
		food.Retinol,
		food.BetaCaroteneEquivalent,
		food.BetaCarotene,
		food.Calciferols,
		food.AlphaTocopherolEquivalent,
		food.AlphaTocopherol,
		food.Phylloquinone,
		food.Thiamine,
		food.Riboflavin,
		food.Niacin,
		food.NiacinEquivalent,
		food.PantothenicAcid,
		food.Pyridoxine,
		food.Biotin,
		food.FolicAcid,
		food.Cobalamin,
		food.AscorbicAcid,
		food.Sodium,
		food.Potassium,
		food.Calcium,
		food.Magnesium,
		food.Phosphorus,
		food.Sulfur,
		food.Chloride,
		food.Iron,
		food.Zinc,
		food.Copper,
		food.Manganese,
		food.Fluoride,
		food.Iodide,
		food.Selenium,
		food.Mannitol,
		food.Sorbitol,
		food.Xylitol,
		food.SugarAlcohols,
		food.Glucose,
		food.Fructose,
		food.Galactose,
		food.Monosaccharides,
		food.Sucrose,
		food.Maltose,
		food.Lactose,
		food.Disaccharides,
		food.TotalSugar,
		food.ResorbableOligosaccharides,
		food.NonResorbableOligosaccharides,
		food.Glycogen,
		food.Starch,
		food.Polysaccharides,
		food.Polypentoses,
		food.Polyhexoses,
		food.PolyuronicAcid,
		food.Cellulose,
		food.Lignin,
		food.WaterSolubleDietaryFiber,
		food.WaterInsolubleDietaryFiber,
		food.Isoleucine,
		food.Leucine,
		food.Lysine,
		food.Methionine,
		food.Cysteine,
		food.Phenylalanine,
		food.Tyrosine,
		food.Threonine,
		food.Tryptophan,
		food.Valine,
		food.Arginine,
		food.Histidine,
		food.EssentialAminoAcids,
		food.Alanine,
		food.AsparticAcid,
		food.GlutamicAcid,
		food.Glycine,
		food.Proline,
		food.Serine,
		food.NonEssentialAminoAcids,
		food.UricAcid,
		food.Purine,
		food.ButyricAcid,
		food.HexanoicAcid,
		food.OctanoicAcid,
		food.DecanoicAcid,
		food.DodecanoicAcid,
		food.TetradecanoicAcid,
		food.PentadecanoicAcid,
		food.HexadecanoicAcid,
		food.HeptadecanoicAcid,
		food.OctadecanoicAcid,
		food.EicosanoicAcid,
		food.DecosanoicAcid,
		food.TetracosanoicAcid,
		food.SaturatedFattyAcids,
		food.TetradecenoicAcid,
		food.PentadecenoicAcid,
		food.HexadecenoicAcid,
		food.HeptadecenoicAcid,
		food.OctadecenoicAcid,
		food.EicosenoicAcid,
		food.DecosenoicAcid,
		food.TetracosenoicAcid,
		food.MonounsaturatedFattyAcids,
		food.HexadecadienoicAcid,
		food.HexadecatetraenoicAcid,
		food.OctadecadienoicAcid,
		food.OctadecatrienoicAcid,
		food.OctadecatetraenoicAcid,
		food.NonadecatrienoicAcid,
		food.EicosadienoicAcid,
		food.EicosatrienoicAcid,
		food.EicosatetraenoicAcid,
		food.EicosapentaenoicAcid,
		food.DocosadienoicAcid,
		food.DocosatrienoicAcid,
		food.DocosatetraenoicAcid,
		food.DocosapentaenoicAcid,
		food.DocosahexaenoicAcid,
		food.PolyunsaturatedFattyAcids,
		food.ShortChainFattyAcids,
		food.MediumChainFattyAcids,
		food.LongChainFattyAcids,
		food.Omega3FattyAcids,
		food.Omega6FattyAcids,
		food.GlycerolAndLipids,
		food.Cholesterol,
		food.Salt,
	)
	if err != nil {
		return nil, err
	}
	return &food, nil
} 

// TODO: Delete and Update food entries (zb doppelte entries?!)

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
