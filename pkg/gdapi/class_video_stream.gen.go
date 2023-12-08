// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VideoStream struct {
  obj gdc.ObjectPtr
}

func createVideoStream(obj gdc.ObjectPtr) *VideoStream {
  return &VideoStream{
    obj: obj,
  }
}

func (me *VideoStream) BaseClass() string {
  return "VideoStream"
}
