// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFLight struct {
  obj gdc.ObjectPtr
}

func (me *GLTFLight) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFLight) BaseClass() string {
  return "GLTFLight"
}
