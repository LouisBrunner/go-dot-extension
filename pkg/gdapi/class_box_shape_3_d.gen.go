// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type BoxShape3D struct {
  obj gdc.ObjectPtr
}

func (me *BoxShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BoxShape3D) BaseClass() string {
  return "BoxShape3D"
}



// Enums

func (me *BoxShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BoxShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BoxShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *BoxShape3D) SetSize(size Vector3, )  {
  classNameV := StringNameFromStr("BoxShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BoxShape3D) GetSize() Vector3 {
  classNameV := StringNameFromStr("BoxShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *BoxShape3D) GetPropSize() Vector3 {
  panic("TODO: implement")
}

func (me *BoxShape3D) SetPropSize(value Vector3) {
  panic("TODO: implement")
}