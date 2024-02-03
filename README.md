# golang create trc20 wallet


example:
```
import 	"github.com/JackGod001/go_trc20_wallet_create"

func main() {
    fmt.Println("生成纯数字随机字符串:")
    wallet, err := walletCreate.GenerateTRCWallet()
	if err != nil {
		t.Error("Test Create TRC20 Address Error:", err)
		return
	}
	fmt.Println("TRC20 Address:", wallet.Address)
	fmt.Println("TRC20 Private Key:", wallet.PrivateKey)
	fmt.Println("TRC20 MNEMONIC:", wallet.Mnemonic)
	fmt.Println("TRC20 Public Key:", wallet.PublicKey)
}


```
