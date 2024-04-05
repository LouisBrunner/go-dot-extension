// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForVisualShaderNodeList struct {
  fnGetDefaultInputPort gdc.MethodBindPtr
  fnSetOutputPortForPreview gdc.MethodBindPtr
  fnGetOutputPortForPreview gdc.MethodBindPtr
  fnSetInputPortDefaultValue gdc.MethodBindPtr
  fnGetInputPortDefaultValue gdc.MethodBindPtr
  fnRemoveInputPortDefaultValue gdc.MethodBindPtr
  fnClearDefaultInputValues gdc.MethodBindPtr
  fnSetDefaultInputValues gdc.MethodBindPtr
  fnGetDefaultInputValues gdc.MethodBindPtr
}

var ptrsForVisualShaderNode ptrsForVisualShaderNodeList

func initVisualShaderNodePtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNode")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_default_input_port")
    defer methodName.Destroy()
    ptrsForVisualShaderNode.fnGetDefaultInputPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1894493699))
  }
  {
    methodName := StringNameFromStr("set_output_port_for_preview")
    defer methodName.Destroy()
    ptrsForVisualShaderNode.fnSetOutputPortForPreview = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_output_port_for_preview")
    defer methodName.Destroy()
    ptrsForVisualShaderNode.fnGetOutputPortForPreview = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_input_port_default_value")
    defer methodName.Destroy()
    ptrsForVisualShaderNode.fnSetInputPortDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 150923387))
  }
  {
    methodName := StringNameFromStr("get_input_port_default_value")
    defer methodName.Destroy()
    ptrsForVisualShaderNode.fnGetInputPortDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
  }
  {
    methodName := StringNameFromStr("remove_input_port_default_value")
    defer methodName.Destroy()
    ptrsForVisualShaderNode.fnRemoveInputPortDefaultValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("clear_default_input_values")
    defer methodName.Destroy()
    ptrsForVisualShaderNode.fnClearDefaultInputValues = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_default_input_values")
    defer methodName.Destroy()
    ptrsForVisualShaderNode.fnSetDefaultInputValues = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_default_input_values")
    defer methodName.Destroy()
    ptrsForVisualShaderNode.fnGetDefaultInputValues = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&type_)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNode.fnGetDefaultInputPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNode) SetOutputPortForPreview(port int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNode.fnSetOutputPortForPreview), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) GetOutputPortForPreview() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNode.fnGetOutputPortForPreview), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNode) SetInputPortDefaultValue(port int64, value Variant, prev_value Variant, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , value.AsCTypePtr(), prev_value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNode.fnSetInputPortDefaultValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) GetInputPortDefaultValue(port int64, ) Variant {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&port)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNode.fnGetInputPortDefaultValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShaderNode) RemoveInputPortDefaultValue(port int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNode.fnRemoveInputPortDefaultValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) ClearDefaultInputValues()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNode.fnClearDefaultInputValues), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) SetDefaultInputValues(values Array, )  {
  cargs := []gdc.ConstTypePtr{values.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNode.fnSetDefaultInputValues), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNode) GetDefaultInputValues() Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNode.fnGetDefaultInputValues), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
