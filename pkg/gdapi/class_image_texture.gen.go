// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForImageTextureList struct {
  fnCreateFromImage gdc.MethodBindPtr
  fnGetFormat gdc.MethodBindPtr
  fnSetImage gdc.MethodBindPtr
  fnUpdate gdc.MethodBindPtr
  fnSetSizeOverride gdc.MethodBindPtr
}

var ptrsForImageTexture ptrsForImageTextureList

func initImageTexturePtrs(iface gdc.Interface) {

  className := StringNameFromStr("ImageTexture")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("create_from_image")
    defer methodName.Destroy()
    ptrsForImageTexture.fnCreateFromImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2775144163))
  }
  {
    methodName := StringNameFromStr("get_format")
    defer methodName.Destroy()
    ptrsForImageTexture.fnGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3847873762))
  }
  {
    methodName := StringNameFromStr("set_image")
    defer methodName.Destroy()
    ptrsForImageTexture.fnSetImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 532598488))
  }
  {
    methodName := StringNameFromStr("update")
    defer methodName.Destroy()
    ptrsForImageTexture.fnUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 532598488))
  }
  {
    methodName := StringNameFromStr("set_size_override")
    defer methodName.Destroy()
    ptrsForImageTexture.fnSetSizeOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1130785943))
  }
}

type ImageTexture struct {
  Texture2D
}

func (me *ImageTexture) BaseClass() string {
  return "ImageTexture"
}

func NewImageTexture() *ImageTexture {
  str := StringNameFromStr("ImageTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ImageTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ImageTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImageTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImageTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  ImageTextureCreateFromImage(image Image, ) ImageTexture {
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewImageTexture()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageTexture.fnCreateFromImage), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ImageTexture) GetFormat() ImageFormat {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ImageFormat

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageTexture.fnGetFormat), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ImageTexture) SetImage(image Image, )  {
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageTexture.fnSetImage), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImageTexture) Update(image Image, )  {
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageTexture.fnUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ImageTexture) SetSizeOverride(size Vector2i, )  {
  cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageTexture.fnSetSizeOverride), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
