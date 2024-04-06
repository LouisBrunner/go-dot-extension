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

type ptrsForAudioEffectFilterList struct {
	fnSetCutoff    gdc.MethodBindPtr
	fnGetCutoff    gdc.MethodBindPtr
	fnSetResonance gdc.MethodBindPtr
	fnGetResonance gdc.MethodBindPtr
	fnSetGain      gdc.MethodBindPtr
	fnGetGain      gdc.MethodBindPtr
	fnSetDb        gdc.MethodBindPtr
	fnGetDb        gdc.MethodBindPtr
}

var ptrsForAudioEffectFilter ptrsForAudioEffectFilterList

func initAudioEffectFilterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectFilter")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_cutoff")
		defer methodName.Destroy()
		ptrsForAudioEffectFilter.fnSetCutoff = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_cutoff")
		defer methodName.Destroy()
		ptrsForAudioEffectFilter.fnGetCutoff = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_resonance")
		defer methodName.Destroy()
		ptrsForAudioEffectFilter.fnSetResonance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_resonance")
		defer methodName.Destroy()
		ptrsForAudioEffectFilter.fnGetResonance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_gain")
		defer methodName.Destroy()
		ptrsForAudioEffectFilter.fnSetGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_gain")
		defer methodName.Destroy()
		ptrsForAudioEffectFilter.fnGetGain = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_db")
		defer methodName.Destroy()
		ptrsForAudioEffectFilter.fnSetDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 771740901))
	}
	{
		methodName := StringNameFromStr("get_db")
		defer methodName.Destroy()
		ptrsForAudioEffectFilter.fnGetDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3981721890))
	}

}

type AudioEffectFilter struct {
	AudioEffect
}

func (me *AudioEffectFilter) BaseClass() string {
	return "AudioEffectFilter"
}

func NewAudioEffectFilter() *AudioEffectFilter {
	str := StringNameFromStr("AudioEffectFilter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectFilter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AudioEffectFilterFilterDB int

const (
	AudioEffectFilterFilterDBFilter6Db  AudioEffectFilterFilterDB = 0
	AudioEffectFilterFilterDBFilter12Db AudioEffectFilterFilterDB = 1
	AudioEffectFilterFilterDBFilter18Db AudioEffectFilterFilterDB = 2
	AudioEffectFilterFilterDBFilter24Db AudioEffectFilterFilterDB = 3
)

func (me *AudioEffectFilter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectFilter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectFilter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectFilter) SetCutoff(freq float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freq)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectFilter.fnSetCutoff), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectFilter) GetCutoff() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectFilter.fnGetCutoff), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectFilter) SetResonance(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectFilter.fnSetResonance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectFilter) GetResonance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectFilter.fnGetResonance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectFilter) SetGain(amount float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectFilter.fnSetGain), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectFilter) GetGain() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectFilter.fnGetGain), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioEffectFilter) SetDb(amount AudioEffectFilterFilterDB) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectFilter.fnSetDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectFilter) GetDb() AudioEffectFilterFilterDB {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AudioEffectFilterFilterDB

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectFilter.fnGetDb), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
