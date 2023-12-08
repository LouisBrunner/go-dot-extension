// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsPointQueryParameters2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsPointQueryParameters2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsPointQueryParameters2D) BaseClass() string {
  return "PhysicsPointQueryParameters2D"
}
