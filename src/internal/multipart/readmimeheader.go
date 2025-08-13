/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2025-08-13 16:08:30
 * @LastEditTime: 2025-08-13 18:49:30
 * @FilePath: \SunnyNet\src\internal\multipart\readmimeheader.go
 */
// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package multipart

import (
	_ "unsafe" // for go:linkname

	"github.com/WyntersN/SunnyNet/src/internal/textproto"
)

// readMIMEHeader is defined in package net/textproto.
//
//go:linkname readMIMEHeader net/textproto.readMIMEHeader
func readMIMEHeader(r *textproto.Reader, maxMemory, maxHeaders int64) (textproto.MIMEHeader, error)
