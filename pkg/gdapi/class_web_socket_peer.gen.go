// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebSocketPeer struct {
  obj gdc.ObjectPtr
}

func (me *WebSocketPeer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WebSocketPeer) BaseClass() string {
  return "WebSocketPeer"
}

type WebSocketPeerWriteMode int
const (
  WebSocketPeerWriteModeText WebSocketPeerWriteMode = 0
  WebSocketPeerWriteModeBinary WebSocketPeerWriteMode = 1
)

type WebSocketPeerState int
const (
  WebSocketPeerStateConnecting WebSocketPeerState = 0
  WebSocketPeerStateOpen WebSocketPeerState = 1
  WebSocketPeerStateClosing WebSocketPeerState = 2
  WebSocketPeerStateClosed WebSocketPeerState = 3
)
