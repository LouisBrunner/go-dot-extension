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

type ptrsForAudioStreamList struct {
	fnXInstantiatePlayback gdc.MethodBindPtr
	fnXGetStreamName       gdc.MethodBindPtr
	fnXGetLength           gdc.MethodBindPtr
	fnXIsMonophonic        gdc.MethodBindPtr
	fnXGetBpm              gdc.MethodBindPtr
	fnXGetBeatCount        gdc.MethodBindPtr
	fnGetLength            gdc.MethodBindPtr
	fnIsMonophonic         gdc.MethodBindPtr
	fnInstantiatePlayback  gdc.MethodBindPtr
}

var ptrsForAudioStream ptrsForAudioStreamList

func initAudioStreamPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioStream")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_length")
		defer methodName.Destroy()
		ptrsForAudioStream.fnGetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("is_monophonic")
		defer methodName.Destroy()
		ptrsForAudioStream.fnIsMonophonic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("instantiate_playback")
		defer methodName.Destroy()
		ptrsForAudioStream.fnInstantiatePlayback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 210135309))
	}

}

type AudioStream struct {
	Resource
}

func (me *AudioStream) BaseClass() string {
	return "AudioStream"
}

func NewAudioStream() *AudioStream {
	str := StringNameFromStr("AudioStream") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioStream{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioStream) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioStream) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioStream) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioStream) GetLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStream.fnGetLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStream) IsMonophonic() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStream.fnIsMonophonic), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AudioStream) InstantiatePlayback() AudioStreamPlayback {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAudioStreamPlayback()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStream.fnInstantiatePlayback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
