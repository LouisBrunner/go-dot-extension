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

type ptrsForAudioStreamPlaybackPlaylistList struct {
}

var ptrsForAudioStreamPlaybackPlaylist ptrsForAudioStreamPlaybackPlaylistList

func initAudioStreamPlaybackPlaylistPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStreamPlaybackPlaylist")
	defer className.Destroy()

}

type AudioStreamPlaybackPlaylist struct {
	AudioStreamPlayback
}

func (me *AudioStreamPlaybackPlaylist) BaseClass() string {
	return "AudioStreamPlaybackPlaylist"
}

func NewAudioStreamPlaybackPlaylist() *AudioStreamPlaybackPlaylist {
	str := StringNameFromStr("AudioStreamPlaybackPlaylist") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStreamPlaybackPlaylist{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioStreamPlaybackPlaylist) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStreamPlaybackPlaylist) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaybackPlaylist) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
