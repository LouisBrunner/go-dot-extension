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

type ptrsForCylinderMeshList struct {
  fnSetTopRadius gdc.MethodBindPtr
  fnGetTopRadius gdc.MethodBindPtr
  fnSetBottomRadius gdc.MethodBindPtr
  fnGetBottomRadius gdc.MethodBindPtr
  fnSetHeight gdc.MethodBindPtr
  fnGetHeight gdc.MethodBindPtr
  fnSetRadialSegments gdc.MethodBindPtr
  fnGetRadialSegments gdc.MethodBindPtr
  fnSetRings gdc.MethodBindPtr
  fnGetRings gdc.MethodBindPtr
  fnSetCapTop gdc.MethodBindPtr
  fnIsCapTop gdc.MethodBindPtr
  fnSetCapBottom gdc.MethodBindPtr
  fnIsCapBottom gdc.MethodBindPtr
}

var ptrsForCylinderMesh ptrsForCylinderMeshList

func initCylinderMeshPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CylinderMesh")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_top_radius")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnSetTopRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_top_radius")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnGetTopRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_bottom_radius")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnSetBottomRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_bottom_radius")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnGetBottomRadius = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_height")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnSetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_height")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_radial_segments")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnSetRadialSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_radial_segments")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnGetRadialSegments = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_rings")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnSetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_rings")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnGetRings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_cap_top")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnSetCapTop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_cap_top")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnIsCapTop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_cap_bottom")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnSetCapBottom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_cap_bottom")
    defer methodName.Destroy()
    ptrsForCylinderMesh.fnIsCapBottom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type CylinderMesh struct {
  PrimitiveMesh
}

func (me *CylinderMesh) BaseClass() string {
  return "CylinderMesh"
}

func NewCylinderMesh() *CylinderMesh {
  str := StringNameFromStr("CylinderMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CylinderMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CylinderMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CylinderMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CylinderMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CylinderMesh) SetTopRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnSetTopRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CylinderMesh) GetTopRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnGetTopRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CylinderMesh) SetBottomRadius(radius float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radius) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnSetBottomRadius), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CylinderMesh) GetBottomRadius() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnGetBottomRadius), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CylinderMesh) SetHeight(height float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&height) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnSetHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CylinderMesh) GetHeight() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CylinderMesh) SetRadialSegments(segments int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&segments) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnSetRadialSegments), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CylinderMesh) GetRadialSegments() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnGetRadialSegments), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CylinderMesh) SetRings(rings int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rings) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnSetRings), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CylinderMesh) GetRings() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnGetRings), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CylinderMesh) SetCapTop(cap_top bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_top) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnSetCapTop), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CylinderMesh) IsCapTop() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnIsCapTop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CylinderMesh) SetCapBottom(cap_bottom bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cap_bottom) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnSetCapBottom), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CylinderMesh) IsCapBottom() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCylinderMesh.fnIsCapBottom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
