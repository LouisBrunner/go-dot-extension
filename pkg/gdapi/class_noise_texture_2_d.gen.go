// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NoiseTexture2D struct {
  obj gdc.ObjectPtr
}

func (me *NoiseTexture2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NoiseTexture2D) BaseClass() string {
  return "NoiseTexture2D"
}
