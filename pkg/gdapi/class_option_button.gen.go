// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OptionButton struct {
  obj gdc.ObjectPtr
}

func (me *OptionButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OptionButton) BaseClass() string {
  return "OptionButton"
}



// Enums

func (me *OptionButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OptionButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OptionButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OptionButton) AddItem(label String, id int, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3043792800) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) AddIconItem(texture Texture2D, label String, id int, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_icon_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3944051090) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) SetItemText(idx int, text String, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) SetItemIcon(idx int, texture Texture2D, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) SetItemDisabled(idx int, disabled bool, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) SetItemId(idx int, id int, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) SetItemMetadata(idx int, metadata Variant, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(metadata.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) SetItemTooltip(idx int, tooltip String, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(tooltip.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) GetItemText(idx int, ) String {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) GetItemIcon(idx int, ) Texture2D {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) GetItemId(idx int, ) int {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) GetItemIndex(id int, ) int {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) GetItemMetadata(idx int, ) Variant {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) GetItemTooltip(idx int, ) String {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) IsItemDisabled(idx int, ) bool {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) IsItemSeparator(idx int, ) bool {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_separator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) AddSeparator(text String, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_separator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3005725572) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) Clear()  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) Select(idx int, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) GetSelected() int {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) GetSelectedId() int {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) GetSelectedMetadata() Variant {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1214101251) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) RemoveItem(idx int, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) GetPopup() PopupMenu {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 229722558) // FIXME: should cache?
  var ret PopupMenu
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) ShowPopup()  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("show_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) SetItemCount(count int, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) GetItemCount() int {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) HasSelectableItems() bool {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_selectable_items")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) GetSelectableItem(from_last bool, ) int {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selectable_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 894402480) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_last), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) SetFitToLongestItem(fit bool, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fit_to_longest_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fit), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) IsFitToLongestItem() bool {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_fit_to_longest_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OptionButton) SetAllowReselect(allow bool, )  {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_reselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OptionButton) GetAllowReselect() bool {
  classNameV := StringNameFromStr("OptionButton")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_allow_reselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type OptionButtonItemSelectedSignalFn func(index int, )

func (me *OptionButton) ConnectItemSelected(subs SignalSubscribers, fn OptionButtonItemSelectedSignalFn) {
  sig := StringNameFromStr("item_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *OptionButton) DisconnectItemSelected(subs SignalSubscribers, fn OptionButtonItemSelectedSignalFn) {
  sig := StringNameFromStr("item_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type OptionButtonItemFocusedSignalFn func(index int, )

func (me *OptionButton) ConnectItemFocused(subs SignalSubscribers, fn OptionButtonItemFocusedSignalFn) {
  sig := StringNameFromStr("item_focused")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *OptionButton) DisconnectItemFocused(subs SignalSubscribers, fn OptionButtonItemFocusedSignalFn) {
  sig := StringNameFromStr("item_focused")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
