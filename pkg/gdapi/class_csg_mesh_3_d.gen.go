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

type ptrsForCSGMesh3DList struct {
  fnSetMesh gdc.MethodBindPtr
  fnGetMesh gdc.MethodBindPtr
  fnSetMaterial gdc.MethodBindPtr
  fnGetMaterial gdc.MethodBindPtr
}

var ptrsForCSGMesh3D ptrsForCSGMesh3DList

func initCSGMesh3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CSGMesh3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_mesh")
    defer methodName.Destroy()
    ptrsForCSGMesh3D.fnSetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 194775623))
  }
  {
    methodName := StringNameFromStr("get_mesh")
    defer methodName.Destroy()
    ptrsForCSGMesh3D.fnGetMesh = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4081188045))
  }
  {
    methodName := StringNameFromStr("set_material")
    defer methodName.Destroy()
    ptrsForCSGMesh3D.fnSetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2757459619))
  }
  {
    methodName := StringNameFromStr("get_material")
    defer methodName.Destroy()
    ptrsForCSGMesh3D.fnGetMaterial = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 5934680))
  }
}

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
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGMesh3D.fnSetMesh), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGMesh3D) GetMesh() Mesh {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMesh()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGMesh3D.fnGetMesh), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CSGMesh3D) SetMaterial(material Material, )  {
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGMesh3D.fnSetMaterial), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CSGMesh3D) GetMaterial() Material {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewMaterial()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCSGMesh3D.fnGetMaterial), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
