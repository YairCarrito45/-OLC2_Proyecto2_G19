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

	// Visit a parse tree produced by VlangParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by VlangParser#expresionStatement.
	VisitExpresionStatement(ctx *ExpresionStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#printStatement.
	VisitPrintStatement(ctx *PrintStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#controlStatement.
	VisitControlStatement(ctx *ControlStatementContext) interface{}

	// Visit a parse tree produced by VlangParser#if_context.
	VisitIf_context(ctx *If_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#for_context.
	VisitFor_context(ctx *For_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#while_context.
	VisitWhile_context(ctx *While_contextContext) interface{}

	// Visit a parse tree produced by VlangParser#ifDcl.
	VisitIfDcl(ctx *IfDclContext) interface{}

	// Visit a parse tree produced by VlangParser#forDcl.
	VisitForDcl(ctx *ForDclContext) interface{}

	// Visit a parse tree produced by VlangParser#whileDcl.
	VisitWhileDcl(ctx *WhileDclContext) interface{}

	// Visit a parse tree produced by VlangParser#multdivmod.
	VisitMultdivmod(ctx *MultdivmodContext) interface{}

	// Visit a parse tree produced by VlangParser#incredecr.
	VisitIncredecr(ctx *IncredecrContext) interface{}

	// Visit a parse tree produced by VlangParser#or.
	VisitOr(ctx *OrContext) interface{}

	// Visit a parse tree produced by VlangParser#valorexpr.
	VisitValorexpr(ctx *ValorexprContext) interface{}

	// Visit a parse tree produced by VlangParser#igualdad.
	VisitIgualdad(ctx *IgualdadContext) interface{}

	// Visit a parse tree produced by VlangParser#asignacionfor.
	VisitAsignacionfor(ctx *AsignacionforContext) interface{}

	// Visit a parse tree produced by VlangParser#expdotexp.
	VisitExpdotexp(ctx *ExpdotexpContext) interface{}

	// Visit a parse tree produced by VlangParser#relacionales.
	VisitRelacionales(ctx *RelacionalesContext) interface{}

	// Visit a parse tree produced by VlangParser#corchetesexpre.
	VisitCorchetesexpre(ctx *CorchetesexpreContext) interface{}

	// Visit a parse tree produced by VlangParser#unario.
	VisitUnario(ctx *UnarioContext) interface{}

	// Visit a parse tree produced by VlangParser#parentesisexpre.
	VisitParentesisexpre(ctx *ParentesisexpreContext) interface{}

	// Visit a parse tree produced by VlangParser#sumres.
	VisitSumres(ctx *SumresContext) interface{}

	// Visit a parse tree produced by VlangParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by VlangParser#expdotexp1.
	VisitExpdotexp1(ctx *Expdotexp1Context) interface{}

	// Visit a parse tree produced by VlangParser#parametros.
	VisitParametros(ctx *ParametrosContext) interface{}

	// Visit a parse tree produced by VlangParser#valores.
	VisitValores(ctx *ValoresContext) interface{}

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

	// Visit a parse tree produced by VlangParser#incremento.
	VisitIncremento(ctx *IncrementoContext) interface{}

	// Visit a parse tree produced by VlangParser#decremento.
	VisitDecremento(ctx *DecrementoContext) interface{}
}
