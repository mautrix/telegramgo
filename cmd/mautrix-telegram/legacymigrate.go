package main

import (
	_ "embed"
)

const legacyMigrateRenameTables = `
ALTER TABLE backfill_queue RENAME TO backfill_queue_old;
ALTER TABLE bot_chat RENAME TO bot_chat_old;
ALTER TABLE contact RENAME TO contact_old;
ALTER TABLE disappearing_message RENAME TO disappearing_message_old;
ALTER TABLE message RENAME TO message_old;
ALTER TABLE portal RENAME TO portal_old;
ALTER TABLE puppet RENAME TO puppet_old;
ALTER TABLE reaction RENAME TO reaction_old;
ALTER TABLE telegram_file RENAME TO telegram_file_old;
ALTER TABLE telethon_entities RENAME TO telethon_entities_old;
ALTER TABLE telethon_sent_files RENAME TO telethon_sent_files_old;
ALTER TABLE telethon_sessions RENAME TO telethon_sessions_old;
ALTER TABLE telethon_update_state RENAME TO telethon_update_state_old;
ALTER TABLE "user" RENAME TO user_old;
ALTER TABLE user_portal RENAME TO user_portal_old;
`

//go:embed legacymigrate.sql
var legacyMigrateCopyData string
