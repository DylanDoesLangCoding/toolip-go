package token

type TokenType string

type Token struct {
	Type  TokenType
	Value string
}

const (
	ILLEGAL = "ILLEGAL"
	NEWLINE = "NEWLINE"
	EOF     = "EOF"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER"
	INTVAL     = "INTVAL"
	FLOAT      = "FLOAT"
	CHAR       = "CHAR"
	STRVAL     = "STRVAL"

	// Operators
	EQUALS     = "=="
	ASSIGN     = "="
	INCREMENT  = "++"
	DECREMENT  = "--"
	LSHIFTEQ   = "<<="
	RSHIFTEQ   = ">>="
	PLUSEQ     = "+="
	MINUSEQ    = "-="
	EXPOEQ     = "**="
	MULTEQ     = "*="
	FDIVEQ     = "//="
	DIVEQ      = "/="
	MODEQ      = "%="
	CONCATEQ   = "..="
	COALESCEEQ = "??="
	LTEQ       = "<="
	GTEQ       = ">="
	BOOLNOTEQ  = "!="
	BITANDEQ   = "&="
	BITOREQ    = "|="
	BITXOREQ   = "^="
	LSHIFT     = "<<"
	RSHIFT     = ">>"
	PLUS       = "+"
	MINUS      = "-"
	EXPO       = "**"
	MULT       = "*"
	FDIV       = "//"
	DIV        = "/"
	MOD        = "%"
	CONCAT     = ".."
	COALESCE   = "??"
	TERNARY    = "?"
	LT         = "<"
	GT         = ">"
	BOOLNOT    = "!"
	BITNOT     = "~" //unary not
	BITAND     = "&"
	BITOR      = "|"
	BITXOR     = "^"
	HASH       = "#" //length of attached function, data structure or string

	//Delimiters
	SEMICOLON = ";"
	COMMA     = ","
	DOT       = "."
	COLON     = ":"
	ARROW     = "->"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"

	//Keywords
	GLOBAL      = "GLOBAL"
	CONST       = "CONST"
	STATIC      = "STATIC"
	DYNAMIC     = "DYNAMIC"
	UNSAFE      = "UNSAFE"
	FUNC        = "FUNC"
	END         = "END"
	SELF        = "SELF"
	FOR         = "FOR"
	DO          = "DO"
	WHILE       = "WHILE"
	LOOP        = "LOOP"
	BREAK       = "BREAK"
	TO          = "TO"
	ETC         = "ETC"
	SWITCH      = "SWITCH"
	CASE        = "CASE"
	INT8TYPE    = "INT8TYPE"
	INT16TYPE   = "INT16TYPE"
	INT32TYPE   = "INT32TYPE"
	INT64TYPE   = "INT64TYPE"
	UINT8TYPE   = "UINT8TYPE"
	UINT16TYPE  = "UINT16TYPE"
	UINT32TYPE  = "UINT32TYPE"
	UINT64TYPE  = "UINT64TYPE"
	FLOAT32TYPE = "FLOAT32TYPE"
	FLOAT64TYPE = "FLOAT64TYPE"
	BOOLTYPE    = "BOOLTYPE"
	CHARTYPE    = "CHARTYPE"
	STRTYPE     = "STRTYPE"
	VOID        = "VOID"
	TABLE       = "TABLE" /* acts as either an array or hash map*/
	SET         = "SET"   /* a stack-queue hybrid */
	LIST        = "LIST"
	TWOLIST     = "TWOLIST"
	ENUM        = "ENUM"
	TYPEDEF     = "TYPEDEF"
	STRUCT      = "STRUCT"
	CLASS       = "CLASS"
	RETURN      = "RETURN"
	AND         = "AND"
	OR          = "OR"
	NOT         = "NOT"
	XOR         = "XOR"
	TRUE        = "TRUE"
	FALSE       = "FALSE"
	NIL         = "NIL"
)

var keywords = map[string]TokenType{
	"global":  GLOBAL,
	"const":   CONST,
	"static":  STATIC,
	"dynamic": DYNAMIC,
	"unsafe":  UNSAFE,
	"func":    FUNC,
	"end":     END,
	"self":    SELF,
	"for":     FOR,
	"do":      DO,
	"while":   WHILE,
	"loop":    LOOP,
	"switch":  SWITCH,
	"case":    CASE,
	"type":    TYPEDEF,
	"struct":  STRUCT,
	"class":   CLASS,
	"table":   TABLE,
	"set":     SET,
	"list":    LIST,
	"twolist": TWOLIST,
	"enum":    ENUM,
	"to":      TO,
	"etc":     ETC,
	"string":  STRTYPE,
	"char":    CHARTYPE,
	"int8":    INT8TYPE,
	"int16":   INT16TYPE,
	"int32":   INT32TYPE,
	"int64":   INT64TYPE,
	"uint8":   UINT8TYPE,
	"uint16":  UINT16TYPE,
	"uint32":  UINT32TYPE,
	"uint64":  UINT64TYPE,
	"flt32":   FLOAT32TYPE,
	"flt64":   FLOAT64TYPE,
	"bool":    BOOLTYPE,
	"void":    VOID,
	"return":  RETURN,
	"break":   BREAK,
	"and":     AND,
	"or":      OR,
	"xor":     XOR,
	"not":     NOT,
	"true":    TRUE,
	"false":   FALSE,
	"nil":     NIL,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
