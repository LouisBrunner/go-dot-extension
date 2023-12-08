// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationAgent2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationAgent2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationAgent2D) BaseClass() string {
  return "NavigationAgent2D"
}
