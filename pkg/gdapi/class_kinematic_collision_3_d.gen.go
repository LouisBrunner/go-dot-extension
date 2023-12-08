// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type KinematicCollision3D struct {
  obj gdc.ObjectPtr
}

func (me *KinematicCollision3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *KinematicCollision3D) BaseClass() string {
  return "KinematicCollision3D"
}
