// Generated from /home/raul/Documentos/OLC2-Proyecto2/Gramar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class GramarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, WS=17, 
		COMS=18, COMM=19, INT=20, DOUBLE=21, BOOLEAN=22;
	public static final int
		RULE_s = 0, RULE_block = 1, RULE_productions = 2, RULE_prin = 3, RULE_swhile = 4, 
		RULE_expr = 5;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "block", "productions", "prin", "swhile", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'print'", "'('", "')'", "'while'", "'!'", "'-'", "'/'", "'*'", 
			"'%'", "'+'", "'<'", "'>'", "'<='", "'>='", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, "WS", "COMS", "COMM", "INT", "DOUBLE", 
			"BOOLEAN"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Gramar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GramarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(GramarParser.EOF, 0); }
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(12);
			block();
			setState(13);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public List<ProductionsContext> productions() {
			return getRuleContexts(ProductionsContext.class);
		}
		public ProductionsContext productions(int i) {
			return getRuleContext(ProductionsContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(18);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0 || _la==T__3) {
				{
				{
				setState(15);
				productions();
				}
				}
				setState(20);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProductionsContext extends ParserRuleContext {
		public PrinContext prin() {
			return getRuleContext(PrinContext.class,0);
		}
		public SwhileContext swhile() {
			return getRuleContext(SwhileContext.class,0);
		}
		public ProductionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_productions; }
	}

	public final ProductionsContext productions() throws RecognitionException {
		ProductionsContext _localctx = new ProductionsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_productions);
		try {
			setState(23);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(21);
				prin();
				}
				break;
			case T__3:
				enterOuterAlt(_localctx, 2);
				{
				setState(22);
				swhile();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrinContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public PrinContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prin; }
	}

	public final PrinContext prin() throws RecognitionException {
		PrinContext _localctx = new PrinContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_prin);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			match(T__0);
			setState(26);
			match(T__1);
			setState(27);
			expr(0);
			setState(28);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SwhileContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public SwhileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_swhile; }
	}

	public final SwhileContext swhile() throws RecognitionException {
		SwhileContext _localctx = new SwhileContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_swhile);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			match(T__3);
			setState(31);
			match(T__1);
			setState(32);
			expr(0);
			setState(33);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ExprContext {
		public TerminalNode INT() { return getToken(GramarParser.INT, 0); }
		public TerminalNode DOUBLE() { return getToken(GramarParser.DOUBLE, 0); }
		public TerminalNode BOOLEAN() { return getToken(GramarParser.BOOLEAN, 0); }
		public LiteralContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AritmeticContext extends ExprContext {
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public AritmeticContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__4:
			case T__5:
				{
				_localctx = new AritmeticContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(36);
				((AritmeticContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__4 || _la==T__5) ) {
					((AritmeticContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(37);
				((AritmeticContext)_localctx).right = expr(9);
				}
				break;
			case INT:
				{
				_localctx = new LiteralContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(38);
				match(INT);
				}
				break;
			case DOUBLE:
				{
				_localctx = new LiteralContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(39);
				match(DOUBLE);
				}
				break;
			case BOOLEAN:
				{
				_localctx = new LiteralContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(40);
				match(BOOLEAN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(60);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(58);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
					case 1:
						{
						_localctx = new AritmeticContext(new ExprContext(_parentctx, _parentState));
						((AritmeticContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(43);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(44);
						((AritmeticContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 896L) != 0)) ) {
							((AritmeticContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(45);
						((AritmeticContext)_localctx).right = expr(9);
						}
						break;
					case 2:
						{
						_localctx = new AritmeticContext(new ExprContext(_parentctx, _parentState));
						((AritmeticContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(46);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(47);
						((AritmeticContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__5 || _la==T__9) ) {
							((AritmeticContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(48);
						((AritmeticContext)_localctx).right = expr(8);
						}
						break;
					case 3:
						{
						_localctx = new AritmeticContext(new ExprContext(_parentctx, _parentState));
						((AritmeticContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(49);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(50);
						((AritmeticContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 30720L) != 0)) ) {
							((AritmeticContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(51);
						((AritmeticContext)_localctx).right = expr(7);
						}
						break;
					case 4:
						{
						_localctx = new AritmeticContext(new ExprContext(_parentctx, _parentState));
						((AritmeticContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(52);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(53);
						((AritmeticContext)_localctx).op = match(T__14);
						setState(54);
						((AritmeticContext)_localctx).right = expr(6);
						}
						break;
					case 5:
						{
						_localctx = new AritmeticContext(new ExprContext(_parentctx, _parentState));
						((AritmeticContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(55);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(56);
						((AritmeticContext)_localctx).op = match(T__15);
						setState(57);
						((AritmeticContext)_localctx).right = expr(5);
						}
						break;
					}
					} 
				}
				setState(62);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 5:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 8);
		case 1:
			return precpred(_ctx, 7);
		case 2:
			return precpred(_ctx, 6);
		case 3:
			return precpred(_ctx, 5);
		case 4:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u0016@\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0005"+
		"\u0001\u0011\b\u0001\n\u0001\f\u0001\u0014\t\u0001\u0001\u0002\u0001\u0002"+
		"\u0003\u0002\u0018\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0003\u0005*\b\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005"+
		";\b\u0005\n\u0005\f\u0005>\t\u0005\u0001\u0005\u0000\u0001\n\u0006\u0000"+
		"\u0002\u0004\u0006\b\n\u0000\u0004\u0001\u0000\u0005\u0006\u0001\u0000"+
		"\u0007\t\u0002\u0000\u0006\u0006\n\n\u0001\u0000\u000b\u000eC\u0000\f"+
		"\u0001\u0000\u0000\u0000\u0002\u0012\u0001\u0000\u0000\u0000\u0004\u0017"+
		"\u0001\u0000\u0000\u0000\u0006\u0019\u0001\u0000\u0000\u0000\b\u001e\u0001"+
		"\u0000\u0000\u0000\n)\u0001\u0000\u0000\u0000\f\r\u0003\u0002\u0001\u0000"+
		"\r\u000e\u0005\u0000\u0000\u0001\u000e\u0001\u0001\u0000\u0000\u0000\u000f"+
		"\u0011\u0003\u0004\u0002\u0000\u0010\u000f\u0001\u0000\u0000\u0000\u0011"+
		"\u0014\u0001\u0000\u0000\u0000\u0012\u0010\u0001\u0000\u0000\u0000\u0012"+
		"\u0013\u0001\u0000\u0000\u0000\u0013\u0003\u0001\u0000\u0000\u0000\u0014"+
		"\u0012\u0001\u0000\u0000\u0000\u0015\u0018\u0003\u0006\u0003\u0000\u0016"+
		"\u0018\u0003\b\u0004\u0000\u0017\u0015\u0001\u0000\u0000\u0000\u0017\u0016"+
		"\u0001\u0000\u0000\u0000\u0018\u0005\u0001\u0000\u0000\u0000\u0019\u001a"+
		"\u0005\u0001\u0000\u0000\u001a\u001b\u0005\u0002\u0000\u0000\u001b\u001c"+
		"\u0003\n\u0005\u0000\u001c\u001d\u0005\u0003\u0000\u0000\u001d\u0007\u0001"+
		"\u0000\u0000\u0000\u001e\u001f\u0005\u0004\u0000\u0000\u001f \u0005\u0002"+
		"\u0000\u0000 !\u0003\n\u0005\u0000!\"\u0005\u0003\u0000\u0000\"\t\u0001"+
		"\u0000\u0000\u0000#$\u0006\u0005\uffff\uffff\u0000$%\u0007\u0000\u0000"+
		"\u0000%*\u0003\n\u0005\t&*\u0005\u0014\u0000\u0000\'*\u0005\u0015\u0000"+
		"\u0000(*\u0005\u0016\u0000\u0000)#\u0001\u0000\u0000\u0000)&\u0001\u0000"+
		"\u0000\u0000)\'\u0001\u0000\u0000\u0000)(\u0001\u0000\u0000\u0000*<\u0001"+
		"\u0000\u0000\u0000+,\n\b\u0000\u0000,-\u0007\u0001\u0000\u0000-;\u0003"+
		"\n\u0005\t./\n\u0007\u0000\u0000/0\u0007\u0002\u0000\u00000;\u0003\n\u0005"+
		"\b12\n\u0006\u0000\u000023\u0007\u0003\u0000\u00003;\u0003\n\u0005\u0007"+
		"45\n\u0005\u0000\u000056\u0005\u000f\u0000\u00006;\u0003\n\u0005\u0006"+
		"78\n\u0004\u0000\u000089\u0005\u0010\u0000\u00009;\u0003\n\u0005\u0005"+
		":+\u0001\u0000\u0000\u0000:.\u0001\u0000\u0000\u0000:1\u0001\u0000\u0000"+
		"\u0000:4\u0001\u0000\u0000\u0000:7\u0001\u0000\u0000\u0000;>\u0001\u0000"+
		"\u0000\u0000<:\u0001\u0000\u0000\u0000<=\u0001\u0000\u0000\u0000=\u000b"+
		"\u0001\u0000\u0000\u0000><\u0001\u0000\u0000\u0000\u0005\u0012\u0017)"+
		":<";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}