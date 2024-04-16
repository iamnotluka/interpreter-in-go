# Language Intepreter In Go

Writing my own language from scratch following book by Thorsten Ball as a guide: https://thorstenball.com/

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

Typing `quit` exits the interactive experience.

### Testing

There is also `lexer_test.go` file which we can run to test that our lexer is doing what it's supposed to.

You can run `go test ./01_lexer/lexer` command. You should get this output if everything is working properly:

```
➜  01_lexer git:(main) ✗ go test ./lexer
ok      example/lexer/lexer     0.474s
```

## Parsing

We've all encountered the parser in some way - most probably by encountering a "parser error" when trying to run our dodgy code.

What is parsing? Parsing is a component that takes some input data (frequently text) and builds up a data structure like tree or some other hierarchical structure. This allows us to represent input with some structure.

In short, parsers take source code as input (in this case as tokens), and represent it as a data structure which represents this source code. The process unavoidable checks and analyses the input that it confirms to the expected structure. For this reason parsing is also refered to as syntatic analysis.
