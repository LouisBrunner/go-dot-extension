// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TextServer struct {
  obj gdc.ObjectPtr
}

func (me *TextServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextServer) BaseClass() string {
  return "TextServer"
}

type TextServerFontAntialiasing int
const (
  TextServerFontAntialiasingNone TextServerFontAntialiasing = 0
  TextServerFontAntialiasingGray TextServerFontAntialiasing = 1
  TextServerFontAntialiasingLcd TextServerFontAntialiasing = 2
)

type TextServerFontLCDSubpixelLayout int
const (
  TextServerFontLcdSubpixelLayoutNone TextServerFontLCDSubpixelLayout = 0
  TextServerFontLcdSubpixelLayoutHrgb TextServerFontLCDSubpixelLayout = 1
  TextServerFontLcdSubpixelLayoutHbgr TextServerFontLCDSubpixelLayout = 2
  TextServerFontLcdSubpixelLayoutVrgb TextServerFontLCDSubpixelLayout = 3
  TextServerFontLcdSubpixelLayoutVbgr TextServerFontLCDSubpixelLayout = 4
  TextServerFontLcdSubpixelLayoutMax TextServerFontLCDSubpixelLayout = 5
)

type TextServerDirection int
const (
  TextServerDirectionAuto TextServerDirection = 0
  TextServerDirectionLtr TextServerDirection = 1
  TextServerDirectionRtl TextServerDirection = 2
  TextServerDirectionInherited TextServerDirection = 3
)

type TextServerOrientation int
const (
  TextServerOrientationHorizontal TextServerOrientation = 0
  TextServerOrientationVertical TextServerOrientation = 1
)

type TextServerJustificationFlag int
const (
  TextServerJustificationNone TextServerJustificationFlag = 0
  TextServerJustificationKashida TextServerJustificationFlag = 1
  TextServerJustificationWordBound TextServerJustificationFlag = 2
  TextServerJustificationTrimEdgeSpaces TextServerJustificationFlag = 4
  TextServerJustificationAfterLastTab TextServerJustificationFlag = 8
  TextServerJustificationConstrainEllipsis TextServerJustificationFlag = 16
  TextServerJustificationSkipLastLine TextServerJustificationFlag = 32
  TextServerJustificationSkipLastLineWithVisibleChars TextServerJustificationFlag = 64
  TextServerJustificationDoNotSkipSingleLine TextServerJustificationFlag = 128
)

type TextServerAutowrapMode int
const (
  TextServerAutowrapOff TextServerAutowrapMode = 0
  TextServerAutowrapArbitrary TextServerAutowrapMode = 1
  TextServerAutowrapWord TextServerAutowrapMode = 2
  TextServerAutowrapWordSmart TextServerAutowrapMode = 3
)

type TextServerLineBreakFlag int
const (
  TextServerBreakNone TextServerLineBreakFlag = 0
  TextServerBreakMandatory TextServerLineBreakFlag = 1
  TextServerBreakWordBound TextServerLineBreakFlag = 2
  TextServerBreakGraphemeBound TextServerLineBreakFlag = 4
  TextServerBreakAdaptive TextServerLineBreakFlag = 8
  TextServerBreakTrimEdgeSpaces TextServerLineBreakFlag = 16
)

type TextServerVisibleCharactersBehavior int
const (
  TextServerVcCharsBeforeShaping TextServerVisibleCharactersBehavior = 0
  TextServerVcCharsAfterShaping TextServerVisibleCharactersBehavior = 1
  TextServerVcGlyphsAuto TextServerVisibleCharactersBehavior = 2
  TextServerVcGlyphsLtr TextServerVisibleCharactersBehavior = 3
  TextServerVcGlyphsRtl TextServerVisibleCharactersBehavior = 4
)

type TextServerOverrunBehavior int
const (
  TextServerOverrunNoTrimming TextServerOverrunBehavior = 0
  TextServerOverrunTrimChar TextServerOverrunBehavior = 1
  TextServerOverrunTrimWord TextServerOverrunBehavior = 2
  TextServerOverrunTrimEllipsis TextServerOverrunBehavior = 3
  TextServerOverrunTrimWordEllipsis TextServerOverrunBehavior = 4
)

type TextServerTextOverrunFlag int
const (
  TextServerOverrunNoTrim TextServerTextOverrunFlag = 0
  TextServerOverrunTrim TextServerTextOverrunFlag = 1
  TextServerOverrunTrimWordOnly TextServerTextOverrunFlag = 2
  TextServerOverrunAddEllipsis TextServerTextOverrunFlag = 4
  TextServerOverrunEnforceEllipsis TextServerTextOverrunFlag = 8
  TextServerOverrunJustificationAware TextServerTextOverrunFlag = 16
)

type TextServerGraphemeFlag int
const (
  TextServerGraphemeIsValid TextServerGraphemeFlag = 1
  TextServerGraphemeIsRtl TextServerGraphemeFlag = 2
  TextServerGraphemeIsVirtual TextServerGraphemeFlag = 4
  TextServerGraphemeIsSpace TextServerGraphemeFlag = 8
  TextServerGraphemeIsBreakHard TextServerGraphemeFlag = 16
  TextServerGraphemeIsBreakSoft TextServerGraphemeFlag = 32
  TextServerGraphemeIsTab TextServerGraphemeFlag = 64
  TextServerGraphemeIsElongation TextServerGraphemeFlag = 128
  TextServerGraphemeIsPunctuation TextServerGraphemeFlag = 256
  TextServerGraphemeIsUnderscore TextServerGraphemeFlag = 512
  TextServerGraphemeIsConnected TextServerGraphemeFlag = 1024
  TextServerGraphemeIsSafeToInsertTatweel TextServerGraphemeFlag = 2048
  TextServerGraphemeIsEmbeddedObject TextServerGraphemeFlag = 4096
)

type TextServerHinting int
const (
  TextServerHintingNone TextServerHinting = 0
  TextServerHintingLight TextServerHinting = 1
  TextServerHintingNormal TextServerHinting = 2
)

type TextServerSubpixelPositioning int
const (
  TextServerSubpixelPositioningDisabled TextServerSubpixelPositioning = 0
  TextServerSubpixelPositioningAuto TextServerSubpixelPositioning = 1
  TextServerSubpixelPositioningOneHalf TextServerSubpixelPositioning = 2
  TextServerSubpixelPositioningOneQuarter TextServerSubpixelPositioning = 3
  TextServerSubpixelPositioningOneHalfMaxSize TextServerSubpixelPositioning = 20
  TextServerSubpixelPositioningOneQuarterMaxSize TextServerSubpixelPositioning = 16
)

type TextServerFeature int
const (
  TextServerFeatureSimpleLayout TextServerFeature = 1
  TextServerFeatureBidiLayout TextServerFeature = 2
  TextServerFeatureVerticalLayout TextServerFeature = 4
  TextServerFeatureShaping TextServerFeature = 8
  TextServerFeatureKashidaJustification TextServerFeature = 16
  TextServerFeatureBreakIterators TextServerFeature = 32
  TextServerFeatureFontBitmap TextServerFeature = 64
  TextServerFeatureFontDynamic TextServerFeature = 128
  TextServerFeatureFontMsdf TextServerFeature = 256
  TextServerFeatureFontSystem TextServerFeature = 512
  TextServerFeatureFontVariable TextServerFeature = 1024
  TextServerFeatureContextSensitiveCaseConversion TextServerFeature = 2048
  TextServerFeatureUseSupportData TextServerFeature = 4096
  TextServerFeatureUnicodeIdentifiers TextServerFeature = 8192
  TextServerFeatureUnicodeSecurity TextServerFeature = 16384
)

type TextServerContourPointTag int
const (
  TextServerContourCurveTagOn TextServerContourPointTag = 1
  TextServerContourCurveTagOffConic TextServerContourPointTag = 0
  TextServerContourCurveTagOffCubic TextServerContourPointTag = 2
)

type TextServerSpacingType int
const (
  TextServerSpacingGlyph TextServerSpacingType = 0
  TextServerSpacingSpace TextServerSpacingType = 1
  TextServerSpacingTop TextServerSpacingType = 2
  TextServerSpacingBottom TextServerSpacingType = 3
  TextServerSpacingMax TextServerSpacingType = 4
)

type TextServerFontStyle int
const (
  TextServerFontBold TextServerFontStyle = 1
  TextServerFontItalic TextServerFontStyle = 2
  TextServerFontFixedWidth TextServerFontStyle = 4
)

type TextServerStructuredTextParser int
const (
  TextServerStructuredTextDefault TextServerStructuredTextParser = 0
  TextServerStructuredTextUri TextServerStructuredTextParser = 1
  TextServerStructuredTextFile TextServerStructuredTextParser = 2
  TextServerStructuredTextEmail TextServerStructuredTextParser = 3
  TextServerStructuredTextList TextServerStructuredTextParser = 4
  TextServerStructuredTextGdscript TextServerStructuredTextParser = 5
  TextServerStructuredTextCustom TextServerStructuredTextParser = 6
)
