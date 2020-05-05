package gtag

import (
	"context"
	"testing"
)

const testDir = "../../test/internal/"

func TestGenerate(t *testing.T) {
	type args struct {
		ctx  context.Context
		file string
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "regular",
			args: args{
				ctx:  context.Background(),
				file: testDir + "regular/user.go",
				name: "User",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "empty",
			args: args{
				ctx:  context.Background(),
				file: testDir + "regular/empty.go",
				name: "Empty",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args.ctx, tt.args.file, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("Generate() got = %s, want %s", got, tt.want)
			//}
			t.Logf("%s", got)
			if err := got.Commit(); err != nil {
				t.Error(err)
			}
		})
	}
}
