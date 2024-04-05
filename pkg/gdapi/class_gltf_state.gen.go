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

type ptrsForGLTFStateList struct {
  fnAddUsedExtension gdc.MethodBindPtr
  fnGetJson gdc.MethodBindPtr
  fnSetJson gdc.MethodBindPtr
  fnGetMajorVersion gdc.MethodBindPtr
  fnSetMajorVersion gdc.MethodBindPtr
  fnGetMinorVersion gdc.MethodBindPtr
  fnSetMinorVersion gdc.MethodBindPtr
  fnGetCopyright gdc.MethodBindPtr
  fnSetCopyright gdc.MethodBindPtr
  fnGetGlbData gdc.MethodBindPtr
  fnSetGlbData gdc.MethodBindPtr
  fnGetUseNamedSkinBinds gdc.MethodBindPtr
  fnSetUseNamedSkinBinds gdc.MethodBindPtr
  fnGetNodes gdc.MethodBindPtr
  fnSetNodes gdc.MethodBindPtr
  fnGetBuffers gdc.MethodBindPtr
  fnSetBuffers gdc.MethodBindPtr
  fnGetBufferViews gdc.MethodBindPtr
  fnSetBufferViews gdc.MethodBindPtr
  fnGetAccessors gdc.MethodBindPtr
  fnSetAccessors gdc.MethodBindPtr
  fnGetMeshes gdc.MethodBindPtr
  fnSetMeshes gdc.MethodBindPtr
  fnGetAnimationPlayersCount gdc.MethodBindPtr
  fnGetAnimationPlayer gdc.MethodBindPtr
  fnGetMaterials gdc.MethodBindPtr
  fnSetMaterials gdc.MethodBindPtr
  fnGetSceneName gdc.MethodBindPtr
  fnSetSceneName gdc.MethodBindPtr
  fnGetBasePath gdc.MethodBindPtr
  fnSetBasePath gdc.MethodBindPtr
  fnGetFilename gdc.MethodBindPtr
  fnSetFilename gdc.MethodBindPtr
  fnGetRootNodes gdc.MethodBindPtr
  fnSetRootNodes gdc.MethodBindPtr
  fnGetTextures gdc.MethodBindPtr
  fnSetTextures gdc.MethodBindPtr
  fnGetTextureSamplers gdc.MethodBindPtr
  fnSetTextureSamplers gdc.MethodBindPtr
  fnGetImages gdc.MethodBindPtr
  fnSetImages gdc.MethodBindPtr
  fnGetSkins gdc.MethodBindPtr
  fnSetSkins gdc.MethodBindPtr
  fnGetCameras gdc.MethodBindPtr
  fnSetCameras gdc.MethodBindPtr
  fnGetLights gdc.MethodBindPtr
  fnSetLights gdc.MethodBindPtr
  fnGetUniqueNames gdc.MethodBindPtr
  fnSetUniqueNames gdc.MethodBindPtr
  fnGetUniqueAnimationNames gdc.MethodBindPtr
  fnSetUniqueAnimationNames gdc.MethodBindPtr
  fnGetSkeletons gdc.MethodBindPtr
  fnSetSkeletons gdc.MethodBindPtr
  fnGetCreateAnimations gdc.MethodBindPtr
  fnSetCreateAnimations gdc.MethodBindPtr
  fnGetAnimations gdc.MethodBindPtr
  fnSetAnimations gdc.MethodBindPtr
  fnGetSceneNode gdc.MethodBindPtr
  fnGetNodeIndex gdc.MethodBindPtr
  fnGetAdditionalData gdc.MethodBindPtr
  fnSetAdditionalData gdc.MethodBindPtr
  fnGetHandleBinaryImage gdc.MethodBindPtr
  fnSetHandleBinaryImage gdc.MethodBindPtr
}

var ptrsForGLTFState ptrsForGLTFStateList

func initGLTFStatePtrs(iface gdc.Interface) {

  className := StringNameFromStr("GLTFState")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("add_used_extension")
    defer methodName.Destroy()
    ptrsForGLTFState.fnAddUsedExtension = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2678287736))
  }
  {
    methodName := StringNameFromStr("get_json")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetJson = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2382534195))
  }
  {
    methodName := StringNameFromStr("set_json")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetJson = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
  }
  {
    methodName := StringNameFromStr("get_major_version")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetMajorVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_major_version")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetMajorVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_minor_version")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetMinorVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_minor_version")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetMinorVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_copyright")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetCopyright = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_copyright")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetCopyright = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_glb_data")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetGlbData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2115431945))
  }
  {
    methodName := StringNameFromStr("set_glb_data")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetGlbData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2971499966))
  }
  {
    methodName := StringNameFromStr("get_use_named_skin_binds")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetUseNamedSkinBinds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_use_named_skin_binds")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetUseNamedSkinBinds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_nodes")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_nodes")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_buffers")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetBuffers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_buffers")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetBuffers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_buffer_views")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetBufferViews = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_buffer_views")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetBufferViews = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_accessors")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetAccessors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_accessors")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetAccessors = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_meshes")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetMeshes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_meshes")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetMeshes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_animation_players_count")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetAnimationPlayersCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
  }
  {
    methodName := StringNameFromStr("get_animation_player")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetAnimationPlayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 925043400))
  }
  {
    methodName := StringNameFromStr("get_materials")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetMaterials = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_materials")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetMaterials = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_scene_name")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetSceneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
  }
  {
    methodName := StringNameFromStr("set_scene_name")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetSceneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_base_path")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetBasePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
  }
  {
    methodName := StringNameFromStr("set_base_path")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetBasePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_filename")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetFilename = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_filename")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetFilename = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_root_nodes")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetRootNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 969006518))
  }
  {
    methodName := StringNameFromStr("set_root_nodes")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetRootNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3614634198))
  }
  {
    methodName := StringNameFromStr("get_textures")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetTextures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_textures")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetTextures = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_texture_samplers")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetTextureSamplers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_texture_samplers")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetTextureSamplers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_images")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetImages = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_images")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetImages = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_skins")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetSkins = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_skins")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetSkins = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_cameras")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetCameras = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_cameras")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetCameras = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_lights")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetLights = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_lights")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetLights = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_unique_names")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetUniqueNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_unique_names")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetUniqueNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_unique_animation_names")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetUniqueAnimationNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_unique_animation_names")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetUniqueAnimationNames = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_skeletons")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetSkeletons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_skeletons")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetSkeletons = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_create_animations")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetCreateAnimations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_create_animations")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetCreateAnimations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_animations")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetAnimations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2915620761))
  }
  {
    methodName := StringNameFromStr("set_animations")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetAnimations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_scene_node")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetSceneNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4253421667))
  }
  {
    methodName := StringNameFromStr("get_node_index")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetNodeIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1205807060))
  }
  {
    methodName := StringNameFromStr("get_additional_data")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetAdditionalData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2138907829))
  }
  {
    methodName := StringNameFromStr("set_additional_data")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetAdditionalData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
  }
  {
    methodName := StringNameFromStr("get_handle_binary_image")
    defer methodName.Destroy()
    ptrsForGLTFState.fnGetHandleBinaryImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_handle_binary_image")
    defer methodName.Destroy()
    ptrsForGLTFState.fnSetHandleBinaryImage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
}

type GLTFState struct {
  Resource
}

func (me *GLTFState) BaseClass() string {
  return "GLTFState"
}

func NewGLTFState() *GLTFState {
  str := StringNameFromStr("GLTFState") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &GLTFState{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  GLTFStateHandleBinaryDiscardTextures = 0
  GLTFStateHandleBinaryExtractTextures = 1
  GLTFStateHandleBinaryEmbedAsBasisu = 2
  GLTFStateHandleBinaryEmbedAsUncompressed = 3
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
  cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr(), gdc.ConstTypePtr(&required) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnAddUsedExtension), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetJson() Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetJson), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetJson(json Dictionary, )  {
  cargs := []gdc.ConstTypePtr{json.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetJson), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetMajorVersion() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetMajorVersion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetMajorVersion(major_version int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&major_version) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetMajorVersion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetMinorVersion() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetMinorVersion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetMinorVersion(minor_version int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&minor_version) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetMinorVersion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetCopyright() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetCopyright), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetCopyright(copyright String, )  {
  cargs := []gdc.ConstTypePtr{copyright.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetCopyright), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetGlbData() PackedByteArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetGlbData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetGlbData(glb_data PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{glb_data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetGlbData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetUseNamedSkinBinds() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetUseNamedSkinBinds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetUseNamedSkinBinds(use_named_skin_binds bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_named_skin_binds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetUseNamedSkinBinds), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetNodes() []GLTFNode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetNodes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFNode](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetNodes(nodes []GLTFNode, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&nodes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetNodes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetBuffers() []PackedByteArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetBuffers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[PackedByteArray](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetBuffers(buffers []PackedByteArray, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetBuffers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetBufferViews() []GLTFBufferView {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetBufferViews), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFBufferView](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetBufferViews(buffer_views []GLTFBufferView, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&buffer_views) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetBufferViews), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetAccessors() []GLTFAccessor {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetAccessors), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFAccessor](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetAccessors(accessors []GLTFAccessor, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&accessors) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetAccessors), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetMeshes() []GLTFMesh {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetMeshes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFMesh](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetMeshes(meshes []GLTFMesh, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&meshes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetMeshes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetAnimationPlayersCount(idx int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetAnimationPlayersCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) GetAnimationPlayer(idx int64, ) AnimationPlayer {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAnimationPlayer()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetAnimationPlayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) GetMaterials() []Material {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetMaterials), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Material](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetMaterials(materials []Material, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&materials) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetMaterials), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetSceneName() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetSceneName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetSceneName(scene_name String, )  {
  cargs := []gdc.ConstTypePtr{scene_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetSceneName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetBasePath() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetBasePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetBasePath(base_path String, )  {
  cargs := []gdc.ConstTypePtr{base_path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetBasePath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetFilename() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetFilename), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetFilename(filename String, )  {
  cargs := []gdc.ConstTypePtr{filename.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetFilename), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetRootNodes() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetRootNodes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetRootNodes(root_nodes PackedInt32Array, )  {
  cargs := []gdc.ConstTypePtr{root_nodes.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetRootNodes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetTextures() []GLTFTexture {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetTextures), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFTexture](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetTextures(textures []GLTFTexture, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&textures) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetTextures), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetTextureSamplers() []GLTFTextureSampler {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetTextureSamplers), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFTextureSampler](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetTextureSamplers(texture_samplers []GLTFTextureSampler, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_samplers) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetTextureSamplers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetImages() []Texture2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetImages), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Texture2D](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetImages(images []Texture2D, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&images) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetImages), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetSkins() []GLTFSkin {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetSkins), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFSkin](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetSkins(skins []GLTFSkin, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skins) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetSkins), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetCameras() []GLTFCamera {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetCameras), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFCamera](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetCameras(cameras []GLTFCamera, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cameras) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetCameras), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetLights() []GLTFLight {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetLights), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFLight](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetLights(lights []GLTFLight, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&lights) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetLights), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetUniqueNames() []String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetUniqueNames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[String](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetUniqueNames(unique_names []String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unique_names) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetUniqueNames), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetUniqueAnimationNames() []String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetUniqueAnimationNames), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[String](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetUniqueAnimationNames(unique_animation_names []String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unique_animation_names) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetUniqueAnimationNames), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetSkeletons() []GLTFSkeleton {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetSkeletons), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFSkeleton](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetSkeletons(skeletons []GLTFSkeleton, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&skeletons) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetSkeletons), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetCreateAnimations() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetCreateAnimations), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetCreateAnimations(create_animations bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&create_animations) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetCreateAnimations), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetAnimations() []GLTFAnimation {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetAnimations), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[GLTFAnimation](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *GLTFState) SetAnimations(animations []GLTFAnimation, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&animations) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetAnimations), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetSceneNode(idx int64, ) Node {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNode()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetSceneNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) GetNodeIndex(scene_node Node, ) int64 {
  cargs := []gdc.ConstTypePtr{scene_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetNodeIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) GetAdditionalData(extension_name StringName, ) Variant {
  cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetAdditionalData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *GLTFState) SetAdditionalData(extension_name StringName, additional_data Variant, )  {
  cargs := []gdc.ConstTypePtr{extension_name.AsCTypePtr(), additional_data.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetAdditionalData), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *GLTFState) GetHandleBinaryImage() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnGetHandleBinaryImage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *GLTFState) SetHandleBinaryImage(method int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&method) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForGLTFState.fnSetHandleBinaryImage), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
