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

type ptrsForAnimationNodeStateMachinePlaybackList struct {
	fnTravel                 gdc.MethodBindPtr
	fnStart                  gdc.MethodBindPtr
	fnNext                   gdc.MethodBindPtr
	fnStop                   gdc.MethodBindPtr
	fnIsPlaying              gdc.MethodBindPtr
	fnGetCurrentNode         gdc.MethodBindPtr
	fnGetCurrentPlayPosition gdc.MethodBindPtr
	fnGetCurrentLength       gdc.MethodBindPtr
	fnGetFadingFromNode      gdc.MethodBindPtr
	fnGetTravelPath          gdc.MethodBindPtr
}

var ptrsForAnimationNodeStateMachinePlayback ptrsForAnimationNodeStateMachinePlaybackList

func initAnimationNodeStateMachinePlaybackPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationNodeStateMachinePlayback")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("travel")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnTravel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3823612587))
	}
	{
		methodName := StringNameFromStr("start")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3823612587))
	}
	{
		methodName := StringNameFromStr("next")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnNext = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("stop")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_playing")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnIsPlaying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_current_node")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnGetCurrentNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("get_current_play_position")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnGetCurrentPlayPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_current_length")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnGetCurrentLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_fading_from_node")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnGetFadingFromNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("get_travel_path")
		defer methodName.Destroy()
		ptrsForAnimationNodeStateMachinePlayback.fnGetTravelPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}

}

type AnimationNodeStateMachinePlayback struct {
	Resource
}

func (me *AnimationNodeStateMachinePlayback) BaseClass() string {
	return "AnimationNodeStateMachinePlayback"
}

func NewAnimationNodeStateMachinePlayback() *AnimationNodeStateMachinePlayback {
	str := StringNameFromStr("AnimationNodeStateMachinePlayback") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationNodeStateMachinePlayback{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *AnimationNodeStateMachinePlayback) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationNodeStateMachinePlayback) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeStateMachinePlayback) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimationNodeStateMachinePlayback) Travel(to_node StringName, reset_on_teleport bool) {
	cargs := []gdc.ConstTypePtr{to_node.AsCTypePtr(), gdc.ConstTypePtr(&reset_on_teleport)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnTravel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachinePlayback) Start(node StringName, reset bool) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), gdc.ConstTypePtr(&reset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnStart), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachinePlayback) Next() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnNext), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachinePlayback) Stop() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationNodeStateMachinePlayback) IsPlaying() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnIsPlaying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeStateMachinePlayback) GetCurrentNode() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnGetCurrentNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeStateMachinePlayback) GetCurrentPlayPosition() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnGetCurrentPlayPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeStateMachinePlayback) GetCurrentLength() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnGetCurrentLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationNodeStateMachinePlayback) GetFadingFromNode() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnGetFadingFromNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationNodeStateMachinePlayback) GetTravelPath() []StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationNodeStateMachinePlayback.fnGetTravelPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[StringName](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

// Signals
