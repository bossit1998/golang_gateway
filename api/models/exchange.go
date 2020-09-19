package models

type GetExchangeRatesModel struct {
  Title string `json:"title"`
  Code string `json:"code"`
  CBPrice string `json:"cb_price"`
  NBUBuyPrice string `json:"nbu_buy_price"`
  NBUCellPrice string `json:"nbu_cell_price"`
  Date string `json:"date"`
}
