// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AESContext struct {
  RefCounted
}

func (me *AESContext) BaseClass() string {
  return "AESContext"
}

func NewAESContext() *AESContext {
  str := StringNameFromStr("AESContext") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AESContext{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AESContextMode int
const (
  AESContextModeModeEcbEncrypt AESContextMode = 0
  AESContextModeModeEcbDecrypt AESContextMode = 1
  AESContextModeModeCbcEncrypt AESContextMode = 2
  AESContextModeModeCbcDecrypt AESContextMode = 3
  AESContextModeModeMax AESContextMode = 4
)

func (me *AESContext) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AESContext) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AESContext) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AESContext) Start(mode AESContextMode, key PackedByteArray, iv PackedByteArray, ) Error {
  classNameV := StringNameFromStr("AESContext")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3122411423) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(iv.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AESContext) Update(src PackedByteArray, ) PackedByteArray {
  classNameV := StringNameFromStr("AESContext")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 527836100) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(src.AsCTypePtr()), }
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AESContext) GetIvState() PackedByteArray {
  classNameV := StringNameFromStr("AESContext")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_iv_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2115431945) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AESContext) Finish()  {
  classNameV := StringNameFromStr("AESContext")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("finish")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
