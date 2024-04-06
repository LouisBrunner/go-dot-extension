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

type ptrsForTabBarList struct {
	fnSetTabCount               gdc.MethodBindPtr
	fnGetTabCount               gdc.MethodBindPtr
	fnSetCurrentTab             gdc.MethodBindPtr
	fnGetCurrentTab             gdc.MethodBindPtr
	fnGetPreviousTab            gdc.MethodBindPtr
	fnSelectPreviousAvailable   gdc.MethodBindPtr
	fnSelectNextAvailable       gdc.MethodBindPtr
	fnSetTabTitle               gdc.MethodBindPtr
	fnGetTabTitle               gdc.MethodBindPtr
	fnSetTabTextDirection       gdc.MethodBindPtr
	fnGetTabTextDirection       gdc.MethodBindPtr
	fnSetTabLanguage            gdc.MethodBindPtr
	fnGetTabLanguage            gdc.MethodBindPtr
	fnSetTabIcon                gdc.MethodBindPtr
	fnGetTabIcon                gdc.MethodBindPtr
	fnSetTabIconMaxWidth        gdc.MethodBindPtr
	fnGetTabIconMaxWidth        gdc.MethodBindPtr
	fnSetTabButtonIcon          gdc.MethodBindPtr
	fnGetTabButtonIcon          gdc.MethodBindPtr
	fnSetTabDisabled            gdc.MethodBindPtr
	fnIsTabDisabled             gdc.MethodBindPtr
	fnSetTabHidden              gdc.MethodBindPtr
	fnIsTabHidden               gdc.MethodBindPtr
	fnSetTabMetadata            gdc.MethodBindPtr
	fnGetTabMetadata            gdc.MethodBindPtr
	fnRemoveTab                 gdc.MethodBindPtr
	fnAddTab                    gdc.MethodBindPtr
	fnGetTabIdxAtPoint          gdc.MethodBindPtr
	fnSetTabAlignment           gdc.MethodBindPtr
	fnGetTabAlignment           gdc.MethodBindPtr
	fnSetClipTabs               gdc.MethodBindPtr
	fnGetClipTabs               gdc.MethodBindPtr
	fnGetTabOffset              gdc.MethodBindPtr
	fnGetOffsetButtonsVisible   gdc.MethodBindPtr
	fnEnsureTabVisible          gdc.MethodBindPtr
	fnGetTabRect                gdc.MethodBindPtr
	fnMoveTab                   gdc.MethodBindPtr
	fnSetTabCloseDisplayPolicy  gdc.MethodBindPtr
	fnGetTabCloseDisplayPolicy  gdc.MethodBindPtr
	fnSetMaxTabWidth            gdc.MethodBindPtr
	fnGetMaxTabWidth            gdc.MethodBindPtr
	fnSetScrollingEnabled       gdc.MethodBindPtr
	fnGetScrollingEnabled       gdc.MethodBindPtr
	fnSetDragToRearrangeEnabled gdc.MethodBindPtr
	fnGetDragToRearrangeEnabled gdc.MethodBindPtr
	fnSetTabsRearrangeGroup     gdc.MethodBindPtr
	fnGetTabsRearrangeGroup     gdc.MethodBindPtr
	fnSetScrollToSelected       gdc.MethodBindPtr
	fnGetScrollToSelected       gdc.MethodBindPtr
	fnSetSelectWithRmb          gdc.MethodBindPtr
	fnGetSelectWithRmb          gdc.MethodBindPtr
	fnClearTabs                 gdc.MethodBindPtr
}

var ptrsForTabBar ptrsForTabBarList

func initTabBarPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TabBar")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_tab_count")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_tab_count")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_current_tab")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetCurrentTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_current_tab")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetCurrentTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_previous_tab")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetPreviousTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("select_previous_available")
		defer methodName.Destroy()
		ptrsForTabBar.fnSelectPreviousAvailable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("select_next_available")
		defer methodName.Destroy()
		ptrsForTabBar.fnSelectNextAvailable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_tab_title")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_tab_title")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_tab_text_direction")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1707680378))
	}
	{
		methodName := StringNameFromStr("get_tab_text_direction")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabTextDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4235602388))
	}
	{
		methodName := StringNameFromStr("set_tab_language")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_tab_language")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabLanguage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_tab_icon")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
	}
	{
		methodName := StringNameFromStr("get_tab_icon")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
	}
	{
		methodName := StringNameFromStr("set_tab_icon_max_width")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabIconMaxWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_tab_icon_max_width")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabIconMaxWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_tab_button_icon")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabButtonIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
	}
	{
		methodName := StringNameFromStr("get_tab_button_icon")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabButtonIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
	}
	{
		methodName := StringNameFromStr("set_tab_disabled")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_tab_disabled")
		defer methodName.Destroy()
		ptrsForTabBar.fnIsTabDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_tab_hidden")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_tab_hidden")
		defer methodName.Destroy()
		ptrsForTabBar.fnIsTabHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_tab_metadata")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
	}
	{
		methodName := StringNameFromStr("get_tab_metadata")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
	}
	{
		methodName := StringNameFromStr("remove_tab")
		defer methodName.Destroy()
		ptrsForTabBar.fnRemoveTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("add_tab")
		defer methodName.Destroy()
		ptrsForTabBar.fnAddTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1465444425))
	}
	{
		methodName := StringNameFromStr("get_tab_idx_at_point")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabIdxAtPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3820158470))
	}
	{
		methodName := StringNameFromStr("set_tab_alignment")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2413632353))
	}
	{
		methodName := StringNameFromStr("get_tab_alignment")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2178122193))
	}
	{
		methodName := StringNameFromStr("set_clip_tabs")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetClipTabs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_clip_tabs")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetClipTabs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_tab_offset")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_offset_buttons_visible")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetOffsetButtonsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("ensure_tab_visible")
		defer methodName.Destroy()
		ptrsForTabBar.fnEnsureTabVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_tab_rect")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabRect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3327874267))
	}
	{
		methodName := StringNameFromStr("move_tab")
		defer methodName.Destroy()
		ptrsForTabBar.fnMoveTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("set_tab_close_display_policy")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabCloseDisplayPolicy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2212906737))
	}
	{
		methodName := StringNameFromStr("get_tab_close_display_policy")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabCloseDisplayPolicy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2956568028))
	}
	{
		methodName := StringNameFromStr("set_max_tab_width")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetMaxTabWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_tab_width")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetMaxTabWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_scrolling_enabled")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetScrollingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_scrolling_enabled")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetScrollingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_drag_to_rearrange_enabled")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetDragToRearrangeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_drag_to_rearrange_enabled")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetDragToRearrangeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tabs_rearrange_group")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetTabsRearrangeGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_tabs_rearrange_group")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetTabsRearrangeGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_scroll_to_selected")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetScrollToSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_scroll_to_selected")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetScrollToSelected = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_select_with_rmb")
		defer methodName.Destroy()
		ptrsForTabBar.fnSetSelectWithRmb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_select_with_rmb")
		defer methodName.Destroy()
		ptrsForTabBar.fnGetSelectWithRmb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("clear_tabs")
		defer methodName.Destroy()
		ptrsForTabBar.fnClearTabs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
}

type TabBar struct {
	Control
}

func (me *TabBar) BaseClass() string {
	return "TabBar"
}

func NewTabBar() *TabBar {
	str := StringNameFromStr("TabBar") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TabBar{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TabBarAlignmentMode int

const (
	TabBarAlignmentModeAlignmentLeft   TabBarAlignmentMode = 0
	TabBarAlignmentModeAlignmentCenter TabBarAlignmentMode = 1
	TabBarAlignmentModeAlignmentRight  TabBarAlignmentMode = 2
	TabBarAlignmentModeAlignmentMax    TabBarAlignmentMode = 3
)

type TabBarCloseButtonDisplayPolicy int

const (
	TabBarCloseButtonDisplayPolicyCloseButtonShowNever      TabBarCloseButtonDisplayPolicy = 0
	TabBarCloseButtonDisplayPolicyCloseButtonShowActiveOnly TabBarCloseButtonDisplayPolicy = 1
	TabBarCloseButtonDisplayPolicyCloseButtonShowAlways     TabBarCloseButtonDisplayPolicy = 2
	TabBarCloseButtonDisplayPolicyCloseButtonMax            TabBarCloseButtonDisplayPolicy = 3
)

func (me *TabBar) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TabBar) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TabBar) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TabBar) SetTabCount(count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabCount), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetCurrentTab(tab_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetCurrentTab), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetCurrentTab() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetCurrentTab), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) GetPreviousTab() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetPreviousTab), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SelectPreviousAvailable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSelectPreviousAvailable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SelectNextAvailable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSelectNextAvailable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetTabTitle(tab_idx int64, title String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabTitle(tab_idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabBar) SetTabTextDirection(tab_idx int64, direction ControlTextDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabTextDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabTextDirection(tab_idx int64) ControlTextDirection {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlTextDirection
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabTextDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TabBar) SetTabLanguage(tab_idx int64, language String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), language.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabLanguage), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabLanguage(tab_idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabLanguage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabBar) SetTabIcon(tab_idx int64, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabIcon(tab_idx int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabBar) SetTabIconMaxWidth(tab_idx int64, width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabIconMaxWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabIconMaxWidth(tab_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabIconMaxWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetTabButtonIcon(tab_idx int64, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabButtonIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabButtonIcon(tab_idx int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabButtonIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabBar) SetTabDisabled(tab_idx int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) IsTabDisabled(tab_idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnIsTabDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetTabHidden(tab_idx int64, hidden bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabHidden), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) IsTabHidden(tab_idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnIsTabHidden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetTabMetadata(tab_idx int64, metadata Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), metadata.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabMetadata), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabMetadata(tab_idx int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabMetadata), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabBar) RemoveTab(tab_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnRemoveTab), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) AddTab(title String, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{title.AsCTypePtr(), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnAddTab), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabIdxAtPoint(point Vector2) int64 {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabIdxAtPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetTabAlignment(alignment TabBarAlignmentMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabAlignment() TabBarAlignmentMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TabBarAlignmentMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TabBar) SetClipTabs(clip_tabs bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_tabs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetClipTabs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetClipTabs() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetClipTabs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) GetTabOffset() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) GetOffsetButtonsVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetOffsetButtonsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) EnsureTabVisible(idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnEnsureTabVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabRect(tab_idx int64) Rect2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRect2()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabRect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabBar) MoveTab(from int64, to int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnMoveTab), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) SetTabCloseDisplayPolicy(policy TabBarCloseButtonDisplayPolicy) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&policy)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabCloseDisplayPolicy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabCloseDisplayPolicy() TabBarCloseButtonDisplayPolicy {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TabBarCloseButtonDisplayPolicy

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabCloseDisplayPolicy), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TabBar) SetMaxTabWidth(width int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetMaxTabWidth), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetMaxTabWidth() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetMaxTabWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetScrollingEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetScrollingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetScrollingEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetScrollingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetDragToRearrangeEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetDragToRearrangeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetDragToRearrangeEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetDragToRearrangeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetTabsRearrangeGroup(group_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetTabsRearrangeGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetTabsRearrangeGroup() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetTabsRearrangeGroup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetScrollToSelected(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetScrollToSelected), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetScrollToSelected() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetScrollToSelected), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) SetSelectWithRmb(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnSetSelectWithRmb), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabBar) GetSelectWithRmb() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnGetSelectWithRmb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabBar) ClearTabs() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabBar.fnClearTabs), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TabBarTabSelectedSignalFn func(tab int)

func (me *TabBar) ConnectTabSelected(subs SignalSubscribers, fn TabBarTabSelectedSignalFn) {
	sig := StringNameFromStr("tab_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabBar) DisconnectTabSelected(subs SignalSubscribers, fn TabBarTabSelectedSignalFn) {
	sig := StringNameFromStr("tab_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabBarTabChangedSignalFn func(tab int)

func (me *TabBar) ConnectTabChanged(subs SignalSubscribers, fn TabBarTabChangedSignalFn) {
	sig := StringNameFromStr("tab_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabBar) DisconnectTabChanged(subs SignalSubscribers, fn TabBarTabChangedSignalFn) {
	sig := StringNameFromStr("tab_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabBarTabClickedSignalFn func(tab int)

func (me *TabBar) ConnectTabClicked(subs SignalSubscribers, fn TabBarTabClickedSignalFn) {
	sig := StringNameFromStr("tab_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabBar) DisconnectTabClicked(subs SignalSubscribers, fn TabBarTabClickedSignalFn) {
	sig := StringNameFromStr("tab_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabBarTabRmbClickedSignalFn func(tab int)

func (me *TabBar) ConnectTabRmbClicked(subs SignalSubscribers, fn TabBarTabRmbClickedSignalFn) {
	sig := StringNameFromStr("tab_rmb_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabBar) DisconnectTabRmbClicked(subs SignalSubscribers, fn TabBarTabRmbClickedSignalFn) {
	sig := StringNameFromStr("tab_rmb_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabBarTabClosePressedSignalFn func(tab int)

func (me *TabBar) ConnectTabClosePressed(subs SignalSubscribers, fn TabBarTabClosePressedSignalFn) {
	sig := StringNameFromStr("tab_close_pressed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabBar) DisconnectTabClosePressed(subs SignalSubscribers, fn TabBarTabClosePressedSignalFn) {
	sig := StringNameFromStr("tab_close_pressed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabBarTabButtonPressedSignalFn func(tab int)

func (me *TabBar) ConnectTabButtonPressed(subs SignalSubscribers, fn TabBarTabButtonPressedSignalFn) {
	sig := StringNameFromStr("tab_button_pressed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabBar) DisconnectTabButtonPressed(subs SignalSubscribers, fn TabBarTabButtonPressedSignalFn) {
	sig := StringNameFromStr("tab_button_pressed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabBarTabHoveredSignalFn func(tab int)

func (me *TabBar) ConnectTabHovered(subs SignalSubscribers, fn TabBarTabHoveredSignalFn) {
	sig := StringNameFromStr("tab_hovered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabBar) DisconnectTabHovered(subs SignalSubscribers, fn TabBarTabHoveredSignalFn) {
	sig := StringNameFromStr("tab_hovered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabBarActiveTabRearrangedSignalFn func(idx_to int)

func (me *TabBar) ConnectActiveTabRearranged(subs SignalSubscribers, fn TabBarActiveTabRearrangedSignalFn) {
	sig := StringNameFromStr("active_tab_rearranged")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabBar) DisconnectActiveTabRearranged(subs SignalSubscribers, fn TabBarActiveTabRearrangedSignalFn) {
	sig := StringNameFromStr("active_tab_rearranged")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
