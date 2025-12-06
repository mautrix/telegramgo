// mautrix-telegram - A Matrix-Telegram puppeting bridge.
// Copyright (C) 2025 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package store

import (
	"context"
	"database/sql"
	"errors"

	"go.mau.fi/util/dbutil"
)

type PhoneNumberQuery struct {
	db *dbutil.Database
}

const (
	getEntityIDForPhoneNumber = "SELECT entity_id FROM telegram_phone_number WHERE phone_number=$1"
	setPhoneNumberQuery       = `
		INSERT INTO telegram_phone_number (phone_number, entity_id)
		VALUES ($1, $2)
		ON CONFLICT (phone_number) DO UPDATE SET entity_id=excluded.entity_id
	`
	clearPhoneNumberQuery = "DELETE FROM telegram_phone_number WHERE entity_id=$1"
)

func (s *PhoneNumberQuery) GetUserID(ctx context.Context, phoneNumber string) (userID int64, err error) {
	err = s.db.QueryRow(ctx, getEntityIDForPhoneNumber, phoneNumber).Scan(&userID)
	if errors.Is(err, sql.ErrNoRows) {
		err = nil
	}
	return
}

func (s *PhoneNumberQuery) Set(ctx context.Context, userID int64, phoneNumber string) (err error) {
	if phoneNumber == "" {
		_, err = s.db.Exec(ctx, clearPhoneNumberQuery, userID)
	} else {
		_, err = s.db.Exec(ctx, setPhoneNumberQuery, phoneNumber, userID)
	}
	return
}
