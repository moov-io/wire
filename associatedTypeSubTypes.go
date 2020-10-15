package wire

// associatedTypeSubTypes models the TypeSubType associations from
// section 15 (Business Function Code Reference) of the Format Reference Guide
type associatedTypeSubTypes []string

// Contains returns true if the target typeSubType string is in the associations list
func (t associatedTypeSubTypes) Contains(target string) bool {
	for _, typeSubType := range t {
		if typeSubType == target {
			return true
		}
	}
	return false
}

// btrTypeSubTypes contains the types/subtypes associated with a BankTransfer BusinessFunctionCode
var btrTypeSubTypes associatedTypeSubTypes = []string{
	FundsTransfer + BasicFundsTransfer,
	FundsTransfer + ReversalTransfer,
	FundsTransfer + ReversalPriorDayTransfer,
	ForeignTransfer + BasicFundsTransfer,
	ForeignTransfer + ReversalTransfer,
	ForeignTransfer + ReversalPriorDayTransfer,
	SettlementTransfer + BasicFundsTransfer,
	SettlementTransfer + ReversalTransfer,
	SettlementTransfer + ReversalPriorDayTransfer,
}

// ctrTypeSubTypes contains the types/subtypes associated with a CustomerTransfer BusinessFunctionCode
var ctrTypeSubTypes associatedTypeSubTypes = []string{
	FundsTransfer + BasicFundsTransfer,
	FundsTransfer + ReversalTransfer,
	FundsTransfer + ReversalPriorDayTransfer,
	ForeignTransfer + BasicFundsTransfer,
	ForeignTransfer + ReversalTransfer,
	ForeignTransfer + ReversalPriorDayTransfer,
	SettlementTransfer + BasicFundsTransfer,
	SettlementTransfer + ReversalTransfer,
	SettlementTransfer + ReversalPriorDayTransfer,
}
