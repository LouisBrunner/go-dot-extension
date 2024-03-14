// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OggPacketSequence struct {
  Resource
}

func (me *OggPacketSequence) BaseClass() string {
  return "OggPacketSequence"
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

func  (me *OggPacketSequence) SetPacketData(packet_data Array, )  {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_packet_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(packet_data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OggPacketSequence) GetPacketData() Array {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OggPacketSequence) SetPacketGranulePositions(granule_positions PackedInt64Array, )  {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_packet_granule_positions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3709968205) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(granule_positions.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OggPacketSequence) GetPacketGranulePositions() PackedInt64Array {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_granule_positions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 235988956) // FIXME: should cache?
  var ret PackedInt64Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OggPacketSequence) SetSamplingRate(sampling_rate float32, )  {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sampling_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sampling_rate), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OggPacketSequence) GetSamplingRate() float32 {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sampling_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OggPacketSequence) GetLength() float32 {
  classNameV := StringNameFromStr("OggPacketSequence")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
