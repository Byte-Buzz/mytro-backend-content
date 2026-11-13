-- ENUMs
CREATE TYPE tile_type AS ENUM ('image', 'video');
CREATE TYPE tile_visibility AS ENUM ('public', 'private', 'unlisted');
CREATE TYPE tile_status AS ENUM ('draft', 'processing', 'ready', 'failed');
CREATE TYPE file_role AS ENUM (
    'original',
    'thumbnail_small',
    'thumbnail_medium',
    'video_preview'
);

-- tiles
CREATE TABLE tiles (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL,
    type tile_type NOT NULL,
    title TEXT,
    description TEXT,
    visibility tile_visibility NOT NULL DEFAULT 'private',
    status tile_status NOT NULL DEFAULT 'draft',
    publish_requested_at TIMESTAMP,
    published_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP NULL
);

-- tile_metadata
CREATE TABLE tile_metadata (
    tile_id UUID PRIMARY KEY
        REFERENCES tiles(id) ON DELETE CASCADE,
    duration_seconds INT,
    dominant_color TEXT
);

-- tile_files
CREATE TABLE tile_files (
    tile_id UUID NOT NULL
        REFERENCES tiles(id) ON DELETE CASCADE,
    file_id UUID NOT NULL,
    role file_role NOT NULL,
    width INT,
    height INT,
    PRIMARY KEY (tile_id, role)
);

-- indexes
CREATE INDEX idx_tiles_owner_id ON tiles(owner_id);
CREATE INDEX idx_tiles_status ON tiles(status);
CREATE INDEX idx_tiles_deleted ON tiles(deleted);
