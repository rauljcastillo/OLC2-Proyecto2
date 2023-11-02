// Generated from /home/raul/Documentos/OLC2-Proyecto2/Gramar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link GramarParser}.
 */
public interface GramarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link GramarParser#s}.
	 * @param ctx the parse tree
	 */
	void enterS(GramarParser.SContext ctx);
	/**
	 * Exit a parse tree produced by {@link GramarParser#s}.
	 * @param ctx the parse tree
	 */
	void exitS(GramarParser.SContext ctx);
	/**
	 * Enter a parse tree produced by {@link GramarParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(GramarParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link GramarParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(GramarParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link GramarParser#productions}.
	 * @param ctx the parse tree
	 */
	void enterProductions(GramarParser.ProductionsContext ctx);
	/**
	 * Exit a parse tree produced by {@link GramarParser#productions}.
	 * @param ctx the parse tree
	 */
	void exitProductions(GramarParser.ProductionsContext ctx);
	/**
	 * Enter a parse tree produced by {@link GramarParser#prin}.
	 * @param ctx the parse tree
	 */
	void enterPrin(GramarParser.PrinContext ctx);
	/**
	 * Exit a parse tree produced by {@link GramarParser#prin}.
	 * @param ctx the parse tree
	 */
	void exitPrin(GramarParser.PrinContext ctx);
	/**
	 * Enter a parse tree produced by {@link GramarParser#swhile}.
	 * @param ctx the parse tree
	 */
	void enterSwhile(GramarParser.SwhileContext ctx);
	/**
	 * Exit a parse tree produced by {@link GramarParser#swhile}.
	 * @param ctx the parse tree
	 */
	void exitSwhile(GramarParser.SwhileContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Literal}
	 * labeled alternative in {@link GramarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterLiteral(GramarParser.LiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Literal}
	 * labeled alternative in {@link GramarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitLiteral(GramarParser.LiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code aritmetic}
	 * labeled alternative in {@link GramarParser#expr}.
	 * @param ctx the parse tree
	 */
	void enterAritmetic(GramarParser.AritmeticContext ctx);
	/**
	 * Exit a parse tree produced by the {@code aritmetic}
	 * labeled alternative in {@link GramarParser#expr}.
	 * @param ctx the parse tree
	 */
	void exitAritmetic(GramarParser.AritmeticContext ctx);
}