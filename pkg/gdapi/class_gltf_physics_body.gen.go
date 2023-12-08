// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFPhysicsBody struct {
  obj gdc.ObjectPtr
}

func (me *GLTFPhysicsBody) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFPhysicsBody) BaseClass() string {
  return "GLTFPhysicsBody"
}
