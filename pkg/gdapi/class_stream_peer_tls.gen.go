// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerTLS struct {
  obj gdc.ObjectPtr
}

func (me *StreamPeerTLS) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StreamPeerTLS) BaseClass() string {
  return "StreamPeerTLS"
}

type StreamPeerTLSStatus int
const (
  StreamPeerTLSStatusDisconnected StreamPeerTLSStatus = 0
  StreamPeerTLSStatusHandshaking StreamPeerTLSStatus = 1
  StreamPeerTLSStatusConnected StreamPeerTLSStatus = 2
  StreamPeerTLSStatusError StreamPeerTLSStatus = 3
  StreamPeerTLSStatusErrorHostnameMismatch StreamPeerTLSStatus = 4
)
