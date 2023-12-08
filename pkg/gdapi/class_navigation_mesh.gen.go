// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationMesh struct {
  obj gdc.ObjectPtr
}

func (me *NavigationMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationMesh) BaseClass() string {
  return "NavigationMesh"
}
