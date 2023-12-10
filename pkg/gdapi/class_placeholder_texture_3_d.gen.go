// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PlaceholderTexture3D struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderTexture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderTexture3D) BaseClass() string {
  return "PlaceholderTexture3D"
}



// Enums

func (me *PlaceholderTexture3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderTexture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTexture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PlaceholderTexture3D) SetSize(size Vector3i, )  {
  classNameV := StringNameFromStr("PlaceholderTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 560364750) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PlaceholderTexture3D) GetSize() Vector3i {
  classNameV := StringNameFromStr("PlaceholderTexture3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2785653706) // FIXME: should cache?
  var ret Vector3i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PlaceholderTexture3D) GetPropSize() Vector3i {
  panic("TODO: implement")
}

func (me *PlaceholderTexture3D) SetPropSize(value Vector3i) {
  panic("TODO: implement")
}