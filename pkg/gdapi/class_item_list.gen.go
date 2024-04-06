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

type ptrsForItemListList struct {
	fnAddItem                gdc.MethodBindPtr
	fnAddIconItem            gdc.MethodBindPtr
	fnSetItemText            gdc.MethodBindPtr
	fnGetItemText            gdc.MethodBindPtr
	fnSetItemIcon            gdc.MethodBindPtr
	fnGetItemIcon            gdc.MethodBindPtr
	fnSetItemTextDirection   gdc.MethodBindPtr
	fnGetItemTextDirection   gdc.MethodBindPtr
	fnSetItemLanguage        gdc.MethodBindPtr
	fnGetItemLanguage        gdc.MethodBindPtr
	fnSetItemIconTransposed  gdc.MethodBindPtr
	fnIsItemIconTransposed   gdc.MethodBindPtr
	fnSetItemIconRegion      gdc.MethodBindPtr
	fnGetItemIconRegion      gdc.MethodBindPtr
	fnSetItemIconModulate    gdc.MethodBindPtr
	fnGetItemIconModulate    gdc.MethodBindPtr
	fnSetItemSelectable      gdc.MethodBindPtr
	fnIsItemSelectable       gdc.MethodBindPtr
	fnSetItemDisabled        gdc.MethodBindPtr
	fnIsItemDisabled         gdc.MethodBindPtr
	fnSetItemMetadata        gdc.MethodBindPtr
	fnGetItemMetadata        gdc.MethodBindPtr
	fnSetItemCustomBgColor   gdc.MethodBindPtr
	fnGetItemCustomBgColor   gdc.MethodBindPtr
	fnSetItemCustomFgColor   gdc.MethodBindPtr
	fnGetItemCustomFgColor   gdc.MethodBindPtr
	fnGetItemRect            gdc.MethodBindPtr
	fnSetItemTooltipEnabled  gdc.MethodBindPtr
	fnIsItemTooltipEnabled   gdc.MethodBindPtr
	fnSetItemTooltip         gdc.MethodBindPtr
	fnGetItemTooltip         gdc.MethodBindPtr
	fnSelect                 gdc.MethodBindPtr
	fnDeselect               gdc.MethodBindPtr
	fnDeselectAll            gdc.MethodBindPtr
	fnIsSelected             gdc.MethodBindPtr
	fnGetSelectedItems       gdc.MethodBindPtr
	fnMoveItem               gdc.MethodBindPtr
	fnSetItemCount           gdc.MethodBindPtr
	fnGetItemCount           gdc.MethodBindPtr
	fnRemoveItem             gdc.MethodBindPtr
	fnClear                  gdc.MethodBindPtr
	fnSortItemsByText        gdc.MethodBindPtr
	fnSetFixedColumnWidth    gdc.MethodBindPtr
	fnGetFixedColumnWidth    gdc.MethodBindPtr
	fnSetSameColumnWidth     gdc.MethodBindPtr
	fnIsSameColumnWidth      gdc.MethodBindPtr
	fnSetMaxTextLines        gdc.MethodBindPtr
	fnGetMaxTextLines        gdc.MethodBindPtr
	fnSetMaxColumns          gdc.MethodBindPtr
	fnGetMaxColumns          gdc.MethodBindPtr
	fnSetSelectMode          gdc.MethodBindPtr
	fnGetSelectMode          gdc.MethodBindPtr
	fnSetIconMode            gdc.MethodBindPtr
	fnGetIconMode            gdc.MethodBindPtr
	fnSetFixedIconSize       gdc.MethodBindPtr
	fnGetFixedIconSize       gdc.MethodBindPtr
	fnSetIconScale           gdc.MethodBindPtr
	fnGetIconScale           gdc.MethodBindPtr
	fnSetAllowRmbSelect      gdc.MethodBindPtr
	fnGetAllowRmbSelect      gdc.MethodBindPtr
	fnSetAllowReselect       gdc.MethodBindPtr
	fnGetAllowReselect       gdc.MethodBindPtr
	fnSetAllowSearch         gdc.MethodBindPtr
	fnGetAllowSearch         gdc.MethodBindPtr
	fnSetAutoHeight          gdc.MethodBindPtr
	fnHasAutoHeight          gdc.MethodBindPtr
	fnIsAnythingSelected     gdc.MethodBindPtr
	fnGetItemAtPosition      gdc.MethodBindPtr
	fnEnsureCurrentIsVisible gdc.MethodBindPtr
	fnGetVScrollBar          gdc.MethodBindPtr
	fnSetTextOverrunBehavior gdc.MethodBindPtr
	fnGetTextOverrunBehavior gdc.MethodBindPtr
	fnForceUpdateListSize    gdc.MethodBindPtr
}

var ptrsForItemList ptrsForItemListList

func initItemListPtrs(iface gdc.Interface) {

	className := StringNameFromStr("ItemList")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_item")
		defer methodName.Destroy()
		ptrsForItemList.fnAddItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 359861678))
	}
	{
		methodName := StringNameFromStr("add_icon_item")
		defer methodName.Destroy()
		ptrsForItemList.fnAddIconItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4256579627))
	}
	{
		methodName := StringNameFromStr("set_item_text")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_item_text")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_item_icon")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
	}
	{
		methodName := StringNameFromStr("get_item_icon")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
	}
	{
		methodName := StringNameFromStr("set_item_text_direction")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1707680378))
	}
	{
		methodName := StringNameFromStr("get_item_text_direction")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4235602388))
	}
	{
		methodName := StringNameFromStr("set_item_language")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_item_language")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_item_icon_transposed")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemIconTransposed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_item_icon_transposed")
		defer methodName.Destroy()
		ptrsForItemList.fnIsItemIconTransposed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_item_icon_region")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemIconRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1356297692))
	}
	{
		methodName := StringNameFromStr("get_item_icon_region")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemIconRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3327874267))
	}
	{
		methodName := StringNameFromStr("set_item_icon_modulate")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemIconModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_item_icon_modulate")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemIconModulate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("set_item_selectable")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemSelectable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_item_selectable")
		defer methodName.Destroy()
		ptrsForItemList.fnIsItemSelectable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_item_disabled")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_item_disabled")
		defer methodName.Destroy()
		ptrsForItemList.fnIsItemDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_item_metadata")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
	}
	{
		methodName := StringNameFromStr("get_item_metadata")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
	}
	{
		methodName := StringNameFromStr("set_item_custom_bg_color")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemCustomBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_item_custom_bg_color")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemCustomBgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("set_item_custom_fg_color")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemCustomFgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("get_item_custom_fg_color")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemCustomFgColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3457211756))
	}
	{
		methodName := StringNameFromStr("get_item_rect")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 159227807))
	}
	{
		methodName := StringNameFromStr("set_item_tooltip_enabled")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemTooltipEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_item_tooltip_enabled")
		defer methodName.Destroy()
		ptrsForItemList.fnIsItemTooltipEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_item_tooltip")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_item_tooltip")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemTooltip = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("select")
		defer methodName.Destroy()
		ptrsForItemList.fnSelect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 972357352))
	}
	{
		methodName := StringNameFromStr("deselect")
		defer methodName.Destroy()
		ptrsForItemList.fnDeselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("deselect_all")
		defer methodName.Destroy()
		ptrsForItemList.fnDeselectAll = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_selected")
		defer methodName.Destroy()
		ptrsForItemList.fnIsSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("get_selected_items")
		defer methodName.Destroy()
		ptrsForItemList.fnGetSelectedItems = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
	}
	{
		methodName := StringNameFromStr("move_item")
		defer methodName.Destroy()
		ptrsForItemList.fnMoveItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_item_count")
		defer methodName.Destroy()
		ptrsForItemList.fnSetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_item_count")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("remove_item")
		defer methodName.Destroy()
		ptrsForItemList.fnRemoveItem = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("clear")
		defer methodName.Destroy()
		ptrsForItemList.fnClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("sort_items_by_text")
		defer methodName.Destroy()
		ptrsForItemList.fnSortItemsByText = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_fixed_column_width")
		defer methodName.Destroy()
		ptrsForItemList.fnSetFixedColumnWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_fixed_column_width")
		defer methodName.Destroy()
		ptrsForItemList.fnGetFixedColumnWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_same_column_width")
		defer methodName.Destroy()
		ptrsForItemList.fnSetSameColumnWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_same_column_width")
		defer methodName.Destroy()
		ptrsForItemList.fnIsSameColumnWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_max_text_lines")
		defer methodName.Destroy()
		ptrsForItemList.fnSetMaxTextLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_text_lines")
		defer methodName.Destroy()
		ptrsForItemList.fnGetMaxTextLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_max_columns")
		defer methodName.Destroy()
		ptrsForItemList.fnSetMaxColumns = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_columns")
		defer methodName.Destroy()
		ptrsForItemList.fnGetMaxColumns = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_select_mode")
		defer methodName.Destroy()
		ptrsForItemList.fnSetSelectMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 928267388))
	}
	{
		methodName := StringNameFromStr("get_select_mode")
		defer methodName.Destroy()
		ptrsForItemList.fnGetSelectMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1191945842))
	}
	{
		methodName := StringNameFromStr("set_icon_mode")
		defer methodName.Destroy()
		ptrsForItemList.fnSetIconMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2025053633))
	}
	{
		methodName := StringNameFromStr("get_icon_mode")
		defer methodName.Destroy()
		ptrsForItemList.fnGetIconMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3353929232))
	}
	{
		methodName := StringNameFromStr("set_fixed_icon_size")
		defer methodName.Destroy()
		ptrsForItemList.fnSetFixedIconSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
	}
	{
		methodName := StringNameFromStr("get_fixed_icon_size")
		defer methodName.Destroy()
		ptrsForItemList.fnGetFixedIconSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3690982128))
	}
	{
		methodName := StringNameFromStr("set_icon_scale")
		defer methodName.Destroy()
		ptrsForItemList.fnSetIconScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_icon_scale")
		defer methodName.Destroy()
		ptrsForItemList.fnGetIconScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_allow_rmb_select")
		defer methodName.Destroy()
		ptrsForItemList.fnSetAllowRmbSelect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_allow_rmb_select")
		defer methodName.Destroy()
		ptrsForItemList.fnGetAllowRmbSelect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_allow_reselect")
		defer methodName.Destroy()
		ptrsForItemList.fnSetAllowReselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_allow_reselect")
		defer methodName.Destroy()
		ptrsForItemList.fnGetAllowReselect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_allow_search")
		defer methodName.Destroy()
		ptrsForItemList.fnSetAllowSearch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_allow_search")
		defer methodName.Destroy()
		ptrsForItemList.fnGetAllowSearch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_auto_height")
		defer methodName.Destroy()
		ptrsForItemList.fnSetAutoHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("has_auto_height")
		defer methodName.Destroy()
		ptrsForItemList.fnHasAutoHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_anything_selected")
		defer methodName.Destroy()
		ptrsForItemList.fnIsAnythingSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_item_at_position")
		defer methodName.Destroy()
		ptrsForItemList.fnGetItemAtPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2300324924))
	}
	{
		methodName := StringNameFromStr("ensure_current_is_visible")
		defer methodName.Destroy()
		ptrsForItemList.fnEnsureCurrentIsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_v_scroll_bar")
		defer methodName.Destroy()
		ptrsForItemList.fnGetVScrollBar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2630340773))
	}
	{
		methodName := StringNameFromStr("set_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForItemList.fnSetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1008890932))
	}
	{
		methodName := StringNameFromStr("get_text_overrun_behavior")
		defer methodName.Destroy()
		ptrsForItemList.fnGetTextOverrunBehavior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3779142101))
	}
	{
		methodName := StringNameFromStr("force_update_list_size")
		defer methodName.Destroy()
		ptrsForItemList.fnForceUpdateListSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

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
	ItemListIconModeIconModeTop  ItemListIconMode = 0
	ItemListIconModeIconModeLeft ItemListIconMode = 1
)

type ItemListSelectMode int

const (
	ItemListSelectModeSelectSingle ItemListSelectMode = 0
	ItemListSelectModeSelectMulti  ItemListSelectMode = 1
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

func (me *ItemList) AddItem(text String, icon Texture2D, selectable bool) int64 {
	cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), icon.AsCTypePtr(), gdc.ConstTypePtr(&selectable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&selectable)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnAddItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) AddIconItem(icon Texture2D, selectable bool) int64 {
	cargs := []gdc.ConstTypePtr{icon.AsCTypePtr(), gdc.ConstTypePtr(&selectable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&selectable)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnAddIconItem), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetItemText(idx int64, text String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), text.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemText(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemText), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetItemIcon(idx int64, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemIcon(idx int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetItemTextDirection(idx int64, direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemTextDirection(idx int64) ControlTextDirection {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ItemList) SetItemLanguage(idx int64, language String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemLanguage(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetItemIconTransposed(idx int64, transposed bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&transposed)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemIconTransposed), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) IsItemIconTransposed(idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnIsItemIconTransposed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetItemIconRegion(idx int64, rect Rect2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemIconRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemIconRegion(idx int64) Rect2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemIconRegion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetItemIconModulate(idx int64, modulate Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), modulate.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemIconModulate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemIconModulate(idx int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemIconModulate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetItemSelectable(idx int64, selectable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&selectable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemSelectable), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) IsItemSelectable(idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnIsItemSelectable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetItemDisabled(idx int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) IsItemDisabled(idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnIsItemDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetItemMetadata(idx int64, metadata Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), metadata.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemMetadata), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemMetadata(idx int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemMetadata), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetItemCustomBgColor(idx int64, custom_bg_color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), custom_bg_color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemCustomBgColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemCustomBgColor(idx int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemCustomBgColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetItemCustomFgColor(idx int64, custom_fg_color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), custom_fg_color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemCustomFgColor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemCustomFgColor(idx int64) Color {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewColor()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemCustomFgColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) GetItemRect(idx int64, expand bool) Rect2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&expand)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()
	pinner.Pin(&idx)
	pinner.Pin(&expand)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetItemTooltipEnabled(idx int64, enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemTooltipEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) IsItemTooltipEnabled(idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnIsItemTooltipEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetItemTooltip(idx int64, tooltip String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), tooltip.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemTooltip), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemTooltip(idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemTooltip), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) Select(idx int64, single bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&single)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSelect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) Deselect(idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnDeselect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) DeselectAll() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnDeselectAll), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) IsSelected(idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnIsSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) GetSelectedItems() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetSelectedItems), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) MoveItem(from_idx int64, to_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_idx), gdc.ConstTypePtr(&to_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnMoveItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) SetItemCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetItemCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetItemCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) RemoveItem(idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnRemoveItem), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) Clear() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnClear), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) SortItemsByText() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSortItemsByText), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) SetFixedColumnWidth(width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetFixedColumnWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetFixedColumnWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetFixedColumnWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetSameColumnWidth(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetSameColumnWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) IsSameColumnWidth() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnIsSameColumnWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetMaxTextLines(lines int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lines)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetMaxTextLines), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetMaxTextLines() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetMaxTextLines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetMaxColumns(amount int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetMaxColumns), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetMaxColumns() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetMaxColumns), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetSelectMode(mode ItemListSelectMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetSelectMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetSelectMode() ItemListSelectMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ItemListSelectMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetSelectMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ItemList) SetIconMode(mode ItemListIconMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetIconMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetIconMode() ItemListIconMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ItemListIconMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetIconMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ItemList) SetFixedIconSize(size Vector2i) {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetFixedIconSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetFixedIconSize() Vector2i {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2i()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetFixedIconSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetIconScale(scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetIconScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetIconScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetIconScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetAllowRmbSelect(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetAllowRmbSelect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetAllowRmbSelect() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetAllowRmbSelect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetAllowReselect(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetAllowReselect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetAllowReselect() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetAllowReselect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetAllowSearch(allow bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&allow)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetAllowSearch), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetAllowSearch() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetAllowSearch), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) SetAutoHeight(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetAutoHeight), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) HasAutoHeight() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnHasAutoHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) IsAnythingSelected() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnIsAnythingSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) GetItemAtPosition(position Vector2, exact bool) int64 {
	cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), gdc.ConstTypePtr(&exact)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&exact)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetItemAtPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *ItemList) EnsureCurrentIsVisible() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnEnsureCurrentIsVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetVScrollBar() VScrollBar {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVScrollBar()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetVScrollBar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *ItemList) SetTextOverrunBehavior(overrun_behavior TextServerOverrunBehavior) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&overrun_behavior)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnSetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *ItemList) GetTextOverrunBehavior() TextServerOverrunBehavior {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TextServerOverrunBehavior

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnGetTextOverrunBehavior), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *ItemList) ForceUpdateListSize() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForItemList.fnForceUpdateListSize), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type ItemListItemSelectedSignalFn func(index int)

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

type ItemListEmptyClickedSignalFn func(at_position Vector2, mouse_button_index int)

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

type ItemListItemClickedSignalFn func(index int, at_position Vector2, mouse_button_index int)

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

type ItemListMultiSelectedSignalFn func(index int, selected bool)

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

type ItemListItemActivatedSignalFn func(index int)

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
