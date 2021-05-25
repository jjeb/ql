// Generated from /Users/juanesparza/src/github.com/ql/parser/QlLexer.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class QlLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		COMMA=1, DOT=2, LR_BRACKET=3, RR_BRACKET=4, ARITHMETIC_OPERATOR=5, PLUS=6, 
		MINUS=7, DIV=8, MOD=9, MUL=10, SHOW=11, AS=12, FROM=13, SINCE=14, UNTIL=15, 
		TIMESPAN_UNIT=16, SECONDS=17, MINUTES=18, HOURS=19, DAYS=20, WEEKS=21, 
		MONTHS=22, QUARTERS=23, YEARS=24, IDENTIFIER=25, SIGNED_NUMERIC_LITERAL=26, 
		NUMERIC_LITERAL=27, YEAR=28, MONTH=29, DAY=30, DATE_LITERAL=31, SPACES=32;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"COMMA", "DOT", "LR_BRACKET", "RR_BRACKET", "ARITHMETIC_OPERATOR", "PLUS", 
			"MINUS", "DIV", "MOD", "MUL", "SHOW", "AS", "FROM", "SINCE", "UNTIL", 
			"TIMESPAN_UNIT", "SECONDS", "MINUTES", "HOURS", "DAYS", "WEEKS", "MONTHS", 
			"QUARTERS", "YEARS", "IDENTIFIER", "SIGNED_NUMERIC_LITERAL", "NUMERIC_LITERAL", 
			"YEAR", "MONTH", "DAY", "DATE_LITERAL", "SPACES", "DIGIT", "A", "B", 
			"C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", 
			"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "','", "'.'", "'('", "')'", null, "'+'", "'-'", "'/'", "'%'", "'*'", 
			null, null, null, null, null, null, "'s'", "'min'", "'h'", "'d'", "'w'", 
			"'m'", "'q'", "'y'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "COMMA", "DOT", "LR_BRACKET", "RR_BRACKET", "ARITHMETIC_OPERATOR", 
			"PLUS", "MINUS", "DIV", "MOD", "MUL", "SHOW", "AS", "FROM", "SINCE", 
			"UNTIL", "TIMESPAN_UNIT", "SECONDS", "MINUTES", "HOURS", "DAYS", "WEEKS", 
			"MONTHS", "QUARTERS", "YEARS", "IDENTIFIER", "SIGNED_NUMERIC_LITERAL", 
			"NUMERIC_LITERAL", "YEAR", "MONTH", "DAY", "DATE_LITERAL", "SPACES"
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


	public QlLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "QlLexer.g4"; }

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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\"\u0140\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\3\2\3"+
		"\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\6\3\6\3\6\5\6\u0087\n\6\3\7\3\7\3"+
		"\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\16"+
		"\3\16\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20"+
		"\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u00b4\n\21\3\22"+
		"\3\22\3\23\3\23\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3\27\3\27\3\30"+
		"\3\30\3\31\3\31\3\32\3\32\7\32\u00ca\n\32\f\32\16\32\u00cd\13\32\3\33"+
		"\3\33\5\33\u00d1\n\33\3\33\3\33\3\34\6\34\u00d6\n\34\r\34\16\34\u00d7"+
		"\3\34\3\34\7\34\u00dc\n\34\f\34\16\34\u00df\13\34\5\34\u00e1\n\34\3\34"+
		"\3\34\6\34\u00e5\n\34\r\34\16\34\u00e6\5\34\u00e9\n\34\3\34\3\34\5\34"+
		"\u00ed\n\34\3\34\6\34\u00f0\n\34\r\34\16\34\u00f1\5\34\u00f4\n\34\3\35"+
		"\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3"+
		"!\3!\3!\3!\3\"\3\"\3#\3#\3$\3$\3%\3%\3&\3&\3\'\3\'\3(\3(\3)\3)\3*\3*\3"+
		"+\3+\3,\3,\3-\3-\3.\3.\3/\3/\3\60\3\60\3\61\3\61\3\62\3\62\3\63\3\63\3"+
		"\64\3\64\3\65\3\65\3\66\3\66\3\67\3\67\38\38\39\39\3:\3:\3;\3;\3<\3<\2"+
		"\2=\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35"+
		"\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36"+
		";\37= ?!A\"C\2E\2G\2I\2K\2M\2O\2Q\2S\2U\2W\2Y\2[\2]\2_\2a\2c\2e\2g\2i"+
		"\2k\2m\2o\2q\2s\2u\2w\2\3\2!\5\2C\\aac|\6\2\62;C\\aac|\4\2--//\5\2\13"+
		"\r\17\17\"\"\3\2\62;\4\2CCcc\4\2DDdd\4\2EEee\4\2FFff\4\2GGgg\4\2HHhh\4"+
		"\2IIii\4\2JJjj\4\2KKkk\4\2LLll\4\2MMmm\4\2NNnn\4\2OOoo\4\2PPpp\4\2QQq"+
		"q\4\2RRrr\4\2SSss\4\2TTtt\4\2UUuu\4\2VVvv\4\2WWww\4\2XXxx\4\2YYyy\4\2"+
		"ZZzz\4\2[[{{\4\2\\\\||\2\u013a\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t"+
		"\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2"+
		"\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2"+
		"\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2"+
		"+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2"+
		"\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\3"+
		"y\3\2\2\2\5{\3\2\2\2\7}\3\2\2\2\t\177\3\2\2\2\13\u0086\3\2\2\2\r\u0088"+
		"\3\2\2\2\17\u008a\3\2\2\2\21\u008c\3\2\2\2\23\u008e\3\2\2\2\25\u0090\3"+
		"\2\2\2\27\u0092\3\2\2\2\31\u0097\3\2\2\2\33\u009a\3\2\2\2\35\u009f\3\2"+
		"\2\2\37\u00a5\3\2\2\2!\u00b3\3\2\2\2#\u00b5\3\2\2\2%\u00b7\3\2\2\2\'\u00bb"+
		"\3\2\2\2)\u00bd\3\2\2\2+\u00bf\3\2\2\2-\u00c1\3\2\2\2/\u00c3\3\2\2\2\61"+
		"\u00c5\3\2\2\2\63\u00c7\3\2\2\2\65\u00d0\3\2\2\2\67\u00e8\3\2\2\29\u00f5"+
		"\3\2\2\2;\u00fa\3\2\2\2=\u00fd\3\2\2\2?\u0100\3\2\2\2A\u0106\3\2\2\2C"+
		"\u010a\3\2\2\2E\u010c\3\2\2\2G\u010e\3\2\2\2I\u0110\3\2\2\2K\u0112\3\2"+
		"\2\2M\u0114\3\2\2\2O\u0116\3\2\2\2Q\u0118\3\2\2\2S\u011a\3\2\2\2U\u011c"+
		"\3\2\2\2W\u011e\3\2\2\2Y\u0120\3\2\2\2[\u0122\3\2\2\2]\u0124\3\2\2\2_"+
		"\u0126\3\2\2\2a\u0128\3\2\2\2c\u012a\3\2\2\2e\u012c\3\2\2\2g\u012e\3\2"+
		"\2\2i\u0130\3\2\2\2k\u0132\3\2\2\2m\u0134\3\2\2\2o\u0136\3\2\2\2q\u0138"+
		"\3\2\2\2s\u013a\3\2\2\2u\u013c\3\2\2\2w\u013e\3\2\2\2yz\7.\2\2z\4\3\2"+
		"\2\2{|\7\60\2\2|\6\3\2\2\2}~\7*\2\2~\b\3\2\2\2\177\u0080\7+\2\2\u0080"+
		"\n\3\2\2\2\u0081\u0087\5\r\7\2\u0082\u0087\5\17\b\2\u0083\u0087\5\21\t"+
		"\2\u0084\u0087\5\23\n\2\u0085\u0087\5\25\13\2\u0086\u0081\3\2\2\2\u0086"+
		"\u0082\3\2\2\2\u0086\u0083\3\2\2\2\u0086\u0084\3\2\2\2\u0086\u0085\3\2"+
		"\2\2\u0087\f\3\2\2\2\u0088\u0089\7-\2\2\u0089\16\3\2\2\2\u008a\u008b\7"+
		"/\2\2\u008b\20\3\2\2\2\u008c\u008d\7\61\2\2\u008d\22\3\2\2\2\u008e\u008f"+
		"\7\'\2\2\u008f\24\3\2\2\2\u0090\u0091\7,\2\2\u0091\26\3\2\2\2\u0092\u0093"+
		"\5i\65\2\u0093\u0094\5S*\2\u0094\u0095\5a\61\2\u0095\u0096\5q9\2\u0096"+
		"\30\3\2\2\2\u0097\u0098\5E#\2\u0098\u0099\5i\65\2\u0099\32\3\2\2\2\u009a"+
		"\u009b\5O(\2\u009b\u009c\5g\64\2\u009c\u009d\5a\61\2\u009d\u009e\5]/\2"+
		"\u009e\34\3\2\2\2\u009f\u00a0\5i\65\2\u00a0\u00a1\5U+\2\u00a1\u00a2\5"+
		"_\60\2\u00a2\u00a3\5I%\2\u00a3\u00a4\5M\'\2\u00a4\36\3\2\2\2\u00a5\u00a6"+
		"\5m\67\2\u00a6\u00a7\5_\60\2\u00a7\u00a8\5k\66\2\u00a8\u00a9\5U+\2\u00a9"+
		"\u00aa\5[.\2\u00aa \3\2\2\2\u00ab\u00b4\5#\22\2\u00ac\u00b4\5%\23\2\u00ad"+
		"\u00b4\5\'\24\2\u00ae\u00b4\5)\25\2\u00af\u00b4\5+\26\2\u00b0\u00b4\5"+
		"-\27\2\u00b1\u00b4\5/\30\2\u00b2\u00b4\5\61\31\2\u00b3\u00ab\3\2\2\2\u00b3"+
		"\u00ac\3\2\2\2\u00b3\u00ad\3\2\2\2\u00b3\u00ae\3\2\2\2\u00b3\u00af\3\2"+
		"\2\2\u00b3\u00b0\3\2\2\2\u00b3\u00b1\3\2\2\2\u00b3\u00b2\3\2\2\2\u00b4"+
		"\"\3\2\2\2\u00b5\u00b6\7u\2\2\u00b6$\3\2\2\2\u00b7\u00b8\7o\2\2\u00b8"+
		"\u00b9\7k\2\2\u00b9\u00ba\7p\2\2\u00ba&\3\2\2\2\u00bb\u00bc\7j\2\2\u00bc"+
		"(\3\2\2\2\u00bd\u00be\7f\2\2\u00be*\3\2\2\2\u00bf\u00c0\7y\2\2\u00c0,"+
		"\3\2\2\2\u00c1\u00c2\7o\2\2\u00c2.\3\2\2\2\u00c3\u00c4\7s\2\2\u00c4\60"+
		"\3\2\2\2\u00c5\u00c6\7{\2\2\u00c6\62\3\2\2\2\u00c7\u00cb\t\2\2\2\u00c8"+
		"\u00ca\t\3\2\2\u00c9\u00c8\3\2\2\2\u00ca\u00cd\3\2\2\2\u00cb\u00c9\3\2"+
		"\2\2\u00cb\u00cc\3\2\2\2\u00cc\64\3\2\2\2\u00cd\u00cb\3\2\2\2\u00ce\u00d1"+
		"\5\r\7\2\u00cf\u00d1\5\17\b\2\u00d0\u00ce\3\2\2\2\u00d0\u00cf\3\2\2\2"+
		"\u00d0\u00d1\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2\u00d3\5\67\34\2\u00d3\66"+
		"\3\2\2\2\u00d4\u00d6\5C\"\2\u00d5\u00d4\3\2\2\2\u00d6\u00d7\3\2\2\2\u00d7"+
		"\u00d5\3\2\2\2\u00d7\u00d8\3\2\2\2\u00d8\u00e0\3\2\2\2\u00d9\u00dd\7\60"+
		"\2\2\u00da\u00dc\5C\"\2\u00db\u00da\3\2\2\2\u00dc\u00df\3\2\2\2\u00dd"+
		"\u00db\3\2\2\2\u00dd\u00de\3\2\2\2\u00de\u00e1\3\2\2\2\u00df\u00dd\3\2"+
		"\2\2\u00e0\u00d9\3\2\2\2\u00e0\u00e1\3\2\2\2\u00e1\u00e9\3\2\2\2\u00e2"+
		"\u00e4\7\60\2\2\u00e3\u00e5\5C\"\2\u00e4\u00e3\3\2\2\2\u00e5\u00e6\3\2"+
		"\2\2\u00e6\u00e4\3\2\2\2\u00e6\u00e7\3\2\2\2\u00e7\u00e9\3\2\2\2\u00e8"+
		"\u00d5\3\2\2\2\u00e8\u00e2\3\2\2\2\u00e9\u00f3\3\2\2\2\u00ea\u00ec\5M"+
		"\'\2\u00eb\u00ed\t\4\2\2\u00ec\u00eb\3\2\2\2\u00ec\u00ed\3\2\2\2\u00ed"+
		"\u00ef\3\2\2\2\u00ee\u00f0\5C\"\2\u00ef\u00ee\3\2\2\2\u00f0\u00f1\3\2"+
		"\2\2\u00f1\u00ef\3\2\2\2\u00f1\u00f2\3\2\2\2\u00f2\u00f4\3\2\2\2\u00f3"+
		"\u00ea\3\2\2\2\u00f3\u00f4\3\2\2\2\u00f48\3\2\2\2\u00f5\u00f6\5C\"\2\u00f6"+
		"\u00f7\5C\"\2\u00f7\u00f8\5C\"\2\u00f8\u00f9\5C\"\2\u00f9:\3\2\2\2\u00fa"+
		"\u00fb\5C\"\2\u00fb\u00fc\5C\"\2\u00fc<\3\2\2\2\u00fd\u00fe\5C\"\2\u00fe"+
		"\u00ff\5C\"\2\u00ff>\3\2\2\2\u0100\u0101\59\35\2\u0101\u0102\7/\2\2\u0102"+
		"\u0103\5;\36\2\u0103\u0104\7/\2\2\u0104\u0105\5=\37\2\u0105@\3\2\2\2\u0106"+
		"\u0107\t\5\2\2\u0107\u0108\3\2\2\2\u0108\u0109\b!\2\2\u0109B\3\2\2\2\u010a"+
		"\u010b\t\6\2\2\u010bD\3\2\2\2\u010c\u010d\t\7\2\2\u010dF\3\2\2\2\u010e"+
		"\u010f\t\b\2\2\u010fH\3\2\2\2\u0110\u0111\t\t\2\2\u0111J\3\2\2\2\u0112"+
		"\u0113\t\n\2\2\u0113L\3\2\2\2\u0114\u0115\t\13\2\2\u0115N\3\2\2\2\u0116"+
		"\u0117\t\f\2\2\u0117P\3\2\2\2\u0118\u0119\t\r\2\2\u0119R\3\2\2\2\u011a"+
		"\u011b\t\16\2\2\u011bT\3\2\2\2\u011c\u011d\t\17\2\2\u011dV\3\2\2\2\u011e"+
		"\u011f\t\20\2\2\u011fX\3\2\2\2\u0120\u0121\t\21\2\2\u0121Z\3\2\2\2\u0122"+
		"\u0123\t\22\2\2\u0123\\\3\2\2\2\u0124\u0125\t\23\2\2\u0125^\3\2\2\2\u0126"+
		"\u0127\t\24\2\2\u0127`\3\2\2\2\u0128\u0129\t\25\2\2\u0129b\3\2\2\2\u012a"+
		"\u012b\t\26\2\2\u012bd\3\2\2\2\u012c\u012d\t\27\2\2\u012df\3\2\2\2\u012e"+
		"\u012f\t\30\2\2\u012fh\3\2\2\2\u0130\u0131\t\31\2\2\u0131j\3\2\2\2\u0132"+
		"\u0133\t\32\2\2\u0133l\3\2\2\2\u0134\u0135\t\33\2\2\u0135n\3\2\2\2\u0136"+
		"\u0137\t\34\2\2\u0137p\3\2\2\2\u0138\u0139\t\35\2\2\u0139r\3\2\2\2\u013a"+
		"\u013b\t\36\2\2\u013bt\3\2\2\2\u013c\u013d\t\37\2\2\u013dv\3\2\2\2\u013e"+
		"\u013f\t \2\2\u013fx\3\2\2\2\17\2\u0086\u00b3\u00cb\u00d0\u00d7\u00dd"+
		"\u00e0\u00e6\u00e8\u00ec\u00f1\u00f3\3\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}