package main

import (
	"log"
	"os"
)

var logI = log.New(os.Stderr, "", 0)
var logE = log.New(os.Stderr, "", log.Llongfile)
