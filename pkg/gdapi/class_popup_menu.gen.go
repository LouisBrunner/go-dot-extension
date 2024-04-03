// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PopupMenu struct {
  Popup
}

func (me *PopupMenu) BaseClass() string {
  return "PopupMenu"
}

func NewPopupMenu() *PopupMenu {
  str := StringNameFromStr("PopupMenu") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PopupMenu{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PopupMenu) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PopupMenu) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PopupMenu) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PopupMenu) ActivateItemByEvent(event InputEvent, for_global_only bool, ) bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("activate_item_by_event")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3716412023) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(event.AsCTypePtr()), gdc.ConstTypePtr(&for_global_only), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) AddItem(label String, id int64, accel Key, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3674230041) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddIconItem(texture Texture2D, label String, id int64, accel Key, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_icon_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1086190128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddCheckItem(label String, id int64, accel Key, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_check_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3674230041) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddIconCheckItem(texture Texture2D, label String, id int64, accel Key, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_icon_check_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1086190128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddRadioCheckItem(label String, id int64, accel Key, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_radio_check_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3674230041) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddIconRadioCheckItem(texture Texture2D, label String, id int64, accel Key, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_icon_radio_check_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1086190128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddMultistateItem(label String, max_states int64, default_state int64, id int64, accel Key, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_multistate_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 150780458) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&max_states), gdc.ConstTypePtr(&default_state), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddShortcut(shortcut Shortcut, id int64, global bool, allow_echo bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3451850107) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shortcut.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global), gdc.ConstTypePtr(&allow_echo), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddIconShortcut(texture Texture2D, shortcut Shortcut, id int64, global bool, allow_echo bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_icon_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2997871092) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(shortcut.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global), gdc.ConstTypePtr(&allow_echo), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddCheckShortcut(shortcut Shortcut, id int64, global bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_check_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1642193386) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shortcut.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddIconCheckShortcut(texture Texture2D, shortcut Shortcut, id int64, global bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_icon_check_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3856247530) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(shortcut.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddRadioCheckShortcut(shortcut Shortcut, id int64, global bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_radio_check_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1642193386) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shortcut.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddIconRadioCheckShortcut(texture Texture2D, shortcut Shortcut, id int64, global bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_icon_radio_check_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3856247530) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), gdc.ConstTypePtr(shortcut.AsCTypePtr()), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddSubmenuItem(label String, submenu String, id int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_submenu_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2979222410) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(submenu.AsCTypePtr()), gdc.ConstTypePtr(&id), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemText(index int64, text String, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(text.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemTextDirection(index int64, direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1707680378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&direction), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemLanguage(index int64, language String, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(language.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemIcon(index int64, icon Texture2D, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(icon.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemIconMaxWidth(index int64, width int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_icon_max_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&width), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemIconModulate(index int64, modulate Color, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_icon_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878471219) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(modulate.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemChecked(index int64, checked bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&checked), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemId(index int64, id int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&id), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemAccelerator(index int64, accel Key, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_accelerator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2992817551) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&accel), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemMetadata(index int64, metadata Variant, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(metadata.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemDisabled(index int64, disabled bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&disabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemSubmenu(index int64, submenu String, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_submenu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(submenu.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemAsSeparator(index int64, enable bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_as_separator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemAsCheckable(index int64, enable bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_as_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemAsRadioCheckable(index int64, enable bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_as_radio_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemTooltip(index int64, tooltip String, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(tooltip.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemShortcut(index int64, shortcut Shortcut, global bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 825127832) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(shortcut.AsCTypePtr()), gdc.ConstTypePtr(&global), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemIndent(index int64, indent int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_indent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&indent), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemMultistate(index int64, state int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_multistate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&state), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetItemShortcutDisabled(index int64, disabled bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_shortcut_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&disabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) ToggleItemChecked(index int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("toggle_item_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) ToggleItemMultistate(index int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("toggle_item_multistate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) GetItemText(index int64, ) String {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PopupMenu) GetItemTextDirection(index int64, ) ControlTextDirection {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4235602388) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  var ret ControlTextDirection

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PopupMenu) GetItemLanguage(index int64, ) String {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PopupMenu) GetItemIcon(index int64, ) Texture2D {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PopupMenu) GetItemIconMaxWidth(index int64, ) int64 {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_icon_max_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) GetItemIconModulate(index int64, ) Color {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_icon_modulate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3457211756) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PopupMenu) IsItemChecked(index int64, ) bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) GetItemId(index int64, ) int64 {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) GetItemIndex(id int64, ) int64 {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) GetItemAccelerator(index int64, ) Key {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_accelerator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 253789942) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  var ret Key

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *PopupMenu) GetItemMetadata(index int64, ) Variant {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PopupMenu) IsItemDisabled(index int64, ) bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) GetItemSubmenu(index int64, ) String {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_submenu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PopupMenu) IsItemSeparator(index int64, ) bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_separator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) IsItemCheckable(index int64, ) bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) IsItemRadioCheckable(index int64, ) bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_radio_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) IsItemShortcutDisabled(index int64, ) bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_item_shortcut_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) GetItemTooltip(index int64, ) String {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PopupMenu) GetItemShortcut(index int64, ) Shortcut {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_shortcut")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1449483325) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewShortcut()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PopupMenu) GetItemIndent(index int64, ) int64 {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_indent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) SetFocusedItem(index int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_focused_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) GetFocusedItem() int64 {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_focused_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) SetItemCount(count int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_item_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) GetItemCount() int64 {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_item_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) ScrollToItem(index int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("scroll_to_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) RemoveItem(index int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) AddSeparator(label String, id int64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_separator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2266703459) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&id), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) Clear(free_submenus bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 107499316) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&free_submenus), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) SetHideOnItemSelection(enable bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hide_on_item_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) IsHideOnItemSelection() bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hide_on_item_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) SetHideOnCheckableItemSelection(enable bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hide_on_checkable_item_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) IsHideOnCheckableItemSelection() bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hide_on_checkable_item_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) SetHideOnStateItemSelection(enable bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_hide_on_state_item_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) IsHideOnStateItemSelection() bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_hide_on_state_item_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) SetSubmenuPopupDelay(seconds float64, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_submenu_popup_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) GetSubmenuPopupDelay() float64 {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_submenu_popup_delay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PopupMenu) SetAllowSearch(allow bool, )  {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_allow_search")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PopupMenu) GetAllowSearch() bool {
  classNameV := StringNameFromStr("PopupMenu")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_allow_search")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type PopupMenuIdPressedSignalFn func(id int, )

func (me *PopupMenu) ConnectIdPressed(subs SignalSubscribers, fn PopupMenuIdPressedSignalFn) {
  sig := StringNameFromStr("id_pressed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *PopupMenu) DisconnectIdPressed(subs SignalSubscribers, fn PopupMenuIdPressedSignalFn) {
  sig := StringNameFromStr("id_pressed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type PopupMenuIdFocusedSignalFn func(id int, )

func (me *PopupMenu) ConnectIdFocused(subs SignalSubscribers, fn PopupMenuIdFocusedSignalFn) {
  sig := StringNameFromStr("id_focused")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *PopupMenu) DisconnectIdFocused(subs SignalSubscribers, fn PopupMenuIdFocusedSignalFn) {
  sig := StringNameFromStr("id_focused")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type PopupMenuIndexPressedSignalFn func(index int, )

func (me *PopupMenu) ConnectIndexPressed(subs SignalSubscribers, fn PopupMenuIndexPressedSignalFn) {
  sig := StringNameFromStr("index_pressed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *PopupMenu) DisconnectIndexPressed(subs SignalSubscribers, fn PopupMenuIndexPressedSignalFn) {
  sig := StringNameFromStr("index_pressed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type PopupMenuMenuChangedSignalFn func()

func (me *PopupMenu) ConnectMenuChanged(subs SignalSubscribers, fn PopupMenuMenuChangedSignalFn) {
  sig := StringNameFromStr("menu_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *PopupMenu) DisconnectMenuChanged(subs SignalSubscribers, fn PopupMenuMenuChangedSignalFn) {
  sig := StringNameFromStr("menu_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
