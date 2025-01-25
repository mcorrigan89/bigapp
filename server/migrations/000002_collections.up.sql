CREATE TABLE IF NOT EXISTS collections (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  collection_name TEXT NOT NULL,
  owner_id UUID NOT NULL REFERENCES users(id),
  public BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS collection_images (
  collection_id UUID NOT NULL REFERENCES collections(id),
  image_id UUID NOT NULL REFERENCES images(id),
  PRIMARY KEY (collection_id, image_id),
  sort_key TEXT NOT NULL
);

ALTER TABLE images ADD COLUMN IF NOT EXISTS owner_id UUID REFERENCES users(id);