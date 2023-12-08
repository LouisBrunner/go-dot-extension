// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Label3D struct {
  obj gdc.ObjectPtr
}

func (me *Label3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Label3D) BaseClass() string {
  return "Label3D"
}

type Label3DDrawFlags int
const (
  Label3DFlagShaded Label3DDrawFlags = 0
  Label3DFlagDoubleSided Label3DDrawFlags = 1
  Label3DFlagDisableDepthTest Label3DDrawFlags = 2
  Label3DFlagFixedSize Label3DDrawFlags = 3
  Label3DFlagMax Label3DDrawFlags = 4
)

type Label3DAlphaCutMode int
const (
  Label3DAlphaCutDisabled Label3DAlphaCutMode = 0
  Label3DAlphaCutDiscard Label3DAlphaCutMode = 1
  Label3DAlphaCutOpaquePrepass Label3DAlphaCutMode = 2
  Label3DAlphaCutHash Label3DAlphaCutMode = 3
)
