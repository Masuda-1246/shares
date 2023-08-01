package config

import "github.com/Masuda-1246/shares/utils"

func Port() string {
	return utils.GetEnvOrDefault("PORT", "8000")
}
