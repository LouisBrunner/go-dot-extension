// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Tweener struct {
  RefCounted
}

func (me *Tweener) BaseClass() string {
  return "Tweener"
}

func NewTweener() *Tweener {
  str := StringNameFromStr("Tweener") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Tweener{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Tweener) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Tweener) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Tweener) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals

type TweenerFinishedSignalFn func()

func (me *Tweener) ConnectFinished(subs SignalSubscribers, fn TweenerFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Tweener) DisconnectFinished(subs SignalSubscribers, fn TweenerFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
