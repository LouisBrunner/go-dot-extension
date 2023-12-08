// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraAttributes struct {
  obj gdc.ObjectPtr
}

func createCameraAttributes(obj gdc.ObjectPtr) *CameraAttributes {
  return &CameraAttributes{
    obj: obj,
  }
}

func (me *CameraAttributes) BaseClass() string {
  return "CameraAttributes"
}
