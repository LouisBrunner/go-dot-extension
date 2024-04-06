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

type ptrsForAudioEffectAmplifyList struct {
	fnSetVolumeDb gdc.MethodBindPtr
	fnGetVolumeDb gdc.MethodBindPtr
}

var ptrsForAudioEffectAmplify ptrsForAudioEffectAmplifyList

func initAudioEffectAmplifyPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectAmplify")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_volume_db")
		defer methodName.Destroy()
		ptrsForAudioEffectAmplify.fnSetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_volume_db")
		defer methodName.Destroy()
		ptrsForAudioEffectAmplify.fnGetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
}

type AudioEffectAmplify struct {
	AudioEffect
}

func (me *AudioEffectAmplify) BaseClass() string {
	return "AudioEffectAmplify"
}

func NewAudioEffectAmplify() *AudioEffectAmplify {
	str := StringNameFromStr("AudioEffectAmplify") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectAmplify{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectAmplify) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectAmplify) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectAmplify) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioEffectAmplify) SetVolumeDb(volume float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&volume)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectAmplify.fnSetVolumeDb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioEffectAmplify) GetVolumeDb() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioEffectAmplify.fnGetVolumeDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
