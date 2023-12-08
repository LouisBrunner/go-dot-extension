// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NoiseTexture3D struct {
  obj gdc.ObjectPtr
}

func (me *NoiseTexture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NoiseTexture3D) BaseClass() string {
  return "NoiseTexture3D"
}
