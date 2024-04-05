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

type ptrsForCSGSphere3DList struct {
  fnSetRadius gdc.MethodBindPtr
  fnGetRadius gdc.MethodBindPtr
  fnSetRadialSegments gdc.MethodBindPtr
  fnGetRadialSegments gdc.MethodBindPtr
  fnSetRings gdc.MethodBindPtr
  fnGetRings gdc.MethodBindPtr
  fnSetSmoothFaces gdc.MethodBindPtr
  fnGetSmoothFaces gdc.MethodBindPtr
  fnSetMaterial gdc.MethodBindPtr
  fnGetMaterial gdc.MethodBindPtr
}

var ptrsForCSGSphere3D ptrsForCSGSphere3DList

func initCSGSphere3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CSGSphere3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_radius")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_radius")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_radial_segments")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnSetRadialSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_radial_segments")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnGetRadialSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_rings")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnSetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_rings")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnGetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_smooth_faces")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnSetSmoothFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_smooth_faces")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnGetSmoothFaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_material")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
  }
  {
    methodName := StringNameFromStr("get_material")
    defer methodName.Destroy()
    ptrsForCSGSphere3D.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
  }
}

type CSGSphere3D struct {
  CSGPrimitive3D
}

func (me *CSGSphere3D) BaseClass() string {
  return "CSGSphere3D"
}

func NewCSGSphere3D() *CSGSphere3D {
  str := StringNameFromStr("CSGSphere3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CSGSphere3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CSGSphere3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGSphere3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGSphere3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CSGSphere3D) SetRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGSphere3D) GetRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGSphere3D) SetRadialSegments(radial_segments int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radial_segments) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnSetRadialSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGSphere3D) GetRadialSegments() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnGetRadialSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGSphere3D) SetRings(rings int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rings) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnSetRings), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGSphere3D) GetRings() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnGetRings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGSphere3D) SetSmoothFaces(smooth_faces bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&smooth_faces) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnSetSmoothFaces), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGSphere3D) GetSmoothFaces() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnGetSmoothFaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CSGSphere3D) SetMaterial(material Material, )  {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGSphere3D) GetMaterial() Material {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGSphere3D.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
