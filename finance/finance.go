// pulls data from yahoo finance and returns data in json format
package finance

// http://localhost:3000/api/v1/book?quotes=BB.TO,ACB.TO
import (
	"github.com/gofiber/fiber"
	"github.com/piquette/finance-go/quote"
	"fmt"
)


// from golang serverless function
type Message struct {
	Data [][]string `json:"data"`
	Columns [8]string `json:"columns"`
	Index []string `json:"index"`
}

func GetTickersPandas(c *fiber.Ctx) {
	
	quotes := queryMulti(c, "quotes")
	iter := quote.List(quotes)
	var stock_data [][]string
	var used_symbols []string
	// Iterate over results. Will exit upon any error.
	columns := [8]string{"Last Price", 
	"Change", "Volume", "Avg Vol (3 Month)", "Vol Ratio", 
	"Dollar", "Market", "Exchange"}
	for iter.Next() {
		q := iter.Quote()
		volume_ratio := float64(q.RegularMarketVolume) / float64(q.AverageDailyVolume3Month)
		used_symbols = append(used_symbols, q.Symbol)
		stock_data = append(stock_data, []string{ 
			fmt.Sprintf("%2.2f", q.RegularMarketPrice),   
			fmt.Sprintf("%2.2f", q.RegularMarketChangePercent), 
			fmt.Sprintf("%d", q.RegularMarketVolume),  
			fmt.Sprintf("%d", q.AverageDailyVolume3Month),
			fmt.Sprintf("%2.2f", volume_ratio),
			q.CurrencyID,
			q.MarketID,
			q.ExchangeID})
	}

	c.JSON(Message{Data: stock_data, Index: used_symbols, Columns: columns})
}
