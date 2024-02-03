package walletCreate

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/JackGod001/go-bip39"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/keys/hd"
	"golang.org/x/crypto/pbkdf2"
)

type Wallet struct {
	Mnemonic   string
	PrivateKey string
	PublicKey  string
	Address    string
}

func NewSeed(mnemonic string, password string) []byte {
	return pbkdf2.Key([]byte(mnemonic), []byte("mnemonic"+password), 2048, 64, sha512.New)
}

// FromMnemonicSeedAndPassphrase derive form mnemonic and passphrase at index
func FromMnemonicSeedAndPassphrase(mnemonic, passphrase string, index int) (*btcec.PrivateKey, *btcec.PublicKey) {
	seed := NewSeed(mnemonic, passphrase)
	master, ch := hd.ComputeMastersFromSeed(seed, []byte("Bitcoin seed"))
	private, _ := hd.DerivePrivateKeyForPath(
		btcec.S256(),
		master,
		ch,
		fmt.Sprintf("44'/195'/0'/0/%d", index),
	)

	return btcec.PrivKeyFromBytes(private[:])
}
func GenerateTRCWallet() (*Wallet, error) {
	// Generate a mnemonic for a random entropy
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return nil, err
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}
	private, _ := FromMnemonicSeedAndPassphrase(mnemonic, "", 0)
	privateString := hex.EncodeToString(private.Serialize())
	if err != nil {
		return nil, err
	}
	//根据私钥获取公钥
	publicKey := private.PubKey()
	// Serialize the public key in compressed format
	serializedPubKey := publicKey.SerializeCompressed()
	publicKeyHex := hex.EncodeToString(serializedPubKey)

	//转化成16进制字符串
	//根据公钥获取地址
	walletAddress := address.PubkeyToAddress(private.ToECDSA().PublicKey).String()
	fmt.Println("walletAddress:", walletAddress)
	fmt.Println("publicKeyHex:", publicKeyHex)
	fmt.Println("privateString:", privateString)
	fmt.Println("mnemonic:", mnemonic)
	return &Wallet{
		Mnemonic:   mnemonic,
		PrivateKey: privateString, //private.ToECDSA().D.String(),
		PublicKey:  publicKeyHex,
		Address:    walletAddress,
	}, nil
}
