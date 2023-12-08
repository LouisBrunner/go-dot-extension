// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraAttributesPhysical struct {
  obj gdc.ObjectPtr
}

func createCameraAttributesPhysical(obj gdc.ObjectPtr) *CameraAttributesPhysical {
  return &CameraAttributesPhysical{
    obj: obj,
  }
}

func (me *CameraAttributesPhysical) BaseClass() string {
  return "CameraAttributesPhysical"
}
