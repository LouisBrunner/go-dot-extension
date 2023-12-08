// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerTCP struct {
  obj gdc.ObjectPtr
}

func createStreamPeerTCP(obj gdc.ObjectPtr) *StreamPeerTCP {
  return &StreamPeerTCP{
    obj: obj,
  }
}

func (me *StreamPeerTCP) BaseClass() string {
  return "StreamPeerTCP"
}
