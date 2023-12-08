// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer2DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer2DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer2DExtension) BaseClass() string {
  return "PhysicsServer2DExtension"
}
