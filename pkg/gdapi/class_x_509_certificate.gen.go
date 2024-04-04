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

type X509Certificate struct {
  Resource
}

func (me *X509Certificate) BaseClass() string {
  return "X509Certificate"
}

func NewX509Certificate() *X509Certificate {
  str := StringNameFromStr("X509Certificate") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &X509Certificate{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *X509Certificate) Load(path String, ) Error {
  classNameV := StringNameFromStr("X509Certificate")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *X509Certificate) SaveToString() String {
  classNameV := StringNameFromStr("X509Certificate")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_to_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *X509Certificate) LoadFromString(string_ String, ) Error {
  classNameV := StringNameFromStr("X509Certificate")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{string_.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

// Signals
