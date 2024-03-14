// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GraphElement struct {
  Container
}

func (me *GraphElement) BaseClass() string {
  return "GraphElement"
}



// Enums

func (me *GraphElement) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GraphElement) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GraphElement) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GraphElement) SetResizable(resizable bool, )  {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resizable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resizable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphElement) IsResizable() bool {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_resizable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphElement) SetDraggable(draggable bool, )  {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_draggable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draggable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphElement) IsDraggable() bool {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_draggable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphElement) SetSelectable(selectable bool, )  {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&selectable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphElement) IsSelectable() bool {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphElement) SetSelected(selected bool, )  {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&selected), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphElement) IsSelected() bool {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *GraphElement) SetPositionOffset(offset Vector2, )  {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *GraphElement) GetPositionOffset() Vector2 {
  classNameV := StringNameFromStr("GraphElement")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type GraphElementNodeSelectedSignalFn func()

func (me *GraphElement) ConnectNodeSelected(subs SignalSubscribers, fn GraphElementNodeSelectedSignalFn) {
  sig := StringNameFromStr("node_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectNodeSelected(subs SignalSubscribers, fn GraphElementNodeSelectedSignalFn) {
  sig := StringNameFromStr("node_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphElementNodeDeselectedSignalFn func()

func (me *GraphElement) ConnectNodeDeselected(subs SignalSubscribers, fn GraphElementNodeDeselectedSignalFn) {
  sig := StringNameFromStr("node_deselected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectNodeDeselected(subs SignalSubscribers, fn GraphElementNodeDeselectedSignalFn) {
  sig := StringNameFromStr("node_deselected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphElementRaiseRequestSignalFn func()

func (me *GraphElement) ConnectRaiseRequest(subs SignalSubscribers, fn GraphElementRaiseRequestSignalFn) {
  sig := StringNameFromStr("raise_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectRaiseRequest(subs SignalSubscribers, fn GraphElementRaiseRequestSignalFn) {
  sig := StringNameFromStr("raise_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphElementDeleteRequestSignalFn func()

func (me *GraphElement) ConnectDeleteRequest(subs SignalSubscribers, fn GraphElementDeleteRequestSignalFn) {
  sig := StringNameFromStr("delete_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectDeleteRequest(subs SignalSubscribers, fn GraphElementDeleteRequestSignalFn) {
  sig := StringNameFromStr("delete_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphElementResizeRequestSignalFn func(new_minsize Vector2, )

func (me *GraphElement) ConnectResizeRequest(subs SignalSubscribers, fn GraphElementResizeRequestSignalFn) {
  sig := StringNameFromStr("resize_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectResizeRequest(subs SignalSubscribers, fn GraphElementResizeRequestSignalFn) {
  sig := StringNameFromStr("resize_request")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphElementDraggedSignalFn func(from Vector2, to Vector2, )

func (me *GraphElement) ConnectDragged(subs SignalSubscribers, fn GraphElementDraggedSignalFn) {
  sig := StringNameFromStr("dragged")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectDragged(subs SignalSubscribers, fn GraphElementDraggedSignalFn) {
  sig := StringNameFromStr("dragged")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type GraphElementPositionOffsetChangedSignalFn func()

func (me *GraphElement) ConnectPositionOffsetChanged(subs SignalSubscribers, fn GraphElementPositionOffsetChangedSignalFn) {
  sig := StringNameFromStr("position_offset_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *GraphElement) DisconnectPositionOffsetChanged(subs SignalSubscribers, fn GraphElementPositionOffsetChangedSignalFn) {
  sig := StringNameFromStr("position_offset_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}