// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VideoStreamPlayer struct {
  obj gdc.ObjectPtr
}

func createVideoStreamPlayer(obj gdc.ObjectPtr) *VideoStreamPlayer {
  return &VideoStreamPlayer{
    obj: obj,
  }
}

func (me *VideoStreamPlayer) BaseClass() string {
  return "VideoStreamPlayer"
}
