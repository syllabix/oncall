package shift

const update = `
	INSERT INTO shifts (user_id, schedule_id, sequence_id, status, started_at)
	VALUES (:user_id, :schedule_id, :sequence_id, :status, :started_at)
	ON CONFLICT(user_id, schedule_id)
	DO UPDATE SET
		user_id = EXCLUDED.user_id,
		schedule_id = EXCLUDED.schedule_id,
		sequence_id = EXCLUDED.sequence_id,
		status = EXCLUDED.status,
		started_at = EXCLUDED.started_at,
		updated_at = now()`
