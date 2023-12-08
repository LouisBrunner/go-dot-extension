// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StyleBoxTexture struct {
  obj gdc.ObjectPtr
}

func (me *StyleBoxTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StyleBoxTexture) BaseClass() string {
  return "StyleBoxTexture"
}

type StyleBoxTextureAxisStretchMode int
const (
  StyleBoxTextureAxisStretchModeStretch StyleBoxTextureAxisStretchMode = 0
  StyleBoxTextureAxisStretchModeTile StyleBoxTextureAxisStretchMode = 1
  StyleBoxTextureAxisStretchModeTileFit StyleBoxTextureAxisStretchMode = 2
)
