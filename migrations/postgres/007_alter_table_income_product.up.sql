ALTER TABLE IF EXISTS income_products
    ADD COLUMN branch_id UUID REFERENCES branches(id) NOT NULL DEFAULT 'aa541fcc-bf74-11ee-ae0b-166244b65504';

