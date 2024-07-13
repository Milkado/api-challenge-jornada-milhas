-- Add the first user to the database
INSERT into `users` (`email`, `name`, `created_at`, `updated_at`) values ("admin@app.com", "Admin", CURRENT_TIMESTAMP(), CURRENT_TIMESTAMP())
