// Code generated by "stringer -type=Type"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ILLEGAL-0]
	_ = x[EOL-1]
	_ = x[IDENT-2]
	_ = x[INT-3]
	_ = x[FLOAT-4]
	_ = x[ASSIGN-5]
	_ = x[PLUS-6]
	_ = x[MINUS-7]
	_ = x[BANG-8]
	_ = x[ASTERISK-9]
	_ = x[SLASH-10]
	_ = x[PERCENT-11]
	_ = x[LT-12]
	_ = x[GT-13]
	_ = x[LTEQ-14]
	_ = x[GTEQ-15]
	_ = x[EQ-16]
	_ = x[NOTEQ-17]
	_ = x[COMMA-18]
	_ = x[SEMICOLON-19]
	_ = x[LPAREN-20]
	_ = x[RPAREN-21]
	_ = x[LBRACE-22]
	_ = x[RBRACE-23]
	_ = x[LBRACKET-24]
	_ = x[RBRACKET-25]
	_ = x[COLON-26]
	_ = x[LINECOMMENT-27]
	_ = x[STARTCOMMENT-28]
	_ = x[ENDCOMMENT-29]
	_ = x[FUNCTION-30]
	_ = x[TRUE-31]
	_ = x[FALSE-32]
	_ = x[IF-33]
	_ = x[ELSE-34]
	_ = x[RETURN-35]
	_ = x[STRING-36]
	_ = x[MACRO-37]
	_ = x[LEN-38]
	_ = x[FIRST-39]
	_ = x[REST-40]
	_ = x[PRINT-41]
	_ = x[LOG-42]
	_ = x[ERROR-43]
	_ = x[EOF-44]
}

const _Type_name = "ILLEGALEOLIDENTINTFLOATASSIGNPLUSMINUSBANGASTERISKSLASHPERCENTLTGTLTEQGTEQEQNOTEQCOMMASEMICOLONLPARENRPARENLBRACERBRACELBRACKETRBRACKETCOLONLINECOMMENTSTARTCOMMENTENDCOMMENTFUNCTIONTRUEFALSEIFELSERETURNSTRINGMACROLENFIRSTRESTPRINTLOGERROREOF"

var _Type_index = [...]uint8{0, 7, 10, 15, 18, 23, 29, 33, 38, 42, 50, 55, 62, 64, 66, 70, 74, 76, 81, 86, 95, 101, 107, 113, 119, 127, 135, 140, 151, 163, 173, 181, 185, 190, 192, 196, 202, 208, 213, 216, 221, 225, 230, 233, 238, 241}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
