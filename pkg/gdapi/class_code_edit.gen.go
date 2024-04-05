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

type ptrsForCodeEditList struct {
  fnXConfirmCodeCompletion gdc.MethodBindPtr
  fnXRequestCodeCompletion gdc.MethodBindPtr
  fnXFilterCodeCompletionCandidates gdc.MethodBindPtr
  fnSetIndentSize gdc.MethodBindPtr
  fnGetIndentSize gdc.MethodBindPtr
  fnSetIndentUsingSpaces gdc.MethodBindPtr
  fnIsIndentUsingSpaces gdc.MethodBindPtr
  fnSetAutoIndentEnabled gdc.MethodBindPtr
  fnIsAutoIndentEnabled gdc.MethodBindPtr
  fnSetAutoIndentPrefixes gdc.MethodBindPtr
  fnGetAutoIndentPrefixes gdc.MethodBindPtr
  fnDoIndent gdc.MethodBindPtr
  fnIndentLines gdc.MethodBindPtr
  fnUnindentLines gdc.MethodBindPtr
  fnConvertIndent gdc.MethodBindPtr
  fnSetAutoBraceCompletionEnabled gdc.MethodBindPtr
  fnIsAutoBraceCompletionEnabled gdc.MethodBindPtr
  fnSetHighlightMatchingBracesEnabled gdc.MethodBindPtr
  fnIsHighlightMatchingBracesEnabled gdc.MethodBindPtr
  fnAddAutoBraceCompletionPair gdc.MethodBindPtr
  fnSetAutoBraceCompletionPairs gdc.MethodBindPtr
  fnGetAutoBraceCompletionPairs gdc.MethodBindPtr
  fnHasAutoBraceCompletionOpenKey gdc.MethodBindPtr
  fnHasAutoBraceCompletionCloseKey gdc.MethodBindPtr
  fnGetAutoBraceCompletionCloseKey gdc.MethodBindPtr
  fnSetDrawBreakpointsGutter gdc.MethodBindPtr
  fnIsDrawingBreakpointsGutter gdc.MethodBindPtr
  fnSetDrawBookmarksGutter gdc.MethodBindPtr
  fnIsDrawingBookmarksGutter gdc.MethodBindPtr
  fnSetDrawExecutingLinesGutter gdc.MethodBindPtr
  fnIsDrawingExecutingLinesGutter gdc.MethodBindPtr
  fnSetLineAsBreakpoint gdc.MethodBindPtr
  fnIsLineBreakpointed gdc.MethodBindPtr
  fnClearBreakpointedLines gdc.MethodBindPtr
  fnGetBreakpointedLines gdc.MethodBindPtr
  fnSetLineAsBookmarked gdc.MethodBindPtr
  fnIsLineBookmarked gdc.MethodBindPtr
  fnClearBookmarkedLines gdc.MethodBindPtr
  fnGetBookmarkedLines gdc.MethodBindPtr
  fnSetLineAsExecuting gdc.MethodBindPtr
  fnIsLineExecuting gdc.MethodBindPtr
  fnClearExecutingLines gdc.MethodBindPtr
  fnGetExecutingLines gdc.MethodBindPtr
  fnSetDrawLineNumbers gdc.MethodBindPtr
  fnIsDrawLineNumbersEnabled gdc.MethodBindPtr
  fnSetLineNumbersZeroPadded gdc.MethodBindPtr
  fnIsLineNumbersZeroPadded gdc.MethodBindPtr
  fnSetDrawFoldGutter gdc.MethodBindPtr
  fnIsDrawingFoldGutter gdc.MethodBindPtr
  fnSetLineFoldingEnabled gdc.MethodBindPtr
  fnIsLineFoldingEnabled gdc.MethodBindPtr
  fnCanFoldLine gdc.MethodBindPtr
  fnFoldLine gdc.MethodBindPtr
  fnUnfoldLine gdc.MethodBindPtr
  fnFoldAllLines gdc.MethodBindPtr
  fnUnfoldAllLines gdc.MethodBindPtr
  fnToggleFoldableLine gdc.MethodBindPtr
  fnIsLineFolded gdc.MethodBindPtr
  fnGetFoldedLines gdc.MethodBindPtr
  fnCreateCodeRegion gdc.MethodBindPtr
  fnGetCodeRegionStartTag gdc.MethodBindPtr
  fnGetCodeRegionEndTag gdc.MethodBindPtr
  fnSetCodeRegionTags gdc.MethodBindPtr
  fnIsLineCodeRegionStart gdc.MethodBindPtr
  fnIsLineCodeRegionEnd gdc.MethodBindPtr
  fnAddStringDelimiter gdc.MethodBindPtr
  fnRemoveStringDelimiter gdc.MethodBindPtr
  fnHasStringDelimiter gdc.MethodBindPtr
  fnSetStringDelimiters gdc.MethodBindPtr
  fnClearStringDelimiters gdc.MethodBindPtr
  fnGetStringDelimiters gdc.MethodBindPtr
  fnIsInString gdc.MethodBindPtr
  fnAddCommentDelimiter gdc.MethodBindPtr
  fnRemoveCommentDelimiter gdc.MethodBindPtr
  fnHasCommentDelimiter gdc.MethodBindPtr
  fnSetCommentDelimiters gdc.MethodBindPtr
  fnClearCommentDelimiters gdc.MethodBindPtr
  fnGetCommentDelimiters gdc.MethodBindPtr
  fnIsInComment gdc.MethodBindPtr
  fnGetDelimiterStartKey gdc.MethodBindPtr
  fnGetDelimiterEndKey gdc.MethodBindPtr
  fnGetDelimiterStartPosition gdc.MethodBindPtr
  fnGetDelimiterEndPosition gdc.MethodBindPtr
  fnSetCodeHint gdc.MethodBindPtr
  fnSetCodeHintDrawBelow gdc.MethodBindPtr
  fnGetTextForCodeCompletion gdc.MethodBindPtr
  fnRequestCodeCompletion gdc.MethodBindPtr
  fnAddCodeCompletionOption gdc.MethodBindPtr
  fnUpdateCodeCompletionOptions gdc.MethodBindPtr
  fnGetCodeCompletionOptions gdc.MethodBindPtr
  fnGetCodeCompletionOption gdc.MethodBindPtr
  fnGetCodeCompletionSelectedIndex gdc.MethodBindPtr
  fnSetCodeCompletionSelectedIndex gdc.MethodBindPtr
  fnConfirmCodeCompletion gdc.MethodBindPtr
  fnCancelCodeCompletion gdc.MethodBindPtr
  fnSetCodeCompletionEnabled gdc.MethodBindPtr
  fnIsCodeCompletionEnabled gdc.MethodBindPtr
  fnSetCodeCompletionPrefixes gdc.MethodBindPtr
  fnGetCodeCompletionPrefixes gdc.MethodBindPtr
  fnSetLineLengthGuidelines gdc.MethodBindPtr
  fnGetLineLengthGuidelines gdc.MethodBindPtr
  fnSetSymbolLookupOnClickEnabled gdc.MethodBindPtr
  fnIsSymbolLookupOnClickEnabled gdc.MethodBindPtr
  fnGetTextForSymbolLookup gdc.MethodBindPtr
  fnGetTextWithCursorChar gdc.MethodBindPtr
  fnSetSymbolLookupWordAsValid gdc.MethodBindPtr
  fnDuplicateLines gdc.MethodBindPtr
}

var ptrsForCodeEdit ptrsForCodeEditList

func initCodeEditPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CodeEdit")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_indent_size")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetIndentSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_indent_size")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetIndentSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_indent_using_spaces")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetIndentUsingSpaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_indent_using_spaces")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsIndentUsingSpaces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_auto_indent_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetAutoIndentEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_auto_indent_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsAutoIndentEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_auto_indent_prefixes")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetAutoIndentPrefixes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_auto_indent_prefixes")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetAutoIndentPrefixes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("do_indent")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnDoIndent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("indent_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIndentLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("unindent_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnUnindentLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("convert_indent")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnConvertIndent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 423910286))
  }
  {
    methodName := StringNameFromStr("set_auto_brace_completion_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetAutoBraceCompletionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_auto_brace_completion_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsAutoBraceCompletionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_highlight_matching_braces_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetHighlightMatchingBracesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_highlight_matching_braces_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsHighlightMatchingBracesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("add_auto_brace_completion_pair")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnAddAutoBraceCompletionPair = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3186203200))
  }
  {
    methodName := StringNameFromStr("set_auto_brace_completion_pairs")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetAutoBraceCompletionPairs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4155329257))
  }
  {
    methodName := StringNameFromStr("get_auto_brace_completion_pairs")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetAutoBraceCompletionPairs = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3102165223))
  }
  {
    methodName := StringNameFromStr("has_auto_brace_completion_open_key")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnHasAutoBraceCompletionOpenKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("has_auto_brace_completion_close_key")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnHasAutoBraceCompletionCloseKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("get_auto_brace_completion_close_key")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetAutoBraceCompletionCloseKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3135753539))
  }
  {
    methodName := StringNameFromStr("set_draw_breakpoints_gutter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetDrawBreakpointsGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_drawing_breakpoints_gutter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsDrawingBreakpointsGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_draw_bookmarks_gutter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetDrawBookmarksGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_drawing_bookmarks_gutter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsDrawingBookmarksGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_draw_executing_lines_gutter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetDrawExecutingLinesGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_drawing_executing_lines_gutter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsDrawingExecutingLinesGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_line_as_breakpoint")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetLineAsBreakpoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("is_line_breakpointed")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsLineBreakpointed = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("clear_breakpointed_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnClearBreakpointedLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_breakpointed_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetBreakpointedLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
  }
  {
    methodName := StringNameFromStr("set_line_as_bookmarked")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetLineAsBookmarked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("is_line_bookmarked")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsLineBookmarked = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("clear_bookmarked_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnClearBookmarkedLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_bookmarked_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetBookmarkedLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
  }
  {
    methodName := StringNameFromStr("set_line_as_executing")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetLineAsExecuting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("is_line_executing")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsLineExecuting = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("clear_executing_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnClearExecutingLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_executing_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetExecutingLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
  }
  {
    methodName := StringNameFromStr("set_draw_line_numbers")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetDrawLineNumbers = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_draw_line_numbers_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsDrawLineNumbersEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_line_numbers_zero_padded")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetLineNumbersZeroPadded = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_line_numbers_zero_padded")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsLineNumbersZeroPadded = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_draw_fold_gutter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetDrawFoldGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_drawing_fold_gutter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsDrawingFoldGutter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_line_folding_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetLineFoldingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_line_folding_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsLineFoldingEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("can_fold_line")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnCanFoldLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("fold_line")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnFoldLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("unfold_line")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnUnfoldLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("fold_all_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnFoldAllLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("unfold_all_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnUnfoldAllLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("toggle_foldable_line")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnToggleFoldableLine = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("is_line_folded")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsLineFolded = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("get_folded_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetFoldedLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("create_code_region")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnCreateCodeRegion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_code_region_start_tag")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetCodeRegionStartTag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_code_region_end_tag")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetCodeRegionEndTag = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("set_code_region_tags")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetCodeRegionTags = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 708800718))
  }
  {
    methodName := StringNameFromStr("is_line_code_region_start")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsLineCodeRegionStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("is_line_code_region_end")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsLineCodeRegionEnd = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("add_string_delimiter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnAddStringDelimiter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3146098955))
  }
  {
    methodName := StringNameFromStr("remove_string_delimiter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnRemoveStringDelimiter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("has_string_delimiter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnHasStringDelimiter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("set_string_delimiters")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetStringDelimiters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("clear_string_delimiters")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnClearStringDelimiters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_string_delimiters")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetStringDelimiters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("is_in_string")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsInString = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 688195400))
  }
  {
    methodName := StringNameFromStr("add_comment_delimiter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnAddCommentDelimiter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3146098955))
  }
  {
    methodName := StringNameFromStr("remove_comment_delimiter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnRemoveCommentDelimiter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("has_comment_delimiter")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnHasCommentDelimiter = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3927539163))
  }
  {
    methodName := StringNameFromStr("set_comment_delimiters")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetCommentDelimiters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("clear_comment_delimiters")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnClearCommentDelimiters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_comment_delimiters")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetCommentDelimiters = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("is_in_comment")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsInComment = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 688195400))
  }
  {
    methodName := StringNameFromStr("get_delimiter_start_key")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetDelimiterStartKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
  }
  {
    methodName := StringNameFromStr("get_delimiter_end_key")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetDelimiterEndKey = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
  }
  {
    methodName := StringNameFromStr("get_delimiter_start_position")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetDelimiterStartPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3016396712))
  }
  {
    methodName := StringNameFromStr("get_delimiter_end_position")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetDelimiterEndPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3016396712))
  }
  {
    methodName := StringNameFromStr("set_code_hint")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetCodeHint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("set_code_hint_draw_below")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetCodeHintDrawBelow = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_text_for_code_completion")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetTextForCodeCompletion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("request_code_completion")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnRequestCodeCompletion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 107499316))
  }
  {
    methodName := StringNameFromStr("add_code_completion_option")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnAddCodeCompletionOption = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 947964390))
  }
  {
    methodName := StringNameFromStr("update_code_completion_options")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnUpdateCodeCompletionOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_code_completion_options")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetCodeCompletionOptions = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("get_code_completion_option")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetCodeCompletionOption = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3485342025))
  }
  {
    methodName := StringNameFromStr("get_code_completion_selected_index")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetCodeCompletionSelectedIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_code_completion_selected_index")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetCodeCompletionSelectedIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("confirm_code_completion")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnConfirmCodeCompletion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 107499316))
  }
  {
    methodName := StringNameFromStr("cancel_code_completion")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnCancelCodeCompletion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("set_code_completion_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetCodeCompletionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_code_completion_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsCodeCompletionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_code_completion_prefixes")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetCodeCompletionPrefixes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_code_completion_prefixes")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetCodeCompletionPrefixes = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("set_line_length_guidelines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetLineLengthGuidelines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
  }
  {
    methodName := StringNameFromStr("get_line_length_guidelines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetLineLengthGuidelines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
  }
  {
    methodName := StringNameFromStr("set_symbol_lookup_on_click_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetSymbolLookupOnClickEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_symbol_lookup_on_click_enabled")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnIsSymbolLookupOnClickEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_text_for_symbol_lookup")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetTextForSymbolLookup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
  }
  {
    methodName := StringNameFromStr("get_text_with_cursor_char")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnGetTextWithCursorChar = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1391810591))
  }
  {
    methodName := StringNameFromStr("set_symbol_lookup_word_as_valid")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnSetSymbolLookupWordAsValid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("duplicate_lines")
    defer methodName.Destroy()
    ptrsForCodeEdit.fnDuplicateLines = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
}

type CodeEdit struct {
  TextEdit
}

func (me *CodeEdit) BaseClass() string {
  return "CodeEdit"
}

func NewCodeEdit() *CodeEdit {
  str := StringNameFromStr("CodeEdit") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CodeEdit{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type CodeEditCodeCompletionKind int
const (
  CodeEditCodeCompletionKindKindClass CodeEditCodeCompletionKind = 0
  CodeEditCodeCompletionKindKindFunction CodeEditCodeCompletionKind = 1
  CodeEditCodeCompletionKindKindSignal CodeEditCodeCompletionKind = 2
  CodeEditCodeCompletionKindKindVariable CodeEditCodeCompletionKind = 3
  CodeEditCodeCompletionKindKindMember CodeEditCodeCompletionKind = 4
  CodeEditCodeCompletionKindKindEnum CodeEditCodeCompletionKind = 5
  CodeEditCodeCompletionKindKindConstant CodeEditCodeCompletionKind = 6
  CodeEditCodeCompletionKindKindNodePath CodeEditCodeCompletionKind = 7
  CodeEditCodeCompletionKindKindFilePath CodeEditCodeCompletionKind = 8
  CodeEditCodeCompletionKindKindPlainText CodeEditCodeCompletionKind = 9
)

type CodeEditCodeCompletionLocation int
const (
  CodeEditCodeCompletionLocationLocationLocal CodeEditCodeCompletionLocation = 0
  CodeEditCodeCompletionLocationLocationParentMask CodeEditCodeCompletionLocation = 256
  CodeEditCodeCompletionLocationLocationOtherUserCode CodeEditCodeCompletionLocation = 512
  CodeEditCodeCompletionLocationLocationOther CodeEditCodeCompletionLocation = 1024
)

func (me *CodeEdit) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CodeEdit) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CodeEdit) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CodeEdit) SetIndentSize(size int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetIndentSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetIndentSize() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetIndentSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetIndentUsingSpaces(use_spaces bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_spaces) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetIndentUsingSpaces), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsIndentUsingSpaces() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsIndentUsingSpaces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetAutoIndentEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetAutoIndentEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsAutoIndentEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsAutoIndentEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetAutoIndentPrefixes(prefixes []String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&prefixes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetAutoIndentPrefixes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetAutoIndentPrefixes() []String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetAutoIndentPrefixes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[String](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *CodeEdit) DoIndent()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnDoIndent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IndentLines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIndentLines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) UnindentLines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnUnindentLines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) ConvertIndent(from_line int64, to_line int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_line) , gdc.ConstTypePtr(&to_line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnConvertIndent), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) SetAutoBraceCompletionEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetAutoBraceCompletionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsAutoBraceCompletionEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsAutoBraceCompletionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetHighlightMatchingBracesEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetHighlightMatchingBracesEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsHighlightMatchingBracesEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsHighlightMatchingBracesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) AddAutoBraceCompletionPair(start_key String, end_key String, )  {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), end_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnAddAutoBraceCompletionPair), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) SetAutoBraceCompletionPairs(pairs Dictionary, )  {
  cargs := []gdc.ConstTypePtr{pairs.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetAutoBraceCompletionPairs), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetAutoBraceCompletionPairs() Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetAutoBraceCompletionPairs), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) HasAutoBraceCompletionOpenKey(open_key String, ) bool {
  cargs := []gdc.ConstTypePtr{open_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnHasAutoBraceCompletionOpenKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) HasAutoBraceCompletionCloseKey(close_key String, ) bool {
  cargs := []gdc.ConstTypePtr{close_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnHasAutoBraceCompletionCloseKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) GetAutoBraceCompletionCloseKey(open_key String, ) String {
  cargs := []gdc.ConstTypePtr{open_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetAutoBraceCompletionCloseKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) SetDrawBreakpointsGutter(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetDrawBreakpointsGutter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsDrawingBreakpointsGutter() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsDrawingBreakpointsGutter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetDrawBookmarksGutter(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetDrawBookmarksGutter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsDrawingBookmarksGutter() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsDrawingBookmarksGutter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetDrawExecutingLinesGutter(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetDrawExecutingLinesGutter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsDrawingExecutingLinesGutter() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsDrawingExecutingLinesGutter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetLineAsBreakpoint(line int64, breakpointed bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&breakpointed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetLineAsBreakpoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsLineBreakpointed(line int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsLineBreakpointed), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) ClearBreakpointedLines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnClearBreakpointedLines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetBreakpointedLines() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetBreakpointedLines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) SetLineAsBookmarked(line int64, bookmarked bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&bookmarked) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetLineAsBookmarked), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsLineBookmarked(line int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsLineBookmarked), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) ClearBookmarkedLines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnClearBookmarkedLines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetBookmarkedLines() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetBookmarkedLines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) SetLineAsExecuting(line int64, executing bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&executing) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetLineAsExecuting), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsLineExecuting(line int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsLineExecuting), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) ClearExecutingLines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnClearExecutingLines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetExecutingLines() PackedInt32Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetExecutingLines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) SetDrawLineNumbers(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetDrawLineNumbers), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsDrawLineNumbersEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsDrawLineNumbersEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetLineNumbersZeroPadded(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetLineNumbersZeroPadded), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsLineNumbersZeroPadded() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsLineNumbersZeroPadded), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetDrawFoldGutter(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetDrawFoldGutter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsDrawingFoldGutter() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsDrawingFoldGutter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetLineFoldingEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetLineFoldingEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsLineFoldingEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsLineFoldingEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) CanFoldLine(line int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnCanFoldLine), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) FoldLine(line int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnFoldLine), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) UnfoldLine(line int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnUnfoldLine), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) FoldAllLines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnFoldAllLines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) UnfoldAllLines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnUnfoldAllLines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) ToggleFoldableLine(line int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnToggleFoldableLine), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsLineFolded(line int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsLineFolded), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) GetFoldedLines() []int {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetFoldedLines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[int](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *CodeEdit) CreateCodeRegion()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnCreateCodeRegion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetCodeRegionStartTag() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetCodeRegionStartTag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) GetCodeRegionEndTag() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetCodeRegionEndTag), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) SetCodeRegionTags(start String, end String, )  {
  cargs := []gdc.ConstTypePtr{start.AsCTypePtr(), end.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetCodeRegionTags), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsLineCodeRegionStart(line int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsLineCodeRegionStart), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) IsLineCodeRegionEnd(line int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&line)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsLineCodeRegionEnd), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) AddStringDelimiter(start_key String, end_key String, line_only bool, )  {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), end_key.AsCTypePtr(), gdc.ConstTypePtr(&line_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnAddStringDelimiter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) RemoveStringDelimiter(start_key String, )  {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnRemoveStringDelimiter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) HasStringDelimiter(start_key String, ) bool {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnHasStringDelimiter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetStringDelimiters(string_delimiters []String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&string_delimiters) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetStringDelimiters), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) ClearStringDelimiters()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnClearStringDelimiters), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetStringDelimiters() []String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetStringDelimiters), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[String](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *CodeEdit) IsInString(line int64, column int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&line)
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsInString), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) AddCommentDelimiter(start_key String, end_key String, line_only bool, )  {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), end_key.AsCTypePtr(), gdc.ConstTypePtr(&line_only) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnAddCommentDelimiter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) RemoveCommentDelimiter(start_key String, )  {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnRemoveCommentDelimiter), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) HasCommentDelimiter(start_key String, ) bool {
  cargs := []gdc.ConstTypePtr{start_key.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnHasCommentDelimiter), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetCommentDelimiters(comment_delimiters []String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&comment_delimiters) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetCommentDelimiters), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) ClearCommentDelimiters()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnClearCommentDelimiters), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetCommentDelimiters() []String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetCommentDelimiters), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[String](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *CodeEdit) IsInComment(line int64, column int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&line)
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsInComment), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) GetDelimiterStartKey(delimiter_index int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delimiter_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&delimiter_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetDelimiterStartKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) GetDelimiterEndKey(delimiter_index int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delimiter_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&delimiter_index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetDelimiterEndKey), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) GetDelimiterStartPosition(line int64, column int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&line)
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetDelimiterStartPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) GetDelimiterEndPosition(line int64, column int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&line)
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetDelimiterEndPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) SetCodeHint(code_hint String, )  {
  cargs := []gdc.ConstTypePtr{code_hint.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetCodeHint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) SetCodeHintDrawBelow(draw_below bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_below) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetCodeHintDrawBelow), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetTextForCodeCompletion() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetTextForCodeCompletion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) RequestCodeCompletion(force bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnRequestCodeCompletion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) AddCodeCompletionOption(type_ CodeEditCodeCompletionKind, display_text String, insert_text String, text_color Color, icon Resource, value Variant, location int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , display_text.AsCTypePtr(), insert_text.AsCTypePtr(), text_color.AsCTypePtr(), icon.AsCTypePtr(), value.AsCTypePtr(), gdc.ConstTypePtr(&location) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnAddCodeCompletionOption), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) UpdateCodeCompletionOptions(force bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&force) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnUpdateCodeCompletionOptions), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetCodeCompletionOptions() []Dictionary {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetCodeCompletionOptions), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *CodeEdit) GetCodeCompletionOption(index int64, ) Dictionary {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetCodeCompletionOption), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) GetCodeCompletionSelectedIndex() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetCodeCompletionSelectedIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetCodeCompletionSelectedIndex(index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetCodeCompletionSelectedIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) ConfirmCodeCompletion(replace bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&replace) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnConfirmCodeCompletion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) CancelCodeCompletion()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnCancelCodeCompletion), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) SetCodeCompletionEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetCodeCompletionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsCodeCompletionEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsCodeCompletionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) SetCodeCompletionPrefixes(prefixes []String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&prefixes) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetCodeCompletionPrefixes), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetCodeCompletionPrefixes() []String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetCodeCompletionPrefixes), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[String](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *CodeEdit) SetLineLengthGuidelines(guideline_columns []int, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&guideline_columns) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetLineLengthGuidelines), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) GetLineLengthGuidelines() []int {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetLineLengthGuidelines), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[int](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *CodeEdit) SetSymbolLookupOnClickEnabled(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetSymbolLookupOnClickEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) IsSymbolLookupOnClickEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnIsSymbolLookupOnClickEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CodeEdit) GetTextForSymbolLookup() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetTextForSymbolLookup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) GetTextWithCursorChar(line int64, column int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&line) , gdc.ConstTypePtr(&column) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&line)
  pinner.Pin(&column)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnGetTextWithCursorChar), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CodeEdit) SetSymbolLookupWordAsValid(valid bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&valid) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnSetSymbolLookupWordAsValid), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CodeEdit) DuplicateLines()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCodeEdit.fnDuplicateLines), me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type CodeEditBreakpointToggledSignalFn func(line int, )

func (me *CodeEdit) ConnectBreakpointToggled(subs SignalSubscribers, fn CodeEditBreakpointToggledSignalFn) {
  sig := StringNameFromStr("breakpoint_toggled")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CodeEdit) DisconnectBreakpointToggled(subs SignalSubscribers, fn CodeEditBreakpointToggledSignalFn) {
  sig := StringNameFromStr("breakpoint_toggled")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CodeEditCodeCompletionRequestedSignalFn func()

func (me *CodeEdit) ConnectCodeCompletionRequested(subs SignalSubscribers, fn CodeEditCodeCompletionRequestedSignalFn) {
  sig := StringNameFromStr("code_completion_requested")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CodeEdit) DisconnectCodeCompletionRequested(subs SignalSubscribers, fn CodeEditCodeCompletionRequestedSignalFn) {
  sig := StringNameFromStr("code_completion_requested")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CodeEditSymbolLookupSignalFn func(symbol String, line int, column int, )

func (me *CodeEdit) ConnectSymbolLookup(subs SignalSubscribers, fn CodeEditSymbolLookupSignalFn) {
  sig := StringNameFromStr("symbol_lookup")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CodeEdit) DisconnectSymbolLookup(subs SignalSubscribers, fn CodeEditSymbolLookupSignalFn) {
  sig := StringNameFromStr("symbol_lookup")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type CodeEditSymbolValidateSignalFn func(symbol String, )

func (me *CodeEdit) ConnectSymbolValidate(subs SignalSubscribers, fn CodeEditSymbolValidateSignalFn) {
  sig := StringNameFromStr("symbol_validate")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *CodeEdit) DisconnectSymbolValidate(subs SignalSubscribers, fn CodeEditSymbolValidateSignalFn) {
  sig := StringNameFromStr("symbol_validate")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
