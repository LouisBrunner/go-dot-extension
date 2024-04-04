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

type CSGMesh3D struct {
  CSGPrimitive3D
}

func (me *CSGMesh3D) BaseClass() string {
  return "CSGMesh3D"
}

func NewCSGMesh3D() *CSGMesh3D {
  str := StringNameFromStr("CSGMesh3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CSGMesh3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CSGMesh3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGMesh3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGMesh3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CSGMesh3D) SetMesh(mesh Mesh, )  {
  classNameV := StringNameFromStr("CSGMesh3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 194775623) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGMesh3D) GetMesh() Mesh {
  classNameV := StringNameFromStr("CSGMesh3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4081188045) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CSGMesh3D) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("CSGMesh3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGMesh3D) GetMaterial() Material {
  classNameV := StringNameFromStr("CSGMesh3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
