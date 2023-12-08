// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationRegion3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationRegion3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationRegion3D) BaseClass() string {
  return "NavigationRegion3D"
}
