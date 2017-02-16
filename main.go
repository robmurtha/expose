package main

import (
	"bytes"
	"os/exec"
	"path"
	"strings"

	"log"
	"os"

	"encoding/json"
	"fmt"

	"flag"

	"github.com/robmurtha/expose/exiv2"
)

var groupFlags = map[string]string{
	"exif": "-PEgnycv",
	"xmp":  "-PXgnycv",
	"iptc": "-PIgnycv",
}

// exiv2 print flags
//-P flgs Print flags for fine control of tag lists ('print' action):
//E : include Exif tags in the list
//I : IPTC datasets
//X : XMP properties
//x : print a column with the tag number
//g : group name
//k : key
//l : tag label
//n : tag name
//y : type
//c : number of components (count)
//s : size in bytes
//v : plain data value
//t : interpreted (translated) data
//h : hexdump of the data

// TBD support labels
// labels introduce a varying number of spaces so they require another pass to parse
// exiv2 -PXgnl testdata/samsung_galaxy.mp4


var verbose = false
var pretty = false

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: expose [-flags] filename")
		fmt.Fprintf(os.Stderr, "\nflags:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func main() {

	flag.BoolVar(&verbose, "v", verbose, "verbose logging")
	flag.BoolVar(&pretty, "p", verbose, "pretty JSON")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
	}

	filename := flag.Arg(0)

	cmdName, err := exec.LookPath("exiv2")
	if err != nil {
		log.Fatalf("exiv2 not found")
	}

	expose := map[string]interface{}{}
	for group, flags := range groupFlags {
		cmdArgs := []string{
			flags, filename,
		}
		if verbose {
			log.Printf("%s %s %s\n", cmdName, flags, filename)
		}
		out, err := exec.Command(cmdName, cmdArgs...).CombinedOutput()
		if verbose {
			log.Print(string(out))
		}
		if err != nil {
			if strings.Contains(string(out), "unknown image type") {
				log.Fatal("expose: unsupported image type try building exiv2 with --enable-video " + path.Ext(filename))
			}
			if verbose {
				log.Println("expose error: " + err.Error())
			}

		}
		if len(out) == 0 {
			continue
		}
		e := exiv2.New(bytes.NewReader(out))
		fields, err := e.Fields()
		if err != nil {
			log.Fatal("expose error: " + err.Error())
		}
		if len(fields) > 0 {
			expose[group] = fields
		}
	}
	if len(expose) > 0 {
		out, err := marshal(map[string]interface{}{"expose": expose}, pretty)
		if err != nil {
			log.Fatal("expose error: " + err.Error())
		}
		fmt.Println(string(out))

	}

}

func marshal(data interface{}, pretty bool) ([]byte, error) {
	if pretty {
		return json.MarshalIndent(data, "", " ")
	}
	return json.Marshal(data)
}
