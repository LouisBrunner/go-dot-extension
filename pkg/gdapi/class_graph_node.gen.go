// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GraphNode struct {
  obj gdc.ObjectPtr
}

func (me *GraphNode) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GraphNode) BaseClass() string {
  return "GraphNode"
}



// Enums

type GraphNodeOverlay int
const (
  GraphNodeOverlayOverlayDisabled GraphNodeOverlay = 0
  GraphNodeOverlayOverlayBreakpoint GraphNodeOverlay = 1
  GraphNodeOverlayOverlayPosition GraphNodeOverlay = 2
)

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
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetTextDirection(direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 119160795) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) GetTextDirection() ControlTextDirection {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797257663) // FIXME: should cache?
  var ret ControlTextDirection
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetLanguage(language String, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) GetLanguage() String {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetSlot(slot_index int, enable_left_port bool, type_left int, color_left Color, enable_right_port bool, type_right int, color_right Color, custom_icon_left Texture2D, custom_icon_right Texture2D, draw_stylebox bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 902131739) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable_left_port), gdc.ConstTypePtr(&type_left), gdc.ConstTypePtr(color_left.AsCTypePtr()), gdc.ConstTypePtr(&enable_right_port), gdc.ConstTypePtr(&type_right), gdc.ConstTypePtr(color_right.AsCTypePtr()), gdc.ConstTypePtr(custom_icon_left.AsCTypePtr()), gdc.ConstTypePtr(custom_icon_right.AsCTypePtr()), gdc.ConstTypePtr(&draw_stylebox), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) ClearSlot(slot_index int, )  {
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

func  (me *GraphNode) SetSlotEnabledLeft(slot_index int, enable bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_enabled_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) IsSlotEnabledLeft(slot_index int, ) bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_slot_enabled_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetSlotTypeLeft(slot_index int, type_ int, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_type_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) GetSlotTypeLeft(slot_index int, ) int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slot_type_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetSlotColorLeft(slot_index int, color Color, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_color_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) GetSlotColorLeft(slot_index int, ) Color {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slot_color_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetSlotEnabledRight(slot_index int, enable bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_enabled_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) IsSlotEnabledRight(slot_index int, ) bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_slot_enabled_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetSlotTypeRight(slot_index int, type_ int, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_type_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&type_), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) GetSlotTypeRight(slot_index int, ) int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slot_type_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetSlotColorRight(slot_index int, color Color, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_color_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) GetSlotColorRight(slot_index int, ) Color {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_slot_color_right")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) IsSlotDrawStylebox(slot_index int, ) bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_slot_draw_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetSlotDrawStylebox(slot_index int, enable bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_slot_draw_stylebox")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&slot_index), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) SetPositionOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) GetPositionOffset() Vector2 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetComment(comment bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_comment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&comment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) IsComment() bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_comment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetResizable(resizable bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resizable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resizable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) IsResizable() bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_resizable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetDraggable(draggable bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draggable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draggable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) IsDraggable() bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draggable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetSelectable(selectable bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&selectable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) IsSelectable() bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetSelected(selected bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&selected), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) IsSelected() bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionInputCount() int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_input_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionInputHeight(port int, ) int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_input_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionInputPosition(port int, ) Vector2 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_input_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3114997196) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionInputType(port int, ) int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_input_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionInputColor(port int, ) Color {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_input_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2624840992) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionInputSlot(port int, ) int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_input_slot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionOutputCount() int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_output_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionOutputHeight(port int, ) int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_output_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionOutputPosition(port int, ) Vector2 {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_output_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3114997196) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionOutputType(port int, ) int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_output_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionOutputColor(port int, ) Color {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_output_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2624840992) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) GetConnectionOutputSlot(port int, ) int {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_connection_output_slot")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&port), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetShowCloseButton(show bool, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_close_button")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&show), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) IsCloseButtonVisible() bool {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_close_button_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphNode) SetOverlay(overlay GraphNodeOverlay, )  {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_overlay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3144190109) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overlay), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphNode) GetOverlay() GraphNodeOverlay {
  classNameV := StringNameFromStr("GraphNode")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_overlay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2854257040) // FIXME: should cache?
  var ret GraphNodeOverlay
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GraphNodePositionOffsetChangedSignalFn func()

func (me *GraphNode) ConnectPositionOffsetChanged(subs SignalSubscribers, fn GraphNodePositionOffsetChangedSignalFn) {
  sig := StringNameFromStr("position_offset_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectPositionOffsetChanged(subs SignalSubscribers, fn GraphNodePositionOffsetChangedSignalFn) {
  sig := StringNameFromStr("position_offset_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphNodeNodeSelectedSignalFn func()

func (me *GraphNode) ConnectNodeSelected(subs SignalSubscribers, fn GraphNodeNodeSelectedSignalFn) {
  sig := StringNameFromStr("node_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectNodeSelected(subs SignalSubscribers, fn GraphNodeNodeSelectedSignalFn) {
  sig := StringNameFromStr("node_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphNodeNodeDeselectedSignalFn func()

func (me *GraphNode) ConnectNodeDeselected(subs SignalSubscribers, fn GraphNodeNodeDeselectedSignalFn) {
  sig := StringNameFromStr("node_deselected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectNodeDeselected(subs SignalSubscribers, fn GraphNodeNodeDeselectedSignalFn) {
  sig := StringNameFromStr("node_deselected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphNodeSlotUpdatedSignalFn func(idx int, )

func (me *GraphNode) ConnectSlotUpdated(subs SignalSubscribers, fn GraphNodeSlotUpdatedSignalFn) {
  sig := StringNameFromStr("slot_updated")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectSlotUpdated(subs SignalSubscribers, fn GraphNodeSlotUpdatedSignalFn) {
  sig := StringNameFromStr("slot_updated")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphNodeDraggedSignalFn func(from Vector2, to Vector2, )

func (me *GraphNode) ConnectDragged(subs SignalSubscribers, fn GraphNodeDraggedSignalFn) {
  sig := StringNameFromStr("dragged")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectDragged(subs SignalSubscribers, fn GraphNodeDraggedSignalFn) {
  sig := StringNameFromStr("dragged")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphNodeRaiseRequestSignalFn func()

func (me *GraphNode) ConnectRaiseRequest(subs SignalSubscribers, fn GraphNodeRaiseRequestSignalFn) {
  sig := StringNameFromStr("raise_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectRaiseRequest(subs SignalSubscribers, fn GraphNodeRaiseRequestSignalFn) {
  sig := StringNameFromStr("raise_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphNodeCloseRequestSignalFn func()

func (me *GraphNode) ConnectCloseRequest(subs SignalSubscribers, fn GraphNodeCloseRequestSignalFn) {
  sig := StringNameFromStr("close_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectCloseRequest(subs SignalSubscribers, fn GraphNodeCloseRequestSignalFn) {
  sig := StringNameFromStr("close_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphNodeResizeRequestSignalFn func(new_minsize Vector2, )

func (me *GraphNode) ConnectResizeRequest(subs SignalSubscribers, fn GraphNodeResizeRequestSignalFn) {
  sig := StringNameFromStr("resize_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphNode) DisconnectResizeRequest(subs SignalSubscribers, fn GraphNodeResizeRequestSignalFn) {
  sig := StringNameFromStr("resize_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
