package autodetect

import "github.com/hahwul/noir/pkg/models"

func initRails() {
	Patterns = append(Patterns, models.AutoDetect{
		Name: "ruby-rails",
		Patterns: []models.AutoDetectPattern{
			{
				File:  "routes.rb",
				Ext:   ".rb",
				Match: "",
			},
		},
	})
}
