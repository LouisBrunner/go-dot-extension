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

type ptrsForTabContainerList struct {
	fnGetTabCount                gdc.MethodBindPtr
	fnSetCurrentTab              gdc.MethodBindPtr
	fnGetCurrentTab              gdc.MethodBindPtr
	fnGetPreviousTab             gdc.MethodBindPtr
	fnSelectPreviousAvailable    gdc.MethodBindPtr
	fnSelectNextAvailable        gdc.MethodBindPtr
	fnGetCurrentTabControl       gdc.MethodBindPtr
	fnGetTabBar                  gdc.MethodBindPtr
	fnGetTabControl              gdc.MethodBindPtr
	fnSetTabAlignment            gdc.MethodBindPtr
	fnGetTabAlignment            gdc.MethodBindPtr
	fnSetTabsPosition            gdc.MethodBindPtr
	fnGetTabsPosition            gdc.MethodBindPtr
	fnSetClipTabs                gdc.MethodBindPtr
	fnGetClipTabs                gdc.MethodBindPtr
	fnSetTabsVisible             gdc.MethodBindPtr
	fnAreTabsVisible             gdc.MethodBindPtr
	fnSetAllTabsInFront          gdc.MethodBindPtr
	fnIsAllTabsInFront           gdc.MethodBindPtr
	fnSetTabTitle                gdc.MethodBindPtr
	fnGetTabTitle                gdc.MethodBindPtr
	fnSetTabIcon                 gdc.MethodBindPtr
	fnGetTabIcon                 gdc.MethodBindPtr
	fnSetTabDisabled             gdc.MethodBindPtr
	fnIsTabDisabled              gdc.MethodBindPtr
	fnSetTabHidden               gdc.MethodBindPtr
	fnIsTabHidden                gdc.MethodBindPtr
	fnSetTabMetadata             gdc.MethodBindPtr
	fnGetTabMetadata             gdc.MethodBindPtr
	fnSetTabButtonIcon           gdc.MethodBindPtr
	fnGetTabButtonIcon           gdc.MethodBindPtr
	fnGetTabIdxAtPoint           gdc.MethodBindPtr
	fnGetTabIdxFromControl       gdc.MethodBindPtr
	fnSetPopup                   gdc.MethodBindPtr
	fnGetPopup                   gdc.MethodBindPtr
	fnSetDragToRearrangeEnabled  gdc.MethodBindPtr
	fnGetDragToRearrangeEnabled  gdc.MethodBindPtr
	fnSetTabsRearrangeGroup      gdc.MethodBindPtr
	fnGetTabsRearrangeGroup      gdc.MethodBindPtr
	fnSetUseHiddenTabsForMinSize gdc.MethodBindPtr
	fnGetUseHiddenTabsForMinSize gdc.MethodBindPtr
	fnSetTabFocusMode            gdc.MethodBindPtr
	fnGetTabFocusMode            gdc.MethodBindPtr
	fnSetDeselectEnabled         gdc.MethodBindPtr
	fnGetDeselectEnabled         gdc.MethodBindPtr
}

var ptrsForTabContainer ptrsForTabContainerList

func initTabContainerPtrs(iface gdc.Interface) {

	className := StringNameFromStr("TabContainer")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_tab_count")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_current_tab")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetCurrentTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_current_tab")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetCurrentTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_previous_tab")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetPreviousTab = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("select_previous_available")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSelectPreviousAvailable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("select_next_available")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSelectNextAvailable = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("get_current_tab_control")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetCurrentTabControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2783021301))
	}
	{
		methodName := StringNameFromStr("get_tab_bar")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabBar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1865451809))
	}
	{
		methodName := StringNameFromStr("get_tab_control")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1065994134))
	}
	{
		methodName := StringNameFromStr("set_tab_alignment")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2413632353))
	}
	{
		methodName := StringNameFromStr("get_tab_alignment")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabAlignment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2178122193))
	}
	{
		methodName := StringNameFromStr("set_tabs_position")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabsPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 256673370))
	}
	{
		methodName := StringNameFromStr("get_tabs_position")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabsPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 919937023))
	}
	{
		methodName := StringNameFromStr("set_clip_tabs")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetClipTabs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_clip_tabs")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetClipTabs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tabs_visible")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("are_tabs_visible")
		defer methodName.Destroy()
		ptrsForTabContainer.fnAreTabsVisible = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_all_tabs_in_front")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetAllTabsInFront = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_all_tabs_in_front")
		defer methodName.Destroy()
		ptrsForTabContainer.fnIsAllTabsInFront = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tab_title")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_tab_title")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabTitle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_tab_icon")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
	}
	{
		methodName := StringNameFromStr("get_tab_icon")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
	}
	{
		methodName := StringNameFromStr("set_tab_disabled")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_tab_disabled")
		defer methodName.Destroy()
		ptrsForTabContainer.fnIsTabDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_tab_hidden")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("is_tab_hidden")
		defer methodName.Destroy()
		ptrsForTabContainer.fnIsTabHidden = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_tab_metadata")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2152698145))
	}
	{
		methodName := StringNameFromStr("get_tab_metadata")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabMetadata = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4227898402))
	}
	{
		methodName := StringNameFromStr("set_tab_button_icon")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabButtonIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
	}
	{
		methodName := StringNameFromStr("get_tab_button_icon")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabButtonIcon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
	}
	{
		methodName := StringNameFromStr("get_tab_idx_at_point")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabIdxAtPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3820158470))
	}
	{
		methodName := StringNameFromStr("get_tab_idx_from_control")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabIdxFromControl = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2787397975))
	}
	{
		methodName := StringNameFromStr("set_popup")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("get_popup")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetPopup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 111095082))
	}
	{
		methodName := StringNameFromStr("set_drag_to_rearrange_enabled")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetDragToRearrangeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_drag_to_rearrange_enabled")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetDragToRearrangeEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tabs_rearrange_group")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabsRearrangeGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_tabs_rearrange_group")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabsRearrangeGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_use_hidden_tabs_for_min_size")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetUseHiddenTabsForMinSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_use_hidden_tabs_for_min_size")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetUseHiddenTabsForMinSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_tab_focus_mode")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetTabFocusMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3232914922))
	}
	{
		methodName := StringNameFromStr("get_tab_focus_mode")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetTabFocusMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2132829277))
	}
	{
		methodName := StringNameFromStr("set_deselect_enabled")
		defer methodName.Destroy()
		ptrsForTabContainer.fnSetDeselectEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_deselect_enabled")
		defer methodName.Destroy()
		ptrsForTabContainer.fnGetDeselectEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}

}

type TabContainer struct {
	Container
}

func (me *TabContainer) BaseClass() string {
	return "TabContainer"
}

func NewTabContainer() *TabContainer {
	str := StringNameFromStr("TabContainer") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &TabContainer{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type TabContainerTabPosition int

const (
	TabContainerTabPositionPositionTop    TabContainerTabPosition = 0
	TabContainerTabPositionPositionBottom TabContainerTabPosition = 1
	TabContainerTabPositionPositionMax    TabContainerTabPosition = 2
)

func (me *TabContainer) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *TabContainer) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *TabContainer) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *TabContainer) GetTabCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetCurrentTab(tab_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetCurrentTab), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetCurrentTab() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetCurrentTab), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) GetPreviousTab() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetPreviousTab), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SelectPreviousAvailable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSelectPreviousAvailable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SelectNextAvailable() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSelectNextAvailable), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) GetCurrentTabControl() Control {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewControl()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetCurrentTabControl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabContainer) GetTabBar() TabBar {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTabBar()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabBar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabContainer) GetTabControl(tab_idx int64) Control {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewControl()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabControl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabContainer) SetTabAlignment(alignment TabBarAlignmentMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabAlignment), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetTabAlignment() TabBarAlignmentMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TabBarAlignmentMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabAlignment), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TabContainer) SetTabsPosition(tabs_position TabContainerTabPosition) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tabs_position)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabsPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetTabsPosition() TabContainerTabPosition {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret TabContainerTabPosition

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabsPosition), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TabContainer) SetClipTabs(clip_tabs bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_tabs)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetClipTabs), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetClipTabs() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetClipTabs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetTabsVisible(visible bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabsVisible), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) AreTabsVisible() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnAreTabsVisible), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetAllTabsInFront(is_front bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&is_front)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetAllTabsInFront), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) IsAllTabsInFront() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnIsAllTabsInFront), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetTabTitle(tab_idx int64, title String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), title.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabTitle), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetTabTitle(tab_idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabTitle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabContainer) SetTabIcon(tab_idx int64, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetTabIcon(tab_idx int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabContainer) SetTabDisabled(tab_idx int64, disabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&disabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) IsTabDisabled(tab_idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnIsTabDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetTabHidden(tab_idx int64, hidden bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&hidden)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabHidden), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) IsTabHidden(tab_idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnIsTabHidden), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetTabMetadata(tab_idx int64, metadata Variant) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), metadata.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabMetadata), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetTabMetadata(tab_idx int64) Variant {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVariant()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabMetadata), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabContainer) SetTabButtonIcon(tab_idx int64, icon Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), icon.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabButtonIcon), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetTabButtonIcon(tab_idx int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&tab_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabButtonIcon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabContainer) GetTabIdxAtPoint(point Vector2) int64 {
	cargs := []gdc.ConstTypePtr{point.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabIdxAtPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) GetTabIdxFromControl(control Control) int64 {
	cargs := []gdc.ConstTypePtr{control.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabIdxFromControl), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetPopup(popup Node) {
	cargs := []gdc.ConstTypePtr{popup.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetPopup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetPopup() Popup {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPopup()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetPopup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *TabContainer) SetDragToRearrangeEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetDragToRearrangeEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetDragToRearrangeEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetDragToRearrangeEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetTabsRearrangeGroup(group_id int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabsRearrangeGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetTabsRearrangeGroup() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabsRearrangeGroup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetUseHiddenTabsForMinSize(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetUseHiddenTabsForMinSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetUseHiddenTabsForMinSize() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetUseHiddenTabsForMinSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *TabContainer) SetTabFocusMode(focus_mode ControlFocusMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&focus_mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetTabFocusMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetTabFocusMode() ControlFocusMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret ControlFocusMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetTabFocusMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *TabContainer) SetDeselectEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnSetDeselectEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *TabContainer) GetDeselectEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForTabContainer.fnGetDeselectEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TabContainerActiveTabRearrangedSignalFn func(idx_to int)

func (me *TabContainer) ConnectActiveTabRearranged(subs SignalSubscribers, fn TabContainerActiveTabRearrangedSignalFn) {
	sig := StringNameFromStr("active_tab_rearranged")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectActiveTabRearranged(subs SignalSubscribers, fn TabContainerActiveTabRearrangedSignalFn) {
	sig := StringNameFromStr("active_tab_rearranged")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabContainerTabChangedSignalFn func(tab int)

func (me *TabContainer) ConnectTabChanged(subs SignalSubscribers, fn TabContainerTabChangedSignalFn) {
	sig := StringNameFromStr("tab_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabChanged(subs SignalSubscribers, fn TabContainerTabChangedSignalFn) {
	sig := StringNameFromStr("tab_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabContainerTabClickedSignalFn func(tab int)

func (me *TabContainer) ConnectTabClicked(subs SignalSubscribers, fn TabContainerTabClickedSignalFn) {
	sig := StringNameFromStr("tab_clicked")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabClicked(subs SignalSubscribers, fn TabContainerTabClickedSignalFn) {
	sig := StringNameFromStr("tab_clicked")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabContainerTabHoveredSignalFn func(tab int)

func (me *TabContainer) ConnectTabHovered(subs SignalSubscribers, fn TabContainerTabHoveredSignalFn) {
	sig := StringNameFromStr("tab_hovered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabHovered(subs SignalSubscribers, fn TabContainerTabHoveredSignalFn) {
	sig := StringNameFromStr("tab_hovered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabContainerTabSelectedSignalFn func(tab int)

func (me *TabContainer) ConnectTabSelected(subs SignalSubscribers, fn TabContainerTabSelectedSignalFn) {
	sig := StringNameFromStr("tab_selected")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabSelected(subs SignalSubscribers, fn TabContainerTabSelectedSignalFn) {
	sig := StringNameFromStr("tab_selected")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabContainerTabButtonPressedSignalFn func(tab int)

func (me *TabContainer) ConnectTabButtonPressed(subs SignalSubscribers, fn TabContainerTabButtonPressedSignalFn) {
	sig := StringNameFromStr("tab_button_pressed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabButtonPressed(subs SignalSubscribers, fn TabContainerTabButtonPressedSignalFn) {
	sig := StringNameFromStr("tab_button_pressed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type TabContainerPrePopupPressedSignalFn func()

func (me *TabContainer) ConnectPrePopupPressed(subs SignalSubscribers, fn TabContainerPrePopupPressedSignalFn) {
	sig := StringNameFromStr("pre_popup_pressed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectPrePopupPressed(subs SignalSubscribers, fn TabContainerPrePopupPressedSignalFn) {
	sig := StringNameFromStr("pre_popup_pressed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
