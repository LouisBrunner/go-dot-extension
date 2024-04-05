// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForOggPacketSequenceList struct {
  fnSetPacketData gdc.MethodBindPtr
  fnGetPacketData gdc.MethodBindPtr
  fnSetPacketGranulePositions gdc.MethodBindPtr
  fnGetPacketGranulePositions gdc.MethodBindPtr
  fnSetSamplingRate gdc.MethodBindPtr
  fnGetSamplingRate gdc.MethodBindPtr
  fnGetLength gdc.MethodBindPtr
}

var ptrsForOggPacketSequence ptrsForOggPacketSequenceList

func initOggPacketSequencePtrs(iface gdc.Interface) {

  className := StringNameFromStr("OggPacketSequence")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_packet_data")
    defer methodName.Destroy()
    ptrsForOggPacketSequence.fnSetPacketData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_packet_data")
    defer methodName.Destroy()
    ptrsForOggPacketSequence.fnGetPacketData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("set_packet_granule_positions")
    defer methodName.Destroy()
    ptrsForOggPacketSequence.fnSetPacketGranulePositions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3709968205))
  }
  {
    methodName := StringNameFromStr("get_packet_granule_positions")
    defer methodName.Destroy()
    ptrsForOggPacketSequence.fnGetPacketGranulePositions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 235988956))
  }
  {
    methodName := StringNameFromStr("set_sampling_rate")
    defer methodName.Destroy()
    ptrsForOggPacketSequence.fnSetSamplingRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_sampling_rate")
    defer methodName.Destroy()
    ptrsForOggPacketSequence.fnGetSamplingRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_length")
    defer methodName.Destroy()
    ptrsForOggPacketSequence.fnGetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&packet_data) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOggPacketSequence.fnSetPacketData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OggPacketSequence) GetPacketData() []Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOggPacketSequence.fnGetPacketData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Array](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *OggPacketSequence) SetPacketGranulePositions(granule_positions PackedInt64Array, )  {
  cargs := []gdc.ConstTypePtr{granule_positions.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOggPacketSequence.fnSetPacketGranulePositions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OggPacketSequence) GetPacketGranulePositions() PackedInt64Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt64Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOggPacketSequence.fnGetPacketGranulePositions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OggPacketSequence) SetSamplingRate(sampling_rate float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&sampling_rate) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOggPacketSequence.fnSetSamplingRate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OggPacketSequence) GetSamplingRate() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOggPacketSequence.fnGetSamplingRate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OggPacketSequence) GetLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOggPacketSequence.fnGetLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
