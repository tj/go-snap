package main

import "github.com/mreiferson/go-snappystream"
import "bufio"
import "fmt"
import "os"
import "io"

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	snap := snappystream.NewWriter(out)

	n, err := io.Copy(snap, in)
	if err != nil {
		panic(err)
	}

	out.Flush()
}
