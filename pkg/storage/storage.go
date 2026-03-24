// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

import (
	"errors"
	"time"
)

// ErrKeyNotFound is a standard error for when a key is not found in the storage engine.
var ErrKeyNotFound = errors.New("key not found")

// AnalyticsHandler defines the interface for analytics.
type AnalyticsHandler interface {
	Connect() bool
	AppendToSetPipelined(string, [][]byte)
	GetAndDeleteSet(string) []interface{}
	SetExp(string, time.Duration) error // Set key expiration
	GetExp(string) (int64, error)       // Returns expiry of a key
}
