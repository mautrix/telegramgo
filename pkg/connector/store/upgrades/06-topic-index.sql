-- v6 (compatible with v2+): Add table for topics

CREATE TABLE telegram_topic (
    channel_id BIGINT NOT NULL,
    topic_id   BIGINT NOT NULL,

    PRIMARY KEY (channel_id, topic_id)
);
