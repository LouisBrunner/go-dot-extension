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

type ptrsForAudioStreamPlaybackSynchronizedList struct {
}

var ptrsForAudioStreamPlaybackSynchronized ptrsForAudioStreamPlaybackSynchronizedList

func initAudioStreamPlaybackSynchronizedPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamPlaybackSynchronized")
	defer className.Destroy()

}

type AudioStreamPlaybackSynchronized struct {
	AudioStreamPlayback
}

func (me *AudioStreamPlaybackSynchronized) BaseClass() string {
	return "AudioStreamPlaybackSynchronized"
}

func NewAudioStreamPlaybackSynchronized() *AudioStreamPlaybackSynchronized {
	str := StringNameFromStr("AudioStreamPlaybackSynchronized") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamPlaybackSynchronized{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioStreamPlaybackSynchronized) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamPlaybackSynchronized) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaybackSynchronized) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
