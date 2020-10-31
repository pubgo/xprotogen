package cmds

import (
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pubgo/xerror"
	"github.com/pubgo/xprotogen/internal/watcher"
	"github.com/spf13/cobra"
	"github.com/valyala/quicktemplate/parser"
)

var logger = log.New(os.Stderr, "xprotogen: ", log.LstdFlags)
var filesCompiled int
var ext = "tpl"
var w bool

func GenCmd() *cobra.Command {
	var (
		cmd = &cobra.Command{
			Use:     "gen",
			Example: "  xprotogen gen [template dir]",
		}
		dir = tplDir
	)

	//cmd.Flags().StringVar(&dir, "dir", tplDir, "Path to directory with template files to compile.")
	cmd.Flags().BoolVarP(&w, "watch", "w", w, "watch the template files to compile.")

	cmd.Run = func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			dir = strings.TrimSpace(args[0])
		}

		if len(dir) == 0 {
			dir = tplDir
		}

		logger.Printf("Compiling *%s template files in directory %q", ext, dir)
		compileDir(dir)
		logger.Printf("Total files compiled: %d", filesCompiled)

		if !w {
			return
		}

		watcher.AddExclude(ext)
		xerror.Exit(watcher.AddRecursive(dir, func(event interface{}) (err error) {
			defer xerror.RespErr(&err)

			if event == nil {
				return nil
			}

			switch event := event.(type) {
			case watcher.Event:
				if watcher.IsUpdateEvent(event) {
					compileFile(event.Name)
				}
			}
			return nil
		}))
		watcher.Start()

		logger.Println("watching:", watcher.List())
		select {}
	}

	return cmd
}

func compileDir(path string) {
	fi, err := os.Stat(path)
	if err != nil {
		logger.Fatalf("cannot compile files in %q: %s", path, err)
	}
	if !fi.IsDir() {
		logger.Fatalf("cannot compile files in %q: it is not directory", path)
	}
	d, err := os.Open(path)
	if err != nil {
		logger.Fatalf("cannot compile files in %q: %s", path, err)
	}
	defer d.Close()

	fis, err := d.Readdir(-1)
	if err != nil {
		logger.Fatalf("cannot read files in %q: %s", path, err)
	}

	var names []string
	for _, fi = range fis {
		name := fi.Name()
		if name == "." || name == ".." {
			continue
		}
		if !fi.IsDir() {
			names = append(names, name)
		} else {
			subPath := filepath.Join(path, name)
			compileDir(subPath)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		if strings.HasSuffix(name, ext) {
			filename := filepath.Join(path, name)
			compileFile(filename)
		}
	}
}

func cleanPath(path string) string {
	return strings.Trim(strings.TrimPrefix(path, filepath.Dir(path)), "/")
}

func compileFile(infile string) {
	outfile := infile + ".go"
	logger.Printf("Compiling %q to %q...", cleanPath(infile), cleanPath(outfile))

	inf, err := os.Open(infile)
	if err != nil {
		logger.Fatalf("cannot open file %q: %s", infile, err)
	}

	tmpfile := outfile + ".tmp"
	outf, err := os.Create(tmpfile)
	if err != nil {
		logger.Fatalf("cannot create file %q: %s", tmpfile, err)
	}

	packageName, err := getPackageName(infile)
	if err != nil {
		logger.Fatalf("cannot determine package name for %q: %s", infile, err)
	}

	if err = parser.Parse(outf, inf, infile, packageName); err != nil {
		logger.Fatalf("error when parsing file %q: %s", infile, err)
	}

	if err = outf.Close(); err != nil {
		logger.Fatalf("error when closing file %q: %s", tmpfile, err)
	}
	if err = inf.Close(); err != nil {
		logger.Fatalf("error when closing file %q: %s", infile, err)
	}

	// prettify the output file
	uglyCode, err := ioutil.ReadFile(tmpfile)
	if err != nil {
		logger.Fatalf("cannot read file %q: %s", tmpfile, err)
	}
	prettyCode, err := format.Source(uglyCode)
	if err != nil {
		logger.Fatalf("error when formatting compiled code for %q: %s. See %q for details", infile, err, tmpfile)
	}
	if err = ioutil.WriteFile(outfile, prettyCode, 0666); err != nil {
		logger.Fatalf("error when writing file %q: %s", outfile, err)
	}
	if err = os.Remove(tmpfile); err != nil {
		logger.Fatalf("error when removing file %q: %s", tmpfile, err)
	}

	filesCompiled++
}

func getPackageName(filename string) (string, error) {
	filenameAbs, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}

	dir, _ := filepath.Split(filenameAbs)
	return filepath.Base(dir), nil
}

func init() {
	rootCmd.AddCommand(GenCmd())
}
