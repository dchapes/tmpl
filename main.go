package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"text/template"

	"github.com/Masterminds/sprig"
)

func main() {
	log.SetPrefix("tmpl: ")
	log.SetFlags(0)
	input := flag.String("f", "-", "Input source")
	flag.Parse()

	in, err := getInput(*input)
	if err != nil {
		log.Fatal(err)
	}
	if err = tmpl(in, os.Stdout, envMap()); err != nil {
		log.Fatal(err)
	}
	if err = in.Close(); err != nil {
		log.Fatal(err)
	}
}

func getInput(path string) (*os.File, error) {
	if path == "-" {
		return os.Stdin, nil
	}
	return os.Open(path)
}

func envMap() map[string]string {
	env := os.Environ()
	result := make(map[string]string, len(env))
	for _, envvar := range env {
		parts := strings.SplitN(envvar, "=", 2)
		result[parts[0]] = parts[1]
	}
	return result
}

func tmpl(in io.Reader, out io.Writer, ctx map[string]string) error {
	i, err := ioutil.ReadAll(in)
	if err != nil {
		return err
	}
	tmpl, err := template.New("format string").Funcs(sprig.TxtFuncMap()).Parse(string(i))
	if err != nil {
		return err
	}
	return tmpl.Execute(out, ctx)
}
