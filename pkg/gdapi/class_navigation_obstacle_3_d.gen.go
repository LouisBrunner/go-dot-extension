// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type NavigationObstacle3D struct {
  obj gdc.ObjectPtr
}

func (me *NavigationObstacle3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *NavigationObstacle3D) BaseClass() string {
  return "NavigationObstacle3D"
}
