# ===================================
# Головний Makefile проекту
# ===================================

# Підключаємо інші Makefile файли
include server.mk

.PHONY: generate fmt help

# ===================================
# Go команди
# ===================================

generate:
	go run ./cmd/generate/main.go

fmt:
	go fmt ./...

# ===================================
# Допомога
# ===================================

help:
	@echo "📋 Доступні команди:"
	@echo ""
	@echo "  Сервер:"
	@echo "  make serve          - Запустити статичний сервер"
	@echo "  make serve-check    - Перевірити доступні сервери"
	@echo "  make serve-open     - Запустити сервер і відкрити браузер"
	@echo "  make serve PORT=3000 - Запустити на іншому порту"
	@echo ""
	@echo "  Go:"
	@echo "  make generate       - Запустити go run ./cmd/generate/main.go"
	@echo "  make fmt            - Форматувати код (go fmt)"
	@echo ""
	@echo "  Інше:"
	@echo "  make help           - Показати цю справку"
