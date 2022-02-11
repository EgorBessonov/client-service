package config

type Config struct {
	TradeServerPort   string `env:"TRADESERVER" envDefault:"localhost:8091"`
	BalanceServerPort string `env:"BALANCESERVER" envDefault:"localhost:8085"`
	UserServerPort    string `env:"USERSERVER" envDefault:"localhost:8087"`
	PriceServerPort   string `env:"PRICESERVER" envDefault:"localhost:8083"`
}
