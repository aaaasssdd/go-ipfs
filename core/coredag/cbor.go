package coredag

import (
	"io"
	"io/ioutil"

	ipldcbor "gx/ipfs/QmRVSCwQtW1rjHCay9NqKXDwbtKTgDcN4iY7PrpSqfKM5D/go-ipld-cbor"
	ipld "gx/ipfs/Qme5bWv7wtjUNGsK2BNGVUFPKiuxWrsqrtvYwCLRw8YFES/go-ipld-format"
)

func cborJSONParser(r io.Reader, mhType uint64, mhLen int) ([]ipld.Node, error) {
	nd, err := ipldcbor.FromJSON(r, mhType, mhLen)
	if err != nil {
		return nil, err
	}

	return []ipld.Node{nd}, nil
}

func cborRawParser(r io.Reader, mhType uint64, mhLen int) ([]ipld.Node, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	nd, err := ipldcbor.Decode(data, mhType, mhLen)
	if err != nil {
		return nil, err
	}

	return []ipld.Node{nd}, nil
}
