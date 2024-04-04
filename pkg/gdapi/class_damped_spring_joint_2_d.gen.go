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

type DampedSpringJoint2D struct {
  Joint2D
}

func (me *DampedSpringJoint2D) BaseClass() string {
  return "DampedSpringJoint2D"
}

func NewDampedSpringJoint2D() *DampedSpringJoint2D {
  str := StringNameFromStr("DampedSpringJoint2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &DampedSpringJoint2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *DampedSpringJoint2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *DampedSpringJoint2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DampedSpringJoint2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *DampedSpringJoint2D) SetLength(length float64, )  {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *DampedSpringJoint2D) GetLength() float64 {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
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

func  (me *DampedSpringJoint2D) SetRestLength(rest_length float64, )  {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rest_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rest_length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *DampedSpringJoint2D) GetRestLength() float64 {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rest_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *DampedSpringJoint2D) SetStiffness(stiffness float64, )  {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stiffness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *DampedSpringJoint2D) GetStiffness() float64 {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *DampedSpringJoint2D) SetDamping(damping float64, )  {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_damping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&damping) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *DampedSpringJoint2D) GetDamping() float64 {
  classNameV := StringNameFromStr("DampedSpringJoint2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_damping")
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
