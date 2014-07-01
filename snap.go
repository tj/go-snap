package main

import snap "github.com/mreiferson/go-snappystream"
import "github.com/visionmedia/go-flags"
import "os"
import "io"

type Options struct {
	Encode bool `short:"e" long:"encode" description:"encode stdin (the default)"`
	Decode bool `short:"d" long:"decode" description:"decode stdin"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	opts := &Options{}
	_, err := flags.Parse(opts)
	check(err)

	if opts.Decode {
		decode()
	} else {
		encode()
	}
}

func encode() {
	_, err := io.Copy(snap.NewWriter(os.Stdout), os.Stdin)
	check(err)
}

func decode() {
	_, err := io.Copy(os.Stdout, snap.NewReader(os.Stdin, snap.VerifyChecksum))
	check(err)
}
