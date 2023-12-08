// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CodeEdit struct {
  obj gdc.ObjectPtr
}

func (me *CodeEdit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CodeEdit) BaseClass() string {
  return "CodeEdit"
}

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

func  (me *CodeEdit) XConfirmCodeCompletion(replace bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) XRequestCodeCompletion(force bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) XFilterCodeCompletionCandidates(candidates Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetIndentSize(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetIndentSize() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetIndentUsingSpaces(use_spaces bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsIndentUsingSpaces() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetAutoIndentEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsAutoIndentEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetAutoIndentPrefixes(prefixes String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetAutoIndentPrefixes() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) DoIndent() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IndentLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) UnindentLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) ConvertIndent(from_line int, to_line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetAutoBraceCompletionEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsAutoBraceCompletionEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetHighlightMatchingBracesEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsHighlightMatchingBracesEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) AddAutoBraceCompletionPair(start_key String, end_key String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetAutoBraceCompletionPairs(pairs Dictionary, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetAutoBraceCompletionPairs() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) HasAutoBraceCompletionOpenKey(open_key String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) HasAutoBraceCompletionCloseKey(close_key String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetAutoBraceCompletionCloseKey(open_key String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetDrawBreakpointsGutter(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsDrawingBreakpointsGutter() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetDrawBookmarksGutter(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsDrawingBookmarksGutter() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetDrawExecutingLinesGutter(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsDrawingExecutingLinesGutter() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetLineAsBreakpoint(line int, breakpointed bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsLineBreakpointed(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) ClearBreakpointedLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetBreakpointedLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetLineAsBookmarked(line int, bookmarked bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsLineBookmarked(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) ClearBookmarkedLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetBookmarkedLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetLineAsExecuting(line int, executing bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsLineExecuting(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) ClearExecutingLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetExecutingLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetDrawLineNumbers(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsDrawLineNumbersEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetLineNumbersZeroPadded(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsLineNumbersZeroPadded() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetDrawFoldGutter(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsDrawingFoldGutter() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetLineFoldingEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsLineFoldingEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) CanFoldLine(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) FoldLine(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) UnfoldLine(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) FoldAllLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) UnfoldAllLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) ToggleFoldableLine(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsLineFolded(line int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetFoldedLines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) AddStringDelimiter(start_key String, end_key String, line_only bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) RemoveStringDelimiter(start_key String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) HasStringDelimiter(start_key String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetStringDelimiters(string_delimiters String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) ClearStringDelimiters() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetStringDelimiters() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsInString(line int, column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) AddCommentDelimiter(start_key String, end_key String, line_only bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) RemoveCommentDelimiter(start_key String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) HasCommentDelimiter(start_key String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetCommentDelimiters(comment_delimiters String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) ClearCommentDelimiters() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetCommentDelimiters() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsInComment(line int, column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetDelimiterStartKey(delimiter_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetDelimiterEndKey(delimiter_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetDelimiterStartPosition(line int, column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetDelimiterEndPosition(line int, column int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetCodeHint(code_hint String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetCodeHintDrawBelow(draw_below bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetTextForCodeCompletion() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) RequestCodeCompletion(force bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) AddCodeCompletionOption(type_ CodeEditCodeCompletionKind, display_text String, insert_text String, text_color Color, icon Resource, value Variant, location int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) UpdateCodeCompletionOptions(force bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetCodeCompletionOptions() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetCodeCompletionOption(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetCodeCompletionSelectedIndex() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetCodeCompletionSelectedIndex(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) ConfirmCodeCompletion(replace bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) CancelCodeCompletion() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetCodeCompletionEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsCodeCompletionEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetCodeCompletionPrefixes(prefixes String, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetCodeCompletionPrefixes() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetLineLengthGuidelines(guideline_columns int, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetLineLengthGuidelines() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetSymbolLookupOnClickEnabled(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) IsSymbolLookupOnClickEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) GetTextForSymbolLookup() { // TODO: return value
  // TODO: implement
}

func  (me *CodeEdit) SetSymbolLookupWordAsValid(valid bool, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
