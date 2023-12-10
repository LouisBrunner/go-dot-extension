// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type InputMap struct {
  obj gdc.ObjectPtr
}

func (me *InputMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputMap) BaseClass() string {
  return "InputMap"
}



// Enums

func (me *InputMap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputMap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputMap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *InputMap) HasAction(action StringName, ) bool {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2619796661) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputMap) GetActions() StringName {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_actions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputMap) AddAction(action StringName, deadzone float32, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 573731101) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&deadzone), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputMap) EraseAction(action StringName, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputMap) ActionSetDeadzone(action StringName, deadzone float32, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_set_deadzone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4135858297) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&deadzone), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputMap) ActionGetDeadzone(action StringName, ) float32 {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_get_deadzone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1391627649) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputMap) ActionAddEvent(action StringName, event InputEvent, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_add_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 518302593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(event.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputMap) ActionHasEvent(action StringName, event InputEvent, ) bool {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_has_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1185871985) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(event.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputMap) ActionEraseEvent(action StringName, event InputEvent, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_erase_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 518302593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(event.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputMap) ActionEraseEvents(action StringName, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_erase_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *InputMap) ActionGetEvents(action StringName, ) InputEvent {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_get_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 689397652) // FIXME: should cache?
  var ret InputEvent
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(action.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputMap) EventIsAction(event InputEvent, action StringName, exact_match bool, ) bool {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("event_is_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3193353650) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(event.AsCTypePtr()), gdc.ConstTypePtr(action.AsCTypePtr()), gdc.ConstTypePtr(&exact_match), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *InputMap) LoadFromProjectSettings()  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_project_settings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties