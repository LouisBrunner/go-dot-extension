// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TorusMesh struct {
  obj gdc.ObjectPtr
}

func (me *TorusMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TorusMesh) BaseClass() string {
  return "TorusMesh"
}
