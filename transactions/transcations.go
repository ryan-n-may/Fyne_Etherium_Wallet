// License: 
package transactions;
import (
	//fmt				"fmt"
	errors			"errors"
	common 			"github.com/ethereum/go-ethereum/common"
	crypto 			"github.com/ethereum/go-ethereum/crypto"
	big		 		"math/big"
	context 		"context"
	ethclient	 	"github.com/ethereum/go-ethereum/ethclient"
	wallet			"wallet/accountops";
	types			"github.com/ethereum/go-ethereum/core/types"
);
/** transfer.go
 *  Author:
 *  Date:
 *  
 **/

 func SuggestGasPrice(url string) *big.Int{
 	client, _ := ethclient.Dial(url);
 	gasPrice_p, _ := client.SuggestGasPrice(context.Background());
 	return gasPrice_p;
 }

func Transfer(a wallet.LocalAccount, url string, eth big.Float, gasLimit uint64, gasPrice big.Int, toAddressHex string) (err error){
	/** Ethereum input value is multiplied by wei 
	 *  gasLimit default is 2100 in UNITS
	 *  gasPrice is set in wei!!!
	 **/
	//success = true;

	client, err := ethclient.Dial(url);
    if err != nil {
        return err;
    }

	if gasLimit == 0{
	 	gasLimit = 21000;
	}
	if gasPrice.Int64() == 0{
		gasPrice_p, _ := client.SuggestGasPrice(context.Background());
		gasPrice = *gasPrice_p;
	}

	// Two ECDSA keys
	privateKey 	:= a.PrivateKey_;
	publicKey 	:= a.PublicKey_;

	fromAddress := crypto.PubkeyToAddress(publicKey);
	if(fromAddress.Hex() != a.AddressHex_){
		err := errors.New("Transfer(): Address generated by public key does not match address stored in LocalAccount struct.")
		return err;
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress);
	if (err != nil){
		return err;
	}

	wei := wallet.ValueToWei(eth);

	toAddress := common.HexToAddress(toAddressHex)

	tx := types.NewTransaction(nonce, toAddress, &wei, gasLimit, &gasPrice, nil);
	chainID, err := client.NetworkID(context.Background());
	if (err != nil){
		return err;
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), &privateKey);
	if (err != nil){
		return err;
	}

	err = client.SendTransaction(context.Background(), signedTx);
	if (err != nil){
		return err;
	}


	return nil;
}