// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StreamPeerTLS struct {
  obj gdc.ObjectPtr
}

func createStreamPeerTLS(obj gdc.ObjectPtr) *StreamPeerTLS {
  return &StreamPeerTLS{
    obj: obj,
  }
}

func (me *StreamPeerTLS) BaseClass() string {
  return "StreamPeerTLS"
}
