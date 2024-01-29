package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	parser "github.com/ZacxDev/evolang/antlr-out"
	"github.com/antlr4-go/antlr/v4"
	"golang.org/x/tools/imports"
)

type EvoLangVisitor struct {
	*parser.BaseEvoLangParserVisitor
	// Add fields to store state as needed
}

func NewEvoLangVisitor() *EvoLangVisitor {
	return &EvoLangVisitor{
		BaseEvoLangParserVisitor: &parser.BaseEvoLangParserVisitor{
			BaseParseTreeVisitor: &antlr.BaseParseTreeVisitor{},
		}, // Proper initialization
	}
}

// Implement other visit methods...

func (v *EvoLangVisitor) VisitProg(ctx *parser.ProgContext) interface{} {
	var goCode string
	for _, child := range ctx.GetChildren() {
		if childResult := v.Visit(child.(antlr.ParseTree)); childResult != nil {
			goCode += childResult.(string)
		}
	}
	return goCode
}

func (v *EvoLangVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	var goCode string

	// Delegate to the appropriate visitor method based on the statement type
	if ctx.ModelDef() != nil {
		modelDefCtx := ctx.ModelDef().(*parser.ModelDefContext) // Type assertion
		goCode = v.VisitModelDef(modelDefCtx).(string)
	} else if ctx.EnumDef() != nil {
		enumDefCtx := ctx.EnumDef().(*parser.EnumDefContext) // Type assertion
		goCode = v.VisitEnumDef(enumDefCtx).(string)
	} else if ctx.MainFunction() != nil {
		mainFunctionCtx := ctx.MainFunction().(*parser.MainFunctionContext) // Type assertion
		goCode = v.VisitMainFunction(mainFunctionCtx).(string)
	}

	// Return the generated Go code for the statement
	return goCode
}

func (v *EvoLangVisitor) VisitModelBody(ctx *parser.ModelBodyContext) interface{} {
	var goCode string
	for _, child := range ctx.GetChildren() {
		// Depending on the structure of a model body in your DSL,
		// delegate to the appropriate visitor methods.
		if fieldCtx, ok := child.(*parser.ModelFieldDefContext); ok {
			goCode += v.VisitModelFieldDef(fieldCtx).(string) + "\n"
		} else if eventCtx, ok := child.(*parser.EventDefContext); ok {
			goCode += v.VisitEventDef(eventCtx).(string) + "\n"
		}
		// ... handle other constructs within a model body ...
	}
	return goCode
}

func (v *EvoLangVisitor) VisitModelFieldDef(ctx *parser.ModelFieldDefContext) interface{} {
	// Example: Check if the type of the field is valid
	fieldType := ctx.TypeDef().GetText()
	if !isValidType(fieldType) {
		log.Fatalf("Invalid type: %s", fieldType)
	}
	return v.VisitChildren(ctx)
}

// Helper function to validate types
func isValidType(t string) bool {
	// Define logic to validate type
	return true
}

func (v *EvoLangVisitor) VisitModelDef(ctx *parser.ModelDefContext) interface{} {
	modelName := ctx.ID().GetText()
	goCode := "type " + modelName + " struct {\n"

	// Correctly access the field definitions
	for _, fieldCtxInterface := range ctx.ModelBody().AllModelFieldDef() {
		// Type assertion to convert IFieldDefContext to *FieldDefContext
		fieldCtx := fieldCtxInterface.(*parser.ModelFieldDefContext)
		fieldCode := v.VisitModelFieldDef(fieldCtx).(string)
		goCode += "\t" + fieldCode + "\n"
	}

	goCode += "}\n"
	return goCode
}

func (v *EvoLangVisitor) VisitEnumDef(ctx *parser.EnumDefContext) interface{} {
	enumName := ctx.ID(0).GetText() // Get the first ID, which should be the enum name
	goCode := "type " + enumName + " int\nconst (\n"

	enumVals := ctx.AllID()[1:] // Skip the first ID as it's the enum name
	for i, enumVal := range enumVals {
		goCode += "\t" + enumVal.GetText() + " " + enumName
		if i == 0 {
			goCode += " = iota"
		}
		goCode += "\n"
	}

	goCode += ")\n"
	return goCode
}

func (v *EvoLangVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var result string
	for _, child := range node.GetChildren() {
		// Type assert child to antlr.ParseTree
		if parseTree, ok := child.(antlr.ParseTree); ok {
			childResult := parseTree.Accept(v)
			if childResult != nil {
				result += childResult.(string)
			}
		}
	}
	return result
}

func (v *EvoLangVisitor) VisitEventDef(ctx *parser.EventDefContext) interface{} {
	// Process event definition
	// This will depend on how events are structured in your DSL
	return v.VisitChildren(ctx)
}

// Similar methods for each expression and literal context
func (v *EvoLangVisitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	return ctx.GetText()
}

func (v *EvoLangVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	return ctx.GetText()
}

func (v *EvoLangVisitor) VisitTypeDef(ctx *parser.TypeDefContext) interface{} {
	// Process type definitions
	return ctx.GetText()
}

func (v *EvoLangVisitor) VisitMulDivExpr(ctx *parser.MulDivExprContext) interface{} {
	left := v.Visit(ctx.GetLeft()).(string)
	right := v.Visit(ctx.GetRight()).(string)
	op := ctx.GetOp().GetText()

	return left + " " + op + " " + right
}

func (v *EvoLangVisitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	left := v.Visit(ctx.GetLeft()).(string)
	right := v.Visit(ctx.GetRight()).(string)
	op := ctx.GetOp().GetText()

	return left + " " + op + " " + right
}

func (v *EvoLangVisitor) VisitLogicalExpr(ctx *parser.LogicalExprContext) interface{} {
	left := v.Visit(ctx.GetLeft()).(string)
	right := v.Visit(ctx.GetRight()).(string)
	op := ctx.GetOp().GetText()

	return left + " " + op + " " + right
}

func (v *EvoLangVisitor) VisitRelationalExpr(ctx *parser.RelationalExprContext) interface{} {
	left := v.Visit(ctx.GetLeft()).(string)
	right := v.Visit(ctx.GetRight()).(string)
	op := ctx.GetOp().GetText()

	return left + " " + op + " " + right
}

func (v *EvoLangVisitor) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	inner := v.Visit(ctx.Expression()).(string)
	return "(" + inner + ")"
}

func (v *EvoLangVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	functionName := ctx.ID().GetText()

	// Perform a type assertion on ctx.ExprList()
	exprListCtx := ctx.ExprList().(*parser.ExprListContext)
	arguments := v.VisitExprList(exprListCtx).(string)

	return functionName + "(" + arguments + ")"
}

func (v *EvoLangVisitor) VisitExprList(ctx *parser.ExprListContext) interface{} {
	var args []string
	for _, expr := range ctx.AllExpression() {
		args = append(args, v.Visit(expr).(string))
	}
	return strings.Join(args, ", ")
}

func (v *EvoLangVisitor) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {
	return ctx.GetText()
}

func (v *EvoLangVisitor) VisitBooleanLiteral(ctx *parser.BooleanLiteralContext) interface{} {
	return ctx.GetText()
}

func (v *EvoLangVisitor) VisitPrintStatement(ctx *parser.PrintStatementContext) interface{} {
	text := ctx.STRING().GetText()
	// Remove the surrounding quotes from the string
	unquotedText := text[1 : len(text)-1]
	return "fmt.Println(\"" + unquotedText + "\")"
}

func (v *EvoLangVisitor) VisitServerStatement(ctx *parser.ServerStatementContext) interface{} {
	serverName := ctx.ID().GetText() // Get the server name
	port := ctx.INT().GetText()      // Get the port number

	// Generate the Go code to set up the server and listen on the specified port
	goCode := fmt.Sprintf("setup%sServer()\n", serverName) // Assuming a setup function for the server
	goCode += fmt.Sprintf("http.ListenAndServe(\":%s\", nil)\n", port)

	return goCode
}

func (v *EvoLangVisitor) VisitEventChannelHandler(ctx *parser.EventChannelHandlerContext) interface{} {
	// Extract the channel name, old value, new value identifiers, and sandbox content
	channelName := ctx.ID(0).GetText() // Assuming first ID is the channel name
	oldValue := ctx.ID(1).GetText()    // Assuming second ID is the old value identifier
	newValue := ctx.ID(2).GetText()    // Assuming third ID is the new value identifier

	// Handle the sandbox content
	sandboxContent := ctx.SANDBOX().GetText()
	return "TODO: implement this lol" + channelName + oldValue + newValue + sandboxContent

	// Process the extracted information as needed
	// ...

	// Return any specific result or simply nil
	return nil
}

// VisitAction handles actions defined in access control.
func (v *EvoLangVisitor) VisitAction(ctx *parser.ActionContext) interface{} {
	// Translate action into a Go function or method call
	return ""
}

// VisitParamList handles parameter lists, possibly in function calls.
func (v *EvoLangVisitor) VisitParamList(ctx *parser.ParamListContext) interface{} {
	// Translate parameter list into Go function parameters
	return ""
}

// VisitBodyDef handles body definitions in endpoints.
func (v *EvoLangVisitor) VisitBodyDef(ctx *parser.BodyDefContext) interface{} {
	// Translate body definition into Go code
	return ""
}

// VisitAuthenticateDef handles authentication definitions in endpoints.
func (v *EvoLangVisitor) VisitAuthenticateDef(ctx *parser.AuthenticateDefContext) interface{} {
	// Implement logic for authentication, translating into Go code
	return ""
}

// VisitReturnDef handles return definitions in endpoints.
func (v *EvoLangVisitor) VisitReturnDef(ctx *parser.ReturnDefContext) interface{} {
	// Translate return definition into Go code
	return ""
}

// VisitArrayDef handles array type definitions.
func (v *EvoLangVisitor) VisitArrayDef(ctx *parser.ArrayDefContext) interface{} {
	// Translate array definitions into Go slice types
	return "[]" + v.Visit(ctx.TypeDef()).(string)
}

// VisitEventType handles event type definitions.
func (v *EvoLangVisitor) VisitEventType(ctx *parser.EventTypeContext) interface{} {
	// Translate event types into Go code
	return ""
}

// VisitMainStatements handles statements in the main function.
func (v *EvoLangVisitor) VisitMainStatements(ctx *parser.MainStatementsContext) interface{} {
	var code string
	for _, stmt := range ctx.AllMainStatement() {
		code += v.Visit(stmt).(string) + "\n"
	}
	return code
}

// VisitMainStatement handles a single statement in the main function.
func (v *EvoLangVisitor) VisitMainStatement(ctx *parser.MainStatementContext) interface{} {
	// Translate main statements into Go code
	return v.VisitChildren(ctx)
}

func (v *EvoLangVisitor) VisitMainFunction(ctx *parser.MainFunctionContext) interface{} {
	goCode := "func main() {\n"
	for _, stmtCtx := range ctx.MainStatements().AllMainStatement() {
		code, ok := v.Visit(stmtCtx).(string)
		if ok {
			goCode += code + "\n"
		}
	}
	goCode += "}\n"
	return goCode
}

func (v *EvoLangVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

const goImports = `package main

import (
    "fmt"
    // other standard imports
)

`

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <file.evo>", os.Args[0])
	}

	// Read EvoLang source file
	filename := os.Args[1]
	source, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	// Parse the source
	inputStream := antlr.NewInputStream(string(source))
	lexer := parser.NewEvoLangLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewEvoLangParser(tokenStream)

	// Use the visitor to generate Go code
	visitor := NewEvoLangVisitor()
	goCode := visitor.Visit(p.Prog()).(string)

	goCodeWithPkgAndStaticImports := goImports + goCode

	outputDir := ".evo-out"
	goFileName := "output.go"
	outputFile := filepath.Join(outputDir, goFileName)

	// Create the output directory if it does not exist
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatalf("Error creating output directory: %s", err)
	}

	// Step 1: Write goCode to the file in the output directory
	err = os.WriteFile(outputFile, []byte(goCodeWithPkgAndStaticImports), 0644)
	if err != nil {
		log.Fatalf("Error writing Go code to file: %s", err)
	}

	// Process the file with goimports
	opt := &imports.Options{
		Comments:   true,
		TabIndent:  true,
		TabWidth:   8,
		FormatOnly: false,
	}

	res, err := imports.Process(outputFile, nil, opt)
	if err != nil {
		log.Fatalf("goimports processing failed: %s", err)
	}

	// Write the processed code back to the file
	err = os.WriteFile(outputFile, res, 0644)
	if err != nil {
		log.Fatalf("Error writing processed code to file: %s", err)
	}

	// Step 2: Compile the Go file into an executable
	execName := filepath.Join(outputDir, "output") // Name of the executable
	cmd := exec.Command("go", "build", "-o", execName, outputFile)
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error compiling Go code: %s", err)
	}

	fmt.Println("Compiled successfully. Executable: ", execName)
}
