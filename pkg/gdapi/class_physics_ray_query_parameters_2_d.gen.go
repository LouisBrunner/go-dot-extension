// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsRayQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsRayQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsRayQueryParameters2D) BaseClass() string {
  return "PhysicsRayQueryParameters2D"
}
