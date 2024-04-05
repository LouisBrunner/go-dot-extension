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

type ptrsForBoxMeshList struct {
  fnSetSize gdc.MethodBindPtr
  fnGetSize gdc.MethodBindPtr
  fnSetSubdivideWidth gdc.MethodBindPtr
  fnGetSubdivideWidth gdc.MethodBindPtr
  fnSetSubdivideHeight gdc.MethodBindPtr
  fnGetSubdivideHeight gdc.MethodBindPtr
  fnSetSubdivideDepth gdc.MethodBindPtr
  fnGetSubdivideDepth gdc.MethodBindPtr
}

var ptrsForBoxMesh ptrsForBoxMeshList

func initBoxMeshPtrs(iface gdc.Interface) {

  className := StringNameFromStr("BoxMesh")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_size")
    defer methodName.Destroy()
    ptrsForBoxMesh.fnSetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
  }
  {
    methodName := StringNameFromStr("get_size")
    defer methodName.Destroy()
    ptrsForBoxMesh.fnGetSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
  }
  {
    methodName := StringNameFromStr("set_subdivide_width")
    defer methodName.Destroy()
    ptrsForBoxMesh.fnSetSubdivideWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_subdivide_width")
    defer methodName.Destroy()
    ptrsForBoxMesh.fnGetSubdivideWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_subdivide_height")
    defer methodName.Destroy()
    ptrsForBoxMesh.fnSetSubdivideHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_subdivide_height")
    defer methodName.Destroy()
    ptrsForBoxMesh.fnGetSubdivideHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_subdivide_depth")
    defer methodName.Destroy()
    ptrsForBoxMesh.fnSetSubdivideDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_subdivide_depth")
    defer methodName.Destroy()
    ptrsForBoxMesh.fnGetSubdivideDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type BoxMesh struct {
  PrimitiveMesh
}

func (me *BoxMesh) BaseClass() string {
  return "BoxMesh"
}

func NewBoxMesh() *BoxMesh {
  str := StringNameFromStr("BoxMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &BoxMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *BoxMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BoxMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BoxMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *BoxMesh) SetSize(size Vector3, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxMesh.fnSetSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoxMesh) GetSize() Vector3 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxMesh.fnGetSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BoxMesh) SetSubdivideWidth(subdivide int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&subdivide) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxMesh.fnSetSubdivideWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoxMesh) GetSubdivideWidth() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxMesh.fnGetSubdivideWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BoxMesh) SetSubdivideHeight(divisions int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&divisions) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxMesh.fnSetSubdivideHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoxMesh) GetSubdivideHeight() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxMesh.fnGetSubdivideHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BoxMesh) SetSubdivideDepth(divisions int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&divisions) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxMesh.fnSetSubdivideDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoxMesh) GetSubdivideDepth() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoxMesh.fnGetSubdivideDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
