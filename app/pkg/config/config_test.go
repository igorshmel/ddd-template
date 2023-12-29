package config

import (
	"errors"
	"testing"
)

func Test_prepareConfigPathToViper(t *testing.T) {
	type result struct {
		pathToFile string
		filename   string
		extension  string
		err        error
	}
	tests := []struct {
		name string
		arg  string
		res  result
	}{
		{
			name: "ok",
			arg:  "config/dev.config.yaml",
			res: result{
				pathToFile: "config",
				filename:   "dev.config",
				extension:  "yaml",
				err:        nil,
			},
		},

		{
			name: "ok",
			arg:  "dev.config.yaml",
			res: result{
				pathToFile: ".",
				filename:   "dev.config",
				extension:  "yaml",
				err:        nil,
			},
		},

		{
			name: "err - wrong extension",
			arg:  "config/yaml",
			res: result{
				err: errFileExtensionEmpty,
			},
		},

		{
			name: "err - wrong extension",
			arg:  "config.",
			res: result{
				err: errFileExtensionEmpty,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pathToFile, filename, ext, err := prepareConfigPathToViper(tt.arg)
			if !errors.Is(err, tt.res.err) {
				t.Errorf("prepareConfigPathToViper() error = %v, wantErr %v", err, tt.res.err)
				return
			}

			if err == nil {
				if pathToFile != tt.res.pathToFile {
					t.Errorf("prepareConfigPathToViper() pathToFile = %v, want %v", pathToFile, tt.res.pathToFile)
				}

				if filename != tt.res.filename {
					t.Errorf("prepareConfigPathToViper() filename = %v, want %v", filename, tt.res.filename)
				}
				if ext != tt.res.extension {
					t.Errorf("prepareConfigPathToViper() ext = %v, want %v", ext, tt.res.extension)
				}
			}
		})
	}
}
