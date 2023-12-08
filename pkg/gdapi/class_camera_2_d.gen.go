// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Camera2D struct {
  obj gdc.ObjectPtr
}

func (me *Camera2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Camera2D) BaseClass() string {
  return "Camera2D"
}

type Camera2DAnchorMode int
const (
  Camera2DAnchorModeFixedTopLeft Camera2DAnchorMode = 0
  Camera2DAnchorModeDragCenter Camera2DAnchorMode = 1
)

type Camera2DCamera2DProcessCallback int
const (
  Camera2DCamera2DProcessPhysics Camera2DCamera2DProcessCallback = 0
  Camera2DCamera2DProcessIdle Camera2DCamera2DProcessCallback = 1
)
