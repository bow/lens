// Copyright (c) 2022 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package store

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/bow/iris/api"
	"github.com/bow/iris/internal/store/opml"
)

const defaultExportTitle = "iris export"

type Subscription []*FeedRecord

func (sub Subscription) Export(title *string) ([]byte, error) {
	et := defaultExportTitle
	if title != nil {
		et = *title
	}
	doc := opml.New(et, time.Now())
	for _, feed := range sub {
		if err := doc.AddOutline(feed); err != nil {
			return nil, err
		}
	}
	return doc.XML()
}

type FeedRecord struct {
	id          ID
	title       string
	description sql.NullString
	feedURL     string
	siteURL     sql.NullString
	subscribed  string
	lastPulled  string
	updated     sql.NullString
	isStarred   bool
	tags        jsonArrayString
	entries     []*EntryRecord
}

func (f *FeedRecord) Outline() (*opml.Outline, error) {
	outl := opml.Outline{
		Text:   f.title,
		Type:   "rss",
		XMLURL: f.feedURL,
	}
	if f.siteURL.Valid {
		outl.HTMLURL = pointer(f.siteURL.String)
	}
	if f.description.Valid {
		outl.Description = pointer(f.description.String)
	}
	if len(f.tags) > 0 {
		outl.Categories = opml.Categories(f.tags)
	}
	if f.isStarred {
		outl.IsStarred = &f.isStarred
	}
	return &outl, nil
}

type FeedBuilder struct {
	id          ID
	title       string
	description *string
	feedURL     string
	siteURL     *string
	subscribed  time.Time
	lastPulled  time.Time
	updated     *time.Time
	isStarred   bool
	tags        []string
	entries     []*EntryRecord
}

func NewFeedBuilder() *FeedBuilder {
	return &FeedBuilder{}
}

func (b *FeedBuilder) Build() *FeedRecord {
	return &FeedRecord{
		id:          b.id,
		title:       b.title,
		description: asNullString(b.description),
		feedURL:     b.feedURL,
		siteURL:     asNullString(b.siteURL),
		subscribed:  *serializeTime(&b.subscribed),
		lastPulled:  *serializeTime(&b.lastPulled),
		updated:     asNullString(serializeTime(b.updated)),
		isStarred:   b.isStarred,
		tags:        jsonArrayString(b.tags),
		entries:     b.entries,
	}
}

func (b *FeedBuilder) ID(value ID) *FeedBuilder {
	b.id = value
	return b
}

func (b *FeedBuilder) Title(value string) *FeedBuilder {
	b.title = value
	return b
}

func (b *FeedBuilder) Description(value *string) *FeedBuilder {
	b.description = value
	return b
}

func (b *FeedBuilder) FeedURL(value string) *FeedBuilder {
	b.feedURL = value
	return b
}

func (b *FeedBuilder) SiteURL(value *string) *FeedBuilder {
	b.siteURL = value
	return b
}

func (b *FeedBuilder) Subscribed(value time.Time) *FeedBuilder {
	b.subscribed = value
	return b
}

func (b *FeedBuilder) LastPulled(value time.Time) *FeedBuilder {
	b.lastPulled = value
	return b
}

func (b *FeedBuilder) Updated(value *time.Time) *FeedBuilder {
	b.updated = value
	return b
}

func (b *FeedBuilder) IsStarred(value bool) *FeedBuilder {
	b.isStarred = value
	return b
}

func (b *FeedBuilder) Tags(value []string) *FeedBuilder {
	b.tags = value
	return b
}

func (b *FeedBuilder) Entries(value []*EntryRecord) *FeedBuilder {
	b.entries = value
	return b
}

type FeedEditOp struct {
	ID          ID
	Title       *string
	Description *string
	Tags        *[]string
	IsStarred   *bool
}

func NewFeedEditOp(proto *api.EditFeedsRequest_Op) *FeedEditOp {
	return &FeedEditOp{
		ID:          proto.Id,
		Title:       proto.Fields.Title,
		Description: proto.Fields.Description,
		Tags:        &proto.Fields.Tags,
		IsStarred:   proto.Fields.IsStarred,
	}
}

type EntryRecord struct {
	ID          ID
	FeedID      ID
	Title       string
	IsRead      bool
	ExtID       string
	Updated     sql.NullString
	Published   sql.NullString
	Description sql.NullString
	Content     sql.NullString
	URL         sql.NullString
}

func (e *EntryRecord) Proto() (*api.Entry, error) {
	proto := api.Entry{
		Id:          e.ID,
		FeedId:      e.FeedID,
		Title:       e.Title,
		IsRead:      e.IsRead,
		ExtId:       e.ExtID,
		Description: fromNullString(e.Description),
		Content:     fromNullString(e.Content),
		Url:         fromNullString(e.URL),
	}

	var err error

	proto.PubTime, err = toProtoTime(fromNullString(e.Published))
	if err != nil {
		return nil, err
	}

	proto.UpdateTime, err = toProtoTime(fromNullString(e.Updated))
	if err != nil {
		return nil, err
	}

	return &proto, nil
}

type EntryEditOp struct {
	ID     ID
	IsRead *bool
}

func NewEntryEditOp(proto *api.EditEntriesRequest_Op) *EntryEditOp {
	return &EntryEditOp{ID: proto.Id, IsRead: proto.Fields.IsRead}
}

type StatsAggregateRecord struct {
	numFeeds             uint32
	numEntries           uint32
	numEntriesUnread     uint32
	lastPullTime         string
	mostRecentUpdateTime sql.NullString
}

// asNullString returns a valid sql.NullString representation of the given string pointer.
func asNullString(v *string) sql.NullString {
	if v == nil {
		return sql.NullString{String: "", Valid: false}
	}
	return toNullString(*v)
}

// toNullString wraps the given string into an sql.NullString value. An empty string input is
// considered a database NULL value.
func toNullString(v string) sql.NullString {
	return sql.NullString{String: v, Valid: v != ""}
}

func resolveFeedUpdateTime(feed *gofeed.Feed) *time.Time {
	// Use feed value if defined.
	var latest = feed.UpdatedParsed
	if latest != nil {
		return latest
	}
	// Otherwise try to infer from entries.
	for _, entry := range feed.Items {
		etv := resolveEntryUpdateTime(entry)
		if latest == nil {
			latest = etv
		}
		if latest != nil && etv != nil {
			if etv.After(*latest) {
				latest = etv
			}
		}
	}
	return latest
}

func resolveEntryUpdateTime(entry *gofeed.Item) *time.Time {
	// Use value if defined.
	if tv := entry.UpdatedParsed; tv != nil {
		return tv
	}
	// Otherwise use published time.
	return entry.PublishedParsed
}

func resolveEntryPublishedTime(entry *gofeed.Item) *time.Time {
	// Use value if defined.
	if tv := entry.PublishedParsed; tv != nil {
		return tv
	}
	// Otherwise use update time.
	return entry.UpdatedParsed
}

func serializeTime(tv *time.Time) *string {
	if tv == nil {
		return nil
	}
	ts := tv.UTC().Format(time.RFC3339)
	return &ts
}

func deserializeTime(v *string) (*time.Time, error) {
	if v == nil {
		return nil, nil
	}
	if *v == "" {
		return nil, nil
	}
	pv, err := time.Parse(time.RFC3339, *v)
	if err != nil {
		return nil, err
	}
	upv := pv.UTC()
	return &upv, nil
}

func toProtoTime(v *string) (*timestamppb.Timestamp, error) {
	tv, err := deserializeTime(v)
	if err != nil {
		return nil, err
	}
	if tv == nil {
		return nil, nil
	}
	return timestamppb.New(*tv), nil
}

// fromNullString unwraps the given sql.NullString value into a string pointer. If the input value
// is NULL (i.e. its `Valid` field is `false`), `nil` is returned.
func fromNullString(v sql.NullString) *string {
	if v.Valid {
		s := v.String
		return &s
	}
	return nil
}

// jsonArrayString is a wrapper type that implements Scan() for database-compatible
// (de)serialization.
type jsonArrayString []string

// Value implements the database valuer interface for serializing into the database.
func (arr *jsonArrayString) Value() (driver.Value, error) {
	if arr == nil {
		return nil, nil
	}
	return json.Marshal([]string(*arr))
}

// Scan implements the database scanner interface for deserialization out of the database.
func (arr *jsonArrayString) Scan(value any) error {
	var bv []byte

	switch v := value.(type) {
	case []byte:
		bv = v
	case string:
		bv = []byte(v)
	default:
		return fmt.Errorf("value of type %T can not be scanned into a string slice", v)
	}

	return json.Unmarshal(bv, arr)
}

func pointer[T any](value T) *T { return &value }
