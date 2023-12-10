// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type HashingContext struct {
  obj gdc.ObjectPtr
}

func (me *HashingContext) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HashingContext) BaseClass() string {
  return "HashingContext"
}



// Enums

type HashingContextHashType int
const (
  HashingContextHashTypeHashMd5 HashingContextHashType = 0
  HashingContextHashTypeHashSha1 HashingContextHashType = 1
  HashingContextHashTypeHashSha256 HashingContextHashType = 2
)

func (me *HashingContext) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HashingContext) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HashingContext) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *HashingContext) Start(type_ HashingContextHashType, ) Error {
  classNameV := StringNameFromStr("HashingContext")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3940338335) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *HashingContext) Update(chunk PackedByteArray, ) Error {
  classNameV := StringNameFromStr("HashingContext")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(chunk.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *HashingContext) Finish() PackedByteArray {
  classNameV := StringNameFromStr("HashingContext")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("finish")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2115431945) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties