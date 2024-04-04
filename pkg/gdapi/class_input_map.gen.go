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

type InputMap struct {
  Object
}

func (me *InputMap) BaseClass() string {
  return "InputMap"
}

func NewInputMap() *InputMap {
  str := StringNameFromStr("InputMap") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &InputMap{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputMap) GetActions() []StringName {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_actions")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[StringName](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *InputMap) AddAction(action StringName, deadzone float64, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4100757082) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&deadzone) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputMap) EraseAction(action StringName, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("erase_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputMap) ActionSetDeadzone(action StringName, deadzone float64, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_set_deadzone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4135858297) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), gdc.ConstTypePtr(&deadzone) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputMap) ActionGetDeadzone(action StringName, ) float64 {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_get_deadzone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1391627649) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputMap) ActionAddEvent(action StringName, event InputEvent, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_add_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 518302593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), event.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputMap) ActionHasEvent(action StringName, event InputEvent, ) bool {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_has_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1185871985) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), event.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputMap) ActionEraseEvent(action StringName, event InputEvent, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_erase_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 518302593) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), event.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputMap) ActionEraseEvents(action StringName, )  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_erase_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *InputMap) ActionGetEvents(action StringName, ) []InputEvent {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("action_get_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 689397652) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{action.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[InputEvent](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *InputMap) EventIsAction(event InputEvent, action StringName, exact_match bool, ) bool {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("event_is_action")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3193353650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), action.AsCTypePtr(), gdc.ConstTypePtr(&exact_match) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&exact_match)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *InputMap) LoadFromProjectSettings()  {
  classNameV := StringNameFromStr("InputMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_project_settings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
