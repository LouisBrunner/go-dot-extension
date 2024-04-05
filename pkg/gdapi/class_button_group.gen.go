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

type ptrsForButtonGroupList struct {
  fnGetPressedButton gdc.MethodBindPtr
  fnGetButtons gdc.MethodBindPtr
  fnSetAllowUnpress gdc.MethodBindPtr
  fnIsAllowUnpress gdc.MethodBindPtr
}

var ptrsForButtonGroup ptrsForButtonGroupList

func initButtonGroupPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ButtonGroup")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_pressed_button")
    defer methodName.Destroy()
    ptrsForButtonGroup.fnGetPressedButton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3886434893))
  }
  {
    methodName := StringNameFromStr("get_buttons")
    defer methodName.Destroy()
    ptrsForButtonGroup.fnGetButtons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_allow_unpress")
    defer methodName.Destroy()
    ptrsForButtonGroup.fnSetAllowUnpress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_allow_unpress")
    defer methodName.Destroy()
    ptrsForButtonGroup.fnIsAllowUnpress = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
}

type ButtonGroup struct {
  Resource
}

func (me *ButtonGroup) BaseClass() string {
  return "ButtonGroup"
}

func NewButtonGroup() *ButtonGroup {
  str := StringNameFromStr("ButtonGroup") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ButtonGroup{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ButtonGroup) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ButtonGroup) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ButtonGroup) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ButtonGroup) GetPressedButton() BaseButton {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBaseButton()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButtonGroup.fnGetPressedButton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ButtonGroup) GetButtons() []BaseButton {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButtonGroup.fnGetButtons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[BaseButton](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *ButtonGroup) SetAllowUnpress(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButtonGroup.fnSetAllowUnpress), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ButtonGroup) IsAllowUnpress() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForButtonGroup.fnIsAllowUnpress), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ButtonGroupPressedSignalFn func(button BaseButton, )

func (me *ButtonGroup) ConnectPressed(subs SignalSubscribers, fn ButtonGroupPressedSignalFn) {
  sig := StringNameFromStr("pressed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ButtonGroup) DisconnectPressed(subs SignalSubscribers, fn ButtonGroupPressedSignalFn) {
  sig := StringNameFromStr("pressed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
