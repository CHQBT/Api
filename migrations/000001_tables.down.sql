-- Drop dependent tables first to avoid constraint issues
DROP TABLE IF EXISTS urls;
DROP TABLE IF EXISTS locations;
DROP TABLE IF EXISTS musics;
DROP TABLE IF EXISTS photos;
DROP TABLE IF EXISTS videos;
DROP TABLE IF EXISTS twit;
DROP TABLE IF EXISTS users;

-- Drop the ENUM type
DROP TYPE IF EXISTS roles;
