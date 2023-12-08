// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFPhysicsBody struct {
  obj gdc.ObjectPtr
}

func createGLTFPhysicsBody(obj gdc.ObjectPtr) *GLTFPhysicsBody {
  return &GLTFPhysicsBody{
    obj: obj,
  }
}

func (me *GLTFPhysicsBody) BaseClass() string {
  return "GLTFPhysicsBody"
}
