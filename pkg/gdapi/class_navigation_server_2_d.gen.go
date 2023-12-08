// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationServer2D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationServer2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationServer2D) BaseClass() string {
  return "NavigationServer2D"
}
