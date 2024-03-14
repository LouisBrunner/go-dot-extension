// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WorldBoundaryShape3D struct {
  Shape3D
}

func (me *WorldBoundaryShape3D) BaseClass() string {
  return "WorldBoundaryShape3D"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(plane.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *WorldBoundaryShape3D) GetPlane() Plane {
  classNameV := StringNameFromStr("WorldBoundaryShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_plane")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2753500971) // FIXME: should cache?
  var ret Plane
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
