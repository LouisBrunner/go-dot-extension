// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Shortcut struct {
  obj gdc.ObjectPtr
}

func (me *Shortcut) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Shortcut) BaseClass() string {
  return "Shortcut"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(events.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Shortcut) GetEvents() Array {
  classNameV := StringNameFromStr("Shortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shortcut) HasValidEvent() bool {
  classNameV := StringNameFromStr("Shortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_valid_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shortcut) MatchesEvent(event InputEvent, ) bool {
  classNameV := StringNameFromStr("Shortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("matches_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3738334489) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(event.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shortcut) GetAsText() String {
  classNameV := StringNameFromStr("Shortcut")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_as_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Shortcut) GetPropEvents() InputEvent {
  panic("TODO: implement")
}

func (me *Shortcut) SetPropEvents(value InputEvent) {
  panic("TODO: implement")
}