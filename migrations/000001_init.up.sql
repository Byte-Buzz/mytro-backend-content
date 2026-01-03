-- ENUMs
CREATE TYPE content_type AS ENUM ('image', 'video');
CREATE TYPE content_visibility AS ENUM ('public', 'private', 'unlisted');
CREATE TYPE content_status AS ENUM ('draft', 'processing', 'ready', 'failed');
CREATE TYPE file_role AS ENUM (
    'original',
    'thumbnail_small',
    'thumbnail_medium',
    'video_preview'
);

-- content
CREATE TABLE content (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL,
    type content_type NOT NULL,
    title TEXT,
    description TEXT,
    visibility content_visibility NOT NULL DEFAULT 'private',
    status content_status NOT NULL DEFAULT 'draft',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP NULL
);

-- content_metadata
CREATE TABLE content_metadata (
    content_id UUID PRIMARY KEY
        REFERENCES content(id) ON DELETE CASCADE,
    duration_seconds INT,
    dominant_color TEXT
);

-- content_files
CREATE TABLE content_files (
    content_id UUID NOT NULL
        REFERENCES content(id) ON DELETE CASCADE,
    file_id UUID NOT NULL,
    role file_role NOT NULL,
    width INT,
    height INT,
    PRIMARY KEY (content_id, role)
);

-- indexes
CREATE INDEX idx_content_owner_id ON content(owner_id);
CREATE INDEX idx_content_status ON content(status);
CREATE INDEX idx_content_deleted ON content(deleted);
