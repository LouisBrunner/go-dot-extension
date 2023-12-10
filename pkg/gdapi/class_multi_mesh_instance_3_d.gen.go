// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type MultiMeshInstance3D struct {
  obj gdc.ObjectPtr
}

func (me *MultiMeshInstance3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *MultiMeshInstance3D) BaseClass() string {
  return "MultiMeshInstance3D"
}



// Enums

func (me *MultiMeshInstance3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *MultiMeshInstance3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *MultiMeshInstance3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *MultiMeshInstance3D) SetMultimesh(multimesh MultiMesh, )  {
  classNameV := StringNameFromStr("MultiMeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_multimesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2246127404) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(multimesh.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *MultiMeshInstance3D) GetMultimesh() MultiMesh {
  classNameV := StringNameFromStr("MultiMeshInstance3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_multimesh")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1385450523) // FIXME: should cache?
  var ret MultiMesh
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *MultiMeshInstance3D) GetPropMultimesh() MultiMesh {
  panic("TODO: implement")
}

func (me *MultiMeshInstance3D) SetPropMultimesh(value MultiMesh) {
  panic("TODO: implement")
}