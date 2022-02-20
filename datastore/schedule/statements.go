package schedule

const upsertStmt = `
	INSERT INTO users (slack_id, slack_handle, email, first_name, last_name, avatar_url, display_name)
	VALUES (:slack_id, :slack_handle, :email, :first_name, :last_name, :avatar_url, :display_name)
	ON CONFLICT(slack_id)
	DO UPDATE SET
		slack_handle=EXCLUDED.slack_handle,
		email=EXCLUDED.email,
		first_name=EXCLUDED.first_name,
		last_name=EXCLUDED.last_name,
		avatar_url=EXCLUDED.avatar_url,
		display_name=EXCLUDED.display_name,
		updated_at = now()
	RETURNING id, created_at, updated_at`
