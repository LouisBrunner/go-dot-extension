// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type KinematicCollision2D struct {
  obj gdc.ObjectPtr
}

func (me *KinematicCollision2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *KinematicCollision2D) BaseClass() string {
  return "KinematicCollision2D"
}
