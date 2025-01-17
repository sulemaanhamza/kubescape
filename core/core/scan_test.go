package core

import (
	"testing"

	"github.com/kubescape/kubescape/v3/core/cautils"
	"github.com/stretchr/testify/assert"
)

func TestIsPrioritizationScanType(t *testing.T) {
	tests := []struct {
		name cautils.ScanTypes
		want bool
	}{
		{
			name: cautils.ScanTypeCluster,
			want: true,
		},
		{
			name: cautils.ScanTypeRepo,
			want: true,
		},
		{
			name: cautils.ScanTypeImage,
			want: false,
		},
		{
			name: cautils.ScanTypeWorkload,
			want: false,
		},
		{
			name: cautils.ScanTypeFramework,
			want: false,
		},
		{
			name: cautils.ScanTypeControl,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(string(tt.name), func(t *testing.T) {
			assert.Equal(t, tt.want, isPrioritizationScanType(tt.name))
		})
	}
}
