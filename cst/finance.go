package cst

// ====================== 财务流水状态 ======================

// FinanceStatus 财务流水状态
type FinanceStatus int

const (
	FinanceStatusUnknown FinanceStatus = iota
	FinanceIncome                      // 收入（支付成功）
	FinanceRefund                      // 退款
	FinanceExpense                     // 支出（手续费、平台扣费等）
	FinanceAdjustment                  // 调整（人工调整、优惠等）
)

func (f FinanceStatus) String() string {
	switch f {
	case FinanceIncome:
		return "收入"
	case FinanceRefund:
		return "退款"
	case FinanceExpense:
		return "支出"
	case FinanceAdjustment:
		return "调整"
	default:
		return "未知"
	}
}
