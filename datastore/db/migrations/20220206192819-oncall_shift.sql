
-- +migrate Up
CREATE TYPE shift_interval AS ENUM ('daily', 'weekly', 'bi-weekly', 'monthly');

CREATE TABLE IF NOT EXISTS schedules (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    slack_channel_id text NOT NULL CONSTRAINT schedules_slack_id_uq UNIQUE,
    team_slack_id text NOT NULL,     
    name text NOT NULL,
    interval shift_interval NOT NULL,
    is_enabled boolean NOT NULL DEFAULT true,
    start_time time with time zone NOT NULL,
    end_time time with time zone NOT NULL,
    active_shift uuid NULL,
    override_shift uuid NULL,
    shifts text[] NOT NULL DEFAULT array[]::text[],
    weekdays_only boolean NOT NULL DEFAULT true,
    created_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    updated_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    deleted_at timestamp with time zone NULL
);

CREATE INDEX schedules_created_at_idx ON schedules (created_at DESC);
CREATE INDEX schedules_deleted_at_idx ON schedules (deleted_at DESC);
CREATE INDEX schedules_active_shift_idx ON schedules (active_shift);
CREATE INDEX schedules_override_shift_idx ON schedules (override_shift);
CREATE INDEX schedules_team_slack_id_idx ON schedules (team_slack_id);
CREATE INDEX schedules_slack_channel_id_idx ON schedules (slack_channel_id);

-- +migrate Down
DROP INDEX IF EXISTS schedule_created_at_idx;
DROP INDEX IF EXISTS schedule_deleted_at_idx;
DROP INDEX IF EXISTS schedule_active_shift_idx;
DROP INDEX IF EXISTS schedule_override_shift_idx;
DROP INDEX IF EXISTS schedule_team_id_idx;
DROP TABLE IF EXISTS schedules;

DROP TYPE IF EXISTS shift_interval;