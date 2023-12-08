// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CameraTexture struct {
  obj gdc.ObjectPtr
}

func (me *CameraTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CameraTexture) BaseClass() string {
  return "CameraTexture"
}
