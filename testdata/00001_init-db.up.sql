------------------------------------------------------------
--                        Enums                           --
------------------------------------------------------------
--CREATE TYPE IF NOT EXISTS user_roles AS ENUM ('admin', 'athlete');
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_roles') THEN
        CREATE TYPE user_roles AS ENUM ('admin', 'athlete');
    END IF;
END
$$;

-- CREATE TYPE user_status AS ENUM ('active', 'inactive');
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_status') THEN
        CREATE TYPE user_status AS ENUM ('active', 'inactive');
    END IF;
END
$$;

------------------------------------------------------------
--                   Users Table                          --
------------------------------------------------------------
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) NOT NULL UNIQUE,
  firstname VARCHAR(50),
  lastname VARCHAR(50),
  date_of_birth DATE,
  password_hash VARCHAR(100),
  status user_status,
  role user_roles,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (email, firstname, lastname, date_of_birth, password_hash, status, role, created_at, updated_at)
VALUES
  ('admin@example.com', 'admin', 'admin', '1990-01-01', 'hash123', 'active', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('test@test.de', 'admin', 'admin', '1990-01-01', 'hash123', 'active', 'admin', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('jan@ullrich.de', 'Jan', 'Ullrich', '1973-12-02', 'hash456', 'inactive', 'athlete', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
  ('mathieu@van-der-pole.be', 'Mathieu', 'van der Poel', '1995-01-19', 'hash456', 'active', 'athlete', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

------------------------------------------------------------
--                   Sessions Table                       --
------------------------------------------------------------
CREATE TABLE IF NOT EXISTS sessions (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  user_id INT REFERENCES users(id) ON DELETE CASCADE 
);


------------------------------------------------------------
--               Global Settings Table                    --
------------------------------------------------------------
CREATE TABLE IF NOT EXISTS globalSettings
(
    SectionName VARCHAR(50),
    SettingName VARCHAR(50),
    SettingValue VARCHAR(1000),
    SettingType SMALLINT DEFAULT 1,

    PRIMARY KEY (SectionName, SettingName)
);

------------------------------------------------------------
--                     Food Table                         --
------------------------------------------------------------
CREATE TABLE foods (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    general_category VARCHAR(50) NOT NULL,
    retention_category VARCHAR(50) NOT NULL,
    index_category VARCHAR(50) NOT NULL,
    kilocalories FLOAT NOT NULL,
    kilojoules FLOAT NOT NULL,
    water FLOAT NOT NULL,
    protein FLOAT NOT NULL,
    fat FLOAT NOT NULL,
    carbohydrates FLOAT NOT NULL,
    dietary_fiber FLOAT NOT NULL,
    minerals FLOAT NOT NULL,
    organic_acids FLOAT NOT NULL,
    alcohol FLOAT NOT NULL,
    retinol_activity_equivalent FLOAT NOT NULL,
    retinol_equivalent FLOAT NOT NULL,
    retinol FLOAT NOT NULL,
    beta_carotene_equivalent FLOAT NOT NULL,
    beta_carotene FLOAT NOT NULL,
    calciferols FLOAT NOT NULL,
    alpha_tocopherol_equivalent FLOAT NOT NULL,
    alpha_tocopherol FLOAT NOT NULL,
    phylloquinone FLOAT NOT NULL,
    thiamine FLOAT NOT NULL,
    riboflavin FLOAT NOT NULL,
    niacin FLOAT NOT NULL,
    niacin_equivalent FLOAT NOT NULL,
    pantothenic_acid FLOAT NOT NULL,
    pyridoxine FLOAT NOT NULL,
    biotin FLOAT NOT NULL,
    folic_acid FLOAT NOT NULL,
    cobalamin FLOAT NOT NULL,
    ascorbic_acid FLOAT NOT NULL,
    sodium FLOAT NOT NULL,
    potassium FLOAT NOT NULL,
    calcium FLOAT NOT NULL,
    magnesium FLOAT NOT NULL,
    phosphorus FLOAT NOT NULL,
    sulfur FLOAT NOT NULL,
    chloride FLOAT NOT NULL,
    iron FLOAT NOT NULL,
    zinc FLOAT NOT NULL,
    copper FLOAT NOT NULL,
    manganese FLOAT NOT NULL,
    fluoride FLOAT NOT NULL,
    iodide FLOAT NOT NULL,
    selenium FLOAT NOT NULL,
    mannitol FLOAT NOT NULL,
    sorbitol FLOAT NOT NULL,
    xylitol FLOAT NOT NULL,
    sugar_alcohols FLOAT NOT NULL,
    glucose FLOAT NOT NULL,
    fructose FLOAT NOT NULL,
    galactose FLOAT NOT NULL,
    monosaccharides FLOAT NOT NULL,
    sucrose FLOAT NOT NULL,
    maltose FLOAT NOT NULL,
    lactose FLOAT NOT NULL,
    disaccharides FLOAT NOT NULL,
    total_sugar FLOAT NOT NULL,
    resorbable_oligosaccharides FLOAT NOT NULL,
    non_resorbable_oligosaccharides FLOAT NOT NULL,
    glycogen FLOAT NOT NULL,
    starch FLOAT NOT NULL,
    polysaccharides FLOAT NOT NULL,
    polypentoses FLOAT NOT NULL,
    polyhexoses FLOAT NOT NULL,
    polyuronic_acid FLOAT NOT NULL,
    cellulose FLOAT NOT NULL,
    lignin FLOAT NOT NULL,
    water_soluble_dietary_fiber FLOAT NOT NULL,
    water_insoluble_dietary_fiber FLOAT NOT NULL,
    isoleucine FLOAT NOT NULL,
    leucine FLOAT NOT NULL,
    lysine FLOAT NOT NULL,
    methionine FLOAT NOT NULL,
    cysteine FLOAT NOT NULL,
    phenylalanine FLOAT NOT NULL,
    tyrosine FLOAT NOT NULL,
    threonine FLOAT NOT NULL,
    tryptophan FLOAT NOT NULL,
    valine FLOAT NOT NULL,
    arginine FLOAT NOT NULL,
    histidine FLOAT NOT NULL,
    essential_amino_acids FLOAT NOT NULL,
    alanine FLOAT NOT NULL,
    aspartic_acid FLOAT NOT NULL,
    glutamic_acid FLOAT NOT NULL,
    glycine FLOAT NOT NULL,
    proline FLOAT NOT NULL,
    serine FLOAT NOT NULL,
    non_essential_amino_acids FLOAT NOT NULL,
    uric_acid FLOAT NOT NULL,
    purine FLOAT NOT NULL,
    butyric_acid FLOAT NOT NULL,
    hexanoic_acid FLOAT NOT NULL,
    octanoic_acid FLOAT NOT NULL,
    decanoic_acid FLOAT NOT NULL,
    dodecanoic_acid FLOAT NOT NULL,
    tetradecanoic_acid FLOAT NOT NULL,
    pentadecanoic_acid FLOAT NOT NULL,
    hexadecanoic_acid FLOAT NOT NULL,
    heptadecanoic_acid FLOAT NOT NULL,
    octadecanoic_acid FLOAT NOT NULL,
    eicosanoic_acid FLOAT NOT NULL,
    decosanoic_acid FLOAT NOT NULL,
    tetracosanoic_acid FLOAT NOT NULL,
    saturated_fatty_acids FLOAT NOT NULL,
    tetradecenoic_acid FLOAT NOT NULL,
    pentadecenoic_acid FLOAT NOT NULL,
    hexadecenoic_acid FLOAT NOT NULL,
    heptadecenoic_acid FLOAT NOT NULL,
    octadecenoic_acid FLOAT NOT NULL,
    eicosenoic_acid FLOAT NOT NULL,
    decosenoic_acid FLOAT NOT NULL,
    tetracosenoic_acid FLOAT NOT NULL,
    monounsaturated_fatty_acids FLOAT NOT NULL,
    hexadecadienoic_acid FLOAT NOT NULL,
    hexadecatetraenoic_acid FLOAT NOT NULL,
    octadecadienoic_acid FLOAT NOT NULL,
    octadecatrienoic_acid FLOAT NOT NULL,
    octadecatetraenoic_acid FLOAT NOT NULL,
    nonadecatrienoic_acid FLOAT NOT NULL,
    eicosadienoic_acid FLOAT NOT NULL,
    eicosatrienoic_acid FLOAT NOT NULL,
    eicosatetraenoic_acid FLOAT NOT NULL,
    eicosapentaenoic_acid FLOAT NOT NULL,
    docosadienoic_acid FLOAT NOT NULL,
    docosatrienoic_acid FLOAT NOT NULL,
    docosatetraenoic_acid FLOAT NOT NULL,
    docosapentaenoic_acid FLOAT NOT NULL,
    docosahexaenoic_acid FLOAT NOT NULL,
    polyunsaturated_fatty_acids FLOAT NOT NULL,
    short_chain_fatty_acids FLOAT NOT NULL,
    medium_chain_fatty_acids FLOAT NOT NULL,
    long_chain_fatty_acids FLOAT NOT NULL,
    omega_3_fatty_acids FLOAT NOT NULL,
    omega_6_fatty_acids FLOAT NOT NULL,
    glycerol_and_lipids FLOAT NOT NULL,
    cholesterol FLOAT NOT NULL,
    salt FLOAT NOT NULL
);

-- Example data with created_at and updated_at fields
INSERT INTO foods (created_at, updated_at, name, general_category, retention_category, index_category, kilocalories, kilojoules, water, protein, fat, carbohydrates, dietary_fiber, minerals, organic_acids, alcohol, retinol_activity_equivalent, retinol_equivalent, retinol, beta_carotene_equivalent, beta_carotene, calciferols, alpha_tocopherol_equivalent, alpha_tocopherol, phylloquinone, thiamine, riboflavin, niacin, niacin_equivalent, pantothenic_acid, pyridoxine, biotin, folic_acid, cobalamin, ascorbic_acid, sodium, potassium, calcium, magnesium, phosphorus, sulfur, chloride, iron, zinc, copper, manganese, fluoride, iodide, selenium, mannitol, sorbitol, xylitol, sugar_alcohols, glucose, fructose, galactose, monosaccharides, sucrose, maltose, lactose, disaccharides, total_sugar, resorbable_oligosaccharides, non_resorbable_oligosaccharides, glycogen, starch, polysaccharides, polypentoses, polyhexoses, polyuronic_acid, cellulose, lignin, water_soluble_dietary_fiber, water_insoluble_dietary_fiber, isoleucine, leucine, lysine, methionine, cysteine, phenylalanine, tyrosine, threonine, tryptophan, valine, arginine, histidine, essential_amino_acids, alanine, aspartic_acid, glutamic_acid, glycine, proline, serine, non_essential_amino_acids, uric_acid, purine, butyric_acid, hexanoic_acid, octanoic_acid, decanoic_acid, dodecanoic_acid, tetradecanoic_acid, pentadecanoic_acid, hexadecanoic_acid, heptadecanoic_acid, octadecanoic_acid, eicosanoic_acid, decosanoic_acid, tetracosanoic_acid, saturated_fatty_acids, tetradecenoic_acid, pentadecenoic_acid, hexadecenoic_acid, heptadecenoic_acid, octadecenoic_acid, eicosenoic_acid, decosenoic_acid, tetracosenoic_acid, monounsaturated_fatty_acids, hexadecadienoic_acid, hexadecatetraenoic_acid, octadecadienoic_acid, octadecatrienoic_acid, octadecatetraenoic_acid, nonadecatrienoic_acid, eicosadienoic_acid, eicosatrienoic_acid, eicosatetraenoic_acid, eicosapentaenoic_acid, docosadienoic_acid, docosatrienoic_acid, docosatetraenoic_acid, docosapentaenoic_acid, docosahexaenoic_acid, polyunsaturated_fatty_acids, short_chain_fatty_acids, medium_chain_fatty_acids, long_chain_fatty_acids, omega_3_fatty_acids, omega_6_fatty_acids, glycerol_and_lipids, cholesterol, salt)
VALUES
  (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Boysenbeere', 'whole_fruit', 'whole_fruit', 'whole_fruit', 37.0, 157.0, 83650.0, 500.0, 300.0, 6950.0, 6200.0, 900.0, 1500.0, 0.0, 0.0, 17.0, 0.0, 0.0, 100.0, 0.0, 400.0, 400.0, 0.0, 20.0, 130.0, 1000.0, 1100.0, 240.0, 60.0, 1.0, 12.0, 0.0, 13000.0, 3.0, 150.0, 25.0, 18.0, 24.0, 16.0, 14.0, 1600.0, 300.0, 140.0, 500.0, 24.0, 0.7, 0.0, 0.0, 0.0, 0.0, 0.0, 2480.0, 3700.0, 0.0, 6180.0, 770.0, 0.0, 0.0, 770.0, 6950.0, 0.0, 0.0, 0.0, 0.0, 0.0, 1916.0, 1159.0, 738.0, 1122.0, 1265.0, 1593.0, 4607.0, 13.0, 29.0, 16.0, 10.0, 6.0, 19.0, 13.0, 19.0, 6.0, 23.0, 35.0, 16.0, 205.0, 29.0, 92.0, 101.0, 25.0, 23.0, 25.0, 295.0, 15.0, 5.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 13.0, 0.0, 5.0, 0.0, 0.0, 0.0, 18.0, 0.0, 0.0, 1.0, 0.0, 36.0, 0.0, 0.0, 0.0, 37.0, 0.0, 0.0, 108.0, 77.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 185.0, 0.0, 0.0, 240.0, 77.0, 108.0, 60.0, 0.0, 8.0),
  (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Brombeere', 'whole_fruit', 'whole_fruit', 'whole_fruit', 41.5, 180.625, 86125.0, 1150.0, 850.0, 6370.0, 4200.0, 510.0, 1700.0, 0.0, 9.0, 32.0, 0.0, 113.0, 184.0, 0.0, 940.0, 600.0, 0.0, 20.0, 35.0, 435.0, 650.0, 230.0, 35.0, 0.4, 25.05, 0.0, 13850.0, 1.5, 195.0, 37.5, 25.0, 27.5, 12.0, 20.0, 650.0, 195.0, 100.0, 970.0, 24.0, 0.4, 0.0, 0.0, 3214.25, 0.0, 3214.25, 1999.5, 2102.75, 0.0, 4102.25, 114.5, 0.0, 0.0, 114.5, 5158.375, 0.0, 0.0, 0.0, 0.0, 0.0, 1088.0, 787.0, 486.0, 582.0, 256.0, 960.0, 2200.0, 31.0, 69.0, 38.0, 23.0, 15.0, 46.0, 31.0, 46.0, 15.0, 53.0, 84.0, 38.0, 489.0, 69.0, 222.0, 245.0, 61.0, 54.0, 61.0, 712.0, 15.0, 5.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 44.0, 0.0, 16.0, 0.0, 0.0, 0.0, 80.0, 0.0, 0.0, 4.0, 0.0, 120.0, 0.0, 0.0, 0.0, 162.0, 0.0, 0.0, 360.0, 256.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 508.0, 0.0, 0.0, 800.0, 256.0, 360.0, 200.0, 0.0, 2.5),
  (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Clementine', 'whole_fruit', 'whole_fruit', 'whole_fruit', 46.0, 192.0, 86102.0, 699.0, 299.0, 9000.0, 2000.0, 700.0, 1200.0, 0.0, 0.0, 50.0, 0.0, 0.0, 300.0, 0.0, 300.0, 300.0, 3.0, 70.0, 20.0, 200.0, 350.0, 200.0, 50.0, 0.5, 15.0, 0.0, 30000.0, 2.0, 180.0, 35.0, 11.0, 20.0, 10.0, 3.0, 300.0, 100.0, 90.0, 40.0, 10.0, 0.8, 0.0, 0.0, 0.0, 0.0, 0.0, 1530.0, 1692.0, 0.0, 3222.0, 5778.0, 0.0, 0.0, 5778.0, 9000.0, 0.0, 0.0, 0.0, 0.0, 0.0, 400.0, 300.0, 600.0, 400.0, 300.0, 788.0, 1212.0, 22.0, 22.0, 41.0, 13.0, 9.0, 28.0, 13.0, 13.0, 9.0, 32.0, 59.0, 13.0, 274.0, 46.0, 129.0, 92.0, 74.0, 50.0, 32.0, 423.0, 20.0, 7.0, 0.0, 0.0, 0.0, 0.0, 0.0, 2.0, 0.0, 57.0, 0.0, 2.0, 0.0, 0.0, 0.0, 61.0, 0.0, 0.0, 7.0, 0.0, 48.0, 0.0, 0.0, 0.0, 55.0, 0.0, 0.0, 87.0, 36.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 123.0, 0.0, 0.0, 239.0, 36.0, 87.0, 60.0, 0.0, 5.0),
  (CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'Erdbeere', 'whole_fruit', 'whole_fruit', 'whole_fruit', 35.5, 148.5, 90035.0, 710.0, 450.0, 5755.0, 2900.0, 500.0, 1000.0, 0.0, 0.0, 2.0, 0.0, 3.0, 9.5, 0.0, 210.0, 120.0, 5.0, 25.5, 37.0, 440.0, 784.75, 210.0, 50.0, 4.0, 42.35, 0.0, 56400.0, 1.0, 152.0, 18.5, 12.5, 24.0, 13.0, 35.5, 421.0, 99.0, 46.0, 400.0, 16.0, 1.9, 0.5, 0.0, 32.0, 28.0, 60.0, 2181.0, 2250.75, 0.0, 4431.75, 1008.0, 0.0, 0.0, 1008.0, 5519.875, 0.0, 0.0, 0.0, 0.0, 0.0, 456.0, 304.0, 810.0, 330.0, 100.0, 580.0, 1050.0, 20.5, 47.75, 37.0, 1.0, 7.75, 27.25, 31.25, 28.25, 16.5, 27.25, 40.0, 17.5, 302.0, 47.75, 208.25, 137.25, 37.0, 29.25, 36.0, 495.5, 21.0, 7.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 26.25, 0.0, 5.0, 0.0, 0.0, 0.0, 31.0, 0.0, 0.0, 1.0, 0.0, 60.5, 0.0, 0.0, 0.0, 61.5, 0.0, 0.0, 130.0, 102.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 232.0, 0.0, 0.0, 324.5, 102.0, 130.0, 78.5, 0.0, 1.5);

  -- Create triggers to enforce lowercase on insert and update
CREATE OR REPLACE FUNCTION lowercase_setting_name()
RETURNS TRIGGER AS $$
BEGIN
    NEW.SectionName := LOWER(NEW.SectionName);
    NEW.SettingName := LOWER(NEW.SettingName);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS lowercase_setting_name_trigger ON globalSettings;
CREATE TRIGGER lowercase_setting_name_trigger
BEFORE INSERT OR UPDATE ON globalSettings
FOR EACH ROW
EXECUTE FUNCTION lowercase_setting_name();

INSERT INTO globalSettings (SectionName, SettingName, SettingValue, SettingType) VALUES ('APP', 'initialized', 'false', 2);
