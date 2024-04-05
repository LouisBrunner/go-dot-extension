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

type ptrsForImageTextureLayeredList struct {
  fnCreateFromImages gdc.MethodBindPtr
  fnUpdateLayer gdc.MethodBindPtr
}

var ptrsForImageTextureLayered ptrsForImageTextureLayeredList

func initImageTextureLayeredPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ImageTextureLayered")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("create_from_images")
    defer methodName.Destroy()
    ptrsForImageTextureLayered.fnCreateFromImages = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2785773503))
  }
  {
    methodName := StringNameFromStr("update_layer")
    defer methodName.Destroy()
    ptrsForImageTextureLayered.fnUpdateLayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3331733361))
  }
}

type ImageTextureLayered struct {
  TextureLayered
}

func (me *ImageTextureLayered) BaseClass() string {
  return "ImageTextureLayered"
}

func NewImageTextureLayered() *ImageTextureLayered {
  str := StringNameFromStr("ImageTextureLayered") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ImageTextureLayered{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ImageTextureLayered) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ImageTextureLayered) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ImageTextureLayered) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ImageTextureLayered) CreateFromImages(images []Image, ) Error {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&images) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&images)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageTextureLayered.fnCreateFromImages), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ImageTextureLayered) UpdateLayer(image Image, layer int64, )  {
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForImageTextureLayered.fnUpdateLayer), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
