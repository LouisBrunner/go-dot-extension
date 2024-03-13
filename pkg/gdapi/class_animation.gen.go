// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *Animation) AddTrack(type_ AnimationTrackType, at_position int, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3843682357) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&at_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) RemoveTrack(track_idx int, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) GetTrackCount() int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_track_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackGetType(track_idx int, ) AnimationTrackType {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_get_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3445944217) // FIXME: should cache?
  var ret AnimationTrackType
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackGetPath(track_idx int, ) NodePath {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_get_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackSetPath(track_idx int, path NodePath, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_set_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761262315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) FindTrack(path NodePath, type_ AnimationTrackType, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 245376003) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackMoveUp(track_idx int, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_move_up")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackMoveDown(track_idx int, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_move_down")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackMoveTo(track_idx int, to_idx int, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_move_to")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&to_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackSwap(track_idx int, with_idx int, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_swap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&with_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackSetImported(track_idx int, imported bool, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_set_imported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&imported), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackIsImported(track_idx int, ) bool {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_is_imported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackSetEnabled(track_idx int, enabled bool, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackIsEnabled(track_idx int, ) bool {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_is_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) PositionTrackInsertKey(track_idx int, time float32, position Vector3, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("position_track_insert_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2540608232) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) RotationTrackInsertKey(track_idx int, time float32, rotation Quaternion, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotation_track_insert_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4165004800) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(rotation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) ScaleTrackInsertKey(track_idx int, time float32, scale Vector3, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scale_track_insert_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2540608232) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(scale.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) BlendShapeTrackInsertKey(track_idx int, time float32, amount float32, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_shape_track_insert_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1534913637) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) PositionTrackInterpolate(track_idx int, time_sec float32, ) Vector3 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("position_track_interpolate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3285246857) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) RotationTrackInterpolate(track_idx int, time_sec float32, ) Quaternion {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotation_track_interpolate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1988711975) // FIXME: should cache?
  var ret Quaternion
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) ScaleTrackInterpolate(track_idx int, time_sec float32, ) Vector3 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scale_track_interpolate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3285246857) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) BlendShapeTrackInterpolate(track_idx int, time_sec float32, ) float32 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("blend_shape_track_interpolate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1900462983) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackInsertKey(track_idx int, time float32, key Variant, transition float32, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_insert_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 808952278) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(&transition), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackRemoveKey(track_idx int, key_idx int, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_remove_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackRemoveKeyAtTime(track_idx int, time float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_remove_key_at_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackSetKeyValue(track_idx int, key int, value Variant, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_set_key_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2060538656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackSetKeyTransition(track_idx int, key_idx int, transition float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_set_key_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), gdc.ConstTypePtr(&transition), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackSetKeyTime(track_idx int, key_idx int, time float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_set_key_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackGetKeyTransition(track_idx int, key_idx int, ) float32 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_get_key_transition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackGetKeyCount(track_idx int, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_get_key_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackGetKeyValue(track_idx int, key_idx int, ) Variant {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_get_key_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 678354945) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackGetKeyTime(track_idx int, key_idx int, ) float32 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_get_key_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackFindKey(track_idx int, time float32, find_mode AnimationFindMode, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_find_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3245197284) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(&find_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackSetInterpolationType(track_idx int, interpolation AnimationInterpolationType, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_set_interpolation_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4112932513) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&interpolation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackGetInterpolationType(track_idx int, ) AnimationInterpolationType {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_get_interpolation_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530756894) // FIXME: should cache?
  var ret AnimationInterpolationType
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackSetInterpolationLoopWrap(track_idx int, interpolation bool, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_set_interpolation_loop_wrap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&interpolation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) TrackGetInterpolationLoopWrap(track_idx int, ) bool {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_get_interpolation_loop_wrap")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) TrackIsCompressed(track_idx int, ) bool {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("track_is_compressed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) ValueTrackSetUpdateMode(track_idx int, mode AnimationUpdateMode, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("value_track_set_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2854058312) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) ValueTrackGetUpdateMode(track_idx int, ) AnimationUpdateMode {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("value_track_get_update_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1440326473) // FIXME: should cache?
  var ret AnimationUpdateMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) ValueTrackInterpolate(track_idx int, time_sec float32, ) Variant {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("value_track_interpolate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 491147702) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) MethodTrackGetName(track_idx int, key_idx int, ) StringName {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("method_track_get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 351665558) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) MethodTrackGetParams(track_idx int, key_idx int, ) Array {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("method_track_get_params")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2345056839) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) BezierTrackInsertKey(track_idx int, time float32, value float32, in_handle Vector2, out_handle Vector2, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bezier_track_insert_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3656773645) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(&value), gdc.ConstTypePtr(in_handle.AsCTypePtr()), gdc.ConstTypePtr(out_handle.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) BezierTrackSetKeyValue(track_idx int, key_idx int, value float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bezier_track_set_key_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) BezierTrackSetKeyInHandle(track_idx int, key_idx int, in_handle Vector2, balanced_value_time_ratio float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bezier_track_set_key_in_handle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1719223284) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), gdc.ConstTypePtr(in_handle.AsCTypePtr()), gdc.ConstTypePtr(&balanced_value_time_ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) BezierTrackSetKeyOutHandle(track_idx int, key_idx int, out_handle Vector2, balanced_value_time_ratio float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bezier_track_set_key_out_handle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1719223284) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), gdc.ConstTypePtr(out_handle.AsCTypePtr()), gdc.ConstTypePtr(&balanced_value_time_ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) BezierTrackGetKeyValue(track_idx int, key_idx int, ) float32 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bezier_track_get_key_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) BezierTrackGetKeyInHandle(track_idx int, key_idx int, ) Vector2 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bezier_track_get_key_in_handle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3016396712) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) BezierTrackGetKeyOutHandle(track_idx int, key_idx int, ) Vector2 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bezier_track_get_key_out_handle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3016396712) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) BezierTrackInterpolate(track_idx int, time float32, ) float32 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("bezier_track_interpolate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1900462983) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) AudioTrackInsertKey(track_idx int, time float32, stream Resource, start_offset float32, end_offset float32, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("audio_track_insert_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4021027286) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(stream.AsCTypePtr()), gdc.ConstTypePtr(&start_offset), gdc.ConstTypePtr(&end_offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) AudioTrackSetKeyStream(track_idx int, key_idx int, stream Resource, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("audio_track_set_key_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3886397084) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), gdc.ConstTypePtr(stream.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) AudioTrackSetKeyStartOffset(track_idx int, key_idx int, offset float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("audio_track_set_key_start_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) AudioTrackSetKeyEndOffset(track_idx int, key_idx int, offset float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("audio_track_set_key_end_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3506521499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), gdc.ConstTypePtr(&offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) AudioTrackGetKeyStream(track_idx int, key_idx int, ) Resource {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("audio_track_get_key_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 635277205) // FIXME: should cache?
  var ret Resource
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) AudioTrackGetKeyStartOffset(track_idx int, key_idx int, ) float32 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("audio_track_get_key_start_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) AudioTrackGetKeyEndOffset(track_idx int, key_idx int, ) float32 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("audio_track_get_key_end_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) AudioTrackSetUseBlend(track_idx int, enable bool, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("audio_track_set_use_blend")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) AudioTrackIsUseBlend(track_idx int, ) bool {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("audio_track_is_use_blend")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) AnimationTrackInsertKey(track_idx int, time float32, animation StringName, ) int {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("animation_track_insert_key")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 158676774) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&time), gdc.ConstTypePtr(animation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) AnimationTrackSetKeyAnimation(track_idx int, key_idx int, animation StringName, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("animation_track_set_key_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 117615382) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), gdc.ConstTypePtr(animation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) AnimationTrackGetKeyAnimation(track_idx int, key_idx int, ) StringName {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("animation_track_get_key_animation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 351665558) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(&key_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) SetLength(time_sec float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&time_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) GetLength() float32 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) SetLoopMode(loop_mode AnimationLoopMode, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3155355575) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) GetLoopMode() AnimationLoopMode {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loop_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1988889481) // FIXME: should cache?
  var ret AnimationLoopMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) SetStep(size_sec float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_sec), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) GetStep() float32 {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_step")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Animation) Clear()  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) CopyTrack(track_idx int, to_animation Animation, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("copy_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 148001024) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&track_idx), gdc.ConstTypePtr(to_animation.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Animation) Compress(page_size int, fps int, split_tolerance float32, )  {
  classNameV := StringNameFromStr("Animation")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("compress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3608408117) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&page_size), gdc.ConstTypePtr(&fps), gdc.ConstTypePtr(&split_tolerance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
