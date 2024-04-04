// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

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
  classNameV := StringNameFromStr("ImageTextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_from_images")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2785773503) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&images) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret Error
  pinner.Pin(&images)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ImageTextureLayered) UpdateLayer(image Image, layer int64, )  {
  classNameV := StringNameFromStr("ImageTextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3331733361) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{image.AsCTypePtr(), gdc.ConstTypePtr(&layer) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
