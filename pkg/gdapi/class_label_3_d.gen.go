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



// Enums

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

func (me *Label3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Label3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Label3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Label3D) SetHorizontalAlignment(alignment HorizontalAlignment, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetHorizontalAlignment()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetVerticalAlignment(alignment VerticalAlignment, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetVerticalAlignment()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetModulate(modulate Color, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetModulate()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetOutlineModulate(modulate Color, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetOutlineModulate()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetText(text String, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetText()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetTextDirection(direction TextServerDirection, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetTextDirection()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetLanguage(language String, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetLanguage()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetStructuredTextBidiOverride(parser TextServerStructuredTextParser, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetStructuredTextBidiOverride()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetStructuredTextBidiOverrideOptions(args Array, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetStructuredTextBidiOverrideOptions()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetUppercase(enable bool, )  {
  panic("TODO: implement")
}

func  (me *Label3D) IsUppercase()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetRenderPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetRenderPriority()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetOutlineRenderPriority(priority int, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetOutlineRenderPriority()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetFont(font Font, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetFont()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetFontSize(size int, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetFontSize()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetOutlineSize(outline_size int, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetOutlineSize()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetLineSpacing(line_spacing float32, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetLineSpacing()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetAutowrapMode(autowrap_mode TextServerAutowrapMode, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetAutowrapMode()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetJustificationFlags(justification_flags TextServerJustificationFlag, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetJustificationFlags()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetWidth(width float32, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetWidth()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetPixelSize(pixel_size float32, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetPixelSize()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetOffset(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetOffset()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetDrawFlag(flag Label3DDrawFlags, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetDrawFlag(flag Label3DDrawFlags, )  {
  panic("TODO: implement")
}

func  (me *Label3D) SetBillboardMode(mode BaseMaterial3DBillboardMode, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetBillboardMode()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetAlphaCutMode(mode Label3DAlphaCutMode, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetAlphaCutMode()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetAlphaScissorThreshold(threshold float32, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetAlphaScissorThreshold()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetAlphaHashScale(threshold float32, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetAlphaHashScale()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetAlphaAntialiasing(alpha_aa BaseMaterial3DAlphaAntiAliasing, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetAlphaAntialiasing()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetAlphaAntialiasingEdge(edge float32, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetAlphaAntialiasingEdge()  {
  panic("TODO: implement")
}

func  (me *Label3D) SetTextureFilter(mode BaseMaterial3DTextureFilter, )  {
  panic("TODO: implement")
}

func  (me *Label3D) GetTextureFilter()  {
  panic("TODO: implement")
}

func  (me *Label3D) GenerateTriangleMesh()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
