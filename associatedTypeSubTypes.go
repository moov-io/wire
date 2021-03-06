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
var btrTypeSubTypes = associatedTypeSubTypes{
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
var ctrTypeSubTypes = associatedTypeSubTypes{
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

var ctpTypeSubTypes = associatedTypeSubTypes{
	FundsTransfer + BasicFundsTransfer,
	FundsTransfer + RequestReversal,
	FundsTransfer + ReversalTransfer,
	FundsTransfer + RequestReversalPriorDayTransfer,
	FundsTransfer + ReversalPriorDayTransfer,
	ForeignTransfer + BasicFundsTransfer,
	ForeignTransfer + RequestReversal,
	ForeignTransfer + ReversalTransfer,
	ForeignTransfer + RequestReversalPriorDayTransfer,
	ForeignTransfer + ReversalPriorDayTransfer,
	SettlementTransfer + BasicFundsTransfer,
	SettlementTransfer + RequestReversal,
	SettlementTransfer + ReversalTransfer,
	SettlementTransfer + RequestReversalPriorDayTransfer,
	SettlementTransfer + ReversalPriorDayTransfer,
}

// these TypeSubType associations are shared between CKS, DEP, FFR, and FFS
var cksTypeSubTypes = associatedTypeSubTypes{
	SettlementTransfer + BasicFundsTransfer,
	SettlementTransfer + ReversalTransfer,
	SettlementTransfer + ReversalPriorDayTransfer,
}
var depTypeSubTypes = cksTypeSubTypes
var ffrTypeSubTypes = cksTypeSubTypes
var ffsTypeSubTypes = cksTypeSubTypes

var drwTypeSubTypes = associatedTypeSubTypes{
	FundsTransfer + FundsTransferRequestCredit,
	SettlementTransfer + FundsTransferRequestCredit,
}

var drbTypeSubTypes = associatedTypeSubTypes{
	SettlementTransfer + RequestCredit,
	SettlementTransfer + RefusalRequestCredit,
}

var drcTypeSubTypes = associatedTypeSubTypes{
	FundsTransfer + RequestCredit,
	FundsTransfer + RefusalRequestCredit,
}

var svcTypeSubTypes = associatedTypeSubTypes{
	FundsTransfer + RequestReversal,
	FundsTransfer + RequestReversalPriorDayTransfer,
	FundsTransfer + RefusalRequestCredit,
	FundsTransfer + SSIServiceMessage,
	ForeignTransfer + RequestReversal,
	ForeignTransfer + RequestReversalPriorDayTransfer,
	ForeignTransfer + SSIServiceMessage,
	SettlementTransfer + RequestReversal,
	SettlementTransfer + RequestReversalPriorDayTransfer,
	SettlementTransfer + RefusalRequestCredit,
	SettlementTransfer + SSIServiceMessage,
}
