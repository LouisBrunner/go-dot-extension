// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GLTFState struct {
  obj gdc.ObjectPtr
}

func (me *GLTFState) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GLTFState) BaseClass() string {
  return "GLTFState"
}

// TODO: needed?
// const (
// // )

var (
  GLTFStateHandleBinaryDiscardTextures = "0" // TODO: construct correctly
  GLTFStateHandleBinaryExtractTextures = "1" // TODO: construct correctly
  GLTFStateHandleBinaryEmbedAsBasisu = "2" // TODO: construct correctly
  GLTFStateHandleBinaryEmbedAsUncompressed = "3" // TODO: construct correctly
)

func  (me *GLTFState) AddUsedExtension(extension_name String, required bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetJson() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetJson(json Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetMajorVersion() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetMajorVersion(major_version int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetMinorVersion() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetMinorVersion(minor_version int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetGlbData() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetGlbData(glb_data PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetUseNamedSkinBinds() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetUseNamedSkinBinds(use_named_skin_binds bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetNodes() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetNodes(nodes GLTFNode, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetBuffers() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetBuffers(buffers PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetBufferViews() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetBufferViews(buffer_views GLTFBufferView, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetAccessors() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetAccessors(accessors GLTFAccessor, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetMeshes() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetMeshes(meshes GLTFMesh, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetAnimationPlayersCount(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetAnimationPlayer(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetMaterials() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetMaterials(materials Material, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetSceneName() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetSceneName(scene_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetBasePath() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetBasePath(base_path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetRootNodes() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetRootNodes(root_nodes PackedInt32Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetTextures() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetTextures(textures GLTFTexture, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetTextureSamplers() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetTextureSamplers(texture_samplers GLTFTextureSampler, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetImages() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetImages(images Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetSkins() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetSkins(skins GLTFSkin, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetCameras() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetCameras(cameras GLTFCamera, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetLights() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetLights(lights GLTFLight, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetUniqueNames() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetUniqueNames(unique_names String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetUniqueAnimationNames() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetUniqueAnimationNames(unique_animation_names String, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetSkeletons() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetSkeletons(skeletons GLTFSkeleton, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetCreateAnimations() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetCreateAnimations(create_animations bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetAnimations() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetAnimations(animations GLTFAnimation, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetSceneNode(idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetNodeIndex(scene_node Node, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetAdditionalData(extension_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetAdditionalData(extension_name StringName, additional_data Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) GetHandleBinaryImage() { // TODO: return value
  // TODO: implement
}

func  (me *GLTFState) SetHandleBinaryImage(method int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
