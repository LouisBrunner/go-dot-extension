// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

const (
  DisplayServerSCREEN_WITH_MOUSE_FOCUS = -4
  DisplayServerSCREEN_WITH_KEYBOARD_FOCUS = -3
  DisplayServerSCREEN_PRIMARY = -2
  DisplayServerSCREEN_OF_MAIN_WINDOW = -1
  DisplayServerMAIN_WINDOW_ID = 0
  DisplayServerINVALID_WINDOW_ID = -1
)

type DisplayServerFeature int
const (
  DisplayServerFeatureGlobalMenu DisplayServerFeature = 0
  DisplayServerFeatureSubwindows DisplayServerFeature = 1
  DisplayServerFeatureTouchscreen DisplayServerFeature = 2
  DisplayServerFeatureMouse DisplayServerFeature = 3
  DisplayServerFeatureMouseWarp DisplayServerFeature = 4
  DisplayServerFeatureClipboard DisplayServerFeature = 5
  DisplayServerFeatureVirtualKeyboard DisplayServerFeature = 6
  DisplayServerFeatureCursorShape DisplayServerFeature = 7
  DisplayServerFeatureCustomCursorShape DisplayServerFeature = 8
  DisplayServerFeatureNativeDialog DisplayServerFeature = 9
  DisplayServerFeatureIme DisplayServerFeature = 10
  DisplayServerFeatureWindowTransparency DisplayServerFeature = 11
  DisplayServerFeatureHidpi DisplayServerFeature = 12
  DisplayServerFeatureIcon DisplayServerFeature = 13
  DisplayServerFeatureNativeIcon DisplayServerFeature = 14
  DisplayServerFeatureOrientation DisplayServerFeature = 15
  DisplayServerFeatureSwapBuffers DisplayServerFeature = 16
  DisplayServerFeatureClipboardPrimary DisplayServerFeature = 18
  DisplayServerFeatureTextToSpeech DisplayServerFeature = 19
  DisplayServerFeatureExtendToTitle DisplayServerFeature = 20
  DisplayServerFeatureScreenCapture DisplayServerFeature = 21
)

type DisplayServerMouseMode int
const (
  DisplayServerMouseModeVisible DisplayServerMouseMode = 0
  DisplayServerMouseModeHidden DisplayServerMouseMode = 1
  DisplayServerMouseModeCaptured DisplayServerMouseMode = 2
  DisplayServerMouseModeConfined DisplayServerMouseMode = 3
  DisplayServerMouseModeConfinedHidden DisplayServerMouseMode = 4
)

type DisplayServerScreenOrientation int
const (
  DisplayServerScreenLandscape DisplayServerScreenOrientation = 0
  DisplayServerScreenPortrait DisplayServerScreenOrientation = 1
  DisplayServerScreenReverseLandscape DisplayServerScreenOrientation = 2
  DisplayServerScreenReversePortrait DisplayServerScreenOrientation = 3
  DisplayServerScreenSensorLandscape DisplayServerScreenOrientation = 4
  DisplayServerScreenSensorPortrait DisplayServerScreenOrientation = 5
  DisplayServerScreenSensor DisplayServerScreenOrientation = 6
)

type DisplayServerVirtualKeyboardType int
const (
  DisplayServerKeyboardTypeDefault DisplayServerVirtualKeyboardType = 0
  DisplayServerKeyboardTypeMultiline DisplayServerVirtualKeyboardType = 1
  DisplayServerKeyboardTypeNumber DisplayServerVirtualKeyboardType = 2
  DisplayServerKeyboardTypeNumberDecimal DisplayServerVirtualKeyboardType = 3
  DisplayServerKeyboardTypePhone DisplayServerVirtualKeyboardType = 4
  DisplayServerKeyboardTypeEmailAddress DisplayServerVirtualKeyboardType = 5
  DisplayServerKeyboardTypePassword DisplayServerVirtualKeyboardType = 6
  DisplayServerKeyboardTypeUrl DisplayServerVirtualKeyboardType = 7
)

type DisplayServerCursorShape int
const (
  DisplayServerCursorArrow DisplayServerCursorShape = 0
  DisplayServerCursorIbeam DisplayServerCursorShape = 1
  DisplayServerCursorPointingHand DisplayServerCursorShape = 2
  DisplayServerCursorCross DisplayServerCursorShape = 3
  DisplayServerCursorWait DisplayServerCursorShape = 4
  DisplayServerCursorBusy DisplayServerCursorShape = 5
  DisplayServerCursorDrag DisplayServerCursorShape = 6
  DisplayServerCursorCanDrop DisplayServerCursorShape = 7
  DisplayServerCursorForbidden DisplayServerCursorShape = 8
  DisplayServerCursorVsize DisplayServerCursorShape = 9
  DisplayServerCursorHsize DisplayServerCursorShape = 10
  DisplayServerCursorBdiagsize DisplayServerCursorShape = 11
  DisplayServerCursorFdiagsize DisplayServerCursorShape = 12
  DisplayServerCursorMove DisplayServerCursorShape = 13
  DisplayServerCursorVsplit DisplayServerCursorShape = 14
  DisplayServerCursorHsplit DisplayServerCursorShape = 15
  DisplayServerCursorHelp DisplayServerCursorShape = 16
  DisplayServerCursorMax DisplayServerCursorShape = 17
)

type DisplayServerWindowMode int
const (
  DisplayServerWindowModeWindowed DisplayServerWindowMode = 0
  DisplayServerWindowModeMinimized DisplayServerWindowMode = 1
  DisplayServerWindowModeMaximized DisplayServerWindowMode = 2
  DisplayServerWindowModeFullscreen DisplayServerWindowMode = 3
  DisplayServerWindowModeExclusiveFullscreen DisplayServerWindowMode = 4
)

type DisplayServerWindowFlags int
const (
  DisplayServerWindowFlagResizeDisabled DisplayServerWindowFlags = 0
  DisplayServerWindowFlagBorderless DisplayServerWindowFlags = 1
  DisplayServerWindowFlagAlwaysOnTop DisplayServerWindowFlags = 2
  DisplayServerWindowFlagTransparent DisplayServerWindowFlags = 3
  DisplayServerWindowFlagNoFocus DisplayServerWindowFlags = 4
  DisplayServerWindowFlagPopup DisplayServerWindowFlags = 5
  DisplayServerWindowFlagExtendToTitle DisplayServerWindowFlags = 6
  DisplayServerWindowFlagMousePassthrough DisplayServerWindowFlags = 7
  DisplayServerWindowFlagMax DisplayServerWindowFlags = 8
)

type DisplayServerWindowEvent int
const (
  DisplayServerWindowEventMouseEnter DisplayServerWindowEvent = 0
  DisplayServerWindowEventMouseExit DisplayServerWindowEvent = 1
  DisplayServerWindowEventFocusIn DisplayServerWindowEvent = 2
  DisplayServerWindowEventFocusOut DisplayServerWindowEvent = 3
  DisplayServerWindowEventCloseRequest DisplayServerWindowEvent = 4
  DisplayServerWindowEventGoBackRequest DisplayServerWindowEvent = 5
  DisplayServerWindowEventDpiChange DisplayServerWindowEvent = 6
  DisplayServerWindowEventTitlebarChange DisplayServerWindowEvent = 7
)

type DisplayServerVSyncMode int
const (
  DisplayServerVsyncDisabled DisplayServerVSyncMode = 0
  DisplayServerVsyncEnabled DisplayServerVSyncMode = 1
  DisplayServerVsyncAdaptive DisplayServerVSyncMode = 2
  DisplayServerVsyncMailbox DisplayServerVSyncMode = 3
)

type DisplayServerHandleType int
const (
  DisplayServerDisplayHandle DisplayServerHandleType = 0
  DisplayServerWindowHandle DisplayServerHandleType = 1
  DisplayServerWindowView DisplayServerHandleType = 2
  DisplayServerOpenglContext DisplayServerHandleType = 3
)

type DisplayServerTTSUtteranceEvent int
const (
  DisplayServerTtsUtteranceStarted DisplayServerTTSUtteranceEvent = 0
  DisplayServerTtsUtteranceEnded DisplayServerTTSUtteranceEvent = 1
  DisplayServerTtsUtteranceCanceled DisplayServerTTSUtteranceEvent = 2
  DisplayServerTtsUtteranceBoundary DisplayServerTTSUtteranceEvent = 3
)
