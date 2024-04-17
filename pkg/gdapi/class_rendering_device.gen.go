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

type ptrsForRenderingDeviceList struct {
	fnTextureCreate                      gdc.MethodBindPtr
	fnTextureCreateShared                gdc.MethodBindPtr
	fnTextureCreateSharedFromSlice       gdc.MethodBindPtr
	fnTextureCreateFromExtension         gdc.MethodBindPtr
	fnTextureUpdate                      gdc.MethodBindPtr
	fnTextureGetData                     gdc.MethodBindPtr
	fnTextureIsFormatSupportedForUsage   gdc.MethodBindPtr
	fnTextureIsShared                    gdc.MethodBindPtr
	fnTextureIsValid                     gdc.MethodBindPtr
	fnTextureCopy                        gdc.MethodBindPtr
	fnTextureClear                       gdc.MethodBindPtr
	fnTextureResolveMultisample          gdc.MethodBindPtr
	fnTextureGetFormat                   gdc.MethodBindPtr
	fnTextureGetNativeHandle             gdc.MethodBindPtr
	fnFramebufferFormatCreate            gdc.MethodBindPtr
	fnFramebufferFormatCreateMultipass   gdc.MethodBindPtr
	fnFramebufferFormatCreateEmpty       gdc.MethodBindPtr
	fnFramebufferFormatGetTextureSamples gdc.MethodBindPtr
	fnFramebufferCreate                  gdc.MethodBindPtr
	fnFramebufferCreateMultipass         gdc.MethodBindPtr
	fnFramebufferCreateEmpty             gdc.MethodBindPtr
	fnFramebufferGetFormat               gdc.MethodBindPtr
	fnFramebufferIsValid                 gdc.MethodBindPtr
	fnSamplerCreate                      gdc.MethodBindPtr
	fnSamplerIsFormatSupportedForFilter  gdc.MethodBindPtr
	fnVertexBufferCreate                 gdc.MethodBindPtr
	fnVertexFormatCreate                 gdc.MethodBindPtr
	fnVertexArrayCreate                  gdc.MethodBindPtr
	fnIndexBufferCreate                  gdc.MethodBindPtr
	fnIndexArrayCreate                   gdc.MethodBindPtr
	fnShaderCompileSpirvFromSource       gdc.MethodBindPtr
	fnShaderCompileBinaryFromSpirv       gdc.MethodBindPtr
	fnShaderCreateFromSpirv              gdc.MethodBindPtr
	fnShaderCreateFromBytecode           gdc.MethodBindPtr
	fnShaderCreatePlaceholder            gdc.MethodBindPtr
	fnShaderGetVertexInputAttributeMask  gdc.MethodBindPtr
	fnUniformBufferCreate                gdc.MethodBindPtr
	fnStorageBufferCreate                gdc.MethodBindPtr
	fnTextureBufferCreate                gdc.MethodBindPtr
	fnUniformSetCreate                   gdc.MethodBindPtr
	fnUniformSetIsValid                  gdc.MethodBindPtr
	fnBufferCopy                         gdc.MethodBindPtr
	fnBufferUpdate                       gdc.MethodBindPtr
	fnBufferClear                        gdc.MethodBindPtr
	fnBufferGetData                      gdc.MethodBindPtr
	fnRenderPipelineCreate               gdc.MethodBindPtr
	fnRenderPipelineIsValid              gdc.MethodBindPtr
	fnComputePipelineCreate              gdc.MethodBindPtr
	fnComputePipelineIsValid             gdc.MethodBindPtr
	fnScreenGetWidth                     gdc.MethodBindPtr
	fnScreenGetHeight                    gdc.MethodBindPtr
	fnScreenGetFramebufferFormat         gdc.MethodBindPtr
	fnDrawListBeginForScreen             gdc.MethodBindPtr
	fnDrawListBegin                      gdc.MethodBindPtr
	fnDrawListBeginSplit                 gdc.MethodBindPtr
	fnDrawListSetBlendConstants          gdc.MethodBindPtr
	fnDrawListBindRenderPipeline         gdc.MethodBindPtr
	fnDrawListBindUniformSet             gdc.MethodBindPtr
	fnDrawListBindVertexArray            gdc.MethodBindPtr
	fnDrawListBindIndexArray             gdc.MethodBindPtr
	fnDrawListSetPushConstant            gdc.MethodBindPtr
	fnDrawListDraw                       gdc.MethodBindPtr
	fnDrawListEnableScissor              gdc.MethodBindPtr
	fnDrawListDisableScissor             gdc.MethodBindPtr
	fnDrawListSwitchToNextPass           gdc.MethodBindPtr
	fnDrawListSwitchToNextPassSplit      gdc.MethodBindPtr
	fnDrawListEnd                        gdc.MethodBindPtr
	fnComputeListBegin                   gdc.MethodBindPtr
	fnComputeListBindComputePipeline     gdc.MethodBindPtr
	fnComputeListSetPushConstant         gdc.MethodBindPtr
	fnComputeListBindUniformSet          gdc.MethodBindPtr
	fnComputeListDispatch                gdc.MethodBindPtr
	fnComputeListDispatchIndirect        gdc.MethodBindPtr
	fnComputeListAddBarrier              gdc.MethodBindPtr
	fnComputeListEnd                     gdc.MethodBindPtr
	fnFreeRid                            gdc.MethodBindPtr
	fnCaptureTimestamp                   gdc.MethodBindPtr
	fnGetCapturedTimestampsCount         gdc.MethodBindPtr
	fnGetCapturedTimestampsFrame         gdc.MethodBindPtr
	fnGetCapturedTimestampGpuTime        gdc.MethodBindPtr
	fnGetCapturedTimestampCpuTime        gdc.MethodBindPtr
	fnGetCapturedTimestampName           gdc.MethodBindPtr
	fnLimitGet                           gdc.MethodBindPtr
	fnGetFrameDelay                      gdc.MethodBindPtr
	fnSubmit                             gdc.MethodBindPtr
	fnSync                               gdc.MethodBindPtr
	fnBarrier                            gdc.MethodBindPtr
	fnFullBarrier                        gdc.MethodBindPtr
	fnCreateLocalDevice                  gdc.MethodBindPtr
	fnSetResourceName                    gdc.MethodBindPtr
	fnDrawCommandBeginLabel              gdc.MethodBindPtr
	fnDrawCommandInsertLabel             gdc.MethodBindPtr
	fnDrawCommandEndLabel                gdc.MethodBindPtr
	fnGetDeviceVendorName                gdc.MethodBindPtr
	fnGetDeviceName                      gdc.MethodBindPtr
	fnGetDevicePipelineCacheUuid         gdc.MethodBindPtr
	fnGetMemoryUsage                     gdc.MethodBindPtr
	fnGetDriverResource                  gdc.MethodBindPtr
}

var ptrsForRenderingDevice ptrsForRenderingDeviceList

func initRenderingDevicePtrs(iface gdc.Interface) {

	className := StringNameFromStr("RenderingDevice")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("texture_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3709173589))
	}
	{
		methodName := StringNameFromStr("texture_create_shared")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureCreateShared = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3178156134))
	}
	{
		methodName := StringNameFromStr("texture_create_shared_from_slice")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureCreateSharedFromSlice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1808971279))
	}
	{
		methodName := StringNameFromStr("texture_create_from_extension")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureCreateFromExtension = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1397171480))
	}
	{
		methodName := StringNameFromStr("texture_update")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1349464008))
	}
	{
		methodName := StringNameFromStr("texture_get_data")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1859412099))
	}
	{
		methodName := StringNameFromStr("texture_is_format_supported_for_usage")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureIsFormatSupportedForUsage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2592520478))
	}
	{
		methodName := StringNameFromStr("texture_is_shared")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureIsShared = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
	}
	{
		methodName := StringNameFromStr("texture_is_valid")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureIsValid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
	}
	{
		methodName := StringNameFromStr("texture_copy")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureCopy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2859522160))
	}
	{
		methodName := StringNameFromStr("texture_clear")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3477703247))
	}
	{
		methodName := StringNameFromStr("texture_resolve_multisample")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureResolveMultisample = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3181288260))
	}
	{
		methodName := StringNameFromStr("texture_get_format")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1374471690))
	}
	{
		methodName := StringNameFromStr("texture_get_native_handle")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureGetNativeHandle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3917799429))
	}
	{
		methodName := StringNameFromStr("framebuffer_format_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFramebufferFormatCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 697032759))
	}
	{
		methodName := StringNameFromStr("framebuffer_format_create_multipass")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFramebufferFormatCreateMultipass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2647479094))
	}
	{
		methodName := StringNameFromStr("framebuffer_format_create_empty")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFramebufferFormatCreateEmpty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 555930169))
	}
	{
		methodName := StringNameFromStr("framebuffer_format_get_texture_samples")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFramebufferFormatGetTextureSamples = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4223391010))
	}
	{
		methodName := StringNameFromStr("framebuffer_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFramebufferCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3284231055))
	}
	{
		methodName := StringNameFromStr("framebuffer_create_multipass")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFramebufferCreateMultipass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1750306695))
	}
	{
		methodName := StringNameFromStr("framebuffer_create_empty")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFramebufferCreateEmpty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3058360618))
	}
	{
		methodName := StringNameFromStr("framebuffer_get_format")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFramebufferGetFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3917799429))
	}
	{
		methodName := StringNameFromStr("framebuffer_is_valid")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFramebufferIsValid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155700596))
	}
	{
		methodName := StringNameFromStr("sampler_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnSamplerCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2327892535))
	}
	{
		methodName := StringNameFromStr("sampler_is_format_supported_for_filter")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnSamplerIsFormatSupportedForFilter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2247922238))
	}
	{
		methodName := StringNameFromStr("vertex_buffer_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnVertexBufferCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3410049843))
	}
	{
		methodName := StringNameFromStr("vertex_format_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnVertexFormatCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1242678479))
	}
	{
		methodName := StringNameFromStr("vertex_array_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnVertexArrayCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3799816279))
	}
	{
		methodName := StringNameFromStr("index_buffer_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnIndexBufferCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3935920523))
	}
	{
		methodName := StringNameFromStr("index_array_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnIndexArrayCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2256026069))
	}
	{
		methodName := StringNameFromStr("shader_compile_spirv_from_source")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnShaderCompileSpirvFromSource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1178973306))
	}
	{
		methodName := StringNameFromStr("shader_compile_binary_from_spirv")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnShaderCompileBinaryFromSpirv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 134910450))
	}
	{
		methodName := StringNameFromStr("shader_create_from_spirv")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnShaderCreateFromSpirv = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 342949005))
	}
	{
		methodName := StringNameFromStr("shader_create_from_bytecode")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnShaderCreateFromBytecode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1687031350))
	}
	{
		methodName := StringNameFromStr("shader_create_placeholder")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnShaderCreatePlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 529393457))
	}
	{
		methodName := StringNameFromStr("shader_get_vertex_input_attribute_mask")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnShaderGetVertexInputAttributeMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3917799429))
	}
	{
		methodName := StringNameFromStr("uniform_buffer_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnUniformBufferCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 34556762))
	}
	{
		methodName := StringNameFromStr("storage_buffer_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnStorageBufferCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2316365934))
	}
	{
		methodName := StringNameFromStr("texture_buffer_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnTextureBufferCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1470338698))
	}
	{
		methodName := StringNameFromStr("uniform_set_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnUniformSetCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2280795797))
	}
	{
		methodName := StringNameFromStr("uniform_set_is_valid")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnUniformSetIsValid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
	}
	{
		methodName := StringNameFromStr("buffer_copy")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnBufferCopy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 864257779))
	}
	{
		methodName := StringNameFromStr("buffer_update")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnBufferUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3454956949))
	}
	{
		methodName := StringNameFromStr("buffer_clear")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnBufferClear = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2452320800))
	}
	{
		methodName := StringNameFromStr("buffer_get_data")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnBufferGetData = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3101830688))
	}
	{
		methodName := StringNameFromStr("render_pipeline_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnRenderPipelineCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2385451958))
	}
	{
		methodName := StringNameFromStr("render_pipeline_is_valid")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnRenderPipelineIsValid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
	}
	{
		methodName := StringNameFromStr("compute_pipeline_create")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputePipelineCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1448838280))
	}
	{
		methodName := StringNameFromStr("compute_pipeline_is_valid")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputePipelineIsValid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3521089500))
	}
	{
		methodName := StringNameFromStr("screen_get_width")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnScreenGetWidth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("screen_get_height")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnScreenGetHeight = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("screen_get_framebuffer_format")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnScreenGetFramebufferFormat = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("draw_list_begin_for_screen")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListBeginForScreen = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3988079995))
	}
	{
		methodName := StringNameFromStr("draw_list_begin")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2686605154))
	}
	{
		methodName := StringNameFromStr("draw_list_begin_split")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListBeginSplit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2406300660))
	}
	{
		methodName := StringNameFromStr("draw_list_set_blend_constants")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListSetBlendConstants = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2878471219))
	}
	{
		methodName := StringNameFromStr("draw_list_bind_render_pipeline")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListBindRenderPipeline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4040184819))
	}
	{
		methodName := StringNameFromStr("draw_list_bind_uniform_set")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListBindUniformSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 749655778))
	}
	{
		methodName := StringNameFromStr("draw_list_bind_vertex_array")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListBindVertexArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4040184819))
	}
	{
		methodName := StringNameFromStr("draw_list_bind_index_array")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListBindIndexArray = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4040184819))
	}
	{
		methodName := StringNameFromStr("draw_list_set_push_constant")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListSetPushConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2772371345))
	}
	{
		methodName := StringNameFromStr("draw_list_draw")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListDraw = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4230067973))
	}
	{
		methodName := StringNameFromStr("draw_list_enable_scissor")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListEnableScissor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 244650101))
	}
	{
		methodName := StringNameFromStr("draw_list_disable_scissor")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListDisableScissor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("draw_list_switch_to_next_pass")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListSwitchToNextPass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("draw_list_switch_to_next_pass_split")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListSwitchToNextPassSplit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2865087369))
	}
	{
		methodName := StringNameFromStr("draw_list_end")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawListEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("compute_list_begin")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputeListBegin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("compute_list_bind_compute_pipeline")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputeListBindComputePipeline = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4040184819))
	}
	{
		methodName := StringNameFromStr("compute_list_set_push_constant")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputeListSetPushConstant = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2772371345))
	}
	{
		methodName := StringNameFromStr("compute_list_bind_uniform_set")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputeListBindUniformSet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 749655778))
	}
	{
		methodName := StringNameFromStr("compute_list_dispatch")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputeListDispatch = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4275841770))
	}
	{
		methodName := StringNameFromStr("compute_list_dispatch_indirect")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputeListDispatchIndirect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 749655778))
	}
	{
		methodName := StringNameFromStr("compute_list_add_barrier")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputeListAddBarrier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("compute_list_end")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnComputeListEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("free_rid")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFreeRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("capture_timestamp")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnCaptureTimestamp = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_captured_timestamps_count")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetCapturedTimestampsCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_captured_timestamps_frame")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetCapturedTimestampsFrame = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_captured_timestamp_gpu_time")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetCapturedTimestampGpuTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_captured_timestamp_cpu_time")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetCapturedTimestampCpuTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("get_captured_timestamp_name")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetCapturedTimestampName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("limit_get")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnLimitGet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1559202131))
	}
	{
		methodName := StringNameFromStr("get_frame_delay")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetFrameDelay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("submit")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnSubmit = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("sync")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnSync = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("barrier")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnBarrier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3718155691))
	}
	{
		methodName := StringNameFromStr("full_barrier")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnFullBarrier = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("create_local_device")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnCreateLocalDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2846302423))
	}
	{
		methodName := StringNameFromStr("set_resource_name")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnSetResourceName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2726140452))
	}
	{
		methodName := StringNameFromStr("draw_command_begin_label")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawCommandBeginLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1636512886))
	}
	{
		methodName := StringNameFromStr("draw_command_insert_label")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawCommandInsertLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1636512886))
	}
	{
		methodName := StringNameFromStr("draw_command_end_label")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnDrawCommandEndLabel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_device_vendor_name")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetDeviceVendorName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_device_name")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetDeviceName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_device_pipeline_cache_uuid")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetDevicePipelineCacheUuid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("get_memory_usage")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetMemoryUsage = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 251690689))
	}
	{
		methodName := StringNameFromStr("get_driver_resource")
		defer methodName.Destroy()
		ptrsForRenderingDevice.fnGetDriverResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501815484))
	}

}

type RenderingDevice struct {
	Object
}

func (me *RenderingDevice) BaseClass() string {
	return "RenderingDevice"
}

func NewRenderingDevice() *RenderingDevice {
	str := StringNameFromStr("RenderingDevice") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &RenderingDevice{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	RenderingDeviceInvalidId       = -1
	RenderingDeviceInvalidFormatId = -1
)

// Enums

type RenderingDeviceDeviceType int

const (
	RenderingDeviceDeviceTypeDeviceTypeOther         RenderingDeviceDeviceType = 0
	RenderingDeviceDeviceTypeDeviceTypeIntegratedGpu RenderingDeviceDeviceType = 1
	RenderingDeviceDeviceTypeDeviceTypeDiscreteGpu   RenderingDeviceDeviceType = 2
	RenderingDeviceDeviceTypeDeviceTypeVirtualGpu    RenderingDeviceDeviceType = 3
	RenderingDeviceDeviceTypeDeviceTypeCpu           RenderingDeviceDeviceType = 4
	RenderingDeviceDeviceTypeDeviceTypeMax           RenderingDeviceDeviceType = 5
)

type RenderingDeviceDriverResource int

const (
	RenderingDeviceDriverResourceDriverResourceLogicalDevice                  RenderingDeviceDriverResource = 0
	RenderingDeviceDriverResourceDriverResourcePhysicalDevice                 RenderingDeviceDriverResource = 1
	RenderingDeviceDriverResourceDriverResourceTopmostObject                  RenderingDeviceDriverResource = 2
	RenderingDeviceDriverResourceDriverResourceCommandQueue                   RenderingDeviceDriverResource = 3
	RenderingDeviceDriverResourceDriverResourceQueueFamily                    RenderingDeviceDriverResource = 4
	RenderingDeviceDriverResourceDriverResourceTexture                        RenderingDeviceDriverResource = 5
	RenderingDeviceDriverResourceDriverResourceTextureView                    RenderingDeviceDriverResource = 6
	RenderingDeviceDriverResourceDriverResourceTextureDataFormat              RenderingDeviceDriverResource = 7
	RenderingDeviceDriverResourceDriverResourceSampler                        RenderingDeviceDriverResource = 8
	RenderingDeviceDriverResourceDriverResourceUniformSet                     RenderingDeviceDriverResource = 9
	RenderingDeviceDriverResourceDriverResourceBuffer                         RenderingDeviceDriverResource = 10
	RenderingDeviceDriverResourceDriverResourceComputePipeline                RenderingDeviceDriverResource = 11
	RenderingDeviceDriverResourceDriverResourceRenderPipeline                 RenderingDeviceDriverResource = 12
	RenderingDeviceDriverResourceDriverResourceVulkanDevice                   RenderingDeviceDriverResource = 0
	RenderingDeviceDriverResourceDriverResourceVulkanPhysicalDevice           RenderingDeviceDriverResource = 1
	RenderingDeviceDriverResourceDriverResourceVulkanInstance                 RenderingDeviceDriverResource = 2
	RenderingDeviceDriverResourceDriverResourceVulkanQueue                    RenderingDeviceDriverResource = 3
	RenderingDeviceDriverResourceDriverResourceVulkanQueueFamilyIndex         RenderingDeviceDriverResource = 4
	RenderingDeviceDriverResourceDriverResourceVulkanImage                    RenderingDeviceDriverResource = 5
	RenderingDeviceDriverResourceDriverResourceVulkanImageView                RenderingDeviceDriverResource = 6
	RenderingDeviceDriverResourceDriverResourceVulkanImageNativeTextureFormat RenderingDeviceDriverResource = 7
	RenderingDeviceDriverResourceDriverResourceVulkanSampler                  RenderingDeviceDriverResource = 8
	RenderingDeviceDriverResourceDriverResourceVulkanDescriptorSet            RenderingDeviceDriverResource = 9
	RenderingDeviceDriverResourceDriverResourceVulkanBuffer                   RenderingDeviceDriverResource = 10
	RenderingDeviceDriverResourceDriverResourceVulkanComputePipeline          RenderingDeviceDriverResource = 11
	RenderingDeviceDriverResourceDriverResourceVulkanRenderPipeline           RenderingDeviceDriverResource = 12
)

type RenderingDeviceDataFormat int

const (
	RenderingDeviceDataFormatDataFormatR4G4UnormPack8                       RenderingDeviceDataFormat = 0
	RenderingDeviceDataFormatDataFormatR4G4B4A4UnormPack16                  RenderingDeviceDataFormat = 1
	RenderingDeviceDataFormatDataFormatB4G4R4A4UnormPack16                  RenderingDeviceDataFormat = 2
	RenderingDeviceDataFormatDataFormatR5G6B5UnormPack16                    RenderingDeviceDataFormat = 3
	RenderingDeviceDataFormatDataFormatB5G6R5UnormPack16                    RenderingDeviceDataFormat = 4
	RenderingDeviceDataFormatDataFormatR5G5B5A1UnormPack16                  RenderingDeviceDataFormat = 5
	RenderingDeviceDataFormatDataFormatB5G5R5A1UnormPack16                  RenderingDeviceDataFormat = 6
	RenderingDeviceDataFormatDataFormatA1R5G5B5UnormPack16                  RenderingDeviceDataFormat = 7
	RenderingDeviceDataFormatDataFormatR8Unorm                              RenderingDeviceDataFormat = 8
	RenderingDeviceDataFormatDataFormatR8Snorm                              RenderingDeviceDataFormat = 9
	RenderingDeviceDataFormatDataFormatR8Uscaled                            RenderingDeviceDataFormat = 10
	RenderingDeviceDataFormatDataFormatR8Sscaled                            RenderingDeviceDataFormat = 11
	RenderingDeviceDataFormatDataFormatR8Uint                               RenderingDeviceDataFormat = 12
	RenderingDeviceDataFormatDataFormatR8Sint                               RenderingDeviceDataFormat = 13
	RenderingDeviceDataFormatDataFormatR8Srgb                               RenderingDeviceDataFormat = 14
	RenderingDeviceDataFormatDataFormatR8G8Unorm                            RenderingDeviceDataFormat = 15
	RenderingDeviceDataFormatDataFormatR8G8Snorm                            RenderingDeviceDataFormat = 16
	RenderingDeviceDataFormatDataFormatR8G8Uscaled                          RenderingDeviceDataFormat = 17
	RenderingDeviceDataFormatDataFormatR8G8Sscaled                          RenderingDeviceDataFormat = 18
	RenderingDeviceDataFormatDataFormatR8G8Uint                             RenderingDeviceDataFormat = 19
	RenderingDeviceDataFormatDataFormatR8G8Sint                             RenderingDeviceDataFormat = 20
	RenderingDeviceDataFormatDataFormatR8G8Srgb                             RenderingDeviceDataFormat = 21
	RenderingDeviceDataFormatDataFormatR8G8B8Unorm                          RenderingDeviceDataFormat = 22
	RenderingDeviceDataFormatDataFormatR8G8B8Snorm                          RenderingDeviceDataFormat = 23
	RenderingDeviceDataFormatDataFormatR8G8B8Uscaled                        RenderingDeviceDataFormat = 24
	RenderingDeviceDataFormatDataFormatR8G8B8Sscaled                        RenderingDeviceDataFormat = 25
	RenderingDeviceDataFormatDataFormatR8G8B8Uint                           RenderingDeviceDataFormat = 26
	RenderingDeviceDataFormatDataFormatR8G8B8Sint                           RenderingDeviceDataFormat = 27
	RenderingDeviceDataFormatDataFormatR8G8B8Srgb                           RenderingDeviceDataFormat = 28
	RenderingDeviceDataFormatDataFormatB8G8R8Unorm                          RenderingDeviceDataFormat = 29
	RenderingDeviceDataFormatDataFormatB8G8R8Snorm                          RenderingDeviceDataFormat = 30
	RenderingDeviceDataFormatDataFormatB8G8R8Uscaled                        RenderingDeviceDataFormat = 31
	RenderingDeviceDataFormatDataFormatB8G8R8Sscaled                        RenderingDeviceDataFormat = 32
	RenderingDeviceDataFormatDataFormatB8G8R8Uint                           RenderingDeviceDataFormat = 33
	RenderingDeviceDataFormatDataFormatB8G8R8Sint                           RenderingDeviceDataFormat = 34
	RenderingDeviceDataFormatDataFormatB8G8R8Srgb                           RenderingDeviceDataFormat = 35
	RenderingDeviceDataFormatDataFormatR8G8B8A8Unorm                        RenderingDeviceDataFormat = 36
	RenderingDeviceDataFormatDataFormatR8G8B8A8Snorm                        RenderingDeviceDataFormat = 37
	RenderingDeviceDataFormatDataFormatR8G8B8A8Uscaled                      RenderingDeviceDataFormat = 38
	RenderingDeviceDataFormatDataFormatR8G8B8A8Sscaled                      RenderingDeviceDataFormat = 39
	RenderingDeviceDataFormatDataFormatR8G8B8A8Uint                         RenderingDeviceDataFormat = 40
	RenderingDeviceDataFormatDataFormatR8G8B8A8Sint                         RenderingDeviceDataFormat = 41
	RenderingDeviceDataFormatDataFormatR8G8B8A8Srgb                         RenderingDeviceDataFormat = 42
	RenderingDeviceDataFormatDataFormatB8G8R8A8Unorm                        RenderingDeviceDataFormat = 43
	RenderingDeviceDataFormatDataFormatB8G8R8A8Snorm                        RenderingDeviceDataFormat = 44
	RenderingDeviceDataFormatDataFormatB8G8R8A8Uscaled                      RenderingDeviceDataFormat = 45
	RenderingDeviceDataFormatDataFormatB8G8R8A8Sscaled                      RenderingDeviceDataFormat = 46
	RenderingDeviceDataFormatDataFormatB8G8R8A8Uint                         RenderingDeviceDataFormat = 47
	RenderingDeviceDataFormatDataFormatB8G8R8A8Sint                         RenderingDeviceDataFormat = 48
	RenderingDeviceDataFormatDataFormatB8G8R8A8Srgb                         RenderingDeviceDataFormat = 49
	RenderingDeviceDataFormatDataFormatA8B8G8R8UnormPack32                  RenderingDeviceDataFormat = 50
	RenderingDeviceDataFormatDataFormatA8B8G8R8SnormPack32                  RenderingDeviceDataFormat = 51
	RenderingDeviceDataFormatDataFormatA8B8G8R8UscaledPack32                RenderingDeviceDataFormat = 52
	RenderingDeviceDataFormatDataFormatA8B8G8R8SscaledPack32                RenderingDeviceDataFormat = 53
	RenderingDeviceDataFormatDataFormatA8B8G8R8UintPack32                   RenderingDeviceDataFormat = 54
	RenderingDeviceDataFormatDataFormatA8B8G8R8SintPack32                   RenderingDeviceDataFormat = 55
	RenderingDeviceDataFormatDataFormatA8B8G8R8SrgbPack32                   RenderingDeviceDataFormat = 56
	RenderingDeviceDataFormatDataFormatA2R10G10B10UnormPack32               RenderingDeviceDataFormat = 57
	RenderingDeviceDataFormatDataFormatA2R10G10B10SnormPack32               RenderingDeviceDataFormat = 58
	RenderingDeviceDataFormatDataFormatA2R10G10B10UscaledPack32             RenderingDeviceDataFormat = 59
	RenderingDeviceDataFormatDataFormatA2R10G10B10SscaledPack32             RenderingDeviceDataFormat = 60
	RenderingDeviceDataFormatDataFormatA2R10G10B10UintPack32                RenderingDeviceDataFormat = 61
	RenderingDeviceDataFormatDataFormatA2R10G10B10SintPack32                RenderingDeviceDataFormat = 62
	RenderingDeviceDataFormatDataFormatA2B10G10R10UnormPack32               RenderingDeviceDataFormat = 63
	RenderingDeviceDataFormatDataFormatA2B10G10R10SnormPack32               RenderingDeviceDataFormat = 64
	RenderingDeviceDataFormatDataFormatA2B10G10R10UscaledPack32             RenderingDeviceDataFormat = 65
	RenderingDeviceDataFormatDataFormatA2B10G10R10SscaledPack32             RenderingDeviceDataFormat = 66
	RenderingDeviceDataFormatDataFormatA2B10G10R10UintPack32                RenderingDeviceDataFormat = 67
	RenderingDeviceDataFormatDataFormatA2B10G10R10SintPack32                RenderingDeviceDataFormat = 68
	RenderingDeviceDataFormatDataFormatR16Unorm                             RenderingDeviceDataFormat = 69
	RenderingDeviceDataFormatDataFormatR16Snorm                             RenderingDeviceDataFormat = 70
	RenderingDeviceDataFormatDataFormatR16Uscaled                           RenderingDeviceDataFormat = 71
	RenderingDeviceDataFormatDataFormatR16Sscaled                           RenderingDeviceDataFormat = 72
	RenderingDeviceDataFormatDataFormatR16Uint                              RenderingDeviceDataFormat = 73
	RenderingDeviceDataFormatDataFormatR16Sint                              RenderingDeviceDataFormat = 74
	RenderingDeviceDataFormatDataFormatR16Sfloat                            RenderingDeviceDataFormat = 75
	RenderingDeviceDataFormatDataFormatR16G16Unorm                          RenderingDeviceDataFormat = 76
	RenderingDeviceDataFormatDataFormatR16G16Snorm                          RenderingDeviceDataFormat = 77
	RenderingDeviceDataFormatDataFormatR16G16Uscaled                        RenderingDeviceDataFormat = 78
	RenderingDeviceDataFormatDataFormatR16G16Sscaled                        RenderingDeviceDataFormat = 79
	RenderingDeviceDataFormatDataFormatR16G16Uint                           RenderingDeviceDataFormat = 80
	RenderingDeviceDataFormatDataFormatR16G16Sint                           RenderingDeviceDataFormat = 81
	RenderingDeviceDataFormatDataFormatR16G16Sfloat                         RenderingDeviceDataFormat = 82
	RenderingDeviceDataFormatDataFormatR16G16B16Unorm                       RenderingDeviceDataFormat = 83
	RenderingDeviceDataFormatDataFormatR16G16B16Snorm                       RenderingDeviceDataFormat = 84
	RenderingDeviceDataFormatDataFormatR16G16B16Uscaled                     RenderingDeviceDataFormat = 85
	RenderingDeviceDataFormatDataFormatR16G16B16Sscaled                     RenderingDeviceDataFormat = 86
	RenderingDeviceDataFormatDataFormatR16G16B16Uint                        RenderingDeviceDataFormat = 87
	RenderingDeviceDataFormatDataFormatR16G16B16Sint                        RenderingDeviceDataFormat = 88
	RenderingDeviceDataFormatDataFormatR16G16B16Sfloat                      RenderingDeviceDataFormat = 89
	RenderingDeviceDataFormatDataFormatR16G16B16A16Unorm                    RenderingDeviceDataFormat = 90
	RenderingDeviceDataFormatDataFormatR16G16B16A16Snorm                    RenderingDeviceDataFormat = 91
	RenderingDeviceDataFormatDataFormatR16G16B16A16Uscaled                  RenderingDeviceDataFormat = 92
	RenderingDeviceDataFormatDataFormatR16G16B16A16Sscaled                  RenderingDeviceDataFormat = 93
	RenderingDeviceDataFormatDataFormatR16G16B16A16Uint                     RenderingDeviceDataFormat = 94
	RenderingDeviceDataFormatDataFormatR16G16B16A16Sint                     RenderingDeviceDataFormat = 95
	RenderingDeviceDataFormatDataFormatR16G16B16A16Sfloat                   RenderingDeviceDataFormat = 96
	RenderingDeviceDataFormatDataFormatR32Uint                              RenderingDeviceDataFormat = 97
	RenderingDeviceDataFormatDataFormatR32Sint                              RenderingDeviceDataFormat = 98
	RenderingDeviceDataFormatDataFormatR32Sfloat                            RenderingDeviceDataFormat = 99
	RenderingDeviceDataFormatDataFormatR32G32Uint                           RenderingDeviceDataFormat = 100
	RenderingDeviceDataFormatDataFormatR32G32Sint                           RenderingDeviceDataFormat = 101
	RenderingDeviceDataFormatDataFormatR32G32Sfloat                         RenderingDeviceDataFormat = 102
	RenderingDeviceDataFormatDataFormatR32G32B32Uint                        RenderingDeviceDataFormat = 103
	RenderingDeviceDataFormatDataFormatR32G32B32Sint                        RenderingDeviceDataFormat = 104
	RenderingDeviceDataFormatDataFormatR32G32B32Sfloat                      RenderingDeviceDataFormat = 105
	RenderingDeviceDataFormatDataFormatR32G32B32A32Uint                     RenderingDeviceDataFormat = 106
	RenderingDeviceDataFormatDataFormatR32G32B32A32Sint                     RenderingDeviceDataFormat = 107
	RenderingDeviceDataFormatDataFormatR32G32B32A32Sfloat                   RenderingDeviceDataFormat = 108
	RenderingDeviceDataFormatDataFormatR64Uint                              RenderingDeviceDataFormat = 109
	RenderingDeviceDataFormatDataFormatR64Sint                              RenderingDeviceDataFormat = 110
	RenderingDeviceDataFormatDataFormatR64Sfloat                            RenderingDeviceDataFormat = 111
	RenderingDeviceDataFormatDataFormatR64G64Uint                           RenderingDeviceDataFormat = 112
	RenderingDeviceDataFormatDataFormatR64G64Sint                           RenderingDeviceDataFormat = 113
	RenderingDeviceDataFormatDataFormatR64G64Sfloat                         RenderingDeviceDataFormat = 114
	RenderingDeviceDataFormatDataFormatR64G64B64Uint                        RenderingDeviceDataFormat = 115
	RenderingDeviceDataFormatDataFormatR64G64B64Sint                        RenderingDeviceDataFormat = 116
	RenderingDeviceDataFormatDataFormatR64G64B64Sfloat                      RenderingDeviceDataFormat = 117
	RenderingDeviceDataFormatDataFormatR64G64B64A64Uint                     RenderingDeviceDataFormat = 118
	RenderingDeviceDataFormatDataFormatR64G64B64A64Sint                     RenderingDeviceDataFormat = 119
	RenderingDeviceDataFormatDataFormatR64G64B64A64Sfloat                   RenderingDeviceDataFormat = 120
	RenderingDeviceDataFormatDataFormatB10G11R11UfloatPack32                RenderingDeviceDataFormat = 121
	RenderingDeviceDataFormatDataFormatE5B9G9R9UfloatPack32                 RenderingDeviceDataFormat = 122
	RenderingDeviceDataFormatDataFormatD16Unorm                             RenderingDeviceDataFormat = 123
	RenderingDeviceDataFormatDataFormatX8D24UnormPack32                     RenderingDeviceDataFormat = 124
	RenderingDeviceDataFormatDataFormatD32Sfloat                            RenderingDeviceDataFormat = 125
	RenderingDeviceDataFormatDataFormatS8Uint                               RenderingDeviceDataFormat = 126
	RenderingDeviceDataFormatDataFormatD16UnormS8Uint                       RenderingDeviceDataFormat = 127
	RenderingDeviceDataFormatDataFormatD24UnormS8Uint                       RenderingDeviceDataFormat = 128
	RenderingDeviceDataFormatDataFormatD32SfloatS8Uint                      RenderingDeviceDataFormat = 129
	RenderingDeviceDataFormatDataFormatBc1RgbUnormBlock                     RenderingDeviceDataFormat = 130
	RenderingDeviceDataFormatDataFormatBc1RgbSrgbBlock                      RenderingDeviceDataFormat = 131
	RenderingDeviceDataFormatDataFormatBc1RgbaUnormBlock                    RenderingDeviceDataFormat = 132
	RenderingDeviceDataFormatDataFormatBc1RgbaSrgbBlock                     RenderingDeviceDataFormat = 133
	RenderingDeviceDataFormatDataFormatBc2UnormBlock                        RenderingDeviceDataFormat = 134
	RenderingDeviceDataFormatDataFormatBc2SrgbBlock                         RenderingDeviceDataFormat = 135
	RenderingDeviceDataFormatDataFormatBc3UnormBlock                        RenderingDeviceDataFormat = 136
	RenderingDeviceDataFormatDataFormatBc3SrgbBlock                         RenderingDeviceDataFormat = 137
	RenderingDeviceDataFormatDataFormatBc4UnormBlock                        RenderingDeviceDataFormat = 138
	RenderingDeviceDataFormatDataFormatBc4SnormBlock                        RenderingDeviceDataFormat = 139
	RenderingDeviceDataFormatDataFormatBc5UnormBlock                        RenderingDeviceDataFormat = 140
	RenderingDeviceDataFormatDataFormatBc5SnormBlock                        RenderingDeviceDataFormat = 141
	RenderingDeviceDataFormatDataFormatBc6HUfloatBlock                      RenderingDeviceDataFormat = 142
	RenderingDeviceDataFormatDataFormatBc6HSfloatBlock                      RenderingDeviceDataFormat = 143
	RenderingDeviceDataFormatDataFormatBc7UnormBlock                        RenderingDeviceDataFormat = 144
	RenderingDeviceDataFormatDataFormatBc7SrgbBlock                         RenderingDeviceDataFormat = 145
	RenderingDeviceDataFormatDataFormatEtc2R8G8B8UnormBlock                 RenderingDeviceDataFormat = 146
	RenderingDeviceDataFormatDataFormatEtc2R8G8B8SrgbBlock                  RenderingDeviceDataFormat = 147
	RenderingDeviceDataFormatDataFormatEtc2R8G8B8A1UnormBlock               RenderingDeviceDataFormat = 148
	RenderingDeviceDataFormatDataFormatEtc2R8G8B8A1SrgbBlock                RenderingDeviceDataFormat = 149
	RenderingDeviceDataFormatDataFormatEtc2R8G8B8A8UnormBlock               RenderingDeviceDataFormat = 150
	RenderingDeviceDataFormatDataFormatEtc2R8G8B8A8SrgbBlock                RenderingDeviceDataFormat = 151
	RenderingDeviceDataFormatDataFormatEacR11UnormBlock                     RenderingDeviceDataFormat = 152
	RenderingDeviceDataFormatDataFormatEacR11SnormBlock                     RenderingDeviceDataFormat = 153
	RenderingDeviceDataFormatDataFormatEacR11G11UnormBlock                  RenderingDeviceDataFormat = 154
	RenderingDeviceDataFormatDataFormatEacR11G11SnormBlock                  RenderingDeviceDataFormat = 155
	RenderingDeviceDataFormatDataFormatAstc4X4UnormBlock                    RenderingDeviceDataFormat = 156
	RenderingDeviceDataFormatDataFormatAstc4X4SrgbBlock                     RenderingDeviceDataFormat = 157
	RenderingDeviceDataFormatDataFormatAstc5X4UnormBlock                    RenderingDeviceDataFormat = 158
	RenderingDeviceDataFormatDataFormatAstc5X4SrgbBlock                     RenderingDeviceDataFormat = 159
	RenderingDeviceDataFormatDataFormatAstc5X5UnormBlock                    RenderingDeviceDataFormat = 160
	RenderingDeviceDataFormatDataFormatAstc5X5SrgbBlock                     RenderingDeviceDataFormat = 161
	RenderingDeviceDataFormatDataFormatAstc6X5UnormBlock                    RenderingDeviceDataFormat = 162
	RenderingDeviceDataFormatDataFormatAstc6X5SrgbBlock                     RenderingDeviceDataFormat = 163
	RenderingDeviceDataFormatDataFormatAstc6X6UnormBlock                    RenderingDeviceDataFormat = 164
	RenderingDeviceDataFormatDataFormatAstc6X6SrgbBlock                     RenderingDeviceDataFormat = 165
	RenderingDeviceDataFormatDataFormatAstc8X5UnormBlock                    RenderingDeviceDataFormat = 166
	RenderingDeviceDataFormatDataFormatAstc8X5SrgbBlock                     RenderingDeviceDataFormat = 167
	RenderingDeviceDataFormatDataFormatAstc8X6UnormBlock                    RenderingDeviceDataFormat = 168
	RenderingDeviceDataFormatDataFormatAstc8X6SrgbBlock                     RenderingDeviceDataFormat = 169
	RenderingDeviceDataFormatDataFormatAstc8X8UnormBlock                    RenderingDeviceDataFormat = 170
	RenderingDeviceDataFormatDataFormatAstc8X8SrgbBlock                     RenderingDeviceDataFormat = 171
	RenderingDeviceDataFormatDataFormatAstc10X5UnormBlock                   RenderingDeviceDataFormat = 172
	RenderingDeviceDataFormatDataFormatAstc10X5SrgbBlock                    RenderingDeviceDataFormat = 173
	RenderingDeviceDataFormatDataFormatAstc10X6UnormBlock                   RenderingDeviceDataFormat = 174
	RenderingDeviceDataFormatDataFormatAstc10X6SrgbBlock                    RenderingDeviceDataFormat = 175
	RenderingDeviceDataFormatDataFormatAstc10X8UnormBlock                   RenderingDeviceDataFormat = 176
	RenderingDeviceDataFormatDataFormatAstc10X8SrgbBlock                    RenderingDeviceDataFormat = 177
	RenderingDeviceDataFormatDataFormatAstc10X10UnormBlock                  RenderingDeviceDataFormat = 178
	RenderingDeviceDataFormatDataFormatAstc10X10SrgbBlock                   RenderingDeviceDataFormat = 179
	RenderingDeviceDataFormatDataFormatAstc12X10UnormBlock                  RenderingDeviceDataFormat = 180
	RenderingDeviceDataFormatDataFormatAstc12X10SrgbBlock                   RenderingDeviceDataFormat = 181
	RenderingDeviceDataFormatDataFormatAstc12X12UnormBlock                  RenderingDeviceDataFormat = 182
	RenderingDeviceDataFormatDataFormatAstc12X12SrgbBlock                   RenderingDeviceDataFormat = 183
	RenderingDeviceDataFormatDataFormatG8B8G8R8422Unorm                     RenderingDeviceDataFormat = 184
	RenderingDeviceDataFormatDataFormatB8G8R8G8422Unorm                     RenderingDeviceDataFormat = 185
	RenderingDeviceDataFormatDataFormatG8B8R83Plane420Unorm                 RenderingDeviceDataFormat = 186
	RenderingDeviceDataFormatDataFormatG8B8R82Plane420Unorm                 RenderingDeviceDataFormat = 187
	RenderingDeviceDataFormatDataFormatG8B8R83Plane422Unorm                 RenderingDeviceDataFormat = 188
	RenderingDeviceDataFormatDataFormatG8B8R82Plane422Unorm                 RenderingDeviceDataFormat = 189
	RenderingDeviceDataFormatDataFormatG8B8R83Plane444Unorm                 RenderingDeviceDataFormat = 190
	RenderingDeviceDataFormatDataFormatR10X6UnormPack16                     RenderingDeviceDataFormat = 191
	RenderingDeviceDataFormatDataFormatR10X6G10X6Unorm2Pack16               RenderingDeviceDataFormat = 192
	RenderingDeviceDataFormatDataFormatR10X6G10X6B10X6A10X6Unorm4Pack16     RenderingDeviceDataFormat = 193
	RenderingDeviceDataFormatDataFormatG10X6B10X6G10X6R10X6422Unorm4Pack16  RenderingDeviceDataFormat = 194
	RenderingDeviceDataFormatDataFormatB10X6G10X6R10X6G10X6422Unorm4Pack16  RenderingDeviceDataFormat = 195
	RenderingDeviceDataFormatDataFormatG10X6B10X6R10X63Plane420Unorm3Pack16 RenderingDeviceDataFormat = 196
	RenderingDeviceDataFormatDataFormatG10X6B10X6R10X62Plane420Unorm3Pack16 RenderingDeviceDataFormat = 197
	RenderingDeviceDataFormatDataFormatG10X6B10X6R10X63Plane422Unorm3Pack16 RenderingDeviceDataFormat = 198
	RenderingDeviceDataFormatDataFormatG10X6B10X6R10X62Plane422Unorm3Pack16 RenderingDeviceDataFormat = 199
	RenderingDeviceDataFormatDataFormatG10X6B10X6R10X63Plane444Unorm3Pack16 RenderingDeviceDataFormat = 200
	RenderingDeviceDataFormatDataFormatR12X4UnormPack16                     RenderingDeviceDataFormat = 201
	RenderingDeviceDataFormatDataFormatR12X4G12X4Unorm2Pack16               RenderingDeviceDataFormat = 202
	RenderingDeviceDataFormatDataFormatR12X4G12X4B12X4A12X4Unorm4Pack16     RenderingDeviceDataFormat = 203
	RenderingDeviceDataFormatDataFormatG12X4B12X4G12X4R12X4422Unorm4Pack16  RenderingDeviceDataFormat = 204
	RenderingDeviceDataFormatDataFormatB12X4G12X4R12X4G12X4422Unorm4Pack16  RenderingDeviceDataFormat = 205
	RenderingDeviceDataFormatDataFormatG12X4B12X4R12X43Plane420Unorm3Pack16 RenderingDeviceDataFormat = 206
	RenderingDeviceDataFormatDataFormatG12X4B12X4R12X42Plane420Unorm3Pack16 RenderingDeviceDataFormat = 207
	RenderingDeviceDataFormatDataFormatG12X4B12X4R12X43Plane422Unorm3Pack16 RenderingDeviceDataFormat = 208
	RenderingDeviceDataFormatDataFormatG12X4B12X4R12X42Plane422Unorm3Pack16 RenderingDeviceDataFormat = 209
	RenderingDeviceDataFormatDataFormatG12X4B12X4R12X43Plane444Unorm3Pack16 RenderingDeviceDataFormat = 210
	RenderingDeviceDataFormatDataFormatG16B16G16R16422Unorm                 RenderingDeviceDataFormat = 211
	RenderingDeviceDataFormatDataFormatB16G16R16G16422Unorm                 RenderingDeviceDataFormat = 212
	RenderingDeviceDataFormatDataFormatG16B16R163Plane420Unorm              RenderingDeviceDataFormat = 213
	RenderingDeviceDataFormatDataFormatG16B16R162Plane420Unorm              RenderingDeviceDataFormat = 214
	RenderingDeviceDataFormatDataFormatG16B16R163Plane422Unorm              RenderingDeviceDataFormat = 215
	RenderingDeviceDataFormatDataFormatG16B16R162Plane422Unorm              RenderingDeviceDataFormat = 216
	RenderingDeviceDataFormatDataFormatG16B16R163Plane444Unorm              RenderingDeviceDataFormat = 217
	RenderingDeviceDataFormatDataFormatMax                                  RenderingDeviceDataFormat = 218
)

type RenderingDeviceBarrierMask int

const (
	RenderingDeviceBarrierMaskBarrierMaskVertex      RenderingDeviceBarrierMask = 1
	RenderingDeviceBarrierMaskBarrierMaskFragment    RenderingDeviceBarrierMask = 8
	RenderingDeviceBarrierMaskBarrierMaskCompute     RenderingDeviceBarrierMask = 2
	RenderingDeviceBarrierMaskBarrierMaskTransfer    RenderingDeviceBarrierMask = 4
	RenderingDeviceBarrierMaskBarrierMaskRaster      RenderingDeviceBarrierMask = 9
	RenderingDeviceBarrierMaskBarrierMaskAllBarriers RenderingDeviceBarrierMask = 32767
	RenderingDeviceBarrierMaskBarrierMaskNoBarrier   RenderingDeviceBarrierMask = 32768
)

type RenderingDeviceTextureType int

const (
	RenderingDeviceTextureTypeTextureType1D        RenderingDeviceTextureType = 0
	RenderingDeviceTextureTypeTextureType2D        RenderingDeviceTextureType = 1
	RenderingDeviceTextureTypeTextureType3D        RenderingDeviceTextureType = 2
	RenderingDeviceTextureTypeTextureTypeCube      RenderingDeviceTextureType = 3
	RenderingDeviceTextureTypeTextureType1DArray   RenderingDeviceTextureType = 4
	RenderingDeviceTextureTypeTextureType2DArray   RenderingDeviceTextureType = 5
	RenderingDeviceTextureTypeTextureTypeCubeArray RenderingDeviceTextureType = 6
	RenderingDeviceTextureTypeTextureTypeMax       RenderingDeviceTextureType = 7
)

type RenderingDeviceTextureSamples int

const (
	RenderingDeviceTextureSamplesTextureSamples1   RenderingDeviceTextureSamples = 0
	RenderingDeviceTextureSamplesTextureSamples2   RenderingDeviceTextureSamples = 1
	RenderingDeviceTextureSamplesTextureSamples4   RenderingDeviceTextureSamples = 2
	RenderingDeviceTextureSamplesTextureSamples8   RenderingDeviceTextureSamples = 3
	RenderingDeviceTextureSamplesTextureSamples16  RenderingDeviceTextureSamples = 4
	RenderingDeviceTextureSamplesTextureSamples32  RenderingDeviceTextureSamples = 5
	RenderingDeviceTextureSamplesTextureSamples64  RenderingDeviceTextureSamples = 6
	RenderingDeviceTextureSamplesTextureSamplesMax RenderingDeviceTextureSamples = 7
)

type RenderingDeviceTextureUsageBits int

const (
	RenderingDeviceTextureUsageBitsTextureUsageSamplingBit               RenderingDeviceTextureUsageBits = 1
	RenderingDeviceTextureUsageBitsTextureUsageColorAttachmentBit        RenderingDeviceTextureUsageBits = 2
	RenderingDeviceTextureUsageBitsTextureUsageDepthStencilAttachmentBit RenderingDeviceTextureUsageBits = 4
	RenderingDeviceTextureUsageBitsTextureUsageStorageBit                RenderingDeviceTextureUsageBits = 8
	RenderingDeviceTextureUsageBitsTextureUsageStorageAtomicBit          RenderingDeviceTextureUsageBits = 16
	RenderingDeviceTextureUsageBitsTextureUsageCpuReadBit                RenderingDeviceTextureUsageBits = 32
	RenderingDeviceTextureUsageBitsTextureUsageCanUpdateBit              RenderingDeviceTextureUsageBits = 64
	RenderingDeviceTextureUsageBitsTextureUsageCanCopyFromBit            RenderingDeviceTextureUsageBits = 128
	RenderingDeviceTextureUsageBitsTextureUsageCanCopyToBit              RenderingDeviceTextureUsageBits = 256
	RenderingDeviceTextureUsageBitsTextureUsageInputAttachmentBit        RenderingDeviceTextureUsageBits = 512
)

type RenderingDeviceTextureSwizzle int

const (
	RenderingDeviceTextureSwizzleTextureSwizzleIdentity RenderingDeviceTextureSwizzle = 0
	RenderingDeviceTextureSwizzleTextureSwizzleZero     RenderingDeviceTextureSwizzle = 1
	RenderingDeviceTextureSwizzleTextureSwizzleOne      RenderingDeviceTextureSwizzle = 2
	RenderingDeviceTextureSwizzleTextureSwizzleR        RenderingDeviceTextureSwizzle = 3
	RenderingDeviceTextureSwizzleTextureSwizzleG        RenderingDeviceTextureSwizzle = 4
	RenderingDeviceTextureSwizzleTextureSwizzleB        RenderingDeviceTextureSwizzle = 5
	RenderingDeviceTextureSwizzleTextureSwizzleA        RenderingDeviceTextureSwizzle = 6
	RenderingDeviceTextureSwizzleTextureSwizzleMax      RenderingDeviceTextureSwizzle = 7
)

type RenderingDeviceTextureSliceType int

const (
	RenderingDeviceTextureSliceTypeTextureSlice2D      RenderingDeviceTextureSliceType = 0
	RenderingDeviceTextureSliceTypeTextureSliceCubemap RenderingDeviceTextureSliceType = 1
	RenderingDeviceTextureSliceTypeTextureSlice3D      RenderingDeviceTextureSliceType = 2
)

type RenderingDeviceSamplerFilter int

const (
	RenderingDeviceSamplerFilterSamplerFilterNearest RenderingDeviceSamplerFilter = 0
	RenderingDeviceSamplerFilterSamplerFilterLinear  RenderingDeviceSamplerFilter = 1
)

type RenderingDeviceSamplerRepeatMode int

const (
	RenderingDeviceSamplerRepeatModeSamplerRepeatModeRepeat            RenderingDeviceSamplerRepeatMode = 0
	RenderingDeviceSamplerRepeatModeSamplerRepeatModeMirroredRepeat    RenderingDeviceSamplerRepeatMode = 1
	RenderingDeviceSamplerRepeatModeSamplerRepeatModeClampToEdge       RenderingDeviceSamplerRepeatMode = 2
	RenderingDeviceSamplerRepeatModeSamplerRepeatModeClampToBorder     RenderingDeviceSamplerRepeatMode = 3
	RenderingDeviceSamplerRepeatModeSamplerRepeatModeMirrorClampToEdge RenderingDeviceSamplerRepeatMode = 4
	RenderingDeviceSamplerRepeatModeSamplerRepeatModeMax               RenderingDeviceSamplerRepeatMode = 5
)

type RenderingDeviceSamplerBorderColor int

const (
	RenderingDeviceSamplerBorderColorSamplerBorderColorFloatTransparentBlack RenderingDeviceSamplerBorderColor = 0
	RenderingDeviceSamplerBorderColorSamplerBorderColorIntTransparentBlack   RenderingDeviceSamplerBorderColor = 1
	RenderingDeviceSamplerBorderColorSamplerBorderColorFloatOpaqueBlack      RenderingDeviceSamplerBorderColor = 2
	RenderingDeviceSamplerBorderColorSamplerBorderColorIntOpaqueBlack        RenderingDeviceSamplerBorderColor = 3
	RenderingDeviceSamplerBorderColorSamplerBorderColorFloatOpaqueWhite      RenderingDeviceSamplerBorderColor = 4
	RenderingDeviceSamplerBorderColorSamplerBorderColorIntOpaqueWhite        RenderingDeviceSamplerBorderColor = 5
	RenderingDeviceSamplerBorderColorSamplerBorderColorMax                   RenderingDeviceSamplerBorderColor = 6
)

type RenderingDeviceVertexFrequency int

const (
	RenderingDeviceVertexFrequencyVertexFrequencyVertex   RenderingDeviceVertexFrequency = 0
	RenderingDeviceVertexFrequencyVertexFrequencyInstance RenderingDeviceVertexFrequency = 1
)

type RenderingDeviceIndexBufferFormat int

const (
	RenderingDeviceIndexBufferFormatIndexBufferFormatUint16 RenderingDeviceIndexBufferFormat = 0
	RenderingDeviceIndexBufferFormatIndexBufferFormatUint32 RenderingDeviceIndexBufferFormat = 1
)

type RenderingDeviceStorageBufferUsage int

const (
	RenderingDeviceStorageBufferUsageStorageBufferUsageDispatchIndirect RenderingDeviceStorageBufferUsage = 1
)

type RenderingDeviceUniformType int

const (
	RenderingDeviceUniformTypeUniformTypeSampler                  RenderingDeviceUniformType = 0
	RenderingDeviceUniformTypeUniformTypeSamplerWithTexture       RenderingDeviceUniformType = 1
	RenderingDeviceUniformTypeUniformTypeTexture                  RenderingDeviceUniformType = 2
	RenderingDeviceUniformTypeUniformTypeImage                    RenderingDeviceUniformType = 3
	RenderingDeviceUniformTypeUniformTypeTextureBuffer            RenderingDeviceUniformType = 4
	RenderingDeviceUniformTypeUniformTypeSamplerWithTextureBuffer RenderingDeviceUniformType = 5
	RenderingDeviceUniformTypeUniformTypeImageBuffer              RenderingDeviceUniformType = 6
	RenderingDeviceUniformTypeUniformTypeUniformBuffer            RenderingDeviceUniformType = 7
	RenderingDeviceUniformTypeUniformTypeStorageBuffer            RenderingDeviceUniformType = 8
	RenderingDeviceUniformTypeUniformTypeInputAttachment          RenderingDeviceUniformType = 9
	RenderingDeviceUniformTypeUniformTypeMax                      RenderingDeviceUniformType = 10
)

type RenderingDeviceRenderPrimitive int

const (
	RenderingDeviceRenderPrimitiveRenderPrimitivePoints                         RenderingDeviceRenderPrimitive = 0
	RenderingDeviceRenderPrimitiveRenderPrimitiveLines                          RenderingDeviceRenderPrimitive = 1
	RenderingDeviceRenderPrimitiveRenderPrimitiveLinesWithAdjacency             RenderingDeviceRenderPrimitive = 2
	RenderingDeviceRenderPrimitiveRenderPrimitiveLinestrips                     RenderingDeviceRenderPrimitive = 3
	RenderingDeviceRenderPrimitiveRenderPrimitiveLinestripsWithAdjacency        RenderingDeviceRenderPrimitive = 4
	RenderingDeviceRenderPrimitiveRenderPrimitiveTriangles                      RenderingDeviceRenderPrimitive = 5
	RenderingDeviceRenderPrimitiveRenderPrimitiveTrianglesWithAdjacency         RenderingDeviceRenderPrimitive = 6
	RenderingDeviceRenderPrimitiveRenderPrimitiveTriangleStrips                 RenderingDeviceRenderPrimitive = 7
	RenderingDeviceRenderPrimitiveRenderPrimitiveTriangleStripsWithAjacency     RenderingDeviceRenderPrimitive = 8
	RenderingDeviceRenderPrimitiveRenderPrimitiveTriangleStripsWithRestartIndex RenderingDeviceRenderPrimitive = 9
	RenderingDeviceRenderPrimitiveRenderPrimitiveTesselationPatch               RenderingDeviceRenderPrimitive = 10
	RenderingDeviceRenderPrimitiveRenderPrimitiveMax                            RenderingDeviceRenderPrimitive = 11
)

type RenderingDevicePolygonCullMode int

const (
	RenderingDevicePolygonCullModePolygonCullDisabled RenderingDevicePolygonCullMode = 0
	RenderingDevicePolygonCullModePolygonCullFront    RenderingDevicePolygonCullMode = 1
	RenderingDevicePolygonCullModePolygonCullBack     RenderingDevicePolygonCullMode = 2
)

type RenderingDevicePolygonFrontFace int

const (
	RenderingDevicePolygonFrontFacePolygonFrontFaceClockwise        RenderingDevicePolygonFrontFace = 0
	RenderingDevicePolygonFrontFacePolygonFrontFaceCounterClockwise RenderingDevicePolygonFrontFace = 1
)

type RenderingDeviceStencilOperation int

const (
	RenderingDeviceStencilOperationStencilOpKeep              RenderingDeviceStencilOperation = 0
	RenderingDeviceStencilOperationStencilOpZero              RenderingDeviceStencilOperation = 1
	RenderingDeviceStencilOperationStencilOpReplace           RenderingDeviceStencilOperation = 2
	RenderingDeviceStencilOperationStencilOpIncrementAndClamp RenderingDeviceStencilOperation = 3
	RenderingDeviceStencilOperationStencilOpDecrementAndClamp RenderingDeviceStencilOperation = 4
	RenderingDeviceStencilOperationStencilOpInvert            RenderingDeviceStencilOperation = 5
	RenderingDeviceStencilOperationStencilOpIncrementAndWrap  RenderingDeviceStencilOperation = 6
	RenderingDeviceStencilOperationStencilOpDecrementAndWrap  RenderingDeviceStencilOperation = 7
	RenderingDeviceStencilOperationStencilOpMax               RenderingDeviceStencilOperation = 8
)

type RenderingDeviceCompareOperator int

const (
	RenderingDeviceCompareOperatorCompareOpNever          RenderingDeviceCompareOperator = 0
	RenderingDeviceCompareOperatorCompareOpLess           RenderingDeviceCompareOperator = 1
	RenderingDeviceCompareOperatorCompareOpEqual          RenderingDeviceCompareOperator = 2
	RenderingDeviceCompareOperatorCompareOpLessOrEqual    RenderingDeviceCompareOperator = 3
	RenderingDeviceCompareOperatorCompareOpGreater        RenderingDeviceCompareOperator = 4
	RenderingDeviceCompareOperatorCompareOpNotEqual       RenderingDeviceCompareOperator = 5
	RenderingDeviceCompareOperatorCompareOpGreaterOrEqual RenderingDeviceCompareOperator = 6
	RenderingDeviceCompareOperatorCompareOpAlways         RenderingDeviceCompareOperator = 7
	RenderingDeviceCompareOperatorCompareOpMax            RenderingDeviceCompareOperator = 8
)

type RenderingDeviceLogicOperation int

const (
	RenderingDeviceLogicOperationLogicOpClear        RenderingDeviceLogicOperation = 0
	RenderingDeviceLogicOperationLogicOpAnd          RenderingDeviceLogicOperation = 1
	RenderingDeviceLogicOperationLogicOpAndReverse   RenderingDeviceLogicOperation = 2
	RenderingDeviceLogicOperationLogicOpCopy         RenderingDeviceLogicOperation = 3
	RenderingDeviceLogicOperationLogicOpAndInverted  RenderingDeviceLogicOperation = 4
	RenderingDeviceLogicOperationLogicOpNoOp         RenderingDeviceLogicOperation = 5
	RenderingDeviceLogicOperationLogicOpXor          RenderingDeviceLogicOperation = 6
	RenderingDeviceLogicOperationLogicOpOr           RenderingDeviceLogicOperation = 7
	RenderingDeviceLogicOperationLogicOpNor          RenderingDeviceLogicOperation = 8
	RenderingDeviceLogicOperationLogicOpEquivalent   RenderingDeviceLogicOperation = 9
	RenderingDeviceLogicOperationLogicOpInvert       RenderingDeviceLogicOperation = 10
	RenderingDeviceLogicOperationLogicOpOrReverse    RenderingDeviceLogicOperation = 11
	RenderingDeviceLogicOperationLogicOpCopyInverted RenderingDeviceLogicOperation = 12
	RenderingDeviceLogicOperationLogicOpOrInverted   RenderingDeviceLogicOperation = 13
	RenderingDeviceLogicOperationLogicOpNand         RenderingDeviceLogicOperation = 14
	RenderingDeviceLogicOperationLogicOpSet          RenderingDeviceLogicOperation = 15
	RenderingDeviceLogicOperationLogicOpMax          RenderingDeviceLogicOperation = 16
)

type RenderingDeviceBlendFactor int

const (
	RenderingDeviceBlendFactorBlendFactorZero                  RenderingDeviceBlendFactor = 0
	RenderingDeviceBlendFactorBlendFactorOne                   RenderingDeviceBlendFactor = 1
	RenderingDeviceBlendFactorBlendFactorSrcColor              RenderingDeviceBlendFactor = 2
	RenderingDeviceBlendFactorBlendFactorOneMinusSrcColor      RenderingDeviceBlendFactor = 3
	RenderingDeviceBlendFactorBlendFactorDstColor              RenderingDeviceBlendFactor = 4
	RenderingDeviceBlendFactorBlendFactorOneMinusDstColor      RenderingDeviceBlendFactor = 5
	RenderingDeviceBlendFactorBlendFactorSrcAlpha              RenderingDeviceBlendFactor = 6
	RenderingDeviceBlendFactorBlendFactorOneMinusSrcAlpha      RenderingDeviceBlendFactor = 7
	RenderingDeviceBlendFactorBlendFactorDstAlpha              RenderingDeviceBlendFactor = 8
	RenderingDeviceBlendFactorBlendFactorOneMinusDstAlpha      RenderingDeviceBlendFactor = 9
	RenderingDeviceBlendFactorBlendFactorConstantColor         RenderingDeviceBlendFactor = 10
	RenderingDeviceBlendFactorBlendFactorOneMinusConstantColor RenderingDeviceBlendFactor = 11
	RenderingDeviceBlendFactorBlendFactorConstantAlpha         RenderingDeviceBlendFactor = 12
	RenderingDeviceBlendFactorBlendFactorOneMinusConstantAlpha RenderingDeviceBlendFactor = 13
	RenderingDeviceBlendFactorBlendFactorSrcAlphaSaturate      RenderingDeviceBlendFactor = 14
	RenderingDeviceBlendFactorBlendFactorSrc1Color             RenderingDeviceBlendFactor = 15
	RenderingDeviceBlendFactorBlendFactorOneMinusSrc1Color     RenderingDeviceBlendFactor = 16
	RenderingDeviceBlendFactorBlendFactorSrc1Alpha             RenderingDeviceBlendFactor = 17
	RenderingDeviceBlendFactorBlendFactorOneMinusSrc1Alpha     RenderingDeviceBlendFactor = 18
	RenderingDeviceBlendFactorBlendFactorMax                   RenderingDeviceBlendFactor = 19
)

type RenderingDeviceBlendOperation int

const (
	RenderingDeviceBlendOperationBlendOpAdd             RenderingDeviceBlendOperation = 0
	RenderingDeviceBlendOperationBlendOpSubtract        RenderingDeviceBlendOperation = 1
	RenderingDeviceBlendOperationBlendOpReverseSubtract RenderingDeviceBlendOperation = 2
	RenderingDeviceBlendOperationBlendOpMinimum         RenderingDeviceBlendOperation = 3
	RenderingDeviceBlendOperationBlendOpMaximum         RenderingDeviceBlendOperation = 4
	RenderingDeviceBlendOperationBlendOpMax             RenderingDeviceBlendOperation = 5
)

type RenderingDevicePipelineDynamicStateFlags int

const (
	RenderingDevicePipelineDynamicStateFlagsDynamicStateLineWidth          RenderingDevicePipelineDynamicStateFlags = 1
	RenderingDevicePipelineDynamicStateFlagsDynamicStateDepthBias          RenderingDevicePipelineDynamicStateFlags = 2
	RenderingDevicePipelineDynamicStateFlagsDynamicStateBlendConstants     RenderingDevicePipelineDynamicStateFlags = 4
	RenderingDevicePipelineDynamicStateFlagsDynamicStateDepthBounds        RenderingDevicePipelineDynamicStateFlags = 8
	RenderingDevicePipelineDynamicStateFlagsDynamicStateStencilCompareMask RenderingDevicePipelineDynamicStateFlags = 16
	RenderingDevicePipelineDynamicStateFlagsDynamicStateStencilWriteMask   RenderingDevicePipelineDynamicStateFlags = 32
	RenderingDevicePipelineDynamicStateFlagsDynamicStateStencilReference   RenderingDevicePipelineDynamicStateFlags = 64
)

type RenderingDeviceInitialAction int

const (
	RenderingDeviceInitialActionInitialActionLoad                RenderingDeviceInitialAction = 0
	RenderingDeviceInitialActionInitialActionClear               RenderingDeviceInitialAction = 1
	RenderingDeviceInitialActionInitialActionDiscard             RenderingDeviceInitialAction = 2
	RenderingDeviceInitialActionInitialActionMax                 RenderingDeviceInitialAction = 3
	RenderingDeviceInitialActionInitialActionClearRegion         RenderingDeviceInitialAction = 1
	RenderingDeviceInitialActionInitialActionClearRegionContinue RenderingDeviceInitialAction = 1
	RenderingDeviceInitialActionInitialActionKeep                RenderingDeviceInitialAction = 0
	RenderingDeviceInitialActionInitialActionDrop                RenderingDeviceInitialAction = 2
	RenderingDeviceInitialActionInitialActionContinue            RenderingDeviceInitialAction = 0
)

type RenderingDeviceFinalAction int

const (
	RenderingDeviceFinalActionFinalActionStore    RenderingDeviceFinalAction = 0
	RenderingDeviceFinalActionFinalActionDiscard  RenderingDeviceFinalAction = 1
	RenderingDeviceFinalActionFinalActionMax      RenderingDeviceFinalAction = 2
	RenderingDeviceFinalActionFinalActionRead     RenderingDeviceFinalAction = 0
	RenderingDeviceFinalActionFinalActionContinue RenderingDeviceFinalAction = 0
)

type RenderingDeviceShaderStage int

const (
	RenderingDeviceShaderStageShaderStageVertex                   RenderingDeviceShaderStage = 0
	RenderingDeviceShaderStageShaderStageFragment                 RenderingDeviceShaderStage = 1
	RenderingDeviceShaderStageShaderStageTesselationControl       RenderingDeviceShaderStage = 2
	RenderingDeviceShaderStageShaderStageTesselationEvaluation    RenderingDeviceShaderStage = 3
	RenderingDeviceShaderStageShaderStageCompute                  RenderingDeviceShaderStage = 4
	RenderingDeviceShaderStageShaderStageMax                      RenderingDeviceShaderStage = 5
	RenderingDeviceShaderStageShaderStageVertexBit                RenderingDeviceShaderStage = 1
	RenderingDeviceShaderStageShaderStageFragmentBit              RenderingDeviceShaderStage = 2
	RenderingDeviceShaderStageShaderStageTesselationControlBit    RenderingDeviceShaderStage = 4
	RenderingDeviceShaderStageShaderStageTesselationEvaluationBit RenderingDeviceShaderStage = 8
	RenderingDeviceShaderStageShaderStageComputeBit               RenderingDeviceShaderStage = 16
)

type RenderingDeviceShaderLanguage int

const (
	RenderingDeviceShaderLanguageShaderLanguageGlsl RenderingDeviceShaderLanguage = 0
	RenderingDeviceShaderLanguageShaderLanguageHlsl RenderingDeviceShaderLanguage = 1
)

type RenderingDevicePipelineSpecializationConstantType int

const (
	RenderingDevicePipelineSpecializationConstantTypePipelineSpecializationConstantTypeBool  RenderingDevicePipelineSpecializationConstantType = 0
	RenderingDevicePipelineSpecializationConstantTypePipelineSpecializationConstantTypeInt   RenderingDevicePipelineSpecializationConstantType = 1
	RenderingDevicePipelineSpecializationConstantTypePipelineSpecializationConstantTypeFloat RenderingDevicePipelineSpecializationConstantType = 2
)

type RenderingDeviceLimit int

const (
	RenderingDeviceLimitLimitMaxBoundUniformSets             RenderingDeviceLimit = 0
	RenderingDeviceLimitLimitMaxFramebufferColorAttachments  RenderingDeviceLimit = 1
	RenderingDeviceLimitLimitMaxTexturesPerUniformSet        RenderingDeviceLimit = 2
	RenderingDeviceLimitLimitMaxSamplersPerUniformSet        RenderingDeviceLimit = 3
	RenderingDeviceLimitLimitMaxStorageBuffersPerUniformSet  RenderingDeviceLimit = 4
	RenderingDeviceLimitLimitMaxStorageImagesPerUniformSet   RenderingDeviceLimit = 5
	RenderingDeviceLimitLimitMaxUniformBuffersPerUniformSet  RenderingDeviceLimit = 6
	RenderingDeviceLimitLimitMaxDrawIndexedIndex             RenderingDeviceLimit = 7
	RenderingDeviceLimitLimitMaxFramebufferHeight            RenderingDeviceLimit = 8
	RenderingDeviceLimitLimitMaxFramebufferWidth             RenderingDeviceLimit = 9
	RenderingDeviceLimitLimitMaxTextureArrayLayers           RenderingDeviceLimit = 10
	RenderingDeviceLimitLimitMaxTextureSize1D                RenderingDeviceLimit = 11
	RenderingDeviceLimitLimitMaxTextureSize2D                RenderingDeviceLimit = 12
	RenderingDeviceLimitLimitMaxTextureSize3D                RenderingDeviceLimit = 13
	RenderingDeviceLimitLimitMaxTextureSizeCube              RenderingDeviceLimit = 14
	RenderingDeviceLimitLimitMaxTexturesPerShaderStage       RenderingDeviceLimit = 15
	RenderingDeviceLimitLimitMaxSamplersPerShaderStage       RenderingDeviceLimit = 16
	RenderingDeviceLimitLimitMaxStorageBuffersPerShaderStage RenderingDeviceLimit = 17
	RenderingDeviceLimitLimitMaxStorageImagesPerShaderStage  RenderingDeviceLimit = 18
	RenderingDeviceLimitLimitMaxUniformBuffersPerShaderStage RenderingDeviceLimit = 19
	RenderingDeviceLimitLimitMaxPushConstantSize             RenderingDeviceLimit = 20
	RenderingDeviceLimitLimitMaxUniformBufferSize            RenderingDeviceLimit = 21
	RenderingDeviceLimitLimitMaxVertexInputAttributeOffset   RenderingDeviceLimit = 22
	RenderingDeviceLimitLimitMaxVertexInputAttributes        RenderingDeviceLimit = 23
	RenderingDeviceLimitLimitMaxVertexInputBindings          RenderingDeviceLimit = 24
	RenderingDeviceLimitLimitMaxVertexInputBindingStride     RenderingDeviceLimit = 25
	RenderingDeviceLimitLimitMinUniformBufferOffsetAlignment RenderingDeviceLimit = 26
	RenderingDeviceLimitLimitMaxComputeSharedMemorySize      RenderingDeviceLimit = 27
	RenderingDeviceLimitLimitMaxComputeWorkgroupCountX       RenderingDeviceLimit = 28
	RenderingDeviceLimitLimitMaxComputeWorkgroupCountY       RenderingDeviceLimit = 29
	RenderingDeviceLimitLimitMaxComputeWorkgroupCountZ       RenderingDeviceLimit = 30
	RenderingDeviceLimitLimitMaxComputeWorkgroupInvocations  RenderingDeviceLimit = 31
	RenderingDeviceLimitLimitMaxComputeWorkgroupSizeX        RenderingDeviceLimit = 32
	RenderingDeviceLimitLimitMaxComputeWorkgroupSizeY        RenderingDeviceLimit = 33
	RenderingDeviceLimitLimitMaxComputeWorkgroupSizeZ        RenderingDeviceLimit = 34
	RenderingDeviceLimitLimitMaxViewportDimensionsX          RenderingDeviceLimit = 35
	RenderingDeviceLimitLimitMaxViewportDimensionsY          RenderingDeviceLimit = 36
)

type RenderingDeviceMemoryType int

const (
	RenderingDeviceMemoryTypeMemoryTextures RenderingDeviceMemoryType = 0
	RenderingDeviceMemoryTypeMemoryBuffers  RenderingDeviceMemoryType = 1
	RenderingDeviceMemoryTypeMemoryTotal    RenderingDeviceMemoryType = 2
)

func (me *RenderingDevice) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *RenderingDevice) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *RenderingDevice) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *RenderingDevice) TextureCreate(format RDTextureFormat, view RDTextureView, data []PackedByteArray) RID {
	cargs := []gdc.ConstTypePtr{format.AsCTypePtr(), view.AsCTypePtr(), gdc.ConstTypePtr(&data)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&data)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) TextureCreateShared(view RDTextureView, with_texture RID) RID {
	cargs := []gdc.ConstTypePtr{view.AsCTypePtr(), with_texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureCreateShared), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) TextureCreateSharedFromSlice(view RDTextureView, with_texture RID, layer int64, mipmap int64, mipmaps int64, slice_type RenderingDeviceTextureSliceType) RID {
	cargs := []gdc.ConstTypePtr{view.AsCTypePtr(), with_texture.AsCTypePtr(), gdc.ConstTypePtr(&layer), gdc.ConstTypePtr(&mipmap), gdc.ConstTypePtr(&mipmaps), gdc.ConstTypePtr(&slice_type)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&layer)
	pinner.Pin(&mipmap)
	pinner.Pin(&mipmaps)
	pinner.Pin(&slice_type)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureCreateSharedFromSlice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) TextureCreateFromExtension(type_ RenderingDeviceTextureType, format RenderingDeviceDataFormat, samples RenderingDeviceTextureSamples, usage_flags RenderingDeviceTextureUsageBits, image int64, width int64, height int64, depth int64, layers int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_), gdc.ConstTypePtr(&format), gdc.ConstTypePtr(&samples), gdc.ConstTypePtr(&usage_flags), gdc.ConstTypePtr(&image), gdc.ConstTypePtr(&width), gdc.ConstTypePtr(&height), gdc.ConstTypePtr(&depth), gdc.ConstTypePtr(&layers)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&type_)
	pinner.Pin(&format)
	pinner.Pin(&samples)
	pinner.Pin(&usage_flags)
	pinner.Pin(&image)
	pinner.Pin(&width)
	pinner.Pin(&height)
	pinner.Pin(&depth)
	pinner.Pin(&layers)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureCreateFromExtension), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) TextureUpdate(texture RID, layer int64, data PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), gdc.ConstTypePtr(&layer), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureUpdate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RenderingDevice) TextureGetData(texture RID, layer int64) PackedByteArray {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), gdc.ConstTypePtr(&layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()
	pinner.Pin(&layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) TextureIsFormatSupportedForUsage(format RenderingDeviceDataFormat, usage_flags RenderingDeviceTextureUsageBits) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format), gdc.ConstTypePtr(&usage_flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&format)
	pinner.Pin(&usage_flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureIsFormatSupportedForUsage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) TextureIsShared(texture RID) bool {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureIsShared), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) TextureIsValid(texture RID) bool {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureIsValid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) TextureCopy(from_texture RID, to_texture RID, from_pos Vector3, to_pos Vector3, size Vector3, src_mipmap int64, dst_mipmap int64, src_layer int64, dst_layer int64) Error {
	cargs := []gdc.ConstTypePtr{from_texture.AsCTypePtr(), to_texture.AsCTypePtr(), from_pos.AsCTypePtr(), to_pos.AsCTypePtr(), size.AsCTypePtr(), gdc.ConstTypePtr(&src_mipmap), gdc.ConstTypePtr(&dst_mipmap), gdc.ConstTypePtr(&src_layer), gdc.ConstTypePtr(&dst_layer)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&src_mipmap)
	pinner.Pin(&dst_mipmap)
	pinner.Pin(&src_layer)
	pinner.Pin(&dst_layer)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureCopy), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RenderingDevice) TextureClear(texture RID, color Color, base_mipmap int64, mipmap_count int64, base_layer int64, layer_count int64) Error {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr(), color.AsCTypePtr(), gdc.ConstTypePtr(&base_mipmap), gdc.ConstTypePtr(&mipmap_count), gdc.ConstTypePtr(&base_layer), gdc.ConstTypePtr(&layer_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&base_mipmap)
	pinner.Pin(&mipmap_count)
	pinner.Pin(&base_layer)
	pinner.Pin(&layer_count)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureClear), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RenderingDevice) TextureResolveMultisample(from_texture RID, to_texture RID) Error {
	cargs := []gdc.ConstTypePtr{from_texture.AsCTypePtr(), to_texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureResolveMultisample), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RenderingDevice) TextureGetFormat(texture RID) RDTextureFormat {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRDTextureFormat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureGetFormat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) TextureGetNativeHandle(texture RID) int64 {
	cargs := []gdc.ConstTypePtr{texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureGetNativeHandle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) FramebufferFormatCreate(attachments []RDAttachmentFormat, view_count int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&attachments), gdc.ConstTypePtr(&view_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&attachments)
	pinner.Pin(&view_count)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFramebufferFormatCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) FramebufferFormatCreateMultipass(attachments []RDAttachmentFormat, passes []RDFramebufferPass, view_count int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&attachments), gdc.ConstTypePtr(&passes), gdc.ConstTypePtr(&view_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&attachments)
	pinner.Pin(&passes)
	pinner.Pin(&view_count)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFramebufferFormatCreateMultipass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) FramebufferFormatCreateEmpty(samples RenderingDeviceTextureSamples) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&samples)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&samples)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFramebufferFormatCreateEmpty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) FramebufferFormatGetTextureSamples(format int64, render_pass int64) RenderingDeviceTextureSamples {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format), gdc.ConstTypePtr(&render_pass)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret RenderingDeviceTextureSamples
	pinner.Pin(&format)
	pinner.Pin(&render_pass)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFramebufferFormatGetTextureSamples), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RenderingDevice) FramebufferCreate(textures []RID, validate_with_format int64, view_count int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&textures), gdc.ConstTypePtr(&validate_with_format), gdc.ConstTypePtr(&view_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&textures)
	pinner.Pin(&validate_with_format)
	pinner.Pin(&view_count)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFramebufferCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) FramebufferCreateMultipass(textures []RID, passes []RDFramebufferPass, validate_with_format int64, view_count int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&textures), gdc.ConstTypePtr(&passes), gdc.ConstTypePtr(&validate_with_format), gdc.ConstTypePtr(&view_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&textures)
	pinner.Pin(&passes)
	pinner.Pin(&validate_with_format)
	pinner.Pin(&view_count)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFramebufferCreateMultipass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) FramebufferCreateEmpty(size Vector2i, samples RenderingDeviceTextureSamples, validate_with_format int64) RID {
	cargs := []gdc.ConstTypePtr{size.AsCTypePtr(), gdc.ConstTypePtr(&samples), gdc.ConstTypePtr(&validate_with_format)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&samples)
	pinner.Pin(&validate_with_format)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFramebufferCreateEmpty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) FramebufferGetFormat(framebuffer RID) int64 {
	cargs := []gdc.ConstTypePtr{framebuffer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFramebufferGetFormat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) FramebufferIsValid(framebuffer RID) bool {
	cargs := []gdc.ConstTypePtr{framebuffer.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFramebufferIsValid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) SamplerCreate(state RDSamplerState) RID {
	cargs := []gdc.ConstTypePtr{state.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnSamplerCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) SamplerIsFormatSupportedForFilter(format RenderingDeviceDataFormat, sampler_filter RenderingDeviceSamplerFilter) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format), gdc.ConstTypePtr(&sampler_filter)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&format)
	pinner.Pin(&sampler_filter)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnSamplerIsFormatSupportedForFilter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) VertexBufferCreate(size_bytes int64, data PackedByteArray, use_as_storage bool) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_bytes), data.AsCTypePtr(), gdc.ConstTypePtr(&use_as_storage)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&size_bytes)
	pinner.Pin(&use_as_storage)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnVertexBufferCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) VertexFormatCreate(vertex_descriptions []RDVertexAttribute) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertex_descriptions)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&vertex_descriptions)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnVertexFormatCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) VertexArrayCreate(vertex_count int64, vertex_format int64, src_buffers []RID, offsets PackedInt64Array) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vertex_count), gdc.ConstTypePtr(&vertex_format), gdc.ConstTypePtr(&src_buffers), offsets.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&vertex_count)
	pinner.Pin(&vertex_format)
	pinner.Pin(&src_buffers)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnVertexArrayCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) IndexBufferCreate(size_indices int64, format RenderingDeviceIndexBufferFormat, data PackedByteArray, use_restart_indices bool) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_indices), gdc.ConstTypePtr(&format), data.AsCTypePtr(), gdc.ConstTypePtr(&use_restart_indices)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&size_indices)
	pinner.Pin(&format)
	pinner.Pin(&use_restart_indices)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnIndexBufferCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) IndexArrayCreate(index_buffer RID, index_offset int64, index_count int64) RID {
	cargs := []gdc.ConstTypePtr{index_buffer.AsCTypePtr(), gdc.ConstTypePtr(&index_offset), gdc.ConstTypePtr(&index_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&index_offset)
	pinner.Pin(&index_count)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnIndexArrayCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) ShaderCompileSpirvFromSource(shader_source RDShaderSource, allow_cache bool) RDShaderSPIRV {
	cargs := []gdc.ConstTypePtr{shader_source.AsCTypePtr(), gdc.ConstTypePtr(&allow_cache)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRDShaderSPIRV()
	pinner.Pin(&allow_cache)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnShaderCompileSpirvFromSource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) ShaderCompileBinaryFromSpirv(spirv_data RDShaderSPIRV, name String) PackedByteArray {
	cargs := []gdc.ConstTypePtr{spirv_data.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnShaderCompileBinaryFromSpirv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) ShaderCreateFromSpirv(spirv_data RDShaderSPIRV, name String) RID {
	cargs := []gdc.ConstTypePtr{spirv_data.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnShaderCreateFromSpirv), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) ShaderCreateFromBytecode(binary_data PackedByteArray, placeholder_rid RID) RID {
	cargs := []gdc.ConstTypePtr{binary_data.AsCTypePtr(), placeholder_rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnShaderCreateFromBytecode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) ShaderCreatePlaceholder() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnShaderCreatePlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) ShaderGetVertexInputAttributeMask(shader RID) int64 {
	cargs := []gdc.ConstTypePtr{shader.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnShaderGetVertexInputAttributeMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) UniformBufferCreate(size_bytes int64, data PackedByteArray) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_bytes), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&size_bytes)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnUniformBufferCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) StorageBufferCreate(size_bytes int64, data PackedByteArray, usage RenderingDeviceStorageBufferUsage) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_bytes), data.AsCTypePtr(), gdc.ConstTypePtr(&usage)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&size_bytes)
	pinner.Pin(&usage)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnStorageBufferCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) TextureBufferCreate(size_bytes int64, format RenderingDeviceDataFormat, data PackedByteArray) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size_bytes), gdc.ConstTypePtr(&format), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&size_bytes)
	pinner.Pin(&format)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnTextureBufferCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) UniformSetCreate(uniforms []RDUniform, shader RID, shader_set int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&uniforms), shader.AsCTypePtr(), gdc.ConstTypePtr(&shader_set)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&uniforms)
	pinner.Pin(&shader_set)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnUniformSetCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) UniformSetIsValid(uniform_set RID) bool {
	cargs := []gdc.ConstTypePtr{uniform_set.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnUniformSetIsValid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) BufferCopy(src_buffer RID, dst_buffer RID, src_offset int64, dst_offset int64, size int64) Error {
	cargs := []gdc.ConstTypePtr{src_buffer.AsCTypePtr(), dst_buffer.AsCTypePtr(), gdc.ConstTypePtr(&src_offset), gdc.ConstTypePtr(&dst_offset), gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&src_offset)
	pinner.Pin(&dst_offset)
	pinner.Pin(&size)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnBufferCopy), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RenderingDevice) BufferUpdate(buffer RID, offset int64, size_bytes int64, data PackedByteArray) Error {
	cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&size_bytes), data.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&offset)
	pinner.Pin(&size_bytes)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnBufferUpdate), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RenderingDevice) BufferClear(buffer RID, offset int64, size_bytes int64) Error {
	cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), gdc.ConstTypePtr(&offset), gdc.ConstTypePtr(&size_bytes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Error
	pinner.Pin(&offset)
	pinner.Pin(&size_bytes)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnBufferClear), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *RenderingDevice) BufferGetData(buffer RID, offset_bytes int64, size_bytes int64) PackedByteArray {
	cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), gdc.ConstTypePtr(&offset_bytes), gdc.ConstTypePtr(&size_bytes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()
	pinner.Pin(&offset_bytes)
	pinner.Pin(&size_bytes)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnBufferGetData), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) RenderPipelineCreate(shader RID, framebuffer_format int64, vertex_format int64, primitive RenderingDeviceRenderPrimitive, rasterization_state RDPipelineRasterizationState, multisample_state RDPipelineMultisampleState, stencil_state RDPipelineDepthStencilState, color_blend_state RDPipelineColorBlendState, dynamic_state_flags RenderingDevicePipelineDynamicStateFlags, for_render_pass int64, specialization_constants []RDPipelineSpecializationConstant) RID {
	cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), gdc.ConstTypePtr(&framebuffer_format), gdc.ConstTypePtr(&vertex_format), gdc.ConstTypePtr(&primitive), rasterization_state.AsCTypePtr(), multisample_state.AsCTypePtr(), stencil_state.AsCTypePtr(), color_blend_state.AsCTypePtr(), gdc.ConstTypePtr(&dynamic_state_flags), gdc.ConstTypePtr(&for_render_pass), gdc.ConstTypePtr(&specialization_constants)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&framebuffer_format)
	pinner.Pin(&vertex_format)
	pinner.Pin(&primitive)
	pinner.Pin(&dynamic_state_flags)
	pinner.Pin(&for_render_pass)
	pinner.Pin(&specialization_constants)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnRenderPipelineCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) RenderPipelineIsValid(render_pipeline RID) bool {
	cargs := []gdc.ConstTypePtr{render_pipeline.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnRenderPipelineIsValid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) ComputePipelineCreate(shader RID, specialization_constants []RDPipelineSpecializationConstant) RID {
	cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), gdc.ConstTypePtr(&specialization_constants)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&specialization_constants)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputePipelineCreate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) ComputePipelineIsValid(compute_pipeline RID) bool {
	cargs := []gdc.ConstTypePtr{compute_pipeline.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputePipelineIsValid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) ScreenGetWidth(screen int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnScreenGetWidth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) ScreenGetHeight(screen int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnScreenGetHeight), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) ScreenGetFramebufferFormat(screen int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnScreenGetFramebufferFormat), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) DrawListBeginForScreen(screen int64, clear_color Color) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&screen), clear_color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&screen)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListBeginForScreen), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) DrawListBegin(framebuffer RID, initial_color_action RenderingDeviceInitialAction, final_color_action RenderingDeviceFinalAction, initial_depth_action RenderingDeviceInitialAction, final_depth_action RenderingDeviceFinalAction, clear_color_values PackedColorArray, clear_depth float64, clear_stencil int64, region Rect2) int64 {
	cargs := []gdc.ConstTypePtr{framebuffer.AsCTypePtr(), gdc.ConstTypePtr(&initial_color_action), gdc.ConstTypePtr(&final_color_action), gdc.ConstTypePtr(&initial_depth_action), gdc.ConstTypePtr(&final_depth_action), clear_color_values.AsCTypePtr(), gdc.ConstTypePtr(&clear_depth), gdc.ConstTypePtr(&clear_stencil), region.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&initial_color_action)
	pinner.Pin(&final_color_action)
	pinner.Pin(&initial_depth_action)
	pinner.Pin(&final_depth_action)
	pinner.Pin(&clear_depth)
	pinner.Pin(&clear_stencil)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) DrawListBeginSplit(framebuffer RID, splits int64, initial_color_action RenderingDeviceInitialAction, final_color_action RenderingDeviceFinalAction, initial_depth_action RenderingDeviceInitialAction, final_depth_action RenderingDeviceFinalAction, clear_color_values PackedColorArray, clear_depth float64, clear_stencil int64, region Rect2, storage_textures []RID) PackedInt64Array {
	cargs := []gdc.ConstTypePtr{framebuffer.AsCTypePtr(), gdc.ConstTypePtr(&splits), gdc.ConstTypePtr(&initial_color_action), gdc.ConstTypePtr(&final_color_action), gdc.ConstTypePtr(&initial_depth_action), gdc.ConstTypePtr(&final_depth_action), clear_color_values.AsCTypePtr(), gdc.ConstTypePtr(&clear_depth), gdc.ConstTypePtr(&clear_stencil), region.AsCTypePtr(), gdc.ConstTypePtr(&storage_textures)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt64Array()
	pinner.Pin(&splits)
	pinner.Pin(&initial_color_action)
	pinner.Pin(&final_color_action)
	pinner.Pin(&initial_depth_action)
	pinner.Pin(&final_depth_action)
	pinner.Pin(&clear_depth)
	pinner.Pin(&clear_stencil)
	pinner.Pin(&storage_textures)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListBeginSplit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) DrawListSetBlendConstants(draw_list int64, color Color) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_list), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListSetBlendConstants), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawListBindRenderPipeline(draw_list int64, render_pipeline RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_list), render_pipeline.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListBindRenderPipeline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawListBindUniformSet(draw_list int64, uniform_set RID, set_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_list), uniform_set.AsCTypePtr(), gdc.ConstTypePtr(&set_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListBindUniformSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawListBindVertexArray(draw_list int64, vertex_array RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_list), vertex_array.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListBindVertexArray), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawListBindIndexArray(draw_list int64, index_array RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_list), index_array.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListBindIndexArray), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawListSetPushConstant(draw_list int64, buffer PackedByteArray, size_bytes int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_list), buffer.AsCTypePtr(), gdc.ConstTypePtr(&size_bytes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListSetPushConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawListDraw(draw_list int64, use_indices bool, instances int64, procedural_vertex_count int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_list), gdc.ConstTypePtr(&use_indices), gdc.ConstTypePtr(&instances), gdc.ConstTypePtr(&procedural_vertex_count)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListDraw), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawListEnableScissor(draw_list int64, rect Rect2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_list), rect.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListEnableScissor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawListDisableScissor(draw_list int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_list)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListDisableScissor), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawListSwitchToNextPass() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListSwitchToNextPass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) DrawListSwitchToNextPassSplit(splits int64) PackedInt64Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&splits)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt64Array()
	pinner.Pin(&splits)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListSwitchToNextPassSplit), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) DrawListEnd() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawListEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) ComputeListBegin() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputeListBegin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) ComputeListBindComputePipeline(compute_list int64, compute_pipeline RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&compute_list), compute_pipeline.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputeListBindComputePipeline), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) ComputeListSetPushConstant(compute_list int64, buffer PackedByteArray, size_bytes int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&compute_list), buffer.AsCTypePtr(), gdc.ConstTypePtr(&size_bytes)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputeListSetPushConstant), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) ComputeListBindUniformSet(compute_list int64, uniform_set RID, set_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&compute_list), uniform_set.AsCTypePtr(), gdc.ConstTypePtr(&set_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputeListBindUniformSet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) ComputeListDispatch(compute_list int64, x_groups int64, y_groups int64, z_groups int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&compute_list), gdc.ConstTypePtr(&x_groups), gdc.ConstTypePtr(&y_groups), gdc.ConstTypePtr(&z_groups)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputeListDispatch), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) ComputeListDispatchIndirect(compute_list int64, buffer RID, offset int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&compute_list), buffer.AsCTypePtr(), gdc.ConstTypePtr(&offset)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputeListDispatchIndirect), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) ComputeListAddBarrier(compute_list int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&compute_list)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputeListAddBarrier), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) ComputeListEnd() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnComputeListEnd), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) FreeRid(rid RID) {
	cargs := []gdc.ConstTypePtr{rid.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFreeRid), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) CaptureTimestamp(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnCaptureTimestamp), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) GetCapturedTimestampsCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetCapturedTimestampsCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) GetCapturedTimestampsFrame() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetCapturedTimestampsFrame), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) GetCapturedTimestampGpuTime(index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetCapturedTimestampGpuTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) GetCapturedTimestampCpuTime(index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetCapturedTimestampCpuTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) GetCapturedTimestampName(index int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetCapturedTimestampName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) LimitGet(limit RenderingDeviceLimit) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&limit)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&limit)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnLimitGet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) GetFrameDelay() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetFrameDelay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) Submit() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnSubmit), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) Sync() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnSync), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) Barrier(from RenderingDeviceBarrierMask, to RenderingDeviceBarrierMask) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnBarrier), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) FullBarrier() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnFullBarrier), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) CreateLocalDevice() RenderingDevice {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRenderingDevice()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnCreateLocalDevice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) SetResourceName(id RID, name String) {
	cargs := []gdc.ConstTypePtr{id.AsCTypePtr(), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnSetResourceName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawCommandBeginLabel(name String, color Color) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawCommandBeginLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawCommandInsertLabel(name String, color Color) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), color.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawCommandInsertLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) DrawCommandEndLabel() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnDrawCommandEndLabel), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *RenderingDevice) GetDeviceVendorName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetDeviceVendorName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) GetDeviceName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetDeviceName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) GetDevicePipelineCacheUuid() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetDevicePipelineCacheUuid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *RenderingDevice) GetMemoryUsage(type_ RenderingDeviceMemoryType) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&type_)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetMemoryUsage), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *RenderingDevice) GetDriverResource(resource RenderingDeviceDriverResource, rid RID, index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&resource), rid.AsCTypePtr(), gdc.ConstTypePtr(&index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&resource)
	pinner.Pin(&index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForRenderingDevice.fnGetDriverResource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
