// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package memstore

import (
	"context"
	"time"

	"github.com/kian99/candid/meeting"
)

var PutAtTime = func(ctx context.Context, s meeting.Store, id, address string, now time.Time) error {
	return s.(*meetingStore).put(ctx, id, address, now)
}
