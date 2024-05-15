package scripts

import (
	"encoding/csv"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/newtri-science/synonym-tool/model"
)


func GenerateFoodEntries(file *multipart.FileHeader) ([]model.Food, error) {
	// test for .csv
    if !strings.HasSuffix(file.Filename, ".csv") {
		return nil, fmt.Errorf("file is no .csv file")
    }

	// open file
    src, err := file.Open()
    if err != nil {
        return nil, fmt.Errorf("error while opening the csv file: %s", err)
    }
    defer src.Close()

    // read file content
    content, err := io.ReadAll(src)
    if err != nil {
        return nil, fmt.Errorf("error while reading the csv stream: %s", err)
    }

    // read csv
    csvReader := csv.NewReader(strings.NewReader(string(content)))
	csvReader.Comma = ';'
    records, err := csvReader.ReadAll()
    if err != nil {
        return nil, fmt.Errorf("error while reading the csv file, please use SEMICOLON as separator: %s", err)
    }

	requiredColumns := []string{
     "name",
     "general_category",
     "retention_category",
     "index_category",
     "kilocalories",
     "kilojoules",
     "water",
     "protein",
     "fat",
     "carbohydrates",
     "dietary_fiber",
     "minerals",
     "organic_acids",
     "alcohol",
     "retinol_activity_equivalent",
     "retinol_equivalent",
     "retinol",
     "beta_carotene_equivalent",
     "beta_carotene",
     "calciferols",
     "alpha_tocopherol_equivalent",
     "alpha_tocopherol",
     "phylloquinone",
     "thiamine",
     "riboflavin",
     "niacin",
     "niacin_equivalent",
     "pantothenic_acid",
     "pyridoxine",
     "biotin",
     "folic_acid",
     "cobalamin",
     "ascorbic_acid",
     "sodium",
     "potassium",
     "calcium",
     "magnesium",
     "phosphorus",
     "sulfur",
     "chloride",
     "iron",
     "zinc",
     "copper",
     "manganese",
     "fluoride",
     "iodide",
     "selenium",
     "mannitol",
     "sorbitol",
     "xylitol",
     "sugar_alcohols",
     "glucose",
     "fructose",
     "galactose",
     "monosaccharides",
     "sucrose",
     "maltose",
     "lactose",
     "disaccharides",
     "total_sugar",
     "resorbable_oligosaccharides",
     "non_resorbable_oligosaccharides",
     "glycogen",
     "starch",
     "polysaccharides",
     "polypentoses",
     "polyhexoses",
     "polyuronic_acid",
     "cellulose",
     "lignin",
     "water_soluble_dietary_fiber",
     "water_insoluble_dietary_fiber",
     "isoleucine",
     "leucine",
     "lysine",
     "methionine",
     "cysteine",
     "phenylalanine",
     "tyrosine",
     "threonine",
     "tryptophan",
     "valine",
     "arginine",
     "histidine",
     "essential_amino_acids",
     "alanine",
     "aspartic_acid",
     "glutamic_acid",
     "glycine",
     "proline",
     "serine",
     "non_essential_amino_acids",
     "uric_acid",
     "purine",
     "butyric_acid",
     "hexanoic_acid",
     "octanoic_acid",
     "decanoic_acid",
     "dodecanoic_acid",
     "tetradecanoic_acid",
     "pentadecanoic_acid",
     "hexadecanoic_acid",
     "heptadecanoic_acid",
     "octadecanoic_acid",
     "eicosanoic_acid",
     "decosanoic_acid",
     "tetracosanoic_acid",
     "saturated_fatty_acids",
     "tetradecenoic_acid",
     "pentadecenoic_acid",
     "hexadecenoic_acid",
     "heptadecenoic_acid",
     "octadecenoic_acid",
     "eicosenoic_acid",
     "decosenoic_acid",
     "tetracosenoic_acid",
     "monounsaturated_fatty_acids",
     "hexadecadienoic_acid",
     "hexadecatetraenoic_acid",
     "octadecadienoic_acid",
     "octadecatrienoic_acid",
     "octadecatetraenoic_acid",
     "nonadecatrienoic_acid",
     "eicosadienoic_acid",
     "eicosatrienoic_acid",
     "eicosatetraenoic_acid",
     "eicosapentaenoic_acid",
     "docosadienoic_acid",
     "docosatrienoic_acid",
     "docosatetraenoic_acid",
     "docosapentaenoic_acid",
     "docosahexaenoic_acid",
     "polyunsaturated_fatty_acids",
     "short_chain_fatty_acids",
     "medium_chain_fatty_acids",
     "long_chain_fatty_acids",
     "omega_3_fatty_acids",
     "omega_6_fatty_acids",
     "glycerol_and_lipids",
     "cholesterol",
     "salt",
	}

	fmt.Println("start controlling csv columns for name & position")

	// Check if all required columns are present
	for i, col := range requiredColumns {
		found := false
		for j, headerCol := range records[0] {
			if strings.EqualFold(headerCol, col) {
				found = true
                if (i+1 != j) {
                    return nil, fmt.Errorf("column '%s' is not in the right order", col)
                }
				break
			}
		}
		if !found {
			return nil, fmt.Errorf("required column '%s' not found in CSV", col)
		}
	}

    // control number of columns, +1 because of the id column
    if len(records[0]) != len(requiredColumns) +1 {
        return nil, fmt.Errorf("number of columns in CSV is not equal to the number of required columns")
    }

    // Remove special characters from the entries
    fmt.Println("start remove special characters data")
    for i, row := range records {
        for j, entry := range row {
            records[i][j] = removeSpecialCharacters(entry)
        }
    }

    // TODO? ad own, unique id to each entry
    fmt.Println("TODO? add unique id to each entry")

    // Transform entries to []Model.Food
    fmt.Println("transform entries to []Model.Food")
    var foods []model.Food
    for _, row := range records[1:] {
        var food model.Food

        food.Name = row[1]
		food.GeneralCategory = row[2]
		food.RetentionCategory = row[3]
        food.IndexCategory = row[4]

        kilocalories, _ := strconv.ParseFloat(row[5], 32)
        food.Kilocalories = float32(kilocalories)

        kilojoules, _ := strconv.ParseFloat(row[6], 32)
        food.Kilojoules = float32(kilojoules)

        water, _ := strconv.ParseFloat(row[7], 32)
        food.Water = float32(water)

        protein, _ := strconv.ParseFloat(row[8], 32)
        food.Protein = float32(protein)

        fat, _ := strconv.ParseFloat(row[9], 32)
        food.Fat = float32(fat)

        carbohydrates, _ := strconv.ParseFloat(row[10], 32)
        food.Carbohydrates = float32(carbohydrates)

        dietaryFiber, _ := strconv.ParseFloat(row[11], 32)
        food.DietaryFiber = float32(dietaryFiber)

        minerals, _ := strconv.ParseFloat(row[12], 32)
        food.Minerals = float32(minerals)

        organicAcids, _ := strconv.ParseFloat(row[13], 32)
        food.OrganicAcids = float32(organicAcids)

        alcohol, _ := strconv.ParseFloat(row[14], 32)
        food.Alcohol = float32(alcohol)

        retinolActivityEquivalent, _ := strconv.ParseFloat(row[15], 32)
        food.RetinolActivityEquivalent = float32(retinolActivityEquivalent)

        retinolEquivalent, _ := strconv.ParseFloat(row[16], 32)
        food.RetinolEquivalent = float32(retinolEquivalent)

        retinol, _ := strconv.ParseFloat(row[17], 32)
        food.Retinol = float32(retinol)

        betaCaroteneEquivalent, _ := strconv.ParseFloat(row[18], 32)
        food.BetaCaroteneEquivalent = float32(betaCaroteneEquivalent)

        betaCarotene, _ := strconv.ParseFloat(row[19], 32)
        food.BetaCarotene = float32(betaCarotene)

        calciferols, _ := strconv.ParseFloat(row[20], 32)
        food.Calciferols = float32(calciferols)

        alphaTocopherolEquivalent, _ := strconv.ParseFloat(row[21], 32)
        food.AlphaTocopherolEquivalent = float32(alphaTocopherolEquivalent)

        alphaTocopherol, _ := strconv.ParseFloat(row[22], 32)
        food.AlphaTocopherol = float32(alphaTocopherol)

        phylloquinone, _ := strconv.ParseFloat(row[23], 32)
        food.Phylloquinone = float32(phylloquinone)

        thiamine, _ := strconv.ParseFloat(row[24], 32)
        food.Thiamine = float32(thiamine)

        riboflavin, _ := strconv.ParseFloat(row[25], 32)
        food.Riboflavin = float32(riboflavin)

        niacin, _ := strconv.ParseFloat(row[26], 32)
        food.Niacin = float32(niacin)

        niacinEquivalent, _ := strconv.ParseFloat(row[27], 32)
        food.NiacinEquivalent = float32(niacinEquivalent)

        pantothenicAcid, _ := strconv.ParseFloat(row[28], 32)
        food.PantothenicAcid = float32(pantothenicAcid)

        pyridoxine, _ := strconv.ParseFloat(row[29], 32)
        food.Pyridoxine = float32(pyridoxine)

        biotin, _ := strconv.ParseFloat(row[30], 32)
        food.Biotin = float32(biotin)

        folicAcid, _ := strconv.ParseFloat(row[31], 32)
        food.FolicAcid = float32(folicAcid)

        cobalamin, _ := strconv.ParseFloat(row[32], 32)
        food.Cobalamin = float32(cobalamin)

        ascorbicAcid, _ := strconv.ParseFloat(row[33], 32)
        food.AscorbicAcid = float32(ascorbicAcid)

        sodium, _ := strconv.ParseFloat(row[34], 32)
        food.Sodium = float32(sodium)

        potassium, _ := strconv.ParseFloat(row[35], 32)
        food.Potassium = float32(potassium)

        calcium, _ := strconv.ParseFloat(row[36], 32)
        food.Calcium = float32(calcium)

        magnesium, _ := strconv.ParseFloat(row[37], 32)
        food.Magnesium = float32(magnesium)

        phosphorus, _ := strconv.ParseFloat(row[38], 32)
        food.Phosphorus = float32(phosphorus)

        sulfur, _ := strconv.ParseFloat(row[39], 32)
        food.Sulfur = float32(sulfur)

        chloride, _ := strconv.ParseFloat(row[40], 32)
        food.Chloride = float32(chloride)

        iron, _ := strconv.ParseFloat(row[41], 32)
        food.Iron = float32(iron)

        zinc, _ := strconv.ParseFloat(row[42], 32)
        food.Zinc = float32(zinc)

        copper, _ := strconv.ParseFloat(row[43], 32)
        food.Copper = float32(copper)

        manganese, _ := strconv.ParseFloat(row[44], 32)
        food.Manganese = float32(manganese)

        fluoride, _ := strconv.ParseFloat(row[45], 32)
        food.Fluoride = float32(fluoride)

        iodide, _ := strconv.ParseFloat(row[46], 32)
        food.Iodide = float32(iodide)

        selenium, _ := strconv.ParseFloat(row[47], 32)
        food.Selenium = float32(selenium)

        mannitol, _ := strconv.ParseFloat(row[48], 32)
        food.Mannitol = float32(mannitol)

        sorbitol, _ := strconv.ParseFloat(row[49], 32)
        food.Sorbitol = float32(sorbitol)

        xylitol, _ := strconv.ParseFloat(row[50], 32)
        food.Xylitol = float32(xylitol)

        sugarAlcohols, _ := strconv.ParseFloat(row[51], 32)
        food.SugarAlcohols = float32(sugarAlcohols)

        glucose, _ := strconv.ParseFloat(row[52], 32)
        food.Glucose = float32(glucose)

        fructose, _ := strconv.ParseFloat(row[53], 32)
        food.Fructose = float32(fructose)

        galactose, _ := strconv.ParseFloat(row[54], 32)
        food.Galactose = float32(galactose)

        monosaccharides, _ := strconv.ParseFloat(row[55], 32)
        food.Monosaccharides = float32(monosaccharides)

        sucrose, _ := strconv.ParseFloat(row[56], 32)
        food.Sucrose = float32(sucrose)

        maltose, _ := strconv.ParseFloat(row[57], 32)
        food.Maltose = float32(maltose)

        lactose, _ := strconv.ParseFloat(row[58], 32)
        food.Lactose = float32(lactose)

        disaccharides, _ := strconv.ParseFloat(row[59], 32)
        food.Disaccharides = float32(disaccharides)

        totalSugar, _ := strconv.ParseFloat(row[60], 32)
        food.TotalSugar = float32(totalSugar)

        resorbableOligosaccharides, _ := strconv.ParseFloat(row[61], 32)
        food.ResorbableOligosaccharides = float32(resorbableOligosaccharides)

        nonResorbableOligosaccharides, _ := strconv.ParseFloat(row[62], 32)
        food.NonResorbableOligosaccharides = float32(nonResorbableOligosaccharides)


        glycogen, _ := strconv.ParseFloat(row[63], 32)
        food.Glycogen = float32(glycogen)

        starch, _ := strconv.ParseFloat(row[64], 32)
        food.Starch = float32(starch)

        polysaccharides, _ := strconv.ParseFloat(row[65], 32)
        food.Polysaccharides = float32(polysaccharides)

        polypentoses, _ := strconv.ParseFloat(row[66], 32)
        food.Polypentoses = float32(polypentoses)

        polyhexoses, _ := strconv.ParseFloat(row[67], 32)
        food.Polyhexoses = float32(polyhexoses)

        polyuronicAcid, _ := strconv.ParseFloat(row[68], 32)
        food.PolyuronicAcid = float32(polyuronicAcid)

        cellulose, _ := strconv.ParseFloat(row[69], 32)
        food.Cellulose = float32(cellulose)

        lignin, _ := strconv.ParseFloat(row[70], 32)
        food.Lignin = float32(lignin)

        waterSolubleDietaryFiber, _ := strconv.ParseFloat(row[71], 32)
        food.WaterSolubleDietaryFiber = float32(waterSolubleDietaryFiber)

        waterInsolubleDietaryFiber, _ := strconv.ParseFloat(row[72], 32)
        food.WaterInsolubleDietaryFiber = float32(waterInsolubleDietaryFiber)

        isoleucine, _ := strconv.ParseFloat(row[73], 32)
        food.Isoleucine = float32(isoleucine)

        leucine, _ := strconv.ParseFloat(row[74], 32)
        food.Leucine = float32(leucine)

        lysine, _ := strconv.ParseFloat(row[75], 32)
        food.Lysine = float32(lysine)

        methionine, _ := strconv.ParseFloat(row[76], 32)
        food.Methionine = float32(methionine)

        cysteine, _ := strconv.ParseFloat(row[77], 32)
        food.Cysteine = float32(cysteine)

        phenylalanine, _ := strconv.ParseFloat(row[78], 32)
        food.Phenylalanine = float32(phenylalanine)

        tyrosine, _ := strconv.ParseFloat(row[79], 32)
        food.Tyrosine = float32(tyrosine)

        threonine, _ := strconv.ParseFloat(row[80], 32)
        food.Threonine = float32(threonine)

        tryptophan, _ := strconv.ParseFloat(row[81], 32)
        food.Tryptophan = float32(tryptophan)

        valine, _ := strconv.ParseFloat(row[82], 32)
        food.Valine = float32(valine)

        arginine, _ := strconv.ParseFloat(row[83], 32)
        food.Arginine = float32(arginine)

        histidine, _ := strconv.ParseFloat(row[84], 32)
        food.Histidine = float32(histidine)

        essentialAminoAcids, _ := strconv.ParseFloat(row[85], 32)
        food.EssentialAminoAcids = float32(essentialAminoAcids)

        alanine, _ := strconv.ParseFloat(row[86], 32)
        food.Alanine = float32(alanine)

        asparticAcid, _ := strconv.ParseFloat(row[87], 32)
        food.AsparticAcid = float32(asparticAcid)

        glutamicAcid, _ := strconv.ParseFloat(row[88], 32)
        food.GlutamicAcid = float32(glutamicAcid)

        glycine, _ := strconv.ParseFloat(row[89], 32)
        food.Glycine = float32(glycine)

        proline, _ := strconv.ParseFloat(row[90], 32)
        food.Proline = float32(proline)

        serine, _ := strconv.ParseFloat(row[91], 32)
        food.Serine = float32(serine)

        nonEssentialAminoAcids, _ := strconv.ParseFloat(row[92], 32)
        food.NonEssentialAminoAcids = float32(nonEssentialAminoAcids)

        uricAcid, _ := strconv.ParseFloat(row[93], 32)
        food.UricAcid = float32(uricAcid)

        purine, _ := strconv.ParseFloat(row[94], 32)
        food.Purine = float32(purine)

        butyricAcid, _ := strconv.ParseFloat(row[95], 32)
        food.ButyricAcid = float32(butyricAcid)

        hexanoicAcid, _ := strconv.ParseFloat(row[96], 32)
        food.HexanoicAcid = float32(hexanoicAcid)

        octanoicAcid, _ := strconv.ParseFloat(row[97], 32)
        food.OctanoicAcid = float32(octanoicAcid)

        decanoicAcid, _ := strconv.ParseFloat(row[98], 32)
        food.DecanoicAcid = float32(decanoicAcid)

        dodecanoicAcid, _ := strconv.ParseFloat(row[99], 32)
        food.DodecanoicAcid = float32(dodecanoicAcid)

        tetradecanoicAcid, _ := strconv.ParseFloat(row[100], 32)
        food.TetradecanoicAcid = float32(tetradecanoicAcid)

        pentadecanoicAcid, _ := strconv.ParseFloat(row[101], 32)
        food.PentadecanoicAcid = float32(pentadecanoicAcid)

        hexadecanoicAcid, _ := strconv.ParseFloat(row[102], 32)
        food.HexadecanoicAcid = float32(hexadecanoicAcid)

        heptadecanoicAcid, _ := strconv.ParseFloat(row[103], 32)
        food.HeptadecanoicAcid = float32(heptadecanoicAcid)

        octadecanoicAcid, _ := strconv.ParseFloat(row[104], 32)
        food.OctadecanoicAcid = float32(octadecanoicAcid)

        eicosanoicAcid, _ := strconv.ParseFloat(row[105], 32)
        food.EicosanoicAcid = float32(eicosanoicAcid)

        decosanoicAcid, _ := strconv.ParseFloat(row[106], 32)
        food.DecosanoicAcid = float32(decosanoicAcid)

        tetracosanoicAcid, _ := strconv.ParseFloat(row[107], 32)
        food.TetracosanoicAcid = float32(tetracosanoicAcid)

        saturatedFattyAcids, _ := strconv.ParseFloat(row[108], 32)
        food.SaturatedFattyAcids = float32(saturatedFattyAcids)

        tetradecenoicAcid, _ := strconv.ParseFloat(row[109], 32)
        food.TetradecenoicAcid = float32(tetradecenoicAcid)

        pentadecenoicAcid, _ := strconv.ParseFloat(row[110], 32)
        food.PentadecenoicAcid = float32(pentadecenoicAcid)

        hexadecenoicAcid, _ := strconv.ParseFloat(row[111], 32)
        food.HexadecenoicAcid = float32(hexadecenoicAcid)

        heptadecenoicAcid, _ := strconv.ParseFloat(row[112], 32)
        food.HeptadecenoicAcid = float32(heptadecenoicAcid)

        octadecenoicAcid, _ := strconv.ParseFloat(row[113], 32)
        food.OctadecenoicAcid = float32(octadecenoicAcid)

        eicosenoicAcid, _ := strconv.ParseFloat(row[114], 32)
        food.EicosenoicAcid = float32(eicosenoicAcid)

        decosenoicAcid, _ := strconv.ParseFloat(row[115], 32)
        food.DecosenoicAcid = float32(decosenoicAcid)

        tetracosenoicAcid, _ := strconv.ParseFloat(row[116], 32)
        food.TetracosenoicAcid = float32(tetracosenoicAcid)

        monounsaturatedFattyAcids, _ := strconv.ParseFloat(row[117], 32)
        food.MonounsaturatedFattyAcids = float32(monounsaturatedFattyAcids)

        hexadecadienoicAcid, _ := strconv.ParseFloat(row[118], 32)
        food.HexadecadienoicAcid = float32(hexadecadienoicAcid)

        hexadecatetraenoicAcid, _ := strconv.ParseFloat(row[119], 32)
        food.HexadecatetraenoicAcid = float32(hexadecatetraenoicAcid)

        octadecadienoicAcid, _ := strconv.ParseFloat(row[120], 32)
        food.OctadecadienoicAcid = float32(octadecadienoicAcid)

        octadecatrienoicAcid, _ := strconv.ParseFloat(row[121], 32)
        food.OctadecatrienoicAcid = float32(octadecatrienoicAcid)

        octadecatetraenoicAcid, _ := strconv.ParseFloat(row[122], 32)
        food.OctadecatetraenoicAcid = float32(octadecatetraenoicAcid)

        nonadecatrienoicAcid, _ := strconv.ParseFloat(row[123], 32)
        food.NonadecatrienoicAcid = float32(nonadecatrienoicAcid)

        eicosadienoicAcid, _ := strconv.ParseFloat(row[124], 32)
        food.EicosadienoicAcid = float32(eicosadienoicAcid)

        eicosatrienoicAcid, _ := strconv.ParseFloat(row[125], 32)
        food.EicosatrienoicAcid = float32(eicosatrienoicAcid)

        eicosatetraenoicAcid, _ := strconv.ParseFloat(row[126], 32)
        food.EicosatetraenoicAcid = float32(eicosatetraenoicAcid)

        eicosapentaenoicAcid, _ := strconv.ParseFloat(row[127], 32)
        food.EicosapentaenoicAcid = float32(eicosapentaenoicAcid)

        docosadienoicAcid, _ := strconv.ParseFloat(row[128], 32)
        food.DocosadienoicAcid = float32(docosadienoicAcid)

        docosatrienoicAcid, _ := strconv.ParseFloat(row[129], 32)
        food.DocosatrienoicAcid = float32(docosatrienoicAcid)

        docosatetraenoicAcid, _ := strconv.ParseFloat(row[130], 32)
        food.DocosatetraenoicAcid = float32(docosatetraenoicAcid)

        docosapentaenoicAcid, _ := strconv.ParseFloat(row[131], 32)
        food.DocosapentaenoicAcid = float32(docosapentaenoicAcid)

        docosahexaenoicAcid, _ := strconv.ParseFloat(row[132], 32)
        food.DocosahexaenoicAcid = float32(docosahexaenoicAcid)

        polyunsaturatedFattyAcids, _ := strconv.ParseFloat(row[133], 32)
        food.PolyunsaturatedFattyAcids = float32(polyunsaturatedFattyAcids)

        shortChainFattyAcids, _ := strconv.ParseFloat(row[134], 32)
        food.ShortChainFattyAcids = float32(shortChainFattyAcids)

        mediumChainFattyAcids, _ := strconv.ParseFloat(row[135], 32)
        food.MediumChainFattyAcids = float32(mediumChainFattyAcids)

        longChainFattyAcids, _ := strconv.ParseFloat(row[136], 32)
        food.LongChainFattyAcids = float32(longChainFattyAcids)

        omega3FattyAcids, _ := strconv.ParseFloat(row[137], 32)
        food.Omega3FattyAcids = float32(omega3FattyAcids)

        omega6FattyAcids, _ := strconv.ParseFloat(row[138], 32)
        food.Omega6FattyAcids = float32(omega6FattyAcids)

        glycerolAndLipids, _ := strconv.ParseFloat(row[139], 32)
        food.GlycerolAndLipids = float32(glycerolAndLipids)

        cholesterol, _ := strconv.ParseFloat(row[140], 32)
        food.Cholesterol = float32(cholesterol)

        salt, _ := strconv.ParseFloat(row[141], 32)
        food.Salt = float32(salt)

        foods = append(foods, food)
        break // TODO
    }
	
    fmt.Println(foods)

	return foods, nil
}

func removeSpecialCharacters(entry string) string {
	charsToRemove := []string{`"`, "'", "\\"}

	// Replace each character with an empty string
	for _, char := range charsToRemove {
		entry = strings.ReplaceAll(entry, char, "")
	}

	return entry
}

