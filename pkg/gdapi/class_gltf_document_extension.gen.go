// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFDocumentExtension struct {
  obj gdc.ObjectPtr
}

func (me *GLTFDocumentExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFDocumentExtension) BaseClass() string {
  return "GLTFDocumentExtension"
}



// Enums

func (me *GLTFDocumentExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFDocumentExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFDocumentExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GLTFDocumentExtension) XImportPreflight(state GLTFState, extensions PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XGetSupportedExtensions()  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XParseNodeExtensions(state GLTFState, gltf_node GLTFNode, extensions Dictionary, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XParseImageData(state GLTFState, image_data PackedByteArray, mime_type String, ret_image Image, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XParseTextureJson(state GLTFState, texture_json Dictionary, ret_gltf_texture GLTFTexture, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XGenerateSceneNode(state GLTFState, gltf_node GLTFNode, scene_parent Node, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XImportPostParse(state GLTFState, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XImportNode(state GLTFState, gltf_node GLTFNode, json Dictionary, node Node, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XImportPost(state GLTFState, root Node, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XExportPreflight(state GLTFState, root Node, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XConvertSceneNode(state GLTFState, gltf_node GLTFNode, scene_node Node, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XExportNode(state GLTFState, gltf_node GLTFNode, json Dictionary, node Node, )  {
  panic("TODO: implement")
}

func  (me *GLTFDocumentExtension) XExportPost(state GLTFState, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
