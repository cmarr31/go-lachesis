package genesis

import (
	"time"

	"github.com/Fantom-foundation/go-lachesis/ethapi/common"
	"github.com/Fantom-foundation/go-lachesis/crypto"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/pos"
)

var (
	genesisTestTime = inter.Timestamp(1565000000 * time.Second)
)

type Genesis struct {
	Alloc     Accounts
	Time      inter.Timestamp
	ExtraData []byte
}

// FakeGenesis generates fake genesis with n-nodes.
func FakeGenesis(validators, others int) Genesis {
	accounts := make(Accounts, validators+others)

	for i := 0; i < validators; i++ {
		key := crypto.FakeKey(i)
		addr := crypto.PubkeyToAddress(key.PublicKey)
		accounts[addr] = Account{
			Balance:    pos.StakeToBalance(1000 * 1000 * pos.Qualification),
			PrivateKey: key,
		}
	}
	for i := validators; i < validators+others; i++ {
		key := crypto.FakeKey(i)
		addr := crypto.PubkeyToAddress(key.PublicKey)
		accounts[addr] = Account{
			Balance:    pos.StakeToBalance(pos.Qualification - 1),
			PrivateKey: key,
		}
	}

	return Genesis{
		Alloc: accounts,
		Time:  genesisTestTime,
	}
}

// MainGenesis returns builtin genesis keys of mainnet.
func MainGenesis() Genesis {
	return Genesis{
		Time: genesisTestTime,
		Alloc: Accounts{
			// TODO: fill with official keys and balances before release!
			common.HexToAddress("a123456789123456789123456789012345678901"): Account{Balance: pos.StakeToBalance(1000000)},
			common.HexToAddress("a123456789123456789123456789012345678902"): Account{Balance: pos.StakeToBalance(1000000)},
		},
	}
}

// TestGenesis returns builtin genesis keys of testnet.
func TestGenesis() Genesis {
	return Genesis{
		Time: genesisTestTime,
		Alloc: Accounts{
			// TODO: fill with official keys and balances before release!
			common.HexToAddress("b123456789123456789123456789012345678901"): Account{Balance: pos.StakeToBalance(1000000)},
			common.HexToAddress("b123456789123456789123456789012345678902"): Account{Balance: pos.StakeToBalance(1000000)},
		},
	}
}
