// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFPhysicsShape struct {
  obj gdc.ObjectPtr
}

func createGLTFPhysicsShape(obj gdc.ObjectPtr) *GLTFPhysicsShape {
  return &GLTFPhysicsShape{
    obj: obj,
  }
}

func (me *GLTFPhysicsShape) BaseClass() string {
  return "GLTFPhysicsShape"
}
