package eth

import (
	"context"
	"fmt"
	"github.com/defiweb/go-eth/abi"
	//"github.com/defiweb/go-eth/crypto"
	"github.com/defiweb/go-eth/types"
	"math/big"

	"github.com/defiweb/go-eth/rpc"
	"github.com/defiweb/go-eth/rpc/transport"
)

func Connect(ctx context.Context, endpoint string) error {
	// Create transport.
	//
	// There are several other transports available:
	// - HTTP (NewHTTP)
	// - WebSocket (NewWebsocket)
	// - IPC (NewIPC)
	t, err := transport.NewHTTP(transport.HTTPOptions{URL: endpoint})
	if err != nil {
		err = fmt.Errorf("create http transport: %w", err)
		return err
	}

	// Create a JSON-RPC client.
	c, err := rpc.NewClient(rpc.WithTransport(t))
	if err != nil {
		err = fmt.Errorf("create rpc client: %w", err)
		return err
	}

	// Get the latest block number.
	b, err := c.BlockNumber(context.Background())
	if err != nil {
		err = fmt.Errorf("get block number: %w", err)
		return err
	}
	fmt.Println("Latest block number:", b)

	accounts, err := c.Accounts(ctx)
	if err != nil {
		err = fmt.Errorf("get accounts: %w", err)
		return err
	}

	fmt.Println(accounts)

	tx := types.NewTransaction()
	tx.Input = []byte("my test transaction")
	tx.From = types.AddressFromBytesPtr(accounts[0].Bytes())
	tx.To = types.MustAddressFromHexPtr("0x1572E0B1e893E43435cf53C6E54E50942ada9789")
	tx.Value = big.NewInt(1)

	hash, tx, err := c.SendTransaction(ctx, tx)
	if err != nil {
		err = fmt.Errorf("send transaction: %w", err)
		return err
	}

	fmt.Println("send transaction hash:", hash)
	fmt.Println("send transaction tx:", tx)

	return nil

}

//const (
//	contractPath = "envs/ethereum-network/artifacts/contracts/DoorControl.sol/DoorControl.json"
//)

func Mint(ctx context.Context, endpoint, contractPath, pk string) error {
	_, _ = abi.LoadJSON(contractPath)
	contract, err := abi.LoadJSON(contractPath)
	if err != nil {
		err = fmt.Errorf("load contract: %w", err)
		return err
	}

	fmt.Printf("Minting contract at %+v\n", contract)
	fmt.Printf("Minting contract at %+v\n", contract.Methods)
	fmt.Printf("Minting contract at %+v\n", contract.Constructor)

	return nil
}
