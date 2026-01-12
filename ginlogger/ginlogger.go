// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package ginlogger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggingMiddleware is a gin gonic middleware to use logrus for logging
func LoggingMiddleware(ctx *gin.Context) {
	ctx.Next()
	method := ctx.Request.Method
	uri := ctx.Request.RequestURI
	status := ctx.Writer.Status()
	ipAddr := ctx.ClientIP()
	logrus.WithFields(logrus.Fields{
		"method":     method,
		"uri":        uri,
		"status":     status,
		"ip-address": ipAddr,
	}).Info("HTTP Request")
	ctx.Next()
}
