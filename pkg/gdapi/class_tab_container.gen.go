// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TabContainer struct {
  Container
}

func (me *TabContainer) BaseClass() string {
  return "TabContainer"
}



// Enums

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

func  (me *TabContainer) GetTabCount() int {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetCurrentTab(tab_idx int, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetCurrentTab() int {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) GetPreviousTab() int {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_previous_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SelectPreviousAvailable() bool {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_previous_available")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SelectNextAvailable() bool {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_next_available")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) GetCurrentTabControl() Control {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_tab_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2783021301) // FIXME: should cache?
  var ret Control
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) GetTabBar() TabBar {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_bar")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1865451809) // FIXME: should cache?
  var ret TabBar
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) GetTabControl(tab_idx int, ) Control {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1065994134) // FIXME: should cache?
  var ret Control
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabAlignment(alignment TabBarAlignmentMode, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2413632353) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetTabAlignment() TabBarAlignmentMode {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2178122193) // FIXME: should cache?
  var ret TabBarAlignmentMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetClipTabs(clip_tabs bool, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clip_tabs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_tabs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetClipTabs() bool {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clip_tabs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabsVisible(visible bool, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tabs_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&visible), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) AreTabsVisible() bool {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("are_tabs_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetAllTabsInFront(is_front bool, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_all_tabs_in_front")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&is_front), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) IsAllTabsInFront() bool {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_all_tabs_in_front")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabTitle(tab_idx int, title String, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(title.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetTabTitle(tab_idx int, ) String {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabIcon(tab_idx int, icon Texture2D, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(icon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetTabIcon(tab_idx int, ) Texture2D {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabDisabled(tab_idx int, disabled bool, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) IsTabDisabled(tab_idx int, ) bool {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_tab_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabHidden(tab_idx int, hidden bool, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&hidden), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) IsTabHidden(tab_idx int, ) bool {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_tab_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabMetadata(tab_idx int, metadata Variant, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(metadata.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetTabMetadata(tab_idx int, ) Variant {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabButtonIcon(tab_idx int, icon Texture2D, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_button_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(icon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetTabButtonIcon(tab_idx int, ) Texture2D {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_button_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) GetTabIdxAtPoint(point Vector2, ) int {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_idx_at_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3820158470) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(point.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) GetTabIdxFromControl(control Control, ) int {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_idx_from_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2787397975) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(control.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetPopup(popup Node, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1078189570) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(popup.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetPopup() Popup {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 111095082) // FIXME: should cache?
  var ret Popup
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetDragToRearrangeEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_to_rearrange_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetDragToRearrangeEnabled() bool {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drag_to_rearrange_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabsRearrangeGroup(group_id int, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tabs_rearrange_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetTabsRearrangeGroup() int {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tabs_rearrange_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetUseHiddenTabsForMinSize(enabled bool, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_hidden_tabs_for_min_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetUseHiddenTabsForMinSize() bool {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_hidden_tabs_for_min_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TabContainer) SetTabFocusMode(focus_mode ControlFocusMode, )  {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_focus_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3232914922) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&focus_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TabContainer) GetTabFocusMode() ControlFocusMode {
  classNameV := StringNameFromStr("TabContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_focus_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2132829277) // FIXME: should cache?
  var ret ControlFocusMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TabContainerActiveTabRearrangedSignalFn func(idx_to int, )

func (me *TabContainer) ConnectActiveTabRearranged(subs SignalSubscribers, fn TabContainerActiveTabRearrangedSignalFn) {
  sig := StringNameFromStr("active_tab_rearranged")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectActiveTabRearranged(subs SignalSubscribers, fn TabContainerActiveTabRearrangedSignalFn) {
  sig := StringNameFromStr("active_tab_rearranged")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TabContainerTabChangedSignalFn func(tab int, )

func (me *TabContainer) ConnectTabChanged(subs SignalSubscribers, fn TabContainerTabChangedSignalFn) {
  sig := StringNameFromStr("tab_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabChanged(subs SignalSubscribers, fn TabContainerTabChangedSignalFn) {
  sig := StringNameFromStr("tab_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TabContainerTabClickedSignalFn func(tab int, )

func (me *TabContainer) ConnectTabClicked(subs SignalSubscribers, fn TabContainerTabClickedSignalFn) {
  sig := StringNameFromStr("tab_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabClicked(subs SignalSubscribers, fn TabContainerTabClickedSignalFn) {
  sig := StringNameFromStr("tab_clicked")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TabContainerTabHoveredSignalFn func(tab int, )

func (me *TabContainer) ConnectTabHovered(subs SignalSubscribers, fn TabContainerTabHoveredSignalFn) {
  sig := StringNameFromStr("tab_hovered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabHovered(subs SignalSubscribers, fn TabContainerTabHoveredSignalFn) {
  sig := StringNameFromStr("tab_hovered")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TabContainerTabSelectedSignalFn func(tab int, )

func (me *TabContainer) ConnectTabSelected(subs SignalSubscribers, fn TabContainerTabSelectedSignalFn) {
  sig := StringNameFromStr("tab_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabSelected(subs SignalSubscribers, fn TabContainerTabSelectedSignalFn) {
  sig := StringNameFromStr("tab_selected")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TabContainerTabButtonPressedSignalFn func(tab int, )

func (me *TabContainer) ConnectTabButtonPressed(subs SignalSubscribers, fn TabContainerTabButtonPressedSignalFn) {
  sig := StringNameFromStr("tab_button_pressed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectTabButtonPressed(subs SignalSubscribers, fn TabContainerTabButtonPressedSignalFn) {
  sig := StringNameFromStr("tab_button_pressed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}

type TabContainerPrePopupPressedSignalFn func()

func (me *TabContainer) ConnectPrePopupPressed(subs SignalSubscribers, fn TabContainerPrePopupPressedSignalFn) {
  sig := StringNameFromStr("pre_popup_pressed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *TabContainer) DisconnectPrePopupPressed(subs SignalSubscribers, fn TabContainerPrePopupPressedSignalFn) {
  sig := StringNameFromStr("pre_popup_pressed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
