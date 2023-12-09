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

func  (me *ScriptLanguageExtension) XGetName()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XInit()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetType()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetExtension()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XFinish()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetReservedWords()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XIsControlFlowKeyword(keyword String, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetCommentDelimiters()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetStringDelimiters()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XMakeTemplate(template String, class_name String, base_class_name String, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetBuiltInTemplates(object StringName, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XIsUsingTemplates()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XValidate(script String, path String, validate_functions bool, validate_errors bool, validate_warnings bool, validate_safe_lines bool, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XValidatePath(path String, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XCreateScript()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XHasNamedClasses()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XSupportsBuiltinMode()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XSupportsDocumentation()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XCanInheritFromFile()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XFindFunction(class_name String, function_name String, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XMakeFunction(class_name String, function_name String, function_args PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XOpenInExternalEditor(script Script, line int, column int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XOverridesExternalEditor()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XCompleteCode(code String, path String, owner Object, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XLookupCode(code String, symbol String, path String, owner Object, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XAutoIndentCode(code String, from_line int, to_line int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XAddGlobalConstant(name StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XAddNamedGlobalConstant(name StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XRemoveNamedGlobalConstant(name StringName, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XThreadEnter()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XThreadExit()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugGetError()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelCount()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelLine(level int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelFunction(level int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelLocals(level int, max_subitems int, max_depth int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelMembers(level int, max_subitems int, max_depth int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugGetStackLevelInstance(level int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugGetGlobals(max_subitems int, max_depth int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugParseStackLevelExpression(level int, expression String, max_subitems int, max_depth int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XDebugGetCurrentStackInfo()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XReloadAllScripts()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XReloadToolScript(script Script, soft_reload bool, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetRecognizedExtensions()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetPublicFunctions()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetPublicConstants()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetPublicAnnotations()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XProfilingStart()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XProfilingStop()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XProfilingGetAccumulatedData(info_array *ScriptLanguageExtensionProfilingInfo, info_max int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XProfilingGetFrameData(info_array *ScriptLanguageExtensionProfilingInfo, info_max int, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XFrame()  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XHandlesGlobalClassType(type_ String, )  {
  panic("TODO: implement")
}

func  (me *ScriptLanguageExtension) XGetGlobalClassName(path String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
