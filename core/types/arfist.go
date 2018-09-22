// Copyright 2017 The go-ethereum Authors
// This file is part of the goArf library.
//
// The goArf library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The goArf library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the goArf library. If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"errors"
	"io"

	"github.com/arfblocks/goArf/common"
	"github.com/arfblocks/goArf/rlp"
)

var (
	// ArfIstDigest represents a hash of "ArfIst practical byzantine fault tolerance"
	// to identify whether the block is from ArfIst consensus engine
	ArfIstDigest = common.HexToHash("0x63746963616c2062797a616e74696e65206661756c7420746f6c6572616e6365")

	ArfIstExtraVanity = 32 // Fixed number of extra-data bytes reserved for validator vanity
	ArfIstExtraSeal   = 65 // Fixed number of extra-data bytes reserved for validator seal

	// ErrInvalidArfIstHeaderExtra is returned if the length of extra-data is less than 32 bytes
	ErrInvalidArfIstHeaderExtra = errors.New("invalid arfist header extra-data")
)

type ArfIstExtra struct {
	Validators    []common.Address
	Seal          []byte
	CommittedSeal [][]byte
}

// EncodeRLP serializes ist into the Ethereum RLP format.
func (ist *ArfIstExtra) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, []interface{}{
		ist.Validators,
		ist.Seal,
		ist.CommittedSeal,
	})
}

// DecodeRLP implements rlp.Decoder, and load the arfist fields from a RLP stream.
func (ist *ArfIstExtra) DecodeRLP(s *rlp.Stream) error {
	var arfistExtra struct {
		Validators    []common.Address
		Seal          []byte
		CommittedSeal [][]byte
	}
	if err := s.Decode(&arfistExtra); err != nil {
		return err
	}
	ist.Validators, ist.Seal, ist.CommittedSeal = arfistExtra.Validators, arfistExtra.Seal, arfistExtra.CommittedSeal
	return nil
}

// ExtractArfIstExtra extracts all values of the ArfIstExtra from the header. It returns an
// error if the length of the given extra-data is less than 32 bytes or the extra-data can not
// be decoded.
func ExtractArfIstExtra(h *Header) (*ArfIstExtra, error) {
	if len(h.Extra) < ArfIstExtraVanity {
		return nil, ErrInvalidArfIstHeaderExtra
	}

	var arfistExtra *ArfIstExtra
	err := rlp.DecodeBytes(h.Extra[ArfIstExtraVanity:], &arfistExtra)
	if err != nil {
		return nil, err
	}
	return arfistExtra, nil
}

// ArfIstFilteredHeader returns a filtered header which some information (like seal, committed seals)
// are clean to fulfill the ArfIst hash rules. It returns nil if the extra-data cannot be
// decoded/encoded by rlp.
func ArfIstFilteredHeader(h *Header, keepSeal bool) *Header {
	newHeader := CopyHeader(h)
	arfistExtra, err := ExtractArfIstExtra(newHeader)
	if err != nil {
		return nil
	}

	if !keepSeal {
		arfistExtra.Seal = []byte{}
	}
	arfistExtra.CommittedSeal = [][]byte{}

	payload, err := rlp.EncodeToBytes(&arfistExtra)
	if err != nil {
		return nil
	}

	newHeader.Extra = append(newHeader.Extra[:ArfIstExtraVanity], payload...)

	return newHeader
}
