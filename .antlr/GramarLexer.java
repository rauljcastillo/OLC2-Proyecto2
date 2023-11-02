// Generated from /home/raul/Documentos/OLC2-Proyecto2/Gramar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class GramarLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, WS=17, 
		COMS=18, COMM=19, INT=20, DOUBLE=21, BOOLEAN=22;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "WS", "COMS", 
			"COMM", "INT", "DOUBLE", "BOOLEAN"
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


	public GramarLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Gramar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0016\u0091\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014"+
		"\u0002\u0015\u0007\u0015\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001"+
		"\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0005\u0011b\b\u0011\n\u0011\f\u0011e\t\u0011"+
		"\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0005\u0012m\b\u0012\n\u0012\f\u0012p\t\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0004\u0013x\b\u0013"+
		"\u000b\u0013\f\u0013y\u0001\u0014\u0004\u0014}\b\u0014\u000b\u0014\f\u0014"+
		"~\u0001\u0014\u0001\u0014\u0004\u0014\u0083\b\u0014\u000b\u0014\f\u0014"+
		"\u0084\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u0090\b\u0015\u0001"+
		"n\u0000\u0016\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005"+
		"\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019"+
		"\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011#\u0012%\u0013\'\u0014)\u0015"+
		"+\u0016\u0001\u0000\u0003\u0002\u0000\t\n  \u0001\u0000\n\n\u0001\u0000"+
		"09\u0096\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000"+
		"\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000"+
		"\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000"+
		"\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000"+
		"\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000"+
		"\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000"+
		"\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000"+
		"\u001d\u0001\u0000\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000"+
		"!\u0001\u0000\u0000\u0000\u0000#\u0001\u0000\u0000\u0000\u0000%\u0001"+
		"\u0000\u0000\u0000\u0000\'\u0001\u0000\u0000\u0000\u0000)\u0001\u0000"+
		"\u0000\u0000\u0000+\u0001\u0000\u0000\u0000\u0001-\u0001\u0000\u0000\u0000"+
		"\u00033\u0001\u0000\u0000\u0000\u00055\u0001\u0000\u0000\u0000\u00077"+
		"\u0001\u0000\u0000\u0000\t=\u0001\u0000\u0000\u0000\u000b?\u0001\u0000"+
		"\u0000\u0000\rA\u0001\u0000\u0000\u0000\u000fC\u0001\u0000\u0000\u0000"+
		"\u0011E\u0001\u0000\u0000\u0000\u0013G\u0001\u0000\u0000\u0000\u0015I"+
		"\u0001\u0000\u0000\u0000\u0017K\u0001\u0000\u0000\u0000\u0019M\u0001\u0000"+
		"\u0000\u0000\u001bP\u0001\u0000\u0000\u0000\u001dS\u0001\u0000\u0000\u0000"+
		"\u001fV\u0001\u0000\u0000\u0000!Y\u0001\u0000\u0000\u0000#]\u0001\u0000"+
		"\u0000\u0000%h\u0001\u0000\u0000\u0000\'w\u0001\u0000\u0000\u0000)|\u0001"+
		"\u0000\u0000\u0000+\u008f\u0001\u0000\u0000\u0000-.\u0005p\u0000\u0000"+
		"./\u0005r\u0000\u0000/0\u0005i\u0000\u000001\u0005n\u0000\u000012\u0005"+
		"t\u0000\u00002\u0002\u0001\u0000\u0000\u000034\u0005(\u0000\u00004\u0004"+
		"\u0001\u0000\u0000\u000056\u0005)\u0000\u00006\u0006\u0001\u0000\u0000"+
		"\u000078\u0005w\u0000\u000089\u0005h\u0000\u00009:\u0005i\u0000\u0000"+
		":;\u0005l\u0000\u0000;<\u0005e\u0000\u0000<\b\u0001\u0000\u0000\u0000"+
		"=>\u0005!\u0000\u0000>\n\u0001\u0000\u0000\u0000?@\u0005-\u0000\u0000"+
		"@\f\u0001\u0000\u0000\u0000AB\u0005/\u0000\u0000B\u000e\u0001\u0000\u0000"+
		"\u0000CD\u0005*\u0000\u0000D\u0010\u0001\u0000\u0000\u0000EF\u0005%\u0000"+
		"\u0000F\u0012\u0001\u0000\u0000\u0000GH\u0005+\u0000\u0000H\u0014\u0001"+
		"\u0000\u0000\u0000IJ\u0005<\u0000\u0000J\u0016\u0001\u0000\u0000\u0000"+
		"KL\u0005>\u0000\u0000L\u0018\u0001\u0000\u0000\u0000MN\u0005<\u0000\u0000"+
		"NO\u0005=\u0000\u0000O\u001a\u0001\u0000\u0000\u0000PQ\u0005>\u0000\u0000"+
		"QR\u0005=\u0000\u0000R\u001c\u0001\u0000\u0000\u0000ST\u0005&\u0000\u0000"+
		"TU\u0005&\u0000\u0000U\u001e\u0001\u0000\u0000\u0000VW\u0005|\u0000\u0000"+
		"WX\u0005|\u0000\u0000X \u0001\u0000\u0000\u0000YZ\u0007\u0000\u0000\u0000"+
		"Z[\u0001\u0000\u0000\u0000[\\\u0006\u0010\u0000\u0000\\\"\u0001\u0000"+
		"\u0000\u0000]^\u0005/\u0000\u0000^_\u0005/\u0000\u0000_c\u0001\u0000\u0000"+
		"\u0000`b\b\u0001\u0000\u0000a`\u0001\u0000\u0000\u0000be\u0001\u0000\u0000"+
		"\u0000ca\u0001\u0000\u0000\u0000cd\u0001\u0000\u0000\u0000df\u0001\u0000"+
		"\u0000\u0000ec\u0001\u0000\u0000\u0000fg\u0006\u0011\u0000\u0000g$\u0001"+
		"\u0000\u0000\u0000hi\u0005/\u0000\u0000ij\u0005*\u0000\u0000jn\u0001\u0000"+
		"\u0000\u0000km\t\u0000\u0000\u0000lk\u0001\u0000\u0000\u0000mp\u0001\u0000"+
		"\u0000\u0000no\u0001\u0000\u0000\u0000nl\u0001\u0000\u0000\u0000oq\u0001"+
		"\u0000\u0000\u0000pn\u0001\u0000\u0000\u0000qr\u0005*\u0000\u0000rs\u0005"+
		"/\u0000\u0000st\u0001\u0000\u0000\u0000tu\u0006\u0012\u0000\u0000u&\u0001"+
		"\u0000\u0000\u0000vx\u0007\u0002\u0000\u0000wv\u0001\u0000\u0000\u0000"+
		"xy\u0001\u0000\u0000\u0000yw\u0001\u0000\u0000\u0000yz\u0001\u0000\u0000"+
		"\u0000z(\u0001\u0000\u0000\u0000{}\u0007\u0002\u0000\u0000|{\u0001\u0000"+
		"\u0000\u0000}~\u0001\u0000\u0000\u0000~|\u0001\u0000\u0000\u0000~\u007f"+
		"\u0001\u0000\u0000\u0000\u007f\u0080\u0001\u0000\u0000\u0000\u0080\u0082"+
		"\u0005.\u0000\u0000\u0081\u0083\u0007\u0002\u0000\u0000\u0082\u0081\u0001"+
		"\u0000\u0000\u0000\u0083\u0084\u0001\u0000\u0000\u0000\u0084\u0082\u0001"+
		"\u0000\u0000\u0000\u0084\u0085\u0001\u0000\u0000\u0000\u0085*\u0001\u0000"+
		"\u0000\u0000\u0086\u0087\u0005t\u0000\u0000\u0087\u0088\u0005r\u0000\u0000"+
		"\u0088\u0089\u0005u\u0000\u0000\u0089\u0090\u0005e\u0000\u0000\u008a\u008b"+
		"\u0005f\u0000\u0000\u008b\u008c\u0005a\u0000\u0000\u008c\u008d\u0005l"+
		"\u0000\u0000\u008d\u008e\u0005s\u0000\u0000\u008e\u0090\u0005e\u0000\u0000"+
		"\u008f\u0086\u0001\u0000\u0000\u0000\u008f\u008a\u0001\u0000\u0000\u0000"+
		"\u0090,\u0001\u0000\u0000\u0000\u0007\u0000cny~\u0084\u008f\u0001\u0006"+
		"\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}