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
  classNameV := StringNameFromStr("ArrayOccluder3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3233972621) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), indices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayOccluder3D) SetVertices(vertices PackedVector3Array, )  {
  classNameV := StringNameFromStr("ArrayOccluder3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_vertices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 334873810) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{vertices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ArrayOccluder3D) SetIndices(indices PackedInt32Array, )  {
  classNameV := StringNameFromStr("ArrayOccluder3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_indices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614634198) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{indices.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
