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

type ptrsForAudioStreamPlaybackResampledList struct {
	fnXMixResampled          gdc.MethodBindPtr
	fnXGetStreamSamplingRate gdc.MethodBindPtr
	fnBeginResample          gdc.MethodBindPtr
}

var ptrsForAudioStreamPlaybackResampled ptrsForAudioStreamPlaybackResampledList

func initAudioStreamPlaybackResampledPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamPlaybackResampled")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("begin_resample")
		defer methodName.Destroy()
		ptrsForAudioStreamPlaybackResampled.fnBeginResample = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type AudioStreamPlaybackResampled struct {
	AudioStreamPlayback
}

func (me *AudioStreamPlaybackResampled) BaseClass() string {
	return "AudioStreamPlaybackResampled"
}

func NewAudioStreamPlaybackResampled() *AudioStreamPlaybackResampled {
	str := StringNameFromStr("AudioStreamPlaybackResampled") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamPlaybackResampled{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioStreamPlaybackResampled) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamPlaybackResampled) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaybackResampled) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStreamPlaybackResampled) BeginResample() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaybackResampled.fnBeginResample), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
