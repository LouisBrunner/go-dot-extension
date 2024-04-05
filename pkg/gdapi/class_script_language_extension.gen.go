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

type ptrsForScriptLanguageExtensionList struct {
  fnXGetName gdc.MethodBindPtr
  fnXInit gdc.MethodBindPtr
  fnXGetType gdc.MethodBindPtr
  fnXGetExtension gdc.MethodBindPtr
  fnXFinish gdc.MethodBindPtr
  fnXGetReservedWords gdc.MethodBindPtr
  fnXIsControlFlowKeyword gdc.MethodBindPtr
  fnXGetCommentDelimiters gdc.MethodBindPtr
  fnXGetDocCommentDelimiters gdc.MethodBindPtr
  fnXGetStringDelimiters gdc.MethodBindPtr
  fnXMakeTemplate gdc.MethodBindPtr
  fnXGetBuiltInTemplates gdc.MethodBindPtr
  fnXIsUsingTemplates gdc.MethodBindPtr
  fnXValidate gdc.MethodBindPtr
  fnXValidatePath gdc.MethodBindPtr
  fnXCreateScript gdc.MethodBindPtr
  fnXHasNamedClasses gdc.MethodBindPtr
  fnXSupportsBuiltinMode gdc.MethodBindPtr
  fnXSupportsDocumentation gdc.MethodBindPtr
  fnXCanInheritFromFile gdc.MethodBindPtr
  fnXFindFunction gdc.MethodBindPtr
  fnXMakeFunction gdc.MethodBindPtr
  fnXOpenInExternalEditor gdc.MethodBindPtr
  fnXOverridesExternalEditor gdc.MethodBindPtr
  fnXCompleteCode gdc.MethodBindPtr
  fnXLookupCode gdc.MethodBindPtr
  fnXAutoIndentCode gdc.MethodBindPtr
  fnXAddGlobalConstant gdc.MethodBindPtr
  fnXAddNamedGlobalConstant gdc.MethodBindPtr
  fnXRemoveNamedGlobalConstant gdc.MethodBindPtr
  fnXThreadEnter gdc.MethodBindPtr
  fnXThreadExit gdc.MethodBindPtr
  fnXDebugGetError gdc.MethodBindPtr
  fnXDebugGetStackLevelCount gdc.MethodBindPtr
  fnXDebugGetStackLevelLine gdc.MethodBindPtr
  fnXDebugGetStackLevelFunction gdc.MethodBindPtr
  fnXDebugGetStackLevelLocals gdc.MethodBindPtr
  fnXDebugGetStackLevelMembers gdc.MethodBindPtr
  fnXDebugGetStackLevelInstance gdc.MethodBindPtr
  fnXDebugGetGlobals gdc.MethodBindPtr
  fnXDebugParseStackLevelExpression gdc.MethodBindPtr
  fnXDebugGetCurrentStackInfo gdc.MethodBindPtr
  fnXReloadAllScripts gdc.MethodBindPtr
  fnXReloadToolScript gdc.MethodBindPtr
  fnXGetRecognizedExtensions gdc.MethodBindPtr
  fnXGetPublicFunctions gdc.MethodBindPtr
  fnXGetPublicConstants gdc.MethodBindPtr
  fnXGetPublicAnnotations gdc.MethodBindPtr
  fnXProfilingStart gdc.MethodBindPtr
  fnXProfilingStop gdc.MethodBindPtr
  fnXProfilingGetAccumulatedData gdc.MethodBindPtr
  fnXProfilingGetFrameData gdc.MethodBindPtr
  fnXFrame gdc.MethodBindPtr
  fnXHandlesGlobalClassType gdc.MethodBindPtr
  fnXGetGlobalClassName gdc.MethodBindPtr
}

var ptrsForScriptLanguageExtension ptrsForScriptLanguageExtensionList

func initScriptLanguageExtensionPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ScriptLanguageExtension")
  defer className.Destroy()
}

type ScriptLanguageExtension struct {
  ScriptLanguage
}

func (me *ScriptLanguageExtension) BaseClass() string {
  return "ScriptLanguageExtension"
}

func NewScriptLanguageExtension() *ScriptLanguageExtension {
  str := StringNameFromStr("ScriptLanguageExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ScriptLanguageExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ScriptLanguageExtensionLookupResultType int
const (
  ScriptLanguageExtensionLookupResultTypeLookupResultScriptLocation ScriptLanguageExtensionLookupResultType = 0
  ScriptLanguageExtensionLookupResultTypeLookupResultClass ScriptLanguageExtensionLookupResultType = 1
  ScriptLanguageExtensionLookupResultTypeLookupResultClassConstant ScriptLanguageExtensionLookupResultType = 2
  ScriptLanguageExtensionLookupResultTypeLookupResultClassProperty ScriptLanguageExtensionLookupResultType = 3
  ScriptLanguageExtensionLookupResultTypeLookupResultClassMethod ScriptLanguageExtensionLookupResultType = 4
  ScriptLanguageExtensionLookupResultTypeLookupResultClassSignal ScriptLanguageExtensionLookupResultType = 5
  ScriptLanguageExtensionLookupResultTypeLookupResultClassEnum ScriptLanguageExtensionLookupResultType = 6
  ScriptLanguageExtensionLookupResultTypeLookupResultClassTbdGlobalscope ScriptLanguageExtensionLookupResultType = 7
  ScriptLanguageExtensionLookupResultTypeLookupResultClassAnnotation ScriptLanguageExtensionLookupResultType = 8
  ScriptLanguageExtensionLookupResultTypeLookupResultMax ScriptLanguageExtensionLookupResultType = 9
)

type ScriptLanguageExtensionCodeCompletionLocation int
const (
  ScriptLanguageExtensionCodeCompletionLocationLocationLocal ScriptLanguageExtensionCodeCompletionLocation = 0
  ScriptLanguageExtensionCodeCompletionLocationLocationParentMask ScriptLanguageExtensionCodeCompletionLocation = 256
  ScriptLanguageExtensionCodeCompletionLocationLocationOtherUserCode ScriptLanguageExtensionCodeCompletionLocation = 512
  ScriptLanguageExtensionCodeCompletionLocationLocationOther ScriptLanguageExtensionCodeCompletionLocation = 1024
)

type ScriptLanguageExtensionCodeCompletionKind int
const (
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindClass ScriptLanguageExtensionCodeCompletionKind = 0
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindFunction ScriptLanguageExtensionCodeCompletionKind = 1
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindSignal ScriptLanguageExtensionCodeCompletionKind = 2
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindVariable ScriptLanguageExtensionCodeCompletionKind = 3
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindMember ScriptLanguageExtensionCodeCompletionKind = 4
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindEnum ScriptLanguageExtensionCodeCompletionKind = 5
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindConstant ScriptLanguageExtensionCodeCompletionKind = 6
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindNodePath ScriptLanguageExtensionCodeCompletionKind = 7
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindFilePath ScriptLanguageExtensionCodeCompletionKind = 8
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindPlainText ScriptLanguageExtensionCodeCompletionKind = 9
  ScriptLanguageExtensionCodeCompletionKindCodeCompletionKindMax ScriptLanguageExtensionCodeCompletionKind = 10
)

func (me *ScriptLanguageExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ScriptLanguageExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ScriptLanguageExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
