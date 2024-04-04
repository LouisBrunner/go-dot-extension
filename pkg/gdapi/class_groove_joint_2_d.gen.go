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

type GrooveJoint2D struct {
  Joint2D
}

func (me *GrooveJoint2D) BaseClass() string {
  return "GrooveJoint2D"
}

func NewGrooveJoint2D() *GrooveJoint2D {
  str := StringNameFromStr("GrooveJoint2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GrooveJoint2D{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *GrooveJoint2D) SetLength(length float64, )  {
  classNameV := StringNameFromStr("GrooveJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GrooveJoint2D) GetLength() float64 {
  classNameV := StringNameFromStr("GrooveJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GrooveJoint2D) SetInitialOffset(offset float64, )  {
  classNameV := StringNameFromStr("GrooveJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_initial_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GrooveJoint2D) GetInitialOffset() float64 {
  classNameV := StringNameFromStr("GrooveJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_initial_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
