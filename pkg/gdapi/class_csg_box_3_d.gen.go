// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CSGBox3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGBox3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGBox3D) BaseClass() string {
  return "CSGBox3D"
}



// Enums

func (me *CSGBox3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGBox3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGBox3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CSGBox3D) SetSize(size Vector3, )  {
  classNameV := StringNameFromStr("CSGBox3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGBox3D) GetSize() Vector3 {
  classNameV := StringNameFromStr("CSGBox3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CSGBox3D) SetMaterial(material Material, )  {
  classNameV := StringNameFromStr("CSGBox3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_material")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2757459619) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(material.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CSGBox3D) GetMaterial() Material {
  classNameV := StringNameFromStr("CSGBox3D")
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

func (me *CSGBox3D) GetPropSize() Vector3 {
  panic("TODO: implement")
}

func (me *CSGBox3D) SetPropSize(value Vector3) {
  panic("TODO: implement")
}

func (me *CSGBox3D) GetPropMaterial() any {
  panic("TODO: implement")
}

func (me *CSGBox3D) SetPropMaterial(value any) {
  panic("TODO: implement")
}