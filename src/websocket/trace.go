//go:build go1.8
// +build go1.8

/*
 * @Descripttion:
 * @version:
 * @Author: Wynters
 * @Date: 2025-08-13 16:08:30
 * @LastEditTime: 2025-08-13 18:52:54
 * @FilePath: \SunnyNet\src\websocket\trace.go
 */

package websocket

import (
	"github.com/WyntersN/SunnyNet/src/crypto/tls"
	"github.com/WyntersN/SunnyNet/src/http/httptrace"
)

func doHandshakeWithTrace(trace *httptrace.ClientTrace, tlsConn *tls.Conn, cfg *tls.Config) error {
	if trace.TLSHandshakeStart != nil {
		trace.TLSHandshakeStart()
	}
	err := doHandshake(tlsConn, cfg)
	if trace.TLSHandshakeDone != nil {
		trace.TLSHandshakeDone(tlsConn.ConnectionState(), err)
	}
	return err
}
