// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Mesh struct {
  obj gdc.ObjectPtr
}

func (me *Mesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Mesh) BaseClass() string {
  return "Mesh"
}

type MeshPrimitiveType int
const (
  MeshPrimitivePoints MeshPrimitiveType = 0
  MeshPrimitiveLines MeshPrimitiveType = 1
  MeshPrimitiveLineStrip MeshPrimitiveType = 2
  MeshPrimitiveTriangles MeshPrimitiveType = 3
  MeshPrimitiveTriangleStrip MeshPrimitiveType = 4
)

type MeshArrayType int
const (
  MeshArrayVertex MeshArrayType = 0
  MeshArrayNormal MeshArrayType = 1
  MeshArrayTangent MeshArrayType = 2
  MeshArrayColor MeshArrayType = 3
  MeshArrayTexUv MeshArrayType = 4
  MeshArrayTexUv2 MeshArrayType = 5
  MeshArrayCustom0 MeshArrayType = 6
  MeshArrayCustom1 MeshArrayType = 7
  MeshArrayCustom2 MeshArrayType = 8
  MeshArrayCustom3 MeshArrayType = 9
  MeshArrayBones MeshArrayType = 10
  MeshArrayWeights MeshArrayType = 11
  MeshArrayIndex MeshArrayType = 12
  MeshArrayMax MeshArrayType = 13
)

type MeshArrayCustomFormat int
const (
  MeshArrayCustomRgba8Unorm MeshArrayCustomFormat = 0
  MeshArrayCustomRgba8Snorm MeshArrayCustomFormat = 1
  MeshArrayCustomRgHalf MeshArrayCustomFormat = 2
  MeshArrayCustomRgbaHalf MeshArrayCustomFormat = 3
  MeshArrayCustomRFloat MeshArrayCustomFormat = 4
  MeshArrayCustomRgFloat MeshArrayCustomFormat = 5
  MeshArrayCustomRgbFloat MeshArrayCustomFormat = 6
  MeshArrayCustomRgbaFloat MeshArrayCustomFormat = 7
  MeshArrayCustomMax MeshArrayCustomFormat = 8
)

type MeshArrayFormat int
const (
  MeshArrayFormatVertex MeshArrayFormat = 1
  MeshArrayFormatNormal MeshArrayFormat = 2
  MeshArrayFormatTangent MeshArrayFormat = 4
  MeshArrayFormatColor MeshArrayFormat = 8
  MeshArrayFormatTexUv MeshArrayFormat = 16
  MeshArrayFormatTexUv2 MeshArrayFormat = 32
  MeshArrayFormatCustom0 MeshArrayFormat = 64
  MeshArrayFormatCustom1 MeshArrayFormat = 128
  MeshArrayFormatCustom2 MeshArrayFormat = 256
  MeshArrayFormatCustom3 MeshArrayFormat = 512
  MeshArrayFormatBones MeshArrayFormat = 1024
  MeshArrayFormatWeights MeshArrayFormat = 2048
  MeshArrayFormatIndex MeshArrayFormat = 4096
  MeshArrayFormatBlendShapeMask MeshArrayFormat = 7
  MeshArrayFormatCustomBase MeshArrayFormat = 13
  MeshArrayFormatCustomBits MeshArrayFormat = 3
  MeshArrayFormatCustom0Shift MeshArrayFormat = 13
  MeshArrayFormatCustom1Shift MeshArrayFormat = 16
  MeshArrayFormatCustom2Shift MeshArrayFormat = 19
  MeshArrayFormatCustom3Shift MeshArrayFormat = 22
  MeshArrayFormatCustomMask MeshArrayFormat = 7
  MeshArrayCompressFlagsBase MeshArrayFormat = 25
  MeshArrayFlagUse2DVertices MeshArrayFormat = 33554432
  MeshArrayFlagUseDynamicUpdate MeshArrayFormat = 67108864
  MeshArrayFlagUse8BoneWeights MeshArrayFormat = 134217728
  MeshArrayFlagUsesEmptyVertexArray MeshArrayFormat = 268435456
)

type MeshBlendShapeMode int
const (
  MeshBlendShapeModeNormalized MeshBlendShapeMode = 0
  MeshBlendShapeModeRelative MeshBlendShapeMode = 1
)
