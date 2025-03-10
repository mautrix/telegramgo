// mautrix-telegram - A Matrix-Telegram puppeting bridge.
// Copyright (C) 2025 Sumner Evans
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

	"go.mau.fi/util/dbutil"
	"maunium.net/go/mautrix/id"
)

const (
	insertTelegramFileQuery          = "INSERT INTO telegram_file (id, mxc, mime_type, size) VALUES ($1, $2, $3, $4)"
	getTelegramFileSelect            = "SELECT id, mxc, mime_type, size FROM telegram_file "
	getTelegramFileByLocationIDQuery = getTelegramFileSelect + "WHERE id=$1"
	getTelegramFileByMXCQuery        = getTelegramFileSelect + "WHERE mxc=$1"
)

type TelegramFileQuery struct {
	*dbutil.QueryHelper[*TelegramFile]
}

type TelegramFileLocationID string

type TelegramFile struct {
	qh *dbutil.QueryHelper[*TelegramFile]

	LocationID TelegramFileLocationID
	MXC        id.ContentURIString
	MIMEType   string
	Size       int
}

var _ dbutil.DataStruct[*TelegramFile] = (*TelegramFile)(nil)

func newTelegramFile(qh *dbutil.QueryHelper[*TelegramFile]) *TelegramFile {
	return &TelegramFile{qh: qh}
}

func (fq *TelegramFileQuery) GetByLocationID(ctx context.Context, locationID TelegramFileLocationID) (*TelegramFile, error) {
	return fq.QueryOne(ctx, getTelegramFileByLocationIDQuery, locationID)
}

func (fq *TelegramFileQuery) GetByMXC(ctx context.Context, mxc string) (*TelegramFile, error) {
	return fq.QueryOne(ctx, getTelegramFileByMXCQuery, mxc)
}

func (f *TelegramFile) sqlVariables() []any {
	return []any{f.LocationID, f.MXC, f.MIMEType, f.Size}
}

func (f *TelegramFile) Insert(ctx context.Context) error {
	return f.qh.Exec(ctx, insertTelegramFileQuery, f.sqlVariables()...)
}

func (f *TelegramFile) Scan(row dbutil.Scannable) (*TelegramFile, error) {
	return f, row.Scan(&f.LocationID, &f.MXC, &f.MIMEType, &f.Size)
}
