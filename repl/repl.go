package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	//"monkey/token"
	"monkey/parser"
)

const PROMPT = ">> "
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
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
			printParserErrors(out, p.Errors())
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

const DONUT = `      
     888                        888    
     888                        888    
 .d88888 .d88b. 88888b. 888  888888888 
d88" 888d88""88b888 "88b888  888888    
888  888888  888888  888888  888888    
Y88b 888Y88..88P888  888Y88b 888Y88b.  
 "Y88888 "Y88P" 888  888 "Y88888 "Y888 
`
func printParserErrors(out io.Writer, errors []string) {
io.WriteString(out, DONUT)
io.WriteString(out, "Woops! We ran into some trouble \n")
io.WriteString(out, " parser errors:\n")
for _, msg := range errors {
io.WriteString(out, "\t"+msg+"\n")
}
}