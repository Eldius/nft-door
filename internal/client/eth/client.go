package eth

import (
	"context"
	"fmt"
	//"github.com/defiweb/go-eth/crypto"
	"github.com/defiweb/go-eth/types"
	"log"
	"math/big"

	//"github.com/defiweb/go-eth/abi"
	"github.com/defiweb/go-eth/rpc"
	"github.com/defiweb/go-eth/rpc/transport"
	//"github.com/defiweb/go-eth/txmodifier"
	//"github.com/defiweb/go-eth/types"
	//"github.com/defiweb/go-eth/wallet"
	//"log"
	"github.com/ethereum/go-ethereum/crypto"
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

func Mint(ctx context.Context, endpoint, pk string) error {
	privateKey, err := crypto.HexToECDSA(pk) if err != nil { log.Fatalf("Failed to load private key: %v", err)

	return nil
}