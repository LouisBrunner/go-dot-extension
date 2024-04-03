// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&images), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *ImageTextureLayered) UpdateLayer(image Image, layer int64, )  {
  classNameV := StringNameFromStr("ImageTextureLayered")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("update_layer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3331733361) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(image.AsCTypePtr()), gdc.ConstTypePtr(&layer), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
