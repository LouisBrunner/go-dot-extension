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

type GLTFMesh struct {
  Resource
}

func (me *GLTFMesh) BaseClass() string {
  return "GLTFMesh"
}

func NewGLTFMesh() *GLTFMesh {
  str := StringNameFromStr("GLTFMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GLTFMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GLTFMesh) GetMesh() ImporterMesh {
  classNameV := StringNameFromStr("GLTFMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3754628756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImporterMesh()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFMesh) SetMesh(mesh ImporterMesh, )  {
  classNameV := StringNameFromStr("GLTFMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2255166972) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{mesh.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFMesh) GetBlendWeights() PackedFloat32Array {
  classNameV := StringNameFromStr("GLTFMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_blend_weights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2445143706) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFMesh) SetBlendWeights(blend_weights PackedFloat32Array, )  {
  classNameV := StringNameFromStr("GLTFMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_blend_weights")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2899603908) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{blend_weights.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFMesh) GetInstanceMaterials() []Material {
  classNameV := StringNameFromStr("GLTFMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_instance_materials")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Material](ret)
}

func  (me *GLTFMesh) SetInstanceMaterials(instance_materials []Material, )  {
  classNameV := StringNameFromStr("GLTFMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_instance_materials")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&instance_materials) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
