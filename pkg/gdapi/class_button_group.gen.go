// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ButtonGroup struct {
  obj gdc.ObjectPtr
}

func (me *ButtonGroup) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ButtonGroup) BaseClass() string {
  return "ButtonGroup"
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
  classNameV := StringNameFromStr("ButtonGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pressed_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3886434893) // FIXME: should cache?
  var ret BaseButton
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ButtonGroup) GetButtons() BaseButton {
  classNameV := StringNameFromStr("ButtonGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buttons")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2915620761) // FIXME: should cache?
  var ret BaseButton
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ButtonGroup) SetAllowUnpress(enabled bool, )  {
  classNameV := StringNameFromStr("ButtonGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_unpress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ButtonGroup) IsAllowUnpress() bool {
  classNameV := StringNameFromStr("ButtonGroup")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_allow_unpress")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ButtonGroup) GetPropAllowUnpress() bool {
  panic("TODO: implement")
}

func (me *ButtonGroup) SetPropAllowUnpress(value bool) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API