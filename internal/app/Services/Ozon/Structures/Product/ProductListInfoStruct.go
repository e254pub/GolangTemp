package Product

import "time"

type ProductInfoListBody struct {
	OfferId   []string `json:"offer_id"`
	ProductId []int64  `json:"product_id"`
	Sku       []int64  `json:"sku"`
}

type ProductInfoListResult struct {
	Result struct {
		Items []struct {
			ID               int           `json:"id"`
			Name             string        `json:"name"`
			OfferID          string        `json:"offer_id"`
			Barcode          string        `json:"barcode"`
			BuyboxPrice      string        `json:"buybox_price"`
			CategoryID       int           `json:"category_id"`
			CreatedAt        time.Time     `json:"created_at"`
			Images           []interface{} `json:"images"`
			MarketingPrice   string        `json:"marketing_price"`
			MinPrice         string        `json:"min_price"`
			OldPrice         string        `json:"old_price"`
			PremiumPrice     string        `json:"premium_price"`
			Price            string        `json:"price"`
			RecommendedPrice string        `json:"recommended_price"`
			Sources          []struct {
				IsEnabled bool   `json:"is_enabled"`
				Sku       int    `json:"sku"`
				Source    string `json:"source"`
			} `json:"sources"`
			State  string `json:"state"`
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
				Reasons       struct {
				} `json:"reasons"`
			} `json:"visibility_details"`
			PriceIndex   string        `json:"price_index"`
			Images360    []interface{} `json:"images360"`
			ColorImage   string        `json:"color_image"`
			PrimaryImage string        `json:"primary_image"`
			Status       struct {
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
		} `json:"items"`
	} `json:"result"`
}
