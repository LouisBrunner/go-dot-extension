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

type ptrsForAudioEffectEQList struct {
	fnSetBandGainDb gdc.MethodBindPtr
	fnGetBandGainDb gdc.MethodBindPtr
	fnGetBandCount  gdc.MethodBindPtr
}

var ptrsForAudioEffectEQ ptrsForAudioEffectEQList

func initAudioEffectEQPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectEQ")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_band_gain_db")
		defer methodName.Destroy()
		ptrsForAudioEffectEQ.fnSetBandGainDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_band_gain_db")
		defer methodName.Destroy()
		ptrsForAudioEffectEQ.fnGetBandGainDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("get_band_count")
		defer methodName.Destroy()
		ptrsForAudioEffectEQ.fnGetBandCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}

}

type AudioEffectEQ struct {
	AudioEffect
}

func (me *AudioEffectEQ) BaseClass() string {
	return "AudioEffectEQ"
}

func NewAudioEffectEQ() *AudioEffectEQ {
	str := StringNameFromStr("AudioEffectEQ") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectEQ{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectEQ) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectEQ) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectEQ) SetBandGainDb(band_idx int64, volume_db float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&band_idx), gdc.ConstTypePtr(&volume_db)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectEQ.fnSetBandGainDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectEQ) GetBandGainDb(band_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&band_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&band_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectEQ.fnGetBandGainDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectEQ) GetBandCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectEQ.fnGetBandCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
