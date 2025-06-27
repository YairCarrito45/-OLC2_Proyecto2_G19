grammar Vlang;

// === Axioma principal ===
programa : declaraciones* EOF ;

declaraciones
    : varDcl
    | structDecl
    | funcionStruct 
    | funcion            
    | stmt
    ;


// === Definición de struct ===
structDecl
    : STRUCT ID LBRACE atributos RBRACE #structDeclaration
    ;

atributos : atributo+ ;

atributo : tipo ID SEMICOLON ;


inicializadorStruct : inicializadorCampo (COMMA inicializadorCampo)* ;

inicializadorCampo : ID COLON expresion ;

funcionStruct
    : FUNC LPAREN ID (MUL)? ID RPAREN ID LPAREN parametrosFormales? RPAREN bloque
    ;

funcion
    : FUNC ID LPAREN parametrosFormales? RPAREN tipo? bloque
    ;


parametrosFormales
    : parametroFormal (COMMA parametroFormal)* ;
    
parametroFormal
    : tipoConReferencia tipoRef=ID
    ;

tipoConReferencia
    : ref=MUL? tipoTipo=tipo
    ;




// === Declaración de variables ===
varDcl
    : MUT ID tipo? ((ASSIGN | ASSIGN_DECL) expresion)? #variableDeclaration
    ;

// === Tipos de datos permitidos ===
tipo
    : sliceTipo
    | 'int'
    | 'float64'
    | 'string'
    | 'bool'
    | 'char'
    ;


sliceTipo
    : LSQUARE RSQUARE tipo ;

// === Sentencias ===
stmt
    : expresion                              #expresionStatement
    | ID LSQUARE expresion RSQUARE LSQUARE expresion RSQUARE ASSIGN expresion  #sliceDobleAsignacionStmt
    | PRINT LPAREN parametros? RPAREN        #printStatement
    | BREAK                                  #breakStatement
    | CONTINUE                               #continueStatement
    | RETURN expresion?                      #returnStatement
    | sentencias_control                     #controlStatement
    | bloque                                 #bloqueStatement
    ;

// === Sentencias de control ===
sentencias_control
    : ifDcl         #if_context 
    | forDcl        #for_context 
    | whileDcl      #while_context
    | switchDcl     #switch_context 
    ;

// === IF / ELSE IF / ELSE ===
ifDcl
    : IF LPAREN? expresion RPAREN? bloque (elseifDcl)* (elseDcl)?
    ;

elseifDcl
    : ELSE IF LPAREN? expresion RPAREN? bloque
    ;

elseDcl
    : ELSE bloque
    ;

// === SWITCH / CASE ===
switchDcl
    : SWITCH expresion LBRACE caseBlock* defaultBlock? RBRACE
    ;

caseBlock
    : CASE expresion COLON declaraciones*
    ;

defaultBlock
    : DEFAULT COLON declaraciones*
    ;

// === Ciclos ===
forDcl
    : FOR expresion bloque                                 #forCondicional
    | FOR expresion SEMICOLON expresion SEMICOLON expresion bloque   #forClasico
    | FOR ID COMMA ID IN ID bloque                         #forForeach
    ;

whileDcl
    : WHILE LPAREN? expresion RPAREN? bloque
    ;

// === Definición de bloque ===
bloque
    : LBRACE declaraciones* RBRACE
    ;

// === Expresiones ===
expresion
    : ID PLUSEQ expresion                                #asignacionCompuestaSuma
    | ID MINUSEQ expresion                               #asignacionCompuestaResta
    | ID LPAREN parametros? RPAREN                       #llamadaFuncion
    | ID ASSIGN expresion                                #asignacionfor
    | ID INC                                             #asignacionSuma
    | ID DEC                                             #asignacionResta
    | ID DOT ID ASSIGN expresion                         #asignacionAtributo       // persona.Nombre = "Bob"
    | ID DOT ID                                          #accesoAtributo           // persona.Nombre
    | ID DOT ID LPAREN parametros? RPAREN                #llamadaMetodoStruct
    | ID LBRACE inicializadorStruct? RBRACE              #instanciaStruct
    | expresion op=(MUL | DIV | MOD) expresion           #multdivmod
    | expresion op=(PLUS | MINUS) expresion              #sumres
    | expresion op=(LT | LE | GE | GT) expresion         #relacionales
    | expresion op=(EQ | NEQ) expresion                  #igualdad
    | expresion AND expresion                            #andExpr
    | expresion OR expresion                             #orExpr
    | op=(NOT | MINUS) expresion                         #unario
    | LPAREN expresion RPAREN                            #parentesisexpre
    | valor                                              #valorexpr
    | ID                                                 #id
    | LBRACE parametros? RBRACE                          #literalSliceExpr
    | LSQUARE RSQUARE tipo LBRACE parametros? RBRACE     #sliceCreacion
    | INDEXOF LPAREN expresion COMMA expresion RPAREN    #indexOfExpr
    | JOIN LPAREN expresion COMMA expresion RPAREN       #joinExpr
    | LEN LPAREN expresion RPAREN                        #lenExpr
    | APPEND LPAREN expresion COMMA expresion RPAREN     #appendExpr
    | ID LSQUARE expresion RSQUARE                       #sliceAcceso
    | ID LSQUARE expresion RSQUARE ASSIGN expresion      #sliceAsignacion
    | ID LSQUARE expresion RSQUARE LSQUARE expresion RSQUARE              #sliceDobleAcceso
    | ID LSQUARE expresion RSQUARE LSQUARE expresion RSQUARE ASSIGN expresion  #sliceDobleAsignacion
    ;


// === Parámetros de funciones ===
parametros : expresion (COMMA expresion)* ;

// === Literales primitivos ===
valor
    : ENTERO    #valorEntero
    | DECIMAL   #valorDecimal
    | CADENA    #valorCadena
    | BOOLEANO  #valorBooleano
    | CARACTER  #valorCaracter
    | NIL       #valorNulo 
    ;

// === Palabras clave ===
MUT     : 'mut' ;
STRUCT  : 'struct' ;
FUNC     : 'fn' ;
PRINT   : 'println' ;
IF      : 'if' ;
ELSE    : 'else' ;
FOR     : 'for' ;
WHILE   : 'while' ;
SWITCH  : 'switch' ;
CASE    : 'case' ;
DEFAULT : 'default' ;
BREAK   : 'break' ;
CONTINUE: 'continue' ;
RETURN  : 'return' ;
COLON   : ':' ;
LEN     : 'len' ;
CAP     : 'cap' ;
INDEXOF : 'indexOf' ;
JOIN    : 'join' ;
APPEND  : 'append' ;
IN      : 'in' ;




// === Tokens de valor ===
BOOLEANO: 'true' | 'false' ;
ENTERO  : [0-9]+ ;
DECIMAL : [0-9]+ '.' [0-9]+ ;
CADENA : '"' (~["\\] | '\\' .)* '"' ;
CARACTER: '\'' . '\'' ;
NIL     : 'nil' ;


// === Identificadores ===
ID: [_a-zA-Z][_a-zA-Z0-9]* ;

// === Operadores ===
PLUS        : '+' ;
MINUS       : '-' ;
MUL         : '*' ;
DIV         : '/' ;
MOD         : '%' ;
NOT         : '!' ;
AND         : '&&' ;
OR          : '||' ;
EQ          : '==' ;
NEQ         : '!=' ;
LT          : '<' ;
LE          : '<=' ;
GT          : '>' ;
GE          : '>=' ;
ASSIGN      : '=' ;
ASSIGN_DECL : ':=' ;
PLUSEQ      : '+=' ;
MINUSEQ     : '-=' ;
INC         : '++' ;
DEC         : '--' ;

// === Símbolos ===
LPAREN     : '(' ;
RPAREN     : ')' ;
LBRACE     : '{' ;
RBRACE     : '}' ;
LSQUARE    : '[' ;
RSQUARE    : ']' ;
DOT        : '.' ;
COMMA      : ',' ;
SEMICOLON  : ';' ;

// === Espacios y comentarios ===
WS            : [ \t\r\n]+ -> skip ;
LINE_COMMENT  : '//' ~[\r\n]* -> skip ;
BLOCK_COMMENT : '/*' .*? '*/' -> skip ;
