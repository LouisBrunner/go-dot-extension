// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShader struct {
  obj gdc.ObjectPtr
}

func (me *VisualShader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShader) BaseClass() string {
  return "VisualShader"
}

const (
  VisualShaderNODE_ID_INVALID = -1
  VisualShaderNODE_ID_OUTPUT = 0
)

type VisualShaderType int
const (
  VisualShaderTypeVertex VisualShaderType = 0
  VisualShaderTypeFragment VisualShaderType = 1
  VisualShaderTypeLight VisualShaderType = 2
  VisualShaderTypeStart VisualShaderType = 3
  VisualShaderTypeProcess VisualShaderType = 4
  VisualShaderTypeCollide VisualShaderType = 5
  VisualShaderTypeStartCustom VisualShaderType = 6
  VisualShaderTypeProcessCustom VisualShaderType = 7
  VisualShaderTypeSky VisualShaderType = 8
  VisualShaderTypeFog VisualShaderType = 9
  VisualShaderTypeMax VisualShaderType = 10
)

type VisualShaderVaryingMode int
const (
  VisualShaderVaryingModeVertexToFragLight VisualShaderVaryingMode = 0
  VisualShaderVaryingModeFragToLight VisualShaderVaryingMode = 1
  VisualShaderVaryingModeMax VisualShaderVaryingMode = 2
)

type VisualShaderVaryingType int
const (
  VisualShaderVaryingTypeFloat VisualShaderVaryingType = 0
  VisualShaderVaryingTypeInt VisualShaderVaryingType = 1
  VisualShaderVaryingTypeUint VisualShaderVaryingType = 2
  VisualShaderVaryingTypeVector2D VisualShaderVaryingType = 3
  VisualShaderVaryingTypeVector3D VisualShaderVaryingType = 4
  VisualShaderVaryingTypeVector4D VisualShaderVaryingType = 5
  VisualShaderVaryingTypeBoolean VisualShaderVaryingType = 6
  VisualShaderVaryingTypeTransform VisualShaderVaryingType = 7
  VisualShaderVaryingTypeMax VisualShaderVaryingType = 8
)
