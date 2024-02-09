package wire

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFinancialInstitution_Validate(t *testing.T) {
	tests := []struct {
		desc    string
		fi      FinancialInstitution
		wantErr string
	}{
		{
			desc: "empty model is valid",
			fi:   FinancialInstitution{},
		},
		{
			desc: "IDCode without ID",
			fi: FinancialInstitution{
				IdentificationCode: CHIPSIdentifier,
			},
			wantErr: fieldError("Identifier", ErrFieldRequired).Error(),
		},
		{
			desc: "ID without IDCode",
			fi: FinancialInstitution{
				Identifier: "someIdentifier",
			},
			wantErr: fieldError("IdentificationCode", ErrFieldRequired).Error(),
		},
		{
			desc: "invalid chars in name",
			fi: FinancialInstitution{
				Name: "ℯⰰ",
			},
			wantErr: fieldError("Name", ErrNonAlphanumeric, "ℯⰰ").Error(),
		},
		{
			desc: "invalid chars in address 1",
			fi: FinancialInstitution{
				Address: Address{
					AddressLineOne: "ℯⰰ",
				},
			},
			wantErr: fieldError("AddressLineOne", ErrNonAlphanumeric, "ℯⰰ").Error(),
		},
		{
			desc: "invalid chars in address 2",
			fi: FinancialInstitution{
				Address: Address{
					AddressLineTwo: "ℯⰰ",
				},
			},
			wantErr: fieldError("AddressLineTwo", ErrNonAlphanumeric, "ℯⰰ").Error(),
		},
		{
			desc: "invalid chars in address 3",
			fi: FinancialInstitution{
				Address: Address{
					AddressLineThree: "ℯⰰ",
				},
			},
			wantErr: fieldError("AddressLineThree", ErrNonAlphanumeric, "ℯⰰ").Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := tt.fi.Validate()

			if tt.wantErr != "" {
				require.ErrorContains(t, got, tt.wantErr)
			} else {
				require.NoError(t, got)
			}
		})
	}
}
