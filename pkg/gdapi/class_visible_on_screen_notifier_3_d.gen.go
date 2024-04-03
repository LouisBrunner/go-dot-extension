// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisibleOnScreenNotifier3D struct {
  VisualInstance3D
}

func (me *VisibleOnScreenNotifier3D) BaseClass() string {
  return "VisibleOnScreenNotifier3D"
}

func NewVisibleOnScreenNotifier3D() *VisibleOnScreenNotifier3D {
  str := StringNameFromStr("VisibleOnScreenNotifier3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisibleOnScreenNotifier3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisibleOnScreenNotifier3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisibleOnScreenNotifier3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenNotifier3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisibleOnScreenNotifier3D) SetAabb(rect AABB, )  {
  classNameV := StringNameFromStr("VisibleOnScreenNotifier3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisibleOnScreenNotifier3D) IsOnScreen() bool {
  classNameV := StringNameFromStr("VisibleOnScreenNotifier3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_on_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type VisibleOnScreenNotifier3DScreenEnteredSignalFn func()

func (me *VisibleOnScreenNotifier3D) ConnectScreenEntered(subs SignalSubscribers, fn VisibleOnScreenNotifier3DScreenEnteredSignalFn) {
  sig := StringNameFromStr("screen_entered")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *VisibleOnScreenNotifier3D) DisconnectScreenEntered(subs SignalSubscribers, fn VisibleOnScreenNotifier3DScreenEnteredSignalFn) {
  sig := StringNameFromStr("screen_entered")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type VisibleOnScreenNotifier3DScreenExitedSignalFn func()

func (me *VisibleOnScreenNotifier3D) ConnectScreenExited(subs SignalSubscribers, fn VisibleOnScreenNotifier3DScreenExitedSignalFn) {
  sig := StringNameFromStr("screen_exited")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *VisibleOnScreenNotifier3D) DisconnectScreenExited(subs SignalSubscribers, fn VisibleOnScreenNotifier3DScreenExitedSignalFn) {
  sig := StringNameFromStr("screen_exited")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
