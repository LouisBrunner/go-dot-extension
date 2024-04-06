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

type ptrsForAudioEffectHighShelfFilterList struct {
}

var ptrsForAudioEffectHighShelfFilter ptrsForAudioEffectHighShelfFilterList

func initAudioEffectHighShelfFilterPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioEffectHighShelfFilter")
	defer className.Destroy()

}

type AudioEffectHighShelfFilter struct {
	AudioEffectFilter
}

func (me *AudioEffectHighShelfFilter) BaseClass() string {
	return "AudioEffectHighShelfFilter"
}

func NewAudioEffectHighShelfFilter() *AudioEffectHighShelfFilter {
	str := StringNameFromStr("AudioEffectHighShelfFilter") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioEffectHighShelfFilter{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioEffectHighShelfFilter) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioEffectHighShelfFilter) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioEffectHighShelfFilter) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
