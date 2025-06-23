// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VlangParser.
type VlangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VlangParser#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by VlangParser#declaraciones.
	VisitDeclaraciones(ctx *DeclaracionesContext) interface{}

	// Visit a parse tree produced by VlangParser#structDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) interface{}

	// Visit a parse tree produced by VlangParser#atributos.
	VisitAtributos(ctx *AtributosContext) interface{}

	// Visit a parse tree produced by VlangParser#atributo.
	VisitAtributo(ctx *AtributoContext) interface{}

	// Visit a parse tree produced by VlangParser#inicializadorStruct.
	VisitInicializadorStruct(ctx *InicializadorStructContext) interface{}

	// Visit a parse tree produced by VlangParser#inicializadorCampo.
	VisitInicializadorCampo(ctx *InicializadorCampoContext) interface{}

	// Visit a parse tree produced by VlangParser#funcionStruct.
	VisitFuncionStruct(ctx *FuncionStructContext) interface{}

	// Visit a parse tree produced by VlangParser#funcion.
	VisitFuncion(ctx *FuncionContext) interface{}

	// Visit a parse tree produced by VlangParser#parametrosFormales.
	VisitParametrosFormales(ctx *ParametrosFormalesContext) interface{}

	// Visit a parse tree produced by VlangParser#parametroFormal.
	VisitParametroFormal(ctx *ParametroFormalContext) interface{}

	// Visit a parse tree produced by VlangParser#tipoConReferencia.
	VisitTipoConReferencia(ctx *TipoConReferenciaContext) interface{}

	// Visit a parse tree produced by VlangParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by VlangParser#tipo.
	VisitTipo(ctx *TipoContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceTipo.
	VisitSliceTipo(ctx *SliceTipoContext) interface{}

	// Visit a parse tree produced by VlangParser#expresionStatement.
	VisitExpresionStatement(ctx *ExpresionStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceDobleAsignacionStmt.
	VisitSliceDobleAsignacionStmt(ctx *SliceDobleAsignacionStmtContext) interface{}

	// Visit a parse tree produced by VlangParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#controlStatement.
	VisitControlStatement(ctx *ControlStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#bloqueStatement.
	VisitBloqueStatement(ctx *BloqueStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#if_context.
	VisitIf_context(ctx *If_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#for_context.
	VisitFor_context(ctx *For_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#while_context.
	VisitWhile_context(ctx *While_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#switch_context.
	VisitSwitch_context(ctx *Switch_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#ifDcl.
	VisitIfDcl(ctx *IfDclContext) interface{}

	// Visit a parse tree produced by VlangParser#elseifDcl.
	VisitElseifDcl(ctx *ElseifDclContext) interface{}

	// Visit a parse tree produced by VlangParser#elseDcl.
	VisitElseDcl(ctx *ElseDclContext) interface{}

	// Visit a parse tree produced by VlangParser#switchDcl.
	VisitSwitchDcl(ctx *SwitchDclContext) interface{}

	// Visit a parse tree produced by VlangParser#caseBlock.
	VisitCaseBlock(ctx *CaseBlockContext) interface{}

	// Visit a parse tree produced by VlangParser#defaultBlock.
	VisitDefaultBlock(ctx *DefaultBlockContext) interface{}

	// Visit a parse tree produced by VlangParser#forCondicional.
	VisitForCondicional(ctx *ForCondicionalContext) interface{}

	// Visit a parse tree produced by VlangParser#forClasico.
	VisitForClasico(ctx *ForClasicoContext) interface{}

	// Visit a parse tree produced by VlangParser#forForeach.
	VisitForForeach(ctx *ForForeachContext) interface{}

	// Visit a parse tree produced by VlangParser#whileDcl.
	VisitWhileDcl(ctx *WhileDclContext) interface{}

	// Visit a parse tree produced by VlangParser#bloque.
	VisitBloque(ctx *BloqueContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionAtributo.
	VisitAsignacionAtributo(ctx *AsignacionAtributoContext) interface{}

	// Visit a parse tree produced by VlangParser#multdivmod.
	VisitMultdivmod(ctx *MultdivmodContext) interface{}

	// Visit a parse tree produced by VlangParser#lenExpr.
	VisitLenExpr(ctx *LenExprContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceDobleAcceso.
	VisitSliceDobleAcceso(ctx *SliceDobleAccesoContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionResta.
	VisitAsignacionResta(ctx *AsignacionRestaContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionfor.
	VisitAsignacionfor(ctx *AsignacionforContext) interface{}

	// Visit a parse tree produced by VlangParser#literalSliceExpr.
	VisitLiteralSliceExpr(ctx *LiteralSliceExprContext) interface{}

	// Visit a parse tree produced by VlangParser#unario.
	VisitUnario(ctx *UnarioContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceAsignacion.
	VisitSliceAsignacion(ctx *SliceAsignacionContext) interface{}

	// Visit a parse tree produced by VlangParser#parentesisexpre.
	VisitParentesisexpre(ctx *ParentesisexpreContext) interface{}

	// Visit a parse tree produced by VlangParser#joinExpr.
	VisitJoinExpr(ctx *JoinExprContext) interface{}

	// Visit a parse tree produced by VlangParser#llamadaMetodoStruct.
	VisitLlamadaMetodoStruct(ctx *LlamadaMetodoStructContext) interface{}

	// Visit a parse tree produced by VlangParser#sumres.
	VisitSumres(ctx *SumresContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceCreacion.
	VisitSliceCreacion(ctx *SliceCreacionContext) interface{}

	// Visit a parse tree produced by VlangParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by VlangParser#appendExpr.
	VisitAppendExpr(ctx *AppendExprContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionCompuestaResta.
	VisitAsignacionCompuestaResta(ctx *AsignacionCompuestaRestaContext) interface{}

	// Visit a parse tree produced by VlangParser#valorexpr.
	VisitValorexpr(ctx *ValorexprContext) interface{}

	// Visit a parse tree produced by VlangParser#igualdad.
	VisitIgualdad(ctx *IgualdadContext) interface{}

	// Visit a parse tree produced by VlangParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceDobleAsignacion.
	VisitSliceDobleAsignacion(ctx *SliceDobleAsignacionContext) interface{}

	// Visit a parse tree produced by VlangParser#sliceAcceso.
	VisitSliceAcceso(ctx *SliceAccesoContext) interface{}

	// Visit a parse tree produced by VlangParser#relacionales.
	VisitRelacionales(ctx *RelacionalesContext) interface{}

	// Visit a parse tree produced by VlangParser#llamadaFuncion.
	VisitLlamadaFuncion(ctx *LlamadaFuncionContext) interface{}

	// Visit a parse tree produced by VlangParser#accesoAtributo.
	VisitAccesoAtributo(ctx *AccesoAtributoContext) interface{}

	// Visit a parse tree produced by VlangParser#instanciaStruct.
	VisitInstanciaStruct(ctx *InstanciaStructContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionCompuestaSuma.
	VisitAsignacionCompuestaSuma(ctx *AsignacionCompuestaSumaContext) interface{}

	// Visit a parse tree produced by VlangParser#indexOfExpr.
	VisitIndexOfExpr(ctx *IndexOfExprContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionSuma.
	VisitAsignacionSuma(ctx *AsignacionSumaContext) interface{}

	// Visit a parse tree produced by VlangParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by VlangParser#parametros.
	VisitParametros(ctx *ParametrosContext) interface{}

	// Visit a parse tree produced by VlangParser#valorEntero.
	VisitValorEntero(ctx *ValorEnteroContext) interface{}

	// Visit a parse tree produced by VlangParser#valorDecimal.
	VisitValorDecimal(ctx *ValorDecimalContext) interface{}

	// Visit a parse tree produced by VlangParser#valorCadena.
	VisitValorCadena(ctx *ValorCadenaContext) interface{}

	// Visit a parse tree produced by VlangParser#valorBooleano.
	VisitValorBooleano(ctx *ValorBooleanoContext) interface{}

	// Visit a parse tree produced by VlangParser#valorCaracter.
	VisitValorCaracter(ctx *ValorCaracterContext) interface{}
}
