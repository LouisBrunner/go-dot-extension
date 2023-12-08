// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedScene struct {
  obj gdc.ObjectPtr
}

func (me *PackedScene) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PackedScene) BaseClass() string {
  return "PackedScene"
}
