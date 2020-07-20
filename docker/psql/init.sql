CREATE TABLE reminders
(
    id         UUID PRIMARY KEY,
    guild_id   VARCHAR(64) NOT NULL,
    channel_id VARCHAR(64) NOT NULL,
    time       TIMESTAMP WITH TIME ZONE,
    name       VARCHAR(30),
    done       BOOLEAN DEFAULT FALSE
);