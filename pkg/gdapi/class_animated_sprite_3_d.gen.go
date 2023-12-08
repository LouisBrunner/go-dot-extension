// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatedSprite3D struct {
  obj gdc.ObjectPtr
}

func (me *AnimatedSprite3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatedSprite3D) BaseClass() string {
  return "AnimatedSprite3D"
}
