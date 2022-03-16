package main

import (
	"context"
	"fmt"
	"log"

	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/libsv/go-bk/wif"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/sighash"
	"github.com/libsv/go-bt/v2/unlocker"
)

func main() {

	howManyUTXOs := GetInput("Enter the number of funding UTXOs")

	tx := bt.NewTx()

	var wifKeys = []string{}

	for i, _ := strconv.Atoi(howManyUTXOs); i > 0; i-- {

		utxo := GetInput("Enter funding UTXO [TXID index script value wifKey]")

		utxoFields := strings.Fields(utxo)

		txId := utxoFields[0]
		index, _ := strconv.Atoi(utxoFields[1])
		script := utxoFields[2]
		value, _ := strconv.Atoi(utxoFields[3])
		wifKeys = append(wifKeys, utxoFields[4])

		//Add input UTXO
		_ = tx.From(
			txId,
			uint32(index),
			script,
			uint64(value),
		)

	}

	howManyOutputs := GetInput("Enter the number of new UTXOs")

	for i, _ := strconv.Atoi(howManyOutputs); i > 0; i-- {

		utxo := GetInput("Enter the reciever address and how many satoshis to be sent [address satoshis]")

		utxoFields := strings.Fields(utxo)

		satoshis, _ := strconv.Atoi(utxoFields[1])

		_ = tx.AddP2PKHOutputFromAddress(utxoFields[0], uint64(satoshis))
	}

	changeAddress := GetInput("Enter Change Address")

	//Calc Fee
	_ = tx.ChangeToAddress(changeAddress, bt.NewFeeQuote())

	//Unlock input UTXOs
	for index, _ := range wifKeys {

		decodedWif, _ := wif.DecodeWIF(wifKeys[index])
		err := tx.FillInput(
			context.Background(),
			&unlocker.Simple{
				PrivateKey: decodedWif.PrivKey,
			},
			bt.UnlockerParams{uint32(index), sighash.AllForkID},
		)

		if err != nil {
			log.Fatalf(err.Error())
		}

	}

	fmt.Printf("%s\n", tx)

}

func GetInput(prompt string) string {
	input := ""
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(prompt + "\n     ")
		input, _ = reader.ReadString('\n')

		if input != "" {
			break
		}
	}
	return strings.TrimSpace(input)

}
