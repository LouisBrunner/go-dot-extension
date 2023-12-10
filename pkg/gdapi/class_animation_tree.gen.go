// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationTree struct {
  obj gdc.ObjectPtr
}

func (me *AnimationTree) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationTree) BaseClass() string {
  return "AnimationTree"
}



// Enums

type AnimationTreeAnimationProcessCallback int
const (
  AnimationTreeAnimationProcessCallbackAnimationProcessPhysics AnimationTreeAnimationProcessCallback = 0
  AnimationTreeAnimationProcessCallbackAnimationProcessIdle AnimationTreeAnimationProcessCallback = 1
  AnimationTreeAnimationProcessCallbackAnimationProcessManual AnimationTreeAnimationProcessCallback = 2
)

func (me *AnimationTree) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationTree) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationTree) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationTree) SetActive(active bool, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationTree) IsActive() bool {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) SetTreeRoot(root AnimationNode, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tree_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 712869711) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(root.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationTree) GetTreeRoot() AnimationNode {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tree_root")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1462070895) // FIXME: should cache?
  var ret AnimationNode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) SetProcessCallback(mode AnimationTreeAnimationProcessCallback, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1723352826) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationTree) GetProcessCallback() AnimationTreeAnimationProcessCallback {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_process_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 891317132) // FIXME: should cache?
  var ret AnimationTreeAnimationProcessCallback
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) SetAnimationPlayer(root NodePath, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animation_player")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(root.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationTree) GetAnimationPlayer() NodePath {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animation_player")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) SetAdvanceExpressionBaseNode(node NodePath, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_advance_expression_base_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationTree) GetAdvanceExpressionBaseNode() NodePath {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_advance_expression_base_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) SetRootMotionTrack(path NodePath, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_motion_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationTree) GetRootMotionTrack() NodePath {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_track")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) SetAudioMaxPolyphony(max_polyphony int, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_audio_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_polyphony), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationTree) GetAudioMaxPolyphony() int {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_audio_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) GetRootMotionPosition() Vector3 {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) GetRootMotionRotation() Quaternion {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1222331677) // FIXME: should cache?
  var ret Quaternion
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) GetRootMotionScale() Vector3 {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) GetRootMotionPositionAccumulator() Vector3 {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_position_accumulator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) GetRootMotionRotationAccumulator() Quaternion {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_rotation_accumulator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1222331677) // FIXME: should cache?
  var ret Quaternion
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) GetRootMotionScaleAccumulator() Vector3 {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_motion_scale_accumulator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationTree) Advance(delta float32, )  {
  classNameV := StringNameFromStr("AnimationTree")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("advance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *AnimationTree) GetPropTreeRoot() AnimationRootNode {
  panic("TODO: implement")
}

func (me *AnimationTree) SetPropTreeRoot(value AnimationRootNode) {
  panic("TODO: implement")
}

func (me *AnimationTree) GetPropAnimPlayer() NodePath {
  panic("TODO: implement")
}

func (me *AnimationTree) SetPropAnimPlayer(value NodePath) {
  panic("TODO: implement")
}

func (me *AnimationTree) GetPropAdvanceExpressionBaseNode() NodePath {
  panic("TODO: implement")
}

func (me *AnimationTree) SetPropAdvanceExpressionBaseNode(value NodePath) {
  panic("TODO: implement")
}

func (me *AnimationTree) GetPropActive() bool {
  panic("TODO: implement")
}

func (me *AnimationTree) SetPropActive(value bool) {
  panic("TODO: implement")
}

func (me *AnimationTree) GetPropProcessCallback() int {
  panic("TODO: implement")
}

func (me *AnimationTree) SetPropProcessCallback(value int) {
  panic("TODO: implement")
}

func (me *AnimationTree) GetPropAudioMaxPolyphony() int {
  panic("TODO: implement")
}

func (me *AnimationTree) SetPropAudioMaxPolyphony(value int) {
  panic("TODO: implement")
}

func (me *AnimationTree) GetPropRootMotionTrack() NodePath {
  panic("TODO: implement")
}

func (me *AnimationTree) SetPropRootMotionTrack(value NodePath) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API