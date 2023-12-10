// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type X509Certificate struct {
  obj gdc.ObjectPtr
}

func (me *X509Certificate) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *X509Certificate) BaseClass() string {
  return "X509Certificate"
}



// Enums

func (me *X509Certificate) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *X509Certificate) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *X509Certificate) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *X509Certificate) Save(path String, ) Error {
  classNameV := StringNameFromStr("X509Certificate")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *X509Certificate) Load(path String, ) Error {
  classNameV := StringNameFromStr("X509Certificate")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *X509Certificate) SaveToString() String {
  classNameV := StringNameFromStr("X509Certificate")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_to_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *X509Certificate) LoadFromString(string_ String, ) Error {
  classNameV := StringNameFromStr("X509Certificate")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties