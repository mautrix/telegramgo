-- v5 (compatible with v2+): Remove redundant index on telegram_channel_state
DROP INDEX telegram_channel_state_user_id_idx;
