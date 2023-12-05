package transformer

import (
	"github.com/hernanijr/imersao13/go/internal/market/dto"
	"github.com/hernanijr/imersao13/go/internal/market/entity"
)

func TransformInput(input dto.TradeInput) *entity.Order {
	asset := entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor := entity.NewInvestor(input.InvestorID)
	order := entity.NewOrder(input.OrderID, investor, asset, input.Shares, input.CurrentShares, input.Price, input.OrderType)

	if input.CurrentShares > 0 {
		assetPosition := entity.NewInvestorAssetPosition(input.AssetID, input.CurrentShares)
		investor.AddAssetPosition(assetPosition)
	}

	return order

}

func TransformOutput(order *entity.Order) *dto.OrderOutput {
	output := &dto.OrderOutput{
		OrderID:    order.ID,
		InvestorID: order.Investor.ID,
		AssetID:    order.Asset.ID,
		OrderType:  order.OrderType,
		Status:     order.Status,
		Partial:    order.PendingShares,
		Shares:     order.Shares,
	}

	var transactionsOutput []*dto.TransactionOutput
	for _, transaction := range order.Transactions {
		transactionsOutput = append(transactionsOutput, &dto.TransactionOutput{
			TransactionID: transaction.ID,
			BuyerID:       transaction.BuyingOrder.ID,
			SellerID:      transaction.SellingOrder.ID,
			AssetID:       transaction.SellingOrder.Asset.ID,
			Price:         transaction.Price,
			Shares:        transaction.SellingOrder.Shares - transaction.SellingOrder.PendingShares,
		})
	}

	output.TransactionOutput = transactionsOutput

	return output
}