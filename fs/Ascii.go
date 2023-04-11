package fs

import (
	"strings"
	"web/fs/pkg"
)

func Ascii(text, banner string) string {
	var filename string
	filename = "fs/" + banner + ".txt"
	str, err := pkg.Getstrings(filename)
	pkg.Chekck(err)
	er := strings.ReplaceAll(text, "\n", "")
	res := strings.Split(er, "\r")
	mymap := pkg.Map(str)
	raz := pkg.Print(mymap, res)
	return raz
}
