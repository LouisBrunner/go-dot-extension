// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer3DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer3DExtension) BaseClass() string {
  return "PhysicsServer3DExtension"
}
