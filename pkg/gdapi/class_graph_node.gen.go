// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *GraphNode) GetTitle()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetTextDirection(direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetTextDirection()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetLanguage(language String, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetLanguage()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSlot(slot_index int, enable_left_port bool, type_left int, color_left Color, enable_right_port bool, type_right int, color_right Color, custom_icon_left Texture2D, custom_icon_right Texture2D, draw_stylebox bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) ClearSlot(slot_index int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) ClearAllSlots()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSlotEnabledLeft(slot_index int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) IsSlotEnabledLeft(slot_index int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSlotTypeLeft(slot_index int, type_ int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetSlotTypeLeft(slot_index int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSlotColorLeft(slot_index int, color Color, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetSlotColorLeft(slot_index int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSlotEnabledRight(slot_index int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) IsSlotEnabledRight(slot_index int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSlotTypeRight(slot_index int, type_ int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetSlotTypeRight(slot_index int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSlotColorRight(slot_index int, color Color, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetSlotColorRight(slot_index int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) IsSlotDrawStylebox(slot_index int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSlotDrawStylebox(slot_index int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetPositionOffset(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetPositionOffset()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetComment(comment bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) IsComment()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetResizable(resizable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) IsResizable()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetDraggable(draggable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) IsDraggable()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSelectable(selectable bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) IsSelectable()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetSelected(selected bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) IsSelected()  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionInputCount()  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionInputHeight(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionInputPosition(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionInputType(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionInputColor(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionInputSlot(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionOutputCount()  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionOutputHeight(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionOutputPosition(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionOutputType(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionOutputColor(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetConnectionOutputSlot(port int, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetShowCloseButton(show bool, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) IsCloseButtonVisible()  {
  panic("TODO: implement")
}

func  (me *GraphNode) SetOverlay(overlay GraphNodeOverlay, )  {
  panic("TODO: implement")
}

func  (me *GraphNode) GetOverlay()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
