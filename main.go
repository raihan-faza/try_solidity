package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/raihan-faza/try_solidity/voting"
)

func main() {
	// this is a code to initiate client
	client, err := ethclient.Dial("http://127.0.0.1:1337")
	if err != nil {
		fmt.Printf("failed to dial server, err:%v", err)
		return
	}

	// generating private key, dont ask me how
	privateKey, err := crypto.HexToECDSA("ceritanyaprivatekeytapijangandicontohahahyak")
	if err != nil {
		fmt.Printf("failed to generate privatekey, err:%v", err)
		return
	}

	// generating publich key from private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("public key is not ECDSA")
	}

	// converting publick key into address
	fromAddres := crypto.PubkeyToAddress(*publicKeyECDSA)
	chainID := big.NewInt(31337)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		fmt.Printf("failed to create new transactor, err:%v", err)
	}

	address, tx, contract, err := voting.DeployVoting(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deployed contract at:", address.Hex())
	fmt.Println("Deployment tx:", tx.Hash().Hex())

	// Read candidate
	candidate, err := contract.Candidates(&bind.CallOpts{
		From:    fromAddres,
		Context: context.Background(),
	}, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Candidate: %s, Votes: %d\n", candidate.Name, candidate.VoteCount)

	// Vote
	tx2, err := contract.Vote(auth, big.NewInt(1))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Voted with tx:", tx2.Hash().Hex())
}
