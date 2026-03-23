# ==============================================================================
# Makefile helper functions for generate necessary files
#

.PHONY: gen.run
gen.run: gen.errcode

.PHONY: gen.errcode
gen.errcode: gen.errcode.doc

.PHONY: gen.errcode.doc
gen.errcode.doc: tools.verify.codegen
	@echo "===========> Generating error code markdown documentation"
	@codegen -type=int -doc \
		-output ${ROOT_DIR}/docs/guide/zh-CN/api/error_code_generated.md ${ROOT_DIR}/internal/pkg/code