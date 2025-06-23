// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// BaseVlangListener is a complete listener for a parse tree produced by VlangParser.
type BaseVlangListener struct{}

var _ VlangListener = &BaseVlangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVlangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVlangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVlangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVlangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseVlangListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseVlangListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterDeclaraciones is called when production declaraciones is entered.
func (s *BaseVlangListener) EnterDeclaraciones(ctx *DeclaracionesContext) {}

// ExitDeclaraciones is called when production declaraciones is exited.
func (s *BaseVlangListener) ExitDeclaraciones(ctx *DeclaracionesContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseVlangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseVlangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterExpresionStatement is called when production expresionStatement is entered.
func (s *BaseVlangListener) EnterExpresionStatement(ctx *ExpresionStatementContext) {}

// ExitExpresionStatement is called when production expresionStatement is exited.
func (s *BaseVlangListener) ExitExpresionStatement(ctx *ExpresionStatementContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseVlangListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseVlangListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterControlStatement is called when production controlStatement is entered.
func (s *BaseVlangListener) EnterControlStatement(ctx *ControlStatementContext) {}

// ExitControlStatement is called when production controlStatement is exited.
func (s *BaseVlangListener) ExitControlStatement(ctx *ControlStatementContext) {}

// EnterIf_context is called when production if_context is entered.
func (s *BaseVlangListener) EnterIf_context(ctx *If_contextContext) {}

// ExitIf_context is called when production if_context is exited.
func (s *BaseVlangListener) ExitIf_context(ctx *If_contextContext) {}

// EnterFor_context is called when production for_context is entered.
func (s *BaseVlangListener) EnterFor_context(ctx *For_contextContext) {}

// ExitFor_context is called when production for_context is exited.
func (s *BaseVlangListener) ExitFor_context(ctx *For_contextContext) {}

// EnterWhile_context is called when production while_context is entered.
func (s *BaseVlangListener) EnterWhile_context(ctx *While_contextContext) {}

// ExitWhile_context is called when production while_context is exited.
func (s *BaseVlangListener) ExitWhile_context(ctx *While_contextContext) {}

// EnterIfDcl is called when production ifDcl is entered.
func (s *BaseVlangListener) EnterIfDcl(ctx *IfDclContext) {}

// ExitIfDcl is called when production ifDcl is exited.
func (s *BaseVlangListener) ExitIfDcl(ctx *IfDclContext) {}

// EnterForDcl is called when production forDcl is entered.
func (s *BaseVlangListener) EnterForDcl(ctx *ForDclContext) {}

// ExitForDcl is called when production forDcl is exited.
func (s *BaseVlangListener) ExitForDcl(ctx *ForDclContext) {}

// EnterWhileDcl is called when production whileDcl is entered.
func (s *BaseVlangListener) EnterWhileDcl(ctx *WhileDclContext) {}

// ExitWhileDcl is called when production whileDcl is exited.
func (s *BaseVlangListener) ExitWhileDcl(ctx *WhileDclContext) {}

// EnterMultdivmod is called when production multdivmod is entered.
func (s *BaseVlangListener) EnterMultdivmod(ctx *MultdivmodContext) {}

// ExitMultdivmod is called when production multdivmod is exited.
func (s *BaseVlangListener) ExitMultdivmod(ctx *MultdivmodContext) {}

// EnterIncredecr is called when production incredecr is entered.
func (s *BaseVlangListener) EnterIncredecr(ctx *IncredecrContext) {}

// ExitIncredecr is called when production incredecr is exited.
func (s *BaseVlangListener) ExitIncredecr(ctx *IncredecrContext) {}

// EnterOr is called when production or is entered.
func (s *BaseVlangListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseVlangListener) ExitOr(ctx *OrContext) {}

// EnterValorexpr is called when production valorexpr is entered.
func (s *BaseVlangListener) EnterValorexpr(ctx *ValorexprContext) {}

// ExitValorexpr is called when production valorexpr is exited.
func (s *BaseVlangListener) ExitValorexpr(ctx *ValorexprContext) {}

// EnterIgualdad is called when production igualdad is entered.
func (s *BaseVlangListener) EnterIgualdad(ctx *IgualdadContext) {}

// ExitIgualdad is called when production igualdad is exited.
func (s *BaseVlangListener) ExitIgualdad(ctx *IgualdadContext) {}

// EnterAsignacionfor is called when production asignacionfor is entered.
func (s *BaseVlangListener) EnterAsignacionfor(ctx *AsignacionforContext) {}

// ExitAsignacionfor is called when production asignacionfor is exited.
func (s *BaseVlangListener) ExitAsignacionfor(ctx *AsignacionforContext) {}

// EnterExpdotexp is called when production expdotexp is entered.
func (s *BaseVlangListener) EnterExpdotexp(ctx *ExpdotexpContext) {}

// ExitExpdotexp is called when production expdotexp is exited.
func (s *BaseVlangListener) ExitExpdotexp(ctx *ExpdotexpContext) {}

// EnterRelacionales is called when production relacionales is entered.
func (s *BaseVlangListener) EnterRelacionales(ctx *RelacionalesContext) {}

// ExitRelacionales is called when production relacionales is exited.
func (s *BaseVlangListener) ExitRelacionales(ctx *RelacionalesContext) {}

// EnterCorchetesexpre is called when production corchetesexpre is entered.
func (s *BaseVlangListener) EnterCorchetesexpre(ctx *CorchetesexpreContext) {}

// ExitCorchetesexpre is called when production corchetesexpre is exited.
func (s *BaseVlangListener) ExitCorchetesexpre(ctx *CorchetesexpreContext) {}

// EnterUnario is called when production unario is entered.
func (s *BaseVlangListener) EnterUnario(ctx *UnarioContext) {}

// ExitUnario is called when production unario is exited.
func (s *BaseVlangListener) ExitUnario(ctx *UnarioContext) {}

// EnterParentesisexpre is called when production parentesisexpre is entered.
func (s *BaseVlangListener) EnterParentesisexpre(ctx *ParentesisexpreContext) {}

// ExitParentesisexpre is called when production parentesisexpre is exited.
func (s *BaseVlangListener) ExitParentesisexpre(ctx *ParentesisexpreContext) {}

// EnterSumres is called when production sumres is entered.
func (s *BaseVlangListener) EnterSumres(ctx *SumresContext) {}

// ExitSumres is called when production sumres is exited.
func (s *BaseVlangListener) ExitSumres(ctx *SumresContext) {}

// EnterId is called when production id is entered.
func (s *BaseVlangListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseVlangListener) ExitId(ctx *IdContext) {}

// EnterExpdotexp1 is called when production expdotexp1 is entered.
func (s *BaseVlangListener) EnterExpdotexp1(ctx *Expdotexp1Context) {}

// ExitExpdotexp1 is called when production expdotexp1 is exited.
func (s *BaseVlangListener) ExitExpdotexp1(ctx *Expdotexp1Context) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseVlangListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseVlangListener) ExitParametros(ctx *ParametrosContext) {}

// EnterValores is called when production valores is entered.
func (s *BaseVlangListener) EnterValores(ctx *ValoresContext) {}

// ExitValores is called when production valores is exited.
func (s *BaseVlangListener) ExitValores(ctx *ValoresContext) {}

// EnterValorEntero is called when production valorEntero is entered.
func (s *BaseVlangListener) EnterValorEntero(ctx *ValorEnteroContext) {}

// ExitValorEntero is called when production valorEntero is exited.
func (s *BaseVlangListener) ExitValorEntero(ctx *ValorEnteroContext) {}

// EnterValorDecimal is called when production valorDecimal is entered.
func (s *BaseVlangListener) EnterValorDecimal(ctx *ValorDecimalContext) {}

// ExitValorDecimal is called when production valorDecimal is exited.
func (s *BaseVlangListener) ExitValorDecimal(ctx *ValorDecimalContext) {}

// EnterValorCadena is called when production valorCadena is entered.
func (s *BaseVlangListener) EnterValorCadena(ctx *ValorCadenaContext) {}

// ExitValorCadena is called when production valorCadena is exited.
func (s *BaseVlangListener) ExitValorCadena(ctx *ValorCadenaContext) {}

// EnterValorBooleano is called when production valorBooleano is entered.
func (s *BaseVlangListener) EnterValorBooleano(ctx *ValorBooleanoContext) {}

// ExitValorBooleano is called when production valorBooleano is exited.
func (s *BaseVlangListener) ExitValorBooleano(ctx *ValorBooleanoContext) {}

// EnterValorCaracter is called when production valorCaracter is entered.
func (s *BaseVlangListener) EnterValorCaracter(ctx *ValorCaracterContext) {}

// ExitValorCaracter is called when production valorCaracter is exited.
func (s *BaseVlangListener) ExitValorCaracter(ctx *ValorCaracterContext) {}

// EnterIncremento is called when production incremento is entered.
func (s *BaseVlangListener) EnterIncremento(ctx *IncrementoContext) {}

// ExitIncremento is called when production incremento is exited.
func (s *BaseVlangListener) ExitIncremento(ctx *IncrementoContext) {}

// EnterDecremento is called when production decremento is entered.
func (s *BaseVlangListener) EnterDecremento(ctx *DecrementoContext) {}

// ExitDecremento is called when production decremento is exited.
func (s *BaseVlangListener) ExitDecremento(ctx *DecrementoContext) {}
