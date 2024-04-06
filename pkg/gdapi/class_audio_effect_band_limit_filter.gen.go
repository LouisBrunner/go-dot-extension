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

type ptrsForAudioEffectBandLimitFilterList struct {
}

var ptrsForAudioEffectBandLimitFilter ptrsForAudioEffectBandLimitFilterList

func initAudioEffectBandLimitFilterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectBandLimitFilter")
	defer className.Destroy()
}

type AudioEffectBandLimitFilter struct {
	AudioEffectFilter
}

func (me *AudioEffectBandLimitFilter) BaseClass() string {
	return "AudioEffectBandLimitFilter"
}

func NewAudioEffectBandLimitFilter() *AudioEffectBandLimitFilter {
	str := StringNameFromStr("AudioEffectBandLimitFilter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectBandLimitFilter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectBandLimitFilter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectBandLimitFilter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectBandLimitFilter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
