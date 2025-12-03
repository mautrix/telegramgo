package cached

import (
	"cmp"
	"slices"

	"go.mau.fi/mautrix-telegram/pkg/gotd/telegram/query/hasher"
	"go.mau.fi/mautrix-telegram/pkg/gotd/tg"
)

func (s *ContactsGetContacts) computeHash(v *tg.ContactsContacts) int64 {
	cts := v.Contacts

	slices.SortStableFunc(cts, func(a, b tg.Contact) int {
		return cmp.Compare(a.UserID, b.UserID)
	})
	h := hasher.Hasher{}
	for _, contact := range cts {
		h.Update(uint32(contact.UserID))
	}

	return h.Sum()
}

func (s *MessagesGetQuickReplies) computeHash(v *tg.MessagesQuickReplies) int64 {
	r := v.QuickReplies

	slices.SortStableFunc(r, func(a, b tg.QuickReply) int {
		return cmp.Compare(a.ShortcutID, b.ShortcutID)
	})
	h := hasher.Hasher{}
	for _, contact := range r {
		h.Update(uint32(contact.ShortcutID))
	}

	return h.Sum()
}

func (s *PaymentsGetStarGiftCollections) computeHash(v *tg.PaymentsStarGiftCollections) int64 {
	collections := v.Collections

	slices.SortStableFunc(collections, func(a, b tg.StarGiftCollection) int {
		return cmp.Compare(a.CollectionID, b.CollectionID)
	})

	h := hasher.Hasher{}
	for _, collection := range collections {
		h.Update(uint32(collection.CollectionID))
	}

	return h.Sum()
}

func (s *AccountGetSavedMusicIDs) computeHash(v *tg.AccountSavedMusicIDs) int64 {
	ids := v.IDs

	slices.Sort(ids)

	h := hasher.Hasher{}
	for _, id := range ids {
		h.Update(uint32(id))
	}

	return h.Sum()
}

func (s *PaymentsGetStarGiftActiveAuctions) computeHash(v *tg.PaymentsStarGiftActiveAuctions) int64 {
	auctions := v.Auctions

	slices.SortStableFunc(auctions, func(a, b tg.StarGiftActiveAuctionState) int {
		return cmp.Compare(a.GetGift().GetID(), b.GetGift().GetID())
	})

	h := hasher.Hasher{}
	for _, auction := range auctions {
		h.Update(uint32(auction.GetGift().GetID()))
	}

	return h.Sum()
}
