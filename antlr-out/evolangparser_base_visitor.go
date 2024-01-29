// Code generated from ./EvoLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLangParser

import "github.com/antlr4-go/antlr/v4"

type BaseEvoLangParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseEvoLangParserVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitModelDef(ctx *ModelDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitModelBody(ctx *ModelBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitModelFieldDef(ctx *ModelFieldDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitRuleDef(ctx *RuleDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitRuleBody(ctx *RuleBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitRuleStatement(ctx *RuleStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitRuleParameters(ctx *RuleParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitFunctionBody(ctx *FunctionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitAssignmentStatement(ctx *AssignmentStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitConditionalStatement(ctx *ConditionalStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitRuleConditionExpression(ctx *RuleConditionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitEqualityExpression(ctx *EqualityExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitInExpression(ctx *InExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitRuleExpressionPrimary(ctx *RuleExpressionPrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMutationDef(ctx *MutationDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMutationBody(ctx *MutationBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMutationStatement(ctx *MutationStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMutationParameters(ctx *MutationParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMutationInnerBlock(ctx *MutationInnerBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMiddlewareBlock(ctx *MiddlewareBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitFunctionBlock(ctx *FunctionBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitInputDef(ctx *InputDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitInputBody(ctx *InputBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitInputFieldDef(ctx *InputFieldDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitEventDef(ctx *EventDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitEventChannelHandler(ctx *EventChannelHandlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitEnumDef(ctx *EnumDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitReferences(ctx *ReferencesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitLinkedBy(ctx *LinkedByContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitAction(ctx *ActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitBodyDef(ctx *BodyDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitAuthenticateDef(ctx *AuthenticateDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitReturnDef(ctx *ReturnDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitTypeDef(ctx *TypeDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitArrayDef(ctx *ArrayDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitEventType(ctx *EventTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMulDivExpr(ctx *MulDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitPropertyAccessExpr(ctx *PropertyAccessExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitLogicalExpr(ctx *LogicalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMethodCallExpr(ctx *MethodCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMethodCallExpression(ctx *MethodCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitSimpleIdExpr(ctx *SimpleIdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitPrimaryFunctionCallExpr(ctx *PrimaryFunctionCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitPrimaryPropertyAccessExpr(ctx *PrimaryPropertyAccessExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitPropertyAccess(ctx *PropertyAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMainFunction(ctx *MainFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMainStatements(ctx *MainStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitMainStatement(ctx *MainStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitServerStatement(ctx *ServerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitSandboxBlock(ctx *SandboxBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitSandboxAttributes(ctx *SandboxAttributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitSandboxAttribute(ctx *SandboxAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitSandboxLangAttribute(ctx *SandboxLangAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitSandboxDefineAttribute(ctx *SandboxDefineAttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangParserVisitor) VisitSandboxDefineContent(ctx *SandboxDefineContentContext) interface{} {
	return v.VisitChildren(ctx)
}
