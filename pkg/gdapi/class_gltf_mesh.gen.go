// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFMesh struct {
  obj gdc.ObjectPtr
}

func (me *GLTFMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFMesh) BaseClass() string {
  return "GLTFMesh"
}
