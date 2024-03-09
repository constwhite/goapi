package api

//Coin balance params
type CoinBalanceParams struct {
	Username string
}

//coin balance repsonse
type CoinBalanceResponse struct {
	//response code
	Code int
	//Account balance
	Balance int64
}

// error repsonse
type Error struct {
	Code    int
	MEssage string
}
