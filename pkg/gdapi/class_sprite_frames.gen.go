// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpriteFrames struct {
  obj gdc.ObjectPtr
}

func createSpriteFrames(obj gdc.ObjectPtr) *SpriteFrames {
  return &SpriteFrames{
    obj: obj,
  }
}

func (me *SpriteFrames) BaseClass() string {
  return "SpriteFrames"
}
