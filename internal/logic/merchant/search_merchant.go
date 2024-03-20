package merchant

import (
	"cloudCustomerService/internal/dao"
	"cloudCustomerService/internal/model"
	"context"
)

func (s *sMerchant) SearchMerchant(ctx context.Context, in *model.MerchantSearchInput) (*model.SearchMerchantOutput, error) {

	out := &model.SearchMerchantOutput{
		Page: in.Page,
		Size: in.Size,
	}
	err := dao.Merchants.Ctx(ctx).Page(in.Page, in.Size).OrderDesc(dao.Merchants.Columns().Id).ScanAndCount(&out.List, &out.Total, false)
	return out, err

}
