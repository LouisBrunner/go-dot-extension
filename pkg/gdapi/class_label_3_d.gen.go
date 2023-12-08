// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Label3D struct {
  obj gdc.ObjectPtr
}

func (me *Label3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Label3D) BaseClass() string {
  return "Label3D"
}
