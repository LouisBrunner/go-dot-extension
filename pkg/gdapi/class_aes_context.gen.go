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

type AESContextMode int
const (
  AESContextModeModeEcbEncrypt AESContextMode = 0
  AESContextModeModeEcbDecrypt AESContextMode = 1
  AESContextModeModeCbcEncrypt AESContextMode = 2
  AESContextModeModeCbcDecrypt AESContextMode = 3
  AESContextModeModeMax AESContextMode = 4
)

func  (me *AESContext) Start(mode AESContextMode, key PackedByteArray, iv PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *AESContext) Update(src PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *AESContext) GetIvState() { // TODO: return value
  // TODO: implement
}

func  (me *AESContext) Finish() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
