package gtag

import (
	"context"
	"testing"
)

const testDir = "../../test/internal/"

func TestGenerate(t *testing.T) {
	type args struct {
		ctx   context.Context
		files []string
		types []string
	}
	tests := []struct {
		name    string
		args    args
		want    []*GenerateResult
		wantErr bool
	}{
		{
			name: "reguler",
			args: args{
				ctx:   context.Background(),
				files: []string{testDir + "regular/user.go", testDir + "regular/empty.go"},
				types: []string{"User", "Empty"},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.ctx, tt.args.files, tt.args.types)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Generate() got = %v, want %v", got, tt.want)
			//}
			for _, result := range got {
				t.Log(result)
			}
		})
	}
}
