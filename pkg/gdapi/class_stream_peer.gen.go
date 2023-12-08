// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeer struct {
  obj gdc.ObjectPtr
}

func createStreamPeer(obj gdc.ObjectPtr) *StreamPeer {
  return &StreamPeer{
    obj: obj,
  }
}

func (me *StreamPeer) BaseClass() string {
  return "StreamPeer"
}
