// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  BackBufferCopyCopyModeCopyModeDisabled BackBufferCopyCopyMode = 0
  BackBufferCopyCopyModeCopyModeRect BackBufferCopyCopyMode = 1
  BackBufferCopyCopyModeCopyModeViewport BackBufferCopyCopyMode = 2
)

func  (me *BackBufferCopy) SetRect(rect Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *BackBufferCopy) GetRect() { // TODO: return value
  // TODO: implement
}

func  (me *BackBufferCopy) SetCopyMode(copy_mode BackBufferCopyCopyMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *BackBufferCopy) GetCopyMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
