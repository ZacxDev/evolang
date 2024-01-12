// Code generated from ./EvoLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLangParser

import "github.com/antlr4-go/antlr/v4"

// BaseEvoLangParserListener is a complete listener for a parse tree produced by EvoLangParser.
type BaseEvoLangParserListener struct{}

var _ EvoLangParserListener = &BaseEvoLangParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvoLangParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvoLangParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvoLangParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvoLangParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseEvoLangParserListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseEvoLangParserListener) ExitProg(ctx *ProgContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseEvoLangParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseEvoLangParserListener) ExitStatement(ctx *StatementContext) {}

// EnterModelDef is called when production modelDef is entered.
func (s *BaseEvoLangParserListener) EnterModelDef(ctx *ModelDefContext) {}

// ExitModelDef is called when production modelDef is exited.
func (s *BaseEvoLangParserListener) ExitModelDef(ctx *ModelDefContext) {}

// EnterModelBody is called when production modelBody is entered.
func (s *BaseEvoLangParserListener) EnterModelBody(ctx *ModelBodyContext) {}

// ExitModelBody is called when production modelBody is exited.
func (s *BaseEvoLangParserListener) ExitModelBody(ctx *ModelBodyContext) {}

// EnterModelFieldDef is called when production modelFieldDef is entered.
func (s *BaseEvoLangParserListener) EnterModelFieldDef(ctx *ModelFieldDefContext) {}

// ExitModelFieldDef is called when production modelFieldDef is exited.
func (s *BaseEvoLangParserListener) ExitModelFieldDef(ctx *ModelFieldDefContext) {}

// EnterRuleDef is called when production ruleDef is entered.
func (s *BaseEvoLangParserListener) EnterRuleDef(ctx *RuleDefContext) {}

// ExitRuleDef is called when production ruleDef is exited.
func (s *BaseEvoLangParserListener) ExitRuleDef(ctx *RuleDefContext) {}

// EnterRuleBody is called when production ruleBody is entered.
func (s *BaseEvoLangParserListener) EnterRuleBody(ctx *RuleBodyContext) {}

// ExitRuleBody is called when production ruleBody is exited.
func (s *BaseEvoLangParserListener) ExitRuleBody(ctx *RuleBodyContext) {}

// EnterRuleStatement is called when production ruleStatement is entered.
func (s *BaseEvoLangParserListener) EnterRuleStatement(ctx *RuleStatementContext) {}

// ExitRuleStatement is called when production ruleStatement is exited.
func (s *BaseEvoLangParserListener) ExitRuleStatement(ctx *RuleStatementContext) {}

// EnterRuleParameters is called when production ruleParameters is entered.
func (s *BaseEvoLangParserListener) EnterRuleParameters(ctx *RuleParametersContext) {}

// ExitRuleParameters is called when production ruleParameters is exited.
func (s *BaseEvoLangParserListener) ExitRuleParameters(ctx *RuleParametersContext) {}

// EnterFunctionBody is called when production functionBody is entered.
func (s *BaseEvoLangParserListener) EnterFunctionBody(ctx *FunctionBodyContext) {}

// ExitFunctionBody is called when production functionBody is exited.
func (s *BaseEvoLangParserListener) ExitFunctionBody(ctx *FunctionBodyContext) {}

// EnterConditionalStatement is called when production conditionalStatement is entered.
func (s *BaseEvoLangParserListener) EnterConditionalStatement(ctx *ConditionalStatementContext) {}

// ExitConditionalStatement is called when production conditionalStatement is exited.
func (s *BaseEvoLangParserListener) ExitConditionalStatement(ctx *ConditionalStatementContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseEvoLangParserListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseEvoLangParserListener) ExitCondition(ctx *ConditionContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseEvoLangParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseEvoLangParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterRuleConditionExpression is called when production ruleConditionExpression is entered.
func (s *BaseEvoLangParserListener) EnterRuleConditionExpression(ctx *RuleConditionExpressionContext) {
}

// ExitRuleConditionExpression is called when production ruleConditionExpression is exited.
func (s *BaseEvoLangParserListener) ExitRuleConditionExpression(ctx *RuleConditionExpressionContext) {
}

// EnterEqualityExpression is called when production equalityExpression is entered.
func (s *BaseEvoLangParserListener) EnterEqualityExpression(ctx *EqualityExpressionContext) {}

// ExitEqualityExpression is called when production equalityExpression is exited.
func (s *BaseEvoLangParserListener) ExitEqualityExpression(ctx *EqualityExpressionContext) {}

// EnterInExpression is called when production inExpression is entered.
func (s *BaseEvoLangParserListener) EnterInExpression(ctx *InExpressionContext) {}

// ExitInExpression is called when production inExpression is exited.
func (s *BaseEvoLangParserListener) ExitInExpression(ctx *InExpressionContext) {}

// EnterRuleExpressionPrimary is called when production ruleExpressionPrimary is entered.
func (s *BaseEvoLangParserListener) EnterRuleExpressionPrimary(ctx *RuleExpressionPrimaryContext) {}

// ExitRuleExpressionPrimary is called when production ruleExpressionPrimary is exited.
func (s *BaseEvoLangParserListener) ExitRuleExpressionPrimary(ctx *RuleExpressionPrimaryContext) {}

// EnterMutationDef is called when production mutationDef is entered.
func (s *BaseEvoLangParserListener) EnterMutationDef(ctx *MutationDefContext) {}

// ExitMutationDef is called when production mutationDef is exited.
func (s *BaseEvoLangParserListener) ExitMutationDef(ctx *MutationDefContext) {}

// EnterMutationBody is called when production mutationBody is entered.
func (s *BaseEvoLangParserListener) EnterMutationBody(ctx *MutationBodyContext) {}

// ExitMutationBody is called when production mutationBody is exited.
func (s *BaseEvoLangParserListener) ExitMutationBody(ctx *MutationBodyContext) {}

// EnterMutationStatement is called when production mutationStatement is entered.
func (s *BaseEvoLangParserListener) EnterMutationStatement(ctx *MutationStatementContext) {}

// ExitMutationStatement is called when production mutationStatement is exited.
func (s *BaseEvoLangParserListener) ExitMutationStatement(ctx *MutationStatementContext) {}

// EnterMutationDetails is called when production mutationDetails is entered.
func (s *BaseEvoLangParserListener) EnterMutationDetails(ctx *MutationDetailsContext) {}

// ExitMutationDetails is called when production mutationDetails is exited.
func (s *BaseEvoLangParserListener) ExitMutationDetails(ctx *MutationDetailsContext) {}

// EnterMutationRulesApplication is called when production mutationRulesApplication is entered.
func (s *BaseEvoLangParserListener) EnterMutationRulesApplication(ctx *MutationRulesApplicationContext) {
}

// ExitMutationRulesApplication is called when production mutationRulesApplication is exited.
func (s *BaseEvoLangParserListener) ExitMutationRulesApplication(ctx *MutationRulesApplicationContext) {
}

// EnterOrExpression is called when production orExpression is entered.
func (s *BaseEvoLangParserListener) EnterOrExpression(ctx *OrExpressionContext) {}

// ExitOrExpression is called when production orExpression is exited.
func (s *BaseEvoLangParserListener) ExitOrExpression(ctx *OrExpressionContext) {}

// EnterAndExpression is called when production andExpression is entered.
func (s *BaseEvoLangParserListener) EnterAndExpression(ctx *AndExpressionContext) {}

// ExitAndExpression is called when production andExpression is exited.
func (s *BaseEvoLangParserListener) ExitAndExpression(ctx *AndExpressionContext) {}

// EnterMutationRulesApplicationExpressionPrimary is called when production mutationRulesApplicationExpressionPrimary is entered.
func (s *BaseEvoLangParserListener) EnterMutationRulesApplicationExpressionPrimary(ctx *MutationRulesApplicationExpressionPrimaryContext) {
}

// ExitMutationRulesApplicationExpressionPrimary is called when production mutationRulesApplicationExpressionPrimary is exited.
func (s *BaseEvoLangParserListener) ExitMutationRulesApplicationExpressionPrimary(ctx *MutationRulesApplicationExpressionPrimaryContext) {
}

// EnterInputDef is called when production inputDef is entered.
func (s *BaseEvoLangParserListener) EnterInputDef(ctx *InputDefContext) {}

// ExitInputDef is called when production inputDef is exited.
func (s *BaseEvoLangParserListener) ExitInputDef(ctx *InputDefContext) {}

// EnterInputBody is called when production inputBody is entered.
func (s *BaseEvoLangParserListener) EnterInputBody(ctx *InputBodyContext) {}

// ExitInputBody is called when production inputBody is exited.
func (s *BaseEvoLangParserListener) ExitInputBody(ctx *InputBodyContext) {}

// EnterInputFieldDef is called when production inputFieldDef is entered.
func (s *BaseEvoLangParserListener) EnterInputFieldDef(ctx *InputFieldDefContext) {}

// ExitInputFieldDef is called when production inputFieldDef is exited.
func (s *BaseEvoLangParserListener) ExitInputFieldDef(ctx *InputFieldDefContext) {}

// EnterEventDef is called when production eventDef is entered.
func (s *BaseEvoLangParserListener) EnterEventDef(ctx *EventDefContext) {}

// ExitEventDef is called when production eventDef is exited.
func (s *BaseEvoLangParserListener) ExitEventDef(ctx *EventDefContext) {}

// EnterEventChannelHandler is called when production eventChannelHandler is entered.
func (s *BaseEvoLangParserListener) EnterEventChannelHandler(ctx *EventChannelHandlerContext) {}

// ExitEventChannelHandler is called when production eventChannelHandler is exited.
func (s *BaseEvoLangParserListener) ExitEventChannelHandler(ctx *EventChannelHandlerContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BaseEvoLangParserListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BaseEvoLangParserListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterReferences is called when production references is entered.
func (s *BaseEvoLangParserListener) EnterReferences(ctx *ReferencesContext) {}

// ExitReferences is called when production references is exited.
func (s *BaseEvoLangParserListener) ExitReferences(ctx *ReferencesContext) {}

// EnterLinkedBy is called when production linkedBy is entered.
func (s *BaseEvoLangParserListener) EnterLinkedBy(ctx *LinkedByContext) {}

// ExitLinkedBy is called when production linkedBy is exited.
func (s *BaseEvoLangParserListener) ExitLinkedBy(ctx *LinkedByContext) {}

// EnterAction is called when production action is entered.
func (s *BaseEvoLangParserListener) EnterAction(ctx *ActionContext) {}

// ExitAction is called when production action is exited.
func (s *BaseEvoLangParserListener) ExitAction(ctx *ActionContext) {}

// EnterParamList is called when production paramList is entered.
func (s *BaseEvoLangParserListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production paramList is exited.
func (s *BaseEvoLangParserListener) ExitParamList(ctx *ParamListContext) {}

// EnterBodyDef is called when production bodyDef is entered.
func (s *BaseEvoLangParserListener) EnterBodyDef(ctx *BodyDefContext) {}

// ExitBodyDef is called when production bodyDef is exited.
func (s *BaseEvoLangParserListener) ExitBodyDef(ctx *BodyDefContext) {}

// EnterAuthenticateDef is called when production authenticateDef is entered.
func (s *BaseEvoLangParserListener) EnterAuthenticateDef(ctx *AuthenticateDefContext) {}

// ExitAuthenticateDef is called when production authenticateDef is exited.
func (s *BaseEvoLangParserListener) ExitAuthenticateDef(ctx *AuthenticateDefContext) {}

// EnterReturnDef is called when production returnDef is entered.
func (s *BaseEvoLangParserListener) EnterReturnDef(ctx *ReturnDefContext) {}

// ExitReturnDef is called when production returnDef is exited.
func (s *BaseEvoLangParserListener) ExitReturnDef(ctx *ReturnDefContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BaseEvoLangParserListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BaseEvoLangParserListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterArrayDef is called when production arrayDef is entered.
func (s *BaseEvoLangParserListener) EnterArrayDef(ctx *ArrayDefContext) {}

// ExitArrayDef is called when production arrayDef is exited.
func (s *BaseEvoLangParserListener) ExitArrayDef(ctx *ArrayDefContext) {}

// EnterEventType is called when production eventType is entered.
func (s *BaseEvoLangParserListener) EnterEventType(ctx *EventTypeContext) {}

// ExitEventType is called when production eventType is exited.
func (s *BaseEvoLangParserListener) ExitEventType(ctx *EventTypeContext) {}

// EnterMulDivExpr is called when production MulDivExpr is entered.
func (s *BaseEvoLangParserListener) EnterMulDivExpr(ctx *MulDivExprContext) {}

// ExitMulDivExpr is called when production MulDivExpr is exited.
func (s *BaseEvoLangParserListener) ExitMulDivExpr(ctx *MulDivExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseEvoLangParserListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseEvoLangParserListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseEvoLangParserListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseEvoLangParserListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBooleanLiteral is called when production BooleanLiteral is entered.
func (s *BaseEvoLangParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production BooleanLiteral is exited.
func (s *BaseEvoLangParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterIntLiteral is called when production IntLiteral is entered.
func (s *BaseEvoLangParserListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production IntLiteral is exited.
func (s *BaseEvoLangParserListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseEvoLangParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseEvoLangParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterRelationalExpr is called when production RelationalExpr is entered.
func (s *BaseEvoLangParserListener) EnterRelationalExpr(ctx *RelationalExprContext) {}

// ExitRelationalExpr is called when production RelationalExpr is exited.
func (s *BaseEvoLangParserListener) ExitRelationalExpr(ctx *RelationalExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseEvoLangParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseEvoLangParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterAddSubExpr is called when production AddSubExpr is entered.
func (s *BaseEvoLangParserListener) EnterAddSubExpr(ctx *AddSubExprContext) {}

// ExitAddSubExpr is called when production AddSubExpr is exited.
func (s *BaseEvoLangParserListener) ExitAddSubExpr(ctx *AddSubExprContext) {}

// EnterLogicalExpr is called when production LogicalExpr is entered.
func (s *BaseEvoLangParserListener) EnterLogicalExpr(ctx *LogicalExprContext) {}

// ExitLogicalExpr is called when production LogicalExpr is exited.
func (s *BaseEvoLangParserListener) ExitLogicalExpr(ctx *LogicalExprContext) {}

// EnterExprList is called when production exprList is entered.
func (s *BaseEvoLangParserListener) EnterExprList(ctx *ExprListContext) {}

// ExitExprList is called when production exprList is exited.
func (s *BaseEvoLangParserListener) ExitExprList(ctx *ExprListContext) {}

// EnterMainFunction is called when production mainFunction is entered.
func (s *BaseEvoLangParserListener) EnterMainFunction(ctx *MainFunctionContext) {}

// ExitMainFunction is called when production mainFunction is exited.
func (s *BaseEvoLangParserListener) ExitMainFunction(ctx *MainFunctionContext) {}

// EnterMainStatements is called when production mainStatements is entered.
func (s *BaseEvoLangParserListener) EnterMainStatements(ctx *MainStatementsContext) {}

// ExitMainStatements is called when production mainStatements is exited.
func (s *BaseEvoLangParserListener) ExitMainStatements(ctx *MainStatementsContext) {}

// EnterMainStatement is called when production mainStatement is entered.
func (s *BaseEvoLangParserListener) EnterMainStatement(ctx *MainStatementContext) {}

// ExitMainStatement is called when production mainStatement is exited.
func (s *BaseEvoLangParserListener) ExitMainStatement(ctx *MainStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseEvoLangParserListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseEvoLangParserListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterServerStatement is called when production serverStatement is entered.
func (s *BaseEvoLangParserListener) EnterServerStatement(ctx *ServerStatementContext) {}

// ExitServerStatement is called when production serverStatement is exited.
func (s *BaseEvoLangParserListener) ExitServerStatement(ctx *ServerStatementContext) {}
