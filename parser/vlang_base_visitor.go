// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

type BaseVlangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVlangVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDeclaraciones(ctx *DeclaracionesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitStructDeclaration(ctx *StructDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAtributos(ctx *AtributosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAtributo(ctx *AtributoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitInicializadorStruct(ctx *InicializadorStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitInicializadorCampo(ctx *InicializadorCampoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncionStruct(ctx *FuncionStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFuncion(ctx *FuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParametrosFormales(ctx *ParametrosFormalesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParametroFormal(ctx *ParametroFormalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitTipoConReferencia(ctx *TipoConReferenciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitTipo(ctx *TipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceTipo(ctx *SliceTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitExpresionStatement(ctx *ExpresionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceDobleAsignacionStmt(ctx *SliceDobleAsignacionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitPrintStatement(ctx *PrintStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitControlStatement(ctx *ControlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBloqueStatement(ctx *BloqueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIf_context(ctx *If_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitFor_context(ctx *For_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitWhile_context(ctx *While_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSwitch_context(ctx *Switch_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIfDcl(ctx *IfDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitElseifDcl(ctx *ElseifDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitElseDcl(ctx *ElseDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSwitchDcl(ctx *SwitchDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitCaseBlock(ctx *CaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitDefaultBlock(ctx *DefaultBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitForCondicional(ctx *ForCondicionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitForClasico(ctx *ForClasicoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitForForeach(ctx *ForForeachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitWhileDcl(ctx *WhileDclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitBloque(ctx *BloqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacionAtributo(ctx *AsignacionAtributoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitMultdivmod(ctx *MultdivmodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitLenExpr(ctx *LenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceDobleAcceso(ctx *SliceDobleAccesoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacionResta(ctx *AsignacionRestaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacionfor(ctx *AsignacionforContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitLiteralSliceExpr(ctx *LiteralSliceExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitUnario(ctx *UnarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceAsignacion(ctx *SliceAsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParentesisexpre(ctx *ParentesisexpreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitJoinExpr(ctx *JoinExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitLlamadaMetodoStruct(ctx *LlamadaMetodoStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSumres(ctx *SumresContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceCreacion(ctx *SliceCreacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAppendExpr(ctx *AppendExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacionCompuestaResta(ctx *AsignacionCompuestaRestaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorexpr(ctx *ValorexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIgualdad(ctx *IgualdadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceDobleAsignacion(ctx *SliceDobleAsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitSliceAcceso(ctx *SliceAccesoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitRelacionales(ctx *RelacionalesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitLlamadaFuncion(ctx *LlamadaFuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAccesoAtributo(ctx *AccesoAtributoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitInstanciaStruct(ctx *InstanciaStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacionCompuestaSuma(ctx *AsignacionCompuestaSumaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitIndexOfExpr(ctx *IndexOfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAsignacionSuma(ctx *AsignacionSumaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitParametros(ctx *ParametrosContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorEntero(ctx *ValorEnteroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorDecimal(ctx *ValorDecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorCadena(ctx *ValorCadenaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorBooleano(ctx *ValorBooleanoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVlangVisitor) VisitValorCaracter(ctx *ValorCaracterContext) interface{} {
	return v.VisitChildren(ctx)
}
