-- indexes
DROP INDEX idx_content_deleted;
DROP INDEX idx_content_status;
DROP INDEX idx_content_owner_id;

-- tables
DROP TABLE content_files;
DROP TABLE content_metadata;
DROP TABLE content;

-- ENUMs
DROP TYPE file_role;
DROP TYPE content_status;
DROP TYPE content_visibility;
DROP TYPE content_type;
