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

type ptrsForArrayOccluder3DList struct {
  fnSetArrays gdc.MethodBindPtr
  fnSetVertices gdc.MethodBindPtr
  fnSetIndices gdc.MethodBindPtr
}

var ptrsForArrayOccluder3D ptrsForArrayOccluder3DList

func initArrayOccluder3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ArrayOccluder3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_arrays")
    defer methodName.Destroy()
    ptrsForArrayOccluder3D.fnSetArrays = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3233972621))
  }
  {
    methodName := StringNameFromStr("set_vertices")
    defer methodName.Destroy()
    ptrsForArrayOccluder3D.fnSetVertices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 334873810))
  }
  {
    methodName := StringNameFromStr("set_indices")
    defer methodName.Destroy()
    ptrsForArrayOccluder3D.fnSetIndices = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
}

type ArrayOccluder3D struct {
  Occluder3D
}

func (me *ArrayOccluder3D) BaseClass() string {
  return "ArrayOccluder3D"
}

func NewArrayOccluder3D() *ArrayOccluder3D {
  str := StringNameFromStr("ArrayOccluder3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ArrayOccluder3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ArrayOccluder3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ArrayOccluder3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ArrayOccluder3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ArrayOccluder3D) SetArrays(vertices PackedVector3Array, indices PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), indices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayOccluder3D.fnSetArrays), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayOccluder3D) SetVertices(vertices PackedVector3Array, )  {
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayOccluder3D.fnSetVertices), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayOccluder3D) SetIndices(indices PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{indices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForArrayOccluder3D.fnSetIndices), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
