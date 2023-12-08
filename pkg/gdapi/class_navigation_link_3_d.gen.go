// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationLink3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationLink3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationLink3D) BaseClass() string {
  return "NavigationLink3D"
}
