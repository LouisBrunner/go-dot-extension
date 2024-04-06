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

type ptrsForAudioListener2DList struct {
	fnMakeCurrent  gdc.MethodBindPtr
	fnClearCurrent gdc.MethodBindPtr
	fnIsCurrent    gdc.MethodBindPtr
}

var ptrsForAudioListener2D ptrsForAudioListener2DList

func initAudioListener2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AudioListener2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("make_current")
		defer methodName.Destroy()
		ptrsForAudioListener2D.fnMakeCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear_current")
		defer methodName.Destroy()
		ptrsForAudioListener2D.fnClearCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_current")
		defer methodName.Destroy()
		ptrsForAudioListener2D.fnIsCurrent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type AudioListener2D struct {
	Node2D
}

func (me *AudioListener2D) BaseClass() string {
	return "AudioListener2D"
}

func NewAudioListener2D() *AudioListener2D {
	str := StringNameFromStr("AudioListener2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AudioListener2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AudioListener2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AudioListener2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AudioListener2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AudioListener2D) MakeCurrent() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioListener2D.fnMakeCurrent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioListener2D) ClearCurrent() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioListener2D.fnClearCurrent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AudioListener2D) IsCurrent() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioListener2D.fnIsCurrent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
