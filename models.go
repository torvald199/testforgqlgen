package model

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Statistics struct {
	ID             string    `json:"id"`
	CoinID         uint64    `json:"coinId"`
	UserID         uint64    `json:"userId"`
	WorkerID       string    `json:"workerId"`
	DateTimeUpdate time.Time `json:"dateTimeUpdate"`
	DateAdded      time.Time `json:"dateAdded"`
	Hashes         []big.Int `json:"hashes"`
}

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	fmt.Println(">>>>>>начала работу MarshalTimestamp")
	timestamp := t.Unix() * 1000
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(timestamp, 10))
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	fmt.Println(">>>>>>начала работу UnmarshalTimestamp")
	if tmpStr, ok := v.(int); ok {
		return time.Unix(int64(tmpStr), 0), nil
	}
	return time.Time{}, errors.New("Timestamp")
}
