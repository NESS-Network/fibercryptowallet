package skycoin

import (
	"fmt"
	"strconv"

	"github.com/fibercrypto/FiberCryptoWallet/src/core"
	"github.com/fibercrypto/FiberCryptoWallet/src/util"
	"github.com/skycoin/skycoin/src/readable"
)

/*
SkycoinPendingTransaction
*/
type SkycoinPendingTransaction struct { //Implements Transaction interface
	Transaction readable.UnconfirmedTransactionVerbose
}

func (txn *SkycoinPendingTransaction) SupportedAssets() []string {
	return []string{Sky, CoinHour, CalculatedHour}
}

func (txn *SkycoinPendingTransaction) GetTimestamp() core.Timestamp {
	return core.Timestamp(txn.Transaction.Received.Unix())
}

func (txn *SkycoinPendingTransaction) GetStatus() core.TransactionStatus {
	return core.TXN_STATUS_PENDING
}

func (txn *SkycoinPendingTransaction) GetInputs() []core.TransactionInput {
	inputs := make([]core.TransactionInput, 0)
	for _, input := range txn.Transaction.Transaction.In {
		inputs = append(inputs, &SkycoinPendingTransactionInput{Input: input})
	}
	return inputs
}

func (txn *SkycoinPendingTransaction) GetOutputs() []core.TransactionOutput {
	outputs := make([]core.TransactionOutput, 0)
	for _, output := range txn.Transaction.Transaction.Out {
		outputs = append(outputs, &SkycoinPendingTransactionOutput{Output: output})
	}
	return outputs
}

func (txn *SkycoinPendingTransaction) GetId() string {
	return txn.Transaction.Transaction.Hash
}

func (txn *SkycoinPendingTransaction) ComputeFee(ticker string) (uint64, error) {
	if ticker == CoinHour {
		return txn.Transaction.Transaction.Fee, nil
	} else if util.StringInList(ticker, txn.SupportedAssets()) {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

/**
 * SkycoinTransactionIterator
 */
type SkycoinTransactionIterator struct { //Implements TransactionIterator interface
	Current      int
	Transactions []core.Transaction
}

func (it *SkycoinTransactionIterator) Value() core.Transaction {
	return it.Transactions[it.Current]
}

func (it *SkycoinTransactionIterator) Next() bool {
	if it.HasNext() {
		it.Current++
		return true
	}
	return false
}

func (it *SkycoinTransactionIterator) HasNext() bool {
	return (it.Current + 1) < len(it.Transactions)
}

func NewSkycoinTransactionIterator(transactions []core.Transaction) *SkycoinTransactionIterator {
	return &SkycoinTransactionIterator{Transactions: transactions, Current: -1}
}

/**
 * SkycoinPendingTransactionInput
 */
type SkycoinPendingTransactionInput struct { //Implements TransactionInput interface
	Input readable.TransactionInput
}

func (in *SkycoinPendingTransactionInput) GetId() string {
	return ""
}

func (in *SkycoinPendingTransactionInput) IsSpent() bool {
	return true
}

func (in *SkycoinPendingTransactionInput) GetSpentOutput() core.TransactionOutput {
	return nil
}

func (in *SkycoinPendingTransactionInput) GetCoins(ticker string) (uint64, error) {
	if util.StringInList(ticker, []string{Sky, CoinHour, CalculatedHour}) {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

/**
 * SkycoinPendingTransactionOutput
 */
type SkycoinPendingTransactionOutput struct { //Implements TransactionOutput interface
	Output readable.TransactionOutput
}

func (sto *SkycoinPendingTransactionOutput) GetId() string {
	return sto.Output.Hash
}

func (sto *SkycoinPendingTransactionOutput) IsSpent() bool {
	return false
}

func (sto *SkycoinPendingTransactionOutput) GetAddress() core.Address {
	return &SkycoinAddress{address: sto.Output.Address}
}

func (sto *SkycoinPendingTransactionOutput) GetCoins(ticker string) (uint64, error) {
	accuracy, err := util.AltcoinQuotient(ticker)
	if err != nil {
		return uint64(0), err
	}
	if ticker == Sky {
		coin, err2 := strconv.ParseFloat(sto.Output.Coins, 64)
		if err2 != nil {
			return uint64(0), err2
		}
		return uint64(coin * float64(accuracy)), nil
	} else if ticker == CoinHour {
		return sto.Output.Hours * accuracy, nil
	} else if ticker == CalculatedHour {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

/**
 * SkycoinTransactionOutputIterator
 */
type SkycoinTransactionOutputIterator struct { //Implements TransactionOutputIterator interface
	Current int
	Outputs []core.TransactionOutput
}

func (it *SkycoinTransactionOutputIterator) Value() core.TransactionOutput {
	return it.Outputs[it.Current]
}

func (it *SkycoinTransactionOutputIterator) Next() bool {
	if it.HasNext() {
		it.Current++
		return true
	}
	return false
}

func (it *SkycoinTransactionOutputIterator) HasNext() bool {
	return (it.Current + 1) < len(it.Outputs)
}

func NewSkycoinTransactionOutputIterator(outputs []core.TransactionOutput) *SkycoinTransactionOutputIterator {
	return &SkycoinTransactionOutputIterator{Outputs: outputs, Current: -1}
}

/*
SkycoinTransaction
*/
type SkycoinTransaction struct {
	skyTxn  readable.TransactionVerbose
	status  core.TransactionStatus
	inputs  []core.TransactionInput
	outputs []core.TransactionOutput
}

func (txn *SkycoinTransaction) SupportedAssets() []string {
	return []string{Sky, CoinHour, CalculatedHour}
}

func (txn *SkycoinTransaction) GetTimestamp() core.Timestamp {
	return core.Timestamp(txn.skyTxn.Timestamp)

}

func (txn *SkycoinTransaction) GetStatus() core.TransactionStatus {

	if txn.status == core.TXN_STATUS_CONFIRMED {
		return txn.status
	}

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return 0
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	txnU, _ := c.Transaction(txn.skyTxn.Hash)
	if txnU.Status.Confirmed {
		txn.status = core.TXN_STATUS_CONFIRMED
		return txn.status
	}
	txn.status = core.TXN_STATUS_PENDING
	return txn.status
}

func (txn *SkycoinTransaction) GetInputs() []core.TransactionInput {
	if txn.inputs == nil {

		c, err := NewSkycoinApiClient(PoolSection)
		if err != nil {
			return nil
		}
		defer core.GetMultiPool().Return(PoolSection, c)
		transaction, err := c.TransactionVerbose(txn.skyTxn.Hash)
		if err != nil {
			return nil
		}
		txn.inputs = make([]core.TransactionInput, 0)
		for _, in := range transaction.Transaction.In {
			txn.inputs = append(txn.inputs, &SkycoinTransactionInput{
				skyIn:       in,
				spentOutput: nil,
			})
		}

	}

	return txn.inputs
}

func (txn *SkycoinTransaction) GetOutputs() []core.TransactionOutput {
	if txn.outputs == nil {
		txn.outputs = make([]core.TransactionOutput, 0)
		for _, out := range txn.skyTxn.Out {
			txn.outputs = append(txn.outputs, &SkycoinTransactionOutput{
				skyOut: out,
				spent:  false,
			})
		}
	}
	return txn.outputs
}

func (txn *SkycoinTransaction) GetId() string {
	return txn.skyTxn.Hash
}

func (txn *SkycoinTransaction) ComputeFee(ticker string) (uint64, error) {
	if ticker == CoinHour {
		return txn.skyTxn.Fee, nil
	} else if util.StringInList(ticker, txn.SupportedAssets()) {
		return uint64(0), nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

type SkycoinTransactionInput struct {
	skyIn       readable.TransactionInput
	spentOutput *SkycoinTransactionOutput
}

func (in *SkycoinTransactionInput) GetId() string {
	return in.skyIn.Hash
}

func (in *SkycoinTransactionInput) GetSpentOutput() core.TransactionOutput {
	if in.spentOutput == nil {

		c, err := NewSkycoinApiClient(PoolSection)
		if err != nil {
			return nil
		}
		defer core.GetMultiPool().Return(PoolSection, c)
		out, err := c.UxOut(in.skyIn.Hash)
		if err != nil {
			return nil
		}
		skyAccuracy, _ := util.AltcoinQuotient("SKY")

		skyOut := &SkycoinTransactionOutput{
			skyOut: readable.TransactionOutput{
				Address: out.OwnerAddress,
				Coins:   strconv.FormatFloat(float64(out.Coins)/float64(skyAccuracy), 'f', -1, 64),
				Hours:   out.Hours,
				Hash:    out.Uxid,
			},
			spent: true}
		in.spentOutput = skyOut

	}
	return in.spentOutput

}

func (in *SkycoinTransactionInput) GetCoins(ticker string) (uint64, error) {
	accuracy, err2 := util.AltcoinQuotient(ticker)
	if err2 != nil {
		return uint64(0), err2
	}
	if ticker == Sky {
		skyf, err := strconv.ParseFloat(in.skyIn.Coins, 64)
		if err != nil {
			return 0, err
		}
		return uint64(skyf * float64(accuracy)), nil
	} else if ticker == CoinHour {
		return in.skyIn.Hours * accuracy, nil
	} else if ticker == CalculatedHour {
		return in.skyIn.CalculatedHours * accuracy, nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

/**
 * SkycoinTransactionInputIterator
 */
type SkycoinTransactionInputIterator struct {
	current int
	data    []core.TransactionInput
}

func (iter *SkycoinTransactionInputIterator) Value() core.TransactionInput {
	return iter.data[iter.current]
}

func (iter *SkycoinTransactionInputIterator) Next() bool {
	if iter.HasNext() {
		iter.current++
		return true
	}
	return false
}

func (iter *SkycoinTransactionInputIterator) HasNext() bool {
	return (iter.current + 1) < len(iter.data)
}

func NewSkycoinTransactioninputIterator(ins []core.TransactionInput) *SkycoinTransactionInputIterator {
	return &SkycoinTransactionInputIterator{data: ins, current: -1}
}

/**
 * SkycoinTransactionOutput
 */
type SkycoinTransactionOutput struct {
	skyOut          readable.TransactionOutput
	spent           bool
	calculatedHours uint64
}

func (out *SkycoinTransactionOutput) GetId() string {
	return out.skyOut.Hash

}

func (out *SkycoinTransactionOutput) GetAddress() core.Address {
	return &SkycoinAddress{address: out.skyOut.Address}
}

func (out *SkycoinTransactionOutput) GetCoins(ticker string) (uint64, error) {
	accuracy, err2 := util.AltcoinQuotient(ticker)
	if err2 != nil {
		return uint64(0), err2
	}
	if ticker == Sky {
		skyf, err := strconv.ParseFloat(out.skyOut.Coins, 64)
		if err != nil {
			return 0, err
		}
		return uint64(skyf * float64(accuracy)), nil
	} else if ticker == CoinHour {
		return out.skyOut.Hours * accuracy, nil
	} else if ticker == CalculatedHour {
		return out.calculatedHours * accuracy, nil
	}
	return uint64(0), fmt.Errorf("Invalid ticker %v\n", ticker)
}

func (out *SkycoinTransactionOutput) IsSpent() bool {
	if out.spent {
		return true
	}

	c, err := NewSkycoinApiClient(PoolSection)
	if err != nil {
		return true
	}
	defer core.GetMultiPool().Return(PoolSection, c)
	ou, err := c.UxOut(out.skyOut.Hash)
	if err != nil {
		return false
	}
	if ou.SpentTxnID != "0000000000000000000000000000000000000000000000000000000000000000" {
		out.spent = true
		return true
	}
	return false
}