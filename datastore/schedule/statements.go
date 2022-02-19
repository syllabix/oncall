package schedule

const upsertStmt = `
	WITH e AS(
		INSERT INTO users (slack_id, slack_handle, email, first_name, last_name, avatar_url, display_name)
		VALUES (:slack_id, :slack_handle, :email, :first_name, :last_name, :avatar_url, :display_name)
		ON CONFLICT(slack_id) DO NOTHING
		RETURNING id, created_at, updated_at
	)
	SELECT * FROM e
	UNION
		SELECT id, created_at, updated_at 
		FROM users 
		WHERE slack_id=:slack_id;`
