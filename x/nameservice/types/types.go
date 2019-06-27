package types

import(
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Whois struct {
	Value string 	`json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins `json:"price"`
}

//Initial starting price for name that is not owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken",1)}

func NewWhois() Whois {
	return Whois {
		Price: MinNamePrice,
	}
}

func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s 
	Value: %s 
	Price: %s`, w.Owner, w.Value, w.Price))
}
