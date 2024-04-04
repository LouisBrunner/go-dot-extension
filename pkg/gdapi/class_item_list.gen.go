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

type ItemList struct {
  Control
}

func (me *ItemList) BaseClass() string {
  return "ItemList"
}

func NewItemList() *ItemList {
  str := StringNameFromStr("ItemList") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ItemList{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ItemListIconMode int
const (
  ItemListIconModeIconModeTop ItemListIconMode = 0
  ItemListIconModeIconModeLeft ItemListIconMode = 1
)

type ItemListSelectMode int
const (
  ItemListSelectModeSelectSingle ItemListSelectMode = 0
  ItemListSelectModeSelectMulti ItemListSelectMode = 1
)

func (me *ItemList) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ItemList) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ItemList) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ItemList) AddItem(text String, icon Texture2D, selectable bool, ) int64 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 359861678) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), icon.AsCTypePtr(), gdc.ConstTypePtr(&selectable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&selectable)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) AddIconItem(icon Texture2D, selectable bool, ) int64 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_icon_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4256579627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{icon.AsCTypePtr(), gdc.ConstTypePtr(&selectable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&selectable)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetItemText(idx int64, text String, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemText(idx int64, ) String {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetItemIcon(idx int64, icon Texture2D, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , icon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemIcon(idx int64, ) Texture2D {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetItemTextDirection(idx int64, direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1707680378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemTextDirection(idx int64, ) ControlTextDirection {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4235602388) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ControlTextDirection
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ItemList) SetItemLanguage(idx int64, language String, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , language.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemLanguage(idx int64, ) String {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetItemIconTransposed(idx int64, transposed bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_icon_transposed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&transposed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) IsItemIconTransposed(idx int64, ) bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_icon_transposed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetItemIconRegion(idx int64, rect Rect2, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_icon_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1356297692) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , rect.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemIconRegion(idx int64, ) Rect2 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_icon_region")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3327874267) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetItemIconModulate(idx int64, modulate Color, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_icon_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , modulate.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemIconModulate(idx int64, ) Color {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_icon_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetItemSelectable(idx int64, selectable bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&selectable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) IsItemSelectable(idx int64, ) bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_selectable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetItemDisabled(idx int64, disabled bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) IsItemDisabled(idx int64, ) bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetItemMetadata(idx int64, metadata Variant, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , metadata.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemMetadata(idx int64, ) Variant {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetItemCustomBgColor(idx int64, custom_bg_color Color, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_custom_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , custom_bg_color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemCustomBgColor(idx int64, ) Color {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_custom_bg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetItemCustomFgColor(idx int64, custom_fg_color Color, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_custom_fg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , custom_fg_color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemCustomFgColor(idx int64, ) Color {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_custom_fg_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) GetItemRect(idx int64, expand bool, ) Rect2 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 159227807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&expand) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()
  pinner.Pin(&idx)
  pinner.Pin(&expand)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetItemTooltipEnabled(idx int64, enable bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_tooltip_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) IsItemTooltipEnabled(idx int64, ) bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_tooltip_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetItemTooltip(idx int64, tooltip String, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , tooltip.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemTooltip(idx int64, ) String {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) Select(idx int64, single bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 972357352) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&single) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) Deselect(idx int64, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) DeselectAll()  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("deselect_all")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) IsSelected(idx int64, ) bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) GetSelectedItems() PackedInt32Array {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_selected_items")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 969006518) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) MoveItem(from_idx int64, to_idx int64, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_idx) , gdc.ConstTypePtr(&to_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) SetItemCount(count int64, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetItemCount() int64 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) RemoveItem(idx int64, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) Clear()  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) SortItemsByText()  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("sort_items_by_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) SetFixedColumnWidth(width int64, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fixed_column_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetFixedColumnWidth() int64 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fixed_column_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetSameColumnWidth(enable bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_same_column_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) IsSameColumnWidth() bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_same_column_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetMaxTextLines(lines int64, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_text_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lines) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetMaxTextLines() int64 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_text_lines")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetMaxColumns(amount int64, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetMaxColumns() int64 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_columns")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetSelectMode(mode ItemListSelectMode, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_select_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 928267388) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetSelectMode() ItemListSelectMode {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_select_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1191945842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ItemListSelectMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ItemList) SetIconMode(mode ItemListIconMode, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2025053633) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetIconMode() ItemListIconMode {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3353929232) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ItemListIconMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ItemList) SetFixedIconSize(size Vector2i, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fixed_icon_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetFixedIconSize() Vector2i {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fixed_icon_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2i()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetIconScale(scale float64, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetIconScale() float64 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_icon_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetAllowRmbSelect(allow bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_rmb_select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetAllowRmbSelect() bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_allow_rmb_select")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetAllowReselect(allow bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_reselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetAllowReselect() bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_allow_reselect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetAllowSearch(allow bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_search")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetAllowSearch() bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_allow_search")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) SetAutoHeight(enable bool, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) HasAutoHeight() bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_auto_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) IsAnythingSelected() bool {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_anything_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) GetItemAtPosition(position Vector2, exact bool, ) int64 {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_at_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2300324924) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&exact) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&exact)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *ItemList) EnsureCurrentIsVisible()  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("ensure_current_is_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetVScrollBar() VScrollBar {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_v_scroll_bar")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2630340773) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVScrollBar()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ItemList) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior, )  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1008890932) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ItemList) GetTextOverrunBehavior() TextServerOverrunBehavior {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_text_overrun_behavior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3779142101) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret TextServerOverrunBehavior

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ItemList) ForceUpdateListSize()  {
  classNameV := StringNameFromStr("ItemList")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update_list_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ItemListItemSelectedSignalFn func(index int, )

func (me *ItemList) ConnectItemSelected(subs SignalSubscribers, fn ItemListItemSelectedSignalFn) {
  sig := StringNameFromStr("item_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ItemList) DisconnectItemSelected(subs SignalSubscribers, fn ItemListItemSelectedSignalFn) {
  sig := StringNameFromStr("item_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ItemListEmptyClickedSignalFn func(at_position Vector2, mouse_button_index int, )

func (me *ItemList) ConnectEmptyClicked(subs SignalSubscribers, fn ItemListEmptyClickedSignalFn) {
  sig := StringNameFromStr("empty_clicked")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ItemList) DisconnectEmptyClicked(subs SignalSubscribers, fn ItemListEmptyClickedSignalFn) {
  sig := StringNameFromStr("empty_clicked")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ItemListItemClickedSignalFn func(index int, at_position Vector2, mouse_button_index int, )

func (me *ItemList) ConnectItemClicked(subs SignalSubscribers, fn ItemListItemClickedSignalFn) {
  sig := StringNameFromStr("item_clicked")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ItemList) DisconnectItemClicked(subs SignalSubscribers, fn ItemListItemClickedSignalFn) {
  sig := StringNameFromStr("item_clicked")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ItemListMultiSelectedSignalFn func(index int, selected bool, )

func (me *ItemList) ConnectMultiSelected(subs SignalSubscribers, fn ItemListMultiSelectedSignalFn) {
  sig := StringNameFromStr("multi_selected")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ItemList) DisconnectMultiSelected(subs SignalSubscribers, fn ItemListMultiSelectedSignalFn) {
  sig := StringNameFromStr("multi_selected")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type ItemListItemActivatedSignalFn func(index int, )

func (me *ItemList) ConnectItemActivated(subs SignalSubscribers, fn ItemListItemActivatedSignalFn) {
  sig := StringNameFromStr("item_activated")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *ItemList) DisconnectItemActivated(subs SignalSubscribers, fn ItemListItemActivatedSignalFn) {
  sig := StringNameFromStr("item_activated")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
