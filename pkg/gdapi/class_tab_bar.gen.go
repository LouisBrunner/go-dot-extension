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



// Enums

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

func  (me *TabBar) SetTabCount(count int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabCount()  {
  panic("TODO: implement")
}

func  (me *TabBar) SetCurrentTab(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetCurrentTab()  {
  panic("TODO: implement")
}

func  (me *TabBar) GetPreviousTab()  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabTitle(tab_idx int, title String, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabTitle(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabTextDirection(tab_idx int, direction ControlTextDirection, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabTextDirection(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabLanguage(tab_idx int, language String, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabLanguage(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabIcon(tab_idx int, icon Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabIcon(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabIconMaxWidth(tab_idx int, width int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabIconMaxWidth(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabButtonIcon(tab_idx int, icon Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabButtonIcon(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabDisabled(tab_idx int, disabled bool, )  {
  panic("TODO: implement")
}

func  (me *TabBar) IsTabDisabled(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabHidden(tab_idx int, hidden bool, )  {
  panic("TODO: implement")
}

func  (me *TabBar) IsTabHidden(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabMetadata(tab_idx int, metadata Variant, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabMetadata(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) RemoveTab(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) AddTab(title String, icon Texture2D, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabIdxAtPoint(point Vector2, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabAlignment(alignment TabBarAlignmentMode, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabAlignment()  {
  panic("TODO: implement")
}

func  (me *TabBar) SetClipTabs(clip_tabs bool, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetClipTabs()  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabOffset()  {
  panic("TODO: implement")
}

func  (me *TabBar) GetOffsetButtonsVisible()  {
  panic("TODO: implement")
}

func  (me *TabBar) EnsureTabVisible(idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabRect(tab_idx int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) MoveTab(from int, to int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabCloseDisplayPolicy(policy TabBarCloseButtonDisplayPolicy, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabCloseDisplayPolicy()  {
  panic("TODO: implement")
}

func  (me *TabBar) SetMaxTabWidth(width int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetMaxTabWidth()  {
  panic("TODO: implement")
}

func  (me *TabBar) SetScrollingEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetScrollingEnabled()  {
  panic("TODO: implement")
}

func  (me *TabBar) SetDragToRearrangeEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetDragToRearrangeEnabled()  {
  panic("TODO: implement")
}

func  (me *TabBar) SetTabsRearrangeGroup(group_id int, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetTabsRearrangeGroup()  {
  panic("TODO: implement")
}

func  (me *TabBar) SetScrollToSelected(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetScrollToSelected()  {
  panic("TODO: implement")
}

func  (me *TabBar) SetSelectWithRmb(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *TabBar) GetSelectWithRmb()  {
  panic("TODO: implement")
}

func  (me *TabBar) ClearTabs()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
