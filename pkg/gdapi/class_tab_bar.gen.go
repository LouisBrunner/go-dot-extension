// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *TabBar) SetTabCount(count int64, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabCount() int64 {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetCurrentTab(tab_idx int64, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetCurrentTab() int64 {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_current_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) GetPreviousTab() int64 {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_previous_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SelectPreviousAvailable() bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_previous_available")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SelectNextAvailable() bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("select_next_available")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetTabTitle(tab_idx int64, title String, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(title.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabTitle(tab_idx int64, ) String {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TabBar) SetTabTextDirection(tab_idx int64, direction ControlTextDirection, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1707680378) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&direction), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabTextDirection(tab_idx int64, ) ControlTextDirection {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_text_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4235602388) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  var ret ControlTextDirection

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TabBar) SetTabLanguage(tab_idx int64, language String, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(language.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabLanguage(tab_idx int64, ) String {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TabBar) SetTabIcon(tab_idx int64, icon Texture2D, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(icon.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabIcon(tab_idx int64, ) Texture2D {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TabBar) SetTabIconMaxWidth(tab_idx int64, width int64, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_icon_max_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&width), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabIconMaxWidth(tab_idx int64, ) int64 {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_icon_max_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetTabButtonIcon(tab_idx int64, icon Texture2D, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_button_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(icon.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabButtonIcon(tab_idx int64, ) Texture2D {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_button_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  ret := NewTexture2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TabBar) SetTabDisabled(tab_idx int64, disabled bool, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&disabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) IsTabDisabled(tab_idx int64, ) bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_tab_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetTabHidden(tab_idx int64, hidden bool, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(&hidden), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) IsTabHidden(tab_idx int64, ) bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_tab_hidden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetTabMetadata(tab_idx int64, metadata Variant, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2152698145) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), gdc.ConstTypePtr(metadata.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabMetadata(tab_idx int64, ) Variant {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_metadata")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4227898402) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TabBar) RemoveTab(tab_idx int64, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) AddTab(title String, icon Texture2D, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1465444425) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(title.AsCTypePtr()), gdc.ConstTypePtr(icon.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabIdxAtPoint(point Vector2, ) int64 {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_idx_at_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3820158470) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(point.AsCTypePtr()), }
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetTabAlignment(alignment TabBarAlignmentMode, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2413632353) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&alignment), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabAlignment() TabBarAlignmentMode {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_alignment")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2178122193) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret TabBarAlignmentMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TabBar) SetClipTabs(clip_tabs bool, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_clip_tabs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&clip_tabs), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetClipTabs() bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_clip_tabs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) GetTabOffset() int64 {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) GetOffsetButtonsVisible() bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_offset_buttons_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) EnsureTabVisible(idx int64, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("ensure_tab_visible")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabRect(tab_idx int64, ) Rect2 {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3327874267) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tab_idx), }
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *TabBar) MoveTab(from int64, to int64, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_tab")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) SetTabCloseDisplayPolicy(policy TabBarCloseButtonDisplayPolicy, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tab_close_display_policy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2212906737) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&policy), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabCloseDisplayPolicy() TabBarCloseButtonDisplayPolicy {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tab_close_display_policy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2956568028) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret TabBarCloseButtonDisplayPolicy

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *TabBar) SetMaxTabWidth(width int64, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_tab_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetMaxTabWidth() int64 {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_tab_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetScrollingEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scrolling_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetScrollingEnabled() bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scrolling_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetDragToRearrangeEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_drag_to_rearrange_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetDragToRearrangeEnabled() bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_drag_to_rearrange_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetTabsRearrangeGroup(group_id int64, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tabs_rearrange_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_id), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetTabsRearrangeGroup() int64 {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tabs_rearrange_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetScrollToSelected(enabled bool, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scroll_to_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetScrollToSelected() bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scroll_to_selected")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) SetSelectWithRmb(enabled bool, )  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_select_with_rmb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *TabBar) GetSelectWithRmb() bool {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_select_with_rmb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *TabBar) ClearTabs()  {
  classNameV := StringNameFromStr("TabBar")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_tabs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type TabBarTabSelectedSignalFn func(tab int, )

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

type TabBarTabChangedSignalFn func(tab int, )

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

type TabBarTabClickedSignalFn func(tab int, )

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

type TabBarTabRmbClickedSignalFn func(tab int, )

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

type TabBarTabClosePressedSignalFn func(tab int, )

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

type TabBarTabButtonPressedSignalFn func(tab int, )

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

type TabBarTabHoveredSignalFn func(tab int, )

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

type TabBarActiveTabRearrangedSignalFn func(idx_to int, )

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
