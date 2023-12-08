// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatedSprite2D struct {
  obj gdc.ObjectPtr
}

func (me *AnimatedSprite2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatedSprite2D) BaseClass() string {
  return "AnimatedSprite2D"
}
