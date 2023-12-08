// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerBuffer struct {
  obj gdc.ObjectPtr
}

func createStreamPeerBuffer(obj gdc.ObjectPtr) *StreamPeerBuffer {
  return &StreamPeerBuffer{
    obj: obj,
  }
}

func (me *StreamPeerBuffer) BaseClass() string {
  return "StreamPeerBuffer"
}
