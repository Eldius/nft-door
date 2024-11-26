package ethereum

import (
	"context"
	"fmt"
	"github.com/eldius/nft-pocs/internal/contracts/doorcontrol"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"time"
)

func Mint(ctx context.Context, endpoint, pk, image string) error {
	fmt.Println("connecting to:", endpoint)
	cl, err := ethclient.Dial(endpoint)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}

	defer cl.Close()

	chainID, err := cl.ChainID(ctx)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainID)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}

	address, tx, instance, err := doorcontrol.DeployDoorControl(auth, cl, common.MaxAddress)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}

	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())

	time.Sleep(250 * time.Millisecond) // Allow it to be processed by the local node :P

	// function call on `instance`. Retrieves pending name
	name, err := instance.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}
	fmt.Println("Pending name:", name)

	dc, err := doorcontrol.NewDoorControl(address, cl)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}

	fmt.Printf("Contract pending deploy: 0x%x\n\n", tx.Hash())
	fmt.Printf("dc: %+v\n", dc)

	session := &doorcontrol.DoorControlSession{
		Contract: dc,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
		TransactOpts: bind.TransactOpts{
			From:     auth.From,
			Signer:   auth.Signer,
			GasLimit: uint64(3141592),
		},
	}

	// Call the previous methods without the option parameters
	tx1, err := session.SafeMint(address, image)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}

	fmt.Println("SafeMint tx1:", tx1)

	return nil
}
