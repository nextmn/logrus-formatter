// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package logger_test

import (
	"github.com/nextmn/logrus-formatter/logger"

	"github.com/sirupsen/logrus"
)

func ExampleInit() {
	appName := "foo"     // name of the application
	logger.Init(appName) // init the logrus formatter
	// ...
	logrus.Info("An event occurred")
}
