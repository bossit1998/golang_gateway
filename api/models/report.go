package models

type GetReportModel struct {
	File string `json:"file"`
}

type OperatorReport struct {
	Name                  string  `json:"name"`
	Username              string  `json:"username"`
	Phone                 string  `json:"phone"`
	TotalOrdersCount      uint64  `json:"total_orders_count"`
	AvgPerHour            float32 `json:"average_per_hour"`
	BotOrdersCount        uint64  `json:"bot_orders_count"`
	AdminPanelOrdersCount uint64  `json:"admin_panel_orders_count"`
	AppOrdersCount        uint64  `json:"app_orders_count"`
	WebsiteOrdersCount    uint64  `json:"website_orders_count"`
}

type OperatorsReport struct {
	Reports []OperatorReport `json:"reports"`
}

type BranchReport struct {
	Name                  string  `json:"name"`
	TotalCount            uint64  `json:"total_count"`
	BotOrdersCount        uint64  `json:"bot_orders_count"`
	AdminPanelOrdersCount uint64  `json:"admin_panel_orders_count"`
	AppOrdersCount        uint64  `json:"app_orders_count"`
	WebsiteOrdersCount    uint64  `json:"website_orders_count"`
	TotalSum              float32 `json:"total_sum"`
	TotalSumCash          float32 `json:"total_sum_cash"`
	TotalSumPayme         float32 `json:"total_sum_payme"`
	TotalSumClick         float32 `json:"total_sum_click"`
}

type BranchesReport struct {
	Reports []BranchReport `json:"reports"`
}

type ShipperReport struct {
	TotalOrdersCount      uint64  `json:"total_orders_count"`
	BotOrdersCount        uint64  `json:"bot_orders_count"`
	AdminPanelOrdersCount uint64  `json:"admin_panel_orders_count"`
	AppOrdersCount        uint64  `json:"app_orders_count"`
	WebsiteOrdersCount    uint64  `json:"website_orders_count"`
	TotalSum              float32 `json:"total_sum"`
	TotalSumCash          float32 `json:"total_sum_cash"`
	TotalSumPayme         float32 `json:"total_sum_payme"`
	TotalSumClick         float32 `json:"total_sum_click"`
}
