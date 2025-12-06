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

	"go.mau.fi/util/dbutil"
	"go.mau.fi/util/exsync"
)

type topicKey struct {
	ChannelID int64
	TopicID   int
}

type TopicQuery struct {
	db             *dbutil.Database
	existingTopics *exsync.Set[topicKey]
}

const (
	getAllTopicsQuery = `SELECT topic_id FROM telegram_topic WHERE channel_id=$1`
	addTopicQuery     = `INSERT INTO telegram_topic (channel_id, topic_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	deleteTopicQuery  = `DELETE FROM telegram_topic WHERE channel_id=$1 AND topic_id=$2`
)

func (s *TopicQuery) Add(ctx context.Context, channelID int64, topicID int) (err error) {
	if channelID == 0 {
		return nil
	}
	if s.existingTopics.Add(topicKey{ChannelID: channelID, TopicID: topicID}) {
		_, err = s.db.Exec(ctx, addTopicQuery, channelID, topicID)
	}
	return
}

func (s *TopicQuery) Delete(ctx context.Context, channelID int64, topicID int) (err error) {
	s.existingTopics.Remove(topicKey{ChannelID: channelID, TopicID: topicID})
	_, err = s.db.Exec(ctx, deleteTopicQuery, channelID, topicID)
	return
}

var intScanner = dbutil.ConvertRowFn[int](dbutil.ScanSingleColumn[int])

func (s *TopicQuery) GetAll(ctx context.Context, channelID int64) (topics []int, err error) {
	return intScanner.NewRowIter(s.db.Query(ctx, getAllTopicsQuery, channelID)).AsList()
}
