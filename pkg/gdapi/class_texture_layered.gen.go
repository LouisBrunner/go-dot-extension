// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextureLayered struct {
  obj gdc.ObjectPtr
}

func (me *TextureLayered) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureLayered) BaseClass() string {
  return "TextureLayered"
}

type TextureLayeredLayeredType int
const (
  TextureLayeredLayeredType2DArray TextureLayeredLayeredType = 0
  TextureLayeredLayeredTypeCubemap TextureLayeredLayeredType = 1
  TextureLayeredLayeredTypeCubemapArray TextureLayeredLayeredType = 2
)
