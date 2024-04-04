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

type PrimitiveMesh struct {
  Mesh
}

func (me *PrimitiveMesh) BaseClass() string {
  return "PrimitiveMesh"
}

func NewPrimitiveMesh() *PrimitiveMesh {
  str := StringNameFromStr("PrimitiveMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PrimitiveMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PrimitiveMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PrimitiveMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PrimitiveMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PrimitiveMesh) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{material.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PrimitiveMesh) GetMaterial() Material {
  classNameV := StringNameFromStr("PrimitiveMesh")
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

func  (me *PrimitiveMesh) GetMeshArrays() Array {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh_arrays")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PrimitiveMesh) SetCustomAabb(aabb AABB, )  {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{aabb.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PrimitiveMesh) GetCustomAabb() AABB {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PrimitiveMesh) SetFlipFaces(flip_faces bool, )  {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_faces) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PrimitiveMesh) GetFlipFaces() bool {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flip_faces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PrimitiveMesh) SetAddUv2(add_uv2 bool, )  {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_add_uv2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&add_uv2) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PrimitiveMesh) GetAddUv2() bool {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_add_uv2")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PrimitiveMesh) SetUv2Padding(uv2_padding float64, )  {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_uv2_padding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&uv2_padding) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PrimitiveMesh) GetUv2Padding() float64 {
  classNameV := StringNameFromStr("PrimitiveMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_uv2_padding")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
