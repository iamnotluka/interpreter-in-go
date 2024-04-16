# interpreter-in-go

Interpreters are magical. Seemingly random characters (letters, numbers and special characters) are feed into it and suddenly become meaningful.

In this project, I have built tree-walking interpreter with it's own lexer, parser and own tree representation, as well as it's own evaluator.

Every interpreter is built to interpret a specific programming language. That's how the programming languages are implemented basically. In this project, I am implementing a language called Monkey and it has list of these features:

- C like syntax
- variable bindings
- integers and booleans
- arithmetic expressions
- built-in functions
- first-class and higher-order functions
- closures
- a string data structure
- an aray data structure
- a hash data structure

We will implement first class functions, what this means is that functions are just values same as integers or strings.

The project will have couple major parts:

- the lexer
- the parser
- the AST (Abstract Syntax Tree)
- the internal object system
- the evaluator

### 1 Lexing

Representing our code in other forms that are easier to work with.

The representation of the source code is done twice before it's able to be evaluated.

Source code -> Tokens -> AST

The source code to tokens transformation is called "lexical analysis"

Example of what comes into the lexer:

```
let x = 5 + 5;
```

And this comes out as the output:

```
[
    LET,
    IDENTIFIER('x),
    EQUAL_SIGN,
    INTEGER(5),
    PLUS_SIGN,
    INTEGER(5),
    SEMICOLON
]
```

#### Token Definition
