// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  HashingContextHashTypeHashMd5 HashingContextHashType = 0
  HashingContextHashTypeHashSha1 HashingContextHashType = 1
  HashingContextHashTypeHashSha256 HashingContextHashType = 2
)

func  (me *HashingContext) Start(type_ HashingContextHashType, ) { // TODO: return value
  // TODO: implement
}

func  (me *HashingContext) Update(chunk PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *HashingContext) Finish() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
