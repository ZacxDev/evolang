// Code generated from ./EvoLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLangParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EvoLangParser.
type EvoLangParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EvoLangParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by EvoLangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by EvoLangParser#modelDef.
	VisitModelDef(ctx *ModelDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#modelBody.
	VisitModelBody(ctx *ModelBodyContext) interface{}

	// Visit a parse tree produced by EvoLangParser#modelFieldDef.
	VisitModelFieldDef(ctx *ModelFieldDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#ruleDef.
	VisitRuleDef(ctx *RuleDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#ruleBody.
	VisitRuleBody(ctx *RuleBodyContext) interface{}

	// Visit a parse tree produced by EvoLangParser#ruleStatement.
	VisitRuleStatement(ctx *RuleStatementContext) interface{}

	// Visit a parse tree produced by EvoLangParser#ruleParameters.
	VisitRuleParameters(ctx *RuleParametersContext) interface{}

	// Visit a parse tree produced by EvoLangParser#functionBody.
	VisitFunctionBody(ctx *FunctionBodyContext) interface{}

	// Visit a parse tree produced by EvoLangParser#assignmentStatement.
	VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{}

	// Visit a parse tree produced by EvoLangParser#conditionalStatement.
	VisitConditionalStatement(ctx *ConditionalStatementContext) interface{}

	// Visit a parse tree produced by EvoLangParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by EvoLangParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by EvoLangParser#ruleConditionExpression.
	VisitRuleConditionExpression(ctx *RuleConditionExpressionContext) interface{}

	// Visit a parse tree produced by EvoLangParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) interface{}

	// Visit a parse tree produced by EvoLangParser#inExpression.
	VisitInExpression(ctx *InExpressionContext) interface{}

	// Visit a parse tree produced by EvoLangParser#ruleExpressionPrimary.
	VisitRuleExpressionPrimary(ctx *RuleExpressionPrimaryContext) interface{}

	// Visit a parse tree produced by EvoLangParser#mutationDef.
	VisitMutationDef(ctx *MutationDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#mutationBody.
	VisitMutationBody(ctx *MutationBodyContext) interface{}

	// Visit a parse tree produced by EvoLangParser#mutationStatement.
	VisitMutationStatement(ctx *MutationStatementContext) interface{}

	// Visit a parse tree produced by EvoLangParser#mutationParameters.
	VisitMutationParameters(ctx *MutationParametersContext) interface{}

	// Visit a parse tree produced by EvoLangParser#mutationInnerBlock.
	VisitMutationInnerBlock(ctx *MutationInnerBlockContext) interface{}

	// Visit a parse tree produced by EvoLangParser#middlewareBlock.
	VisitMiddlewareBlock(ctx *MiddlewareBlockContext) interface{}

	// Visit a parse tree produced by EvoLangParser#functionBlock.
	VisitFunctionBlock(ctx *FunctionBlockContext) interface{}

	// Visit a parse tree produced by EvoLangParser#inputDef.
	VisitInputDef(ctx *InputDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#inputBody.
	VisitInputBody(ctx *InputBodyContext) interface{}

	// Visit a parse tree produced by EvoLangParser#inputFieldDef.
	VisitInputFieldDef(ctx *InputFieldDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#eventDef.
	VisitEventDef(ctx *EventDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#eventChannelHandler.
	VisitEventChannelHandler(ctx *EventChannelHandlerContext) interface{}

	// Visit a parse tree produced by EvoLangParser#enumDef.
	VisitEnumDef(ctx *EnumDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#references.
	VisitReferences(ctx *ReferencesContext) interface{}

	// Visit a parse tree produced by EvoLangParser#linkedBy.
	VisitLinkedBy(ctx *LinkedByContext) interface{}

	// Visit a parse tree produced by EvoLangParser#action.
	VisitAction(ctx *ActionContext) interface{}

	// Visit a parse tree produced by EvoLangParser#paramList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by EvoLangParser#bodyDef.
	VisitBodyDef(ctx *BodyDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#authenticateDef.
	VisitAuthenticateDef(ctx *AuthenticateDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#returnDef.
	VisitReturnDef(ctx *ReturnDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#typeDef.
	VisitTypeDef(ctx *TypeDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#arrayDef.
	VisitArrayDef(ctx *ArrayDefContext) interface{}

	// Visit a parse tree produced by EvoLangParser#eventType.
	VisitEventType(ctx *EventTypeContext) interface{}

	// Visit a parse tree produced by EvoLangParser#FunctionCallExpr.
	VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{}

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

	// Visit a parse tree produced by EvoLangParser#RelationalExpr.
	VisitRelationalExpr(ctx *RelationalExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#PropertyAccessExpr.
	VisitPropertyAccessExpr(ctx *PropertyAccessExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#LogicalExpr.
	VisitLogicalExpr(ctx *LogicalExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#MethodCallExpr.
	VisitMethodCallExpr(ctx *MethodCallExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#methodCallExpression.
	VisitMethodCallExpression(ctx *MethodCallExpressionContext) interface{}

	// Visit a parse tree produced by EvoLangParser#SimpleIdExpr.
	VisitSimpleIdExpr(ctx *SimpleIdExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#PrimaryFunctionCallExpr.
	VisitPrimaryFunctionCallExpr(ctx *PrimaryFunctionCallExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#PrimaryPropertyAccessExpr.
	VisitPrimaryPropertyAccessExpr(ctx *PrimaryPropertyAccessExprContext) interface{}

	// Visit a parse tree produced by EvoLangParser#propertyAccess.
	VisitPropertyAccess(ctx *PropertyAccessContext) interface{}

	// Visit a parse tree produced by EvoLangParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

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

	// Visit a parse tree produced by EvoLangParser#sandboxBlock.
	VisitSandboxBlock(ctx *SandboxBlockContext) interface{}

	// Visit a parse tree produced by EvoLangParser#sandboxAttributes.
	VisitSandboxAttributes(ctx *SandboxAttributesContext) interface{}

	// Visit a parse tree produced by EvoLangParser#sandboxAttribute.
	VisitSandboxAttribute(ctx *SandboxAttributeContext) interface{}

	// Visit a parse tree produced by EvoLangParser#sandboxLangAttribute.
	VisitSandboxLangAttribute(ctx *SandboxLangAttributeContext) interface{}

	// Visit a parse tree produced by EvoLangParser#sandboxDefineAttribute.
	VisitSandboxDefineAttribute(ctx *SandboxDefineAttributeContext) interface{}

	// Visit a parse tree produced by EvoLangParser#sandboxDefineContent.
	VisitSandboxDefineContent(ctx *SandboxDefineContentContext) interface{}
}
