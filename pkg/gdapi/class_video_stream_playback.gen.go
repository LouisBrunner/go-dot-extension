// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VideoStreamPlayback struct {
  Resource
}

func (me *VideoStreamPlayback) BaseClass() string {
  return "VideoStreamPlayback"
}

func NewVideoStreamPlayback() *VideoStreamPlayback {
  str := StringNameFromStr("VideoStreamPlayback") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VideoStreamPlayback{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VideoStreamPlayback) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VideoStreamPlayback) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VideoStreamPlayback) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VideoStreamPlayback) MixAudio(num_frames int64, buffer PackedFloat32Array, offset int64, ) int64 {
  classNameV := StringNameFromStr("VideoStreamPlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("mix_audio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 93876830) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&num_frames), gdc.ConstTypePtr(buffer.AsCTypePtr()), gdc.ConstTypePtr(&offset), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
