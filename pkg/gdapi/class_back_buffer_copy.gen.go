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



// Enums

type BackBufferCopyCopyMode int
const (
  BackBufferCopyCopyModeCopyModeDisabled BackBufferCopyCopyMode = 0
  BackBufferCopyCopyModeCopyModeRect BackBufferCopyCopyMode = 1
  BackBufferCopyCopyModeCopyModeViewport BackBufferCopyCopyMode = 2
)

func (me *BackBufferCopy) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BackBufferCopy) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *BackBufferCopy) SetRect(rect Rect2, )  {
  panic("TODO: implement")
}

func  (me *BackBufferCopy) GetRect()  {
  panic("TODO: implement")
}

func  (me *BackBufferCopy) SetCopyMode(copy_mode BackBufferCopyCopyMode, )  {
  panic("TODO: implement")
}

func  (me *BackBufferCopy) GetCopyMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
