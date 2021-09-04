package structures
test
type Expense struct {
	Bill struct {
		State   string `json:"state"`
		Summary struct {
			DueDate          string `json:"due_date"`
			CloseDate        string `json:"close_date"`
			LateInterestRate string `json:"late_interest_rate"`
			PastBalance      int    `json:"past_balance"`
			LateFee          string `json:"late_fee"`
			EffectiveDueDate string `json:"effective_due_date"`
			TotalBalance     int    `json:"total_balance"`
			InterestRate     string `json:"interest_rate"`
			Interest         int    `json:"interest"`
			TotalCumulative  int    `json:"total_cumulative"`
			Paid             int    `json:"paid"`
			MinimumPayment   int    `json:"minimum_payment"`
			OpenDate         string `json:"open_date"`
		} `json:"summary"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Barcode struct {
				Href string `json:"href"`
			} `json:"barcode"`
			BoletoEmail struct {
				Href string `json:"href"`
			} `json:"boleto_email"`
		} `json:"_links"`
		LineItems []struct {
			Amount     int    `json:"amount"`
			Index      int    `json:"index,omitempty"`
			Title      string `json:"title"`
			PostDate   string `json:"post_date"`
			ID         string `json:"id"`
			Href       string `json:"href,omitempty"`
			Category   string `json:"category"`
			Charges    int    `json:"charges,omitempty"`
			TypeDetail string `json:"type_detail,omitempty"`
		} `json:"line_items"`
	} `json:"bill"`
}
