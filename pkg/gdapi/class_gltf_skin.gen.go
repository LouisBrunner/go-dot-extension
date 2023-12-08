// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFSkin struct {
  obj gdc.ObjectPtr
}

func (me *GLTFSkin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFSkin) BaseClass() string {
  return "GLTFSkin"
}
