// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GrooveJoint2D struct {
  obj gdc.ObjectPtr
}

func (me *GrooveJoint2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GrooveJoint2D) BaseClass() string {
  return "GrooveJoint2D"
}



// Enums

func (me *GrooveJoint2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GrooveJoint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GrooveJoint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GrooveJoint2D) SetLength(length float32, )  {
  classNameV := StringNameFromStr("GrooveJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GrooveJoint2D) GetLength() float32 {
  classNameV := StringNameFromStr("GrooveJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GrooveJoint2D) SetInitialOffset(offset float32, )  {
  classNameV := StringNameFromStr("GrooveJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_initial_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GrooveJoint2D) GetInitialOffset() float32 {
  classNameV := StringNameFromStr("GrooveJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_initial_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *GrooveJoint2D) GetPropLength() float32 {
  panic("TODO: implement")
}

func (me *GrooveJoint2D) SetPropLength(value float32) {
  panic("TODO: implement")
}

func (me *GrooveJoint2D) GetPropInitialOffset() float32 {
  panic("TODO: implement")
}

func (me *GrooveJoint2D) SetPropInitialOffset(value float32) {
  panic("TODO: implement")
}