// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OggPacketSequence struct {
  obj gdc.ObjectPtr
}

func (me *OggPacketSequence) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OggPacketSequence) BaseClass() string {
  return "OggPacketSequence"
}



// Enums

func (me *OggPacketSequence) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OggPacketSequence) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *OggPacketSequence) SetPacketData(packet_data Array, )  {
  panic("TODO: implement")
}

func  (me *OggPacketSequence) GetPacketData()  {
  panic("TODO: implement")
}

func  (me *OggPacketSequence) SetPacketGranulePositions(granule_positions PackedInt64Array, )  {
  panic("TODO: implement")
}

func  (me *OggPacketSequence) GetPacketGranulePositions()  {
  panic("TODO: implement")
}

func  (me *OggPacketSequence) SetSamplingRate(sampling_rate float32, )  {
  panic("TODO: implement")
}

func  (me *OggPacketSequence) GetSamplingRate()  {
  panic("TODO: implement")
}

func  (me *OggPacketSequence) GetLength()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
