// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTextureParameter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTextureParameter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTextureParameter) BaseClass() string {
  return "VisualShaderNodeTextureParameter"
}



// Enums

type VisualShaderNodeTextureParameterTextureType int
const (
  VisualShaderNodeTextureParameterTextureTypeTypeData VisualShaderNodeTextureParameterTextureType = 0
  VisualShaderNodeTextureParameterTextureTypeTypeColor VisualShaderNodeTextureParameterTextureType = 1
  VisualShaderNodeTextureParameterTextureTypeTypeNormalMap VisualShaderNodeTextureParameterTextureType = 2
  VisualShaderNodeTextureParameterTextureTypeTypeAnisotropy VisualShaderNodeTextureParameterTextureType = 3
  VisualShaderNodeTextureParameterTextureTypeTypeMax VisualShaderNodeTextureParameterTextureType = 4
)

type VisualShaderNodeTextureParameterColorDefault int
const (
  VisualShaderNodeTextureParameterColorDefaultColorDefaultWhite VisualShaderNodeTextureParameterColorDefault = 0
  VisualShaderNodeTextureParameterColorDefaultColorDefaultBlack VisualShaderNodeTextureParameterColorDefault = 1
  VisualShaderNodeTextureParameterColorDefaultColorDefaultTransparent VisualShaderNodeTextureParameterColorDefault = 2
  VisualShaderNodeTextureParameterColorDefaultColorDefaultMax VisualShaderNodeTextureParameterColorDefault = 3
)

type VisualShaderNodeTextureParameterTextureFilter int
const (
  VisualShaderNodeTextureParameterTextureFilterFilterDefault VisualShaderNodeTextureParameterTextureFilter = 0
  VisualShaderNodeTextureParameterTextureFilterFilterNearest VisualShaderNodeTextureParameterTextureFilter = 1
  VisualShaderNodeTextureParameterTextureFilterFilterLinear VisualShaderNodeTextureParameterTextureFilter = 2
  VisualShaderNodeTextureParameterTextureFilterFilterNearestMipmap VisualShaderNodeTextureParameterTextureFilter = 3
  VisualShaderNodeTextureParameterTextureFilterFilterLinearMipmap VisualShaderNodeTextureParameterTextureFilter = 4
  VisualShaderNodeTextureParameterTextureFilterFilterNearestMipmapAnisotropic VisualShaderNodeTextureParameterTextureFilter = 5
  VisualShaderNodeTextureParameterTextureFilterFilterLinearMipmapAnisotropic VisualShaderNodeTextureParameterTextureFilter = 6
  VisualShaderNodeTextureParameterTextureFilterFilterMax VisualShaderNodeTextureParameterTextureFilter = 7
)

type VisualShaderNodeTextureParameterTextureRepeat int
const (
  VisualShaderNodeTextureParameterTextureRepeatRepeatDefault VisualShaderNodeTextureParameterTextureRepeat = 0
  VisualShaderNodeTextureParameterTextureRepeatRepeatEnabled VisualShaderNodeTextureParameterTextureRepeat = 1
  VisualShaderNodeTextureParameterTextureRepeatRepeatDisabled VisualShaderNodeTextureParameterTextureRepeat = 2
  VisualShaderNodeTextureParameterTextureRepeatRepeatMax VisualShaderNodeTextureParameterTextureRepeat = 3
)

type VisualShaderNodeTextureParameterTextureSource int
const (
  VisualShaderNodeTextureParameterTextureSourceSourceNone VisualShaderNodeTextureParameterTextureSource = 0
  VisualShaderNodeTextureParameterTextureSourceSourceScreen VisualShaderNodeTextureParameterTextureSource = 1
  VisualShaderNodeTextureParameterTextureSourceSourceDepth VisualShaderNodeTextureParameterTextureSource = 2
  VisualShaderNodeTextureParameterTextureSourceSourceNormalRoughness VisualShaderNodeTextureParameterTextureSource = 3
  VisualShaderNodeTextureParameterTextureSourceSourceMax VisualShaderNodeTextureParameterTextureSource = 4
)

func (me *VisualShaderNodeTextureParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTextureParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeTextureParameter) SetTextureType(type_ VisualShaderNodeTextureParameterTextureType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTextureParameter) GetTextureType()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTextureParameter) SetColorDefault(color VisualShaderNodeTextureParameterColorDefault, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTextureParameter) GetColorDefault()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTextureParameter) SetTextureFilter(filter VisualShaderNodeTextureParameterTextureFilter, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTextureParameter) GetTextureFilter()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTextureParameter) SetTextureRepeat(repeat VisualShaderNodeTextureParameterTextureRepeat, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTextureParameter) GetTextureRepeat()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTextureParameter) SetTextureSource(source VisualShaderNodeTextureParameterTextureSource, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTextureParameter) GetTextureSource()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
