// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerTCP struct {
  obj gdc.ObjectPtr
}

func (me *StreamPeerTCP) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StreamPeerTCP) BaseClass() string {
  return "StreamPeerTCP"
}

type StreamPeerTCPStatus int
const (
  StreamPeerTCPStatusNone StreamPeerTCPStatus = 0
  StreamPeerTCPStatusConnecting StreamPeerTCPStatus = 1
  StreamPeerTCPStatusConnected StreamPeerTCPStatus = 2
  StreamPeerTCPStatusError StreamPeerTCPStatus = 3
)
