// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type BackBufferCopy struct {
  Node2D
}

func (me *BackBufferCopy) BaseClass() string {
  return "BackBufferCopy"
}

func NewBackBufferCopy() *BackBufferCopy {
  str := StringNameFromStr("BackBufferCopy") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &BackBufferCopy{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type BackBufferCopyCopyMode int
const (
  BackBufferCopyCopyModeCopyModeDisabled BackBufferCopyCopyMode = 0
  BackBufferCopyCopyModeCopyModeRect BackBufferCopyCopyMode = 1
  BackBufferCopyCopyModeCopyModeViewport BackBufferCopyCopyMode = 2
)

func (me *BackBufferCopy) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BackBufferCopy) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BackBufferCopy) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *BackBufferCopy) SetRect(rect Rect2, )  {
  classNameV := StringNameFromStr("BackBufferCopy")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2046264180) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BackBufferCopy) GetRect() Rect2 {
  classNameV := StringNameFromStr("BackBufferCopy")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BackBufferCopy) SetCopyMode(copy_mode BackBufferCopyCopyMode, )  {
  classNameV := StringNameFromStr("BackBufferCopy")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_copy_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1713538590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&copy_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BackBufferCopy) GetCopyMode() BackBufferCopyCopyMode {
  classNameV := StringNameFromStr("BackBufferCopy")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_copy_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3271169440) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret BackBufferCopyCopyMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
