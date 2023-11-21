package wire

// ValidateOpts contains specific overrides from the default set of validations
type ValidateOpts struct {
	// SkipMandatoryIMAD skips checking that InputMessageAccountabilityData is mandatory tag.
	SkipMandatoryIMAD bool `json:"skipMandatoryIMAD"`

	// AllowMissingSenderSupplied allows the senderSupplied field to be omitted.
	AllowMissingSenderSupplied bool `json:"allowMissingSenderSupplied"`
}
