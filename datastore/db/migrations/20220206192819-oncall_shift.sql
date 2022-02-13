
-- +migrate Up
CREATE TYPE shift_interval AS ENUM ('daily', 'weekly', 'bi-weekly', 'monthly');

CREATE TABLE IF NOT EXISTS oncall_schedule (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    team_id uuid NOT NULL,     
    name text NOT NULL,
    interval shift_interval NOT NULL,
    is_enabled boolean NOT NULL DEFAULT true,
    start_time time with time zone NOT NULL,
    end_time time with time zone NOT NULL,
    active_shift uuid NULL,
    override_shift uuid NULL,
    created_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    updated_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    deleted_at timestamp with time zone NULL
);

CREATE INDEX oncall_schedule_created_at_idx ON oncall_schedule (created_at DESC);
CREATE INDEX oncall_schedule_deleted_at_idx ON oncall_schedule (deleted_at DESC);
CREATE INDEX oncall_schedule_active_shift_idx ON oncall_schedule (active_shift);
CREATE INDEX oncall_schedule_override_shift_idx ON oncall_schedule (override_shift);
CREATE INDEX oncall_schedule_team_id_idx ON oncall_schedule (team_id);

CREATE TABLE IF NOT EXISTS shift (
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL,     
    oncall_schedule_id uuid NOT NULL,
    next_shift uuid NOT NULL,
    prev_shift uuid NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    updated_at timestamp with time zone NOT NULL DEFAULT ('now'::text)::timestamp with time zone,
    deleted_at timestamp with time zone NULL
);

CREATE INDEX shift_user_id_idx ON shift (user_id);
CREATE INDEX shift_oncall_schedule_id ON shift (oncall_schedule_id);

-- +migrate Down
DROP INDEX IF EXISTS oncall_schedule_created_at_idx;
DROP INDEX IF EXISTS oncall_schedule_deleted_at_idx;
DROP INDEX IF EXISTS oncall_schedule_active_shift_idx;
DROP INDEX IF EXISTS oncall_schedule_override_shift_idx;
DROP INDEX IF EXISTS oncall_schedule_team_id_idx;
DROP TABLE IF EXISTS oncall_schedule;

DROP INDEX IF EXISTS shift_user_id_idx;
DROP INDEX IF EXISTS shift_oncall_schedule_id;
DROP TABLE IF EXISTS shift;

DROP TYPE IF EXISTS shift_interval;