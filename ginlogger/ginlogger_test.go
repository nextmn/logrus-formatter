// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package ginlogger_test

import (
	"github.com/gin-gonic/gin"
	"github.com/nextmn/logrus-formatter/ginlogger"
)

func ExampleLoggingMiddleware() {
	r := gin.New()                     // Create new gin server
	r.Use(gin.Recovery())              // Add a recovery function
	r.Use(ginlogger.LoggingMiddleware) // Use the middleware
	// ...
}
