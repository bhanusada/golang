package xsdtype

//OrderCanonical root tag
type OrderCanonical struct {
	TransactionCode  string `xml:"transactionCode" validate:"required"`
	CurrentTimestamp string `xml:"currentTimestamp" validate:"required"`
	OrderDetails
}

//TransactionCode add or modify
//type TransactionCode struct {
//	TransactionCode string `xml:"transactionCode" validate:"required"`
//}

//CurrentTimestamp timestamp
//type CurrentTimestamp struct {
//	CurrentTimestamp string `xml:"currentTimestamp" validate:"required"`
//}

//OrderDetails order details
type OrderDetails struct {
	OrderDetails GroupOrderDetails `xml:"orderDetails" validate:"required"`
	//OrderHeader    GroupOrderHeader    `xml:"orderHeader"`
	//OrderDiscounts GroupOrderDiscounts `xml:"discounts"`
	//OrderLines     GroupOrderLines     `xml:"orderLines"`
	//OrderSummary   GroupOrderSummary   `xml:"orderSummary"`
}

//GroupOrderDetails group of order details
type GroupOrderDetails struct {
	OrderHeader    GroupOrderHeader    `xml:"orderHeader"`
	OrderDiscounts GroupOrderDiscounts `xml:"discounts"`
	OrderLines     GroupOrderLines     `xml:"orderLines"`
	OrderSummary   GroupOrderSummary   `xml:"orderSummary"`
}

//GroupOrderHeader header details
type GroupOrderHeader struct {
	AccountOrderTypeCode          int16                    `xml:"accountOrderTypeCode"`
	NewStoreFlag                  bool                     `xml:"newStoreFlag" validate:"required"`
	SrcOrderRefID                 int32                    `xml:"srcOrderRefId"`
	RelatedOrders                 []GroupRelatedOrders     `xml:"relatedOrders"`
	ActionType                    string                   `xml:"actionType"`
	TotalOrderRetailAmount        float32                  `xml:"totalOrderRetailAmount"`
	OrderReasonCode               int16                    `xml:"orderReasonCode"`
	ProductDestinationTypeCode    int16                    `xml:"productDestinationTypeCode" validate:"required"`
	OmtOrderRefTypeCode           int16                    `xml:"omtOrderRefTypeCode"`
	TotalOrderWeight              float64                  `xml:"totalOrderWeight"`
	TotalOrderWeightUomCode       int16                    `xml:"totalOrderWeightUomCode"`
	ReceivingLocationNumber       string                   `xml:"receivingLocationNumber" validate:"required"`
	OrderCancelTimestamp          string                   `xml:"orderCancelTimestamp"`
	SpecialOrderFlag              bool                     `xml:"specialOrderFlag" validate:"required"`
	ConsignmentFlag               bool                     `xml:"consignmentFlag" validate:"required"`
	CurrencyTypeCode              string                   `xml:"currencyTypeCode"`
	DistributionServiceTypeCode   int16                    `xml:"distributionServiceTypeCode" validate:"required"`
	OmtOrderTypeCode              int16                    `xml:"omtOrderTypeCode"`
	InboundOrderTypeCode          int16                    `xml:"inboundOrderTypeCode"`
	CommentText                   string                   `xml:"commentText"`
	CreateSystemUserID            string                   `xml:"createSystemUserId"`
	LogID                         string                   `xml:"logId"`
	ShippingLocationTypeCode      string                   `xml:"shippingLocationTypeCode"`
	ReplenishmentMethodCode       int16                    `xml:"replenishmentMethodCode" validate:"required"`
	OrderClass                    int16                    `xml:"orderClass" validate:"required"`
	PaymentTermDetail             GroupPaymentTermDetail   `xml:"paymentTermDetail"`
	UpdatedHeaderFields           GroupUpdatedHeaderFields `xml:"updatedHeaderFields"`
	CreateProgramID               string                   `xml:"createProgramId"`
	OrderClosedTimestamp          string                   `xml:"orderClosedTimestamp"`
	ProductOriginTypeCode         int16                    `xml:"productOriginTypeCode" validate:"required"`
	OrderStatusCode               int16                    `xml:"orderStatusCode" validate:"required"`
	OrderOpenTimestamp            string                   `xml:"orderOpenTimestamp" validate:"required"`
	OrderProductGroupCode         int16                    `xml:"orderProductGroupCode"`
	LastUpdateProgramID           string                   `xml:"lastUpdateProgramId"`
	MerchandisingDepartmentNumber int16                    `xml:"merchandisingDepartmentNumber"`
	OrderRequestingSystemCode     int16                    `xml:"orderRequestingSystemCode" validate:"required"`
	WmsShippingNumber             string                   `xml:"wmsShippingNumber"`
	AllocationID                  int32                    `xml:"allocationId"`
	DocumentTypeCode              string                   `xml:"documentTypeCode"`
	OrderNumber                   int32                    `xml:"orderNumber" validate:"required"`
	OrderCreateTimestamp          string                   `xml:"orderCreateTimestamp" validate:"required"`
	StoreExceptionRequestFlag     bool                     `xml:"storeExceptionRequestFlag"`
	EstimatedShipDate             string                   `xml:"estimatedShipDate"`
	CancelEligibilityDate         string                   `xml:"cancelEligibilityDate"`
	OrderHeaderVersionNumber      int16                    `xml:"orderHeaderVersionNumber" validate:"required"`
	FirstReceiptTimestamp         string                   `xml:"firstReceiptTimestamp"`
	KeyrecTypeIndicator           string                   `xml:"keyrecTypeIndicator"`
	EstimatedOrderCreateDate      string                   `xml:"estimatedOrderCreateDate"`
	ThdAsnID                      int64                    `xml:"thdAsnId"`
	TotalOrderVolume              float64                  `xml:"totalOrderVolume"`
	OrderTypeCode                 string                   `xml:"orderTypeCode"`
	ReceivingLocationTypeCode     string                   `xml:"receivingLocationTypeCode" validate:"required"`
	MerchandisingVendorNumber     int32                    `xml:"merchandisingVendorNumber"`
	ExpectedArrivalDate           string                   `xml:"expectedArrivalDate" validate:"required"`
	LastUpdateSystemUserID        string                   `xml:"lastUpdateSystemUserId"`
	OrderCreateDate               string                   `xml:"orderCreateDate" validate:"required"`
	OrderOpenSystemUserID         string                   `xml:"orderOpenSystemUserId"`
	AddressDetails                []GroupAddressDetails    `xml:"addressDetail"`
	AggExpectedArrivalDate        string                   `xml:"aggExpectedArrivalDate"`
	OrderSourceLocationNumber     int16                    `xml:"orderSourceLocationNumber"`
	TotalOrderCostAmount          string                   `xml:"totalOrderCostAmount"`
	TotalOrderVolumeUomCode       int16                    `xml:"totalOrderVolumeUomCode"`
	OrderTransmissionTypeCode     int16                    `xml:"orderTransmissionTypeCode"`
	FreightTermCode               string                   `xml:"freightTermCode"`
	TransactionPurposeCode        string                   `xml:"transactionPurposeCode" validate:"required"`
	ShippingLocationNumber        string                   `xml:"shippingLocationNumber"`
	LastUpdatedTimestamp          string                   `xml:"lastUpdatedTimestamp"`
	HeaderUpdatedFlag             bool                     `xml:"headerUpdatedFlag" validate:"required"`
	FreightOnboardCode            string                   `xml:"freightOnboardCode"`
	PeggedOrderFlag               bool                     `xml:"peggedOrderFlag" validate:"required"`
	VendorCertifiedFlag           bool                     `xml:"vendorCertifiedFlag"`
	PostReceiptFlag               bool                     `xml:"postReceiptFlag"`
}

//GroupRelatedOrders related details
type GroupRelatedOrders struct {
	RelatedOrderNumber         string `xml:"relatedOrderNumber" validate:"required"`
	RelatedOrderCreateDate     string `xml:"relatedOrderCreateDate" validate:"required"`
	RelatedOrderSourceTypeCode int16  `xml:"relatedOrderSourceTypeCode" validate:"required"`
	OrderRelationTypeCode      int16  `xml:"orderRelationTypeCode" validate:"required"`
}

//GroupPaymentTermDetail payment details
type GroupPaymentTermDetail struct {
	PaymentTermNumber          int16  `xml:"paymentTermNumber" validate:"required"`
	PaymentTermDescription     string `xml:"paymentTermDescription" validate:"required"`
	StoreTermIdentifier        string `xml:"storeTermIdentifier"`
	BasisDateTypeCode          string `xml:"basisDateTypeCode" validate:"required"`
	PaymentTermDiscountPercent string `xml:"paymentTermDiscountPercent" validate:"required"`
	PaymentTermTypeCode        string `xml:"paymentTermTypeCode" validate:"required"`
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
	Discounts []GroupDiscounts `xml:"discount" validate:"required"`
}

//GroupDiscounts discount codes, percent and subdiscount
type GroupDiscounts struct {
	DiscountTypeCode string            `xml:"discountTypeCode" validate:"required"`
	DiscountPercent  string            `xml:"discountPercent" validate:"required"`
	SubDiscounts     GroupSubDiscounts `xml:"subDiscounts"`
}

//GroupSubDiscounts group of subdiscount
type GroupSubDiscounts struct {
	SubDiscounts []GroupSubDiscount `xml:"subDiscount"`
}

//GroupSubDiscount subdiscount code and percent
type GroupSubDiscount struct {
	SubDiscountTypeCode int16  `xml:"subDiscountTypeCode" validate:"required"`
	SubDiscountPercent  string `xml:"subDiscountPercent" validate:"required"`
}

//GroupOrderLines group of order lines
type GroupOrderLines struct {
	OrderLines []GroupOrderLine `xml:"orderLine" validate:"required"`
}

//GroupOrderLine order line details
type GroupOrderLine struct {
	VendorPartNumber                string                     `xml:"vendorPartNumber"`
	VisibilityFeedReceivedTimestamp string                     `xml:"visibilityFeedReceivedTimestamp"`
	FirstReceiptTimestamp           string                     `xml:"firstReceiptTimestamp"`
	LineUpdatedFlag                 bool                       `xml:"lineUpdatedFlag" validate:"required"`
	CreateSystemUserID              string                     `xml:"createSystemUserId"`
	UnitCostAmount                  string                     `xml:"unitCostAmount" validate:"required"`
	SkuNumber                       int32                      `xml:"skuNumber" validate:"required"`
	UpdatedDetailFields             []GroupUpdatedDetailFields `xml:"updatedDetailFields"`
	CreateTimeStamp                 string                     `xml:"createTimeStamp"`
	OnOrderNotShippedQuantity       float64                    `xml:"onOrderNotShippedQuantity" validate:"required"`
	BuyPackWeightUomCode            int16                      `xml:"buyPackWeightUomCode" validate:"required"`
	LastUpdateSystemUserID          string                     `xml:"lastUpdateSystemUserId"`
	OrderLineVersionNumber          int16                      `xml:"orderLineVersionNumber" validate:"required"`
	AdjustedOrderQuantity           float64                    `xml:"adjustedOrderQuantity" validate:"required"`
	HazmatCode                      string                     `xml:"hazmatCode"`
	BuyPackUoiCode                  int16                      `xml:"buyPackUoiCode" validate:"required"`
	InTransitQuantity               float64                    `xml:"inTransitQuantity" validate:"required"`
	HandlingDispositionType         string                     `xml:"handlingDispositionType"`
	ReceivedQuantity                float64                    `xml:"receivedQuantity" validate:"required"`
	BuyPackVolume                   float64                    `xml:"buyPackVolume"`
	IsPrepackSku                    string                     `xml:"isPrepackSku"`
	CostOverrideFlag                bool                       `xml:"costOverrideFlag"`
	UnitRetailAmount                string                     `xml:"unitRetailAmount" validate:"required"`
	SkuShortDescription             string                     `xml:"skuShortDescription" validate:"required"`
	OrderLineStatusCode             int16                      `xml:"orderLineStatusCode"`
	TotalShippedQuantity            float64                    `xml:"totalShippedQuantity" validate:"required"`
	BuyPackWeight                   float64                    `xml:"buyPackWeight" validate:"required"`
	LastUpdateProgramID             string                     `xml:"lastUpdateProgramId"`
	OrderLineNumber                 int16                      `xml:"orderLineNumber" validate:"required"`
	WmsRequestSequenceNumber        int32                      `xml:"wmsRequestSequenceNumber"`
	VisibilityStatusCode            int16                      `xml:"visibilityStatusCode"`
	BuyPackVolumeUomCode            int16                      `xml:"buyPackVolumeUomCode"`
	UnitOfMeasureCode               string                     `xml:"unitOfMeasureCode"`
	SkuDescription                  string                     `xml:"skuDescription" validate:"required"`
	VisibilityLocationCode          int16                      `xml:"visibilityLocationCode"`
	VisibilityExpectedArrivalDate   string                     `xml:"visibilityExpectedArrivalDate"`
	CreateProgramID                 string                     `xml:"createProgramId"`
	LastUpdatedTimestamp            string                     `xml:"lastUpdatedTimestamp"`
	BuyPackSize                     float64                    `xml:"buyPackSize" validate:"required"`
	OriginalOrderQuantity           float64                    `xml:"originalOrderQuantity" validate:"required"`
	AisleBaySequenceID              string                     `xml:"aisleBaySequenceId"`
	UniversalProductCode            string                     `xml:"universalProductCode"`
}

//GroupUpdatedDetailFields group of updated detail fields
type GroupUpdatedDetailFields struct {
	UpdatedDetailFields []string `xml:"updatedDetailField" validate:"required"`
}

//GroupOrderSummary order summary
type GroupOrderSummary struct {
	TotalNumberOfDetailLines int16   `xml:"totalNumberOfDetailLines" validate:"required"`
	TotalOrderWeight         float64 `xml:"totalOrderWeight" validate:"required"`
	TotalOrderVolume         float64 `xml:"totalOrderVolume"`
	TotalOrderQuantity       float64 `xml:"totalOrderQuantity" validate:"required"`
	TotalOrderVolumeUomCode  int16   `xml:"totalOrderVolumeUomCode"`
	TotalNumberOfDiscounts   int16   `xml:"totalNumberOfDiscounts"`
	TotalOrderRetailAmount   string  `xml:"totalOrderRetailAmount" validate:"required"`
	TotalOrderCostAmount     string  `xml:"totalOrderCostAmount" validate:"required"`
	TotalOrderWeightUomCode  int16   `xml:"totalOrderWeightUomCode" validate:"required"`
}
