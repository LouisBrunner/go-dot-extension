// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AESContext struct {
  obj gdc.ObjectPtr
}

func (me *AESContext) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AESContext) BaseClass() string {
  return "AESContext"
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

func  (me *AESContext) Start(mode AESContextMode, key PackedByteArray, iv PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *AESContext) Update(src PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *AESContext) GetIvState()  {
  panic("TODO: implement")
}

func  (me *AESContext) Finish()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
