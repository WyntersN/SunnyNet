//go:build go1.8
// +build go1.8

/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2025-08-13 16:08:30
 * @LastEditTime: 2025-08-13 18:52:40
 * @FilePath: \SunnyNet\src\websocket\client_clone.go
 */
// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import "github.com/WyntersN/SunnyNet/src/crypto/tls"

func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()
}
