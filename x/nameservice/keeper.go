package nameservice

import (
	//"encoding/ascii85"
	"github.com/cosmos/cosmos-sdk/codec"  // Tools to work with Cosmos encoding format, Amino
	"github.com/cosmos/cosmos-sdk/x/bank" //Controls accounts and coin transfers

	sdk "github.com/cosmos/cosmos-sdk/types" //Commonly used types through the SDK
)

type Keeper struct {
	coinKeeper bank.Keeper  //allows the module to call functions from bank module

	storeKey sdk.StoreKey  //Gates access to KVStore (persists the state of application i.e Whois structure)

	cdc *codec.Codec   //pointer to Codec used by Amino to encode or decode binary structs
}

//Constructor function
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper {
		coinKeeper: coinKeeper,
		storeKey: storeKey,
		cdc: cdc,
	}
}

//Sets the Whois interface metadata struct for a name
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois Whois){
	if whois.Owner.Empty(){
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))
}

//Gets the Whois metadata for a name
func (k Keeper) GetWhois(ctx sdk.Context, name string ) Whois {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(name)) {
		return NewWhois()
	}
	bz := store.Get([]byte(name))
	var whois Whois
	k.cdc.MustUnmarshalBinaryBare(bz, &whois)
	return whois
}

func (k Keeper) ResolveName (ctx sdk.Context, name string) string {
	return k.GetWhois(ctx,name).Value
}

func (k Keeper) SetName (ctx sdk.Context, name string , value string){
	whois := k.GetWhois(ctx,name)
	whois.Value = value
	k.SetWhois(ctx,name,whois)
}

func (k Keeper) HasOwner (ctx sdk.Context, name string) bool {
	return !k.GetWhois(ctx,name).Owner.Empty()
}

func (k Keeper) GetOwner (ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetWhois(ctx,name).Owner
}

func (k Keeper) SetOwner (ctx sdk.Context, name string, owner sdk.AccAddress){
	whois := k.GetWhois(ctx,name)
	whois.Owner = owner
	k.SetWhois(ctx,name,whois)
}

func (k Keeper) GetPrice (ctx sdk.Context, name string) sdk.Coins {
	return k.GetWhois(ctx,name).Price
}

func (k Keeper) SetPrice (ctx sdk.Context, name string, price sdk.Coins){
	whois := k.GetWhois(ctx,name)
	whois.Price = price
	k.SetWhois(ctx,name,whois)
}

func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}

