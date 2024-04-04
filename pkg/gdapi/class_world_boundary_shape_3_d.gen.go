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

type WorldBoundaryShape3D struct {
  Shape3D
}

func (me *WorldBoundaryShape3D) BaseClass() string {
  return "WorldBoundaryShape3D"
}

func NewWorldBoundaryShape3D() *WorldBoundaryShape3D {
  str := StringNameFromStr("WorldBoundaryShape3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &WorldBoundaryShape3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *WorldBoundaryShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WorldBoundaryShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WorldBoundaryShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WorldBoundaryShape3D) SetPlane(plane Plane, )  {
  classNameV := StringNameFromStr("WorldBoundaryShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_plane")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3505987427) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{plane.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WorldBoundaryShape3D) GetPlane() Plane {
  classNameV := StringNameFromStr("WorldBoundaryShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plane")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2753500971) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPlane()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
