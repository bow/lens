// Copyright (c) 2023 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bow/iris/internal"
)

func (db *SQLite) ListEntries(
	ctx context.Context,
	feedID internal.ID,
) ([]*internal.Entry, error) {

	recs := make([]*entryRecord, 0)
	dbFunc := func(ctx context.Context, tx *sql.Tx) error {
		_, err := getFeed(ctx, tx, feedID)
		if errors.Is(err, sql.ErrNoRows) {
			return internal.FeedNotFoundError{ID: feedID}
		}
		irecs, err := getAllFeedEntries(ctx, tx, feedID, nil)
		recs = irecs
		return err
	}

	fail := failF("SQLite.ListEntries")

	db.mu.Lock()
	defer db.mu.Unlock()

	err := db.withTx(ctx, dbFunc)
	if err != nil {
		return nil, fail(err)
	}

	return entryRecords(recs).entries(), nil
}