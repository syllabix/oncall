
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS teams (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    slack_id text NOT NULL CONSTRAINT team_slack_id_uq UNIQUE,
    name text NOT NULL,
    domain text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    updated_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    deleted_at timestamp with time zone NULL   
);

CREATE INDEX team_created_at_idx ON teams (created_at DESC);
CREATE INDEX team_deleted_at_idx ON teams (deleted_at DESC);

CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    slack_id text NOT NULL CONSTRAINT user_slack_id_uq UNIQUE, 
    email text NOT NULL,
    first_name text NOT NULL,
    last_name text NOT NULL,    
    avatar_url text NOT NULL,
    display_name text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    updated_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    deleted_at timestamp with time zone NULL
);

CREATE INDEX user_created_at_idx ON users (created_at DESC);
CREATE INDEX user_deleted_at_idx ON users (deleted_at DESC);

CREATE TABLE IF NOT EXISTS team_members (
    user_id uuid NOT NULL,
    team_id uuid NOT NULL,
    PRIMARY KEY (user_id, team_id)
);

-- +migrate Down
DROP TABLE IF EXISTS team_members;

DROP INDEX IF EXISTS team_created_at_idx;
DROP INDEX IF EXISTS team_deleted_at_idx;
DROP TABLE IF EXISTS teams;

DROP INDEX IF EXISTS user_created_at_idx;
DROP INDEX IF EXISTS user_deleted_at_idx;
DROP TABLE IF EXISTS users;
