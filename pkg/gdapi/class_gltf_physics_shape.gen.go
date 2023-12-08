// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFPhysicsShape struct {
  obj gdc.ObjectPtr
}

func (me *GLTFPhysicsShape) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFPhysicsShape) BaseClass() string {
  return "GLTFPhysicsShape"
}
