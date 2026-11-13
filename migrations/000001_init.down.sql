-- indexes
DROP INDEX idx_tiles_deleted;
DROP INDEX idx_tiles_status;
DROP INDEX idx_tiles_owner_id;

-- tables
DROP TABLE tile_files;
DROP TABLE tile_metadata;
DROP TABLE tiles;

-- ENUMs
DROP TYPE file_role;
DROP TYPE tile_status;
DROP TYPE tile_visibility;
DROP TYPE tile_type;
