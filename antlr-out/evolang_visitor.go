// Code generated from ./EvoLang.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLang

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EvoLangParser.
type EvoLangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EvoLangParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by EvoLangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by EvoLangParser#modelDef.
	VisitModelDef(ctx *ModelDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#modelBody.
	VisitModelBody(ctx *ModelBodyContext) interface{}

	// Visit a parse tree produced by EvoLangParser#fieldDef.
	VisitFieldDef(ctx *FieldDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#eventDef.
	VisitEventDef(ctx *EventDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#eventBody.
	VisitEventBody(ctx *EventBodyContext) interface{}

	// Visit a parse tree produced by EvoLangParser#enumDef.
	VisitEnumDef(ctx *EnumDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#accessControlDef.
	VisitAccessControlDef(ctx *AccessControlDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#accessRule.
	VisitAccessRule(ctx *AccessRuleContext) interface{}

	// Visit a parse tree produced by EvoLangParser#action.
	VisitAction(ctx *ActionContext) interface{}

	// Visit a parse tree produced by EvoLangParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by EvoLangParser#endpointDef.
	VisitEndpointDef(ctx *EndpointDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#endpointBody.
	VisitEndpointBody(ctx *EndpointBodyContext) interface{}

	// Visit a parse tree produced by EvoLangParser#serverDef.
	VisitServerDef(ctx *ServerDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#serverBody.
	VisitServerBody(ctx *ServerBodyContext) interface{}

	// Visit a parse tree produced by EvoLangParser#bodyDef.
	VisitBodyDef(ctx *BodyDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#authenticateDef.
	VisitAuthenticateDef(ctx *AuthenticateDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#returnDef.
	VisitReturnDef(ctx *ReturnDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#sandbox.
	VisitSandbox(ctx *SandboxContext) interface{}

	// Visit a parse tree produced by EvoLangParser#goCode.
	VisitGoCode(ctx *GoCodeContext) interface{}

	// Visit a parse tree produced by EvoLangParser#typeDef.
	VisitTypeDef(ctx *TypeDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#arrayDef.
	VisitArrayDef(ctx *ArrayDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#refDef.
	VisitRefDef(ctx *RefDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#eventType.
	VisitEventType(ctx *EventTypeContext) interface{}

	// Visit a parse tree produced by EvoLangParser#ifCondition.
	VisitIfCondition(ctx *IfConditionContext) interface{}

	// Visit a parse tree produced by EvoLangParser#MulDivExpr.
	VisitMulDivExpr(ctx *MulDivExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by EvoLangParser#BooleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by EvoLangParser#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by EvoLangParser#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by EvoLangParser#RelationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#LogicalExpr.
	VisitLogicalExpr(ctx *LogicalExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#exprList.
	VisitExprList(ctx *ExprListContext) interface{}

	// Visit a parse tree produced by EvoLangParser#mainFunction.
	VisitMainFunction(ctx *MainFunctionContext) interface{}

	// Visit a parse tree produced by EvoLangParser#mainStatements.
	VisitMainStatements(ctx *MainStatementsContext) interface{}

	// Visit a parse tree produced by EvoLangParser#mainStatement.
	VisitMainStatement(ctx *MainStatementContext) interface{}

	// Visit a parse tree produced by EvoLangParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by EvoLangParser#serverStatement.
	VisitServerStatement(ctx *ServerStatementContext) interface{}
}
