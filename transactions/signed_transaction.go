// +build !nosigning

package transactions

import (
	// Stdlib
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"time"

	// RPC
	"github.com/shaunmza/steemgo/encoding/transaction"
	"github.com/shaunmza/steemgo/types"

	// Vendor
	"github.com/pkg/errors"
)

type SignedTransaction struct {
	*types.Transaction
}

func NewSignedTransaction(tx *types.Transaction) *SignedTransaction {
	if tx.Expiration == nil {
		expiration := time.Now().Add(30 * time.Second).UTC()
		tx.Expiration = &types.Time{&expiration}
	}

	return &SignedTransaction{tx}
}

func (tx *SignedTransaction) Serialize() ([]byte, error) {
	var b bytes.Buffer
	encoder := transaction.NewEncoder(&b)

	if err := encoder.Encode(tx.Transaction); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (tx *SignedTransaction) Digest(chain *Chain) ([]byte, error) {
	var msgBuffer bytes.Buffer

	// Write the chain ID.
	rawChainID, err := hex.DecodeString(chain.ID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode chain ID: %v", chain.ID)
	}

	if _, err := msgBuffer.Write(rawChainID); err != nil {
		return nil, errors.Wrap(err, "failed to write chain ID")
	}

	// Write the serialized transaction.
	rawTx, err := tx.Serialize()
	if err != nil {
		return nil, err
	}

	if _, err := msgBuffer.Write(rawTx); err != nil {
		return nil, errors.Wrap(err, "failed to write serialized transaction")
	}

	// Compute the digest.
	digest := sha256.Sum256(msgBuffer.Bytes())
	return digest[:], nil
}

// get rid of cgo and lsecp256k1
func (tx *SignedTransaction) Sign(privKeys [][]byte, chain *Chain) error {
	var buf bytes.Buffer
	chainid, _ := hex.DecodeString(chain.ID)
	//fmt.Println(tx.Operations[0])
	//fmt.Println(" ")
	tx_raw, _ := tx.Serialize()
	//fmt.Println(tx_raw)
	//fmt.Println(" ")
	buf.Write(chainid)
	buf.Write(tx_raw)
	data := buf.Bytes()
	//msg_sha := crypto.Sha256(buf.Bytes())

	var sigsHex []string

	for _, priv_b := range privKeys {
		sigBytes := tx.Sign_Single(priv_b, data)
		sigsHex = append(sigsHex, hex.EncodeToString(sigBytes))
	}

	tx.Transaction.Signatures = sigsHex
	return nil
}
