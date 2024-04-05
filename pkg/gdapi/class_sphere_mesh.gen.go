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

type ptrsForSphereMeshList struct {
  fnSetRadius gdc.MethodBindPtr
  fnGetRadius gdc.MethodBindPtr
  fnSetHeight gdc.MethodBindPtr
  fnGetHeight gdc.MethodBindPtr
  fnSetRadialSegments gdc.MethodBindPtr
  fnGetRadialSegments gdc.MethodBindPtr
  fnSetRings gdc.MethodBindPtr
  fnGetRings gdc.MethodBindPtr
  fnSetIsHemisphere gdc.MethodBindPtr
  fnGetIsHemisphere gdc.MethodBindPtr
}

var ptrsForSphereMesh ptrsForSphereMeshList

func initSphereMeshPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SphereMesh")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_radius")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnSetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_radius")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnGetRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_height")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_height")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_radial_segments")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnSetRadialSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_radial_segments")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnGetRadialSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_rings")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnSetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_rings")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnGetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_is_hemisphere")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnSetIsHemisphere = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_is_hemisphere")
    defer methodName.Destroy()
    ptrsForSphereMesh.fnGetIsHemisphere = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type SphereMesh struct {
  PrimitiveMesh
}

func (me *SphereMesh) BaseClass() string {
  return "SphereMesh"
}

func NewSphereMesh() *SphereMesh {
  str := StringNameFromStr("SphereMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SphereMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SphereMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SphereMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SphereMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SphereMesh) SetRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnSetRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SphereMesh) GetRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnGetRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SphereMesh) SetHeight(height float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SphereMesh) GetHeight() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SphereMesh) SetRadialSegments(radial_segments int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radial_segments) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnSetRadialSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SphereMesh) GetRadialSegments() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnGetRadialSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SphereMesh) SetRings(rings int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rings) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnSetRings), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SphereMesh) GetRings() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnGetRings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SphereMesh) SetIsHemisphere(is_hemisphere bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&is_hemisphere) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnSetIsHemisphere), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SphereMesh) GetIsHemisphere() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSphereMesh.fnGetIsHemisphere), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
