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

type ptrsForAnimationList struct {
  fnAddTrack gdc.MethodBindPtr
  fnRemoveTrack gdc.MethodBindPtr
  fnGetTrackCount gdc.MethodBindPtr
  fnTrackGetType gdc.MethodBindPtr
  fnTrackGetPath gdc.MethodBindPtr
  fnTrackSetPath gdc.MethodBindPtr
  fnFindTrack gdc.MethodBindPtr
  fnTrackMoveUp gdc.MethodBindPtr
  fnTrackMoveDown gdc.MethodBindPtr
  fnTrackMoveTo gdc.MethodBindPtr
  fnTrackSwap gdc.MethodBindPtr
  fnTrackSetImported gdc.MethodBindPtr
  fnTrackIsImported gdc.MethodBindPtr
  fnTrackSetEnabled gdc.MethodBindPtr
  fnTrackIsEnabled gdc.MethodBindPtr
  fnPositionTrackInsertKey gdc.MethodBindPtr
  fnRotationTrackInsertKey gdc.MethodBindPtr
  fnScaleTrackInsertKey gdc.MethodBindPtr
  fnBlendShapeTrackInsertKey gdc.MethodBindPtr
  fnPositionTrackInterpolate gdc.MethodBindPtr
  fnRotationTrackInterpolate gdc.MethodBindPtr
  fnScaleTrackInterpolate gdc.MethodBindPtr
  fnBlendShapeTrackInterpolate gdc.MethodBindPtr
  fnTrackInsertKey gdc.MethodBindPtr
  fnTrackRemoveKey gdc.MethodBindPtr
  fnTrackRemoveKeyAtTime gdc.MethodBindPtr
  fnTrackSetKeyValue gdc.MethodBindPtr
  fnTrackSetKeyTransition gdc.MethodBindPtr
  fnTrackSetKeyTime gdc.MethodBindPtr
  fnTrackGetKeyTransition gdc.MethodBindPtr
  fnTrackGetKeyCount gdc.MethodBindPtr
  fnTrackGetKeyValue gdc.MethodBindPtr
  fnTrackGetKeyTime gdc.MethodBindPtr
  fnTrackFindKey gdc.MethodBindPtr
  fnTrackSetInterpolationType gdc.MethodBindPtr
  fnTrackGetInterpolationType gdc.MethodBindPtr
  fnTrackSetInterpolationLoopWrap gdc.MethodBindPtr
  fnTrackGetInterpolationLoopWrap gdc.MethodBindPtr
  fnTrackIsCompressed gdc.MethodBindPtr
  fnValueTrackSetUpdateMode gdc.MethodBindPtr
  fnValueTrackGetUpdateMode gdc.MethodBindPtr
  fnValueTrackInterpolate gdc.MethodBindPtr
  fnMethodTrackGetName gdc.MethodBindPtr
  fnMethodTrackGetParams gdc.MethodBindPtr
  fnBezierTrackInsertKey gdc.MethodBindPtr
  fnBezierTrackSetKeyValue gdc.MethodBindPtr
  fnBezierTrackSetKeyInHandle gdc.MethodBindPtr
  fnBezierTrackSetKeyOutHandle gdc.MethodBindPtr
  fnBezierTrackGetKeyValue gdc.MethodBindPtr
  fnBezierTrackGetKeyInHandle gdc.MethodBindPtr
  fnBezierTrackGetKeyOutHandle gdc.MethodBindPtr
  fnBezierTrackInterpolate gdc.MethodBindPtr
  fnAudioTrackInsertKey gdc.MethodBindPtr
  fnAudioTrackSetKeyStream gdc.MethodBindPtr
  fnAudioTrackSetKeyStartOffset gdc.MethodBindPtr
  fnAudioTrackSetKeyEndOffset gdc.MethodBindPtr
  fnAudioTrackGetKeyStream gdc.MethodBindPtr
  fnAudioTrackGetKeyStartOffset gdc.MethodBindPtr
  fnAudioTrackGetKeyEndOffset gdc.MethodBindPtr
  fnAudioTrackSetUseBlend gdc.MethodBindPtr
  fnAudioTrackIsUseBlend gdc.MethodBindPtr
  fnAnimationTrackInsertKey gdc.MethodBindPtr
  fnAnimationTrackSetKeyAnimation gdc.MethodBindPtr
  fnAnimationTrackGetKeyAnimation gdc.MethodBindPtr
  fnSetLength gdc.MethodBindPtr
  fnGetLength gdc.MethodBindPtr
  fnSetLoopMode gdc.MethodBindPtr
  fnGetLoopMode gdc.MethodBindPtr
  fnSetStep gdc.MethodBindPtr
  fnGetStep gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
  fnCopyTrack gdc.MethodBindPtr
  fnCompress gdc.MethodBindPtr
}

var ptrsForAnimation ptrsForAnimationList

func initAnimationPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Animation")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_track")
    defer methodName.Destroy()
    ptrsForAnimation.fnAddTrack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3843682357))
  }
  {
    methodName := StringNameFromStr("remove_track")
    defer methodName.Destroy()
    ptrsForAnimation.fnRemoveTrack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_track_count")
    defer methodName.Destroy()
    ptrsForAnimation.fnGetTrackCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("track_get_type")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackGetType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3445944217))
  }
  {
    methodName := StringNameFromStr("track_get_path")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
  }
  {
    methodName := StringNameFromStr("track_set_path")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackSetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761262315))
  }
  {
    methodName := StringNameFromStr("find_track")
    defer methodName.Destroy()
    ptrsForAnimation.fnFindTrack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 245376003))
  }
  {
    methodName := StringNameFromStr("track_move_up")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackMoveUp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("track_move_down")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackMoveDown = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("track_move_to")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackMoveTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("track_swap")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackSwap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("track_set_imported")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackSetImported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("track_is_imported")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackIsImported = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("track_set_enabled")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("track_is_enabled")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackIsEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("position_track_insert_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnPositionTrackInsertKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2540608232))
  }
  {
    methodName := StringNameFromStr("rotation_track_insert_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnRotationTrackInsertKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4165004800))
  }
  {
    methodName := StringNameFromStr("scale_track_insert_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnScaleTrackInsertKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2540608232))
  }
  {
    methodName := StringNameFromStr("blend_shape_track_insert_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnBlendShapeTrackInsertKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1534913637))
  }
  {
    methodName := StringNameFromStr("position_track_interpolate")
    defer methodName.Destroy()
    ptrsForAnimation.fnPositionTrackInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3285246857))
  }
  {
    methodName := StringNameFromStr("rotation_track_interpolate")
    defer methodName.Destroy()
    ptrsForAnimation.fnRotationTrackInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1988711975))
  }
  {
    methodName := StringNameFromStr("scale_track_interpolate")
    defer methodName.Destroy()
    ptrsForAnimation.fnScaleTrackInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3285246857))
  }
  {
    methodName := StringNameFromStr("blend_shape_track_interpolate")
    defer methodName.Destroy()
    ptrsForAnimation.fnBlendShapeTrackInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1900462983))
  }
  {
    methodName := StringNameFromStr("track_insert_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackInsertKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 808952278))
  }
  {
    methodName := StringNameFromStr("track_remove_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackRemoveKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("track_remove_key_at_time")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackRemoveKeyAtTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("track_set_key_value")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackSetKeyValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2060538656))
  }
  {
    methodName := StringNameFromStr("track_set_key_transition")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackSetKeyTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
  }
  {
    methodName := StringNameFromStr("track_set_key_time")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackSetKeyTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
  }
  {
    methodName := StringNameFromStr("track_get_key_transition")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackGetKeyTransition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
  }
  {
    methodName := StringNameFromStr("track_get_key_count")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackGetKeyCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("track_get_key_value")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackGetKeyValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 678354945))
  }
  {
    methodName := StringNameFromStr("track_get_key_time")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackGetKeyTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
  }
  {
    methodName := StringNameFromStr("track_find_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackFindKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3245197284))
  }
  {
    methodName := StringNameFromStr("track_set_interpolation_type")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackSetInterpolationType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4112932513))
  }
  {
    methodName := StringNameFromStr("track_get_interpolation_type")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackGetInterpolationType = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530756894))
  }
  {
    methodName := StringNameFromStr("track_set_interpolation_loop_wrap")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackSetInterpolationLoopWrap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("track_get_interpolation_loop_wrap")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackGetInterpolationLoopWrap = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("track_is_compressed")
    defer methodName.Destroy()
    ptrsForAnimation.fnTrackIsCompressed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("value_track_set_update_mode")
    defer methodName.Destroy()
    ptrsForAnimation.fnValueTrackSetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2854058312))
  }
  {
    methodName := StringNameFromStr("value_track_get_update_mode")
    defer methodName.Destroy()
    ptrsForAnimation.fnValueTrackGetUpdateMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1440326473))
  }
  {
    methodName := StringNameFromStr("value_track_interpolate")
    defer methodName.Destroy()
    ptrsForAnimation.fnValueTrackInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 491147702))
  }
  {
    methodName := StringNameFromStr("method_track_get_name")
    defer methodName.Destroy()
    ptrsForAnimation.fnMethodTrackGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 351665558))
  }
  {
    methodName := StringNameFromStr("method_track_get_params")
    defer methodName.Destroy()
    ptrsForAnimation.fnMethodTrackGetParams = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2345056839))
  }
  {
    methodName := StringNameFromStr("bezier_track_insert_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnBezierTrackInsertKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3656773645))
  }
  {
    methodName := StringNameFromStr("bezier_track_set_key_value")
    defer methodName.Destroy()
    ptrsForAnimation.fnBezierTrackSetKeyValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
  }
  {
    methodName := StringNameFromStr("bezier_track_set_key_in_handle")
    defer methodName.Destroy()
    ptrsForAnimation.fnBezierTrackSetKeyInHandle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1719223284))
  }
  {
    methodName := StringNameFromStr("bezier_track_set_key_out_handle")
    defer methodName.Destroy()
    ptrsForAnimation.fnBezierTrackSetKeyOutHandle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1719223284))
  }
  {
    methodName := StringNameFromStr("bezier_track_get_key_value")
    defer methodName.Destroy()
    ptrsForAnimation.fnBezierTrackGetKeyValue = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
  }
  {
    methodName := StringNameFromStr("bezier_track_get_key_in_handle")
    defer methodName.Destroy()
    ptrsForAnimation.fnBezierTrackGetKeyInHandle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3016396712))
  }
  {
    methodName := StringNameFromStr("bezier_track_get_key_out_handle")
    defer methodName.Destroy()
    ptrsForAnimation.fnBezierTrackGetKeyOutHandle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3016396712))
  }
  {
    methodName := StringNameFromStr("bezier_track_interpolate")
    defer methodName.Destroy()
    ptrsForAnimation.fnBezierTrackInterpolate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1900462983))
  }
  {
    methodName := StringNameFromStr("audio_track_insert_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnAudioTrackInsertKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4021027286))
  }
  {
    methodName := StringNameFromStr("audio_track_set_key_stream")
    defer methodName.Destroy()
    ptrsForAnimation.fnAudioTrackSetKeyStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3886397084))
  }
  {
    methodName := StringNameFromStr("audio_track_set_key_start_offset")
    defer methodName.Destroy()
    ptrsForAnimation.fnAudioTrackSetKeyStartOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
  }
  {
    methodName := StringNameFromStr("audio_track_set_key_end_offset")
    defer methodName.Destroy()
    ptrsForAnimation.fnAudioTrackSetKeyEndOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3506521499))
  }
  {
    methodName := StringNameFromStr("audio_track_get_key_stream")
    defer methodName.Destroy()
    ptrsForAnimation.fnAudioTrackGetKeyStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 635277205))
  }
  {
    methodName := StringNameFromStr("audio_track_get_key_start_offset")
    defer methodName.Destroy()
    ptrsForAnimation.fnAudioTrackGetKeyStartOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
  }
  {
    methodName := StringNameFromStr("audio_track_get_key_end_offset")
    defer methodName.Destroy()
    ptrsForAnimation.fnAudioTrackGetKeyEndOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
  }
  {
    methodName := StringNameFromStr("audio_track_set_use_blend")
    defer methodName.Destroy()
    ptrsForAnimation.fnAudioTrackSetUseBlend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("audio_track_is_use_blend")
    defer methodName.Destroy()
    ptrsForAnimation.fnAudioTrackIsUseBlend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("animation_track_insert_key")
    defer methodName.Destroy()
    ptrsForAnimation.fnAnimationTrackInsertKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 158676774))
  }
  {
    methodName := StringNameFromStr("animation_track_set_key_animation")
    defer methodName.Destroy()
    ptrsForAnimation.fnAnimationTrackSetKeyAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 117615382))
  }
  {
    methodName := StringNameFromStr("animation_track_get_key_animation")
    defer methodName.Destroy()
    ptrsForAnimation.fnAnimationTrackGetKeyAnimation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 351665558))
  }
  {
    methodName := StringNameFromStr("set_length")
    defer methodName.Destroy()
    ptrsForAnimation.fnSetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_length")
    defer methodName.Destroy()
    ptrsForAnimation.fnGetLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_loop_mode")
    defer methodName.Destroy()
    ptrsForAnimation.fnSetLoopMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3155355575))
  }
  {
    methodName := StringNameFromStr("get_loop_mode")
    defer methodName.Destroy()
    ptrsForAnimation.fnGetLoopMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1988889481))
  }
  {
    methodName := StringNameFromStr("set_step")
    defer methodName.Destroy()
    ptrsForAnimation.fnSetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_step")
    defer methodName.Destroy()
    ptrsForAnimation.fnGetStep = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForAnimation.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("copy_track")
    defer methodName.Destroy()
    ptrsForAnimation.fnCopyTrack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 148001024))
  }
  {
    methodName := StringNameFromStr("compress")
    defer methodName.Destroy()
    ptrsForAnimation.fnCompress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3608408117))
  }
}

type Animation struct {
  Resource
}

func (me *Animation) BaseClass() string {
  return "Animation"
}

func NewAnimation() *Animation {
  str := StringNameFromStr("Animation") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Animation{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AnimationTrackType int
const (
  AnimationTrackTypeTypeValue AnimationTrackType = 0
  AnimationTrackTypeTypePosition3D AnimationTrackType = 1
  AnimationTrackTypeTypeRotation3D AnimationTrackType = 2
  AnimationTrackTypeTypeScale3D AnimationTrackType = 3
  AnimationTrackTypeTypeBlendShape AnimationTrackType = 4
  AnimationTrackTypeTypeMethod AnimationTrackType = 5
  AnimationTrackTypeTypeBezier AnimationTrackType = 6
  AnimationTrackTypeTypeAudio AnimationTrackType = 7
  AnimationTrackTypeTypeAnimation AnimationTrackType = 8
)

type AnimationInterpolationType int
const (
  AnimationInterpolationTypeInterpolationNearest AnimationInterpolationType = 0
  AnimationInterpolationTypeInterpolationLinear AnimationInterpolationType = 1
  AnimationInterpolationTypeInterpolationCubic AnimationInterpolationType = 2
  AnimationInterpolationTypeInterpolationLinearAngle AnimationInterpolationType = 3
  AnimationInterpolationTypeInterpolationCubicAngle AnimationInterpolationType = 4
)

type AnimationUpdateMode int
const (
  AnimationUpdateModeUpdateContinuous AnimationUpdateMode = 0
  AnimationUpdateModeUpdateDiscrete AnimationUpdateMode = 1
  AnimationUpdateModeUpdateCapture AnimationUpdateMode = 2
)

type AnimationLoopMode int
const (
  AnimationLoopModeLoopNone AnimationLoopMode = 0
  AnimationLoopModeLoopLinear AnimationLoopMode = 1
  AnimationLoopModeLoopPingpong AnimationLoopMode = 2
)

type AnimationLoopedFlag int
const (
  AnimationLoopedFlagLoopedFlagNone AnimationLoopedFlag = 0
  AnimationLoopedFlagLoopedFlagEnd AnimationLoopedFlag = 1
  AnimationLoopedFlagLoopedFlagStart AnimationLoopedFlag = 2
)

type AnimationFindMode int
const (
  AnimationFindModeFindModeNearest AnimationFindMode = 0
  AnimationFindModeFindModeApprox AnimationFindMode = 1
  AnimationFindModeFindModeExact AnimationFindMode = 2
)

func (me *Animation) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Animation) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Animation) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Animation) AddTrack(type_ AnimationTrackType, at_position int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , gdc.ConstTypePtr(&at_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&type_)
  pinner.Pin(&at_position)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAddTrack), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) RemoveTrack(track_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnRemoveTrack), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) GetTrackCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnGetTrackCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackGetType(track_idx int64, ) AnimationTrackType {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationTrackType
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackGetType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Animation) TrackGetPath(track_idx int64, ) NodePath {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) TrackSetPath(track_idx int64, path NodePath, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackSetPath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) FindTrack(path NodePath, type_ AnimationTrackType, ) int64 {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&type_)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnFindTrack), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackMoveUp(track_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackMoveUp), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackMoveDown(track_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackMoveDown), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackMoveTo(track_idx int64, to_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&to_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackMoveTo), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackSwap(track_idx int64, with_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&with_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackSwap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackSetImported(track_idx int64, imported bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&imported) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackSetImported), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackIsImported(track_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackIsImported), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackSetEnabled(track_idx int64, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackIsEnabled(track_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackIsEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) PositionTrackInsertKey(track_idx int64, time float64, position Vector3, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnPositionTrackInsertKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) RotationTrackInsertKey(track_idx int64, time float64, rotation Quaternion, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , rotation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnRotationTrackInsertKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) ScaleTrackInsertKey(track_idx int64, time float64, scale Vector3, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnScaleTrackInsertKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) BlendShapeTrackInsertKey(track_idx int64, time float64, amount float64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)
  pinner.Pin(&amount)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBlendShapeTrackInsertKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) PositionTrackInterpolate(track_idx int64, time_sec float64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&track_idx)
  pinner.Pin(&time_sec)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnPositionTrackInterpolate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) RotationTrackInterpolate(track_idx int64, time_sec float64, ) Quaternion {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewQuaternion()
  pinner.Pin(&track_idx)
  pinner.Pin(&time_sec)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnRotationTrackInterpolate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) ScaleTrackInterpolate(track_idx int64, time_sec float64, ) Vector3 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&track_idx)
  pinner.Pin(&time_sec)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnScaleTrackInterpolate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) BlendShapeTrackInterpolate(track_idx int64, time_sec float64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&track_idx)
  pinner.Pin(&time_sec)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBlendShapeTrackInterpolate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackInsertKey(track_idx int64, time float64, key Variant, transition float64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , key.AsCTypePtr(), gdc.ConstTypePtr(&transition) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)
  pinner.Pin(&transition)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackInsertKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackRemoveKey(track_idx int64, key_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackRemoveKey), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackRemoveKeyAtTime(track_idx int64, time float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackRemoveKeyAtTime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackSetKeyValue(track_idx int64, key int64, value Variant, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key) , value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackSetKeyValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackSetKeyTransition(track_idx int64, key_idx int64, transition float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , gdc.ConstTypePtr(&transition) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackSetKeyTransition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackSetKeyTime(track_idx int64, key_idx int64, time float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackSetKeyTime), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackGetKeyTransition(track_idx int64, key_idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackGetKeyTransition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackGetKeyCount(track_idx int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackGetKeyCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackGetKeyValue(track_idx int64, key_idx int64, ) Variant {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackGetKeyValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) TrackGetKeyTime(track_idx int64, key_idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackGetKeyTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackFindKey(track_idx int64, time float64, find_mode AnimationFindMode, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , gdc.ConstTypePtr(&find_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)
  pinner.Pin(&find_mode)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackFindKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackSetInterpolationType(track_idx int64, interpolation AnimationInterpolationType, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&interpolation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackSetInterpolationType), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackGetInterpolationType(track_idx int64, ) AnimationInterpolationType {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationInterpolationType
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackGetInterpolationType), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Animation) TrackSetInterpolationLoopWrap(track_idx int64, interpolation bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&interpolation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackSetInterpolationLoopWrap), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) TrackGetInterpolationLoopWrap(track_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackGetInterpolationLoopWrap), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) TrackIsCompressed(track_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnTrackIsCompressed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) ValueTrackSetUpdateMode(track_idx int64, mode AnimationUpdateMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnValueTrackSetUpdateMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) ValueTrackGetUpdateMode(track_idx int64, ) AnimationUpdateMode {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationUpdateMode
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnValueTrackGetUpdateMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Animation) ValueTrackInterpolate(track_idx int64, time_sec float64, ) Variant {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&track_idx)
  pinner.Pin(&time_sec)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnValueTrackInterpolate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) MethodTrackGetName(track_idx int64, key_idx int64, ) StringName {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnMethodTrackGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) MethodTrackGetParams(track_idx int64, key_idx int64, ) Array {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnMethodTrackGetParams), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) BezierTrackInsertKey(track_idx int64, time float64, value float64, in_handle Vector2, out_handle Vector2, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , gdc.ConstTypePtr(&value) , in_handle.AsCTypePtr(), out_handle.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)
  pinner.Pin(&value)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBezierTrackInsertKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) BezierTrackSetKeyValue(track_idx int64, key_idx int64, value float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , gdc.ConstTypePtr(&value) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBezierTrackSetKeyValue), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) BezierTrackSetKeyInHandle(track_idx int64, key_idx int64, in_handle Vector2, balanced_value_time_ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , in_handle.AsCTypePtr(), gdc.ConstTypePtr(&balanced_value_time_ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBezierTrackSetKeyInHandle), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) BezierTrackSetKeyOutHandle(track_idx int64, key_idx int64, out_handle Vector2, balanced_value_time_ratio float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , out_handle.AsCTypePtr(), gdc.ConstTypePtr(&balanced_value_time_ratio) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBezierTrackSetKeyOutHandle), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) BezierTrackGetKeyValue(track_idx int64, key_idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBezierTrackGetKeyValue), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) BezierTrackGetKeyInHandle(track_idx int64, key_idx int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBezierTrackGetKeyInHandle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) BezierTrackGetKeyOutHandle(track_idx int64, key_idx int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBezierTrackGetKeyOutHandle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) BezierTrackInterpolate(track_idx int64, time float64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnBezierTrackInterpolate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) AudioTrackInsertKey(track_idx int64, time float64, stream Resource, start_offset float64, end_offset float64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , stream.AsCTypePtr(), gdc.ConstTypePtr(&start_offset) , gdc.ConstTypePtr(&end_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)
  pinner.Pin(&start_offset)
  pinner.Pin(&end_offset)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAudioTrackInsertKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) AudioTrackSetKeyStream(track_idx int64, key_idx int64, stream Resource, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , stream.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAudioTrackSetKeyStream), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) AudioTrackSetKeyStartOffset(track_idx int64, key_idx int64, offset float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAudioTrackSetKeyStartOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) AudioTrackSetKeyEndOffset(track_idx int64, key_idx int64, offset float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAudioTrackSetKeyEndOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) AudioTrackGetKeyStream(track_idx int64, key_idx int64, ) Resource {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewResource()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAudioTrackGetKeyStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) AudioTrackGetKeyStartOffset(track_idx int64, key_idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAudioTrackGetKeyStartOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) AudioTrackGetKeyEndOffset(track_idx int64, key_idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAudioTrackGetKeyEndOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) AudioTrackSetUseBlend(track_idx int64, enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAudioTrackSetUseBlend), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) AudioTrackIsUseBlend(track_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&track_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAudioTrackIsUseBlend), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) AnimationTrackInsertKey(track_idx int64, time float64, animation StringName, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&time) , animation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&track_idx)
  pinner.Pin(&time)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAnimationTrackInsertKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) AnimationTrackSetKeyAnimation(track_idx int64, key_idx int64, animation StringName, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , animation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAnimationTrackSetKeyAnimation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) AnimationTrackGetKeyAnimation(track_idx int64, key_idx int64, ) StringName {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , gdc.ConstTypePtr(&key_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&track_idx)
  pinner.Pin(&key_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnAnimationTrackGetKeyAnimation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Animation) SetLength(time_sec float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnSetLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) GetLength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnGetLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) SetLoopMode(loop_mode AnimationLoopMode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnSetLoopMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) GetLoopMode() AnimationLoopMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationLoopMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnGetLoopMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Animation) SetStep(size_sec float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_sec) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnSetStep), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) GetStep() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnGetStep), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Animation) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) CopyTrack(track_idx int64, to_animation Animation, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx) , to_animation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnCopyTrack), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Animation) Compress(page_size int64, fps int64, split_tolerance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&page_size) , gdc.ConstTypePtr(&fps) , gdc.ConstTypePtr(&split_tolerance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAnimation.fnCompress), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
