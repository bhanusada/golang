package xsdtype

//OrderCanonical root tag
type OrderCanonical struct {
	TransactionCode
	CurrentTimestamp
	OrderDetails
}

//TransactionCode add or modify
type TransactionCode struct {
	TransactionCode string `xml:"transactionCode" validate:"required"`
}

//CurrentTimestamp timestamp
type CurrentTimestamp struct {
	CurrentTimestamp string `xml:"currentTimestamp"`
}

//OrderDetails order details
type OrderDetails struct {
	OrderDetails *GroupOrderDetails `xml:"orderDetails"`
}

//GroupOrderDetails group of order details
type GroupOrderDetails struct {
	OrderHeader    *GroupOrderHeader    `xml:"orderHeader"`
	OrderDiscounts *GroupOrderDiscounts `xml:"discounts"`
	OrderLines     *GroupOrderLines     `xml:"orderLines"`
	OrderSummary   *GroupOrderSummary   `xml:"orderSummary"`
}

//GroupOrderHeader header details
type GroupOrderHeader struct {
	AccountOrderTypeCode          int16                     `xml:"accountOrderTypeCode"`
	NewStoreFlag                  bool                      `xml:"newStoreFlag"`
	SrcOrderRefID                 int32                     `xml:"srcOrderRefId"`
	RelatedOrders                 []*GroupRelatedOrders     `xml:"relatedOrders"`
	ActionType                    string                    `xml:"actionType"`
	TotalOrderRetailAmount        float32                   `xml:"totalOrderRetailAmount"`
	OrderReasonCode               int16                     `xml:"orderReasonCode"`
	ProductDestinationTypeCode    int16                     `xml:"productDestinationTypeCode"`
	OmtOrderRefTypeCode           int16                     `xml:"omtOrderRefTypeCode"`
	TotalOrderWeight              float64                   `xml:"totalOrderWeight"`
	TotalOrderWeightUomCode       int16                     `xml:"totalOrderWeightUomCode"`
	ReceivingLocationNumber       string                    `xml:"receivingLocationNumber"`
	OrderCancelTimestamp          string                    `xml:"orderCancelTimestamp"`
	SpecialOrderFlag              bool                      `xml:"specialOrderFlag"`
	ConsignmentFlag               bool                      `xml:"consignmentFlag"`
	CurrencyTypeCode              string                    `xml:"currencyTypeCode"`
	DistributionServiceTypeCode   int16                     `xml:"distributionServiceTypeCode"`
	OmtOrderTypeCode              int16                     `xml:"omtOrderTypeCode"`
	InboundOrderTypeCode          int16                     `xml:"inboundOrderTypeCode"`
	CommentText                   string                    `xml:"commentText"`
	CreateSystemUserID            string                    `xml:"createSystemUserId"`
	LogID                         string                    `xml:"logId"`
	ShippingLocationTypeCode      string                    `xml:"shippingLocationTypeCode"`
	ReplenishmentMethodCode       int16                     `xml:"replenishmentMethodCode"`
	OrderClass                    int16                     `xml:"orderClass"`
	PaymentTermDetail             *GroupPaymentTermDetail   `xml:"paymentTermDetail"`
	UpdatedHeaderFields           *GroupUpdatedHeaderFields `xml:"updatedHeaderFields"`
	CreateProgramID               string                    `xml:"createProgramId"`
	OrderClosedTimestamp          string                    `xml:"orderClosedTimestamp"`
	ProductOriginTypeCode         int16                     `xml:"productOriginTypeCode"`
	OrderStatusCode               int16                     `xml:"orderStatusCode"`
	OrderOpenTimestamp            string                    `xml:"orderOpenTimestamp"`
	OrderProductGroupCode         int16                     `xml:"orderProductGroupCode"`
	LastUpdateProgramID           string                    `xml:"lastUpdateProgramId"`
	MerchandisingDepartmentNumber int16                     `xml:"merchandisingDepartmentNumber"`
	OrderRequestingSystemCode     int16                     `xml:"orderRequestingSystemCode"`
	WmsShippingNumber             string                    `xml:"wmsShippingNumber"`
	AllocationID                  int32                     `xml:"allocationId"`
	DocumentTypeCode              string                    `xml:"documentTypeCode"`
	OrderNumber                   int32                     `xml:"orderNumber"`
	OrderCreateTimestamp          string                    `xml:"orderCreateTimestamp"`
	StoreExceptionRequestFlag     bool                      `xml:"storeExceptionRequestFlag"`
	EstimatedShipDate             string                    `xml:"estimatedShipDate"`
	CancelEligibilityDate         string                    `xml:"cancelEligibilityDate"`
	OrderHeaderVersionNumber      int16                     `xml:"orderHeaderVersionNumber"`
	FirstReceiptTimestamp         string                    `xml:"firstReceiptTimestamp"`
	KeyrecTypeIndicator           string                    `xml:"keyrecTypeIndicator"`
	EstimatedOrderCreateDate      string                    `xml:"estimatedOrderCreateDate"`
	ThdAsnID                      int64                     `xml:"thdAsnId"`
	TotalOrderVolume              float64                   `xml:"totalOrderVolume"`
	OrderTypeCode                 string                    `xml:"orderTypeCode"`
	ReceivingLocationTypeCode     string                    `xml:"receivingLocationTypeCode"`
	MerchandisingVendorNumber     int32                     `xml:"merchandisingVendorNumber"`
	ExpectedArrivalDate           string                    `xml:"expectedArrivalDate"`
	LastUpdateSystemUserID        string                    `xml:"lastUpdateSystemUserId"`
	OrderCreateDate               string                    `xml:"orderCreateDate"`
	OrderOpenSystemUserID         string                    `xml:"orderOpenSystemUserId"`
	AddressDetails                []*GroupAddressDetails    `xml:"addressDetail"`
	AggExpectedArrivalDate        string                    `xml:"aggExpectedArrivalDate"`
	OrderSourceLocationNumber     int16                     `xml:"orderSourceLocationNumber"`
	TotalOrderCostAmount          string                    `xml:"totalOrderCostAmount"`
	TotalOrderVolumeUomCode       int16                     `xml:"totalOrderVolumeUomCode"`
	OrderTransmissionTypeCode     int16                     `xml:"orderTransmissionTypeCode"`
	FreightTermCode               string                    `xml:"freightTermCode"`
	TransactionPurposeCode        string                    `xml:"transactionPurposeCode" validate:"required"`
	ShippingLocationNumber        string                    `xml:"shippingLocationNumber"`
	LastUpdatedTimestamp          string                    `xml:"lastUpdatedTimestamp"`
	HeaderUpdatedFlag             bool                      `xml:"headerUpdatedFlag"`
	FreightOnboardCode            string                    `xml:"freightOnboardCode"`
	PeggedOrderFlag               bool                      `xml:"peggedOrderFlag"`
	VendorCertifiedFlag           bool                      `xml:"vendorCertifiedFlag"`
	PostReceiptFlag               bool                      `xml:"postReceiptFlag"`
}

//GroupRelatedOrders related details
type GroupRelatedOrders struct {
	RelatedOrderNumber         string `xml:"relatedOrderNumber"`
	RelatedOrderCreateDate     string `xml:"relatedOrderCreateDate"`
	RelatedOrderSourceTypeCode int16  `xml:"relatedOrderSourceTypeCode"`
	OrderRelationTypeCode      int16  `xml:"orderRelationTypeCode"`
}

//GroupPaymentTermDetail payment details
type GroupPaymentTermDetail struct {
	PaymentTermNumber          int16  `xml:"paymentTermNumber"`
	PaymentTermDescription     string `xml:"paymentTermDescription"`
	StoreTermIdentifier        string `xml:"storeTermIdentifier"`
	BasisDateTypeCode          string `xml:"basisDateTypeCode"`
	PaymentTermDiscountPercent string `xml:"paymentTermDiscountPercent"`
	PaymentTermTypeCode        string `xml:"paymentTermTypeCode"`
}

//GroupUpdatedHeaderFields group of updated header detail
type GroupUpdatedHeaderFields struct {
	UpdatedHeaderFields []string `xml:"updatedHeaderField"`
}

//GroupAddressDetails address details
type GroupAddressDetails struct {
	AddressLine1 string `xml:"addressLine1"`
	AddressLine2 string `xml:"addressLine2"`
	City         string `xml:"city"`
	StateCode    string `xml:"stateCode"`
	ZipCode      string `xml:"zipCode"`
}

//GroupOrderDiscounts group of discounts
type GroupOrderDiscounts struct {
	Discounts []*GroupDiscounts `xml:"discount"`
}

//GroupDiscounts discount codes, percent and subdiscount
type GroupDiscounts struct {
	DiscountTypeCode string          `xml:"discountTypeCode"`
	DiscountPercent  string          `xml:"discountPercent"`
	SubDiscounts     *GroupDiscounts `xml:"subDiscounts"`
}

//GroupSubDiscounts group of subdiscount
type GroupSubDiscounts struct {
	SubDiscounts []*GroupSubDiscount `xml:"subDiscount"`
}

//GroupSubDiscount subdiscount code and percent
type GroupSubDiscount struct {
	SubDiscountTypeCode int16  `xml:"subDiscountTypeCode"`
	SubDiscountPercent  string `xml:"subDiscountPercent"`
}

//GroupOrderLines group of order lines
type GroupOrderLines struct {
	OrderLines []*GroupOrderLine `xml:"orderLine"`
}

//GroupOrderLine order line details
type GroupOrderLine struct {
	VendorPartNumber                string                      `xml:"vendorPartNumber"`
	VisibilityFeedReceivedTimestamp string                      `xml:"visibilityFeedReceivedTimestamp"`
	FirstReceiptTimestamp           string                      `xml:"firstReceiptTimestamp"`
	LineUpdatedFlag                 bool                        `xml:"lineUpdatedFlag"`
	CreateSystemUserID              string                      `xml:"createSystemUserId"`
	UnitCostAmount                  string                      `xml:"unitCostAmount"`
	SkuNumber                       int32                       `xml:"skuNumber"`
	UpdatedDetailFields             []*GroupUpdatedDetailFields `xml:"updatedDetailFields"`
	CreateTimeStamp                 string                      `xml:"createTimeStamp"`
	OnOrderNotShippedQuantity       float64                     `xml:"onOrderNotShippedQuantity"`
	BuyPackWeightUomCode            int16                       `xml:"buyPackWeightUomCode"`
	LastUpdateSystemUserID          string                      `xml:"lastUpdateSystemUserId"`
	OrderLineVersionNumber          int16                       `xml:"orderLineVersionNumber"`
	AdjustedOrderQuantity           float64                     `xml:"adjustedOrderQuantity"`
	HazmatCode                      string                      `xml:"hazmatCode"`
	BuyPackUoiCode                  int16                       `xml:"buyPackUoiCode"`
	InTransitQuantity               float64                     `xml:"inTransitQuantity"`
	HandlingDispositionType         string                      `xml:"handlingDispositionType"`
	ReceivedQuantity                float64                     `xml:"receivedQuantity"`
	BuyPackVolume                   float64                     `xml:"buyPackVolume"`
	IsPrepackSku                    string                      `xml:"isPrepackSku"`
	CostOverrideFlag                bool                        `xml:"costOverrideFlag"`
	UnitRetailAmount                string                      `xml:"unitRetailAmount"`
	SkuShortDescription             string                      `xml:"skuShortDescription"`
	OrderLineStatusCode             int16                       `xml:"orderLineStatusCode"`
	TotalShippedQuantity            float64                     `xml:"totalShippedQuantity"`
	BuyPackWeight                   float64                     `xml:"buyPackWeight"`
	LastUpdateProgramID             string                      `xml:"lastUpdateProgramId"`
	OrderLineNumber                 int16                       `xml:"orderLineNumber"`
	WmsRequestSequenceNumber        int32                       `xml:"wmsRequestSequenceNumber"`
	VisibilityStatusCode            int16                       `xml:"visibilityStatusCode"`
	BuyPackVolumeUomCode            int16                       `xml:"buyPackVolumeUomCode"`
	UnitOfMeasureCode               string                      `xml:"unitOfMeasureCode"`
	SkuDescription                  string                      `xml:"skuDescription"`
	VisibilityLocationCode          int16                       `xml:"visibilityLocationCode"`
	VisibilityExpectedArrivalDate   string                      `xml:"visibilityExpectedArrivalDate"`
	CreateProgramID                 string                      `xml:"createProgramId"`
	LastUpdatedTimestamp            string                      `xml:"lastUpdatedTimestamp"`
	BuyPackSize                     float64                     `xml:"buyPackSize"`
	OriginalOrderQuantity           float64                     `xml:"originalOrderQuantity"`
	AisleBaySequenceID              string                      `xml:"aisleBaySequenceId"`
	UniversalProductCode            string                      `xml:"universalProductCode"`
}

//GroupUpdatedDetailFields group of updated detail fields
type GroupUpdatedDetailFields struct {
	UpdatedDetailFields []string `xml:"updatedDetailField"`
}

//GroupOrderSummary order summary
type GroupOrderSummary struct {
	TotalNumberOfDetailLines int16   `xml:"totalNumberOfDetailLines"`
	TotalOrderWeight         float64 `xml:"totalOrderWeight"`
	TotalOrderVolume         float64 `xml:"totalOrderVolume"`
	TotalOrderQuantity       float64 `xml:"totalOrderQuantity"`
	TotalOrderVolumeUomCode  int16   `xml:"totalOrderVolumeUomCode"`
	TotalNumberOfDiscounts   int16   `xml:"totalNumberOfDiscounts"`
	TotalOrderRetailAmount   string  `xml:"totalOrderRetailAmount"`
	TotalOrderCostAmount     string  `xml:"totalOrderCostAmount"`
	TotalOrderWeightUomCode  int16   `xml:"totalOrderWeightUomCode"`
}
