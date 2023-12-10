// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDVertexAttribute struct {
  obj gdc.ObjectPtr
}

func (me *RDVertexAttribute) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDVertexAttribute) BaseClass() string {
  return "RDVertexAttribute"
}



// Enums

func (me *RDVertexAttribute) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDVertexAttribute) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDVertexAttribute) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDVertexAttribute) SetLocation(p_member int, )  {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_location")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDVertexAttribute) GetLocation() int {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_location")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDVertexAttribute) SetOffset(p_member int, )  {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDVertexAttribute) GetOffset() int {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDVertexAttribute) SetFormat(p_member RenderingDeviceDataFormat, )  {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 565531219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDVertexAttribute) GetFormat() RenderingDeviceDataFormat {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2235804183) // FIXME: should cache?
  var ret RenderingDeviceDataFormat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDVertexAttribute) SetStride(p_member int, )  {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stride")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDVertexAttribute) GetStride() int {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stride")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDVertexAttribute) SetFrequency(p_member RenderingDeviceVertexFrequency, )  {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 522141836) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&p_member), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDVertexAttribute) GetFrequency() RenderingDeviceVertexFrequency {
  classNameV := StringNameFromStr("RDVertexAttribute")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_frequency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4154106413) // FIXME: should cache?
  var ret RenderingDeviceVertexFrequency
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *RDVertexAttribute) GetPropLocation() int {
  panic("TODO: implement")
}

func (me *RDVertexAttribute) SetPropLocation(value int) {
  panic("TODO: implement")
}

func (me *RDVertexAttribute) GetPropOffset() int {
  panic("TODO: implement")
}

func (me *RDVertexAttribute) SetPropOffset(value int) {
  panic("TODO: implement")
}

func (me *RDVertexAttribute) GetPropFormat() int {
  panic("TODO: implement")
}

func (me *RDVertexAttribute) SetPropFormat(value int) {
  panic("TODO: implement")
}

func (me *RDVertexAttribute) GetPropStride() int {
  panic("TODO: implement")
}

func (me *RDVertexAttribute) SetPropStride(value int) {
  panic("TODO: implement")
}

func (me *RDVertexAttribute) GetPropFrequency() int {
  panic("TODO: implement")
}

func (me *RDVertexAttribute) SetPropFrequency(value int) {
  panic("TODO: implement")
}