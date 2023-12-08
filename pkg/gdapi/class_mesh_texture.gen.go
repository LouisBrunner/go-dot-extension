// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshTexture struct {
  obj gdc.ObjectPtr
}

func (me *MeshTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MeshTexture) BaseClass() string {
  return "MeshTexture"
}
