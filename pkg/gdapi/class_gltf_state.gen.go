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



// Constants

var (
  GLTFStateHandleBinaryDiscardTextures = "0" // TODO: construct correctly
  GLTFStateHandleBinaryExtractTextures = "1" // TODO: construct correctly
  GLTFStateHandleBinaryEmbedAsBasisu = "2" // TODO: construct correctly
  GLTFStateHandleBinaryEmbedAsUncompressed = "3" // TODO: construct correctly
)

// Enums

func (me *GLTFState) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GLTFState) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GLTFState) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GLTFState) AddUsedExtension(extension_name String, required bool, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetJson()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetJson(json Dictionary, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetMajorVersion()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetMajorVersion(major_version int, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetMinorVersion()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetMinorVersion(minor_version int, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetGlbData()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetGlbData(glb_data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetUseNamedSkinBinds()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetUseNamedSkinBinds(use_named_skin_binds bool, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetNodes()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetNodes(nodes GLTFNode, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetBuffers()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetBuffers(buffers PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetBufferViews()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetBufferViews(buffer_views GLTFBufferView, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetAccessors()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetAccessors(accessors GLTFAccessor, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetMeshes()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetMeshes(meshes GLTFMesh, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetAnimationPlayersCount(idx int, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetAnimationPlayer(idx int, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetMaterials()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetMaterials(materials Material, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetSceneName()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetSceneName(scene_name String, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetBasePath()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetBasePath(base_path String, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetRootNodes()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetRootNodes(root_nodes PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetTextures()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetTextures(textures GLTFTexture, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetTextureSamplers()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetTextureSamplers(texture_samplers GLTFTextureSampler, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetImages()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetImages(images Texture2D, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetSkins()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetSkins(skins GLTFSkin, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetCameras()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetCameras(cameras GLTFCamera, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetLights()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetLights(lights GLTFLight, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetUniqueNames()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetUniqueNames(unique_names String, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetUniqueAnimationNames()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetUniqueAnimationNames(unique_animation_names String, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetSkeletons()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetSkeletons(skeletons GLTFSkeleton, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetCreateAnimations()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetCreateAnimations(create_animations bool, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetAnimations()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetAnimations(animations GLTFAnimation, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetSceneNode(idx int, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetNodeIndex(scene_node Node, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetAdditionalData(extension_name StringName, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetAdditionalData(extension_name StringName, additional_data Variant, )  {
  panic("TODO: implement")
}

func  (me *GLTFState) GetHandleBinaryImage()  {
  panic("TODO: implement")
}

func  (me *GLTFState) SetHandleBinaryImage(method int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
