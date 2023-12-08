// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type DisplayServer struct {
  obj gdc.ObjectPtr
}

func (me *DisplayServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *DisplayServer) BaseClass() string {
  return "DisplayServer"
}

// TODO: needed?
// const (
// // )

var (
  DisplayServerScreenWithMouseFocus = "-4" // TODO: construct correctly
  DisplayServerScreenWithKeyboardFocus = "-3" // TODO: construct correctly
  DisplayServerScreenPrimary = "-2" // TODO: construct correctly
  DisplayServerScreenOfMainWindow = "-1" // TODO: construct correctly
  DisplayServerMainWindowId = "0" // TODO: construct correctly
  DisplayServerInvalidWindowId = "-1" // TODO: construct correctly
)

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

func  (me *DisplayServer) HasFeature(feature DisplayServerFeature, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetName() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuAddSubmenuItem(menu_root String, label String, submenu String, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuAddItem(menu_root String, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuAddCheckItem(menu_root String, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuAddIconItem(menu_root String, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuAddIconCheckItem(menu_root String, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuAddRadioCheckItem(menu_root String, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuAddIconRadioCheckItem(menu_root String, icon Texture2D, label String, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuAddMultistateItem(menu_root String, label String, max_states int, default_state int, callback Callable, key_callback Callable, tag Variant, accelerator Key, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuAddSeparator(menu_root String, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemIndexFromText(menu_root String, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemIndexFromTag(menu_root String, tag Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuIsItemChecked(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuIsItemCheckable(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuIsItemRadioCheckable(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemCallback(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemKeyCallback(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemTag(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemText(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemSubmenu(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemAccelerator(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuIsItemDisabled(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemTooltip(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemState(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemMaxStates(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemIcon(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemIndentationLevel(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemChecked(menu_root String, idx int, checked bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemCheckable(menu_root String, idx int, checkable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemRadioCheckable(menu_root String, idx int, checkable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemCallback(menu_root String, idx int, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemKeyCallback(menu_root String, idx int, key_callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemTag(menu_root String, idx int, tag Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemText(menu_root String, idx int, text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemSubmenu(menu_root String, idx int, submenu String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemAccelerator(menu_root String, idx int, keycode Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemDisabled(menu_root String, idx int, disabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemTooltip(menu_root String, idx int, tooltip String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemState(menu_root String, idx int, state int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemMaxStates(menu_root String, idx int, max_states int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemIcon(menu_root String, idx int, icon Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuSetItemIndentationLevel(menu_root String, idx int, level int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuGetItemCount(menu_root String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuRemoveItem(menu_root String, idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GlobalMenuClear(menu_root String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TtsIsSpeaking() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TtsIsPaused() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TtsGetVoices() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TtsGetVoicesForLanguage(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TtsSpeak(text String, voice String, volume int, pitch float32, rate float32, utterance_id int, interrupt bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TtsPause() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TtsResume() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TtsStop() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TtsSetUtteranceCallback(event DisplayServerTTSUtteranceEvent, callable Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) IsDarkModeSupported() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) IsDarkMode() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetAccentColor() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) MouseSetMode(mouse_mode DisplayServerMouseMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) MouseGetMode() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WarpMouse(position Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) MouseGetPosition() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) MouseGetButtonState() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ClipboardSet(clipboard String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ClipboardGet() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ClipboardHas() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ClipboardSetPrimary(clipboard_primary String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ClipboardGetPrimary() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetDisplayCutouts() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetDisplaySafeArea() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetScreenCount() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetPrimaryScreen() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetKeyboardFocusScreen() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetScreenFromRect(rect Rect2, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetPosition(screen int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetSize(screen int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetUsableRect(screen int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetDpi(screen int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetScale(screen int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) IsTouchscreenAvailable() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetMaxScale() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetRefreshRate(screen int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetPixel(position Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetImage(screen int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenSetOrientation(orientation DisplayServerScreenOrientation, screen int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenGetOrientation(screen int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenSetKeepOn(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ScreenIsKeptOn() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetWindowList() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetWindowAtScreenPosition(position Vector2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetNativeHandle(handle_type DisplayServerHandleType, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetActivePopup() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetPopupSafeRect(window int, rect Rect2i, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetPopupSafeRect(window int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetTitle(title String, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetMousePassthrough(region PackedVector2Array, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetCurrentScreen(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetCurrentScreen(screen int, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetPosition(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetPositionWithDecorations(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetPosition(position Vector2i, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetSize(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetSize(size Vector2i, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetRectChangedCallback(callback Callable, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetWindowEventCallback(callback Callable, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetInputEventCallback(callback Callable, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetInputTextCallback(callback Callable, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetDropFilesCallback(callback Callable, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetAttachedInstanceId(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetMaxSize(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetMaxSize(max_size Vector2i, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetMinSize(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetMinSize(min_size Vector2i, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetSizeWithDecorations(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetMode(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetMode(mode DisplayServerWindowMode, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetFlag(flag DisplayServerWindowFlags, enabled bool, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetFlag(flag DisplayServerWindowFlags, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetWindowButtonsOffset(offset Vector2i, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetSafeTitleMargins(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowRequestAttention(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowMoveToForeground(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowIsFocused(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowCanDraw(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetTransient(window_id int, parent_window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetExclusive(window_id int, exclusive bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetImeActive(active bool, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetImePosition(position Vector2i, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowSetVsyncMode(vsync_mode DisplayServerVSyncMode, window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowGetVsyncMode(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowIsMaximizeAllowed(window_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowMaximizeOnTitleDblClick() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) WindowMinimizeOnTitleDblClick() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ImeGetSelection() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ImeGetText() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) VirtualKeyboardShow(existing_text String, position Rect2, type_ DisplayServerVirtualKeyboardType, max_length int, cursor_start int, cursor_end int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) VirtualKeyboardHide() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) VirtualKeyboardGetHeight() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) CursorSetShape(shape DisplayServerCursorShape, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) CursorGetShape() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) CursorSetCustomImage(cursor Resource, shape DisplayServerCursorShape, hotspot Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) GetSwapCancelOk() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) EnableForStealingFocus(process_id int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) DialogShow(title String, description String, buttons PackedStringArray, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) DialogInputText(title String, description String, existing_text String, callback Callable, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) KeyboardGetLayoutCount() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) KeyboardGetCurrentLayout() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) KeyboardSetCurrentLayout(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) KeyboardGetLayoutLanguage(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) KeyboardGetLayoutName(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) KeyboardGetKeycodeFromPhysical(keycode Key, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ProcessEvents() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) ForceProcessAndDropEvents() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) SetNativeIcon(filename String, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) SetIcon(image Image, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TabletGetDriverCount() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TabletGetDriverName(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TabletGetCurrentDriver() { // TODO: return value
  // TODO: implement
}

func  (me *DisplayServer) TabletSetCurrentDriver(name String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
