// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDUniform struct {
  obj gdc.ObjectPtr
}

func (me *RDUniform) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDUniform) BaseClass() string {
  return "RDUniform"
}



// Enums

func (me *RDUniform) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDUniform) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDUniform) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDUniform) SetUniformType(p_member RenderingDeviceUniformType, )  {
  classNameV := StringNameFromStr("RDUniform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uniform_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1664894931) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDUniform) GetUniformType() RenderingDeviceUniformType {
  classNameV := StringNameFromStr("RDUniform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uniform_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 475470040) // FIXME: should cache?
  var ret RenderingDeviceUniformType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDUniform) SetBinding(p_member int, )  {
  classNameV := StringNameFromStr("RDUniform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_binding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDUniform) GetBinding() int {
  classNameV := StringNameFromStr("RDUniform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_binding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDUniform) AddId(id RID, )  {
  classNameV := StringNameFromStr("RDUniform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(id.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDUniform) ClearIds()  {
  classNameV := StringNameFromStr("RDUniform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_ids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDUniform) GetIds() RID {
  classNameV := StringNameFromStr("RDUniform")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ids")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
