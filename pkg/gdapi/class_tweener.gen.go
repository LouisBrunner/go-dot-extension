// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Tweener struct {
  obj gdc.ObjectPtr
}

func (me *Tweener) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Tweener) BaseClass() string {
  return "Tweener"
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
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Tweener) DisconnectFinished(subs SignalSubscribers, fn TweenerFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
