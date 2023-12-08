// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WebSocketMultiplayerPeer struct {
  obj gdc.ObjectPtr
}

func createWebSocketMultiplayerPeer(obj gdc.ObjectPtr) *WebSocketMultiplayerPeer {
  return &WebSocketMultiplayerPeer{
    obj: obj,
  }
}

func (me *WebSocketMultiplayerPeer) BaseClass() string {
  return "WebSocketMultiplayerPeer"
}
