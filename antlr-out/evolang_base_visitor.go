// Code generated from ./EvoLang.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLang

import "github.com/antlr4-go/antlr/v4"

type BaseEvoLangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseEvoLangVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitModelDef(ctx *ModelDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitModelBody(ctx *ModelBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitFieldDef(ctx *FieldDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitEventDef(ctx *EventDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitEventBody(ctx *EventBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitEnumDef(ctx *EnumDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitAccessControlDef(ctx *AccessControlDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitAccessRule(ctx *AccessRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitAction(ctx *ActionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitEndpointDef(ctx *EndpointDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitEndpointBody(ctx *EndpointBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitServerDef(ctx *ServerDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitServerBody(ctx *ServerBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitBodyDef(ctx *BodyDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitAuthenticateDef(ctx *AuthenticateDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitReturnDef(ctx *ReturnDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitSandbox(ctx *SandboxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitGoCode(ctx *GoCodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitTypeDef(ctx *TypeDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitArrayDef(ctx *ArrayDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitRefDef(ctx *RefDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitEventType(ctx *EventTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitIfCondition(ctx *IfConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitMulDivExpr(ctx *MulDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitRelationalExpr(ctx *RelationalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitLogicalExpr(ctx *LogicalExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitExprList(ctx *ExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitMainFunction(ctx *MainFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitMainStatements(ctx *MainStatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitMainStatement(ctx *MainStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseEvoLangVisitor) VisitServerStatement(ctx *ServerStatementContext) interface{} {
	return v.VisitChildren(ctx)
}
