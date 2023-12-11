// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextServer struct {
  obj gdc.ObjectPtr
}

func (me *TextServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextServer) BaseClass() string {
  return "TextServer"
}



// Enums

type TextServerFontAntialiasing int
const (
  TextServerFontAntialiasingFontAntialiasingNone TextServerFontAntialiasing = 0
  TextServerFontAntialiasingFontAntialiasingGray TextServerFontAntialiasing = 1
  TextServerFontAntialiasingFontAntialiasingLcd TextServerFontAntialiasing = 2
)

type TextServerFontLCDSubpixelLayout int
const (
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutNone TextServerFontLCDSubpixelLayout = 0
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutHrgb TextServerFontLCDSubpixelLayout = 1
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutHbgr TextServerFontLCDSubpixelLayout = 2
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutVrgb TextServerFontLCDSubpixelLayout = 3
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutVbgr TextServerFontLCDSubpixelLayout = 4
  TextServerFontLCDSubpixelLayoutFontLcdSubpixelLayoutMax TextServerFontLCDSubpixelLayout = 5
)

type TextServerDirection int
const (
  TextServerDirectionDirectionAuto TextServerDirection = 0
  TextServerDirectionDirectionLtr TextServerDirection = 1
  TextServerDirectionDirectionRtl TextServerDirection = 2
  TextServerDirectionDirectionInherited TextServerDirection = 3
)

type TextServerOrientation int
const (
  TextServerOrientationOrientationHorizontal TextServerOrientation = 0
  TextServerOrientationOrientationVertical TextServerOrientation = 1
)

type TextServerJustificationFlag int
const (
  TextServerJustificationFlagJustificationNone TextServerJustificationFlag = 0
  TextServerJustificationFlagJustificationKashida TextServerJustificationFlag = 1
  TextServerJustificationFlagJustificationWordBound TextServerJustificationFlag = 2
  TextServerJustificationFlagJustificationTrimEdgeSpaces TextServerJustificationFlag = 4
  TextServerJustificationFlagJustificationAfterLastTab TextServerJustificationFlag = 8
  TextServerJustificationFlagJustificationConstrainEllipsis TextServerJustificationFlag = 16
  TextServerJustificationFlagJustificationSkipLastLine TextServerJustificationFlag = 32
  TextServerJustificationFlagJustificationSkipLastLineWithVisibleChars TextServerJustificationFlag = 64
  TextServerJustificationFlagJustificationDoNotSkipSingleLine TextServerJustificationFlag = 128
)

type TextServerAutowrapMode int
const (
  TextServerAutowrapModeAutowrapOff TextServerAutowrapMode = 0
  TextServerAutowrapModeAutowrapArbitrary TextServerAutowrapMode = 1
  TextServerAutowrapModeAutowrapWord TextServerAutowrapMode = 2
  TextServerAutowrapModeAutowrapWordSmart TextServerAutowrapMode = 3
)

type TextServerLineBreakFlag int
const (
  TextServerLineBreakFlagBreakNone TextServerLineBreakFlag = 0
  TextServerLineBreakFlagBreakMandatory TextServerLineBreakFlag = 1
  TextServerLineBreakFlagBreakWordBound TextServerLineBreakFlag = 2
  TextServerLineBreakFlagBreakGraphemeBound TextServerLineBreakFlag = 4
  TextServerLineBreakFlagBreakAdaptive TextServerLineBreakFlag = 8
  TextServerLineBreakFlagBreakTrimEdgeSpaces TextServerLineBreakFlag = 16
)

type TextServerVisibleCharactersBehavior int
const (
  TextServerVisibleCharactersBehaviorVcCharsBeforeShaping TextServerVisibleCharactersBehavior = 0
  TextServerVisibleCharactersBehaviorVcCharsAfterShaping TextServerVisibleCharactersBehavior = 1
  TextServerVisibleCharactersBehaviorVcGlyphsAuto TextServerVisibleCharactersBehavior = 2
  TextServerVisibleCharactersBehaviorVcGlyphsLtr TextServerVisibleCharactersBehavior = 3
  TextServerVisibleCharactersBehaviorVcGlyphsRtl TextServerVisibleCharactersBehavior = 4
)

type TextServerOverrunBehavior int
const (
  TextServerOverrunBehaviorOverrunNoTrimming TextServerOverrunBehavior = 0
  TextServerOverrunBehaviorOverrunTrimChar TextServerOverrunBehavior = 1
  TextServerOverrunBehaviorOverrunTrimWord TextServerOverrunBehavior = 2
  TextServerOverrunBehaviorOverrunTrimEllipsis TextServerOverrunBehavior = 3
  TextServerOverrunBehaviorOverrunTrimWordEllipsis TextServerOverrunBehavior = 4
)

type TextServerTextOverrunFlag int
const (
  TextServerTextOverrunFlagOverrunNoTrim TextServerTextOverrunFlag = 0
  TextServerTextOverrunFlagOverrunTrim TextServerTextOverrunFlag = 1
  TextServerTextOverrunFlagOverrunTrimWordOnly TextServerTextOverrunFlag = 2
  TextServerTextOverrunFlagOverrunAddEllipsis TextServerTextOverrunFlag = 4
  TextServerTextOverrunFlagOverrunEnforceEllipsis TextServerTextOverrunFlag = 8
  TextServerTextOverrunFlagOverrunJustificationAware TextServerTextOverrunFlag = 16
)

type TextServerGraphemeFlag int
const (
  TextServerGraphemeFlagGraphemeIsValid TextServerGraphemeFlag = 1
  TextServerGraphemeFlagGraphemeIsRtl TextServerGraphemeFlag = 2
  TextServerGraphemeFlagGraphemeIsVirtual TextServerGraphemeFlag = 4
  TextServerGraphemeFlagGraphemeIsSpace TextServerGraphemeFlag = 8
  TextServerGraphemeFlagGraphemeIsBreakHard TextServerGraphemeFlag = 16
  TextServerGraphemeFlagGraphemeIsBreakSoft TextServerGraphemeFlag = 32
  TextServerGraphemeFlagGraphemeIsTab TextServerGraphemeFlag = 64
  TextServerGraphemeFlagGraphemeIsElongation TextServerGraphemeFlag = 128
  TextServerGraphemeFlagGraphemeIsPunctuation TextServerGraphemeFlag = 256
  TextServerGraphemeFlagGraphemeIsUnderscore TextServerGraphemeFlag = 512
  TextServerGraphemeFlagGraphemeIsConnected TextServerGraphemeFlag = 1024
  TextServerGraphemeFlagGraphemeIsSafeToInsertTatweel TextServerGraphemeFlag = 2048
  TextServerGraphemeFlagGraphemeIsEmbeddedObject TextServerGraphemeFlag = 4096
)

type TextServerHinting int
const (
  TextServerHintingHintingNone TextServerHinting = 0
  TextServerHintingHintingLight TextServerHinting = 1
  TextServerHintingHintingNormal TextServerHinting = 2
)

type TextServerSubpixelPositioning int
const (
  TextServerSubpixelPositioningSubpixelPositioningDisabled TextServerSubpixelPositioning = 0
  TextServerSubpixelPositioningSubpixelPositioningAuto TextServerSubpixelPositioning = 1
  TextServerSubpixelPositioningSubpixelPositioningOneHalf TextServerSubpixelPositioning = 2
  TextServerSubpixelPositioningSubpixelPositioningOneQuarter TextServerSubpixelPositioning = 3
  TextServerSubpixelPositioningSubpixelPositioningOneHalfMaxSize TextServerSubpixelPositioning = 20
  TextServerSubpixelPositioningSubpixelPositioningOneQuarterMaxSize TextServerSubpixelPositioning = 16
)

type TextServerFeature int
const (
  TextServerFeatureFeatureSimpleLayout TextServerFeature = 1
  TextServerFeatureFeatureBidiLayout TextServerFeature = 2
  TextServerFeatureFeatureVerticalLayout TextServerFeature = 4
  TextServerFeatureFeatureShaping TextServerFeature = 8
  TextServerFeatureFeatureKashidaJustification TextServerFeature = 16
  TextServerFeatureFeatureBreakIterators TextServerFeature = 32
  TextServerFeatureFeatureFontBitmap TextServerFeature = 64
  TextServerFeatureFeatureFontDynamic TextServerFeature = 128
  TextServerFeatureFeatureFontMsdf TextServerFeature = 256
  TextServerFeatureFeatureFontSystem TextServerFeature = 512
  TextServerFeatureFeatureFontVariable TextServerFeature = 1024
  TextServerFeatureFeatureContextSensitiveCaseConversion TextServerFeature = 2048
  TextServerFeatureFeatureUseSupportData TextServerFeature = 4096
  TextServerFeatureFeatureUnicodeIdentifiers TextServerFeature = 8192
  TextServerFeatureFeatureUnicodeSecurity TextServerFeature = 16384
)

type TextServerContourPointTag int
const (
  TextServerContourPointTagContourCurveTagOn TextServerContourPointTag = 1
  TextServerContourPointTagContourCurveTagOffConic TextServerContourPointTag = 0
  TextServerContourPointTagContourCurveTagOffCubic TextServerContourPointTag = 2
)

type TextServerSpacingType int
const (
  TextServerSpacingTypeSpacingGlyph TextServerSpacingType = 0
  TextServerSpacingTypeSpacingSpace TextServerSpacingType = 1
  TextServerSpacingTypeSpacingTop TextServerSpacingType = 2
  TextServerSpacingTypeSpacingBottom TextServerSpacingType = 3
  TextServerSpacingTypeSpacingMax TextServerSpacingType = 4
)

type TextServerFontStyle int
const (
  TextServerFontStyleFontBold TextServerFontStyle = 1
  TextServerFontStyleFontItalic TextServerFontStyle = 2
  TextServerFontStyleFontFixedWidth TextServerFontStyle = 4
)

type TextServerStructuredTextParser int
const (
  TextServerStructuredTextParserStructuredTextDefault TextServerStructuredTextParser = 0
  TextServerStructuredTextParserStructuredTextUri TextServerStructuredTextParser = 1
  TextServerStructuredTextParserStructuredTextFile TextServerStructuredTextParser = 2
  TextServerStructuredTextParserStructuredTextEmail TextServerStructuredTextParser = 3
  TextServerStructuredTextParserStructuredTextList TextServerStructuredTextParser = 4
  TextServerStructuredTextParserStructuredTextGdscript TextServerStructuredTextParser = 5
  TextServerStructuredTextParserStructuredTextCustom TextServerStructuredTextParser = 6
)

func (me *TextServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *TextServer) HasFeature(feature TextServerFeature, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_feature")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3967367083) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&feature), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) GetName() String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) GetFeatures() int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_features")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) LoadSupportData(filename String, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_support_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323990056) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filename.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) GetSupportDataFilename() String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_support_data_filename")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) GetSupportDataInfo() String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_support_data_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) SaveSupportData(filename String, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_support_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(filename.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) IsLocaleRightToLeft(locale String, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_locale_right_to_left")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(locale.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) NameToTag(name String, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("name_to_tag")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) TagToName(tag int, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("tag_to_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&tag), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) Has(rid RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3521089500) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FreeRid(rid RID, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("free_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) CreateFont() RID {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 529393457) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetData(font_rid RID, data PackedByteArray, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1355495400) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontSetFaceIndex(font_rid RID, face_index int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_face_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&face_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetFaceIndex(font_rid RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_face_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetFaceCount(font_rid RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_face_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetStyle(font_rid RID, style TextServerFontStyle, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_style")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 898466325) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&style), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetStyle(font_rid RID, ) TextServerFontStyle {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_style")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3082502592) // FIXME: should cache?
  var ret TextServerFontStyle
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetName(font_rid RID, name String, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2726140452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetName(font_rid RID, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642473191) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetOtNameStrings(font_rid RID, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_ot_name_strings")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1882737106) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetStyleName(font_rid RID, name String, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_style_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2726140452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetStyleName(font_rid RID, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_style_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642473191) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetWeight(font_rid RID, weight int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_weight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&weight), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetWeight(font_rid RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_weight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetStretch(font_rid RID, weight int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&weight), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetStretch(font_rid RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetAntialiasing(font_rid RID, antialiasing TextServerFontAntialiasing, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 958337235) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&antialiasing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetAntialiasing(font_rid RID, ) TextServerFontAntialiasing {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_antialiasing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3389420495) // FIXME: should cache?
  var ret TextServerFontAntialiasing
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetGenerateMipmaps(font_rid RID, generate_mipmaps bool, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_generate_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&generate_mipmaps), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetGenerateMipmaps(font_rid RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_generate_mipmaps")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetMultichannelSignedDistanceField(font_rid RID, msdf bool, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_multichannel_signed_distance_field")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&msdf), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontIsMultichannelSignedDistanceField(font_rid RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_is_multichannel_signed_distance_field")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetMsdfPixelRange(font_rid RID, msdf_pixel_range int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_msdf_pixel_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&msdf_pixel_range), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetMsdfPixelRange(font_rid RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_msdf_pixel_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetMsdfSize(font_rid RID, msdf_size int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_msdf_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&msdf_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetMsdfSize(font_rid RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_msdf_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetFixedSize(font_rid RID, fixed_size int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_fixed_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&fixed_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetFixedSize(font_rid RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_fixed_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetAllowSystemFallback(font_rid RID, allow_system_fallback bool, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_allow_system_fallback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&allow_system_fallback), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontIsAllowSystemFallback(font_rid RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_is_allow_system_fallback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetForceAutohinter(font_rid RID, force_autohinter bool, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_force_autohinter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&force_autohinter), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontIsForceAutohinter(font_rid RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_is_force_autohinter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetHinting(font_rid RID, hinting TextServerHinting, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_hinting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1520010864) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&hinting), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetHinting(font_rid RID, ) TextServerHinting {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_hinting")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3971592737) // FIXME: should cache?
  var ret TextServerHinting
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetSubpixelPositioning(font_rid RID, subpixel_positioning TextServerSubpixelPositioning, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_subpixel_positioning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3830459669) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&subpixel_positioning), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetSubpixelPositioning(font_rid RID, ) TextServerSubpixelPositioning {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_subpixel_positioning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2752233671) // FIXME: should cache?
  var ret TextServerSubpixelPositioning
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetEmbolden(font_rid RID, strength float32, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_embolden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&strength), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetEmbolden(font_rid RID, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_embolden")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetTransform(font_rid RID, transform Transform2D, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1246044741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(transform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetTransform(font_rid RID, ) Transform2D {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 213527486) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetVariationCoordinates(font_rid RID, variation_coordinates Dictionary, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_variation_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1217542888) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(variation_coordinates.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetVariationCoordinates(font_rid RID, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_variation_coordinates")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1882737106) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetOversampling(font_rid RID, oversampling float32, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1794382983) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&oversampling), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetOversampling(font_rid RID, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetSizeCacheList(font_rid RID, ) Vector2i {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_size_cache_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontClearSizeCache(font_rid RID, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_clear_size_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontRemoveSizeCache(font_rid RID, size Vector2i, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_remove_size_cache")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2450610377) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontSetAscent(font_rid RID, size int, ascent float32, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_ascent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1892459533) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&ascent), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetAscent(font_rid RID, size int, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_ascent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 755457166) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetDescent(font_rid RID, size int, descent float32, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_descent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1892459533) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&descent), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetDescent(font_rid RID, size int, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_descent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 755457166) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetUnderlinePosition(font_rid RID, size int, underline_position float32, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_underline_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1892459533) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&underline_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetUnderlinePosition(font_rid RID, size int, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_underline_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 755457166) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetUnderlineThickness(font_rid RID, size int, underline_thickness float32, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_underline_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1892459533) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&underline_thickness), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetUnderlineThickness(font_rid RID, size int, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_underline_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 755457166) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetScale(font_rid RID, size int, scale float32, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1892459533) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetScale(font_rid RID, size int, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 755457166) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetTextureCount(font_rid RID, size Vector2i, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_texture_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1311001310) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontClearTextures(font_rid RID, size Vector2i, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_clear_textures")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2450610377) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontRemoveTexture(font_rid RID, size Vector2i, texture_index int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_remove_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3810512262) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontSetTextureImage(font_rid RID, size Vector2i, texture_index int, image Image, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_texture_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2354485091) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), gdc.ConstTypePtr(image.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetTextureImage(font_rid RID, size Vector2i, texture_index int, ) Image {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_texture_image")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2451761155) // FIXME: should cache?
  var ret Image
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetTextureOffsets(font_rid RID, size Vector2i, texture_index int, offset PackedInt32Array, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_texture_offsets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3005398047) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetTextureOffsets(font_rid RID, size Vector2i, texture_index int, ) PackedInt32Array {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_texture_offsets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3420028887) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&texture_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetGlyphList(font_rid RID, size Vector2i, ) PackedInt32Array {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 46086620) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontClearGlyphs(font_rid RID, size Vector2i, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_clear_glyphs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2450610377) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontRemoveGlyph(font_rid RID, size Vector2i, glyph int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_remove_glyph")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3810512262) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetGlyphAdvance(font_rid RID, size int, glyph int, ) Vector2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_advance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2555689501) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetGlyphAdvance(font_rid RID, size int, glyph int, advance Vector2, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_glyph_advance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3219397315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(advance.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetGlyphOffset(font_rid RID, size Vector2i, glyph int, ) Vector2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 513728628) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetGlyphOffset(font_rid RID, size Vector2i, glyph int, offset Vector2, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_glyph_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1812632090) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetGlyphSize(font_rid RID, size Vector2i, glyph int, ) Vector2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 513728628) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetGlyphSize(font_rid RID, size Vector2i, glyph int, gl_size Vector2, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_glyph_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1812632090) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(gl_size.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetGlyphUvRect(font_rid RID, size Vector2i, glyph int, ) Rect2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_uv_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2274268786) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetGlyphUvRect(font_rid RID, size Vector2i, glyph int, uv_rect Rect2, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_glyph_uv_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1973324081) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(uv_rect.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetGlyphTextureIdx(font_rid RID, size Vector2i, glyph int, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_texture_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4292800474) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetGlyphTextureIdx(font_rid RID, size Vector2i, glyph int, texture_idx int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_glyph_texture_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4254580980) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), gdc.ConstTypePtr(&texture_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetGlyphTextureRid(font_rid RID, size Vector2i, glyph int, ) RID {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_texture_rid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1451696141) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetGlyphTextureSize(font_rid RID, size Vector2i, glyph int, ) Vector2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_texture_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 513728628) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&glyph), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetGlyphContours(font RID, size int, index int, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_contours")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2903964473) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetKerningList(font_rid RID, size int, ) Vector2i {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_kerning_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1778388067) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontClearKerningMap(font_rid RID, size int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_clear_kerning_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3411492887) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontRemoveKerning(font_rid RID, size int, glyph_pair Vector2i, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_remove_kerning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2141860016) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(glyph_pair.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontSetKerning(font_rid RID, size int, glyph_pair Vector2i, kerning Vector2, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_kerning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3630965883) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(glyph_pair.AsCTypePtr()), gdc.ConstTypePtr(kerning.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetKerning(font_rid RID, size int, glyph_pair Vector2i, ) Vector2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_kerning")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1019980169) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(glyph_pair.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetGlyphIndex(font_rid RID, size int, char int, variation_selector int, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_glyph_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1765635060) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&char), gdc.ConstTypePtr(&variation_selector), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetCharFromGlyphIndex(font_rid RID, size int, glyph_index int, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_char_from_glyph_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2156738276) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&glyph_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontHasChar(font_rid RID, char int, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_has_char")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3120086654) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(&char), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetSupportedChars(font_rid RID, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_supported_chars")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642473191) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontRenderRange(font_rid RID, size Vector2i, start int, end int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_render_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4254580980) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontRenderGlyph(font_rid RID, size Vector2i, index int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_render_glyph")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3810512262) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontDrawGlyph(font_rid RID, canvas RID, size int, pos Vector2, index int, color Color, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_draw_glyph")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1821196351) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&index), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontDrawGlyphOutline(font_rid RID, canvas RID, size int, outline_size int, pos Vector2, index int, color Color, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_draw_glyph_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1124898203) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&outline_size), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&index), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontIsLanguageSupported(font_rid RID, language String, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_is_language_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3199320846) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetLanguageSupportOverride(font_rid RID, language String, supported bool, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_language_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2313957094) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), gdc.ConstTypePtr(&supported), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetLanguageSupportOverride(font_rid RID, language String, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_language_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2829184646) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontRemoveLanguageSupportOverride(font_rid RID, language String, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_remove_language_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2726140452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetLanguageSupportOverrides(font_rid RID, ) PackedStringArray {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_language_support_overrides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2801473409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontIsScriptSupported(font_rid RID, script String, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_is_script_supported")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3199320846) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(script.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetScriptSupportOverride(font_rid RID, script String, supported bool, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_script_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2313957094) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(script.AsCTypePtr()), gdc.ConstTypePtr(&supported), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetScriptSupportOverride(font_rid RID, script String, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_script_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2829184646) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(script.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontRemoveScriptSupportOverride(font_rid RID, script String, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_remove_script_support_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2726140452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(script.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetScriptSupportOverrides(font_rid RID, ) PackedStringArray {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_script_support_overrides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2801473409) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetOpentypeFeatureOverrides(font_rid RID, overrides Dictionary, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_opentype_feature_overrides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1217542888) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), gdc.ConstTypePtr(overrides.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) FontGetOpentypeFeatureOverrides(font_rid RID, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_opentype_feature_overrides")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1882737106) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSupportedFeatureList(font_rid RID, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_supported_feature_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1882737106) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSupportedVariationList(font_rid RID, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_supported_variation_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1882737106) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(font_rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontGetGlobalOversampling() float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_get_global_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FontSetGlobalOversampling(oversampling float32, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("font_set_global_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&oversampling), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) GetHexCodeBoxSize(size int, index int, ) Vector2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_hex_code_box_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3016396712) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) DrawHexCodeBox(canvas RID, size int, pos Vector2, index int, color Color, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw_hex_code_box")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602046441) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&index), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) CreateShapedText(direction TextServerDirection, orientation TextServerOrientation, ) RID {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_shaped_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1231398698) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&direction), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextClear(rid RID, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_clear")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(rid.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextSetDirection(shaped RID, direction TextServerDirection, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_set_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2616949700) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextGetDirection(shaped RID, ) TextServerDirection {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3065904362) // FIXME: should cache?
  var ret TextServerDirection
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetInferredDirection(shaped RID, ) TextServerDirection {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_inferred_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3065904362) // FIXME: should cache?
  var ret TextServerDirection
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextSetBidiOverride(shaped RID, override Array, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_set_bidi_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 684822712) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(override.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextSetCustomPunctuation(shaped RID, punct String, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_set_custom_punctuation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2726140452) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(punct.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextGetCustomPunctuation(shaped RID, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_custom_punctuation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 642473191) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextSetOrientation(shaped RID, orientation TextServerOrientation, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_set_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 104095128) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&orientation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextGetOrientation(shaped RID, ) TextServerOrientation {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_orientation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3142708106) // FIXME: should cache?
  var ret TextServerOrientation
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextSetPreserveInvalid(shaped RID, enabled bool, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_set_preserve_invalid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextGetPreserveInvalid(shaped RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_preserve_invalid")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextSetPreserveControl(shaped RID, enabled bool, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_set_preserve_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1265174801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextGetPreserveControl(shaped RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_preserve_control")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextSetSpacing(shaped RID, spacing TextServerSpacingType, value int, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_set_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1307259930) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&spacing), gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextGetSpacing(shaped RID, spacing TextServerSpacingType, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_spacing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1213653558) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&spacing), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextAddString(shaped RID, text String, fonts RID, size int, opentype_features Dictionary, language String, meta Variant, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_add_string")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2621279422) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), gdc.ConstTypePtr(fonts.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(opentype_features.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), gdc.ConstTypePtr(meta.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextAddObject(shaped RID, key Variant, size Vector2, inline_align InlineAlignment, length int, baseline float32, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_add_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2838446185) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&length), gdc.ConstTypePtr(&baseline), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextResizeObject(shaped RID, key Variant, size Vector2, inline_align InlineAlignment, baseline float32, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_resize_object")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2353789835) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(key.AsCTypePtr()), gdc.ConstTypePtr(size.AsCTypePtr()), gdc.ConstTypePtr(&inline_align), gdc.ConstTypePtr(&baseline), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedGetSpanCount(shaped RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_get_span_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedGetSpanMeta(shaped RID, index int, ) Variant {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_get_span_meta")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4069510997) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedSetSpanUpdateFont(shaped RID, index int, fonts RID, size int, opentype_features Dictionary, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_set_span_update_font")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1578983057) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&index), gdc.ConstTypePtr(fonts.AsCTypePtr()), gdc.ConstTypePtr(&size), gdc.ConstTypePtr(opentype_features.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextSubstr(shaped RID, start int, length int, ) RID {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_substr")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1937682086) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetParent(shaped RID, ) RID {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3814569979) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextFitToWidth(shaped RID, width float32, justification_flags TextServerJustificationFlag, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_fit_to_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 603718830) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&justification_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextTabAlign(shaped RID, tab_stops PackedFloat32Array, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_tab_align")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1283669550) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(tab_stops.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextShape(shaped RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3521089500) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextIsReady(shaped RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_is_ready")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextHasVisibleChars(shaped RID, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_has_visible_chars")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4155700596) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetGlyphs(shaped RID, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_glyphs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextSortLogical(shaped RID, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_sort_logical")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2670461153) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetGlyphCount(shaped RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_glyph_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetRange(shaped RID, ) Vector2i {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 733700038) // FIXME: should cache?
  var ret Vector2i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetLineBreaksAdv(shaped RID, width PackedFloat32Array, start int, once bool, break_flags TextServerLineBreakFlag, ) PackedInt32Array {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_line_breaks_adv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4206849830) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(width.AsCTypePtr()), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&once), gdc.ConstTypePtr(&break_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetLineBreaks(shaped RID, width float32, start int, break_flags TextServerLineBreakFlag, ) PackedInt32Array {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_line_breaks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 303410369) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&break_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetWordBreaks(shaped RID, grapheme_flags TextServerGraphemeFlag, ) PackedInt32Array {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_word_breaks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3299477123) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&grapheme_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetTrimPos(shaped RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_trim_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetEllipsisPos(shaped RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_ellipsis_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetEllipsisGlyphs(shaped RID, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_ellipsis_glyphs")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetEllipsisGlyphCount(shaped RID, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_ellipsis_glyph_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2198884583) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextOverrunTrimToWidth(shaped RID, width float32, overrun_trim_flags TextServerTextOverrunFlag, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_overrun_trim_to_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1572579718) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&overrun_trim_flags), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextGetObjects(shaped RID, ) Array {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_objects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2684255073) // FIXME: should cache?
  var ret Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetObjectRect(shaped RID, key Variant, ) Rect2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_object_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 447978354) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(key.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetSize(shaped RID, ) Vector2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2440833711) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetAscent(shaped RID, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_ascent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetDescent(shaped RID, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_descent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetWidth(shaped RID, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetUnderlinePosition(shaped RID, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_underline_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetUnderlineThickness(shaped RID, ) float32 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_underline_thickness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 866169185) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetCarets(shaped RID, position int, ) Dictionary {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_carets")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1574219346) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetSelection(shaped RID, start int, end int, ) PackedVector2Array {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_selection")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3714187733) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextHitTestGrapheme(shaped RID, coords float32, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_hit_test_grapheme")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3149310417) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&coords), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextHitTestPosition(shaped RID, coords float32, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_hit_test_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3149310417) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&coords), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextGetGraphemeBounds(shaped RID, pos int, ) Vector2 {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_grapheme_bounds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2546185844) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&pos), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextNextGraphemePos(shaped RID, pos int, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_next_grapheme_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1120910005) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&pos), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextPrevGraphemePos(shaped RID, pos int, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_prev_grapheme_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1120910005) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&pos), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ShapedTextDraw(shaped RID, canvas RID, pos Vector2, clip_l float32, clip_r float32, color Color, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 70679950) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&clip_l), gdc.ConstTypePtr(&clip_r), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextDrawOutline(shaped RID, canvas RID, pos Vector2, clip_l float32, clip_r float32, outline_size int, color Color, )  {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_draw_outline")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2673671346) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(canvas.AsCTypePtr()), gdc.ConstTypePtr(pos.AsCTypePtr()), gdc.ConstTypePtr(&clip_l), gdc.ConstTypePtr(&clip_r), gdc.ConstTypePtr(&outline_size), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *TextServer) ShapedTextGetDominantDirectionInRange(shaped RID, start int, end int, ) TextServerDirection {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("shaped_text_get_dominant_direction_in_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3326907668) // FIXME: should cache?
  var ret TextServerDirection
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shaped.AsCTypePtr()), gdc.ConstTypePtr(&start), gdc.ConstTypePtr(&end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) FormatNumber(number String, language String, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("format_number")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2305636099) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(number.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ParseNumber(number String, language String, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_number")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2305636099) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(number.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) PercentSign(language String, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("percent_sign")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 993269549) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) StringGetWordBreaks(string_ String, language String, chars_per_line int, ) PackedInt32Array {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("string_get_word_breaks")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1398910359) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), gdc.ConstTypePtr(&chars_per_line), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) IsConfusable(string_ String, dict PackedStringArray, ) int {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_confusable")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1433197768) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), gdc.ConstTypePtr(dict.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) SpoofCheck(string_ String, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("spoof_check")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) StripDiacritics(string_ String, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("strip_diacritics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3135753539) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) IsValidIdentifier(string_ String, ) bool {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_valid_identifier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3927539163) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) StringToUpper(string_ String, language String, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("string_to_upper")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2305636099) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) StringToLower(string_ String, language String, ) String {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("string_to_lower")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2305636099) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(string_.AsCTypePtr()), gdc.ConstTypePtr(language.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *TextServer) ParseStructuredText(parser_type TextServerStructuredTextParser, args Array, text String, ) Vector3i {
  classNameV := StringNameFromStr("TextServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("parse_structured_text")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3310685015) // FIXME: should cache?
  var ret Vector3i
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&parser_type), gdc.ConstTypePtr(args.AsCTypePtr()), gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
