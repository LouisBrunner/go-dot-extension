// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeParticleMeshEmitter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParticleMeshEmitter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParticleMeshEmitter) BaseClass() string {
  return "VisualShaderNodeParticleMeshEmitter"
}



// Enums

func (me *VisualShaderNodeParticleMeshEmitter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleMeshEmitter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleMeshEmitter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParticleMeshEmitter) SetMesh(mesh Mesh, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMeshEmitter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 194775623) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(mesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParticleMeshEmitter) GetMesh() Mesh {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMeshEmitter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1808005922) // FIXME: should cache?
  var ret Mesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeParticleMeshEmitter) SetUseAllSurfaces(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMeshEmitter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_all_surfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParticleMeshEmitter) IsUseAllSurfaces() bool {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMeshEmitter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_use_all_surfaces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeParticleMeshEmitter) SetSurfaceIndex(surface_index int, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMeshEmitter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_surface_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&surface_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParticleMeshEmitter) GetSurfaceIndex() int {
  classNameV := StringNameFromStr("VisualShaderNodeParticleMeshEmitter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surface_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeParticleMeshEmitter) GetPropMesh() Mesh {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParticleMeshEmitter) SetPropMesh(value Mesh) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParticleMeshEmitter) GetPropUseAllSurfaces() bool {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParticleMeshEmitter) SetPropUseAllSurfaces(value bool) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParticleMeshEmitter) GetPropSurfaceIndex() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParticleMeshEmitter) SetPropSurfaceIndex(value int) {
  panic("TODO: implement")
}