// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CurveTexture struct {
  obj gdc.ObjectPtr
}

func (me *CurveTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CurveTexture) BaseClass() string {
  return "CurveTexture"
}

type CurveTextureTextureMode int
const (
  CurveTextureTextureModeRgb CurveTextureTextureMode = 0
  CurveTextureTextureModeRed CurveTextureTextureMode = 1
)
