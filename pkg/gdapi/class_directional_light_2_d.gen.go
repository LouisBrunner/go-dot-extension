// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DirectionalLight2D struct {
  obj gdc.ObjectPtr
}

func (me *DirectionalLight2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DirectionalLight2D) BaseClass() string {
  return "DirectionalLight2D"
}
