// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioListener2D struct {
  obj gdc.ObjectPtr
}

func (me *AudioListener2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioListener2D) BaseClass() string {
  return "AudioListener2D"
}
