package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/himetani/mymonkey/evaluator"
	"github.com/himetani/mymonkey/lexer"
	"github.com/himetani/mymonkey/object"
	"github.com/himetani/mymonkey/parser"
)

const (
	PROMPT      = ">> "
	MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ O | O /  /   / |
  \ '- ,\.-""""""-./, -'  /
   ''-' /_   ^ ^   \ '_''
       |  \._   _./ |
       \   \ '~' /  /
        ,._ '_=_' _.,
           '-----'
`
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
