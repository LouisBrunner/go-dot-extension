// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Label3D struct {
  obj gdc.ObjectPtr
}

func (me *Label3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Label3D) BaseClass() string {
  return "Label3D"
}

type Label3DDrawFlags int
const (
  Label3DDrawFlagsFlagShaded Label3DDrawFlags = 0
  Label3DDrawFlagsFlagDoubleSided Label3DDrawFlags = 1
  Label3DDrawFlagsFlagDisableDepthTest Label3DDrawFlags = 2
  Label3DDrawFlagsFlagFixedSize Label3DDrawFlags = 3
  Label3DDrawFlagsFlagMax Label3DDrawFlags = 4
)

type Label3DAlphaCutMode int
const (
  Label3DAlphaCutModeAlphaCutDisabled Label3DAlphaCutMode = 0
  Label3DAlphaCutModeAlphaCutDiscard Label3DAlphaCutMode = 1
  Label3DAlphaCutModeAlphaCutOpaquePrepass Label3DAlphaCutMode = 2
  Label3DAlphaCutModeAlphaCutHash Label3DAlphaCutMode = 3
)

func  (me *Label3D) SetHorizontalAlignment(alignment HorizontalAlignment, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetHorizontalAlignment() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetVerticalAlignment(alignment VerticalAlignment, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetVerticalAlignment() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetModulate(modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetModulate() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetOutlineModulate(modulate Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetOutlineModulate() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetText(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetText() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetTextDirection(direction TextServerDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetTextDirection() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetLanguage(language String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetLanguage() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetStructuredTextBidiOverride() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetStructuredTextBidiOverrideOptions(args Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetStructuredTextBidiOverrideOptions() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetUppercase(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) IsUppercase() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetRenderPriority(priority int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetRenderPriority() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetOutlineRenderPriority(priority int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetOutlineRenderPriority() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetFont(font Font, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetFont() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetFontSize(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetFontSize() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetOutlineSize(outline_size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetOutlineSize() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetLineSpacing(line_spacing float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetLineSpacing() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetAutowrapMode() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetJustificationFlags(justification_flags TextServerJustificationFlag, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetJustificationFlags() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetWidth(width float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetWidth() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetPixelSize(pixel_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetPixelSize() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetOffset(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetOffset() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetDrawFlag(flag Label3DDrawFlags, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetDrawFlag(flag Label3DDrawFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetBillboardMode() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetAlphaCutMode(mode Label3DAlphaCutMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetAlphaCutMode() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetAlphaScissorThreshold(threshold float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetAlphaScissorThreshold() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetAlphaHashScale(threshold float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetAlphaHashScale() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetAlphaAntialiasing() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetAlphaAntialiasingEdge(edge float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetAlphaAntialiasingEdge() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, ) { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GetTextureFilter() { // TODO: return value
  // TODO: implement
}

func  (me *Label3D) GenerateTriangleMesh() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
