// Code generated from ./EvoLang.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLang

import "github.com/antlr4-go/antlr/v4"

// BaseEvoLangListener is a complete listener for a parse tree produced by EvoLangParser.
type BaseEvoLangListener struct{}

var _ EvoLangListener = &BaseEvoLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvoLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvoLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvoLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvoLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseEvoLangListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseEvoLangListener) ExitProg(ctx *ProgContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseEvoLangListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseEvoLangListener) ExitStatement(ctx *StatementContext) {}

// EnterModelDef is called when production modelDef is entered.
func (s *BaseEvoLangListener) EnterModelDef(ctx *ModelDefContext) {}

// ExitModelDef is called when production modelDef is exited.
func (s *BaseEvoLangListener) ExitModelDef(ctx *ModelDefContext) {}

// EnterModelBody is called when production modelBody is entered.
func (s *BaseEvoLangListener) EnterModelBody(ctx *ModelBodyContext) {}

// ExitModelBody is called when production modelBody is exited.
func (s *BaseEvoLangListener) ExitModelBody(ctx *ModelBodyContext) {}

// EnterFieldDef is called when production fieldDef is entered.
func (s *BaseEvoLangListener) EnterFieldDef(ctx *FieldDefContext) {}

// ExitFieldDef is called when production fieldDef is exited.
func (s *BaseEvoLangListener) ExitFieldDef(ctx *FieldDefContext) {}

// EnterEventDef is called when production eventDef is entered.
func (s *BaseEvoLangListener) EnterEventDef(ctx *EventDefContext) {}

// ExitEventDef is called when production eventDef is exited.
func (s *BaseEvoLangListener) ExitEventDef(ctx *EventDefContext) {}

// EnterEventBody is called when production eventBody is entered.
func (s *BaseEvoLangListener) EnterEventBody(ctx *EventBodyContext) {}

// ExitEventBody is called when production eventBody is exited.
func (s *BaseEvoLangListener) ExitEventBody(ctx *EventBodyContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BaseEvoLangListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BaseEvoLangListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterAccessControlDef is called when production accessControlDef is entered.
func (s *BaseEvoLangListener) EnterAccessControlDef(ctx *AccessControlDefContext) {}

// ExitAccessControlDef is called when production accessControlDef is exited.
func (s *BaseEvoLangListener) ExitAccessControlDef(ctx *AccessControlDefContext) {}

// EnterAccessRule is called when production accessRule is entered.
func (s *BaseEvoLangListener) EnterAccessRule(ctx *AccessRuleContext) {}

// ExitAccessRule is called when production accessRule is exited.
func (s *BaseEvoLangListener) ExitAccessRule(ctx *AccessRuleContext) {}

// EnterAction is called when production action is entered.
func (s *BaseEvoLangListener) EnterAction(ctx *ActionContext) {}

// ExitAction is called when production action is exited.
func (s *BaseEvoLangListener) ExitAction(ctx *ActionContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseEvoLangListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseEvoLangListener) ExitParamList(ctx *ParamListContext) {}

// EnterEndpointDef is called when production endpointDef is entered.
func (s *BaseEvoLangListener) EnterEndpointDef(ctx *EndpointDefContext) {}

// ExitEndpointDef is called when production endpointDef is exited.
func (s *BaseEvoLangListener) ExitEndpointDef(ctx *EndpointDefContext) {}

// EnterEndpointBody is called when production endpointBody is entered.
func (s *BaseEvoLangListener) EnterEndpointBody(ctx *EndpointBodyContext) {}

// ExitEndpointBody is called when production endpointBody is exited.
func (s *BaseEvoLangListener) ExitEndpointBody(ctx *EndpointBodyContext) {}

// EnterServerDef is called when production serverDef is entered.
func (s *BaseEvoLangListener) EnterServerDef(ctx *ServerDefContext) {}

// ExitServerDef is called when production serverDef is exited.
func (s *BaseEvoLangListener) ExitServerDef(ctx *ServerDefContext) {}

// EnterServerBody is called when production serverBody is entered.
func (s *BaseEvoLangListener) EnterServerBody(ctx *ServerBodyContext) {}

// ExitServerBody is called when production serverBody is exited.
func (s *BaseEvoLangListener) ExitServerBody(ctx *ServerBodyContext) {}

// EnterBodyDef is called when production bodyDef is entered.
func (s *BaseEvoLangListener) EnterBodyDef(ctx *BodyDefContext) {}

// ExitBodyDef is called when production bodyDef is exited.
func (s *BaseEvoLangListener) ExitBodyDef(ctx *BodyDefContext) {}

// EnterAuthenticateDef is called when production authenticateDef is entered.
func (s *BaseEvoLangListener) EnterAuthenticateDef(ctx *AuthenticateDefContext) {}

// ExitAuthenticateDef is called when production authenticateDef is exited.
func (s *BaseEvoLangListener) ExitAuthenticateDef(ctx *AuthenticateDefContext) {}

// EnterReturnDef is called when production returnDef is entered.
func (s *BaseEvoLangListener) EnterReturnDef(ctx *ReturnDefContext) {}

// ExitReturnDef is called when production returnDef is exited.
func (s *BaseEvoLangListener) ExitReturnDef(ctx *ReturnDefContext) {}

// EnterSandbox is called when production sandbox is entered.
func (s *BaseEvoLangListener) EnterSandbox(ctx *SandboxContext) {}

// ExitSandbox is called when production sandbox is exited.
func (s *BaseEvoLangListener) ExitSandbox(ctx *SandboxContext) {}

// EnterGoCode is called when production goCode is entered.
func (s *BaseEvoLangListener) EnterGoCode(ctx *GoCodeContext) {}

// ExitGoCode is called when production goCode is exited.
func (s *BaseEvoLangListener) ExitGoCode(ctx *GoCodeContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BaseEvoLangListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BaseEvoLangListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterArrayDef is called when production arrayDef is entered.
func (s *BaseEvoLangListener) EnterArrayDef(ctx *ArrayDefContext) {}

// ExitArrayDef is called when production arrayDef is exited.
func (s *BaseEvoLangListener) ExitArrayDef(ctx *ArrayDefContext) {}

// EnterRefDef is called when production refDef is entered.
func (s *BaseEvoLangListener) EnterRefDef(ctx *RefDefContext) {}

// ExitRefDef is called when production refDef is exited.
func (s *BaseEvoLangListener) ExitRefDef(ctx *RefDefContext) {}

// EnterEventType is called when production eventType is entered.
func (s *BaseEvoLangListener) EnterEventType(ctx *EventTypeContext) {}

// ExitEventType is called when production eventType is exited.
func (s *BaseEvoLangListener) ExitEventType(ctx *EventTypeContext) {}

// EnterIfCondition is called when production ifCondition is entered.
func (s *BaseEvoLangListener) EnterIfCondition(ctx *IfConditionContext) {}

// ExitIfCondition is called when production ifCondition is exited.
func (s *BaseEvoLangListener) ExitIfCondition(ctx *IfConditionContext) {}

// EnterMulDivExpr is called when production MulDivExpr is entered.
func (s *BaseEvoLangListener) EnterMulDivExpr(ctx *MulDivExprContext) {}

// ExitMulDivExpr is called when production MulDivExpr is exited.
func (s *BaseEvoLangListener) ExitMulDivExpr(ctx *MulDivExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseEvoLangListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseEvoLangListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseEvoLangListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseEvoLangListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production BooleanLiteral is entered.
func (s *BaseEvoLangListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production BooleanLiteral is exited.
func (s *BaseEvoLangListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterIntLiteral is called when production IntLiteral is entered.
func (s *BaseEvoLangListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production IntLiteral is exited.
func (s *BaseEvoLangListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseEvoLangListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseEvoLangListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterRelationalExpr is called when production RelationalExpr is entered.
func (s *BaseEvoLangListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production RelationalExpr is exited.
func (s *BaseEvoLangListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseEvoLangListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseEvoLangListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterAddSubExpr is called when production AddSubExpr is entered.
func (s *BaseEvoLangListener) EnterAddSubExpr(ctx *AddSubExprContext) {}

// ExitAddSubExpr is called when production AddSubExpr is exited.
func (s *BaseEvoLangListener) ExitAddSubExpr(ctx *AddSubExprContext) {}

// EnterLogicalExpr is called when production LogicalExpr is entered.
func (s *BaseEvoLangListener) EnterLogicalExpr(ctx *LogicalExprContext) {}

// ExitLogicalExpr is called when production LogicalExpr is exited.
func (s *BaseEvoLangListener) ExitLogicalExpr(ctx *LogicalExprContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseEvoLangListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseEvoLangListener) ExitExprList(ctx *ExprListContext) {}

// EnterMainFunction is called when production mainFunction is entered.
func (s *BaseEvoLangListener) EnterMainFunction(ctx *MainFunctionContext) {}

// ExitMainFunction is called when production mainFunction is exited.
func (s *BaseEvoLangListener) ExitMainFunction(ctx *MainFunctionContext) {}

// EnterMainStatements is called when production mainStatements is entered.
func (s *BaseEvoLangListener) EnterMainStatements(ctx *MainStatementsContext) {}

// ExitMainStatements is called when production mainStatements is exited.
func (s *BaseEvoLangListener) ExitMainStatements(ctx *MainStatementsContext) {}

// EnterMainStatement is called when production mainStatement is entered.
func (s *BaseEvoLangListener) EnterMainStatement(ctx *MainStatementContext) {}

// ExitMainStatement is called when production mainStatement is exited.
func (s *BaseEvoLangListener) ExitMainStatement(ctx *MainStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseEvoLangListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseEvoLangListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterServerStatement is called when production serverStatement is entered.
func (s *BaseEvoLangListener) EnterServerStatement(ctx *ServerStatementContext) {}

// ExitServerStatement is called when production serverStatement is exited.
func (s *BaseEvoLangListener) ExitServerStatement(ctx *ServerStatementContext) {}
