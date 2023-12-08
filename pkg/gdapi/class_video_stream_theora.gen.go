// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VideoStreamTheora struct {
  obj gdc.ObjectPtr
}

func createVideoStreamTheora(obj gdc.ObjectPtr) *VideoStreamTheora {
  return &VideoStreamTheora{
    obj: obj,
  }
}

func (me *VideoStreamTheora) BaseClass() string {
  return "VideoStreamTheora"
}
