// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type HMACContext struct {
  obj gdc.ObjectPtr
}

func (me *HMACContext) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HMACContext) BaseClass() string {
  return "HMACContext"
}



// Enums

func (me *HMACContext) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HMACContext) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HMACContext) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *HMACContext) Start(hash_type HashingContextHashType, key PackedByteArray, ) Error {
  classNameV := StringNameFromStr("HMACContext")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3537364598) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hash_type), gdc.ConstTypePtr(key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *HMACContext) Update(data PackedByteArray, ) Error {
  classNameV := StringNameFromStr("HMACContext")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 680677267) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *HMACContext) Finish() PackedByteArray {
  classNameV := StringNameFromStr("HMACContext")
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