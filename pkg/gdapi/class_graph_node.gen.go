// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GraphNode struct {
  GraphElement
}

func (me *GraphNode) BaseClass() string {
  return "GraphNode"
}

func NewGraphNode() *GraphNode {
  str := StringNameFromStr("GraphNode") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GraphNode{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *GraphNode) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GraphNode) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GraphNode) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GraphNode) SetTitle(title String, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(title.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) GetTitle() String {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphNode) GetTitlebarHbox() HBoxContainer {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_titlebar_hbox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3590609951) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewHBoxContainer()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphNode) SetSlot(slot_index int64, enable_left_port bool, type_left int64, color_left Color, enable_right_port bool, type_right int64, color_right Color, custom_icon_left Texture2D, custom_icon_right Texture2D, draw_stylebox bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2873310869) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable_left_port), gdc.ConstTypePtr(&type_left), gdc.ConstTypePtr(color_left.AsCTypePtr()), gdc.ConstTypePtr(&enable_right_port), gdc.ConstTypePtr(&type_right), gdc.ConstTypePtr(color_right.AsCTypePtr()), gdc.ConstTypePtr(custom_icon_left.AsCTypePtr()), gdc.ConstTypePtr(custom_icon_right.AsCTypePtr()), gdc.ConstTypePtr(&draw_stylebox), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) ClearSlot(slot_index int64, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_slot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) ClearAllSlots()  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_all_slots")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) IsSlotEnabledLeft(slot_index int64, ) bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_slot_enabled_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) SetSlotEnabledLeft(slot_index int64, enable bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_enabled_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) SetSlotTypeLeft(slot_index int64, type_ int64, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_type_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&type_), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) GetSlotTypeLeft(slot_index int64, ) int64 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slot_type_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) SetSlotColorLeft(slot_index int64, color Color, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_color_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(color.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) GetSlotColorLeft(slot_index int64, ) Color {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slot_color_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphNode) IsSlotEnabledRight(slot_index int64, ) bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_slot_enabled_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) SetSlotEnabledRight(slot_index int64, enable bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_enabled_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) SetSlotTypeRight(slot_index int64, type_ int64, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_type_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&type_), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) GetSlotTypeRight(slot_index int64, ) int64 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slot_type_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) SetSlotColorRight(slot_index int64, color Color, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_color_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(color.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) GetSlotColorRight(slot_index int64, ) Color {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slot_color_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphNode) IsSlotDrawStylebox(slot_index int64, ) bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_slot_draw_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) SetSlotDrawStylebox(slot_index int64, enable bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_draw_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GraphNode) GetInputPortCount() int64 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_port_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) GetInputPortPosition(port_idx int64, ) Vector2 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_port_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3114997196) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx), }
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphNode) GetInputPortType(port_idx int64, ) int64 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_port_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) GetInputPortColor(port_idx int64, ) Color {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_port_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2624840992) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx), }
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphNode) GetInputPortSlot(port_idx int64, ) int64 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_port_slot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) GetOutputPortCount() int64 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_port_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) GetOutputPortPosition(port_idx int64, ) Vector2 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_port_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3114997196) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx), }
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphNode) GetOutputPortType(port_idx int64, ) int64 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_port_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GraphNode) GetOutputPortColor(port_idx int64, ) Color {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_port_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2624840992) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx), }
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GraphNode) GetOutputPortSlot(port_idx int64, ) int64 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_port_slot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port_idx), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GraphNodeSlotUpdatedSignalFn func(slot_index int, )

func (me *GraphNode) ConnectSlotUpdated(subs SignalSubscribers, fn GraphNodeSlotUpdatedSignalFn) {
  sig := StringNameFromStr("slot_updated")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectSlotUpdated(subs SignalSubscribers, fn GraphNodeSlotUpdatedSignalFn) {
  sig := StringNameFromStr("slot_updated")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
