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

type OggPacketSequence struct {
  Resource
}

func (me *OggPacketSequence) BaseClass() string {
  return "OggPacketSequence"
}

func NewOggPacketSequence() *OggPacketSequence {
  str := StringNameFromStr("OggPacketSequence") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OggPacketSequence{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *OggPacketSequence) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OggPacketSequence) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OggPacketSequence) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OggPacketSequence) SetPacketData(packet_data []Array, )  {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_packet_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&packet_data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OggPacketSequence) GetPacketData() []Array {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[Array](ret)
}

func  (me *OggPacketSequence) SetPacketGranulePositions(granule_positions PackedInt64Array, )  {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_packet_granule_positions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3709968205) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{granule_positions.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OggPacketSequence) GetPacketGranulePositions() PackedInt64Array {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_granule_positions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 235988956) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OggPacketSequence) SetSamplingRate(sampling_rate float64, )  {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sampling_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sampling_rate) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OggPacketSequence) GetSamplingRate() float64 {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sampling_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OggPacketSequence) GetLength() float64 {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
