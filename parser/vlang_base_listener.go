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

// EnterStructDeclaration is called when production structDeclaration is entered.
func (s *BaseVlangListener) EnterStructDeclaration(ctx *StructDeclarationContext) {}

// ExitStructDeclaration is called when production structDeclaration is exited.
func (s *BaseVlangListener) ExitStructDeclaration(ctx *StructDeclarationContext) {}

// EnterAtributos is called when production atributos is entered.
func (s *BaseVlangListener) EnterAtributos(ctx *AtributosContext) {}

// ExitAtributos is called when production atributos is exited.
func (s *BaseVlangListener) ExitAtributos(ctx *AtributosContext) {}

// EnterAtributo is called when production atributo is entered.
func (s *BaseVlangListener) EnterAtributo(ctx *AtributoContext) {}

// ExitAtributo is called when production atributo is exited.
func (s *BaseVlangListener) ExitAtributo(ctx *AtributoContext) {}

// EnterInicializadorStruct is called when production inicializadorStruct is entered.
func (s *BaseVlangListener) EnterInicializadorStruct(ctx *InicializadorStructContext) {}

// ExitInicializadorStruct is called when production inicializadorStruct is exited.
func (s *BaseVlangListener) ExitInicializadorStruct(ctx *InicializadorStructContext) {}

// EnterInicializadorCampo is called when production inicializadorCampo is entered.
func (s *BaseVlangListener) EnterInicializadorCampo(ctx *InicializadorCampoContext) {}

// ExitInicializadorCampo is called when production inicializadorCampo is exited.
func (s *BaseVlangListener) ExitInicializadorCampo(ctx *InicializadorCampoContext) {}

// EnterFuncionStruct is called when production funcionStruct is entered.
func (s *BaseVlangListener) EnterFuncionStruct(ctx *FuncionStructContext) {}

// ExitFuncionStruct is called when production funcionStruct is exited.
func (s *BaseVlangListener) ExitFuncionStruct(ctx *FuncionStructContext) {}

// EnterFuncion is called when production funcion is entered.
func (s *BaseVlangListener) EnterFuncion(ctx *FuncionContext) {}

// ExitFuncion is called when production funcion is exited.
func (s *BaseVlangListener) ExitFuncion(ctx *FuncionContext) {}

// EnterParametrosFormales is called when production parametrosFormales is entered.
func (s *BaseVlangListener) EnterParametrosFormales(ctx *ParametrosFormalesContext) {}

// ExitParametrosFormales is called when production parametrosFormales is exited.
func (s *BaseVlangListener) ExitParametrosFormales(ctx *ParametrosFormalesContext) {}

// EnterParametroFormal is called when production parametroFormal is entered.
func (s *BaseVlangListener) EnterParametroFormal(ctx *ParametroFormalContext) {}

// ExitParametroFormal is called when production parametroFormal is exited.
func (s *BaseVlangListener) ExitParametroFormal(ctx *ParametroFormalContext) {}

// EnterTipoConReferencia is called when production tipoConReferencia is entered.
func (s *BaseVlangListener) EnterTipoConReferencia(ctx *TipoConReferenciaContext) {}

// ExitTipoConReferencia is called when production tipoConReferencia is exited.
func (s *BaseVlangListener) ExitTipoConReferencia(ctx *TipoConReferenciaContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseVlangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseVlangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseVlangListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseVlangListener) ExitTipo(ctx *TipoContext) {}

// EnterSliceTipo is called when production sliceTipo is entered.
func (s *BaseVlangListener) EnterSliceTipo(ctx *SliceTipoContext) {}

// ExitSliceTipo is called when production sliceTipo is exited.
func (s *BaseVlangListener) ExitSliceTipo(ctx *SliceTipoContext) {}

// EnterExpresionStatement is called when production expresionStatement is entered.
func (s *BaseVlangListener) EnterExpresionStatement(ctx *ExpresionStatementContext) {}

// ExitExpresionStatement is called when production expresionStatement is exited.
func (s *BaseVlangListener) ExitExpresionStatement(ctx *ExpresionStatementContext) {}

// EnterSliceDobleAsignacionStmt is called when production sliceDobleAsignacionStmt is entered.
func (s *BaseVlangListener) EnterSliceDobleAsignacionStmt(ctx *SliceDobleAsignacionStmtContext) {}

// ExitSliceDobleAsignacionStmt is called when production sliceDobleAsignacionStmt is exited.
func (s *BaseVlangListener) ExitSliceDobleAsignacionStmt(ctx *SliceDobleAsignacionStmtContext) {}

// EnterPrintStatement is called when production printStatement is entered.
func (s *BaseVlangListener) EnterPrintStatement(ctx *PrintStatementContext) {}

// ExitPrintStatement is called when production printStatement is exited.
func (s *BaseVlangListener) ExitPrintStatement(ctx *PrintStatementContext) {}

// EnterBreakStatement is called when production breakStatement is entered.
func (s *BaseVlangListener) EnterBreakStatement(ctx *BreakStatementContext) {}

// ExitBreakStatement is called when production breakStatement is exited.
func (s *BaseVlangListener) ExitBreakStatement(ctx *BreakStatementContext) {}

// EnterContinueStatement is called when production continueStatement is entered.
func (s *BaseVlangListener) EnterContinueStatement(ctx *ContinueStatementContext) {}

// ExitContinueStatement is called when production continueStatement is exited.
func (s *BaseVlangListener) ExitContinueStatement(ctx *ContinueStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseVlangListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseVlangListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterControlStatement is called when production controlStatement is entered.
func (s *BaseVlangListener) EnterControlStatement(ctx *ControlStatementContext) {}

// ExitControlStatement is called when production controlStatement is exited.
func (s *BaseVlangListener) ExitControlStatement(ctx *ControlStatementContext) {}

// EnterBloqueStatement is called when production bloqueStatement is entered.
func (s *BaseVlangListener) EnterBloqueStatement(ctx *BloqueStatementContext) {}

// ExitBloqueStatement is called when production bloqueStatement is exited.
func (s *BaseVlangListener) ExitBloqueStatement(ctx *BloqueStatementContext) {}

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

// EnterSwitch_context is called when production switch_context is entered.
func (s *BaseVlangListener) EnterSwitch_context(ctx *Switch_contextContext) {}

// ExitSwitch_context is called when production switch_context is exited.
func (s *BaseVlangListener) ExitSwitch_context(ctx *Switch_contextContext) {}

// EnterIfDcl is called when production ifDcl is entered.
func (s *BaseVlangListener) EnterIfDcl(ctx *IfDclContext) {}

// ExitIfDcl is called when production ifDcl is exited.
func (s *BaseVlangListener) ExitIfDcl(ctx *IfDclContext) {}

// EnterElseifDcl is called when production elseifDcl is entered.
func (s *BaseVlangListener) EnterElseifDcl(ctx *ElseifDclContext) {}

// ExitElseifDcl is called when production elseifDcl is exited.
func (s *BaseVlangListener) ExitElseifDcl(ctx *ElseifDclContext) {}

// EnterElseDcl is called when production elseDcl is entered.
func (s *BaseVlangListener) EnterElseDcl(ctx *ElseDclContext) {}

// ExitElseDcl is called when production elseDcl is exited.
func (s *BaseVlangListener) ExitElseDcl(ctx *ElseDclContext) {}

// EnterSwitchDcl is called when production switchDcl is entered.
func (s *BaseVlangListener) EnterSwitchDcl(ctx *SwitchDclContext) {}

// ExitSwitchDcl is called when production switchDcl is exited.
func (s *BaseVlangListener) ExitSwitchDcl(ctx *SwitchDclContext) {}

// EnterCaseBlock is called when production caseBlock is entered.
func (s *BaseVlangListener) EnterCaseBlock(ctx *CaseBlockContext) {}

// ExitCaseBlock is called when production caseBlock is exited.
func (s *BaseVlangListener) ExitCaseBlock(ctx *CaseBlockContext) {}

// EnterDefaultBlock is called when production defaultBlock is entered.
func (s *BaseVlangListener) EnterDefaultBlock(ctx *DefaultBlockContext) {}

// ExitDefaultBlock is called when production defaultBlock is exited.
func (s *BaseVlangListener) ExitDefaultBlock(ctx *DefaultBlockContext) {}

// EnterForCondicional is called when production forCondicional is entered.
func (s *BaseVlangListener) EnterForCondicional(ctx *ForCondicionalContext) {}

// ExitForCondicional is called when production forCondicional is exited.
func (s *BaseVlangListener) ExitForCondicional(ctx *ForCondicionalContext) {}

// EnterForClasico is called when production forClasico is entered.
func (s *BaseVlangListener) EnterForClasico(ctx *ForClasicoContext) {}

// ExitForClasico is called when production forClasico is exited.
func (s *BaseVlangListener) ExitForClasico(ctx *ForClasicoContext) {}

// EnterForForeach is called when production forForeach is entered.
func (s *BaseVlangListener) EnterForForeach(ctx *ForForeachContext) {}

// ExitForForeach is called when production forForeach is exited.
func (s *BaseVlangListener) ExitForForeach(ctx *ForForeachContext) {}

// EnterWhileDcl is called when production whileDcl is entered.
func (s *BaseVlangListener) EnterWhileDcl(ctx *WhileDclContext) {}

// ExitWhileDcl is called when production whileDcl is exited.
func (s *BaseVlangListener) ExitWhileDcl(ctx *WhileDclContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BaseVlangListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BaseVlangListener) ExitBloque(ctx *BloqueContext) {}

// EnterAsignacionAtributo is called when production asignacionAtributo is entered.
func (s *BaseVlangListener) EnterAsignacionAtributo(ctx *AsignacionAtributoContext) {}

// ExitAsignacionAtributo is called when production asignacionAtributo is exited.
func (s *BaseVlangListener) ExitAsignacionAtributo(ctx *AsignacionAtributoContext) {}

// EnterMultdivmod is called when production multdivmod is entered.
func (s *BaseVlangListener) EnterMultdivmod(ctx *MultdivmodContext) {}

// ExitMultdivmod is called when production multdivmod is exited.
func (s *BaseVlangListener) ExitMultdivmod(ctx *MultdivmodContext) {}

// EnterLenExpr is called when production lenExpr is entered.
func (s *BaseVlangListener) EnterLenExpr(ctx *LenExprContext) {}

// ExitLenExpr is called when production lenExpr is exited.
func (s *BaseVlangListener) ExitLenExpr(ctx *LenExprContext) {}

// EnterSliceDobleAcceso is called when production sliceDobleAcceso is entered.
func (s *BaseVlangListener) EnterSliceDobleAcceso(ctx *SliceDobleAccesoContext) {}

// ExitSliceDobleAcceso is called when production sliceDobleAcceso is exited.
func (s *BaseVlangListener) ExitSliceDobleAcceso(ctx *SliceDobleAccesoContext) {}

// EnterAsignacionResta is called when production asignacionResta is entered.
func (s *BaseVlangListener) EnterAsignacionResta(ctx *AsignacionRestaContext) {}

// ExitAsignacionResta is called when production asignacionResta is exited.
func (s *BaseVlangListener) ExitAsignacionResta(ctx *AsignacionRestaContext) {}

// EnterAsignacionfor is called when production asignacionfor is entered.
func (s *BaseVlangListener) EnterAsignacionfor(ctx *AsignacionforContext) {}

// ExitAsignacionfor is called when production asignacionfor is exited.
func (s *BaseVlangListener) ExitAsignacionfor(ctx *AsignacionforContext) {}

// EnterLiteralSliceExpr is called when production literalSliceExpr is entered.
func (s *BaseVlangListener) EnterLiteralSliceExpr(ctx *LiteralSliceExprContext) {}

// ExitLiteralSliceExpr is called when production literalSliceExpr is exited.
func (s *BaseVlangListener) ExitLiteralSliceExpr(ctx *LiteralSliceExprContext) {}

// EnterUnario is called when production unario is entered.
func (s *BaseVlangListener) EnterUnario(ctx *UnarioContext) {}

// ExitUnario is called when production unario is exited.
func (s *BaseVlangListener) ExitUnario(ctx *UnarioContext) {}

// EnterSliceAsignacion is called when production sliceAsignacion is entered.
func (s *BaseVlangListener) EnterSliceAsignacion(ctx *SliceAsignacionContext) {}

// ExitSliceAsignacion is called when production sliceAsignacion is exited.
func (s *BaseVlangListener) ExitSliceAsignacion(ctx *SliceAsignacionContext) {}

// EnterParentesisexpre is called when production parentesisexpre is entered.
func (s *BaseVlangListener) EnterParentesisexpre(ctx *ParentesisexpreContext) {}

// ExitParentesisexpre is called when production parentesisexpre is exited.
func (s *BaseVlangListener) ExitParentesisexpre(ctx *ParentesisexpreContext) {}

// EnterJoinExpr is called when production joinExpr is entered.
func (s *BaseVlangListener) EnterJoinExpr(ctx *JoinExprContext) {}

// ExitJoinExpr is called when production joinExpr is exited.
func (s *BaseVlangListener) ExitJoinExpr(ctx *JoinExprContext) {}

// EnterLlamadaMetodoStruct is called when production llamadaMetodoStruct is entered.
func (s *BaseVlangListener) EnterLlamadaMetodoStruct(ctx *LlamadaMetodoStructContext) {}

// ExitLlamadaMetodoStruct is called when production llamadaMetodoStruct is exited.
func (s *BaseVlangListener) ExitLlamadaMetodoStruct(ctx *LlamadaMetodoStructContext) {}

// EnterSumres is called when production sumres is entered.
func (s *BaseVlangListener) EnterSumres(ctx *SumresContext) {}

// ExitSumres is called when production sumres is exited.
func (s *BaseVlangListener) ExitSumres(ctx *SumresContext) {}

// EnterSliceCreacion is called when production sliceCreacion is entered.
func (s *BaseVlangListener) EnterSliceCreacion(ctx *SliceCreacionContext) {}

// ExitSliceCreacion is called when production sliceCreacion is exited.
func (s *BaseVlangListener) ExitSliceCreacion(ctx *SliceCreacionContext) {}

// EnterId is called when production id is entered.
func (s *BaseVlangListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseVlangListener) ExitId(ctx *IdContext) {}

// EnterAppendExpr is called when production appendExpr is entered.
func (s *BaseVlangListener) EnterAppendExpr(ctx *AppendExprContext) {}

// ExitAppendExpr is called when production appendExpr is exited.
func (s *BaseVlangListener) ExitAppendExpr(ctx *AppendExprContext) {}

// EnterAsignacionCompuestaResta is called when production asignacionCompuestaResta is entered.
func (s *BaseVlangListener) EnterAsignacionCompuestaResta(ctx *AsignacionCompuestaRestaContext) {}

// ExitAsignacionCompuestaResta is called when production asignacionCompuestaResta is exited.
func (s *BaseVlangListener) ExitAsignacionCompuestaResta(ctx *AsignacionCompuestaRestaContext) {}

// EnterValorexpr is called when production valorexpr is entered.
func (s *BaseVlangListener) EnterValorexpr(ctx *ValorexprContext) {}

// ExitValorexpr is called when production valorexpr is exited.
func (s *BaseVlangListener) ExitValorexpr(ctx *ValorexprContext) {}

// EnterIgualdad is called when production igualdad is entered.
func (s *BaseVlangListener) EnterIgualdad(ctx *IgualdadContext) {}

// ExitIgualdad is called when production igualdad is exited.
func (s *BaseVlangListener) ExitIgualdad(ctx *IgualdadContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseVlangListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseVlangListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterSliceDobleAsignacion is called when production sliceDobleAsignacion is entered.
func (s *BaseVlangListener) EnterSliceDobleAsignacion(ctx *SliceDobleAsignacionContext) {}

// ExitSliceDobleAsignacion is called when production sliceDobleAsignacion is exited.
func (s *BaseVlangListener) ExitSliceDobleAsignacion(ctx *SliceDobleAsignacionContext) {}

// EnterSliceAcceso is called when production sliceAcceso is entered.
func (s *BaseVlangListener) EnterSliceAcceso(ctx *SliceAccesoContext) {}

// ExitSliceAcceso is called when production sliceAcceso is exited.
func (s *BaseVlangListener) ExitSliceAcceso(ctx *SliceAccesoContext) {}

// EnterRelacionales is called when production relacionales is entered.
func (s *BaseVlangListener) EnterRelacionales(ctx *RelacionalesContext) {}

// ExitRelacionales is called when production relacionales is exited.
func (s *BaseVlangListener) ExitRelacionales(ctx *RelacionalesContext) {}

// EnterLlamadaFuncion is called when production llamadaFuncion is entered.
func (s *BaseVlangListener) EnterLlamadaFuncion(ctx *LlamadaFuncionContext) {}

// ExitLlamadaFuncion is called when production llamadaFuncion is exited.
func (s *BaseVlangListener) ExitLlamadaFuncion(ctx *LlamadaFuncionContext) {}

// EnterAccesoAtributo is called when production accesoAtributo is entered.
func (s *BaseVlangListener) EnterAccesoAtributo(ctx *AccesoAtributoContext) {}

// ExitAccesoAtributo is called when production accesoAtributo is exited.
func (s *BaseVlangListener) ExitAccesoAtributo(ctx *AccesoAtributoContext) {}

// EnterInstanciaStruct is called when production instanciaStruct is entered.
func (s *BaseVlangListener) EnterInstanciaStruct(ctx *InstanciaStructContext) {}

// ExitInstanciaStruct is called when production instanciaStruct is exited.
func (s *BaseVlangListener) ExitInstanciaStruct(ctx *InstanciaStructContext) {}

// EnterAsignacionCompuestaSuma is called when production asignacionCompuestaSuma is entered.
func (s *BaseVlangListener) EnterAsignacionCompuestaSuma(ctx *AsignacionCompuestaSumaContext) {}

// ExitAsignacionCompuestaSuma is called when production asignacionCompuestaSuma is exited.
func (s *BaseVlangListener) ExitAsignacionCompuestaSuma(ctx *AsignacionCompuestaSumaContext) {}

// EnterIndexOfExpr is called when production indexOfExpr is entered.
func (s *BaseVlangListener) EnterIndexOfExpr(ctx *IndexOfExprContext) {}

// ExitIndexOfExpr is called when production indexOfExpr is exited.
func (s *BaseVlangListener) ExitIndexOfExpr(ctx *IndexOfExprContext) {}

// EnterAsignacionSuma is called when production asignacionSuma is entered.
func (s *BaseVlangListener) EnterAsignacionSuma(ctx *AsignacionSumaContext) {}

// ExitAsignacionSuma is called when production asignacionSuma is exited.
func (s *BaseVlangListener) ExitAsignacionSuma(ctx *AsignacionSumaContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseVlangListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseVlangListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterParametros is called when production parametros is entered.
func (s *BaseVlangListener) EnterParametros(ctx *ParametrosContext) {}

// ExitParametros is called when production parametros is exited.
func (s *BaseVlangListener) ExitParametros(ctx *ParametrosContext) {}

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
