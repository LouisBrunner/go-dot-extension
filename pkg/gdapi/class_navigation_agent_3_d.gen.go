// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationAgent3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationAgent3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationAgent3D) BaseClass() string {
  return "NavigationAgent3D"
}
