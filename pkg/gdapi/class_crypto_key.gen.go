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

type CryptoKey struct {
  Resource
}

func (me *CryptoKey) BaseClass() string {
  return "CryptoKey"
}

func NewCryptoKey() *CryptoKey {
  str := StringNameFromStr("CryptoKey") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CryptoKey{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CryptoKey) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CryptoKey) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CryptoKey) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CryptoKey) Save(path String, public_only bool, ) Error {
  classNameV := StringNameFromStr("CryptoKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 885841341) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&public_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&public_only)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CryptoKey) Load(path String, public_only bool, ) Error {
  classNameV := StringNameFromStr("CryptoKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 885841341) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&public_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&public_only)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *CryptoKey) IsPublicOnly() bool {
  classNameV := StringNameFromStr("CryptoKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_public_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CryptoKey) SaveToString(public_only bool, ) String {
  classNameV := StringNameFromStr("CryptoKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_to_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 32795936) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&public_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&public_only)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CryptoKey) LoadFromString(string_key String, public_only bool, ) Error {
  classNameV := StringNameFromStr("CryptoKey")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 885841341) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{string_key.AsCTypePtr(), gdc.ConstTypePtr(&public_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&public_only)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

// Signals
