CREATE TABLE
    IF NOT EXISTS messages (
        id UUID PRIMARY KEY, 
        room_id UUID NOT NULL REFERENCES rooms (id),
        author_id UUID NOT NULL REFERENCES users (id),
        content VARCHAR(2000) NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );
