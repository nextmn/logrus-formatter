// Copyright Louis Royer and the NextMN contributors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
// SPDX-License-Identifier: MIT

package httplog_test

import (
	"log"
	"net/http"

	"github.com/nextmn/logrus-formatter/httplog"
)

func ExampleNewRequestLoggerMiddleware() {
	mux := http.NewServeMux()
	logger := httplog.NewRequestLoggerMiddleware(mux)
	log.Fatal(http.ListenAndServe(":8080", logger))
}
