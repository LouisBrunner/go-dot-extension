// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeGroupBase struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeGroupBase) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeGroupBase) BaseClass() string {
  return "VisualShaderNodeGroupBase"
}



// Enums

func (me *VisualShaderNodeGroupBase) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeGroupBase) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeGroupBase) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeGroupBase) SetInputs(inputs String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_inputs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(inputs.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) GetInputs() String {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_inputs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeGroupBase) SetOutputs(outputs String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_outputs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(outputs.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) GetOutputs() String {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_outputs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeGroupBase) IsValidPortName(name String, ) bool {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_valid_port_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeGroupBase) AddInputPort(id int, type_ int, name String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_input_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2285447957) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) RemoveInputPort(id int, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_input_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) GetInputPortCount() int {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_port_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeGroupBase) HasInputPort(id int, ) bool {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_input_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeGroupBase) ClearInputPorts()  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_input_ports")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) AddOutputPort(id int, type_ int, name String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_output_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2285447957) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) RemoveOutputPort(id int, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_output_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) GetOutputPortCount() int {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_port_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeGroupBase) HasOutputPort(id int, ) bool {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_output_port")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeGroupBase) ClearOutputPorts()  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_output_ports")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) SetInputPortName(id int, name String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_port_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) SetInputPortType(id int, type_ int, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_port_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) SetOutputPortName(id int, name String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_output_port_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) SetOutputPortType(id int, type_ int, )  {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_output_port_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeGroupBase) GetFreeInputPortId() int {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_free_input_port_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeGroupBase) GetFreeOutputPortId() int {
  classNameV := StringNameFromStr("VisualShaderNodeGroupBase")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_free_output_port_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
