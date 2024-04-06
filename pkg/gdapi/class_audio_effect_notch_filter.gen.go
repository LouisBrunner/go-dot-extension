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

type ptrsForAudioEffectNotchFilterList struct {
}

var ptrsForAudioEffectNotchFilter ptrsForAudioEffectNotchFilterList

func initAudioEffectNotchFilterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectNotchFilter")
	defer className.Destroy()
}

type AudioEffectNotchFilter struct {
	AudioEffectFilter
}

func (me *AudioEffectNotchFilter) BaseClass() string {
	return "AudioEffectNotchFilter"
}

func NewAudioEffectNotchFilter() *AudioEffectNotchFilter {
	str := StringNameFromStr("AudioEffectNotchFilter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectNotchFilter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectNotchFilter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectNotchFilter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectNotchFilter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
