package types

// Chain 支持的公链类型
type Chain string

const (
	ChainBTC  Chain = "BTC"
	ChainETH  Chain = "ETH"
	ChainTRON Chain = "TRON"
	ChainHECO Chain = "HECO"
	ChainBSC  Chain = "BSC"
)

func (chain Chain) IsEVMChain() bool {

	return chain == ChainETH || chain == ChainBSC || chain == ChainHECO
}

func (chain Chain) String() string {

	switch chain {

	case ChainBTC:

		return "BTC"
	case ChainETH:

		return "ETH"
	case ChainTRON:

		return "TRON"
	case ChainHECO:

		return "HECO"
	case ChainBSC:

		return "BSC"
	default:
		return "UNKNOWN"
	}
}

func (chain Chain) ChainSymbol() string {

	switch chain {

	case ChainBTC:

		return "BTC"
	case ChainETH:

		return "ETH"
	case ChainTRON:

		return "TRX"
	case ChainHECO:

		return "HT"
	case ChainBSC:

		return "BNB"
	default:
		return "UNKNOWN"
	}
}

func (chain Chain) GetUSDT() string {
	switch chain {

	case ChainBTC:

		return "USDT"
	case ChainETH:

		return "USDT-ETH"
	case ChainBSC:

		return "USDT-BSC"
	case ChainHECO:

		return "USDT-HECO"
	case ChainTRON:

		return "USDT-TRON"
	default:

		return "UNKNOWN"
	}
}

func NewChain(value string) Chain {

	switch value {

	case "BTC":

		return "BTC"
	case "ETH":

		return "ETH"
	case "TRON":

		return "TRON"
	case "HECO":

		return "HECO"
	case "BSC":

		return "BSC"
	default:
		return "UNKNOWN"
	}
}
