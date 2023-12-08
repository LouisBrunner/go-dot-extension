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

func  (me *GLTFDocumentExtension) XImportPreflight(state GLTFState, extensions PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XGetSupportedExtensions() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XParseNodeExtensions(state GLTFState, gltf_node GLTFNode, extensions Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XParseImageData(state GLTFState, image_data PackedByteArray, mime_type String, ret_image Image, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XParseTextureJson(state GLTFState, texture_json Dictionary, ret_gltf_texture GLTFTexture, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XGenerateSceneNode(state GLTFState, gltf_node GLTFNode, scene_parent Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XImportPostParse(state GLTFState, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XImportNode(state GLTFState, gltf_node GLTFNode, json Dictionary, node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XImportPost(state GLTFState, root Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XExportPreflight(state GLTFState, root Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XConvertSceneNode(state GLTFState, gltf_node GLTFNode, scene_node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XExportNode(state GLTFState, gltf_node GLTFNode, json Dictionary, node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFDocumentExtension) XExportPost(state GLTFState, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
