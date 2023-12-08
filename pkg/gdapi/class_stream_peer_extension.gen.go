// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerExtension struct {
  obj gdc.ObjectPtr
}

func createStreamPeerExtension(obj gdc.ObjectPtr) *StreamPeerExtension {
  return &StreamPeerExtension{
    obj: obj,
  }
}

func (me *StreamPeerExtension) BaseClass() string {
  return "StreamPeerExtension"
}
