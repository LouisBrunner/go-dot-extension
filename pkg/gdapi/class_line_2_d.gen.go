// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Line2D struct {
  obj gdc.ObjectPtr
}

func (me *Line2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Line2D) BaseClass() string {
  return "Line2D"
}

type Line2DLineJointMode int
const (
  Line2DLineJointSharp Line2DLineJointMode = 0
  Line2DLineJointBevel Line2DLineJointMode = 1
  Line2DLineJointRound Line2DLineJointMode = 2
)

type Line2DLineCapMode int
const (
  Line2DLineCapNone Line2DLineCapMode = 0
  Line2DLineCapBox Line2DLineCapMode = 1
  Line2DLineCapRound Line2DLineCapMode = 2
)

type Line2DLineTextureMode int
const (
  Line2DLineTextureNone Line2DLineTextureMode = 0
  Line2DLineTextureTile Line2DLineTextureMode = 1
  Line2DLineTextureStretch Line2DLineTextureMode = 2
)
