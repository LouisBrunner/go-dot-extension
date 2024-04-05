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

type ptrsForSpringArm3DList struct {
  fnGetHitLength gdc.MethodBindPtr
  fnSetLength gdc.MethodBindPtr
  fnGetLength gdc.MethodBindPtr
  fnSetShape gdc.MethodBindPtr
  fnGetShape gdc.MethodBindPtr
  fnAddExcludedObject gdc.MethodBindPtr
  fnRemoveExcludedObject gdc.MethodBindPtr
  fnClearExcludedObjects gdc.MethodBindPtr
  fnSetCollisionMask gdc.MethodBindPtr
  fnGetCollisionMask gdc.MethodBindPtr
  fnSetMargin gdc.MethodBindPtr
  fnGetMargin gdc.MethodBindPtr
}

var ptrsForSpringArm3D ptrsForSpringArm3DList

func initSpringArm3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SpringArm3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_hit_length")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnGetHitLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_length")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnSetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_length")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnGetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_shape")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1549710052))
  }
  {
    methodName := StringNameFromStr("get_shape")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3214262478))
  }
  {
    methodName := StringNameFromStr("add_excluded_object")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnAddExcludedObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
  }
  {
    methodName := StringNameFromStr("remove_excluded_object")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnRemoveExcludedObject = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
  }
  {
    methodName := StringNameFromStr("clear_excluded_objects")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnClearExcludedObjects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_collision_mask")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_collision_mask")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_margin")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnSetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_margin")
    defer methodName.Destroy()
    ptrsForSpringArm3D.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
}

type SpringArm3D struct {
  Node3D
}

func (me *SpringArm3D) BaseClass() string {
  return "SpringArm3D"
}

func NewSpringArm3D() *SpringArm3D {
  str := StringNameFromStr("SpringArm3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SpringArm3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SpringArm3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SpringArm3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SpringArm3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SpringArm3D) GetHitLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnGetHitLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpringArm3D) SetLength(length float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnSetLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpringArm3D) GetLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnGetLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpringArm3D) SetShape(shape Shape3D, )  {
  cargs := []gdc.ConstTypePtr{shape.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpringArm3D) GetShape() Shape3D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShape3D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SpringArm3D) AddExcludedObject(RID RID, )  {
  cargs := []gdc.ConstTypePtr{RID.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnAddExcludedObject), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpringArm3D) RemoveExcludedObject(RID RID, ) bool {
  cargs := []gdc.ConstTypePtr{RID.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnRemoveExcludedObject), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpringArm3D) ClearExcludedObjects()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnClearExcludedObjects), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpringArm3D) SetCollisionMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpringArm3D) GetCollisionMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SpringArm3D) SetMargin(margin float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnSetMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SpringArm3D) GetMargin() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSpringArm3D.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
