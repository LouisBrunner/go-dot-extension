// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type Shortcut struct {
  Resource
}

func (me *Shortcut) BaseClass() string {
  return "Shortcut"
}

func NewShortcut() *Shortcut {
  str := StringNameFromStr("Shortcut") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Shortcut{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Shortcut) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Shortcut) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Shortcut) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Shortcut) SetEvents(events Array, )  {
  classNameV := StringNameFromStr("Shortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 381264803) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{events.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Shortcut) GetEvents() Array {
  classNameV := StringNameFromStr("Shortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Shortcut) HasValidEvent() bool {
  classNameV := StringNameFromStr("Shortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_valid_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Shortcut) MatchesEvent(event InputEvent, ) bool {
  classNameV := StringNameFromStr("Shortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("matches_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3738334489) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Shortcut) GetAsText() String {
  classNameV := StringNameFromStr("Shortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_as_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
