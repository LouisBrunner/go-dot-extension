// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScriptLanguageExtension struct {
  obj gdc.ObjectPtr
}

func (me *ScriptLanguageExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ScriptLanguageExtension) BaseClass() string {
  return "ScriptLanguageExtension"
}

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

func  (me *ScriptLanguageExtension) XGetName() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XInit() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetType() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetExtension() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XFinish() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetReservedWords() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XIsControlFlowKeyword(keyword String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetCommentDelimiters() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetStringDelimiters() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XMakeTemplate(template String, class_name String, base_class_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetBuiltInTemplates(object StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XIsUsingTemplates() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XValidate(script String, path String, validate_functions bool, validate_errors bool, validate_warnings bool, validate_safe_lines bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XValidatePath(path String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XCreateScript() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XHasNamedClasses() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XSupportsBuiltinMode() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XSupportsDocumentation() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XCanInheritFromFile() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XFindFunction(class_name String, function_name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XMakeFunction(class_name String, function_name String, function_args PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XOpenInExternalEditor(script Script, line int, column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XOverridesExternalEditor() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XCompleteCode(code String, path String, owner Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XLookupCode(code String, symbol String, path String, owner Object, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XAutoIndentCode(code String, from_line int, to_line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XAddGlobalConstant(name StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XAddNamedGlobalConstant(name StringName, value Variant, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XRemoveNamedGlobalConstant(name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XThreadEnter() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XThreadExit() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugGetError() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelCount() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelLine(level int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelFunction(level int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelLocals(level int, max_subitems int, max_depth int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelMembers(level int, max_subitems int, max_depth int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelInstance(level int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugGetGlobals(max_subitems int, max_depth int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugParseStackLevelExpression(level int, expression String, max_subitems int, max_depth int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XDebugGetCurrentStackInfo() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XReloadAllScripts() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XReloadToolScript(script Script, soft_reload bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetRecognizedExtensions() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetPublicFunctions() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetPublicConstants() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetPublicAnnotations() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XProfilingStart() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XProfilingStop() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XProfilingGetAccumulatedData(info_array *ScriptLanguageExtensionProfilingInfo, info_max int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XProfilingGetFrameData(info_array *ScriptLanguageExtensionProfilingInfo, info_max int, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XFrame() { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XHandlesGlobalClassType(type_ String, ) { // TODO: return value
  // TODO: implement
}

func  (me *ScriptLanguageExtension) XGetGlobalClassName(path String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
