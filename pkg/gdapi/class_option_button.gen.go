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

type ptrsForOptionButtonList struct {
  fnAddItem gdc.MethodBindPtr
  fnAddIconItem gdc.MethodBindPtr
  fnSetItemText gdc.MethodBindPtr
  fnSetItemIcon gdc.MethodBindPtr
  fnSetItemDisabled gdc.MethodBindPtr
  fnSetItemId gdc.MethodBindPtr
  fnSetItemMetadata gdc.MethodBindPtr
  fnSetItemTooltip gdc.MethodBindPtr
  fnGetItemText gdc.MethodBindPtr
  fnGetItemIcon gdc.MethodBindPtr
  fnGetItemId gdc.MethodBindPtr
  fnGetItemIndex gdc.MethodBindPtr
  fnGetItemMetadata gdc.MethodBindPtr
  fnGetItemTooltip gdc.MethodBindPtr
  fnIsItemDisabled gdc.MethodBindPtr
  fnIsItemSeparator gdc.MethodBindPtr
  fnAddSeparator gdc.MethodBindPtr
  fnClear gdc.MethodBindPtr
  fnSelect gdc.MethodBindPtr
  fnGetSelected gdc.MethodBindPtr
  fnGetSelectedId gdc.MethodBindPtr
  fnGetSelectedMetadata gdc.MethodBindPtr
  fnRemoveItem gdc.MethodBindPtr
  fnGetPopup gdc.MethodBindPtr
  fnShowPopup gdc.MethodBindPtr
  fnSetItemCount gdc.MethodBindPtr
  fnGetItemCount gdc.MethodBindPtr
  fnHasSelectableItems gdc.MethodBindPtr
  fnGetSelectableItem gdc.MethodBindPtr
  fnSetFitToLongestItem gdc.MethodBindPtr
  fnIsFitToLongestItem gdc.MethodBindPtr
  fnSetAllowReselect gdc.MethodBindPtr
  fnGetAllowReselect gdc.MethodBindPtr
  fnSetDisableShortcuts gdc.MethodBindPtr
}

var ptrsForOptionButton ptrsForOptionButtonList

func initOptionButtonPtrs(iface gdc.Interface) {

  className := StringNameFromStr("OptionButton")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_item")
    defer methodName.Destroy()
    ptrsForOptionButton.fnAddItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2697778442))
  }
  {
    methodName := StringNameFromStr("add_icon_item")
    defer methodName.Destroy()
    ptrsForOptionButton.fnAddIconItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3781678508))
  }
  {
    methodName := StringNameFromStr("set_item_text")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
  }
  {
    methodName := StringNameFromStr("set_item_icon")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
  }
  {
    methodName := StringNameFromStr("set_item_disabled")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("set_item_id")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetItemId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("set_item_metadata")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetItemMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
  }
  {
    methodName := StringNameFromStr("set_item_tooltip")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
  }
  {
    methodName := StringNameFromStr("get_item_text")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
  }
  {
    methodName := StringNameFromStr("get_item_icon")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
  }
  {
    methodName := StringNameFromStr("get_item_id")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetItemId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("get_item_index")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetItemIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("get_item_metadata")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetItemMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
  }
  {
    methodName := StringNameFromStr("get_item_tooltip")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
  }
  {
    methodName := StringNameFromStr("is_item_disabled")
    defer methodName.Destroy()
    ptrsForOptionButton.fnIsItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("is_item_separator")
    defer methodName.Destroy()
    ptrsForOptionButton.fnIsItemSeparator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("add_separator")
    defer methodName.Destroy()
    ptrsForOptionButton.fnAddSeparator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3005725572))
  }
  {
    methodName := StringNameFromStr("clear")
    defer methodName.Destroy()
    ptrsForOptionButton.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("select")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSelect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_selected")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_selected_id")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetSelectedId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_selected_metadata")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetSelectedMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1214101251))
  }
  {
    methodName := StringNameFromStr("remove_item")
    defer methodName.Destroy()
    ptrsForOptionButton.fnRemoveItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_popup")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 229722558))
  }
  {
    methodName := StringNameFromStr("show_popup")
    defer methodName.Destroy()
    ptrsForOptionButton.fnShowPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_item_count")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_item_count")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("has_selectable_items")
    defer methodName.Destroy()
    ptrsForOptionButton.fnHasSelectableItems = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_selectable_item")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetSelectableItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 894402480))
  }
  {
    methodName := StringNameFromStr("set_fit_to_longest_item")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetFitToLongestItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_fit_to_longest_item")
    defer methodName.Destroy()
    ptrsForOptionButton.fnIsFitToLongestItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_allow_reselect")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetAllowReselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_allow_reselect")
    defer methodName.Destroy()
    ptrsForOptionButton.fnGetAllowReselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_disable_shortcuts")
    defer methodName.Destroy()
    ptrsForOptionButton.fnSetDisableShortcuts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
}

type OptionButton struct {
  Button
}

func (me *OptionButton) BaseClass() string {
  return "OptionButton"
}

func NewOptionButton() *OptionButton {
  str := StringNameFromStr("OptionButton") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OptionButton{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *OptionButton) AddItem(label String, id int64, )  {
  cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnAddItem), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) AddIconItem(texture Texture2D, label String, id int64, )  {
  cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), label.AsCTypePtr(), gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnAddIconItem), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) SetItemText(idx int64, text String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetItemText), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) SetItemIcon(idx int64, texture Texture2D, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , texture.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetItemIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) SetItemDisabled(idx int64, disabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetItemDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) SetItemId(idx int64, id int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetItemId), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) SetItemMetadata(idx int64, metadata Variant, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , metadata.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetItemMetadata), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) SetItemTooltip(idx int64, tooltip String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , tooltip.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetItemTooltip), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) GetItemText(idx int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetItemText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OptionButton) GetItemIcon(idx int64, ) Texture2D {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetItemIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OptionButton) GetItemId(idx int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetItemId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) GetItemIndex(id int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&id)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetItemIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) GetItemMetadata(idx int64, ) Variant {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetItemMetadata), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OptionButton) GetItemTooltip(idx int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetItemTooltip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OptionButton) IsItemDisabled(idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnIsItemDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) IsItemSeparator(idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnIsItemSeparator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) AddSeparator(text String, )  {
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnAddSeparator), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) Clear()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) Select(idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSelect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) GetSelected() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) GetSelectedId() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetSelectedId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) GetSelectedMetadata() Variant {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetSelectedMetadata), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OptionButton) RemoveItem(idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnRemoveItem), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) GetPopup() PopupMenu {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPopupMenu()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetPopup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *OptionButton) ShowPopup()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnShowPopup), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) SetItemCount(count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetItemCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) GetItemCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetItemCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) HasSelectableItems() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnHasSelectableItems), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) GetSelectableItem(from_last bool, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_last) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&from_last)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetSelectableItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) SetFitToLongestItem(fit bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fit) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetFitToLongestItem), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) IsFitToLongestItem() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnIsFitToLongestItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) SetAllowReselect(allow bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetAllowReselect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OptionButton) GetAllowReselect() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnGetAllowReselect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OptionButton) SetDisableShortcuts(disabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForOptionButton.fnSetDisableShortcuts), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type OptionButtonItemSelectedSignalFn func(index int, )

func (me *OptionButton) ConnectItemSelected(subs SignalSubscribers, fn OptionButtonItemSelectedSignalFn) {
  sig := StringNameFromStr("item_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *OptionButton) DisconnectItemSelected(subs SignalSubscribers, fn OptionButtonItemSelectedSignalFn) {
  sig := StringNameFromStr("item_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type OptionButtonItemFocusedSignalFn func(index int, )

func (me *OptionButton) ConnectItemFocused(subs SignalSubscribers, fn OptionButtonItemFocusedSignalFn) {
  sig := StringNameFromStr("item_focused")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *OptionButton) DisconnectItemFocused(subs SignalSubscribers, fn OptionButtonItemFocusedSignalFn) {
  sig := StringNameFromStr("item_focused")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
