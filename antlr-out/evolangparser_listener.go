// Code generated from ./EvoLangParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // EvoLangParser

import "github.com/antlr4-go/antlr/v4"

// EvoLangParserListener is a complete listener for a parse tree produced by EvoLangParser.
type EvoLangParserListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterModelDef is called when entering the modelDef production.
	EnterModelDef(c *ModelDefContext)

	// EnterModelBody is called when entering the modelBody production.
	EnterModelBody(c *ModelBodyContext)

	// EnterModelFieldDef is called when entering the modelFieldDef production.
	EnterModelFieldDef(c *ModelFieldDefContext)

	// EnterRuleDef is called when entering the ruleDef production.
	EnterRuleDef(c *RuleDefContext)

	// EnterRuleBody is called when entering the ruleBody production.
	EnterRuleBody(c *RuleBodyContext)

	// EnterRuleStatement is called when entering the ruleStatement production.
	EnterRuleStatement(c *RuleStatementContext)

	// EnterRuleParameters is called when entering the ruleParameters production.
	EnterRuleParameters(c *RuleParametersContext)

	// EnterFunctionBody is called when entering the functionBody production.
	EnterFunctionBody(c *FunctionBodyContext)

	// EnterAssignmentStatement is called when entering the assignmentStatement production.
	EnterAssignmentStatement(c *AssignmentStatementContext)

	// EnterConditionalStatement is called when entering the conditionalStatement production.
	EnterConditionalStatement(c *ConditionalStatementContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterRuleConditionExpression is called when entering the ruleConditionExpression production.
	EnterRuleConditionExpression(c *RuleConditionExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterInExpression is called when entering the inExpression production.
	EnterInExpression(c *InExpressionContext)

	// EnterRuleExpressionPrimary is called when entering the ruleExpressionPrimary production.
	EnterRuleExpressionPrimary(c *RuleExpressionPrimaryContext)

	// EnterMutationDef is called when entering the mutationDef production.
	EnterMutationDef(c *MutationDefContext)

	// EnterMutationBody is called when entering the mutationBody production.
	EnterMutationBody(c *MutationBodyContext)

	// EnterMutationStatement is called when entering the mutationStatement production.
	EnterMutationStatement(c *MutationStatementContext)

	// EnterMutationParameters is called when entering the mutationParameters production.
	EnterMutationParameters(c *MutationParametersContext)

	// EnterMutationInnerBlock is called when entering the mutationInnerBlock production.
	EnterMutationInnerBlock(c *MutationInnerBlockContext)

	// EnterMiddlewareBlock is called when entering the middlewareBlock production.
	EnterMiddlewareBlock(c *MiddlewareBlockContext)

	// EnterFunctionBlock is called when entering the functionBlock production.
	EnterFunctionBlock(c *FunctionBlockContext)

	// EnterInputDef is called when entering the inputDef production.
	EnterInputDef(c *InputDefContext)

	// EnterInputBody is called when entering the inputBody production.
	EnterInputBody(c *InputBodyContext)

	// EnterInputFieldDef is called when entering the inputFieldDef production.
	EnterInputFieldDef(c *InputFieldDefContext)

	// EnterEventDef is called when entering the eventDef production.
	EnterEventDef(c *EventDefContext)

	// EnterEventChannelHandler is called when entering the eventChannelHandler production.
	EnterEventChannelHandler(c *EventChannelHandlerContext)

	// EnterEnumDef is called when entering the enumDef production.
	EnterEnumDef(c *EnumDefContext)

	// EnterReferences is called when entering the references production.
	EnterReferences(c *ReferencesContext)

	// EnterLinkedBy is called when entering the linkedBy production.
	EnterLinkedBy(c *LinkedByContext)

	// EnterAction is called when entering the action production.
	EnterAction(c *ActionContext)

	// EnterParamList is called when entering the paramList production.
	EnterParamList(c *ParamListContext)

	// EnterBodyDef is called when entering the bodyDef production.
	EnterBodyDef(c *BodyDefContext)

	// EnterAuthenticateDef is called when entering the authenticateDef production.
	EnterAuthenticateDef(c *AuthenticateDefContext)

	// EnterReturnDef is called when entering the returnDef production.
	EnterReturnDef(c *ReturnDefContext)

	// EnterTypeDef is called when entering the typeDef production.
	EnterTypeDef(c *TypeDefContext)

	// EnterArrayDef is called when entering the arrayDef production.
	EnterArrayDef(c *ArrayDefContext)

	// EnterEventType is called when entering the eventType production.
	EnterEventType(c *EventTypeContext)

	// EnterFunctionCallExpr is called when entering the FunctionCallExpr production.
	EnterFunctionCallExpr(c *FunctionCallExprContext)

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

	// EnterRelationalExpr is called when entering the RelationalExpr production.
	EnterRelationalExpr(c *RelationalExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterPropertyAccessExpr is called when entering the PropertyAccessExpr production.
	EnterPropertyAccessExpr(c *PropertyAccessExprContext)

	// EnterAddSubExpr is called when entering the AddSubExpr production.
	EnterAddSubExpr(c *AddSubExprContext)

	// EnterLogicalExpr is called when entering the LogicalExpr production.
	EnterLogicalExpr(c *LogicalExprContext)

	// EnterMethodCallExpr is called when entering the MethodCallExpr production.
	EnterMethodCallExpr(c *MethodCallExprContext)

	// EnterMethodCallExpression is called when entering the methodCallExpression production.
	EnterMethodCallExpression(c *MethodCallExpressionContext)

	// EnterSimpleIdExpr is called when entering the SimpleIdExpr production.
	EnterSimpleIdExpr(c *SimpleIdExprContext)

	// EnterPrimaryFunctionCallExpr is called when entering the PrimaryFunctionCallExpr production.
	EnterPrimaryFunctionCallExpr(c *PrimaryFunctionCallExprContext)

	// EnterPrimaryPropertyAccessExpr is called when entering the PrimaryPropertyAccessExpr production.
	EnterPrimaryPropertyAccessExpr(c *PrimaryPropertyAccessExprContext)

	// EnterPropertyAccess is called when entering the propertyAccess production.
	EnterPropertyAccess(c *PropertyAccessContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

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

	// EnterSandboxBlock is called when entering the sandboxBlock production.
	EnterSandboxBlock(c *SandboxBlockContext)

	// EnterSandboxAttributes is called when entering the sandboxAttributes production.
	EnterSandboxAttributes(c *SandboxAttributesContext)

	// EnterSandboxAttribute is called when entering the sandboxAttribute production.
	EnterSandboxAttribute(c *SandboxAttributeContext)

	// EnterSandboxLangAttribute is called when entering the sandboxLangAttribute production.
	EnterSandboxLangAttribute(c *SandboxLangAttributeContext)

	// EnterSandboxDefineAttribute is called when entering the sandboxDefineAttribute production.
	EnterSandboxDefineAttribute(c *SandboxDefineAttributeContext)

	// EnterSandboxDefineContent is called when entering the sandboxDefineContent production.
	EnterSandboxDefineContent(c *SandboxDefineContentContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitModelDef is called when exiting the modelDef production.
	ExitModelDef(c *ModelDefContext)

	// ExitModelBody is called when exiting the modelBody production.
	ExitModelBody(c *ModelBodyContext)

	// ExitModelFieldDef is called when exiting the modelFieldDef production.
	ExitModelFieldDef(c *ModelFieldDefContext)

	// ExitRuleDef is called when exiting the ruleDef production.
	ExitRuleDef(c *RuleDefContext)

	// ExitRuleBody is called when exiting the ruleBody production.
	ExitRuleBody(c *RuleBodyContext)

	// ExitRuleStatement is called when exiting the ruleStatement production.
	ExitRuleStatement(c *RuleStatementContext)

	// ExitRuleParameters is called when exiting the ruleParameters production.
	ExitRuleParameters(c *RuleParametersContext)

	// ExitFunctionBody is called when exiting the functionBody production.
	ExitFunctionBody(c *FunctionBodyContext)

	// ExitAssignmentStatement is called when exiting the assignmentStatement production.
	ExitAssignmentStatement(c *AssignmentStatementContext)

	// ExitConditionalStatement is called when exiting the conditionalStatement production.
	ExitConditionalStatement(c *ConditionalStatementContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitRuleConditionExpression is called when exiting the ruleConditionExpression production.
	ExitRuleConditionExpression(c *RuleConditionExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitInExpression is called when exiting the inExpression production.
	ExitInExpression(c *InExpressionContext)

	// ExitRuleExpressionPrimary is called when exiting the ruleExpressionPrimary production.
	ExitRuleExpressionPrimary(c *RuleExpressionPrimaryContext)

	// ExitMutationDef is called when exiting the mutationDef production.
	ExitMutationDef(c *MutationDefContext)

	// ExitMutationBody is called when exiting the mutationBody production.
	ExitMutationBody(c *MutationBodyContext)

	// ExitMutationStatement is called when exiting the mutationStatement production.
	ExitMutationStatement(c *MutationStatementContext)

	// ExitMutationParameters is called when exiting the mutationParameters production.
	ExitMutationParameters(c *MutationParametersContext)

	// ExitMutationInnerBlock is called when exiting the mutationInnerBlock production.
	ExitMutationInnerBlock(c *MutationInnerBlockContext)

	// ExitMiddlewareBlock is called when exiting the middlewareBlock production.
	ExitMiddlewareBlock(c *MiddlewareBlockContext)

	// ExitFunctionBlock is called when exiting the functionBlock production.
	ExitFunctionBlock(c *FunctionBlockContext)

	// ExitInputDef is called when exiting the inputDef production.
	ExitInputDef(c *InputDefContext)

	// ExitInputBody is called when exiting the inputBody production.
	ExitInputBody(c *InputBodyContext)

	// ExitInputFieldDef is called when exiting the inputFieldDef production.
	ExitInputFieldDef(c *InputFieldDefContext)

	// ExitEventDef is called when exiting the eventDef production.
	ExitEventDef(c *EventDefContext)

	// ExitEventChannelHandler is called when exiting the eventChannelHandler production.
	ExitEventChannelHandler(c *EventChannelHandlerContext)

	// ExitEnumDef is called when exiting the enumDef production.
	ExitEnumDef(c *EnumDefContext)

	// ExitReferences is called when exiting the references production.
	ExitReferences(c *ReferencesContext)

	// ExitLinkedBy is called when exiting the linkedBy production.
	ExitLinkedBy(c *LinkedByContext)

	// ExitAction is called when exiting the action production.
	ExitAction(c *ActionContext)

	// ExitParamList is called when exiting the paramList production.
	ExitParamList(c *ParamListContext)

	// ExitBodyDef is called when exiting the bodyDef production.
	ExitBodyDef(c *BodyDefContext)

	// ExitAuthenticateDef is called when exiting the authenticateDef production.
	ExitAuthenticateDef(c *AuthenticateDefContext)

	// ExitReturnDef is called when exiting the returnDef production.
	ExitReturnDef(c *ReturnDefContext)

	// ExitTypeDef is called when exiting the typeDef production.
	ExitTypeDef(c *TypeDefContext)

	// ExitArrayDef is called when exiting the arrayDef production.
	ExitArrayDef(c *ArrayDefContext)

	// ExitEventType is called when exiting the eventType production.
	ExitEventType(c *EventTypeContext)

	// ExitFunctionCallExpr is called when exiting the FunctionCallExpr production.
	ExitFunctionCallExpr(c *FunctionCallExprContext)

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

	// ExitRelationalExpr is called when exiting the RelationalExpr production.
	ExitRelationalExpr(c *RelationalExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitPropertyAccessExpr is called when exiting the PropertyAccessExpr production.
	ExitPropertyAccessExpr(c *PropertyAccessExprContext)

	// ExitAddSubExpr is called when exiting the AddSubExpr production.
	ExitAddSubExpr(c *AddSubExprContext)

	// ExitLogicalExpr is called when exiting the LogicalExpr production.
	ExitLogicalExpr(c *LogicalExprContext)

	// ExitMethodCallExpr is called when exiting the MethodCallExpr production.
	ExitMethodCallExpr(c *MethodCallExprContext)

	// ExitMethodCallExpression is called when exiting the methodCallExpression production.
	ExitMethodCallExpression(c *MethodCallExpressionContext)

	// ExitSimpleIdExpr is called when exiting the SimpleIdExpr production.
	ExitSimpleIdExpr(c *SimpleIdExprContext)

	// ExitPrimaryFunctionCallExpr is called when exiting the PrimaryFunctionCallExpr production.
	ExitPrimaryFunctionCallExpr(c *PrimaryFunctionCallExprContext)

	// ExitPrimaryPropertyAccessExpr is called when exiting the PrimaryPropertyAccessExpr production.
	ExitPrimaryPropertyAccessExpr(c *PrimaryPropertyAccessExprContext)

	// ExitPropertyAccess is called when exiting the propertyAccess production.
	ExitPropertyAccess(c *PropertyAccessContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

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

	// ExitSandboxBlock is called when exiting the sandboxBlock production.
	ExitSandboxBlock(c *SandboxBlockContext)

	// ExitSandboxAttributes is called when exiting the sandboxAttributes production.
	ExitSandboxAttributes(c *SandboxAttributesContext)

	// ExitSandboxAttribute is called when exiting the sandboxAttribute production.
	ExitSandboxAttribute(c *SandboxAttributeContext)

	// ExitSandboxLangAttribute is called when exiting the sandboxLangAttribute production.
	ExitSandboxLangAttribute(c *SandboxLangAttributeContext)

	// ExitSandboxDefineAttribute is called when exiting the sandboxDefineAttribute production.
	ExitSandboxDefineAttribute(c *SandboxDefineAttributeContext)

	// ExitSandboxDefineContent is called when exiting the sandboxDefineContent production.
	ExitSandboxDefineContent(c *SandboxDefineContentContext)
}
