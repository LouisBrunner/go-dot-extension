// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type KinematicCollision2D struct {
  obj gdc.ObjectPtr
}

func createKinematicCollision2D(obj gdc.ObjectPtr) *KinematicCollision2D {
  return &KinematicCollision2D{
    obj: obj,
  }
}

func (me *KinematicCollision2D) BaseClass() string {
  return "KinematicCollision2D"
}
