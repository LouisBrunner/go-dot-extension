// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type DisplayServer struct {
  obj gdc.ObjectPtr
}

func (me *DisplayServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DisplayServer) BaseClass() string {
  return "DisplayServer"
}



// Constants

var (
  DisplayServerScreenWithMouseFocus = "-4" // TODO: construct correctly
  DisplayServerScreenWithKeyboardFocus = "-3" // TODO: construct correctly
  DisplayServerScreenPrimary = "-2" // TODO: construct correctly
  DisplayServerScreenOfMainWindow = "-1" // TODO: construct correctly
  DisplayServerMainWindowId = "0" // TODO: construct correctly
  DisplayServerInvalidWindowId = "-1" // TODO: construct correctly
)

// Enums

type DisplayServerFeature int
const (
  DisplayServerFeatureFeatureGlobalMenu DisplayServerFeature = 0
  DisplayServerFeatureFeatureSubwindows DisplayServerFeature = 1
  DisplayServerFeatureFeatureTouchscreen DisplayServerFeature = 2
  DisplayServerFeatureFeatureMouse DisplayServerFeature = 3
  DisplayServerFeatureFeatureMouseWarp DisplayServerFeature = 4
  DisplayServerFeatureFeatureClipboard DisplayServerFeature = 5
  DisplayServerFeatureFeatureVirtualKeyboard DisplayServerFeature = 6
  DisplayServerFeatureFeatureCursorShape DisplayServerFeature = 7
  DisplayServerFeatureFeatureCustomCursorShape DisplayServerFeature = 8
  DisplayServerFeatureFeatureNativeDialog DisplayServerFeature = 9
  DisplayServerFeatureFeatureIme DisplayServerFeature = 10
  DisplayServerFeatureFeatureWindowTransparency DisplayServerFeature = 11
  DisplayServerFeatureFeatureHidpi DisplayServerFeature = 12
  DisplayServerFeatureFeatureIcon DisplayServerFeature = 13
  DisplayServerFeatureFeatureNativeIcon DisplayServerFeature = 14
  DisplayServerFeatureFeatureOrientation DisplayServerFeature = 15
  DisplayServerFeatureFeatureSwapBuffers DisplayServerFeature = 16
  DisplayServerFeatureFeatureClipboardPrimary DisplayServerFeature = 18
  DisplayServerFeatureFeatureTextToSpeech DisplayServerFeature = 19
  DisplayServerFeatureFeatureExtendToTitle DisplayServerFeature = 20
  DisplayServerFeatureFeatureScreenCapture DisplayServerFeature = 21
)

type DisplayServerMouseMode int
const (
  DisplayServerMouseModeMouseModeVisible DisplayServerMouseMode = 0
  DisplayServerMouseModeMouseModeHidden DisplayServerMouseMode = 1
  DisplayServerMouseModeMouseModeCaptured DisplayServerMouseMode = 2
  DisplayServerMouseModeMouseModeConfined DisplayServerMouseMode = 3
  DisplayServerMouseModeMouseModeConfinedHidden DisplayServerMouseMode = 4
)

type DisplayServerScreenOrientation int
const (
  DisplayServerScreenOrientationScreenLandscape DisplayServerScreenOrientation = 0
  DisplayServerScreenOrientationScreenPortrait DisplayServerScreenOrientation = 1
  DisplayServerScreenOrientationScreenReverseLandscape DisplayServerScreenOrientation = 2
  DisplayServerScreenOrientationScreenReversePortrait DisplayServerScreenOrientation = 3
  DisplayServerScreenOrientationScreenSensorLandscape DisplayServerScreenOrientation = 4
  DisplayServerScreenOrientationScreenSensorPortrait DisplayServerScreenOrientation = 5
  DisplayServerScreenOrientationScreenSensor DisplayServerScreenOrientation = 6
)

type DisplayServerVirtualKeyboardType int
const (
  DisplayServerVirtualKeyboardTypeKeyboardTypeDefault DisplayServerVirtualKeyboardType = 0
  DisplayServerVirtualKeyboardTypeKeyboardTypeMultiline DisplayServerVirtualKeyboardType = 1
  DisplayServerVirtualKeyboardTypeKeyboardTypeNumber DisplayServerVirtualKeyboardType = 2
  DisplayServerVirtualKeyboardTypeKeyboardTypeNumberDecimal DisplayServerVirtualKeyboardType = 3
  DisplayServerVirtualKeyboardTypeKeyboardTypePhone DisplayServerVirtualKeyboardType = 4
  DisplayServerVirtualKeyboardTypeKeyboardTypeEmailAddress DisplayServerVirtualKeyboardType = 5
  DisplayServerVirtualKeyboardTypeKeyboardTypePassword DisplayServerVirtualKeyboardType = 6
  DisplayServerVirtualKeyboardTypeKeyboardTypeUrl DisplayServerVirtualKeyboardType = 7
)

type DisplayServerCursorShape int
const (
  DisplayServerCursorShapeCursorArrow DisplayServerCursorShape = 0
  DisplayServerCursorShapeCursorIbeam DisplayServerCursorShape = 1
  DisplayServerCursorShapeCursorPointingHand DisplayServerCursorShape = 2
  DisplayServerCursorShapeCursorCross DisplayServerCursorShape = 3
  DisplayServerCursorShapeCursorWait DisplayServerCursorShape = 4
  DisplayServerCursorShapeCursorBusy DisplayServerCursorShape = 5
  DisplayServerCursorShapeCursorDrag DisplayServerCursorShape = 6
  DisplayServerCursorShapeCursorCanDrop DisplayServerCursorShape = 7
  DisplayServerCursorShapeCursorForbidden DisplayServerCursorShape = 8
  DisplayServerCursorShapeCursorVsize DisplayServerCursorShape = 9
  DisplayServerCursorShapeCursorHsize DisplayServerCursorShape = 10
  DisplayServerCursorShapeCursorBdiagsize DisplayServerCursorShape = 11
  DisplayServerCursorShapeCursorFdiagsize DisplayServerCursorShape = 12
  DisplayServerCursorShapeCursorMove DisplayServerCursorShape = 13
  DisplayServerCursorShapeCursorVsplit DisplayServerCursorShape = 14
  DisplayServerCursorShapeCursorHsplit DisplayServerCursorShape = 15
  DisplayServerCursorShapeCursorHelp DisplayServerCursorShape = 16
  DisplayServerCursorShapeCursorMax DisplayServerCursorShape = 17
)

type DisplayServerWindowMode int
const (
  DisplayServerWindowModeWindowModeWindowed DisplayServerWindowMode = 0
  DisplayServerWindowModeWindowModeMinimized DisplayServerWindowMode = 1
  DisplayServerWindowModeWindowModeMaximized DisplayServerWindowMode = 2
  DisplayServerWindowModeWindowModeFullscreen DisplayServerWindowMode = 3
  DisplayServerWindowModeWindowModeExclusiveFullscreen DisplayServerWindowMode = 4
)

type DisplayServerWindowFlags int
const (
  DisplayServerWindowFlagsWindowFlagResizeDisabled DisplayServerWindowFlags = 0
  DisplayServerWindowFlagsWindowFlagBorderless DisplayServerWindowFlags = 1
  DisplayServerWindowFlagsWindowFlagAlwaysOnTop DisplayServerWindowFlags = 2
  DisplayServerWindowFlagsWindowFlagTransparent DisplayServerWindowFlags = 3
  DisplayServerWindowFlagsWindowFlagNoFocus DisplayServerWindowFlags = 4
  DisplayServerWindowFlagsWindowFlagPopup DisplayServerWindowFlags = 5
  DisplayServerWindowFlagsWindowFlagExtendToTitle DisplayServerWindowFlags = 6
  DisplayServerWindowFlagsWindowFlagMousePassthrough DisplayServerWindowFlags = 7
  DisplayServerWindowFlagsWindowFlagMax DisplayServerWindowFlags = 8
)

type DisplayServerWindowEvent int
const (
  DisplayServerWindowEventWindowEventMouseEnter DisplayServerWindowEvent = 0
  DisplayServerWindowEventWindowEventMouseExit DisplayServerWindowEvent = 1
  DisplayServerWindowEventWindowEventFocusIn DisplayServerWindowEvent = 2
  DisplayServerWindowEventWindowEventFocusOut DisplayServerWindowEvent = 3
  DisplayServerWindowEventWindowEventCloseRequest DisplayServerWindowEvent = 4
  DisplayServerWindowEventWindowEventGoBackRequest DisplayServerWindowEvent = 5
  DisplayServerWindowEventWindowEventDpiChange DisplayServerWindowEvent = 6
  DisplayServerWindowEventWindowEventTitlebarChange DisplayServerWindowEvent = 7
)

type DisplayServerVSyncMode int
const (
  DisplayServerVSyncModeVsyncDisabled DisplayServerVSyncMode = 0
  DisplayServerVSyncModeVsyncEnabled DisplayServerVSyncMode = 1
  DisplayServerVSyncModeVsyncAdaptive DisplayServerVSyncMode = 2
  DisplayServerVSyncModeVsyncMailbox DisplayServerVSyncMode = 3
)

type DisplayServerHandleType int
const (
  DisplayServerHandleTypeDisplayHandle DisplayServerHandleType = 0
  DisplayServerHandleTypeWindowHandle DisplayServerHandleType = 1
  DisplayServerHandleTypeWindowView DisplayServerHandleType = 2
  DisplayServerHandleTypeOpenglContext DisplayServerHandleType = 3
)

type DisplayServerTTSUtteranceEvent int
const (
  DisplayServerTTSUtteranceEventTtsUtteranceStarted DisplayServerTTSUtteranceEvent = 0
  DisplayServerTTSUtteranceEventTtsUtteranceEnded DisplayServerTTSUtteranceEvent = 1
  DisplayServerTTSUtteranceEventTtsUtteranceCanceled DisplayServerTTSUtteranceEvent = 2
  DisplayServerTTSUtteranceEventTtsUtteranceBoundary DisplayServerTTSUtteranceEvent = 3
)

func (me *DisplayServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *DisplayServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *DisplayServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *DisplayServer) HasFeature(feature DisplayServerFeature, ) bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_feature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 334065950) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetName() String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuAddSubmenuItem(menu_root String, label String, submenu String, index int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_add_submenu_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3806306913) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(submenu.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuAddItem(menu_root String, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_add_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3415468211) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(key_callback.AsCTypePtr()), gdc.ConstTypePtr(tag.AsCTypePtr()), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuAddCheckItem(menu_root String, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_add_check_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3415468211) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(key_callback.AsCTypePtr()), gdc.ConstTypePtr(tag.AsCTypePtr()), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuAddIconItem(menu_root String, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_add_icon_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1700867534) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(icon.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(key_callback.AsCTypePtr()), gdc.ConstTypePtr(tag.AsCTypePtr()), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuAddIconCheckItem(menu_root String, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_add_icon_check_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1700867534) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(icon.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(key_callback.AsCTypePtr()), gdc.ConstTypePtr(tag.AsCTypePtr()), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuAddRadioCheckItem(menu_root String, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_add_radio_check_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3415468211) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(key_callback.AsCTypePtr()), gdc.ConstTypePtr(tag.AsCTypePtr()), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuAddIconRadioCheckItem(menu_root String, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_add_icon_radio_check_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1700867534) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(icon.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(key_callback.AsCTypePtr()), gdc.ConstTypePtr(tag.AsCTypePtr()), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuAddMultistateItem(menu_root String, label String, max_states int, default_state int, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_add_multistate_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 635750054) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(label.AsCTypePtr()), gdc.ConstTypePtr(&max_states), gdc.ConstTypePtr(&default_state), gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(key_callback.AsCTypePtr()), gdc.ConstTypePtr(tag.AsCTypePtr()), gdc.ConstTypePtr(&accelerator), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuAddSeparator(menu_root String, index int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_add_separator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1041533178) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemIndexFromText(menu_root String, text String, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_index_from_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2878152881) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemIndexFromTag(menu_root String, tag Variant, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_index_from_tag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2941063483) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(tag.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuIsItemChecked(menu_root String, idx int, ) bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_is_item_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3511468594) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuIsItemCheckable(menu_root String, idx int, ) bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_is_item_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3511468594) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuIsItemRadioCheckable(menu_root String, idx int, ) bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_is_item_radio_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3511468594) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemCallback(menu_root String, idx int, ) Callable {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 748666903) // FIXME: should cache?
  var ret Callable
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemKeyCallback(menu_root String, idx int, ) Callable {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_key_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 748666903) // FIXME: should cache?
  var ret Callable
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemTag(menu_root String, idx int, ) Variant {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_tag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 330672633) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemText(menu_root String, idx int, ) String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 591067909) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemSubmenu(menu_root String, idx int, ) String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_submenu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 591067909) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemAccelerator(menu_root String, idx int, ) Key {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_accelerator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 936065394) // FIXME: should cache?
  var ret Key
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuIsItemDisabled(menu_root String, idx int, ) bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_is_item_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3511468594) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemTooltip(menu_root String, idx int, ) String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 591067909) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemState(menu_root String, idx int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3422818498) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemMaxStates(menu_root String, idx int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_max_states")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3422818498) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemIcon(menu_root String, idx int, ) Texture2D {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3591713183) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuGetItemIndentationLevel(menu_root String, idx int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_indentation_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3422818498) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuSetItemChecked(menu_root String, idx int, checked bool, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_checked")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4108344793) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&checked), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemCheckable(menu_root String, idx int, checkable bool, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4108344793) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&checkable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemRadioCheckable(menu_root String, idx int, checkable bool, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_radio_checkable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4108344793) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&checkable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemCallback(menu_root String, idx int, callback Callable, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3809915389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemKeyCallback(menu_root String, idx int, key_callback Callable, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_key_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3809915389) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(key_callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemTag(menu_root String, idx int, tag Variant, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_tag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 453659863) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(tag.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemText(menu_root String, idx int, text String, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 965966136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemSubmenu(menu_root String, idx int, submenu String, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_submenu")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 965966136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(submenu.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemAccelerator(menu_root String, idx int, keycode Key, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_accelerator")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 566943293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&keycode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemDisabled(menu_root String, idx int, disabled bool, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4108344793) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemTooltip(menu_root String, idx int, tooltip String, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_tooltip")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 965966136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(tooltip.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemState(menu_root String, idx int, state int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3474840532) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&state), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemMaxStates(menu_root String, idx int, max_states int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_max_states")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3474840532) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&max_states), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemIcon(menu_root String, idx int, icon Texture2D, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3201338066) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(icon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuSetItemIndentationLevel(menu_root String, idx int, level int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_set_item_indentation_level")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3474840532) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&level), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuGetItemCount(menu_root String, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_get_item_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GlobalMenuRemoveItem(menu_root String, idx int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_remove_item")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2956805083) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GlobalMenuClear(menu_root String, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_menu_clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(menu_root.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) TtsIsSpeaking() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tts_is_speaking")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) TtsIsPaused() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tts_is_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) TtsGetVoices() Dictionary {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tts_get_voices")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) TtsGetVoicesForLanguage(language String, ) PackedStringArray {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tts_get_voices_for_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4291131558) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) TtsSpeak(text String, voice String, volume int, pitch float32, rate float32, utterance_id int, interrupt bool, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tts_speak")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3741216677) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(voice.AsCTypePtr()), gdc.ConstTypePtr(&volume), gdc.ConstTypePtr(&pitch), gdc.ConstTypePtr(&rate), gdc.ConstTypePtr(&utterance_id), gdc.ConstTypePtr(&interrupt), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) TtsPause()  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tts_pause")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) TtsResume()  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tts_resume")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) TtsStop()  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tts_stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) TtsSetUtteranceCallback(event DisplayServerTTSUtteranceEvent, callable Callable, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tts_set_utterance_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 109679083) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&event), gdc.ConstTypePtr(callable.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) IsDarkModeSupported() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_dark_mode_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) IsDarkMode() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_dark_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetAccentColor() Color {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_accent_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) MouseSetMode(mouse_mode DisplayServerMouseMode, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("mouse_set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 348288463) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mouse_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) MouseGetMode() DisplayServerMouseMode {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("mouse_get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1353961651) // FIXME: should cache?
  var ret DisplayServerMouseMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WarpMouse(position Vector2i, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("warp_mouse")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1130785943) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) MouseGetPosition() Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("mouse_get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) MouseGetButtonState() MouseButtonMask {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("mouse_get_button_state")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2512161324) // FIXME: should cache?
  var ret MouseButtonMask
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ClipboardSet(clipboard String, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clipboard_set")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(clipboard.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) ClipboardGet() String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clipboard_get")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ClipboardHas() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clipboard_has")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ClipboardSetPrimary(clipboard_primary String, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clipboard_set_primary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(clipboard_primary.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) ClipboardGetPrimary() String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clipboard_get_primary")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetDisplayCutouts() Rect2 {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_display_cutouts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetDisplaySafeArea() Rect2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_display_safe_area")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 410525958) // FIXME: should cache?
  var ret Rect2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetScreenCount() int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetPrimaryScreen() int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_primary_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetKeyboardFocusScreen() int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_keyboard_focus_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetScreenFromRect(rect Rect2, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_screen_from_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 741354659) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenGetPosition(screen int, ) Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1725937825) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenGetSize(screen int, ) Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1725937825) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenGetUsableRect(screen int, ) Rect2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_usable_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2439012528) // FIXME: should cache?
  var ret Rect2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenGetDpi(screen int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_dpi")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 181039630) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenGetScale(screen int, ) float32 {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 909105437) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) IsTouchscreenAvailable() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_touchscreen_available")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4162880507) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenGetMaxScale() float32 {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_max_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenGetRefreshRate(screen int, ) float32 {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_refresh_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 909105437) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenGetPixel(position Vector2i, ) Color {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_pixel")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1532707496) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenGetImage(screen int, ) Image {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3813388802) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenSetOrientation(orientation DisplayServerScreenOrientation, screen int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_set_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2629526904) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&orientation), gdc.ConstTypePtr(&screen), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) ScreenGetOrientation(screen int, ) DisplayServerScreenOrientation {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_get_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 133818562) // FIXME: should cache?
  var ret DisplayServerScreenOrientation
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ScreenSetKeepOn(enable bool, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_set_keep_on")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) ScreenIsKeptOn() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("screen_is_kept_on")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetWindowList() PackedInt32Array {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_window_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) GetWindowAtScreenPosition(position Vector2i, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_window_at_screen_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2485466453) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowGetNativeHandle(handle_type DisplayServerHandleType, window_id int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_native_handle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2709193271) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&handle_type), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowGetActivePopup() int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_active_popup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetPopupSafeRect(window int, rect Rect2i, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_popup_safe_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3317281434) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window), gdc.ConstTypePtr(rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetPopupSafeRect(window int, ) Rect2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_popup_safe_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2161169500) // FIXME: should cache?
  var ret Rect2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetTitle(title String, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_title")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3043792800) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(title.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetMousePassthrough(region PackedVector2Array, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_mouse_passthrough")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3958815166) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(region.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetCurrentScreen(window_id int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_current_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetCurrentScreen(screen int, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_current_screen")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3023605688) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetPosition(window_id int, ) Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 763922886) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowGetPositionWithDecorations(window_id int, ) Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_position_with_decorations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 763922886) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetPosition(position Vector2i, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614040015) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetSize(window_id int, ) Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 763922886) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetSize(size Vector2i, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614040015) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetRectChangedCallback(callback Callable, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_rect_changed_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3653650673) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetWindowEventCallback(callback Callable, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_window_event_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3653650673) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetInputEventCallback(callback Callable, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_input_event_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3653650673) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetInputTextCallback(callback Callable, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_input_text_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3653650673) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetDropFilesCallback(callback Callable, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_drop_files_callback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3653650673) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(callback.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetAttachedInstanceId(window_id int, ) int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_attached_instance_id")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1591665591) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowGetMaxSize(window_id int, ) Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 763922886) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetMaxSize(max_size Vector2i, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_max_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614040015) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(max_size.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetMinSize(window_id int, ) Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_min_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 763922886) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetMinSize(min_size Vector2i, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_min_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614040015) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(min_size.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetSizeWithDecorations(window_id int, ) Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_size_with_decorations")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 763922886) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowGetMode(window_id int, ) DisplayServerWindowMode {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2185728461) // FIXME: should cache?
  var ret DisplayServerWindowMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetMode(mode DisplayServerWindowMode, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2942569511) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetFlag(flag DisplayServerWindowFlags, enabled bool, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3971592565) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&enabled), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetFlag(flag DisplayServerWindowFlags, window_id int, ) bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_flag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2662949986) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flag), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetWindowButtonsOffset(offset Vector2i, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_window_buttons_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614040015) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(offset.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetSafeTitleMargins(window_id int, ) Vector3i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_safe_title_margins")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2295066620) // FIXME: should cache?
  var ret Vector3i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowRequestAttention(window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_request_attention")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowMoveToForeground(window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_move_to_foreground")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1995695955) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowIsFocused(window_id int, ) bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_is_focused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1051549951) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowCanDraw(window_id int, ) bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_can_draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1051549951) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowSetTransient(window_id int, parent_window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_transient")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), gdc.ConstTypePtr(&parent_window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetExclusive(window_id int, exclusive bool, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_exclusive")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), gdc.ConstTypePtr(&exclusive), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetImeActive(active bool, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_ime_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 450484987) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&active), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetImePosition(position Vector2i, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_ime_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3614040015) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowSetVsyncMode(vsync_mode DisplayServerVSyncMode, window_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_set_vsync_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1708924624) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vsync_mode), gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) WindowGetVsyncMode(window_id int, ) DisplayServerVSyncMode {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_get_vsync_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 578873795) // FIXME: should cache?
  var ret DisplayServerVSyncMode
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowIsMaximizeAllowed(window_id int, ) bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_is_maximize_allowed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1051549951) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&window_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowMaximizeOnTitleDblClick() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_maximize_on_title_dbl_click")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) WindowMinimizeOnTitleDblClick() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("window_minimize_on_title_dbl_click")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ImeGetSelection() Vector2i {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("ime_get_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3690982128) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ImeGetText() String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("ime_get_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) VirtualKeyboardShow(existing_text String, position Rect2, type_ DisplayServerVirtualKeyboardType, max_length int, cursor_start int, cursor_end int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("virtual_keyboard_show")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 860410478) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(existing_text.AsCTypePtr()), gdc.ConstTypePtr(position.AsCTypePtr()), gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&max_length), gdc.ConstTypePtr(&cursor_start), gdc.ConstTypePtr(&cursor_end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) VirtualKeyboardHide()  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("virtual_keyboard_hide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) VirtualKeyboardGetHeight() int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("virtual_keyboard_get_height")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) CursorSetShape(shape DisplayServerCursorShape, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cursor_set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2026291549) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&shape), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) CursorGetShape() DisplayServerCursorShape {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cursor_get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1087724927) // FIXME: should cache?
  var ret DisplayServerCursorShape
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) CursorSetCustomImage(cursor Resource, shape DisplayServerCursorShape, hotspot Vector2, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cursor_set_custom_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1358907026) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(cursor.AsCTypePtr()), gdc.ConstTypePtr(&shape), gdc.ConstTypePtr(hotspot.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) GetSwapCancelOk() bool {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_swap_cancel_ok")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) EnableForStealingFocus(process_id int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("enable_for_stealing_focus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&process_id), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) DialogShow(title String, description String, buttons PackedStringArray, callback Callable, ) Error {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("dialog_show")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4115553226) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(title.AsCTypePtr()), gdc.ConstTypePtr(description.AsCTypePtr()), gdc.ConstTypePtr(buttons.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) DialogInputText(title String, description String, existing_text String, callback Callable, ) Error {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("dialog_input_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3088703427) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(title.AsCTypePtr()), gdc.ConstTypePtr(description.AsCTypePtr()), gdc.ConstTypePtr(existing_text.AsCTypePtr()), gdc.ConstTypePtr(callback.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) KeyboardGetLayoutCount() int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("keyboard_get_layout_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) KeyboardGetCurrentLayout() int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("keyboard_get_current_layout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) KeyboardSetCurrentLayout(index int, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("keyboard_set_current_layout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) KeyboardGetLayoutLanguage(index int, ) String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("keyboard_get_layout_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) KeyboardGetLayoutName(index int, ) String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("keyboard_get_layout_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) KeyboardGetKeycodeFromPhysical(keycode Key, ) Key {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("keyboard_get_keycode_from_physical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3447613187) // FIXME: should cache?
  var ret Key
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&keycode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) ProcessEvents()  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("process_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) ForceProcessAndDropEvents()  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_process_and_drop_events")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) SetNativeIcon(filename String, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_native_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filename.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) SetIcon(image Image, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_icon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 532598488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(image.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *DisplayServer) TabletGetDriverCount() int {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tablet_get_driver_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) TabletGetDriverName(idx int, ) String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tablet_get_driver_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) TabletGetCurrentDriver() String {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tablet_get_current_driver")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *DisplayServer) TabletSetCurrentDriver(name String, )  {
  classNameV := StringNameFromStr("DisplayServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tablet_set_current_driver")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
