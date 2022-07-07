package Product

import "time"

type ProductInfoBody struct {
	OfferId   string `json:"offer_id"`
	ProductId int64  `json:"product_id"`
	Sku       int64  `json:"sku"`
}

type ProductInfoResult struct {
	Result struct {
		ID               int       `json:"id"`
		Name             string    `json:"name"`
		OfferID          string    `json:"offer_id"`
		Barcode          string    `json:"barcode"`
		BuyboxPrice      string    `json:"buybox_price"`
		CategoryID       int       `json:"category_id"`
		CreatedAt        time.Time `json:"created_at"`
		Images           []string  `json:"images"`
		MarketingPrice   string    `json:"marketing_price"`
		MinOzonPrice     string    `json:"min_ozon_price"`
		OldPrice         string    `json:"old_price"`
		PremiumPrice     string    `json:"premium_price"`
		Price            string    `json:"price"`
		RecommendedPrice string    `json:"recommended_price"`
		MinPrice         string    `json:"min_price"`
		Sources          []struct {
			IsEnabled bool   `json:"is_enabled"`
			Sku       int    `json:"sku"`
			Source    string `json:"source"`
		} `json:"sources"`
		Stocks struct {
			Coming   int `json:"coming"`
			Present  int `json:"present"`
			Reserved int `json:"reserved"`
		} `json:"stocks"`
		Errors            []interface{} `json:"errors"`
		Vat               string        `json:"vat"`
		Visible           bool          `json:"visible"`
		VisibilityDetails struct {
			HasPrice      bool `json:"has_price"`
			HasStock      bool `json:"has_stock"`
			ActiveProduct bool `json:"active_product"`
		} `json:"visibility_details"`
		PriceIndex  string `json:"price_index"`
		Commissions []struct {
			Percent        int     `json:"percent"`
			MinValue       int     `json:"min_value"`
			Value          float64 `json:"value"`
			SaleSchema     string  `json:"sale_schema"`
			DeliveryAmount int     `json:"delivery_amount"`
			ReturnAmount   int     `json:"return_amount"`
		} `json:"commissions"`
		VolumeWeight        float64       `json:"volume_weight"`
		IsPrepayment        bool          `json:"is_prepayment"`
		IsPrepaymentAllowed bool          `json:"is_prepayment_allowed"`
		Images360           []interface{} `json:"images360"`
		ColorImage          string        `json:"color_image"`
		PrimaryImage        string        `json:"primary_image"`
		Status              struct {
			State            string        `json:"state"`
			StateFailed      string        `json:"state_failed"`
			ModerateStatus   string        `json:"moderate_status"`
			DeclineReasons   []interface{} `json:"decline_reasons"`
			ValidationState  string        `json:"validation_state"`
			StateName        string        `json:"state_name"`
			StateDescription string        `json:"state_description"`
			IsFailed         bool          `json:"is_failed"`
			IsCreated        bool          `json:"is_created"`
			StateTooltip     string        `json:"state_tooltip"`
			ItemErrors       []interface{} `json:"item_errors"`
			StateUpdatedAt   time.Time     `json:"state_updated_at"`
		} `json:"status"`
		State       string `json:"state"`
		ServiceType string `json:"service_type"`
		FboSku      int    `json:"fbo_sku"`
		FbsSku      int    `json:"fbs_sku"`
	} `json:"result"`
}
