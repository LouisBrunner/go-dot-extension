// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *VisualShaderNodeTextureParameter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTextureParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTextureParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeTextureParameter) SetTextureType(type_ VisualShaderNodeTextureParameterTextureType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2227296876) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTextureParameter) GetTextureType() VisualShaderNodeTextureParameterTextureType {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 367922070) // FIXME: should cache?
  var ret VisualShaderNodeTextureParameterTextureType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeTextureParameter) SetColorDefault(color VisualShaderNodeTextureParameterColorDefault, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color_default")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4217624432) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&color), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTextureParameter) GetColorDefault() VisualShaderNodeTextureParameterColorDefault {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color_default")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3837060134) // FIXME: should cache?
  var ret VisualShaderNodeTextureParameterColorDefault
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeTextureParameter) SetTextureFilter(filter VisualShaderNodeTextureParameterTextureFilter, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2147684752) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&filter), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTextureParameter) GetTextureFilter() VisualShaderNodeTextureParameterTextureFilter {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_filter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4184490817) // FIXME: should cache?
  var ret VisualShaderNodeTextureParameterTextureFilter
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeTextureParameter) SetTextureRepeat(repeat VisualShaderNodeTextureParameterTextureRepeat, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2036143070) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&repeat), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTextureParameter) GetTextureRepeat() VisualShaderNodeTextureParameterTextureRepeat {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_repeat")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1690132794) // FIXME: should cache?
  var ret VisualShaderNodeTextureParameterTextureRepeat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeTextureParameter) SetTextureSource(source VisualShaderNodeTextureParameterTextureSource, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1212687372) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&source), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTextureParameter) GetTextureSource() VisualShaderNodeTextureParameterTextureSource {
  classNameV := StringNameFromStr("VisualShaderNodeTextureParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2039092262) // FIXME: should cache?
  var ret VisualShaderNodeTextureParameterTextureSource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
