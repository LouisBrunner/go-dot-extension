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

type CollisionShape3D struct {
  Node3D
}

func (me *CollisionShape3D) BaseClass() string {
  return "CollisionShape3D"
}

func NewCollisionShape3D() *CollisionShape3D {
  str := StringNameFromStr("CollisionShape3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CollisionShape3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CollisionShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CollisionShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CollisionShape3D) ResourceChanged(resource Resource, )  {
  classNameV := StringNameFromStr("CollisionShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("resource_changed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 968641751) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{resource.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionShape3D) SetShape(shape Shape3D, )  {
  classNameV := StringNameFromStr("CollisionShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1549710052) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{shape.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionShape3D) GetShape() Shape3D {
  classNameV := StringNameFromStr("CollisionShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3214262478) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShape3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CollisionShape3D) SetDisabled(enable bool, )  {
  classNameV := StringNameFromStr("CollisionShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionShape3D) IsDisabled() bool {
  classNameV := StringNameFromStr("CollisionShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionShape3D) MakeConvexFromSiblings()  {
  classNameV := StringNameFromStr("CollisionShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("make_convex_from_siblings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
