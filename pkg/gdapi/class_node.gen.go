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

type ptrsForNodeList struct {
	fnXProcess                        gdc.MethodBindPtr
	fnXPhysicsProcess                 gdc.MethodBindPtr
	fnXEnterTree                      gdc.MethodBindPtr
	fnXExitTree                       gdc.MethodBindPtr
	fnXReady                          gdc.MethodBindPtr
	fnXGetConfigurationWarnings       gdc.MethodBindPtr
	fnXInput                          gdc.MethodBindPtr
	fnXShortcutInput                  gdc.MethodBindPtr
	fnXUnhandledInput                 gdc.MethodBindPtr
	fnXUnhandledKeyInput              gdc.MethodBindPtr
	fnPrintOrphanNodes                gdc.MethodBindPtr
	fnAddSibling                      gdc.MethodBindPtr
	fnSetName                         gdc.MethodBindPtr
	fnGetName                         gdc.MethodBindPtr
	fnAddChild                        gdc.MethodBindPtr
	fnRemoveChild                     gdc.MethodBindPtr
	fnReparent                        gdc.MethodBindPtr
	fnGetChildCount                   gdc.MethodBindPtr
	fnGetChildren                     gdc.MethodBindPtr
	fnGetChild                        gdc.MethodBindPtr
	fnHasNode                         gdc.MethodBindPtr
	fnGetNode                         gdc.MethodBindPtr
	fnGetNodeOrNull                   gdc.MethodBindPtr
	fnGetParent                       gdc.MethodBindPtr
	fnFindChild                       gdc.MethodBindPtr
	fnFindChildren                    gdc.MethodBindPtr
	fnFindParent                      gdc.MethodBindPtr
	fnHasNodeAndResource              gdc.MethodBindPtr
	fnGetNodeAndResource              gdc.MethodBindPtr
	fnIsInsideTree                    gdc.MethodBindPtr
	fnIsAncestorOf                    gdc.MethodBindPtr
	fnIsGreaterThan                   gdc.MethodBindPtr
	fnGetPath                         gdc.MethodBindPtr
	fnGetPathTo                       gdc.MethodBindPtr
	fnAddToGroup                      gdc.MethodBindPtr
	fnRemoveFromGroup                 gdc.MethodBindPtr
	fnIsInGroup                       gdc.MethodBindPtr
	fnMoveChild                       gdc.MethodBindPtr
	fnGetGroups                       gdc.MethodBindPtr
	fnSetOwner                        gdc.MethodBindPtr
	fnGetOwner                        gdc.MethodBindPtr
	fnGetIndex                        gdc.MethodBindPtr
	fnPrintTree                       gdc.MethodBindPtr
	fnPrintTreePretty                 gdc.MethodBindPtr
	fnGetTreeString                   gdc.MethodBindPtr
	fnGetTreeStringPretty             gdc.MethodBindPtr
	fnSetSceneFilePath                gdc.MethodBindPtr
	fnGetSceneFilePath                gdc.MethodBindPtr
	fnPropagateNotification           gdc.MethodBindPtr
	fnPropagateCall                   gdc.MethodBindPtr
	fnSetPhysicsProcess               gdc.MethodBindPtr
	fnGetPhysicsProcessDeltaTime      gdc.MethodBindPtr
	fnIsPhysicsProcessing             gdc.MethodBindPtr
	fnGetProcessDeltaTime             gdc.MethodBindPtr
	fnSetProcess                      gdc.MethodBindPtr
	fnSetProcessPriority              gdc.MethodBindPtr
	fnGetProcessPriority              gdc.MethodBindPtr
	fnSetPhysicsProcessPriority       gdc.MethodBindPtr
	fnGetPhysicsProcessPriority       gdc.MethodBindPtr
	fnIsProcessing                    gdc.MethodBindPtr
	fnSetProcessInput                 gdc.MethodBindPtr
	fnIsProcessingInput               gdc.MethodBindPtr
	fnSetProcessShortcutInput         gdc.MethodBindPtr
	fnIsProcessingShortcutInput       gdc.MethodBindPtr
	fnSetProcessUnhandledInput        gdc.MethodBindPtr
	fnIsProcessingUnhandledInput      gdc.MethodBindPtr
	fnSetProcessUnhandledKeyInput     gdc.MethodBindPtr
	fnIsProcessingUnhandledKeyInput   gdc.MethodBindPtr
	fnSetProcessMode                  gdc.MethodBindPtr
	fnGetProcessMode                  gdc.MethodBindPtr
	fnCanProcess                      gdc.MethodBindPtr
	fnSetProcessThreadGroup           gdc.MethodBindPtr
	fnGetProcessThreadGroup           gdc.MethodBindPtr
	fnSetProcessThreadMessages        gdc.MethodBindPtr
	fnGetProcessThreadMessages        gdc.MethodBindPtr
	fnSetProcessThreadGroupOrder      gdc.MethodBindPtr
	fnGetProcessThreadGroupOrder      gdc.MethodBindPtr
	fnSetDisplayFolded                gdc.MethodBindPtr
	fnIsDisplayedFolded               gdc.MethodBindPtr
	fnSetProcessInternal              gdc.MethodBindPtr
	fnIsProcessingInternal            gdc.MethodBindPtr
	fnSetPhysicsProcessInternal       gdc.MethodBindPtr
	fnIsPhysicsProcessingInternal     gdc.MethodBindPtr
	fnGetWindow                       gdc.MethodBindPtr
	fnGetLastExclusiveWindow          gdc.MethodBindPtr
	fnGetTree                         gdc.MethodBindPtr
	fnCreateTween                     gdc.MethodBindPtr
	fnDuplicate                       gdc.MethodBindPtr
	fnReplaceBy                       gdc.MethodBindPtr
	fnSetSceneInstanceLoadPlaceholder gdc.MethodBindPtr
	fnGetSceneInstanceLoadPlaceholder gdc.MethodBindPtr
	fnSetEditableInstance             gdc.MethodBindPtr
	fnIsEditableInstance              gdc.MethodBindPtr
	fnGetViewport                     gdc.MethodBindPtr
	fnQueueFree                       gdc.MethodBindPtr
	fnRequestReady                    gdc.MethodBindPtr
	fnIsNodeReady                     gdc.MethodBindPtr
	fnSetMultiplayerAuthority         gdc.MethodBindPtr
	fnGetMultiplayerAuthority         gdc.MethodBindPtr
	fnIsMultiplayerAuthority          gdc.MethodBindPtr
	fnGetMultiplayer                  gdc.MethodBindPtr
	fnRpcConfig                       gdc.MethodBindPtr
	fnSetEditorDescription            gdc.MethodBindPtr
	fnGetEditorDescription            gdc.MethodBindPtr
	fnSetUniqueNameInOwner            gdc.MethodBindPtr
	fnIsUniqueNameInOwner             gdc.MethodBindPtr
	fnRpc                             gdc.MethodBindPtr
	fnRpcId                           gdc.MethodBindPtr
	fnUpdateConfigurationWarnings     gdc.MethodBindPtr
	fnCallDeferredThreadGroup         gdc.MethodBindPtr
	fnSetDeferredThreadGroup          gdc.MethodBindPtr
	fnNotifyDeferredThreadGroup       gdc.MethodBindPtr
	fnCallThreadSafe                  gdc.MethodBindPtr
	fnSetThreadSafe                   gdc.MethodBindPtr
	fnNotifyThreadSafe                gdc.MethodBindPtr
}

var ptrsForNode ptrsForNodeList

func initNodePtrs(iface gdc.Interface) {

	className := StringNameFromStr("Node")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("print_orphan_nodes")
		defer methodName.Destroy()
		ptrsForNode.fnPrintOrphanNodes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("add_sibling")
		defer methodName.Destroy()
		ptrsForNode.fnAddSibling = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2570952461))
	}
	{
		methodName := StringNameFromStr("set_name")
		defer methodName.Destroy()
		ptrsForNode.fnSetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_name")
		defer methodName.Destroy()
		ptrsForNode.fnGetName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("add_child")
		defer methodName.Destroy()
		ptrsForNode.fnAddChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3863233950))
	}
	{
		methodName := StringNameFromStr("remove_child")
		defer methodName.Destroy()
		ptrsForNode.fnRemoveChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("reparent")
		defer methodName.Destroy()
		ptrsForNode.fnReparent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3685795103))
	}
	{
		methodName := StringNameFromStr("get_child_count")
		defer methodName.Destroy()
		ptrsForNode.fnGetChildCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 894402480))
	}
	{
		methodName := StringNameFromStr("get_children")
		defer methodName.Destroy()
		ptrsForNode.fnGetChildren = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 873284517))
	}
	{
		methodName := StringNameFromStr("get_child")
		defer methodName.Destroy()
		ptrsForNode.fnGetChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 541253412))
	}
	{
		methodName := StringNameFromStr("has_node")
		defer methodName.Destroy()
		ptrsForNode.fnHasNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 861721659))
	}
	{
		methodName := StringNameFromStr("get_node")
		defer methodName.Destroy()
		ptrsForNode.fnGetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2734337346))
	}
	{
		methodName := StringNameFromStr("get_node_or_null")
		defer methodName.Destroy()
		ptrsForNode.fnGetNodeOrNull = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2734337346))
	}
	{
		methodName := StringNameFromStr("get_parent")
		defer methodName.Destroy()
		ptrsForNode.fnGetParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
	}
	{
		methodName := StringNameFromStr("find_child")
		defer methodName.Destroy()
		ptrsForNode.fnFindChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2008217037))
	}
	{
		methodName := StringNameFromStr("find_children")
		defer methodName.Destroy()
		ptrsForNode.fnFindChildren = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2560337219))
	}
	{
		methodName := StringNameFromStr("find_parent")
		defer methodName.Destroy()
		ptrsForNode.fnFindParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1140089439))
	}
	{
		methodName := StringNameFromStr("has_node_and_resource")
		defer methodName.Destroy()
		ptrsForNode.fnHasNodeAndResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 861721659))
	}
	{
		methodName := StringNameFromStr("get_node_and_resource")
		defer methodName.Destroy()
		ptrsForNode.fnGetNodeAndResource = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 502563882))
	}
	{
		methodName := StringNameFromStr("is_inside_tree")
		defer methodName.Destroy()
		ptrsForNode.fnIsInsideTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("is_ancestor_of")
		defer methodName.Destroy()
		ptrsForNode.fnIsAncestorOf = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3093956946))
	}
	{
		methodName := StringNameFromStr("is_greater_than")
		defer methodName.Destroy()
		ptrsForNode.fnIsGreaterThan = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3093956946))
	}
	{
		methodName := StringNameFromStr("get_path")
		defer methodName.Destroy()
		ptrsForNode.fnGetPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("get_path_to")
		defer methodName.Destroy()
		ptrsForNode.fnGetPathTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 498846349))
	}
	{
		methodName := StringNameFromStr("add_to_group")
		defer methodName.Destroy()
		ptrsForNode.fnAddToGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3683006648))
	}
	{
		methodName := StringNameFromStr("remove_from_group")
		defer methodName.Destroy()
		ptrsForNode.fnRemoveFromGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("is_in_group")
		defer methodName.Destroy()
		ptrsForNode.fnIsInGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2619796661))
	}
	{
		methodName := StringNameFromStr("move_child")
		defer methodName.Destroy()
		ptrsForNode.fnMoveChild = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3315886247))
	}
	{
		methodName := StringNameFromStr("get_groups")
		defer methodName.Destroy()
		ptrsForNode.fnGetGroups = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_owner")
		defer methodName.Destroy()
		ptrsForNode.fnSetOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1078189570))
	}
	{
		methodName := StringNameFromStr("get_owner")
		defer methodName.Destroy()
		ptrsForNode.fnGetOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3160264692))
	}
	{
		methodName := StringNameFromStr("get_index")
		defer methodName.Destroy()
		ptrsForNode.fnGetIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 894402480))
	}
	{
		methodName := StringNameFromStr("print_tree")
		defer methodName.Destroy()
		ptrsForNode.fnPrintTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("print_tree_pretty")
		defer methodName.Destroy()
		ptrsForNode.fnPrintTreePretty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_tree_string")
		defer methodName.Destroy()
		ptrsForNode.fnGetTreeString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("get_tree_string_pretty")
		defer methodName.Destroy()
		ptrsForNode.fnGetTreeStringPretty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
	}
	{
		methodName := StringNameFromStr("set_scene_file_path")
		defer methodName.Destroy()
		ptrsForNode.fnSetSceneFilePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_scene_file_path")
		defer methodName.Destroy()
		ptrsForNode.fnGetSceneFilePath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("propagate_notification")
		defer methodName.Destroy()
		ptrsForNode.fnPropagateNotification = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("propagate_call")
		defer methodName.Destroy()
		ptrsForNode.fnPropagateCall = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1871007965))
	}
	{
		methodName := StringNameFromStr("set_physics_process")
		defer methodName.Destroy()
		ptrsForNode.fnSetPhysicsProcess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_physics_process_delta_time")
		defer methodName.Destroy()
		ptrsForNode.fnGetPhysicsProcessDeltaTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("is_physics_processing")
		defer methodName.Destroy()
		ptrsForNode.fnIsPhysicsProcessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_process_delta_time")
		defer methodName.Destroy()
		ptrsForNode.fnGetProcessDeltaTime = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_process")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("set_process_priority")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_process_priority")
		defer methodName.Destroy()
		ptrsForNode.fnGetProcessPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_physics_process_priority")
		defer methodName.Destroy()
		ptrsForNode.fnSetPhysicsProcessPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_physics_process_priority")
		defer methodName.Destroy()
		ptrsForNode.fnGetPhysicsProcessPriority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("is_processing")
		defer methodName.Destroy()
		ptrsForNode.fnIsProcessing = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_process_input")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_processing_input")
		defer methodName.Destroy()
		ptrsForNode.fnIsProcessingInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_process_shortcut_input")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessShortcutInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_processing_shortcut_input")
		defer methodName.Destroy()
		ptrsForNode.fnIsProcessingShortcutInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_process_unhandled_input")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessUnhandledInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_processing_unhandled_input")
		defer methodName.Destroy()
		ptrsForNode.fnIsProcessingUnhandledInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_process_unhandled_key_input")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessUnhandledKeyInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_processing_unhandled_key_input")
		defer methodName.Destroy()
		ptrsForNode.fnIsProcessingUnhandledKeyInput = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_process_mode")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1841290486))
	}
	{
		methodName := StringNameFromStr("get_process_mode")
		defer methodName.Destroy()
		ptrsForNode.fnGetProcessMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 739966102))
	}
	{
		methodName := StringNameFromStr("can_process")
		defer methodName.Destroy()
		ptrsForNode.fnCanProcess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_process_thread_group")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessThreadGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2275442745))
	}
	{
		methodName := StringNameFromStr("get_process_thread_group")
		defer methodName.Destroy()
		ptrsForNode.fnGetProcessThreadGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1866404740))
	}
	{
		methodName := StringNameFromStr("set_process_thread_messages")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessThreadMessages = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1357280998))
	}
	{
		methodName := StringNameFromStr("get_process_thread_messages")
		defer methodName.Destroy()
		ptrsForNode.fnGetProcessThreadMessages = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4228993612))
	}
	{
		methodName := StringNameFromStr("set_process_thread_group_order")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessThreadGroupOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_process_thread_group_order")
		defer methodName.Destroy()
		ptrsForNode.fnGetProcessThreadGroupOrder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_display_folded")
		defer methodName.Destroy()
		ptrsForNode.fnSetDisplayFolded = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_displayed_folded")
		defer methodName.Destroy()
		ptrsForNode.fnIsDisplayedFolded = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_process_internal")
		defer methodName.Destroy()
		ptrsForNode.fnSetProcessInternal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_processing_internal")
		defer methodName.Destroy()
		ptrsForNode.fnIsProcessingInternal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_physics_process_internal")
		defer methodName.Destroy()
		ptrsForNode.fnSetPhysicsProcessInternal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_physics_processing_internal")
		defer methodName.Destroy()
		ptrsForNode.fnIsPhysicsProcessingInternal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_window")
		defer methodName.Destroy()
		ptrsForNode.fnGetWindow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1757182445))
	}
	{
		methodName := StringNameFromStr("get_last_exclusive_window")
		defer methodName.Destroy()
		ptrsForNode.fnGetLastExclusiveWindow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1757182445))
	}
	{
		methodName := StringNameFromStr("get_tree")
		defer methodName.Destroy()
		ptrsForNode.fnGetTree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2958820483))
	}
	{
		methodName := StringNameFromStr("create_tween")
		defer methodName.Destroy()
		ptrsForNode.fnCreateTween = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3426978995))
	}
	{
		methodName := StringNameFromStr("duplicate")
		defer methodName.Destroy()
		ptrsForNode.fnDuplicate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3511555459))
	}
	{
		methodName := StringNameFromStr("replace_by")
		defer methodName.Destroy()
		ptrsForNode.fnReplaceBy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2570952461))
	}
	{
		methodName := StringNameFromStr("set_scene_instance_load_placeholder")
		defer methodName.Destroy()
		ptrsForNode.fnSetSceneInstanceLoadPlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_scene_instance_load_placeholder")
		defer methodName.Destroy()
		ptrsForNode.fnGetSceneInstanceLoadPlaceholder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_editable_instance")
		defer methodName.Destroy()
		ptrsForNode.fnSetEditableInstance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2731852923))
	}
	{
		methodName := StringNameFromStr("is_editable_instance")
		defer methodName.Destroy()
		ptrsForNode.fnIsEditableInstance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3093956946))
	}
	{
		methodName := StringNameFromStr("get_viewport")
		defer methodName.Destroy()
		ptrsForNode.fnGetViewport = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3596683776))
	}
	{
		methodName := StringNameFromStr("queue_free")
		defer methodName.Destroy()
		ptrsForNode.fnQueueFree = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("request_ready")
		defer methodName.Destroy()
		ptrsForNode.fnRequestReady = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_node_ready")
		defer methodName.Destroy()
		ptrsForNode.fnIsNodeReady = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_multiplayer_authority")
		defer methodName.Destroy()
		ptrsForNode.fnSetMultiplayerAuthority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 972357352))
	}
	{
		methodName := StringNameFromStr("get_multiplayer_authority")
		defer methodName.Destroy()
		ptrsForNode.fnGetMultiplayerAuthority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("is_multiplayer_authority")
		defer methodName.Destroy()
		ptrsForNode.fnIsMultiplayerAuthority = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("get_multiplayer")
		defer methodName.Destroy()
		ptrsForNode.fnGetMultiplayer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 406750475))
	}
	{
		methodName := StringNameFromStr("rpc_config")
		defer methodName.Destroy()
		ptrsForNode.fnRpcConfig = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}
	{
		methodName := StringNameFromStr("set_editor_description")
		defer methodName.Destroy()
		ptrsForNode.fnSetEditorDescription = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_editor_description")
		defer methodName.Destroy()
		ptrsForNode.fnGetEditorDescription = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_unique_name_in_owner")
		defer methodName.Destroy()
		ptrsForNode.fnSetUniqueNameInOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_unique_name_in_owner")
		defer methodName.Destroy()
		ptrsForNode.fnIsUniqueNameInOwner = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("rpc")
		defer methodName.Destroy()
		ptrsForNode.fnRpc = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4047867050))
	}
	{
		methodName := StringNameFromStr("rpc_id")
		defer methodName.Destroy()
		ptrsForNode.fnRpcId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 361499283))
	}
	{
		methodName := StringNameFromStr("update_configuration_warnings")
		defer methodName.Destroy()
		ptrsForNode.fnUpdateConfigurationWarnings = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("call_deferred_thread_group")
		defer methodName.Destroy()
		ptrsForNode.fnCallDeferredThreadGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3400424181))
	}
	{
		methodName := StringNameFromStr("set_deferred_thread_group")
		defer methodName.Destroy()
		ptrsForNode.fnSetDeferredThreadGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}
	{
		methodName := StringNameFromStr("notify_deferred_thread_group")
		defer methodName.Destroy()
		ptrsForNode.fnNotifyDeferredThreadGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("call_thread_safe")
		defer methodName.Destroy()
		ptrsForNode.fnCallThreadSafe = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3400424181))
	}
	{
		methodName := StringNameFromStr("set_thread_safe")
		defer methodName.Destroy()
		ptrsForNode.fnSetThreadSafe = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3776071444))
	}
	{
		methodName := StringNameFromStr("notify_thread_safe")
		defer methodName.Destroy()
		ptrsForNode.fnNotifyThreadSafe = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}

}

type Node struct {
	Object
}

func (me *Node) BaseClass() string {
	return "Node"
}

func NewNode() *Node {
	str := StringNameFromStr("Node") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &Node{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Constants

var (
	NodeNotificationEnterTree              = 10
	NodeNotificationExitTree               = 11
	NodeNotificationMovedInParent          = 12
	NodeNotificationReady                  = 13
	NodeNotificationPaused                 = 14
	NodeNotificationUnpaused               = 15
	NodeNotificationPhysicsProcess         = 16
	NodeNotificationProcess                = 17
	NodeNotificationParented               = 18
	NodeNotificationUnparented             = 19
	NodeNotificationSceneInstantiated      = 20
	NodeNotificationDragBegin              = 21
	NodeNotificationDragEnd                = 22
	NodeNotificationPathRenamed            = 23
	NodeNotificationChildOrderChanged      = 24
	NodeNotificationInternalProcess        = 25
	NodeNotificationInternalPhysicsProcess = 26
	NodeNotificationPostEnterTree          = 27
	NodeNotificationDisabled               = 28
	NodeNotificationEnabled                = 29
	NodeNotificationEditorPreSave          = 9001
	NodeNotificationEditorPostSave         = 9002
	NodeNotificationWmMouseEnter           = 1002
	NodeNotificationWmMouseExit            = 1003
	NodeNotificationWmWindowFocusIn        = 1004
	NodeNotificationWmWindowFocusOut       = 1005
	NodeNotificationWmCloseRequest         = 1006
	NodeNotificationWmGoBackRequest        = 1007
	NodeNotificationWmSizeChanged          = 1008
	NodeNotificationWmDpiChange            = 1009
	NodeNotificationVpMouseEnter           = 1010
	NodeNotificationVpMouseExit            = 1011
	NodeNotificationOsMemoryWarning        = 2009
	NodeNotificationTranslationChanged     = 2010
	NodeNotificationWmAbout                = 2011
	NodeNotificationCrash                  = 2012
	NodeNotificationOsImeUpdate            = 2013
	NodeNotificationApplicationResumed     = 2014
	NodeNotificationApplicationPaused      = 2015
	NodeNotificationApplicationFocusIn     = 2016
	NodeNotificationApplicationFocusOut    = 2017
	NodeNotificationTextServerChanged      = 2018
)

// Enums

type NodeProcessMode int

const (
	NodeProcessModeProcessModeInherit    NodeProcessMode = 0
	NodeProcessModeProcessModePausable   NodeProcessMode = 1
	NodeProcessModeProcessModeWhenPaused NodeProcessMode = 2
	NodeProcessModeProcessModeAlways     NodeProcessMode = 3
	NodeProcessModeProcessModeDisabled   NodeProcessMode = 4
)

type NodeProcessThreadGroup int

const (
	NodeProcessThreadGroupProcessThreadGroupInherit    NodeProcessThreadGroup = 0
	NodeProcessThreadGroupProcessThreadGroupMainThread NodeProcessThreadGroup = 1
	NodeProcessThreadGroupProcessThreadGroupSubThread  NodeProcessThreadGroup = 2
)

type NodeProcessThreadMessages int

const (
	NodeProcessThreadMessagesFlagProcessThreadMessages        NodeProcessThreadMessages = 1
	NodeProcessThreadMessagesFlagProcessThreadMessagesPhysics NodeProcessThreadMessages = 2
	NodeProcessThreadMessagesFlagProcessThreadMessagesAll     NodeProcessThreadMessages = 3
)

type NodeDuplicateFlags int

const (
	NodeDuplicateFlagsDuplicateSignals          NodeDuplicateFlags = 1
	NodeDuplicateFlagsDuplicateGroups           NodeDuplicateFlags = 2
	NodeDuplicateFlagsDuplicateScripts          NodeDuplicateFlags = 4
	NodeDuplicateFlagsDuplicateUseInstantiation NodeDuplicateFlags = 8
)

type NodeInternalMode int

const (
	NodeInternalModeInternalModeDisabled NodeInternalMode = 0
	NodeInternalModeInternalModeFront    NodeInternalMode = 1
	NodeInternalModeInternalModeBack     NodeInternalMode = 2
)

func (me *Node) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *Node) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *Node) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func NodePrintOrphanNodes() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnPrintOrphanNodes), nil, unsafe.SliceData(cargs), nil)

}

func (me *Node) AddSibling(sibling Node, force_readable_name bool) {
	cargs := []gdc.ConstTypePtr{sibling.AsCTypePtr(), gdc.ConstTypePtr(&force_readable_name)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnAddSibling), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) SetName(name String) {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetName() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) AddChild(node Node, force_readable_name bool, internal NodeInternalMode) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), gdc.ConstTypePtr(&force_readable_name), gdc.ConstTypePtr(&internal)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnAddChild), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) RemoveChild(node Node) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnRemoveChild), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) Reparent(new_parent Node, keep_global_transform bool) {
	cargs := []gdc.ConstTypePtr{new_parent.AsCTypePtr(), gdc.ConstTypePtr(&keep_global_transform)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnReparent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetChildCount(include_internal bool) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&include_internal)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&include_internal)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetChildCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) GetChildren(include_internal bool) []Node {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&include_internal)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&include_internal)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetChildren), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Node](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Node) GetChild(idx int64, include_internal bool) Node {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), gdc.ConstTypePtr(&include_internal)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()
	pinner.Pin(&idx)
	pinner.Pin(&include_internal)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetChild), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) HasNode(path NodePath) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnHasNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) GetNode(path NodePath) Node {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) GetNodeOrNull(path NodePath) Node {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetNodeOrNull), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) GetParent() Node {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) FindChild(pattern String, recursive bool, owned bool) Node {
	cargs := []gdc.ConstTypePtr{pattern.AsCTypePtr(), gdc.ConstTypePtr(&recursive), gdc.ConstTypePtr(&owned)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()
	pinner.Pin(&recursive)
	pinner.Pin(&owned)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnFindChild), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) FindChildren(pattern String, type_ String, recursive bool, owned bool) []Node {
	cargs := []gdc.ConstTypePtr{pattern.AsCTypePtr(), type_.AsCTypePtr(), gdc.ConstTypePtr(&recursive), gdc.ConstTypePtr(&owned)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()
	pinner.Pin(&recursive)
	pinner.Pin(&owned)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnFindChildren), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[Node](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Node) FindParent(pattern String) Node {
	cargs := []gdc.ConstTypePtr{pattern.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnFindParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) HasNodeAndResource(path NodePath) bool {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnHasNodeAndResource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) GetNodeAndResource(path NodePath) Array {
	cargs := []gdc.ConstTypePtr{path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetNodeAndResource), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) IsInsideTree() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsInsideTree), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) IsAncestorOf(node Node) bool {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsAncestorOf), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) IsGreaterThan(node Node) bool {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsGreaterThan), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) GetPath() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) GetPathTo(node Node, use_unique_path bool) NodePath {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), gdc.ConstTypePtr(&use_unique_path)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()
	pinner.Pin(&use_unique_path)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetPathTo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) AddToGroup(group StringName, persistent bool) {
	cargs := []gdc.ConstTypePtr{group.AsCTypePtr(), gdc.ConstTypePtr(&persistent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnAddToGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) RemoveFromGroup(group StringName) {
	cargs := []gdc.ConstTypePtr{group.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnRemoveFromGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsInGroup(group StringName) bool {
	cargs := []gdc.ConstTypePtr{group.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsInGroup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) MoveChild(child_node Node, to_index int64) {
	cargs := []gdc.ConstTypePtr{child_node.AsCTypePtr(), gdc.ConstTypePtr(&to_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnMoveChild), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetGroups() []StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetGroups), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[StringName](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *Node) SetOwner(owner Node) {
	cargs := []gdc.ConstTypePtr{owner.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetOwner), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetOwner() Node {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) GetIndex(include_internal bool) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&include_internal)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&include_internal)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) PrintTree() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnPrintTree), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) PrintTreePretty() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnPrintTreePretty), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetTreeString() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetTreeString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) GetTreeStringPretty() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetTreeStringPretty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) SetSceneFilePath(scene_file_path String) {
	cargs := []gdc.ConstTypePtr{scene_file_path.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetSceneFilePath), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetSceneFilePath() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetSceneFilePath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) PropagateNotification(what int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&what)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnPropagateNotification), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) PropagateCall(method StringName, args Array, parent_first bool) {
	cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), args.AsCTypePtr(), gdc.ConstTypePtr(&parent_first)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnPropagateCall), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) SetPhysicsProcess(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetPhysicsProcess), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetPhysicsProcessDeltaTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetPhysicsProcessDeltaTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) IsPhysicsProcessing() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsPhysicsProcessing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) GetProcessDeltaTime() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetProcessDeltaTime), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetProcess(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcess), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) SetProcessPriority(priority int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetProcessPriority() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetProcessPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetPhysicsProcessPriority(priority int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetPhysicsProcessPriority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetPhysicsProcessPriority() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetPhysicsProcessPriority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) IsProcessing() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsProcessing), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetProcessInput(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsProcessingInput() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsProcessingInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetProcessShortcutInput(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessShortcutInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsProcessingShortcutInput() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsProcessingShortcutInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetProcessUnhandledInput(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessUnhandledInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsProcessingUnhandledInput() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsProcessingUnhandledInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetProcessUnhandledKeyInput(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessUnhandledKeyInput), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsProcessingUnhandledKeyInput() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsProcessingUnhandledKeyInput), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetProcessMode(mode NodeProcessMode) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessMode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetProcessMode() NodeProcessMode {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NodeProcessMode

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetProcessMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Node) CanProcess() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnCanProcess), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetProcessThreadGroup(mode NodeProcessThreadGroup) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessThreadGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetProcessThreadGroup() NodeProcessThreadGroup {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NodeProcessThreadGroup

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetProcessThreadGroup), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Node) SetProcessThreadMessages(flags NodeProcessThreadMessages) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessThreadMessages), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetProcessThreadMessages() NodeProcessThreadMessages {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret NodeProcessThreadMessages

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetProcessThreadMessages), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Node) SetProcessThreadGroupOrder(order int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&order)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessThreadGroupOrder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetProcessThreadGroupOrder() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetProcessThreadGroupOrder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetDisplayFolded(fold bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&fold)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetDisplayFolded), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsDisplayedFolded() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsDisplayedFolded), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetProcessInternal(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetProcessInternal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsProcessingInternal() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsProcessingInternal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetPhysicsProcessInternal(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetPhysicsProcessInternal), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsPhysicsProcessingInternal() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsPhysicsProcessingInternal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) GetWindow() Window {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewWindow()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetWindow), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) GetLastExclusiveWindow() Window {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewWindow()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetLastExclusiveWindow), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) GetTree() SceneTree {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSceneTree()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetTree), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) CreateTween() Tween {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTween()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnCreateTween), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) Duplicate(flags int64) Node {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flags)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNode()
	pinner.Pin(&flags)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnDuplicate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) ReplaceBy(node Node, keep_groups bool) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), gdc.ConstTypePtr(&keep_groups)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnReplaceBy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) SetSceneInstanceLoadPlaceholder(load_placeholder bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&load_placeholder)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetSceneInstanceLoadPlaceholder), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetSceneInstanceLoadPlaceholder() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetSceneInstanceLoadPlaceholder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetEditableInstance(node Node, is_editable bool) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr(), gdc.ConstTypePtr(&is_editable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetEditableInstance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsEditableInstance(node Node) bool {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsEditableInstance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) GetViewport() Viewport {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewViewport()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetViewport), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) QueueFree() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnQueueFree), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) RequestReady() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnRequestReady), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsNodeReady() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsNodeReady), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) SetMultiplayerAuthority(id int64, recursive bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&id), gdc.ConstTypePtr(&recursive)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetMultiplayerAuthority), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetMultiplayerAuthority() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetMultiplayerAuthority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) IsMultiplayerAuthority() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsMultiplayerAuthority), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) GetMultiplayer() MultiplayerAPI {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewMultiplayerAPI()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetMultiplayer), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) RpcConfig(method StringName, config Variant) {
	cargs := []gdc.ConstTypePtr{method.AsCTypePtr(), config.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnRpcConfig), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) SetEditorDescription(editor_description String) {
	cargs := []gdc.ConstTypePtr{editor_description.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetEditorDescription), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) GetEditorDescription() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnGetEditorDescription), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Node) SetUniqueNameInOwner(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetUniqueNameInOwner), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) IsUniqueNameInOwner() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnIsUniqueNameInOwner), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Node) Rpc(method StringName, varargs ...Variant) Error {
	cargs := make([]gdc.ConstVariantPtr, 0, 1+len(varargs))
	var0 := method.AsVariant()
	defer var0.Destroy()
	cargs = append(cargs, var0.AsCPtr())
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}
	ret := NewVariant()
	defer ret.Destroy()
	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForNode.fnRpc), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return Error(-1)
	}
	retInt, err := ret.AsInt()
	if err != nil {
		log.Printf("Error converting return value to int enum: %v", err) // FIXME: bad logging
		return Error(-1)
	}
	return Error(retInt.Get())
}

func (me *Node) RpcId(peer_id int64, method StringName, varargs ...Variant) Error {
	cargs := make([]gdc.ConstVariantPtr, 0, 2+len(varargs))
	intVar0 := NewIntFromInt(peer_id)
	defer intVar0.Destroy()
	var0 := intVar0.AsVariant()
	defer var0.Destroy()
	cargs = append(cargs, var0.AsCPtr())
	var1 := method.AsVariant()
	defer var1.Destroy()
	cargs = append(cargs, var1.AsCPtr())
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}
	ret := NewVariant()
	defer ret.Destroy()
	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForNode.fnRpcId), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return Error(-1)
	}
	retInt, err := ret.AsInt()
	if err != nil {
		log.Printf("Error converting return value to int enum: %v", err) // FIXME: bad logging
		return Error(-1)
	}
	return Error(retInt.Get())
}

func (me *Node) UpdateConfigurationWarnings() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnUpdateConfigurationWarnings), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) CallDeferredThreadGroup(method StringName, varargs ...Variant) Variant {
	cargs := make([]gdc.ConstVariantPtr, 0, 1+len(varargs))
	var0 := method.AsVariant()
	defer var0.Destroy()
	cargs = append(cargs, var0.AsCPtr())
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}
	ret := NewVariant()
	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForNode.fnCallDeferredThreadGroup), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return *ret
	}
	return *ret
}

func (me *Node) SetDeferredThreadGroup(property StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetDeferredThreadGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) NotifyDeferredThreadGroup(what int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&what)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnNotifyDeferredThreadGroup), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) CallThreadSafe(method StringName, varargs ...Variant) Variant {
	cargs := make([]gdc.ConstVariantPtr, 0, 1+len(varargs))
	var0 := method.AsVariant()
	defer var0.Destroy()
	cargs = append(cargs, var0.AsCPtr())
	for _, v := range varargs {
		cargs = append(cargs, v.AsCPtr())
	}
	ret := NewVariant()
	cerr := &gdc.CallError{}
	giface.ObjectMethodBindCall(ensurePtr(ptrsForNode.fnCallThreadSafe), me.obj, unsafe.SliceData(cargs), gdc.Int(len(cargs)), ret.asUninitialized(), cerr)
	if cerr.Error != gdc.CallOk {
		log.Printf("Error calling method: %v", cerr) // FIXME: bad logging
		return *ret
	}
	return *ret
}

func (me *Node) SetThreadSafe(property StringName, value Variant) {
	cargs := []gdc.ConstTypePtr{property.AsCTypePtr(), value.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnSetThreadSafe), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Node) NotifyThreadSafe(what int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&what)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForNode.fnNotifyThreadSafe), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type NodeReadySignalFn func()

func (me *Node) ConnectReady(subs SignalSubscribers, fn NodeReadySignalFn) {
	sig := StringNameFromStr("ready")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node) DisconnectReady(subs SignalSubscribers, fn NodeReadySignalFn) {
	sig := StringNameFromStr("ready")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NodeRenamedSignalFn func()

func (me *Node) ConnectRenamed(subs SignalSubscribers, fn NodeRenamedSignalFn) {
	sig := StringNameFromStr("renamed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node) DisconnectRenamed(subs SignalSubscribers, fn NodeRenamedSignalFn) {
	sig := StringNameFromStr("renamed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NodeTreeEnteredSignalFn func()

func (me *Node) ConnectTreeEntered(subs SignalSubscribers, fn NodeTreeEnteredSignalFn) {
	sig := StringNameFromStr("tree_entered")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node) DisconnectTreeEntered(subs SignalSubscribers, fn NodeTreeEnteredSignalFn) {
	sig := StringNameFromStr("tree_entered")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NodeTreeExitingSignalFn func()

func (me *Node) ConnectTreeExiting(subs SignalSubscribers, fn NodeTreeExitingSignalFn) {
	sig := StringNameFromStr("tree_exiting")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node) DisconnectTreeExiting(subs SignalSubscribers, fn NodeTreeExitingSignalFn) {
	sig := StringNameFromStr("tree_exiting")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NodeTreeExitedSignalFn func()

func (me *Node) ConnectTreeExited(subs SignalSubscribers, fn NodeTreeExitedSignalFn) {
	sig := StringNameFromStr("tree_exited")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node) DisconnectTreeExited(subs SignalSubscribers, fn NodeTreeExitedSignalFn) {
	sig := StringNameFromStr("tree_exited")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NodeChildEnteredTreeSignalFn func(node Node)

func (me *Node) ConnectChildEnteredTree(subs SignalSubscribers, fn NodeChildEnteredTreeSignalFn) {
	sig := StringNameFromStr("child_entered_tree")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node) DisconnectChildEnteredTree(subs SignalSubscribers, fn NodeChildEnteredTreeSignalFn) {
	sig := StringNameFromStr("child_entered_tree")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NodeChildExitingTreeSignalFn func(node Node)

func (me *Node) ConnectChildExitingTree(subs SignalSubscribers, fn NodeChildExitingTreeSignalFn) {
	sig := StringNameFromStr("child_exiting_tree")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node) DisconnectChildExitingTree(subs SignalSubscribers, fn NodeChildExitingTreeSignalFn) {
	sig := StringNameFromStr("child_exiting_tree")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NodeChildOrderChangedSignalFn func()

func (me *Node) ConnectChildOrderChanged(subs SignalSubscribers, fn NodeChildOrderChangedSignalFn) {
	sig := StringNameFromStr("child_order_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node) DisconnectChildOrderChanged(subs SignalSubscribers, fn NodeChildOrderChangedSignalFn) {
	sig := StringNameFromStr("child_order_changed")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type NodeReplacingBySignalFn func(node Node)

func (me *Node) ConnectReplacingBy(subs SignalSubscribers, fn NodeReplacingBySignalFn) {
	sig := StringNameFromStr("replacing_by")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Node) DisconnectReplacingBy(subs SignalSubscribers, fn NodeReplacingBySignalFn) {
	sig := StringNameFromStr("replacing_by")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
