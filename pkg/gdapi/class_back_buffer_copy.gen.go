// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BackBufferCopy struct {
  obj gdc.ObjectPtr
}

func (me *BackBufferCopy) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BackBufferCopy) BaseClass() string {
  return "BackBufferCopy"
}

type BackBufferCopyCopyMode int
const (
  BackBufferCopyCopyModeDisabled BackBufferCopyCopyMode = 0
  BackBufferCopyCopyModeRect BackBufferCopyCopyMode = 1
  BackBufferCopyCopyModeViewport BackBufferCopyCopyMode = 2
)
