// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VideoStreamPlayback struct {
  obj gdc.ObjectPtr
}

func createVideoStreamPlayback(obj gdc.ObjectPtr) *VideoStreamPlayback {
  return &VideoStreamPlayback{
    obj: obj,
  }
}

func (me *VideoStreamPlayback) BaseClass() string {
  return "VideoStreamPlayback"
}
