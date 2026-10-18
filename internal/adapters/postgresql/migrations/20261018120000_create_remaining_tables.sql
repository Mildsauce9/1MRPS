-- +goose Up

-- lookup tables: id is the primary key, nothing references out of them
CREATE TABLE IF NOT EXISTS category_mapping (
    category INTEGER PRIMARY KEY,
    category_description TEXT
);

CREATE TABLE IF NOT EXISTS genres (
    id BIGINT PRIMARY KEY,
    created_at DATE,
    name TEXT,
    slug TEXT,
    updated_at DATE,
    url TEXT,
    checksum TEXT
);

CREATE TABLE IF NOT EXISTS themes (
    id BIGINT PRIMARY KEY,
    created_at DATE,
    name TEXT,
    slug TEXT,
    updated_at DATE,
    url TEXT,
    checksum TEXT
);

CREATE TABLE IF NOT EXISTS player_perspectives (
    id BIGINT PRIMARY KEY,
    created_at DATE,
    name TEXT,
    slug TEXT,
    updated_at DATE,
    url TEXT,
    checksum TEXT
);

CREATE TABLE IF NOT EXISTS platforms (
    id BIGINT PRIMARY KEY,
    alternative_name TEXT,
    category INTEGER,
    created_at TIMESTAMP,
    name TEXT,
    platform_logo BIGINT,
    slug TEXT,
    updated_at TIMESTAMP,
    url TEXT,
    versions JSONB,
    websites JSONB,
    checksum TEXT,
    generation INTEGER,
    platform_family BIGINT,
    abbreviation TEXT,
    summary TEXT
);

CREATE TABLE IF NOT EXISTS age_ratings (
    id BIGINT PRIMARY KEY,
    category INTEGER,
    rating_system TEXT,
    rating INTEGER,
    rating_description TEXT,
    age_group TEXT,
    checksum TEXT,
    content_descriptions JSONB, -- list of age_rating_content_descriptions.id
    synopsis TEXT
);

CREATE TABLE IF NOT EXISTS age_rating_content_descriptions (
    csv_index INTEGER,
    id BIGINT PRIMARY KEY,
    category INTEGER,
    description TEXT,
    checksum TEXT
);

-- companies
CREATE TABLE IF NOT EXISTS companies (
    csv_index INTEGER,
    id BIGINT PRIMARY KEY,
    change_date_category INTEGER,
    country INTEGER,
    created_at BIGINT,
    description TEXT,
    developed JSONB, -- list of games.id
    logo BIGINT,
    name TEXT,
    slug TEXT,
    start_date BIGINT,
    start_date_category INTEGER,
    updated_at BIGINT,
    url TEXT,
    websites JSONB,
    checksum TEXT,
    published JSONB, -- list of games.id
    parent BIGINT REFERENCES companies (id),
    change_date BIGINT,
    changed_company_id BIGINT REFERENCES companies (id)
);

-- join table between games and companies
CREATE TABLE IF NOT EXISTS involved_companies (
    csv_index INTEGER,
    id BIGINT PRIMARY KEY,
    company BIGINT REFERENCES companies (id),
    created_at BIGINT,
    developer BOOLEAN,
    game BIGINT REFERENCES games (id),
    porting BOOLEAN,
    publisher BOOLEAN,
    supporting BOOLEAN,
    updated_at BIGINT,
    checksum TEXT
);

CREATE INDEX IF NOT EXISTS idx_involved_companies_company ON involved_companies (company);
CREATE INDEX IF NOT EXISTS idx_involved_companies_game ON involved_companies (game);
CREATE INDEX IF NOT EXISTS idx_companies_parent ON companies (parent);

-- +goose Down

DROP TABLE IF EXISTS involved_companies;
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS age_rating_content_descriptions;
DROP TABLE IF EXISTS age_ratings;
DROP TABLE IF EXISTS platforms;
DROP TABLE IF EXISTS player_perspectives;
DROP TABLE IF EXISTS themes;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS category_mapping;
