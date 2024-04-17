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

type ptrsForAudioStreamPlaybackInteractiveList struct {
	fnSwitchToClipByName gdc.MethodBindPtr
	fnSwitchToClip       gdc.MethodBindPtr
}

var ptrsForAudioStreamPlaybackInteractive ptrsForAudioStreamPlaybackInteractiveList

func initAudioStreamPlaybackInteractivePtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamPlaybackInteractive")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("switch_to_clip_by_name")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaybackInteractive.fnSwitchToClipByName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("switch_to_clip")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaybackInteractive.fnSwitchToClip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}

}

type AudioStreamPlaybackInteractive struct {
	AudioStreamPlayback
}

func (me *AudioStreamPlaybackInteractive) BaseClass() string {
	return "AudioStreamPlaybackInteractive"
}

func NewAudioStreamPlaybackInteractive() *AudioStreamPlaybackInteractive {
	str := StringNameFromStr("AudioStreamPlaybackInteractive") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamPlaybackInteractive{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioStreamPlaybackInteractive) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamPlaybackInteractive) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaybackInteractive) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamPlaybackInteractive) SwitchToClipByName(clip_name StringName) {
	cargs := []gdc.ConstTypePtr{clip_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaybackInteractive.fnSwitchToClipByName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioStreamPlaybackInteractive) SwitchToClip(clip_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaybackInteractive.fnSwitchToClip), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
