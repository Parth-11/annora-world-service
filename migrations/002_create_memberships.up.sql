CREATE TABLE world_memberships (
    world_id UUID NOT NULL,
    user_id UUID NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('member', 'moderator', 'owner')),
    reputation INT NOT NULL DEFAULT 0,
    joined_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (world_id, user_id),
    CONSTRAINT fk_world_memberships_world
        FOREIGN KEY (world_id)
        REFERENCES worlds(id)
        ON DELETE CASCADE
);