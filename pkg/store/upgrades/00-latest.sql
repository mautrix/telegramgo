-- v0 -> v1: Latest revision

CREATE TABLE telegram_session (
    user_id      INTEGER PRIMARY KEY,
    session_data BYTEA NOT NULL
);

CREATE TABLE telegram_user_state (
    user_id INTEGER PRIMARY KEY,
    pts     INTEGER NOT NULL,
    qts     INTEGER NOT NULL,
    date    INTEGER NOT NULL,
    seq     INTEGER NOT NULL
);

CREATE TABLE telegram_channel_state (
    user_id    INTEGER,
    channel_id INTEGER,
    pts        INTEGER NOT NULL,

    PRIMARY KEY (user_id, channel_id)
);

CREATE INDEX idx_telegram_channel_state_user_id ON telegram_channel_state (user_id);

CREATE TABLE telegram_channel_access_hashes (
    user_id     INTEGER,
    channel_id  INTEGER,
    access_hash INTEGER NOT NULL,

    PRIMARY KEY (user_id, channel_id)
);
