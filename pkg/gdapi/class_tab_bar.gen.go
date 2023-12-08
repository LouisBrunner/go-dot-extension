// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TabBar struct {
  obj gdc.ObjectPtr
}

func (me *TabBar) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TabBar) BaseClass() string {
  return "TabBar"
}

type TabBarAlignmentMode int
const (
  TabBarAlignmentModeAlignmentLeft TabBarAlignmentMode = 0
  TabBarAlignmentModeAlignmentCenter TabBarAlignmentMode = 1
  TabBarAlignmentModeAlignmentRight TabBarAlignmentMode = 2
  TabBarAlignmentModeAlignmentMax TabBarAlignmentMode = 3
)

type TabBarCloseButtonDisplayPolicy int
const (
  TabBarCloseButtonDisplayPolicyCloseButtonShowNever TabBarCloseButtonDisplayPolicy = 0
  TabBarCloseButtonDisplayPolicyCloseButtonShowActiveOnly TabBarCloseButtonDisplayPolicy = 1
  TabBarCloseButtonDisplayPolicyCloseButtonShowAlways TabBarCloseButtonDisplayPolicy = 2
  TabBarCloseButtonDisplayPolicyCloseButtonMax TabBarCloseButtonDisplayPolicy = 3
)

func  (me *TabBar) SetTabCount(count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabCount() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetCurrentTab(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetCurrentTab() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetPreviousTab() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabTitle(tab_idx int, title String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabTitle(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabTextDirection(tab_idx int, direction ControlTextDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabTextDirection(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabLanguage(tab_idx int, language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabLanguage(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabIcon(tab_idx int, icon Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabIcon(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabIconMaxWidth(tab_idx int, width int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabIconMaxWidth(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabButtonIcon(tab_idx int, icon Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabButtonIcon(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabDisabled(tab_idx int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) IsTabDisabled(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabHidden(tab_idx int, hidden bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) IsTabHidden(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabMetadata(tab_idx int, metadata Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabMetadata(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) RemoveTab(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) AddTab(title String, icon Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabIdxAtPoint(point Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabAlignment(alignment TabBarAlignmentMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabAlignment() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetClipTabs(clip_tabs bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetClipTabs() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabOffset() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetOffsetButtonsVisible() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) EnsureTabVisible(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabRect(tab_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) MoveTab(from int, to int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabCloseDisplayPolicy(policy TabBarCloseButtonDisplayPolicy, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabCloseDisplayPolicy() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetMaxTabWidth(width int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetMaxTabWidth() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetScrollingEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetScrollingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetDragToRearrangeEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetDragToRearrangeEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetTabsRearrangeGroup(group_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetTabsRearrangeGroup() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetScrollToSelected(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetScrollToSelected() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) SetSelectWithRmb(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) GetSelectWithRmb() { // TODO: return value
  // TODO: implement
}

func  (me *TabBar) ClearTabs() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
