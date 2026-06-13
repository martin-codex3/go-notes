-- +goose Up
CREATE IF NOT EXISTS notes (
    notes_id SERIAL PRIMARY KEY,
    notes_title VARCHAR NOT NULL,
    notes_description TEXT NOT NULL,
    is_complete BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE notes;
