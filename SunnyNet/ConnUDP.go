package SunnyNet

import (
	"github.com/WyntersN/SunnyNet/src/CrossCompiled"
	"github.com/WyntersN/SunnyNet/src/Interface"
	NFapi "github.com/WyntersN/SunnyNet/src/ProcessDrv/nfapi"
)

type ConnUDP Interface.ConnUDPCall

type udpConn struct {
	sunnyContext  int
	theology      int64 // 唯一ID
	messageId     int   // 消息ID
	_type         int   // 请求类型 例如 public.SunnyNetUDPType...
	pid           int
	localAddress  string
	remoteAddress string
	data          []byte
	_Display      bool
}

func (U udpConn) SetDisplay(Display bool) {
	U._Display = Display
}

func (U udpConn) GetSocket5User() Interface.AuthUser {
	return Interface.AuthUser{}
}

func (U udpConn) GetProcessName() string {
	if U.pid == 0 {
		return "代理连接"
	}
	return CrossCompiled.GetPidName(int32(U.pid))
}

// SetBody 修改消息
func (U udpConn) SetBody(i []byte) bool {
	U.data = i
	return true
}

func (U udpConn) BodyLen() int {
	return len(U.data)
}

func (U udpConn) Context() int {
	return U.sunnyContext
}

func (U udpConn) Type() int {
	return U._type
}

func (U udpConn) MessageId() int {
	return U.messageId
}

func (U udpConn) Theology() int {
	return int(U.theology)
}

func (U udpConn) PID() int {
	return U.pid
}

func (U udpConn) LocalAddress() string {
	return U.localAddress
}

func (U udpConn) RemoteAddress() string {
	return U.remoteAddress
}

func (U udpConn) Body() []byte {
	return U.data
}

// SendToServer 主动向服务器发送消息
func (U udpConn) SendToServer(data []byte) bool {
	return NFapi.UdpSendToServer(U.theology, data)
}

// SendToClient 主动向客户端发送消息
func (U udpConn) SendToClient(data []byte) bool {
	return NFapi.UdpSendToClient(U.theology, data)
}
