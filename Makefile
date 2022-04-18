.PHONY: models generate

# ==============================================================================
# Go migrate postgresql

# ==============================================================================
# Swagger Models
models:
	$(call print-target)
	swagger generate model -f docs/boss_swagger.yml -m internal/api

generate: models
