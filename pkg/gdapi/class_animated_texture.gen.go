// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatedTexture struct {
  obj gdc.ObjectPtr
}

func (me *AnimatedTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatedTexture) BaseClass() string {
  return "AnimatedTexture"
}
