package xsdtype

type OrderCanonical struct {
	TransactionCode
	CurrentTimestamp
	OrderDetails
}

type TransactionCode struct {
	TransactionCode string `xml:"transactionCode" validate:"required"`
}

type CurrentTimestamp struct {
	CurrentTimestamp string `xml:"currentTimestamp"`
}

type OrderDetails struct {
	OrderDetails *GroupOrderDetails `xml:"orderDetails"`
}

type GroupOrderDetails struct {
	OrderHeader    *GroupOrderHeader    `xml:"orderHeader"`
	OrderDiscounts *GroupOrderDiscounts `xml:"discounts"`
	OrderLines     *GroupOrderLines     `xml:"orderLines"`
	OrderSummary   *GroupOrderSummary   `xml:"orderSummary"`
}

type GroupOrderHeader struct {
	AccountOrderTypeCode        int16               `xml:"accountOrderTypeCode"`
	NewStoreFlag                bool                `xml:"newStoreFlag"`
	SrcOrderRefId               xsdt.Int            `xml:"srcOrderRefId"`
	RelatedOrders               *GroupRelatedOrders `xml:"relatedOrders"`
	ActionType                  string              `xml:"actionType"`
	TotalOrderRetailAmount      float32             `xml:"totalOrderRetailAmount"`
	OrderReasonCode             int16               `xml:"orderReasonCode"`
	ProductDestinationTypeCode  int16               `xml:"productDestinationTypeCode"`
	OmtOrderRefTypeCode         int16               `xml:"omtOrderRefTypeCode"`
	TotalOrderWeight            xsdt.Double         `xml:"totalOrderWeight"`
	TotalOrderWeightUomCode     xsdt.Short          `xml:"totalOrderWeightUomCode"`
	ReceivingLocationNumber     xsdt.String         `xml:"receivingLocationNumber"`
	OrderCancelTimestamp        xsdt.DateTime       `xml:"orderCancelTimestamp"`
	SpecialOrderFlag            xsdt.Boolean        `xml:"specialOrderFlag"`
	ConsignmentFlag             xsdt.Boolean        `xml:"consignmentFlag"`
	CurrencyTypeCode            xsdt.String         `xml:"currencyTypeCode"`
	DistributionServiceTypeCode xsdt.Short          `xml:"distributionServiceTypeCode"`
}
