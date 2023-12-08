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

type GraphNodeOverlay int
const (
  GraphNodeOverlayOverlayDisabled GraphNodeOverlay = 0
  GraphNodeOverlayOverlayBreakpoint GraphNodeOverlay = 1
  GraphNodeOverlayOverlayPosition GraphNodeOverlay = 2
)

func  (me *GraphNode) SetTitle(title String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetTitle() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetTextDirection(direction ControlTextDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetTextDirection() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetLanguage(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetLanguage() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSlot(slot_index int, enable_left_port bool, type_left int, color_left Color, enable_right_port bool, type_right int, color_right Color, custom_icon_left Texture2D, custom_icon_right Texture2D, draw_stylebox bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) ClearSlot(slot_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) ClearAllSlots() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSlotEnabledLeft(slot_index int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) IsSlotEnabledLeft(slot_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSlotTypeLeft(slot_index int, type_ int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetSlotTypeLeft(slot_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSlotColorLeft(slot_index int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetSlotColorLeft(slot_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSlotEnabledRight(slot_index int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) IsSlotEnabledRight(slot_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSlotTypeRight(slot_index int, type_ int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetSlotTypeRight(slot_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSlotColorRight(slot_index int, color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetSlotColorRight(slot_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) IsSlotDrawStylebox(slot_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSlotDrawStylebox(slot_index int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetPositionOffset(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetPositionOffset() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetComment(comment bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) IsComment() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetResizable(resizable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) IsResizable() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetDraggable(draggable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) IsDraggable() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSelectable(selectable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) IsSelectable() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetSelected(selected bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) IsSelected() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionInputCount() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionInputHeight(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionInputPosition(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionInputType(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionInputColor(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionInputSlot(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionOutputCount() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionOutputHeight(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionOutputPosition(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionOutputType(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionOutputColor(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetConnectionOutputSlot(port int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetShowCloseButton(show bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) IsCloseButtonVisible() { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) SetOverlay(overlay GraphNodeOverlay, ) { // TODO: return value
  // TODO: implement
}

func  (me *GraphNode) GetOverlay() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
