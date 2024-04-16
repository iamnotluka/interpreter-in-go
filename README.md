# interpreter-in-go

Interpreters are magical. Seemingly random characters (letters, numbers and special characters) are feed into it and suddenly become meaningful. In this project, I have built tree-walking interpreter with it's own lexer, parser and own tree representation, as well as it's own evaluator. Every interpreter is built to interpret a specific programming language. That's how the programming languages are implemented basically.

I will implement first class functions, what this means is that functions are just values same as integers or strings.

The project will have couple major parts:

- The Lexer
- The Parser
- The AST (Abstract Syntax Tree)
- The Internal Object System
- The Evaluator

## Lexer

Representing our code in other forms that are easier to work with.

The representation of the source code is done twice before it's able to be evaluated.

Source code -> Tokens -> AST

The source code to tokens transformation is called "lexical analysis".

Below is the example of what comes into the lexer and outputs.

If you wish to run the lexer yourself just run the main.go function in our lexer folder.

```
cd 01_lexer && go run main.go
```

```
➜  01_lexer git:(main) ✗ go run main.go
Hello zoricl! This is Luka's programming language called fluxo. Welcome!
Feel free to start typing in commands.
>> let add = fn(x, y) { x + y; };
{Type:LET Literal:let}
{Type:IDENT Literal:add}
{Type:= Literal:=}
{Type:FUNCTION Literal:fn}
{Type:( Literal:(}
{Type:IDENT Literal:x}
{Type:, Literal:,}
{Type:IDENT Literal:y}
{Type:) Literal:)}
{Type:{ Literal:{}
{Type:IDENT Literal:x}
{Type:+ Literal:+}
{Type:IDENT Literal:y}
{Type:; Literal:;}
{Type:} Literal:}}
{Type:; Literal:;}
```

## Parsing
