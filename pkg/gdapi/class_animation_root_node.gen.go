// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationRootNode struct {
  obj gdc.ObjectPtr
}

func (me *AnimationRootNode) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationRootNode) BaseClass() string {
  return "AnimationRootNode"
}
