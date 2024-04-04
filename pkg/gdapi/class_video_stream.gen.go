// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VideoStream struct {
  Resource
}

func (me *VideoStream) BaseClass() string {
  return "VideoStream"
}

func NewVideoStream() *VideoStream {
  str := StringNameFromStr("VideoStream") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VideoStream{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VideoStream) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VideoStream) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VideoStream) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VideoStream) SetFile(file String, )  {
  classNameV := StringNameFromStr("VideoStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{file.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VideoStream) GetFile() String {
  classNameV := StringNameFromStr("VideoStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
