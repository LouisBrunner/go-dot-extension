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

type ptrsForAudioEffectPhaserList struct {
	fnSetRangeMinHz gdc.MethodBindPtr
	fnGetRangeMinHz gdc.MethodBindPtr
	fnSetRangeMaxHz gdc.MethodBindPtr
	fnGetRangeMaxHz gdc.MethodBindPtr
	fnSetRateHz     gdc.MethodBindPtr
	fnGetRateHz     gdc.MethodBindPtr
	fnSetFeedback   gdc.MethodBindPtr
	fnGetFeedback   gdc.MethodBindPtr
	fnSetDepth      gdc.MethodBindPtr
	fnGetDepth      gdc.MethodBindPtr
}

var ptrsForAudioEffectPhaser ptrsForAudioEffectPhaserList

func initAudioEffectPhaserPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectPhaser")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_range_min_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnSetRangeMinHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_range_min_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnGetRangeMinHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_range_max_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnSetRangeMaxHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_range_max_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnGetRangeMaxHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_rate_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnSetRateHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_rate_hz")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnGetRateHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_feedback")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnSetFeedback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_feedback")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnGetFeedback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_depth")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnSetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_depth")
		defer methodName.Destroy()
		ptrsForAudioEffectPhaser.fnGetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type AudioEffectPhaser struct {
	AudioEffect
}

func (me *AudioEffectPhaser) BaseClass() string {
	return "AudioEffectPhaser"
}

func NewAudioEffectPhaser() *AudioEffectPhaser {
	str := StringNameFromStr("AudioEffectPhaser") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectPhaser{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectPhaser) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectPhaser) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectPhaser) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectPhaser) SetRangeMinHz(hz float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnSetRangeMinHz), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectPhaser) GetRangeMinHz() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnGetRangeMinHz), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectPhaser) SetRangeMaxHz(hz float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnSetRangeMaxHz), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectPhaser) GetRangeMaxHz() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnGetRangeMaxHz), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectPhaser) SetRateHz(hz float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnSetRateHz), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectPhaser) GetRateHz() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnGetRateHz), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectPhaser) SetFeedback(fbk float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fbk)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnSetFeedback), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectPhaser) GetFeedback() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnGetFeedback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectPhaser) SetDepth(depth float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&depth)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnSetDepth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectPhaser) GetDepth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectPhaser.fnGetDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
