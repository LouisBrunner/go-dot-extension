// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNode struct {
  Resource
}

func (me *VisualShaderNode) BaseClass() string {
  return "VisualShaderNode"
}

func NewVisualShaderNode() *VisualShaderNode {
  str := StringNameFromStr("VisualShaderNode") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNode{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodePortType int
const (
  VisualShaderNodePortTypePortTypeScalar VisualShaderNodePortType = 0
  VisualShaderNodePortTypePortTypeScalarInt VisualShaderNodePortType = 1
  VisualShaderNodePortTypePortTypeScalarUint VisualShaderNodePortType = 2
  VisualShaderNodePortTypePortTypeVector2D VisualShaderNodePortType = 3
  VisualShaderNodePortTypePortTypeVector3D VisualShaderNodePortType = 4
  VisualShaderNodePortTypePortTypeVector4D VisualShaderNodePortType = 5
  VisualShaderNodePortTypePortTypeBoolean VisualShaderNodePortType = 6
  VisualShaderNodePortTypePortTypeTransform VisualShaderNodePortType = 7
  VisualShaderNodePortTypePortTypeSampler VisualShaderNodePortType = 8
  VisualShaderNodePortTypePortTypeMax VisualShaderNodePortType = 9
)

func (me *VisualShaderNode) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNode) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNode) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNode) GetDefaultInputPort(type_ VisualShaderNodePortType, ) int64 {
  classNameV := StringNameFromStr("VisualShaderNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_input_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1894493699) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&type_)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNode) SetOutputPortForPreview(port int64, )  {
  classNameV := StringNameFromStr("VisualShaderNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_output_port_for_preview")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) GetOutputPortForPreview() int64 {
  classNameV := StringNameFromStr("VisualShaderNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_port_for_preview")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNode) SetInputPortDefaultValue(port int64, value Variant, prev_value Variant, )  {
  classNameV := StringNameFromStr("VisualShaderNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_port_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 150923387) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , value.AsCTypePtr(), prev_value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) GetInputPortDefaultValue(port int64, ) Variant {
  classNameV := StringNameFromStr("VisualShaderNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_port_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShaderNode) RemoveInputPortDefaultValue(port int64, )  {
  classNameV := StringNameFromStr("VisualShaderNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_input_port_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) ClearDefaultInputValues()  {
  classNameV := StringNameFromStr("VisualShaderNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_default_input_values")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) SetDefaultInputValues(values Array, )  {
  classNameV := StringNameFromStr("VisualShaderNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_input_values")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{values.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) GetDefaultInputValues() Array {
  classNameV := StringNameFromStr("VisualShaderNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_input_values")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
