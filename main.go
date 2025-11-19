package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// Esta funcion se encarga de generar el bocadillo de la frase de fortune con el uso de los arrays de "borders" y "ret".
// El contador hace una iteracion con el array de "lines" para calcular la longitud del output de fortune y asi saber cuantas lineas debe usar.
func buildBalloon(lines []string, maxwidth int) string {
	var borders []string
	count := len(lines)
	var ret []string

	borders = []string{"/", "\\", "\\", "/", "|", "<", ">"}

	top := " " + strings.Repeat("_", maxwidth+2)
	bottom := " " + strings.Repeat("-", maxwidth+2)

	ret = append(ret, top)
	if count == 1 {
		s := fmt.Sprintf("%s %s %s", borders[5], lines[0], borders[6])
		ret = append(ret, s)

	} else {
		s := fmt.Sprintf(`%s %s %s`, borders[0], lines[0], borders[1])
		ret = append(ret, s)
		i := 1
		for ; i < count-1; i++ {
			s = fmt.Sprintf(`%s %s %s`, borders[4], lines[i], borders[4])
			ret = append(ret, s)
		}
		s = fmt.Sprintf(`%s %s %s`, borders[2], lines[i], borders[3])
		ret = append(ret, s)
	}
	ret = append(ret, bottom)
	return strings.Join(ret, "\n")
}

// Esta funcion se encarga de cambiar las tabulaciones por espacios.
func tabsToSpaces(lines []string) []string {
	var ret []string
	for _, l := range lines {
		l = strings.Replace(l, "\t", "    ", -1)
		ret = append(ret, l)
	}
	return ret
}

// Esta funcion se encarga de calcular la dimension del output.
func calculateMaxWidth(lines []string) int {
	w := 0
	for _, l := range lines {
		len := utf8.RuneCountInString(l)
		if len > w {
			w = len
		}
	}

	return w
}

// Esta funcion se encarga de cambiar la longitud del output segun la dimension.
func normalizeStringsLength(lines []string, maxwidth int) []string {
	var ret []string
	for _, l := range lines {
		s := l + strings.Repeat(" ", maxwidth-utf8.RuneCountInString(l))
		ret = append(ret, s)
	}
	return ret
}

// Esta funcion es para las variables de las figuras que se van a imprimir, se pueden añadir figuras mas adelante.
func printFigure(name string) {
	var cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`
	var stegosaurus = `         \                      .       .
          \                    / ` + "`" + `.   .' "
           \           .---.  <    > <    >  .---.
            \          |    \  \ - ~ ~ - /  /    |
          _____           ..-~             ~-..-~
         |     |   \~~~\\.'                    ` + "`" + `./~~~/
        ---------   \__/                         \__/
       .'  O    \     /               /       \  "
      (_____,    ` + "`" + `._.'               |         }  \/~~~/
       ` + "`" + `----.          /       }     |        /    \__/
             ` + "`" + `-.      |       /      |       /      ` + "`" + `. ,~~|
                 ~-.__|      /_ - ~ ^|      /- _      ` + "`" + `..-'
                      |     /        |     /     ~-.     ` + "`" + `-. _  _  _
                      |_____|        |_____|         ~ - . _ _ _ _ _>

	`
	var tux = `
   \
    \
        .--.
       |o_o |
       |:_/ |
      //   \ \
     (|     | )
    /'\_   _/'\
    \___)=(___/
`
	var kitten = `
   \
    \

     |\_/|
     |o o|__
     --*--__\
     C_C_(___)
`
	var whale = `
   \
    \
     \
                '-.
      .---._     \ .--'
    /       '-..__)  ,-'
   |    0           /
    --.__,   .__.,'
     '-.___'._\_.'
`

	switch name {
	case "cow":
		fmt.Println(cow)
	case "stegosaurus":
		fmt.Println(stegosaurus)
	case "tux":
		fmt.Println(tux)
	case "kitten":
		fmt.Println(kitten)
	case "whale":
		fmt.Println(whale)
	default:
		fmt.Println("Figura desconocida")
	}
}

// Esta funcion es la principal del codigo, se encarga de administrar el output de fortune en la terminal.
// Ademas tambien se encarga de las flags para el comando.
func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("El comando debe funcionar con las pipes")
		fmt.Println("Usa fortune/figlet | gocowsay")
		return
	}

	var lines []string

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	var figure string
	flag.StringVar(&figure, "f", "cow", "el nombre de la figura debe ser entre cow, stegosaurus, tux, kitten, whale ")
	flag.Parse()

	lines = tabsToSpaces(lines)
	maxwidth := calculateMaxWidth(lines)
	messages := normalizeStringsLength(lines, maxwidth)
	balloon := buildBalloon(messages, maxwidth)
	fmt.Println(balloon)
	printFigure(figure)
	fmt.Println()
}
