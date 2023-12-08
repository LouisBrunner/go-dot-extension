// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DirectionalLight3D struct {
  obj gdc.ObjectPtr
}

func (me *DirectionalLight3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DirectionalLight3D) BaseClass() string {
  return "DirectionalLight3D"
}

type DirectionalLight3DShadowMode int
const (
  DirectionalLight3DShadowOrthogonal DirectionalLight3DShadowMode = 0
  DirectionalLight3DShadowParallel2Splits DirectionalLight3DShadowMode = 1
  DirectionalLight3DShadowParallel4Splits DirectionalLight3DShadowMode = 2
)

type DirectionalLight3DSkyMode int
const (
  DirectionalLight3DSkyModeLightAndSky DirectionalLight3DSkyMode = 0
  DirectionalLight3DSkyModeLightOnly DirectionalLight3DSkyMode = 1
  DirectionalLight3DSkyModeSkyOnly DirectionalLight3DSkyMode = 2
)
