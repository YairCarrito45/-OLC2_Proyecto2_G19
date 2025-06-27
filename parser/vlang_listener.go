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

	// EnterStructDeclaration is called when entering the structDeclaration production.
	EnterStructDeclaration(c *StructDeclarationContext)

	// EnterAtributos is called when entering the atributos production.
	EnterAtributos(c *AtributosContext)

	// EnterAtributo is called when entering the atributo production.
	EnterAtributo(c *AtributoContext)

	// EnterInicializadorStruct is called when entering the inicializadorStruct production.
	EnterInicializadorStruct(c *InicializadorStructContext)

	// EnterInicializadorCampo is called when entering the inicializadorCampo production.
	EnterInicializadorCampo(c *InicializadorCampoContext)

	// EnterFuncionStruct is called when entering the funcionStruct production.
	EnterFuncionStruct(c *FuncionStructContext)

	// EnterFuncion is called when entering the funcion production.
	EnterFuncion(c *FuncionContext)

	// EnterParametrosFormales is called when entering the parametrosFormales production.
	EnterParametrosFormales(c *ParametrosFormalesContext)

	// EnterParametroFormal is called when entering the parametroFormal production.
	EnterParametroFormal(c *ParametroFormalContext)

	// EnterTipoConReferencia is called when entering the tipoConReferencia production.
	EnterTipoConReferencia(c *TipoConReferenciaContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterSliceTipo is called when entering the sliceTipo production.
	EnterSliceTipo(c *SliceTipoContext)

	// EnterExpresionStatement is called when entering the expresionStatement production.
	EnterExpresionStatement(c *ExpresionStatementContext)

	// EnterSliceDobleAsignacionStmt is called when entering the sliceDobleAsignacionStmt production.
	EnterSliceDobleAsignacionStmt(c *SliceDobleAsignacionStmtContext)

	// EnterPrintStatement is called when entering the printStatement production.
	EnterPrintStatement(c *PrintStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterControlStatement is called when entering the controlStatement production.
	EnterControlStatement(c *ControlStatementContext)

	// EnterBloqueStatement is called when entering the bloqueStatement production.
	EnterBloqueStatement(c *BloqueStatementContext)

	// EnterIf_context is called when entering the if_context production.
	EnterIf_context(c *If_contextContext)

	// EnterFor_context is called when entering the for_context production.
	EnterFor_context(c *For_contextContext)

	// EnterWhile_context is called when entering the while_context production.
	EnterWhile_context(c *While_contextContext)

	// EnterSwitch_context is called when entering the switch_context production.
	EnterSwitch_context(c *Switch_contextContext)

	// EnterIfDcl is called when entering the ifDcl production.
	EnterIfDcl(c *IfDclContext)

	// EnterElseifDcl is called when entering the elseifDcl production.
	EnterElseifDcl(c *ElseifDclContext)

	// EnterElseDcl is called when entering the elseDcl production.
	EnterElseDcl(c *ElseDclContext)

	// EnterSwitchDcl is called when entering the switchDcl production.
	EnterSwitchDcl(c *SwitchDclContext)

	// EnterCaseBlock is called when entering the caseBlock production.
	EnterCaseBlock(c *CaseBlockContext)

	// EnterDefaultBlock is called when entering the defaultBlock production.
	EnterDefaultBlock(c *DefaultBlockContext)

	// EnterForCondicional is called when entering the forCondicional production.
	EnterForCondicional(c *ForCondicionalContext)

	// EnterForClasico is called when entering the forClasico production.
	EnterForClasico(c *ForClasicoContext)

	// EnterForForeach is called when entering the forForeach production.
	EnterForForeach(c *ForForeachContext)

	// EnterWhileDcl is called when entering the whileDcl production.
	EnterWhileDcl(c *WhileDclContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterAsignacionAtributo is called when entering the asignacionAtributo production.
	EnterAsignacionAtributo(c *AsignacionAtributoContext)

	// EnterMultdivmod is called when entering the multdivmod production.
	EnterMultdivmod(c *MultdivmodContext)

	// EnterLenExpr is called when entering the lenExpr production.
	EnterLenExpr(c *LenExprContext)

	// EnterSliceDobleAcceso is called when entering the sliceDobleAcceso production.
	EnterSliceDobleAcceso(c *SliceDobleAccesoContext)

	// EnterAsignacionResta is called when entering the asignacionResta production.
	EnterAsignacionResta(c *AsignacionRestaContext)

	// EnterAsignacionfor is called when entering the asignacionfor production.
	EnterAsignacionfor(c *AsignacionforContext)

	// EnterLiteralSliceExpr is called when entering the literalSliceExpr production.
	EnterLiteralSliceExpr(c *LiteralSliceExprContext)

	// EnterUnario is called when entering the unario production.
	EnterUnario(c *UnarioContext)

	// EnterSliceAsignacion is called when entering the sliceAsignacion production.
	EnterSliceAsignacion(c *SliceAsignacionContext)

	// EnterParentesisexpre is called when entering the parentesisexpre production.
	EnterParentesisexpre(c *ParentesisexpreContext)

	// EnterJoinExpr is called when entering the joinExpr production.
	EnterJoinExpr(c *JoinExprContext)

	// EnterLlamadaMetodoStruct is called when entering the llamadaMetodoStruct production.
	EnterLlamadaMetodoStruct(c *LlamadaMetodoStructContext)

	// EnterSumres is called when entering the sumres production.
	EnterSumres(c *SumresContext)

	// EnterSliceCreacion is called when entering the sliceCreacion production.
	EnterSliceCreacion(c *SliceCreacionContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterAppendExpr is called when entering the appendExpr production.
	EnterAppendExpr(c *AppendExprContext)

	// EnterAsignacionCompuestaResta is called when entering the asignacionCompuestaResta production.
	EnterAsignacionCompuestaResta(c *AsignacionCompuestaRestaContext)

	// EnterValorexpr is called when entering the valorexpr production.
	EnterValorexpr(c *ValorexprContext)

	// EnterIgualdad is called when entering the igualdad production.
	EnterIgualdad(c *IgualdadContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterSliceDobleAsignacion is called when entering the sliceDobleAsignacion production.
	EnterSliceDobleAsignacion(c *SliceDobleAsignacionContext)

	// EnterSliceAcceso is called when entering the sliceAcceso production.
	EnterSliceAcceso(c *SliceAccesoContext)

	// EnterRelacionales is called when entering the relacionales production.
	EnterRelacionales(c *RelacionalesContext)

	// EnterLlamadaFuncion is called when entering the llamadaFuncion production.
	EnterLlamadaFuncion(c *LlamadaFuncionContext)

	// EnterAccesoAtributo is called when entering the accesoAtributo production.
	EnterAccesoAtributo(c *AccesoAtributoContext)

	// EnterInstanciaStruct is called when entering the instanciaStruct production.
	EnterInstanciaStruct(c *InstanciaStructContext)

	// EnterAsignacionCompuestaSuma is called when entering the asignacionCompuestaSuma production.
	EnterAsignacionCompuestaSuma(c *AsignacionCompuestaSumaContext)

	// EnterIndexOfExpr is called when entering the indexOfExpr production.
	EnterIndexOfExpr(c *IndexOfExprContext)

	// EnterAsignacionSuma is called when entering the asignacionSuma production.
	EnterAsignacionSuma(c *AsignacionSumaContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

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

	// EnterValorNulo is called when entering the valorNulo production.
	EnterValorNulo(c *ValorNuloContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitDeclaraciones is called when exiting the declaraciones production.
	ExitDeclaraciones(c *DeclaracionesContext)

	// ExitStructDeclaration is called when exiting the structDeclaration production.
	ExitStructDeclaration(c *StructDeclarationContext)

	// ExitAtributos is called when exiting the atributos production.
	ExitAtributos(c *AtributosContext)

	// ExitAtributo is called when exiting the atributo production.
	ExitAtributo(c *AtributoContext)

	// ExitInicializadorStruct is called when exiting the inicializadorStruct production.
	ExitInicializadorStruct(c *InicializadorStructContext)

	// ExitInicializadorCampo is called when exiting the inicializadorCampo production.
	ExitInicializadorCampo(c *InicializadorCampoContext)

	// ExitFuncionStruct is called when exiting the funcionStruct production.
	ExitFuncionStruct(c *FuncionStructContext)

	// ExitFuncion is called when exiting the funcion production.
	ExitFuncion(c *FuncionContext)

	// ExitParametrosFormales is called when exiting the parametrosFormales production.
	ExitParametrosFormales(c *ParametrosFormalesContext)

	// ExitParametroFormal is called when exiting the parametroFormal production.
	ExitParametroFormal(c *ParametroFormalContext)

	// ExitTipoConReferencia is called when exiting the tipoConReferencia production.
	ExitTipoConReferencia(c *TipoConReferenciaContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitSliceTipo is called when exiting the sliceTipo production.
	ExitSliceTipo(c *SliceTipoContext)

	// ExitExpresionStatement is called when exiting the expresionStatement production.
	ExitExpresionStatement(c *ExpresionStatementContext)

	// ExitSliceDobleAsignacionStmt is called when exiting the sliceDobleAsignacionStmt production.
	ExitSliceDobleAsignacionStmt(c *SliceDobleAsignacionStmtContext)

	// ExitPrintStatement is called when exiting the printStatement production.
	ExitPrintStatement(c *PrintStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitControlStatement is called when exiting the controlStatement production.
	ExitControlStatement(c *ControlStatementContext)

	// ExitBloqueStatement is called when exiting the bloqueStatement production.
	ExitBloqueStatement(c *BloqueStatementContext)

	// ExitIf_context is called when exiting the if_context production.
	ExitIf_context(c *If_contextContext)

	// ExitFor_context is called when exiting the for_context production.
	ExitFor_context(c *For_contextContext)

	// ExitWhile_context is called when exiting the while_context production.
	ExitWhile_context(c *While_contextContext)

	// ExitSwitch_context is called when exiting the switch_context production.
	ExitSwitch_context(c *Switch_contextContext)

	// ExitIfDcl is called when exiting the ifDcl production.
	ExitIfDcl(c *IfDclContext)

	// ExitElseifDcl is called when exiting the elseifDcl production.
	ExitElseifDcl(c *ElseifDclContext)

	// ExitElseDcl is called when exiting the elseDcl production.
	ExitElseDcl(c *ElseDclContext)

	// ExitSwitchDcl is called when exiting the switchDcl production.
	ExitSwitchDcl(c *SwitchDclContext)

	// ExitCaseBlock is called when exiting the caseBlock production.
	ExitCaseBlock(c *CaseBlockContext)

	// ExitDefaultBlock is called when exiting the defaultBlock production.
	ExitDefaultBlock(c *DefaultBlockContext)

	// ExitForCondicional is called when exiting the forCondicional production.
	ExitForCondicional(c *ForCondicionalContext)

	// ExitForClasico is called when exiting the forClasico production.
	ExitForClasico(c *ForClasicoContext)

	// ExitForForeach is called when exiting the forForeach production.
	ExitForForeach(c *ForForeachContext)

	// ExitWhileDcl is called when exiting the whileDcl production.
	ExitWhileDcl(c *WhileDclContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitAsignacionAtributo is called when exiting the asignacionAtributo production.
	ExitAsignacionAtributo(c *AsignacionAtributoContext)

	// ExitMultdivmod is called when exiting the multdivmod production.
	ExitMultdivmod(c *MultdivmodContext)

	// ExitLenExpr is called when exiting the lenExpr production.
	ExitLenExpr(c *LenExprContext)

	// ExitSliceDobleAcceso is called when exiting the sliceDobleAcceso production.
	ExitSliceDobleAcceso(c *SliceDobleAccesoContext)

	// ExitAsignacionResta is called when exiting the asignacionResta production.
	ExitAsignacionResta(c *AsignacionRestaContext)

	// ExitAsignacionfor is called when exiting the asignacionfor production.
	ExitAsignacionfor(c *AsignacionforContext)

	// ExitLiteralSliceExpr is called when exiting the literalSliceExpr production.
	ExitLiteralSliceExpr(c *LiteralSliceExprContext)

	// ExitUnario is called when exiting the unario production.
	ExitUnario(c *UnarioContext)

	// ExitSliceAsignacion is called when exiting the sliceAsignacion production.
	ExitSliceAsignacion(c *SliceAsignacionContext)

	// ExitParentesisexpre is called when exiting the parentesisexpre production.
	ExitParentesisexpre(c *ParentesisexpreContext)

	// ExitJoinExpr is called when exiting the joinExpr production.
	ExitJoinExpr(c *JoinExprContext)

	// ExitLlamadaMetodoStruct is called when exiting the llamadaMetodoStruct production.
	ExitLlamadaMetodoStruct(c *LlamadaMetodoStructContext)

	// ExitSumres is called when exiting the sumres production.
	ExitSumres(c *SumresContext)

	// ExitSliceCreacion is called when exiting the sliceCreacion production.
	ExitSliceCreacion(c *SliceCreacionContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitAppendExpr is called when exiting the appendExpr production.
	ExitAppendExpr(c *AppendExprContext)

	// ExitAsignacionCompuestaResta is called when exiting the asignacionCompuestaResta production.
	ExitAsignacionCompuestaResta(c *AsignacionCompuestaRestaContext)

	// ExitValorexpr is called when exiting the valorexpr production.
	ExitValorexpr(c *ValorexprContext)

	// ExitIgualdad is called when exiting the igualdad production.
	ExitIgualdad(c *IgualdadContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitSliceDobleAsignacion is called when exiting the sliceDobleAsignacion production.
	ExitSliceDobleAsignacion(c *SliceDobleAsignacionContext)

	// ExitSliceAcceso is called when exiting the sliceAcceso production.
	ExitSliceAcceso(c *SliceAccesoContext)

	// ExitRelacionales is called when exiting the relacionales production.
	ExitRelacionales(c *RelacionalesContext)

	// ExitLlamadaFuncion is called when exiting the llamadaFuncion production.
	ExitLlamadaFuncion(c *LlamadaFuncionContext)

	// ExitAccesoAtributo is called when exiting the accesoAtributo production.
	ExitAccesoAtributo(c *AccesoAtributoContext)

	// ExitInstanciaStruct is called when exiting the instanciaStruct production.
	ExitInstanciaStruct(c *InstanciaStructContext)

	// ExitAsignacionCompuestaSuma is called when exiting the asignacionCompuestaSuma production.
	ExitAsignacionCompuestaSuma(c *AsignacionCompuestaSumaContext)

	// ExitIndexOfExpr is called when exiting the indexOfExpr production.
	ExitIndexOfExpr(c *IndexOfExprContext)

	// ExitAsignacionSuma is called when exiting the asignacionSuma production.
	ExitAsignacionSuma(c *AsignacionSumaContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

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

	// ExitValorNulo is called when exiting the valorNulo production.
	ExitValorNulo(c *ValorNuloContext)
}
