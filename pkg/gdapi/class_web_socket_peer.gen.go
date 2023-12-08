// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebSocketPeer struct {
  obj gdc.ObjectPtr
}

func createWebSocketPeer(obj gdc.ObjectPtr) *WebSocketPeer {
  return &WebSocketPeer{
    obj: obj,
  }
}

func (me *WebSocketPeer) BaseClass() string {
  return "WebSocketPeer"
}
