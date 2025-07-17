-- campaigns table
CREATE TABLE IF NOT EXISTS campaigns (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    image TEXT NOT NULL,
    cta TEXT NOT NULL,
    status TEXT NOT NULL CHECK (status IN ('ACTIVE', 'INACTIVE'))
);

-- targeting_rules table
CREATE TABLE IF NOT EXISTS targeting_rules (
    id UUID PRIMARY KEY,
    campaign_id UUID NOT NULL REFERENCES campaigns(id) ON DELETE CASCADE,
    dimension TEXT NOT NULL CHECK (dimension IN ('app', 'country', 'os')),
    rule_type TEXT NOT NULL CHECK (rule_type IN ('INCLUDE', 'EXCLUDE')),
    values TEXT[] NOT NULL
);

-- Indexes to speed up read queries
CREATE INDEX IF NOT EXISTS idx_campaign_status ON campaigns(status);
CREATE INDEX IF NOT EXISTS idx_rules_campaign_id ON targeting_rules(campaign_id);


-- Campaigns
INSERT INTO campaigns (id, name, image, cta, status)
VALUES 
  ('8f5bfc02-bcc1-41fc-9238-ecbe8392d2a2', 'spotify', 'https://image.spotify.com', 'Listen Now', 'ACTIVE'),
  ('44e6c308-44c9-4b60-b83d-fc93f4fc11a6', 'duolingo', 'https://image.duolingo.com', 'Start Learning', 'ACTIVE'),
  ('6c01b5c6-90a9-45aa-b3d4-bde219624702', 'subwaysurfer', 'https://image.subwaysurfer.com', 'Play Now', 'ACTIVE');

-- Targeting Rules
INSERT INTO targeting_rules (id, campaign_id, dimension, rule_type, values)
VALUES 
  (gen_random_uuid(), '8f5bfc02-bcc1-41fc-9238-ecbe8392d2a2', 'country', 'INCLUDE', ARRAY['US', 'Canada']),
  (gen_random_uuid(), '44e6c308-44c9-4b60-b83d-fc93f4fc11a6', 'os', 'INCLUDE', ARRAY['Android', 'iOS']),
  (gen_random_uuid(), '44e6c308-44c9-4b60-b83d-fc93f4fc11a6', 'country', 'EXCLUDE', ARRAY['US']),
  (gen_random_uuid(), '6c01b5c6-90a9-45aa-b3d4-bde219624702', 'os', 'INCLUDE', ARRAY['Android']),
  (gen_random_uuid(), '6c01b5c6-90a9-45aa-b3d4-bde219624702', 'app', 'INCLUDE', ARRAY['com.gametion.ludokinggame']);


-- Enable UUID generation if not already
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
