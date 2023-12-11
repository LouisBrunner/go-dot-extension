// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GLTFLight struct {
  obj gdc.ObjectPtr
}

func (me *GLTFLight) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFLight) BaseClass() string {
  return "GLTFLight"
}



// Enums

func (me *GLTFLight) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFLight) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFLight) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  GLTFLightFromNode(light_node Light3D, ) GLTFLight {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3907677874) // FIXME: should cache?
  var ret GLTFLight
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(light_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFLight) ToNode() Light3D {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2040811672) // FIXME: should cache?
  var ret Light3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  GLTFLightFromDictionary(dictionary Dictionary, ) GLTFLight {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("from_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4057087208) // FIXME: should cache?
  var ret GLTFLight
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(dictionary.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFLight) ToDictionary() Dictionary {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_dictionary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3102165223) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFLight) GetColor() Color {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3200896285) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFLight) SetColor(color Color, )  {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFLight) GetIntensity() float32 {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFLight) SetIntensity(intensity float32, )  {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_intensity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&intensity), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFLight) GetLightType() String {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_light_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFLight) SetLightType(light_type String, )  {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_light_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(light_type.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFLight) GetRange() float32 {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFLight) SetRange(range_ float32, )  {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&range_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFLight) GetInnerConeAngle() float32 {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inner_cone_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFLight) SetInnerConeAngle(inner_cone_angle float32, )  {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inner_cone_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&inner_cone_angle), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GLTFLight) GetOuterConeAngle() float32 {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outer_cone_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GLTFLight) SetOuterConeAngle(outer_cone_angle float32, )  {
  classNameV := StringNameFromStr("GLTFLight")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outer_cone_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&outer_cone_angle), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
