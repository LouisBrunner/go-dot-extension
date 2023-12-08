// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  AESContextModeEcbEncrypt AESContextMode = 0
  AESContextModeEcbDecrypt AESContextMode = 1
  AESContextModeCbcEncrypt AESContextMode = 2
  AESContextModeCbcDecrypt AESContextMode = 3
  AESContextModeMax AESContextMode = 4
)
