// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"math/rand"
	"time"

	"github.com/robinlg/iam/internal/watcher"
	_ "go.uber.org/automaxprocs"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	watcher.NewApp("iam-watcher").Run()
}
