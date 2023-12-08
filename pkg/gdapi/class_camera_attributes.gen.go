// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraAttributes struct {
  obj gdc.ObjectPtr
}

func (me *CameraAttributes) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CameraAttributes) BaseClass() string {
  return "CameraAttributes"
}
