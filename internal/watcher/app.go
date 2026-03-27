// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package watcher

import (
	"github.com/robinlg/iam/internal/watcher/config"
	"github.com/robinlg/iam/internal/watcher/options"
	"github.com/robinlg/iam/pkg/app"
	"github.com/robinlg/iam/pkg/log"
)

const commandDesc = `IAM Watcher is a pluggable watcher service used to do some periodic work like cron job. 
But the difference with cron job is iam-watcher also support sleep some duration after previous job done.

Find more iam-pump information at:
    https://github.com/robinlg/iam/blob/master/docs/guide/en-US/cmd/iam-watcher.md`

// NewApp creates an App object with default parameters.
func NewApp(basename string) *app.App {
	opts := options.NewOptions()
	application := app.NewApp("IAM watcher server",
		basename,
		app.WithOptions(opts),
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		log.Init(opts.Log)
		defer log.Flush()

		cfg, err := config.CreateConfigFromOptions(opts)
		if err != nil {
			return err
		}

		return Run(cfg)
	}
}
