// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HashingContext struct {
  obj gdc.ObjectPtr
}

func (me *HashingContext) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HashingContext) BaseClass() string {
  return "HashingContext"
}

type HashingContextHashType int
const (
  HashingContextHashMd5 HashingContextHashType = 0
  HashingContextHashSha1 HashingContextHashType = 1
  HashingContextHashSha256 HashingContextHashType = 2
)
