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

type ptrsForPopupMenuList struct {
	fnActivateItemByEvent             gdc.MethodBindPtr
	fnAddItem                         gdc.MethodBindPtr
	fnAddIconItem                     gdc.MethodBindPtr
	fnAddCheckItem                    gdc.MethodBindPtr
	fnAddIconCheckItem                gdc.MethodBindPtr
	fnAddRadioCheckItem               gdc.MethodBindPtr
	fnAddIconRadioCheckItem           gdc.MethodBindPtr
	fnAddMultistateItem               gdc.MethodBindPtr
	fnAddShortcut                     gdc.MethodBindPtr
	fnAddIconShortcut                 gdc.MethodBindPtr
	fnAddCheckShortcut                gdc.MethodBindPtr
	fnAddIconCheckShortcut            gdc.MethodBindPtr
	fnAddRadioCheckShortcut           gdc.MethodBindPtr
	fnAddIconRadioCheckShortcut       gdc.MethodBindPtr
	fnAddSubmenuItem                  gdc.MethodBindPtr
	fnSetItemText                     gdc.MethodBindPtr
	fnSetItemTextDirection            gdc.MethodBindPtr
	fnSetItemLanguage                 gdc.MethodBindPtr
	fnSetItemIcon                     gdc.MethodBindPtr
	fnSetItemIconMaxWidth             gdc.MethodBindPtr
	fnSetItemIconModulate             gdc.MethodBindPtr
	fnSetItemChecked                  gdc.MethodBindPtr
	fnSetItemId                       gdc.MethodBindPtr
	fnSetItemAccelerator              gdc.MethodBindPtr
	fnSetItemMetadata                 gdc.MethodBindPtr
	fnSetItemDisabled                 gdc.MethodBindPtr
	fnSetItemSubmenu                  gdc.MethodBindPtr
	fnSetItemAsSeparator              gdc.MethodBindPtr
	fnSetItemAsCheckable              gdc.MethodBindPtr
	fnSetItemAsRadioCheckable         gdc.MethodBindPtr
	fnSetItemTooltip                  gdc.MethodBindPtr
	fnSetItemShortcut                 gdc.MethodBindPtr
	fnSetItemIndent                   gdc.MethodBindPtr
	fnSetItemMultistate               gdc.MethodBindPtr
	fnSetItemShortcutDisabled         gdc.MethodBindPtr
	fnToggleItemChecked               gdc.MethodBindPtr
	fnToggleItemMultistate            gdc.MethodBindPtr
	fnGetItemText                     gdc.MethodBindPtr
	fnGetItemTextDirection            gdc.MethodBindPtr
	fnGetItemLanguage                 gdc.MethodBindPtr
	fnGetItemIcon                     gdc.MethodBindPtr
	fnGetItemIconMaxWidth             gdc.MethodBindPtr
	fnGetItemIconModulate             gdc.MethodBindPtr
	fnIsItemChecked                   gdc.MethodBindPtr
	fnGetItemId                       gdc.MethodBindPtr
	fnGetItemIndex                    gdc.MethodBindPtr
	fnGetItemAccelerator              gdc.MethodBindPtr
	fnGetItemMetadata                 gdc.MethodBindPtr
	fnIsItemDisabled                  gdc.MethodBindPtr
	fnGetItemSubmenu                  gdc.MethodBindPtr
	fnIsItemSeparator                 gdc.MethodBindPtr
	fnIsItemCheckable                 gdc.MethodBindPtr
	fnIsItemRadioCheckable            gdc.MethodBindPtr
	fnIsItemShortcutDisabled          gdc.MethodBindPtr
	fnGetItemTooltip                  gdc.MethodBindPtr
	fnGetItemShortcut                 gdc.MethodBindPtr
	fnGetItemIndent                   gdc.MethodBindPtr
	fnSetFocusedItem                  gdc.MethodBindPtr
	fnGetFocusedItem                  gdc.MethodBindPtr
	fnSetItemCount                    gdc.MethodBindPtr
	fnGetItemCount                    gdc.MethodBindPtr
	fnScrollToItem                    gdc.MethodBindPtr
	fnRemoveItem                      gdc.MethodBindPtr
	fnAddSeparator                    gdc.MethodBindPtr
	fnClear                           gdc.MethodBindPtr
	fnSetHideOnItemSelection          gdc.MethodBindPtr
	fnIsHideOnItemSelection           gdc.MethodBindPtr
	fnSetHideOnCheckableItemSelection gdc.MethodBindPtr
	fnIsHideOnCheckableItemSelection  gdc.MethodBindPtr
	fnSetHideOnStateItemSelection     gdc.MethodBindPtr
	fnIsHideOnStateItemSelection      gdc.MethodBindPtr
	fnSetSubmenuPopupDelay            gdc.MethodBindPtr
	fnGetSubmenuPopupDelay            gdc.MethodBindPtr
	fnSetAllowSearch                  gdc.MethodBindPtr
	fnGetAllowSearch                  gdc.MethodBindPtr
}

var ptrsForPopupMenu ptrsForPopupMenuList

func initPopupMenuPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PopupMenu")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("activate_item_by_event")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnActivateItemByEvent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3716412023))
	}
	{
		methodName := StringNameFromStr("add_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3674230041))
	}
	{
		methodName := StringNameFromStr("add_icon_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddIconItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1086190128))
	}
	{
		methodName := StringNameFromStr("add_check_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3674230041))
	}
	{
		methodName := StringNameFromStr("add_icon_check_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddIconCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1086190128))
	}
	{
		methodName := StringNameFromStr("add_radio_check_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddRadioCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3674230041))
	}
	{
		methodName := StringNameFromStr("add_icon_radio_check_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddIconRadioCheckItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1086190128))
	}
	{
		methodName := StringNameFromStr("add_multistate_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddMultistateItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 150780458))
	}
	{
		methodName := StringNameFromStr("add_shortcut")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3451850107))
	}
	{
		methodName := StringNameFromStr("add_icon_shortcut")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddIconShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2997871092))
	}
	{
		methodName := StringNameFromStr("add_check_shortcut")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddCheckShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1642193386))
	}
	{
		methodName := StringNameFromStr("add_icon_check_shortcut")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddIconCheckShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3856247530))
	}
	{
		methodName := StringNameFromStr("add_radio_check_shortcut")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddRadioCheckShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1642193386))
	}
	{
		methodName := StringNameFromStr("add_icon_radio_check_shortcut")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddIconRadioCheckShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3856247530))
	}
	{
		methodName := StringNameFromStr("add_submenu_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddSubmenuItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2979222410))
	}
	{
		methodName := StringNameFromStr("set_item_text")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("set_item_text_direction")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1707680378))
	}
	{
		methodName := StringNameFromStr("set_item_language")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("set_item_icon")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
	}
	{
		methodName := StringNameFromStr("set_item_icon_max_width")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemIconMaxWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_item_icon_modulate")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemIconModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("set_item_checked")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("set_item_id")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_item_accelerator")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemAccelerator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2992817551))
	}
	{
		methodName := StringNameFromStr("set_item_metadata")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
	}
	{
		methodName := StringNameFromStr("set_item_disabled")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("set_item_submenu")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemSubmenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("set_item_as_separator")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemAsSeparator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("set_item_as_checkable")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemAsCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("set_item_as_radio_checkable")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemAsRadioCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("set_item_tooltip")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("set_item_shortcut")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 825127832))
	}
	{
		methodName := StringNameFromStr("set_item_indent")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemIndent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_item_multistate")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemMultistate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_item_shortcut_disabled")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemShortcutDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("toggle_item_checked")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnToggleItemChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("toggle_item_multistate")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnToggleItemMultistate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_item_text")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_item_text_direction")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4235602388))
	}
	{
		methodName := StringNameFromStr("get_item_language")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_item_icon")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
	}
	{
		methodName := StringNameFromStr("get_item_icon_max_width")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemIconMaxWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_item_icon_modulate")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemIconModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("is_item_checked")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnIsItemChecked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_item_id")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_item_index")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_item_accelerator")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemAccelerator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 253789942))
	}
	{
		methodName := StringNameFromStr("get_item_metadata")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
	}
	{
		methodName := StringNameFromStr("is_item_disabled")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnIsItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_item_submenu")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemSubmenu = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("is_item_separator")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnIsItemSeparator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("is_item_checkable")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnIsItemCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("is_item_radio_checkable")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnIsItemRadioCheckable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("is_item_shortcut_disabled")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnIsItemShortcutDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_item_tooltip")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("get_item_shortcut")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemShortcut = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1449483325))
	}
	{
		methodName := StringNameFromStr("get_item_indent")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemIndent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_focused_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetFocusedItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_focused_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetFocusedItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_item_count")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_item_count")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("scroll_to_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnScrollToItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("remove_item")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnRemoveItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("add_separator")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnAddSeparator = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2266703459))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 107499316))
	}
	{
		methodName := StringNameFromStr("set_hide_on_item_selection")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetHideOnItemSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hide_on_item_selection")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnIsHideOnItemSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hide_on_checkable_item_selection")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetHideOnCheckableItemSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hide_on_checkable_item_selection")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnIsHideOnCheckableItemSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hide_on_state_item_selection")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetHideOnStateItemSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hide_on_state_item_selection")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnIsHideOnStateItemSelection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_submenu_popup_delay")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetSubmenuPopupDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_submenu_popup_delay")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetSubmenuPopupDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_allow_search")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnSetAllowSearch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_allow_search")
		defer methodName.Destroy()
		ptrsForPopupMenu.fnGetAllowSearch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

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

func (me *PopupMenu) ActivateItemByEvent(event InputEvent, for_global_only bool) bool {
	cargs := []gdc.ConstTypePtr{event.AsCTypePtr(), gdc.ConstTypePtr(&for_global_only)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&for_global_only)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnActivateItemByEvent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) AddItem(label String, id int64, accel Key) {
	cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddIconItem(texture Texture2D, label String, id int64, accel Key) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), label.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddIconItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddCheckItem(label String, id int64, accel Key) {
	cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddCheckItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddIconCheckItem(texture Texture2D, label String, id int64, accel Key) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), label.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddIconCheckItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddRadioCheckItem(label String, id int64, accel Key) {
	cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddRadioCheckItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddIconRadioCheckItem(texture Texture2D, label String, id int64, accel Key) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), label.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddIconRadioCheckItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddMultistateItem(label String, max_states int64, default_state int64, id int64, accel Key) {
	cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), gdc.ConstTypePtr(&max_states), gdc.ConstTypePtr(&default_state), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&accel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddMultistateItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddShortcut(shortcut Shortcut, id int64, global bool, allow_echo bool) {
	cargs := []gdc.ConstTypePtr{shortcut.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global), gdc.ConstTypePtr(&allow_echo)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddShortcut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddIconShortcut(texture Texture2D, shortcut Shortcut, id int64, global bool, allow_echo bool) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), shortcut.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global), gdc.ConstTypePtr(&allow_echo)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddIconShortcut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddCheckShortcut(shortcut Shortcut, id int64, global bool) {
	cargs := []gdc.ConstTypePtr{shortcut.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddCheckShortcut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddIconCheckShortcut(texture Texture2D, shortcut Shortcut, id int64, global bool) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), shortcut.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddIconCheckShortcut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddRadioCheckShortcut(shortcut Shortcut, id int64, global bool) {
	cargs := []gdc.ConstTypePtr{shortcut.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddRadioCheckShortcut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddIconRadioCheckShortcut(texture Texture2D, shortcut Shortcut, id int64, global bool) {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), shortcut.AsCTypePtr(), gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&global)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddIconRadioCheckShortcut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddSubmenuItem(label String, submenu String, id int64) {
	cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), submenu.AsCTypePtr(), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddSubmenuItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemText(index int64, text String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemTextDirection(index int64, direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemLanguage(index int64, language String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemIcon(index int64, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemIconMaxWidth(index int64, width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemIconMaxWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemIconModulate(index int64, modulate Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemIconModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemChecked(index int64, checked bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&checked)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemChecked), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemId(index int64, id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemId), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemAccelerator(index int64, accel Key) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&accel)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemAccelerator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemMetadata(index int64, metadata Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), metadata.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemMetadata), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemDisabled(index int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemSubmenu(index int64, submenu String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), submenu.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemSubmenu), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemAsSeparator(index int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemAsSeparator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemAsCheckable(index int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemAsCheckable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemAsRadioCheckable(index int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemAsRadioCheckable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemTooltip(index int64, tooltip String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), tooltip.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemTooltip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemShortcut(index int64, shortcut Shortcut, global bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), shortcut.AsCTypePtr(), gdc.ConstTypePtr(&global)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemShortcut), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemIndent(index int64, indent int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&indent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemIndent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemMultistate(index int64, state int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&state)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemMultistate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetItemShortcutDisabled(index int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemShortcutDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) ToggleItemChecked(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnToggleItemChecked), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) ToggleItemMultistate(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnToggleItemMultistate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) GetItemText(index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PopupMenu) GetItemTextDirection(index int64) ControlTextDirection {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PopupMenu) GetItemLanguage(index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PopupMenu) GetItemIcon(index int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PopupMenu) GetItemIconMaxWidth(index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemIconMaxWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) GetItemIconModulate(index int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemIconModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PopupMenu) IsItemChecked(index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnIsItemChecked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) GetItemId(index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) GetItemIndex(id int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&id)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) GetItemAccelerator(index int64) Key {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Key
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemAccelerator), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *PopupMenu) GetItemMetadata(index int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemMetadata), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PopupMenu) IsItemDisabled(index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnIsItemDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) GetItemSubmenu(index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemSubmenu), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PopupMenu) IsItemSeparator(index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnIsItemSeparator), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) IsItemCheckable(index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnIsItemCheckable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) IsItemRadioCheckable(index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnIsItemRadioCheckable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) IsItemShortcutDisabled(index int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnIsItemShortcutDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) GetItemTooltip(index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemTooltip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PopupMenu) GetItemShortcut(index int64) Shortcut {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewShortcut()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemShortcut), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PopupMenu) GetItemIndent(index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemIndent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) SetFocusedItem(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetFocusedItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) GetFocusedItem() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetFocusedItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) SetItemCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetItemCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) GetItemCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetItemCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) ScrollToItem(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnScrollToItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) RemoveItem(index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnRemoveItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) AddSeparator(label String, id int64) {
	cargs := []gdc.ConstTypePtr{label.AsCTypePtr(), gdc.ConstTypePtr(&id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnAddSeparator), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) Clear(free_submenus bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&free_submenus)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) SetHideOnItemSelection(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetHideOnItemSelection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) IsHideOnItemSelection() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnIsHideOnItemSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) SetHideOnCheckableItemSelection(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetHideOnCheckableItemSelection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) IsHideOnCheckableItemSelection() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnIsHideOnCheckableItemSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) SetHideOnStateItemSelection(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetHideOnStateItemSelection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) IsHideOnStateItemSelection() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnIsHideOnStateItemSelection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) SetSubmenuPopupDelay(seconds float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetSubmenuPopupDelay), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) GetSubmenuPopupDelay() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetSubmenuPopupDelay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PopupMenu) SetAllowSearch(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnSetAllowSearch), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PopupMenu) GetAllowSearch() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPopupMenu.fnGetAllowSearch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type PopupMenuIdPressedSignalFn func(id int)

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

type PopupMenuIdFocusedSignalFn func(id int)

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

type PopupMenuIndexPressedSignalFn func(index int)

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
