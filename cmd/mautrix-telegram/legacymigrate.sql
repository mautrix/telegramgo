INSERT INTO "user" (bridge_id, mxid)
SELECT '', mxid FROM user_old;

ALTER TABLE telethon_sessions_old ADD COLUMN json_data jsonb;
UPDATE telethon_sessions_old SET json_data=
    -- only: postgres
    jsonb_build_object
    -- only: sqlite (line commented)
--  json_object
    (
        -- only: postgres
        'auth_key', encode(auth_key, 'base64'),
        -- only: sqlite (line commented)
--      'auth_key', replace(base64(auth_key), x'0A', ''),
        'dc_id', dc_id,
        'server_address', server_address,
        'port', port
    );

INSERT INTO user_login (bridge_id, user_mxid, id, remote_name, remote_profile, space_room, metadata)
SELECT
    '', -- bridge_id
    mxid, -- user_mxid
    CAST(tgid AS TEXT), -- id
    COALESCE(tg_username, tg_phone, ''), -- remote_name
    '{}', -- remote_profile
    '', -- space_room
    -- only: postgres
    jsonb_build_object
    -- only: sqlite (line commented)
--  json_object
    (
        'phone', COALESCE(tg_phone, ''),
        'session', json((SELECT json_data FROM telethon_sessions_old WHERE session_id=tgid))
    ) -- metadata
FROM user_old
WHERE tgid IS NOT NULL;

INSERT INTO ghost (
    bridge_id, id, name, avatar_id, avatar_hash, avatar_mxc,
    name_set, avatar_set, contact_info_set, is_bot, identifiers, metadata
)
SELECT
    '', -- bridge_id
    CAST(id AS TEXT), -- id
    displayname, -- name
    photo_id, -- avatar_id
    '', -- avatar_hash
    avatar_url, -- avatar_mxc
    name_set,
    avatar_set,
    contact_info_set,
    is_bot,
    '[]', -- identifiers
    -- only: postgres
    jsonb_build_object
    -- only: sqlite (line commented)
--  json_object
    (
        'is_premium', is_premium,
        'is_channel', is_channel,
        'phone', phone,
        'name_source', displayname_source,
        'name_quality', displayname_quality,
        'name_not_contact', CASE WHEN displayname_contact THEN json('false') ELSE json('true') END,
    ) -- metadata
FROM puppet_old;

DELETE FROM user_portal_old WHERE portal IN (SELECT tgid FROM portal_old WHERE peer_type<>'channel');

UPDATE portal_old
SET tg_receiver=COALESCE((SELECT "user" FROM user_portal_old WHERE portal=portal_old.tgid LIMIT 1), tg_receiver)
WHERE peer_type='chat' AND tgid=tg_receiver;

UPDATE portal_old
SET tg_receiver=COALESCE((SELECT tgid FROM user_old WHERE tgid IS NOT NULL LIMIT 1), tg_receiver)
WHERE peer_type='chat' AND tgid=tg_receiver;

DELETE FROM portal_old WHERE peer_type='chat' AND tgid=tg_receiver;

INSERT INTO portal (
    bridge_id, id, receiver, mxid, other_user_id, name, topic, avatar_id, avatar_hash, avatar_mxc,
    name_set, avatar_set, topic_set, name_is_custom, in_space, room_type, metadata
)
SELECT
    '', -- bridge_id
    CAST(tgid AS TEXT), -- id
    CAST(tg_receiver AS TEXT), -- receiver
    mxid, -- mxid
    CASE WHEN peer_type='user' THEN CAST(tgid AS TEXT) END, -- other_user_id
    title, -- name
    about, -- topic
    photo_id, -- avatar_id
    '', -- avatar_hash
    avatar_url, -- avatar_mxc
    name_set, -- name_set
    avatar_set, -- avatar_set
    false, -- topic_set
    peer_type<>'user', -- name_is_custom
    false, -- in_space
    CASE WHEN peer_type='user' THEN 'dm' ELSE '' END, -- room_type
    -- only: postgres
    jsonb_build_object
    -- only: sqlite (line commented)
--  json_object
    (
        -- TODO
    ) -- metadata
FROM portal_old;

INSERT INTO user_portal (bridge_id, user_mxid, login_id, portal_id, portal_receiver, in_space, preferred)
SELECT
    '', -- bridge_id
    user_old.mxid, -- user_mxid
    CAST(user_portal_old.user AS TEXT), -- login_id
    CAST(user_portal_old.portal AS TEXT), -- portal_id
    CAST(user_portal_old.portal_receiver AS TEXT), -- portal_receiver
    false, -- in_space
    false -- preferred
FROM user_portal_old
INNER JOIN user_old ON user_portal_old."user" = user_old.tgid;

INSERT INTO user_portal (bridge_id, user_mxid, login_id, portal_id, portal_receiver, in_space, preferred)
SELECT
    '', -- bridge_id
    user_old.mxid, -- user_mxid
    CAST(portal_old.tg_receiver AS TEXT), -- login_id
    CAST(portal_old.tgid AS TEXT), -- portal_id
    CAST(portal_old.tg_receiver AS TEXT), -- portal_receiver
    false, -- in_space
    false -- preferred
FROM portal_old
INNER JOIN user_old ON portal_old.tg_receiver = user_old.tgid
WHERE portal_old.tg_receiver<>portal_old.tgid
ON CONFLICT (bridge_id, user_mxid, login_id, portal_id, portal_receiver) DO NOTHING;

INSERT INTO message (
    bridge_id, id, part_id, mxid, room_id, room_receiver, sender_id, sender_mxid, timestamp, edit_count, metadata
)
SELECT
    '', -- bridge_id
    CASE WHEN tg_space=portal_old.tgid THEN (CAST(tg_space AS TEXT) || '.') ELSE '' END || CAST(message_old.tgid AS TEXT), -- id
    '', -- part_id
    message_old.mxid, -- mxid
    CAST(portal_old.tgid AS TEXT), -- room_id
    CAST(portal_old.tg_receiver AS TEXT), -- room_receiver
    CAST(sender AS TEXT), -- sender_id
    sender_mxid,
    0, -- timestamp
    edit_index, -- edit_count
    '{}' -- metadata
FROM message_old
INNER JOIN portal_old ON mx_room=portal_old.mxid
WHERE tg_space=portal_old.tgid OR tg_space=portal_old.tg_receiver;

INSERT INTO reaction (
    bridge_id, message_id, message_part_id, sender_id, emoji_id, room_id, room_receiver, mxid, timestamp, emoji, metadata
)
SELECT
    '', -- bridge_id
    message.id, -- message_id
    message.part_id, -- message_part_id
    CAST(tg_sender AS TEXT), -- sender_id
    reaction, -- emoji_id
    message.room_id, -- room_id
    message.room_receiver, -- room_receiver
    reaction_old.mxid, -- mxid
    0, -- timestamp
    reaction, -- emoji
    '{}' -- metadata
FROM reaction_old
INNER JOIN message ON reaction_old.msg_mxid=message.mxid;

-- Python -> Go mx_ table migration
ALTER TABLE mx_room_state DROP COLUMN is_encrypted;
ALTER TABLE mx_room_state RENAME COLUMN has_full_member_list TO members_fetched;

-- only: postgres until "end only"
ALTER TABLE mx_room_state ALTER COLUMN power_levels TYPE jsonb USING power_levels::jsonb;
ALTER TABLE mx_room_state ALTER COLUMN encryption TYPE jsonb USING encryption::jsonb;
ALTER TABLE mx_room_state ALTER COLUMN members_fetched SET DEFAULT false;
ALTER TABLE mx_room_state ALTER COLUMN members_fetched SET NOT NULL;
-- end only postgres

ALTER TABLE mx_user_profile ADD COLUMN name_skeleton bytea;
CREATE INDEX mx_user_profile_membership_idx ON mx_user_profile (room_id, membership);
CREATE INDEX mx_user_profile_name_skeleton_idx ON mx_user_profile (room_id, name_skeleton);

UPDATE mx_user_profile SET displayname='' WHERE displayname IS NULL;
UPDATE mx_user_profile SET avatar_url='' WHERE avatar_url IS NULL;

CREATE TABLE mx_registrations (
    user_id TEXT PRIMARY KEY
);

UPDATE mx_version SET version=7;