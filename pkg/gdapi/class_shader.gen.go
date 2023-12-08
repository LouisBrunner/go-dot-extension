// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shader struct {
  obj gdc.ObjectPtr
}

func (me *Shader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Shader) BaseClass() string {
  return "Shader"
}

type ShaderMode int
const (
  ShaderModeSpatial ShaderMode = 0
  ShaderModeCanvasItem ShaderMode = 1
  ShaderModeParticles ShaderMode = 2
  ShaderModeSky ShaderMode = 3
  ShaderModeFog ShaderMode = 4
)
