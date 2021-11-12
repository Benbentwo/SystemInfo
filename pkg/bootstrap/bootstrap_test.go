package bootstrap

import (
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
		want    *[]BootstrapConfig
		wantErr bool
	}{
		{
			name: "Load Bootstrap.json",
			args: args{inputFile: "../../test/data/bootstrap.json"},
			want: &[]BootstrapConfig{
				{
					OperatingSystem: []string{"darwin"},
					Name:            "brew",
					Options:         []string{"--cask"},
					Command:         "",
					Packages:        []string{"spotify", "goland"},
				},
				{
					OperatingSystem: []string{"darwin"},
					Name:            "custom",
					Command:         "/bin/bash",
					Options:         []string{"-c"},
					Packages: []string{
						"curl my-bash.sh | sh",
						"echo 'hello world'",
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
				t.Errorf("LoadBootstrapConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
