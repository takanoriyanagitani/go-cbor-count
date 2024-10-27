package dec2iter

import (
	"io"
	"iter"

	ac "github.com/fxamacker/cbor/v2"

	ci "github.com/takanoriyanagitani/go-cbor-count/iter/cbor2iter"
)

type DecIter struct {
	*ac.Decoder
}

func (d DecIter) ToEmptyIterArray() iter.Seq[struct{}] {
	return func(yield func(struct{}) bool) {
		var buf []any
		var err error
		for {
			err = d.Decoder.Decode(&buf)
			if nil != err {
				return
			}

			if !yield(struct{}{}) {
				return
			}
		}
	}
}

func (d DecIter) ToEmptyIterMap() iter.Seq[struct{}] {
	return func(yield func(struct{}) bool) {
		var buf map[string]any
		var err error
		for {
			err = d.Decoder.Decode(&buf)
			if nil != err {
				return
			}

			if !yield(struct{}{}) {
				return
			}
		}
	}
}

func (d DecIter) ToIterSourceEmptyArray() ci.IterSourceEmpty {
	return func() iter.Seq[struct{}] {
		return d.ToEmptyIterArray()
	}
}

func (d DecIter) ToIterSourceEmptyMap() ci.IterSourceEmpty {
	return func() iter.Seq[struct{}] {
		return d.ToEmptyIterMap()
	}
}

func DecIterNew(rdr io.Reader) DecIter {
	return DecIter{Decoder: ac.NewDecoder(rdr)}
}
