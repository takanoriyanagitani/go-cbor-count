package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	cc "github.com/takanoriyanagitani/go-cbor-count"

	ci "github.com/takanoriyanagitani/go-cbor-count/iter/cbor2iter"
	ca "github.com/takanoriyanagitani/go-cbor-count/iter/cbor2iter/amacker"
)

func count2stdout(_ context.Context, cnt uint64) error {
	_, e := fmt.Printf("%v\n", cnt)
	return e
}

var coutput cc.CountOutput = count2stdout

var counter ci.Counter = ci.CounterDefault

func main() {
	var r io.Reader = os.Stdin
	var br io.Reader = bufio.NewReader(r)

	isrc := ca.DecIterNew(br)
	var ise ci.IterSourceEmpty = isrc.ToIterSourceEmptyArray()
	var csource cc.CountSource = ise.ToCountSource(counter)

	e := csource.OutputCount(context.Background(), coutput)
	if nil != e {
		log.Printf("%v\n", e)
	}
}
