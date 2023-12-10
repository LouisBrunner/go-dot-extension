// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CSGMesh3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGMesh3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGMesh3D) BaseClass() string {
  return "CSGMesh3D"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGMesh3D) GetMesh() Mesh {
  classNameV := StringNameFromStr("CSGMesh3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4081188045) // FIXME: should cache?
  var ret Mesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGMesh3D) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("CSGMesh3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGMesh3D) GetMaterial() Material {
  classNameV := StringNameFromStr("CSGMesh3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 5934680) // FIXME: should cache?
  var ret Material
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CSGMesh3D) GetPropMesh() Mesh {
  panic("TODO: implement")
}

func (me *CSGMesh3D) SetPropMesh(value Mesh) {
  panic("TODO: implement")
}

func (me *CSGMesh3D) GetPropMaterial() any {
  panic("TODO: implement")
}

func (me *CSGMesh3D) SetPropMaterial(value any) {
  panic("TODO: implement")
}