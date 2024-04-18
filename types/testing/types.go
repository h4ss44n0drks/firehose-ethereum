// Copyright 2021 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ct

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/go-testing-interface"
	pbeth "github.com/streamingfast/firehose-ethereum/types/pb/sf/ethereum/type/v2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type from hexString

func From(in string) from     { return from(newHexString(in)) }
func FromFull(in string) from { return from(newHexString(in, Full, length(20))) }

type to hexString

func To(in string) to     { return to(newHexString(in)) }
func ToFull(in string) to { return to(newHexString(in, Full, length(20))) }

type previousHash hexString

func PreviousHash(in string) previousHash { return previousHash(newHexString(in)) }
func PreviousHashFull(in string) previousHash {
	return previousHash(newHexString(in, Full, length(20)))
}

func Block(t testing.T, blkHash string, components ...interface{}) *pbeth.Block {
	// This is for testing purposes, so it's easier to convey the id and the num from a single element
	ref := newBlockRefFromID(blkHash)

	pbblock := &pbeth.Block{
		Ver:    2,
		Hash:   toBytes(t, ref.ID),
		Number: ref.Number,
	}

	blockTime, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05.5Z")
	require.NoError(t, err)

	pbblock.Header = &pbeth.BlockHeader{
		Hash:       toBytes(t, ref.ID),
		Number:     ref.Number,
		ParentHash: toBytes(t, fmt.Sprintf("%08x%s", pbblock.Number-1, blkHash[8:])),
		Timestamp:  timestamppb.New(blockTime),
	}

	for _, component := range components {
		switch v := component.(type) {
		case *pbeth.TransactionTrace:
			pbblock.TransactionTraces = append(pbblock.TransactionTraces, v)
		case previousHash:
			pbblock.Header.ParentHash = hexString(v).bytes(t)
		case *pbeth.BalanceChange:
			pbblock.BalanceChanges = append(pbblock.BalanceChanges, v)
		case *pbeth.CodeChange:
			pbblock.CodeChanges = append(pbblock.CodeChanges, v)
		default:
			failInvalidComponent(t, "block", component)
		}
	}

	if os.Getenv("DEBUG") != "" {
		out, err := json.MarshalIndent(pbblock, "", "  ")
		require.NoError(t, err)

		// We re-normalize to a plain map[string]interface{} so it's printed as JSON and not a proto default String implementation
		normalizedOut := map[string]interface{}{}
		require.NoError(t, json.Unmarshal([]byte(out), &normalizedOut))

		fmt.Println("created test block", normalizedOut)
	}

	return pbblock
}

type Nonce uint64
type InputData string
type GasLimit uint64
type GasPrice string
type Value string

func (p GasPrice) ToBigInt(t testing.T) *big.Int {
	return Ether(p).ToBigInt(t)
}

func (v Value) ToBigInt(t testing.T) *big.Int {
	return Ether(v).ToBigInt(t)
}

var b1e18 = new(big.Int).Exp(big.NewInt(1), big.NewInt(18), nil)

type Ether string

func (e Ether) ToBigInt(t testing.T) *big.Int {
	in := string(e)
	if strings.HasSuffix(in, " ETH") {
		raw := strings.TrimSuffix(in, " ETH")

		dotIndex := strings.Index(in, ".")
		if dotIndex >= 0 {
			raw = raw[0:dotIndex] + raw[dotIndex+1:]
		}

		if len(raw) < 19 {
			raw = raw + strings.Repeat("0", 19-len(raw))
		} else if len(raw) > 19 {
			raw = raw[0:19]
		}

		out, worked := new(big.Int).SetString(raw, 10)
		require.True(t, worked, "Conversion of %q to big.Int failed", raw)

		return out.Mul(out, b1e18)
	}

	out, worked := new(big.Int).SetString(in, 0)
	require.True(t, worked, "Conversion of %q to big.Int failed", in)
	return out
}

// func Trx(t testing.T, components ...interface{}) *pbeth.Transaction {
// 	trx := &pbeth.Transaction{}
// 	for _, component := range components {
// 		switch v := component.(type) {
// 		case hash:
// 			trx.Hash = hexString(v).bytes(t)
// 		case from:
// 			trx.From = hexString(v).bytes(t)
// 		case to:
// 			trx.To = hexString(v).bytes(t)
// 		case InputData:
// 			trx.Input = toBytes(t, string(v))
// 		case Nonce:
// 			trx.Nonce = uint64(v)
// 		case GasLimit:
// 			trx.GasLimit = uint64(v)
// 		case GasPrice:
// 			trx.GasPrice = pbeth.BigIntFromNative(v.ToBigInt(t))
// 		case Value:
// 			trx.Value = pbeth.BigIntFromNative(v.ToBigInt(t))
// 		default:
// 			failInvalidComponent(t, "trx", component)
// 		}
// 	}

// 	return trx
// }

func TrxTrace(t testing.T, components ...interface{}) *pbeth.TransactionTrace {
	trace := &pbeth.TransactionTrace{
		Receipt: &pbeth.TransactionReceipt{},
	}

	for _, component := range components {
		switch v := component.(type) {
		case hash:
			trace.Hash = hexString(v).bytes(t)
		case from:
			trace.From = hexString(v).bytes(t)
		case to:
			trace.To = hexString(v).bytes(t)
		case GasPrice:
			trace.GasPrice = pbeth.BigIntFromNative(v.ToBigInt(t))
		case Nonce:
			trace.Nonce = uint64(v)
		case *pbeth.Call:
			trace.Calls = append(trace.Calls, v)
		case TrxTraceComponent:
			v.Apply(trace)
		default:
			failInvalidComponent(t, "trx_trace", component)
		}
	}

	return trace
}

func TrxTraceConfig(fn func(trxTrace *pbeth.TransactionTrace)) TrxTraceComponent {
	return trxTraceConfig(fn)
}

type trxTraceConfig func(trxTrace *pbeth.TransactionTrace)

func (c trxTraceConfig) Apply(trxTrace *pbeth.TransactionTrace) {
	c(trxTrace)
}

type TrxTraceComponent interface {
	Apply(trxTrace *pbeth.TransactionTrace)
}

type caller hexString

func Caller(in string) caller     { return caller(newHexString(in)) }
func CallerFull(in string) caller { return caller(newHexString(in, Full, length(20))) }

func Call(t testing.T, components ...interface{}) *pbeth.Call {
	call := &pbeth.Call{}
	for _, component := range components {
		switch v := component.(type) {
		case from:
			call.Caller = hexString(v).bytes(t)
		case caller:
			call.Caller = hexString(v).bytes(t)
		case address:
			call.Address = hexString(v).bytes(t)
		case to:
			call.Address = hexString(v).bytes(t)
		case *pbeth.BalanceChange:
			call.BalanceChanges = append(call.BalanceChanges, v)
		case *pbeth.NonceChange:
			call.NonceChanges = append(call.NonceChanges, v)
		case *pbeth.StorageChange:
			call.StorageChanges = append(call.StorageChanges, v)
		case *pbeth.CodeChange:
			call.CodeChanges = append(call.CodeChanges, v)
		case *pbeth.Log:
			call.Logs = append(call.Logs, v)
		case CallComponent:
			v.Apply(call)
		default:
			failInvalidComponent(t, "call", component)
		}
	}

	if call.Value == nil {
		call.Value = pbeth.BigIntFromNative(big.NewInt(0))
	}

	if call.CallType == pbeth.CallType_UNSPECIFIED {
		call.CallType = pbeth.CallType_CALL
	}

	return call
}

func CallConfig(fn func(call *pbeth.Call)) CallComponent {
	return callConfig(fn)
}

type callConfig func(call *pbeth.Call)

func (c callConfig) Apply(call *pbeth.Call) {
	c(call)
}

type CallComponent interface {
	Apply(call *pbeth.Call)
}

type Ordinal uint64

func BalanceChange(t testing.T, address address, values string, components ...interface{}) *pbeth.BalanceChange {
	datas := strings.Split(values, "/")

	balanceChange := &pbeth.BalanceChange{
		Address: hexString(address).bytes(t),
	}

	toBigIntBytes := func(value string) []byte {
		bigValue, succeed := new(big.Int).SetString(value, 10)
		require.True(t, succeed, "unable to convert value to BigInt")

		return bigValue.Bytes()
	}

	if datas[0] != "" {
		balanceChange.OldValue = pbeth.BigIntFromBytes(toBigIntBytes(datas[0]))
	}

	if datas[1] != "" {
		balanceChange.NewValue = pbeth.BigIntFromBytes(toBigIntBytes(datas[1]))
	}

	for _, component := range components {
		switch v := component.(type) {
		case pbeth.BalanceChange_Reason:
			balanceChange.Reason = v
		case Ordinal:
			balanceChange.Ordinal = uint64(v)
		default:
			failInvalidComponent(t, "balanceChange", component)
		}
	}

	return balanceChange
}

func NonceChange(t testing.T, address address, values string, components ...interface{}) *pbeth.NonceChange {
	datas := strings.Split(values, "/")

	nonceChange := &pbeth.NonceChange{
		Address: hexString(address).bytes(t),
	}

	toUint64 := func(value string) uint64 {
		nonce, err := strconv.ParseUint(value, 10, 64)
		require.NoError(t, err, "unable to convert nonce to uint64")

		return nonce
	}

	if datas[0] != "" {
		nonceChange.OldValue = toUint64(datas[0])
	}

	if datas[1] != "" {
		nonceChange.NewValue = toUint64(datas[1])
	}

	for _, component := range components {
		switch v := component.(type) {
		case Ordinal:
			nonceChange.Ordinal = uint64(v)
		default:
			failInvalidComponent(t, "nonceChange", component)
		}
	}

	return nonceChange
}

func StorageChange(t testing.T, address address, key hash, data string, components ...interface{}) *pbeth.StorageChange {
	datas := strings.Split(data, "/")

	storageChange := &pbeth.StorageChange{
		Address: hexString(address).bytes(t),
		Key:     hexString(key).bytes(t),
	}

	if datas[0] != "" {
		storageChange.OldValue = toFilledBytes(t, datas[0], 32)
	}

	if datas[1] != "" {
		storageChange.NewValue = toFilledBytes(t, datas[1], 32)
	}

	for _, component := range components {
		switch v := component.(type) {
		case Ordinal:
			storageChange.Ordinal = uint64(v)
		default:
			failInvalidComponent(t, "storageChange", component)
		}
	}

	return storageChange
}

func CodeChange(t testing.T, address address, old, new []byte, components ...interface{}) *pbeth.CodeChange {
	codeChange := &pbeth.CodeChange{
		Address: hexString(address).bytes(t),
	}

	if old != nil {
		codeChange.OldHash = sha256Sum(old)
		codeChange.OldCode = old
	}

	if new != nil {
		codeChange.NewHash = sha256Sum(new)
		codeChange.NewCode = new
	}

	for _, component := range components {
		switch v := component.(type) {
		case Ordinal:
			codeChange.Ordinal = uint64(v)
		default:
			failInvalidComponent(t, "codeChange", component)
		}
	}

	return codeChange
}

func sha256Sum(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

type logTopic hexString

func LogTopic(in string) logTopic     { return logTopic(newHexString(in)) }
func LogTopicFull(in string) logTopic { return logTopic(newHexString(in, Full, length(32))) }

type LogData string

func Log(t testing.T, address address, components ...interface{}) *pbeth.Log {
	log := &pbeth.Log{
		Address: hexString(address).bytes(t),
	}

	for _, component := range components {
		switch v := component.(type) {
		case logTopic:
			log.Topics = append(log.Topics, hexString(v).bytes(t))
		case LogData:
			log.Data = toBytes(t, string(v))

		default:
			failInvalidComponent(t, "log", component)
		}
	}

	return log
}

func ToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

type address hexString

func Address(in string) address     { return address(newHexString(in)) }
func AddressFull(in string) address { return address(newHexString(in, Full, length(20))) }

func (a address) Bytes(t testing.T) []byte  { return hexString(a).bytes(t) }
func (a address) String(t testing.T) string { return hexString(a).string(t) }

type hash hexString

func Hash(in string) hash     { return hash(newHexString(in)) }
func HashFull(in string) hash { return hash(newHexString(in, Full, length(32))) }

func (a hash) Bytes(t testing.T) []byte  { return hexString(a).bytes(t) }
func (a hash) String(t testing.T) string { return hexString(a).string(t) }

func toBytes(t testing.T, in string) []byte {
	out, err := hex.DecodeString(sanitizeHex(in))
	require.NoError(t, err)

	return out
}

func toFilledBytes(t testing.T, in string, length int) []byte {
	out := toBytes(t, in)
	if len(out) == length {
		return out
	}

	if len(out) < length {
		copied := make([]byte, length)
		copy(copied, out)
		out = copied
	} else {
		// Necessarly longer
		out = out[0:length]
	}

	return out
}

type expand bool
type length int

const Full expand = true

type hexString struct {
	in     string
	expand bool
	length int
}

func newHexString(in string, opts ...interface{}) (out hexString) {
	out.in = in
	for _, opt := range opts {
		switch v := opt.(type) {
		case expand:
			out.expand = bool(v)
		case length:
			out.length = int(v)
		}
	}
	return
}

func (h hexString) bytes(t testing.T) []byte {
	if h.expand {
		return toFilledBytes(t, h.in, h.length)
	}

	return toBytes(t, h.in)
}

func (h hexString) string(t testing.T) string {
	return hex.EncodeToString(h.bytes(t))
}

type ignoreComponent func(v interface{}) bool

func failInvalidComponent(t testing.T, tag string, component interface{}, options ...interface{}) {
	shouldIgnore := ignoreComponent(func(v interface{}) bool { return false })
	for _, option := range options {
		switch v := option.(type) {
		case ignoreComponent:
			shouldIgnore = v
		}
	}

	if shouldIgnore(component) {
		return
	}

	require.FailNowf(t, "invalid component", "Invalid %s component of type %T", tag, component)
}

type blockRef struct {
	ID     string
	Number uint64
}

func newBlockRefFromID(id string) blockRef {
	if len(id) < 8 {
		return blockRef{id, 0}
	}

	bin, err := hex.DecodeString(string(id)[:8])
	if err != nil {
		return blockRef{id, 0}
	}

	return blockRef{id, uint64(binary.BigEndian.Uint32(bin))}
}

func sanitizeHex(input string) string {
	if has0xPrefix(input) {
		input = input[2:]
	}

	if len(input)%2 != 0 {
		input = "0" + input
	}

	return strings.ToLower(input)
}

func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}
