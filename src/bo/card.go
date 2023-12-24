package bo

//type BalanceBo struct {
//	Balance  string
//	Unsettle string
//	Record   ConsumptionRecord
//}

type ElectricCharge struct {
	Balance string
	Bonus   string
}

type Allowance struct {
	Type   string
	Num    string
	Unit   string
	Amount string
}

//type ConsumptionRecord = []map[string]string

type BalanceData struct {
	Balance  string            // 账户余额
	Unsettle string            // 未结算金额
	Record   ConsumptionRecord // 消费记录
}

type ConsumptionRecord = []OBJData

type OBJData struct {
	Address     string
	Balance     string
	Cost        string
	Time        string
	Transaction string
}
