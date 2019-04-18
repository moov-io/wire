package wire

type FWMValidator struct {
	FWMValidator FedWireMessage
}

// NewFWMValidator returns a new FWMValidator
func NewFWMValidator() FWMValidator {
	fwmValidator := FWMValidator{}
	return fwmValidator
}

// isMandatory validates mandatory tags for a FedWireMessage are defined
func (fwmv *FWMValidator) isMandatory() error {
	if fwmv.FWMValidator.SenderSupplied == nil {
		return fieldError("SenderSupplied", ErrFieldRequired)
	}
	if fwmv.FWMValidator.TypeSubType == nil {
		return fieldError("TypeSubType", ErrFieldRequired)
	}
	if fwmv.FWMValidator.InputMessageAccountabilityData == nil {
		return fieldError("InputMessageAccountabilityData", ErrFieldRequired)
	}
	if fwmv.FWMValidator.Amount == nil {
		return fieldError("Amount", ErrFieldRequired)
	}
	if fwmv.FWMValidator.SenderDepositoryInstitution == nil {
		return fieldError("SenderDepositoryInstitution", ErrFieldRequired)
	}
	if fwmv.FWMValidator.ReceiverDepositoryInstitution == nil {
		return fieldError("ReceiverDepositoryInstitution", ErrFieldRequired)
	}
	if fwmv.FWMValidator.BusinessFunctionCode == nil {
		return fieldError("BusinessFunctionCode", ErrFieldRequired)
	}
	return nil
}
