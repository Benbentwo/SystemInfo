package bootstrap

import (
	_ "github.com/kylelemons/godebug/pretty"
	"reflect"
	"testing"
)

func TestLoadBootstrapConfig(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name    string
		args    args
		want    []Profile
		wantErr bool
	}{
		{
			name: "Load Bootstrap.yaml",
			args: args{inputFile: "data/bootstrap.yaml"},
			want: []Profile{
				{"personal_computer", []Installer{
					{"darwin", "brew", []string{"--cask"}, []string{"spotify", "goland"}, "", ""},
					{"darwin", "custom", []string(nil), []string{"curl my-bash.sh | sh", `echo "hello world"`}, "", "/bin/bash -c"},
				},
				},
				{"work_computer", []Installer{
					{"darwin", "brew", []string{"--cask"}, []string{"spotify", "goland"}, "", ""},
				},
				},
			}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadBootstrapConfig(tt.args.inputFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadBootstrapConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadBootstrapConfig() \n\tgot:\n\t\t%#v\n\twant: \n\t\t%#v\n", got, tt.want)
			}
		})
	}
}
