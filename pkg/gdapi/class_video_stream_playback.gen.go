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

func  (me *VideoStreamPlayback) MixAudio(num_frames int, buffer PackedFloat32Array, offset int, ) int {
  classNameV := StringNameFromStr("VideoStreamPlayback")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("mix_audio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 93876830) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&num_frames), gdc.ConstTypePtr(buffer.AsCTypePtr()), gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
