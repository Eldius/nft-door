package ethereum

import (
	"context"
	"fmt"
	"github.com/eldius/nft-pocs/internal/contracts/doorcontrol"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getClient(endpoint string) (*ethclient.Client, error) {
	fmt.Println("connecting to:", endpoint)
	cl, err := ethclient.Dial(endpoint)
	if err != nil {
		err = fmt.Errorf("failed to connect eth client: %v", err)
		return nil, err
	}

	return cl, nil
}

func authenticate(ctx context.Context, cl *ethclient.Client, pk string) (*bind.TransactOpts, error) {
	chainID, err := cl.ChainID(ctx)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return nil, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(ecdsaPrivateKey, chainID)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return nil, err
	}

	return auth, nil
}

func deployContract(ctx context.Context, auth *bind.TransactOpts, cl *ethclient.Client) (common.Address, *types.Transaction, *doorcontrol.DoorControl, error) {
	return doorcontrol.DeployDoorControl(auth, cl, auth.From)
}

func DeployContract(ctx context.Context, endpoint, pk, link string) error {
	return nil
}

func Mint(ctx context.Context, endpoint, pk, link string) error {
	cl, err := getClient(endpoint)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}

	defer cl.Close()

	auth, err := authenticate(ctx, cl, pk)
	if err != nil {
		err = fmt.Errorf("failed to authenticate: %v", err)
		return err
	}

	//address, tx, instance, err := doorcontrol.DeployDoorControl(auth, cl, common.MaxAddress)
	//if err != nil {
	//	err = fmt.Errorf("failed to connect to eth client: %v", err)
	//	return err
	//}
	//
	//fmt.Printf("Contract pending deploy: 0x%x\n", address)
	//fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	//fmt.Printf("Contract pending deploy: 0x%x\n\n", tx.Hash())
	//fmt.Printf("instance: 0x%+v\n\n", instance)

	address, tx, instance, err := deployContract(ctx, auth, cl)

	fmt.Printf("tx: %+v\n\n", tx)
	fmt.Printf("instance: %+v\n\n", instance)
	fmt.Printf("address: %+v\n\n", address)

	session := &doorcontrol.DoorControlSession{
		Contract: instance,
		CallOpts: bind.CallOpts{
			Pending: true,
			Context: ctx,
		},
		TransactOpts: *auth,
	}

	// Call the previous methods without the option parameters
	tx1, err := session.SafeMint(auth.From, link)
	if err != nil {
		err = fmt.Errorf("failed to mint nft: %v", err)
		return err
	}

	fmt.Printf("tx1.Hash(): 0x%x\n", tx1.Hash())
	fmt.Printf("tx1.Value(): 0x%x\n", tx1.Value())

	return nil
}

func mint(instance *doorcontrol.DoorControlSession, auth *bind.TransactOpts, link string) error {
	minted, err := instance.SafeMint(auth.From, link)
	if err != nil {
		err = fmt.Errorf("failed to connect to eth client: %v", err)
		return err
	}

	fmt.Printf("Contract minted: 0x%x\n\n", minted)

	return nil
}
