grammar Gramar;

WS: [ \n\t] -> skip;
COMS: '//'~[\n]* -> skip;
COMM: '/*' .*? '*/'->skip;

/*Palabras reservadas */
BOOLEAN: ('true' | 'false');
RVAR: 'var';
RLET: 'let';
RINT: 'Int';
RFLOAT: 'Float';
RSTRING: 'String';
RBOOL: 'Bool';
RChar: 'Character';
GUION: '_';
INOUT: 'inout';
PUNTE: '&';
ABR: '[';
CIER: ']';
POINTS: '...';
BRK: 'break';


/*Expresiones regulares */
ID: [a-zA-Z_][a-zA-Z0-9]*;
INT: [0-9]+;
DOUBLE: [0-9]+'.'[0-9]+;
CADENA: '"' ( ~["\\] | '\\' [nrt\\"\\] )* '"';
CHAR: '"' [a-zA-Z0-9]? '"';


s: block EOF;

block: (productions)*;

productions: prin
    | swhile
    | pif
    | pdeclara
    | asign
    | pfor
    | pswitch
    | pfuncion
    | pllamada
    | pbrk
;

pbrk: BRK
;
pswitch: 'switch' expr '{' (ccase)+ pdefault '}';

ccase: 'case' expr ':' block 'break'?
;

 pdefault: 'default' ':' block 'break'?
; 

pfor: 'for' ID 'in' ((expr POINTS expr)| expr) '{' block '}';

prin: 'print' '(' expr ')'
    | 'print' '(' cexpr ')'';'?
;

cexpr: expr ',' cexpr
    | expr
;

swhile: 'while'  expr '{' block '}';


asign: ID op=('=' | '+=' | '-=') expr ;

pdeclara: (RVAR | RLET) ID ':' tipo '=' expr ';'?
    | (RVAR | RLET) ID ':' tipo '?'
    | (RVAR | RLET) ID '=' expr
;

tipo: RINT
    | RFLOAT
    | RSTRING
    | RBOOL
    | RChar
;


pif: 'if' expr '{' block '}'
    | 'if' expr '{' block '}' pelse;

pelse: 'else' pif
    | 'else' '{' block '}'
;

pfuncion: 'func' ID '(' ')' ('->' tipo)? '{' block '}'
    | 'func' ID '(' pparams ')' ('->' tipo)? '{' block '}'
;

pparams: (pparamet (',' pparamet)*)?
;

pparamet: pid=(ID | GUION) ID ':' (INOUT)? tipo
    | pid=(ID | GUION) ID ':' (INOUT)? ABR tipo CIER
;

pllamada: ID '(' ')' ';'?           
    | ID '(' argumento ')' ';'?    
;

argumento: tipoArg ',' argumento
    | tipoArg
;

tipoArg: PUNTE exp1=expr
    | ID ':' exp2=expr   
    | exp3=expr
;



expr: op=('!'|'-') right=expr                              #aritmetic
    | '(' expr ')'                                         #parent
    | left=expr op=('/'|'*'|'%') right=expr                #aritmetic
    | left=expr op=('+' | '-') right=expr                  #aritmetic
    | left=expr op=('<' | '>' | '<=' | '>=') right=expr    #aritmetic
    | left=expr op=('=='| '!=') right=expr                 #aritmetic
    | left=expr op='&&' right=expr                         #aritmetic
    | left=expr op='||' right=expr                         #aritmetic
    | INT                                                  #Literal
    | DOUBLE                                               #Literal
    | BOOLEAN                                              #Literal
    | CADENA                                               #Literal
    | CHAR                                                 #Literal
    | ID                                                   #Acceso

    
;