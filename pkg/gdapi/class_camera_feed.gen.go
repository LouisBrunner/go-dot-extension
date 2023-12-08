// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraFeed struct {
  obj gdc.ObjectPtr
}

func createCameraFeed(obj gdc.ObjectPtr) *CameraFeed {
  return &CameraFeed{
    obj: obj,
  }
}

func (me *CameraFeed) BaseClass() string {
  return "CameraFeed"
}
