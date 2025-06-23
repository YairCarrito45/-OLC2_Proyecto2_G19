// Generated from /home/harold/archivos/OLC2_Proyecto1_202204578/parser/Vlang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link VlangParser}.
 */
public interface VlangListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link VlangParser#programa}.
	 * @param ctx the parse tree
	 */
	void enterPrograma(VlangParser.ProgramaContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#programa}.
	 * @param ctx the parse tree
	 */
	void exitPrograma(VlangParser.ProgramaContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#declaraciones}.
	 * @param ctx the parse tree
	 */
	void enterDeclaraciones(VlangParser.DeclaracionesContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#declaraciones}.
	 * @param ctx the parse tree
	 */
	void exitDeclaraciones(VlangParser.DeclaracionesContext ctx);
	/**
	 * Enter a parse tree produced by the {@code variableDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void enterVariableDeclaration(VlangParser.VariableDeclarationContext ctx);
	/**
	 * Exit a parse tree produced by the {@code variableDeclaration}
	 * labeled alternative in {@link VlangParser#varDcl}.
	 * @param ctx the parse tree
	 */
	void exitVariableDeclaration(VlangParser.VariableDeclarationContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expresionStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterExpresionStatement(VlangParser.ExpresionStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expresionStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitExpresionStatement(VlangParser.ExpresionStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code printStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterPrintStatement(VlangParser.PrintStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code printStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitPrintStatement(VlangParser.PrintStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code controlStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void enterControlStatement(VlangParser.ControlStatementContext ctx);
	/**
	 * Exit a parse tree produced by the {@code controlStatement}
	 * labeled alternative in {@link VlangParser#stmt}.
	 * @param ctx the parse tree
	 */
	void exitControlStatement(VlangParser.ControlStatementContext ctx);
	/**
	 * Enter a parse tree produced by the {@code if_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void enterIf_context(VlangParser.If_contextContext ctx);
	/**
	 * Exit a parse tree produced by the {@code if_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void exitIf_context(VlangParser.If_contextContext ctx);
	/**
	 * Enter a parse tree produced by the {@code for_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void enterFor_context(VlangParser.For_contextContext ctx);
	/**
	 * Exit a parse tree produced by the {@code for_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void exitFor_context(VlangParser.For_contextContext ctx);
	/**
	 * Enter a parse tree produced by the {@code while_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void enterWhile_context(VlangParser.While_contextContext ctx);
	/**
	 * Exit a parse tree produced by the {@code while_context}
	 * labeled alternative in {@link VlangParser#sentencias_control}.
	 * @param ctx the parse tree
	 */
	void exitWhile_context(VlangParser.While_contextContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#ifDcl}.
	 * @param ctx the parse tree
	 */
	void enterIfDcl(VlangParser.IfDclContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#ifDcl}.
	 * @param ctx the parse tree
	 */
	void exitIfDcl(VlangParser.IfDclContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#forDcl}.
	 * @param ctx the parse tree
	 */
	void enterForDcl(VlangParser.ForDclContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#forDcl}.
	 * @param ctx the parse tree
	 */
	void exitForDcl(VlangParser.ForDclContext ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#whileDcl}.
	 * @param ctx the parse tree
	 */
	void enterWhileDcl(VlangParser.WhileDclContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#whileDcl}.
	 * @param ctx the parse tree
	 */
	void exitWhileDcl(VlangParser.WhileDclContext ctx);
	/**
	 * Enter a parse tree produced by the {@code multdivmod}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterMultdivmod(VlangParser.MultdivmodContext ctx);
	/**
	 * Exit a parse tree produced by the {@code multdivmod}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitMultdivmod(VlangParser.MultdivmodContext ctx);
	/**
	 * Enter a parse tree produced by the {@code or}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterOr(VlangParser.OrContext ctx);
	/**
	 * Exit a parse tree produced by the {@code or}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitOr(VlangParser.OrContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorexpr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterValorexpr(VlangParser.ValorexprContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorexpr}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitValorexpr(VlangParser.ValorexprContext ctx);
	/**
	 * Enter a parse tree produced by the {@code igualdad}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterIgualdad(VlangParser.IgualdadContext ctx);
	/**
	 * Exit a parse tree produced by the {@code igualdad}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitIgualdad(VlangParser.IgualdadContext ctx);
	/**
	 * Enter a parse tree produced by the {@code asignacionfor}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionfor(VlangParser.AsignacionforContext ctx);
	/**
	 * Exit a parse tree produced by the {@code asignacionfor}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionfor(VlangParser.AsignacionforContext ctx);
	/**
	 * Enter a parse tree produced by the {@code incremento}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterIncremento(VlangParser.IncrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code incremento}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitIncremento(VlangParser.IncrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expdotexp}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpdotexp(VlangParser.ExpdotexpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code expdotexp}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpdotexp(VlangParser.ExpdotexpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code relacionales}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterRelacionales(VlangParser.RelacionalesContext ctx);
	/**
	 * Exit a parse tree produced by the {@code relacionales}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitRelacionales(VlangParser.RelacionalesContext ctx);
	/**
	 * Enter a parse tree produced by the {@code decremento}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterDecremento(VlangParser.DecrementoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code decremento}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitDecremento(VlangParser.DecrementoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code corchetesexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterCorchetesexpre(VlangParser.CorchetesexpreContext ctx);
	/**
	 * Exit a parse tree produced by the {@code corchetesexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitCorchetesexpre(VlangParser.CorchetesexpreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code unario}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterUnario(VlangParser.UnarioContext ctx);
	/**
	 * Exit a parse tree produced by the {@code unario}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitUnario(VlangParser.UnarioContext ctx);
	/**
	 * Enter a parse tree produced by the {@code parentesisexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterParentesisexpre(VlangParser.ParentesisexpreContext ctx);
	/**
	 * Exit a parse tree produced by the {@code parentesisexpre}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitParentesisexpre(VlangParser.ParentesisexpreContext ctx);
	/**
	 * Enter a parse tree produced by the {@code sumres}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterSumres(VlangParser.SumresContext ctx);
	/**
	 * Exit a parse tree produced by the {@code sumres}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitSumres(VlangParser.SumresContext ctx);
	/**
	 * Enter a parse tree produced by the {@code id}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterId(VlangParser.IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code id}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitId(VlangParser.IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code expdotexp1}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void enterExpdotexp1(VlangParser.Expdotexp1Context ctx);
	/**
	 * Exit a parse tree produced by the {@code expdotexp1}
	 * labeled alternative in {@link VlangParser#expresion}.
	 * @param ctx the parse tree
	 */
	void exitExpdotexp1(VlangParser.Expdotexp1Context ctx);
	/**
	 * Enter a parse tree produced by {@link VlangParser#parametros}.
	 * @param ctx the parse tree
	 */
	void enterParametros(VlangParser.ParametrosContext ctx);
	/**
	 * Exit a parse tree produced by {@link VlangParser#parametros}.
	 * @param ctx the parse tree
	 */
	void exitParametros(VlangParser.ParametrosContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorEntero}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorEntero(VlangParser.ValorEnteroContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorEntero}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorEntero(VlangParser.ValorEnteroContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorDecimal}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorDecimal(VlangParser.ValorDecimalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorDecimal}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorDecimal(VlangParser.ValorDecimalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorCadena}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorCadena(VlangParser.ValorCadenaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorCadena}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorCadena(VlangParser.ValorCadenaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorBooleano}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorBooleano(VlangParser.ValorBooleanoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorBooleano}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorBooleano(VlangParser.ValorBooleanoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code valorCaracter}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void enterValorCaracter(VlangParser.ValorCaracterContext ctx);
	/**
	 * Exit a parse tree produced by the {@code valorCaracter}
	 * labeled alternative in {@link VlangParser#valor}.
	 * @param ctx the parse tree
	 */
	void exitValorCaracter(VlangParser.ValorCaracterContext ctx);
}