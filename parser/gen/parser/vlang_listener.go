// Code generated from parser/Vlang.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // Vlang
import "github.com/antlr4-go/antlr/v4"

// VlangListener is a complete listener for a parse tree produced by VlangParser.
type VlangListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterDeclaraciones is called when entering the declaraciones production.
	EnterDeclaraciones(c *DeclaracionesContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterExpresionStatement is called when entering the expresionStatement production.
	EnterExpresionStatement(c *ExpresionStatementContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterControlStatement is called when entering the controlStatement production.
	EnterControlStatement(c *ControlStatementContext)

	// EnterIf_context is called when entering the if_context production.
	EnterIf_context(c *If_contextContext)

	// EnterFor_context is called when entering the for_context production.
	EnterFor_context(c *For_contextContext)

	// EnterWhile_context is called when entering the while_context production.
	EnterWhile_context(c *While_contextContext)

	// EnterIfDcl is called when entering the ifDcl production.
	EnterIfDcl(c *IfDclContext)

	// EnterForDcl is called when entering the forDcl production.
	EnterForDcl(c *ForDclContext)

	// EnterWhileDcl is called when entering the whileDcl production.
	EnterWhileDcl(c *WhileDclContext)

	// EnterMultdivmod is called when entering the multdivmod production.
	EnterMultdivmod(c *MultdivmodContext)

	// EnterIncredecr is called when entering the incredecr production.
	EnterIncredecr(c *IncredecrContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterValorexpr is called when entering the valorexpr production.
	EnterValorexpr(c *ValorexprContext)

	// EnterIgualdad is called when entering the igualdad production.
	EnterIgualdad(c *IgualdadContext)

	// EnterAsignacionfor is called when entering the asignacionfor production.
	EnterAsignacionfor(c *AsignacionforContext)

	// EnterExpdotexp is called when entering the expdotexp production.
	EnterExpdotexp(c *ExpdotexpContext)

	// EnterRelacionales is called when entering the relacionales production.
	EnterRelacionales(c *RelacionalesContext)

	// EnterCorchetesexpre is called when entering the corchetesexpre production.
	EnterCorchetesexpre(c *CorchetesexpreContext)

	// EnterUnario is called when entering the unario production.
	EnterUnario(c *UnarioContext)

	// EnterParentesisexpre is called when entering the parentesisexpre production.
	EnterParentesisexpre(c *ParentesisexpreContext)

	// EnterSumres is called when entering the sumres production.
	EnterSumres(c *SumresContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterExpdotexp1 is called when entering the expdotexp1 production.
	EnterExpdotexp1(c *Expdotexp1Context)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterValores is called when entering the valores production.
	EnterValores(c *ValoresContext)

	// EnterValorEntero is called when entering the valorEntero production.
	EnterValorEntero(c *ValorEnteroContext)

	// EnterValorDecimal is called when entering the valorDecimal production.
	EnterValorDecimal(c *ValorDecimalContext)

	// EnterValorCadena is called when entering the valorCadena production.
	EnterValorCadena(c *ValorCadenaContext)

	// EnterValorBooleano is called when entering the valorBooleano production.
	EnterValorBooleano(c *ValorBooleanoContext)

	// EnterValorCaracter is called when entering the valorCaracter production.
	EnterValorCaracter(c *ValorCaracterContext)

	// EnterIncremento is called when entering the incremento production.
	EnterIncremento(c *IncrementoContext)

	// EnterDecremento is called when entering the decremento production.
	EnterDecremento(c *DecrementoContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitDeclaraciones is called when exiting the declaraciones production.
	ExitDeclaraciones(c *DeclaracionesContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitExpresionStatement is called when exiting the expresionStatement production.
	ExitExpresionStatement(c *ExpresionStatementContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitControlStatement is called when exiting the controlStatement production.
	ExitControlStatement(c *ControlStatementContext)

	// ExitIf_context is called when exiting the if_context production.
	ExitIf_context(c *If_contextContext)

	// ExitFor_context is called when exiting the for_context production.
	ExitFor_context(c *For_contextContext)

	// ExitWhile_context is called when exiting the while_context production.
	ExitWhile_context(c *While_contextContext)

	// ExitIfDcl is called when exiting the ifDcl production.
	ExitIfDcl(c *IfDclContext)

	// ExitForDcl is called when exiting the forDcl production.
	ExitForDcl(c *ForDclContext)

	// ExitWhileDcl is called when exiting the whileDcl production.
	ExitWhileDcl(c *WhileDclContext)

	// ExitMultdivmod is called when exiting the multdivmod production.
	ExitMultdivmod(c *MultdivmodContext)

	// ExitIncredecr is called when exiting the incredecr production.
	ExitIncredecr(c *IncredecrContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitValorexpr is called when exiting the valorexpr production.
	ExitValorexpr(c *ValorexprContext)

	// ExitIgualdad is called when exiting the igualdad production.
	ExitIgualdad(c *IgualdadContext)

	// ExitAsignacionfor is called when exiting the asignacionfor production.
	ExitAsignacionfor(c *AsignacionforContext)

	// ExitExpdotexp is called when exiting the expdotexp production.
	ExitExpdotexp(c *ExpdotexpContext)

	// ExitRelacionales is called when exiting the relacionales production.
	ExitRelacionales(c *RelacionalesContext)

	// ExitCorchetesexpre is called when exiting the corchetesexpre production.
	ExitCorchetesexpre(c *CorchetesexpreContext)

	// ExitUnario is called when exiting the unario production.
	ExitUnario(c *UnarioContext)

	// ExitParentesisexpre is called when exiting the parentesisexpre production.
	ExitParentesisexpre(c *ParentesisexpreContext)

	// ExitSumres is called when exiting the sumres production.
	ExitSumres(c *SumresContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitExpdotexp1 is called when exiting the expdotexp1 production.
	ExitExpdotexp1(c *Expdotexp1Context)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitValores is called when exiting the valores production.
	ExitValores(c *ValoresContext)

	// ExitValorEntero is called when exiting the valorEntero production.
	ExitValorEntero(c *ValorEnteroContext)

	// ExitValorDecimal is called when exiting the valorDecimal production.
	ExitValorDecimal(c *ValorDecimalContext)

	// ExitValorCadena is called when exiting the valorCadena production.
	ExitValorCadena(c *ValorCadenaContext)

	// ExitValorBooleano is called when exiting the valorBooleano production.
	ExitValorBooleano(c *ValorBooleanoContext)

	// ExitValorCaracter is called when exiting the valorCaracter production.
	ExitValorCaracter(c *ValorCaracterContext)

	// ExitIncremento is called when exiting the incremento production.
	ExitIncremento(c *IncrementoContext)

	// ExitDecremento is called when exiting the decremento production.
	ExitDecremento(c *DecrementoContext)
}
