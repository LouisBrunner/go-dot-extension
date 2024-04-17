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

type ptrsForAnimationMixerList struct {
	fnXPostProcessKeyValue             gdc.MethodBindPtr
	fnAddAnimationLibrary              gdc.MethodBindPtr
	fnRemoveAnimationLibrary           gdc.MethodBindPtr
	fnRenameAnimationLibrary           gdc.MethodBindPtr
	fnHasAnimationLibrary              gdc.MethodBindPtr
	fnGetAnimationLibrary              gdc.MethodBindPtr
	fnGetAnimationLibraryList          gdc.MethodBindPtr
	fnHasAnimation                     gdc.MethodBindPtr
	fnGetAnimation                     gdc.MethodBindPtr
	fnGetAnimationList                 gdc.MethodBindPtr
	fnSetActive                        gdc.MethodBindPtr
	fnIsActive                         gdc.MethodBindPtr
	fnSetDeterministic                 gdc.MethodBindPtr
	fnIsDeterministic                  gdc.MethodBindPtr
	fnSetRootNode                      gdc.MethodBindPtr
	fnGetRootNode                      gdc.MethodBindPtr
	fnSetCallbackModeProcess           gdc.MethodBindPtr
	fnGetCallbackModeProcess           gdc.MethodBindPtr
	fnSetCallbackModeMethod            gdc.MethodBindPtr
	fnGetCallbackModeMethod            gdc.MethodBindPtr
	fnSetCallbackModeDiscrete          gdc.MethodBindPtr
	fnGetCallbackModeDiscrete          gdc.MethodBindPtr
	fnSetAudioMaxPolyphony             gdc.MethodBindPtr
	fnGetAudioMaxPolyphony             gdc.MethodBindPtr
	fnSetRootMotionTrack               gdc.MethodBindPtr
	fnGetRootMotionTrack               gdc.MethodBindPtr
	fnGetRootMotionPosition            gdc.MethodBindPtr
	fnGetRootMotionRotation            gdc.MethodBindPtr
	fnGetRootMotionScale               gdc.MethodBindPtr
	fnGetRootMotionPositionAccumulator gdc.MethodBindPtr
	fnGetRootMotionRotationAccumulator gdc.MethodBindPtr
	fnGetRootMotionScaleAccumulator    gdc.MethodBindPtr
	fnClearCaches                      gdc.MethodBindPtr
	fnAdvance                          gdc.MethodBindPtr
	fnCapture                          gdc.MethodBindPtr
	fnSetResetOnSaveEnabled            gdc.MethodBindPtr
	fnIsResetOnSaveEnabled             gdc.MethodBindPtr
	fnFindAnimation                    gdc.MethodBindPtr
	fnFindAnimationLibrary             gdc.MethodBindPtr
}

var ptrsForAnimationMixer ptrsForAnimationMixerList

func initAnimationMixerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("AnimationMixer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_animation_library")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnAddAnimationLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 618909818))
	}
	{
		methodName := StringNameFromStr("remove_animation_library")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnRemoveAnimationLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("rename_animation_library")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnRenameAnimationLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("has_animation_library")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnHasAnimationLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_animation_library")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetAnimationLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 147342321))
	}
	{
		methodName := StringNameFromStr("get_animation_library_list")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetAnimationLibraryList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("has_animation")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnHasAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("get_animation")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2933122410))
	}
	{
		methodName := StringNameFromStr("get_animation_list")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetAnimationList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1139954409))
	}
	{
		methodName := StringNameFromStr("set_active")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnSetActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_active")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnIsActive = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_deterministic")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnSetDeterministic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_deterministic")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnIsDeterministic = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_root_node")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnSetRootNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_root_node")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetRootNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_callback_mode_process")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnSetCallbackModeProcess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2153733086))
	}
	{
		methodName := StringNameFromStr("get_callback_mode_process")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetCallbackModeProcess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1394468472))
	}
	{
		methodName := StringNameFromStr("set_callback_mode_method")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnSetCallbackModeMethod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 742218271))
	}
	{
		methodName := StringNameFromStr("get_callback_mode_method")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetCallbackModeMethod = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 489449656))
	}
	{
		methodName := StringNameFromStr("set_callback_mode_discrete")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnSetCallbackModeDiscrete = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1998944670))
	}
	{
		methodName := StringNameFromStr("get_callback_mode_discrete")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetCallbackModeDiscrete = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3493168860))
	}
	{
		methodName := StringNameFromStr("set_audio_max_polyphony")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnSetAudioMaxPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_audio_max_polyphony")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetAudioMaxPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_root_motion_track")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnSetRootMotionTrack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_root_motion_track")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetRootMotionTrack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("get_root_motion_position")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetRootMotionPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_root_motion_rotation")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetRootMotionRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1222331677))
	}
	{
		methodName := StringNameFromStr("get_root_motion_scale")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetRootMotionScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_root_motion_position_accumulator")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetRootMotionPositionAccumulator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_root_motion_rotation_accumulator")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetRootMotionRotationAccumulator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1222331677))
	}
	{
		methodName := StringNameFromStr("get_root_motion_scale_accumulator")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnGetRootMotionScaleAccumulator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("clear_caches")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnClearCaches = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("advance")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnAdvance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("capture")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnCapture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1333632127))
	}
	{
		methodName := StringNameFromStr("set_reset_on_save_enabled")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnSetResetOnSaveEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_reset_on_save_enabled")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnIsResetOnSaveEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("find_animation")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnFindAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1559484580))
	}
	{
		methodName := StringNameFromStr("find_animation_library")
		defer methodName.Destroy()
		ptrsForAnimationMixer.fnFindAnimationLibrary = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1559484580))
	}

}

type AnimationMixer struct {
	Node
}

func (me *AnimationMixer) BaseClass() string {
	return "AnimationMixer"
}

func NewAnimationMixer() *AnimationMixer {
	str := StringNameFromStr("AnimationMixer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &AnimationMixer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type AnimationMixerAnimationCallbackModeProcess int

const (
	AnimationMixerAnimationCallbackModeProcessAnimationCallbackModeProcessPhysics AnimationMixerAnimationCallbackModeProcess = 0
	AnimationMixerAnimationCallbackModeProcessAnimationCallbackModeProcessIdle    AnimationMixerAnimationCallbackModeProcess = 1
	AnimationMixerAnimationCallbackModeProcessAnimationCallbackModeProcessManual  AnimationMixerAnimationCallbackModeProcess = 2
)

type AnimationMixerAnimationCallbackModeMethod int

const (
	AnimationMixerAnimationCallbackModeMethodAnimationCallbackModeMethodDeferred  AnimationMixerAnimationCallbackModeMethod = 0
	AnimationMixerAnimationCallbackModeMethodAnimationCallbackModeMethodImmediate AnimationMixerAnimationCallbackModeMethod = 1
)

type AnimationMixerAnimationCallbackModeDiscrete int

const (
	AnimationMixerAnimationCallbackModeDiscreteAnimationCallbackModeDiscreteDominant        AnimationMixerAnimationCallbackModeDiscrete = 0
	AnimationMixerAnimationCallbackModeDiscreteAnimationCallbackModeDiscreteRecessive       AnimationMixerAnimationCallbackModeDiscrete = 1
	AnimationMixerAnimationCallbackModeDiscreteAnimationCallbackModeDiscreteForceContinuous AnimationMixerAnimationCallbackModeDiscrete = 2
)

func (me *AnimationMixer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *AnimationMixer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *AnimationMixer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *AnimationMixer) AddAnimationLibrary(name StringName, library AnimationLibrary) Error {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), library.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnAddAnimationLibrary), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationMixer) RemoveAnimationLibrary(name StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnRemoveAnimationLibrary), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) RenameAnimationLibrary(name StringName, newname StringName) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), newname.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnRenameAnimationLibrary), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) HasAnimationLibrary(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnHasAnimationLibrary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationMixer) GetAnimationLibrary(name StringName) AnimationLibrary {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAnimationLibrary()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetAnimationLibrary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) GetAnimationLibraryList() []StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetAnimationLibraryList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[StringName](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *AnimationMixer) HasAnimation(name StringName) bool {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnHasAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationMixer) GetAnimation(name StringName) Animation {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAnimation()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) GetAnimationList() PackedStringArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedStringArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetAnimationList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) SetActive(active bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnSetActive), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) IsActive() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnIsActive), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationMixer) SetDeterministic(deterministic bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&deterministic)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnSetDeterministic), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) IsDeterministic() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnIsDeterministic), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationMixer) SetRootNode(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnSetRootNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) GetRootNode() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetRootNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) SetCallbackModeProcess(mode AnimationMixerAnimationCallbackModeProcess) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnSetCallbackModeProcess), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) GetCallbackModeProcess() AnimationMixerAnimationCallbackModeProcess {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationMixerAnimationCallbackModeProcess

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetCallbackModeProcess), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationMixer) SetCallbackModeMethod(mode AnimationMixerAnimationCallbackModeMethod) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnSetCallbackModeMethod), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) GetCallbackModeMethod() AnimationMixerAnimationCallbackModeMethod {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationMixerAnimationCallbackModeMethod

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetCallbackModeMethod), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationMixer) SetCallbackModeDiscrete(mode AnimationMixerAnimationCallbackModeDiscrete) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnSetCallbackModeDiscrete), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) GetCallbackModeDiscrete() AnimationMixerAnimationCallbackModeDiscrete {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret AnimationMixerAnimationCallbackModeDiscrete

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetCallbackModeDiscrete), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *AnimationMixer) SetAudioMaxPolyphony(max_polyphony int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_polyphony)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnSetAudioMaxPolyphony), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) GetAudioMaxPolyphony() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetAudioMaxPolyphony), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationMixer) SetRootMotionTrack(path NodePath) {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnSetRootMotionTrack), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) GetRootMotionTrack() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetRootMotionTrack), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) GetRootMotionPosition() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetRootMotionPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) GetRootMotionRotation() Quaternion {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewQuaternion()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetRootMotionRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) GetRootMotionScale() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetRootMotionScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) GetRootMotionPositionAccumulator() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetRootMotionPositionAccumulator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) GetRootMotionRotationAccumulator() Quaternion {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewQuaternion()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetRootMotionRotationAccumulator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) GetRootMotionScaleAccumulator() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnGetRootMotionScaleAccumulator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) ClearCaches() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnClearCaches), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) Advance(delta float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnAdvance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) Capture(name StringName, duration float64, trans_type TweenTransitionType, ease_type TweenEaseType) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&duration), gdc.ConstTypePtr(&trans_type), gdc.ConstTypePtr(&ease_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnCapture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) SetResetOnSaveEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnSetResetOnSaveEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *AnimationMixer) IsResetOnSaveEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnIsResetOnSaveEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *AnimationMixer) FindAnimation(animation Animation) StringName {
	cargs := []gdc.ConstTypePtr{animation.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnFindAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *AnimationMixer) FindAnimationLibrary(animation Animation) StringName {
	cargs := []gdc.ConstTypePtr{animation.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimationMixer.fnFindAnimationLibrary), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationMixerAnimationListChangedSignalFn func()

func (me *AnimationMixer) ConnectAnimationListChanged(subs SignalSubscribers, fn AnimationMixerAnimationListChangedSignalFn) {
	sig := StringNameFromStr("animation_list_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectAnimationListChanged(subs SignalSubscribers, fn AnimationMixerAnimationListChangedSignalFn) {
	sig := StringNameFromStr("animation_list_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerAnimationLibrariesUpdatedSignalFn func()

func (me *AnimationMixer) ConnectAnimationLibrariesUpdated(subs SignalSubscribers, fn AnimationMixerAnimationLibrariesUpdatedSignalFn) {
	sig := StringNameFromStr("animation_libraries_updated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectAnimationLibrariesUpdated(subs SignalSubscribers, fn AnimationMixerAnimationLibrariesUpdatedSignalFn) {
	sig := StringNameFromStr("animation_libraries_updated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerAnimationFinishedSignalFn func(anim_name StringName)

func (me *AnimationMixer) ConnectAnimationFinished(subs SignalSubscribers, fn AnimationMixerAnimationFinishedSignalFn) {
	sig := StringNameFromStr("animation_finished")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectAnimationFinished(subs SignalSubscribers, fn AnimationMixerAnimationFinishedSignalFn) {
	sig := StringNameFromStr("animation_finished")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerAnimationStartedSignalFn func(anim_name StringName)

func (me *AnimationMixer) ConnectAnimationStarted(subs SignalSubscribers, fn AnimationMixerAnimationStartedSignalFn) {
	sig := StringNameFromStr("animation_started")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectAnimationStarted(subs SignalSubscribers, fn AnimationMixerAnimationStartedSignalFn) {
	sig := StringNameFromStr("animation_started")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerCachesClearedSignalFn func()

func (me *AnimationMixer) ConnectCachesCleared(subs SignalSubscribers, fn AnimationMixerCachesClearedSignalFn) {
	sig := StringNameFromStr("caches_cleared")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectCachesCleared(subs SignalSubscribers, fn AnimationMixerCachesClearedSignalFn) {
	sig := StringNameFromStr("caches_cleared")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerMixerAppliedSignalFn func()

func (me *AnimationMixer) ConnectMixerApplied(subs SignalSubscribers, fn AnimationMixerMixerAppliedSignalFn) {
	sig := StringNameFromStr("mixer_applied")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectMixerApplied(subs SignalSubscribers, fn AnimationMixerMixerAppliedSignalFn) {
	sig := StringNameFromStr("mixer_applied")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type AnimationMixerMixerUpdatedSignalFn func()

func (me *AnimationMixer) ConnectMixerUpdated(subs SignalSubscribers, fn AnimationMixerMixerUpdatedSignalFn) {
	sig := StringNameFromStr("mixer_updated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationMixer) DisconnectMixerUpdated(subs SignalSubscribers, fn AnimationMixerMixerUpdatedSignalFn) {
	sig := StringNameFromStr("mixer_updated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
