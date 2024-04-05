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

type ptrsForVisualShaderNodeGroupBaseList struct {
  fnSetInputs gdc.MethodBindPtr
  fnGetInputs gdc.MethodBindPtr
  fnSetOutputs gdc.MethodBindPtr
  fnGetOutputs gdc.MethodBindPtr
  fnIsValidPortName gdc.MethodBindPtr
  fnAddInputPort gdc.MethodBindPtr
  fnRemoveInputPort gdc.MethodBindPtr
  fnGetInputPortCount gdc.MethodBindPtr
  fnHasInputPort gdc.MethodBindPtr
  fnClearInputPorts gdc.MethodBindPtr
  fnAddOutputPort gdc.MethodBindPtr
  fnRemoveOutputPort gdc.MethodBindPtr
  fnGetOutputPortCount gdc.MethodBindPtr
  fnHasOutputPort gdc.MethodBindPtr
  fnClearOutputPorts gdc.MethodBindPtr
  fnSetInputPortName gdc.MethodBindPtr
  fnSetInputPortType gdc.MethodBindPtr
  fnSetOutputPortName gdc.MethodBindPtr
  fnSetOutputPortType gdc.MethodBindPtr
  fnGetFreeInputPortId gdc.MethodBindPtr
  fnGetFreeOutputPortId gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeGroupBase ptrsForVisualShaderNodeGroupBaseList

func initVisualShaderNodeGroupBasePtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeGroupBase")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_inputs")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnSetInputs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_inputs")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnGetInputs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_outputs")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnSetOutputs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_outputs")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnGetOutputs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("is_valid_port_name")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnIsValidPortName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("add_input_port")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnAddInputPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2285447957))
  }
  {
    methodName := StringNameFromStr("remove_input_port")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnRemoveInputPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_input_port_count")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnGetInputPortCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("has_input_port")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnHasInputPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("clear_input_ports")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnClearInputPorts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("add_output_port")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnAddOutputPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2285447957))
  }
  {
    methodName := StringNameFromStr("remove_output_port")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnRemoveOutputPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_output_port_count")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnGetOutputPortCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("has_output_port")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnHasOutputPort = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("clear_output_ports")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnClearOutputPorts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_input_port_name")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnSetInputPortName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
  }
  {
    methodName := StringNameFromStr("set_input_port_type")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnSetInputPortType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("set_output_port_name")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnSetOutputPortName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
  }
  {
    methodName := StringNameFromStr("set_output_port_type")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnSetOutputPortType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("get_free_input_port_id")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnGetFreeInputPortId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_free_output_port_id")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeGroupBase.fnGetFreeOutputPortId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type VisualShaderNodeGroupBase struct {
  VisualShaderNodeResizableBase
}

func (me *VisualShaderNodeGroupBase) BaseClass() string {
  return "VisualShaderNodeGroupBase"
}

func NewVisualShaderNodeGroupBase() *VisualShaderNodeGroupBase {
  str := StringNameFromStr("VisualShaderNodeGroupBase") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeGroupBase{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{inputs.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnSetInputs), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) GetInputs() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnGetInputs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShaderNodeGroupBase) SetOutputs(outputs String, )  {
  cargs := []gdc.ConstTypePtr{outputs.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnSetOutputs), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) GetOutputs() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnGetOutputs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShaderNodeGroupBase) IsValidPortName(name String, ) bool {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnIsValidPortName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeGroupBase) AddInputPort(id int64, type_ int64, name String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&type_) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnAddInputPort), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) RemoveInputPort(id int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnRemoveInputPort), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) GetInputPortCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnGetInputPortCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeGroupBase) HasInputPort(id int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnHasInputPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeGroupBase) ClearInputPorts()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnClearInputPorts), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) AddOutputPort(id int64, type_ int64, name String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&type_) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnAddOutputPort), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) RemoveOutputPort(id int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnRemoveOutputPort), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) GetOutputPortCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnGetOutputPortCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeGroupBase) HasOutputPort(id int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnHasOutputPort), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeGroupBase) ClearOutputPorts()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnClearOutputPorts), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) SetInputPortName(id int64, name String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnSetInputPortName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) SetInputPortType(id int64, type_ int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnSetInputPortType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) SetOutputPortName(id int64, name String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnSetOutputPortName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) SetOutputPortType(id int64, type_ int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnSetOutputPortType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeGroupBase) GetFreeInputPortId() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnGetFreeInputPortId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeGroupBase) GetFreeOutputPortId() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeGroupBase.fnGetFreeOutputPortId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
