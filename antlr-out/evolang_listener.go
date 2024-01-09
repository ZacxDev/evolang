// Code generated from ./EvoLang.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLang

import "github.com/antlr4-go/antlr/v4"

// EvoLangListener is a complete listener for a parse tree produced by EvoLangParser.
type EvoLangListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterModelDef is called when entering the modelDef production.
	EnterModelDef(c *ModelDefContext)

	// EnterModelBody is called when entering the modelBody production.
	EnterModelBody(c *ModelBodyContext)

	// EnterFieldDef is called when entering the fieldDef production.
	EnterFieldDef(c *FieldDefContext)

	// EnterEventDef is called when entering the eventDef production.
	EnterEventDef(c *EventDefContext)

	// EnterEventBody is called when entering the eventBody production.
	EnterEventBody(c *EventBodyContext)

	// EnterEnumDef is called when entering the enumDef production.
	EnterEnumDef(c *EnumDefContext)

	// EnterAccessControlDef is called when entering the accessControlDef production.
	EnterAccessControlDef(c *AccessControlDefContext)

	// EnterAccessRule is called when entering the accessRule production.
	EnterAccessRule(c *AccessRuleContext)

	// EnterAction is called when entering the action production.
	EnterAction(c *ActionContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterEndpointDef is called when entering the endpointDef production.
	EnterEndpointDef(c *EndpointDefContext)

	// EnterEndpointBody is called when entering the endpointBody production.
	EnterEndpointBody(c *EndpointBodyContext)

	// EnterServerDef is called when entering the serverDef production.
	EnterServerDef(c *ServerDefContext)

	// EnterServerBody is called when entering the serverBody production.
	EnterServerBody(c *ServerBodyContext)

	// EnterBodyDef is called when entering the bodyDef production.
	EnterBodyDef(c *BodyDefContext)

	// EnterAuthenticateDef is called when entering the authenticateDef production.
	EnterAuthenticateDef(c *AuthenticateDefContext)

	// EnterReturnDef is called when entering the returnDef production.
	EnterReturnDef(c *ReturnDefContext)

	// EnterSandbox is called when entering the sandbox production.
	EnterSandbox(c *SandboxContext)

	// EnterGoCode is called when entering the goCode production.
	EnterGoCode(c *GoCodeContext)

	// EnterTypeDef is called when entering the typeDef production.
	EnterTypeDef(c *TypeDefContext)

	// EnterArrayDef is called when entering the arrayDef production.
	EnterArrayDef(c *ArrayDefContext)

	// EnterRefDef is called when entering the refDef production.
	EnterRefDef(c *RefDefContext)

	// EnterEventType is called when entering the eventType production.
	EnterEventType(c *EventTypeContext)

	// EnterIfCondition is called when entering the ifCondition production.
	EnterIfCondition(c *IfConditionContext)

	// EnterMulDivExpr is called when entering the MulDivExpr production.
	EnterMulDivExpr(c *MulDivExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBooleanLiteral is called when entering the BooleanLiteral production.
	EnterBooleanLiteral(c *BooleanLiteralContext)

	// EnterIntLiteral is called when entering the IntLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterRelationalExpr is called when entering the RelationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterAddSubExpr is called when entering the AddSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterLogicalExpr is called when entering the LogicalExpr production.
	EnterLogicalExpr(c *LogicalExprContext)

	// EnterExprList is called when entering the exprList production.
	EnterExprList(c *ExprListContext)

	// EnterMainFunction is called when entering the mainFunction production.
	EnterMainFunction(c *MainFunctionContext)

	// EnterMainStatements is called when entering the mainStatements production.
	EnterMainStatements(c *MainStatementsContext)

	// EnterMainStatement is called when entering the mainStatement production.
	EnterMainStatement(c *MainStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterServerStatement is called when entering the serverStatement production.
	EnterServerStatement(c *ServerStatementContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitModelDef is called when exiting the modelDef production.
	ExitModelDef(c *ModelDefContext)

	// ExitModelBody is called when exiting the modelBody production.
	ExitModelBody(c *ModelBodyContext)

	// ExitFieldDef is called when exiting the fieldDef production.
	ExitFieldDef(c *FieldDefContext)

	// ExitEventDef is called when exiting the eventDef production.
	ExitEventDef(c *EventDefContext)

	// ExitEventBody is called when exiting the eventBody production.
	ExitEventBody(c *EventBodyContext)

	// ExitEnumDef is called when exiting the enumDef production.
	ExitEnumDef(c *EnumDefContext)

	// ExitAccessControlDef is called when exiting the accessControlDef production.
	ExitAccessControlDef(c *AccessControlDefContext)

	// ExitAccessRule is called when exiting the accessRule production.
	ExitAccessRule(c *AccessRuleContext)

	// ExitAction is called when exiting the action production.
	ExitAction(c *ActionContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitEndpointDef is called when exiting the endpointDef production.
	ExitEndpointDef(c *EndpointDefContext)

	// ExitEndpointBody is called when exiting the endpointBody production.
	ExitEndpointBody(c *EndpointBodyContext)

	// ExitServerDef is called when exiting the serverDef production.
	ExitServerDef(c *ServerDefContext)

	// ExitServerBody is called when exiting the serverBody production.
	ExitServerBody(c *ServerBodyContext)

	// ExitBodyDef is called when exiting the bodyDef production.
	ExitBodyDef(c *BodyDefContext)

	// ExitAuthenticateDef is called when exiting the authenticateDef production.
	ExitAuthenticateDef(c *AuthenticateDefContext)

	// ExitReturnDef is called when exiting the returnDef production.
	ExitReturnDef(c *ReturnDefContext)

	// ExitSandbox is called when exiting the sandbox production.
	ExitSandbox(c *SandboxContext)

	// ExitGoCode is called when exiting the goCode production.
	ExitGoCode(c *GoCodeContext)

	// ExitTypeDef is called when exiting the typeDef production.
	ExitTypeDef(c *TypeDefContext)

	// ExitArrayDef is called when exiting the arrayDef production.
	ExitArrayDef(c *ArrayDefContext)

	// ExitRefDef is called when exiting the refDef production.
	ExitRefDef(c *RefDefContext)

	// ExitEventType is called when exiting the eventType production.
	ExitEventType(c *EventTypeContext)

	// ExitIfCondition is called when exiting the ifCondition production.
	ExitIfCondition(c *IfConditionContext)

	// ExitMulDivExpr is called when exiting the MulDivExpr production.
	ExitMulDivExpr(c *MulDivExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBooleanLiteral is called when exiting the BooleanLiteral production.
	ExitBooleanLiteral(c *BooleanLiteralContext)

	// ExitIntLiteral is called when exiting the IntLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitRelationalExpr is called when exiting the RelationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitAddSubExpr is called when exiting the AddSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitLogicalExpr is called when exiting the LogicalExpr production.
	ExitLogicalExpr(c *LogicalExprContext)

	// ExitExprList is called when exiting the exprList production.
	ExitExprList(c *ExprListContext)

	// ExitMainFunction is called when exiting the mainFunction production.
	ExitMainFunction(c *MainFunctionContext)

	// ExitMainStatements is called when exiting the mainStatements production.
	ExitMainStatements(c *MainStatementsContext)

	// ExitMainStatement is called when exiting the mainStatement production.
	ExitMainStatement(c *MainStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitServerStatement is called when exiting the serverStatement production.
	ExitServerStatement(c *ServerStatementContext)
}
