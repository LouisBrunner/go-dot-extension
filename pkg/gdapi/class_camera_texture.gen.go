// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraTexture struct {
  obj gdc.ObjectPtr
}

func createCameraTexture(obj gdc.ObjectPtr) *CameraTexture {
  return &CameraTexture{
    obj: obj,
  }
}

func (me *CameraTexture) BaseClass() string {
  return "CameraTexture"
}