// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Animation struct {
  obj gdc.ObjectPtr
}

func (me *Animation) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Animation) BaseClass() string {
  return "Animation"
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

func  (me *Animation) AddTrack(type_ AnimationTrackType, at_position int, )  {
  panic("TODO: implement")
}

func  (me *Animation) RemoveTrack(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) GetTrackCount()  {
  panic("TODO: implement")
}

func  (me *Animation) TrackGetType(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackGetPath(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackSetPath(track_idx int, path NodePath, )  {
  panic("TODO: implement")
}

func  (me *Animation) FindTrack(path NodePath, type_ AnimationTrackType, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackMoveUp(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackMoveDown(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackMoveTo(track_idx int, to_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackSwap(track_idx int, with_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackSetImported(track_idx int, imported bool, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackIsImported(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackSetEnabled(track_idx int, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackIsEnabled(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) PositionTrackInsertKey(track_idx int, time float32, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *Animation) RotationTrackInsertKey(track_idx int, time float32, rotation Quaternion, )  {
  panic("TODO: implement")
}

func  (me *Animation) ScaleTrackInsertKey(track_idx int, time float32, scale Vector3, )  {
  panic("TODO: implement")
}

func  (me *Animation) BlendShapeTrackInsertKey(track_idx int, time float32, amount float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) PositionTrackInterpolate(track_idx int, time_sec float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) RotationTrackInterpolate(track_idx int, time_sec float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) ScaleTrackInterpolate(track_idx int, time_sec float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) BlendShapeTrackInterpolate(track_idx int, time_sec float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackInsertKey(track_idx int, time float32, key Variant, transition float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackRemoveKey(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackRemoveKeyAtTime(track_idx int, time float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackSetKeyValue(track_idx int, key int, value Variant, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackSetKeyTransition(track_idx int, key_idx int, transition float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackSetKeyTime(track_idx int, key_idx int, time float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackGetKeyTransition(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackGetKeyCount(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackGetKeyValue(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackGetKeyTime(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackFindKey(track_idx int, time float32, find_mode AnimationFindMode, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackSetInterpolationType(track_idx int, interpolation AnimationInterpolationType, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackGetInterpolationType(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackSetInterpolationLoopWrap(track_idx int, interpolation bool, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackGetInterpolationLoopWrap(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) TrackIsCompressed(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) ValueTrackSetUpdateMode(track_idx int, mode AnimationUpdateMode, )  {
  panic("TODO: implement")
}

func  (me *Animation) ValueTrackGetUpdateMode(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) ValueTrackInterpolate(track_idx int, time_sec float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) MethodTrackGetName(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) MethodTrackGetParams(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) BezierTrackInsertKey(track_idx int, time float32, value float32, in_handle Vector2, out_handle Vector2, )  {
  panic("TODO: implement")
}

func  (me *Animation) BezierTrackSetKeyValue(track_idx int, key_idx int, value float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) BezierTrackSetKeyInHandle(track_idx int, key_idx int, in_handle Vector2, balanced_value_time_ratio float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) BezierTrackSetKeyOutHandle(track_idx int, key_idx int, out_handle Vector2, balanced_value_time_ratio float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) BezierTrackGetKeyValue(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) BezierTrackGetKeyInHandle(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) BezierTrackGetKeyOutHandle(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) BezierTrackInterpolate(track_idx int, time float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) AudioTrackInsertKey(track_idx int, time float32, stream Resource, start_offset float32, end_offset float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) AudioTrackSetKeyStream(track_idx int, key_idx int, stream Resource, )  {
  panic("TODO: implement")
}

func  (me *Animation) AudioTrackSetKeyStartOffset(track_idx int, key_idx int, offset float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) AudioTrackSetKeyEndOffset(track_idx int, key_idx int, offset float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) AudioTrackGetKeyStream(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) AudioTrackGetKeyStartOffset(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) AudioTrackGetKeyEndOffset(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) AudioTrackSetUseBlend(track_idx int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *Animation) AudioTrackIsUseBlend(track_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) AnimationTrackInsertKey(track_idx int, time float32, animation StringName, )  {
  panic("TODO: implement")
}

func  (me *Animation) AnimationTrackSetKeyAnimation(track_idx int, key_idx int, animation StringName, )  {
  panic("TODO: implement")
}

func  (me *Animation) AnimationTrackGetKeyAnimation(track_idx int, key_idx int, )  {
  panic("TODO: implement")
}

func  (me *Animation) SetLength(time_sec float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) GetLength()  {
  panic("TODO: implement")
}

func  (me *Animation) SetLoopMode(loop_mode AnimationLoopMode, )  {
  panic("TODO: implement")
}

func  (me *Animation) GetLoopMode()  {
  panic("TODO: implement")
}

func  (me *Animation) SetStep(size_sec float32, )  {
  panic("TODO: implement")
}

func  (me *Animation) GetStep()  {
  panic("TODO: implement")
}

func  (me *Animation) Clear()  {
  panic("TODO: implement")
}

func  (me *Animation) CopyTrack(track_idx int, to_animation Animation, )  {
  panic("TODO: implement")
}

func  (me *Animation) Compress(page_size int, fps int, split_tolerance float32, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
