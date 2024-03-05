package handler

import (
	"context"

	"github.com/zhangnan2016/payment-service/common"
	"github.com/zhangnan2016/payment-service/domain/model"
	"github.com/zhangnan2016/payment-service/domain/service"
	payment "github.com/zhangnan2016/payment-service/proto/payment"
)

type Payment struct {
	PaymentDataService service.IPaymentDataService
}

// AddPayment 添加支付记录
func (p Payment) AddPayment(ctx context.Context, request *payment.PaymentInfo,
	response *payment.PaymentID) error {
	addPayment := &model.Payment{}
	if err := common.SwapTo(request, addPayment); err != nil {
		common.Error(err)
	}
	paymentID, err := p.PaymentDataService.AddPayment(addPayment)
	if err != nil {
		common.Error(err)
	}
	response.PaymentId = paymentID
	return nil
}

// UpdatePayment 更新支付信息
func (p Payment) UpdatePayment(ctx context.Context, request *payment.PaymentInfo,
	response *payment.Response) error {
	updatePayment := &model.Payment{}
	if err := common.SwapTo(request, updatePayment); err != nil {
		common.Error(err)
	}
	return p.PaymentDataService.UpdatePayment(updatePayment)
}

// DeletePaymentByID 删除支付信息
func (p Payment) DeletePaymentByID(ctx context.Context, request *payment.PaymentID,
	response *payment.Response) error {
	return p.PaymentDataService.DeletePayment(request.PaymentId)
}

// FindPaymentByID 根据支付单号查询支付信息
func (p Payment) FindPaymentByID(ctx context.Context, request *payment.PaymentID,
	response *payment.PaymentInfo) error {
	findPayment, err := p.PaymentDataService.FindPaymentByID(request.PaymentId)
	if err != nil {
		common.Error(err)
	}
	return common.SwapTo(findPayment, response)
}

// FindAllPayment 查找全部支付订单信息
func (p Payment) FindAllPayment(ctx context.Context, request *payment.All,
	response *payment.PaymentAll) error {
	allPayment, err := p.PaymentDataService.FindAllPayment()
	if err != nil {
		common.Error(err)
	}
	for _, v := range allPayment {
		paymentInfo := &payment.PaymentInfo{}
		if err := common.SwapTo(v, paymentInfo); err != nil {
			common.Error(err)
		}
		response.PaymentInfo = append(response.PaymentInfo, paymentInfo)
	}
	return nil
}
