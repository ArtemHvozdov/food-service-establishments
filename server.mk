# ===================================
# Статичний сервер для локальної розробки
# ===================================

.PHONY: serve serve-check serve-open

# Порт для сервера
PORT ?= 8000
PUBLIC_DIR ?= public

serve: serve-check
	@if command -v python3 &> /dev/null; then \
		echo "🐍 Запускаю сервер через Python 3 на http://localhost:$(PORT)"; \
		cd $(PUBLIC_DIR) && python3 -m http.server $(PORT); \
	elif command -v python &> /dev/null; then \
		echo "🐍 Запускаю сервер через Python 2 на http://localhost:$(PORT)"; \
		cd $(PUBLIC_DIR) && python -m SimpleHTTPServer $(PORT); \
	elif command -v php &> /dev/null; then \
		echo "🐘 Запускаю сервер через PHP на http://localhost:$(PORT)"; \
		cd $(PUBLIC_DIR) && php -S localhost:$(PORT); \
	elif command -v ruby &> /dev/null; then \
		echo "💎 Запускаю сервер через Ruby на http://localhost:$(PORT)"; \
		cd $(PUBLIC_DIR) && ruby -run -ehttpd . -p $(PORT); \
	elif command -v npx &> /dev/null; then \
		echo "📦 Запускаю сервер через Node.js (http-server) на http://localhost:$(PORT)"; \
		npx http-server ./$(PUBLIC_DIR) -p $(PORT); \
	elif command -v busybox &> /dev/null; then \
		echo "🔧 Запускаю сервер через BusyBox на http://localhost:$(PORT)"; \
		busybox httpd -f -p $(PORT) -h $(shell pwd)/$(PUBLIC_DIR); \
	else \
		echo "❌ Помилка: не знайдено жодного серверу!"; \
		echo "Встановіть один з: Python, PHP, Ruby, Node.js або BusyBox"; \
		exit 1; \
	fi

serve-check:
	@echo "🔍 Перевіряю доступні серверні..."
	@command -v python3 >/dev/null 2>&1 && echo "✓ Python 3 знайдено" || true
	@command -v python >/dev/null 2>&1 && echo "✓ Python знайдено" || true
	@command -v php >/dev/null 2>&1 && echo "✓ PHP знайдено" || true
	@command -v ruby >/dev/null 2>&1 && echo "✓ Ruby знайдено" || true
	@command -v npx >/dev/null 2>&1 && echo "✓ Node.js знайдено" || true
	@command -v busybox >/dev/null 2>&1 && echo "✓ BusyBox знайдено" || true
	@echo ""

# Альтернативна команда з автовідкриттям браузера
serve-open: serve-check
	@if command -v python3 &> /dev/null; then \
		echo "🐍 Запускаю Python 3 на http://localhost:$(PORT)"; \
		(sleep 2 && $(call open_browser)) & \
		cd $(PUBLIC_DIR) && python3 -m http.server $(PORT); \
	elif command -v php &> /dev/null; then \
		echo "🐘 Запускаю PHP на http://localhost:$(PORT)"; \
		(sleep 2 && $(call open_browser)) & \
		cd $(PUBLIC_DIR) && php -S localhost:$(PORT); \
	elif command -v npx &> /dev/null; then \
		echo "📦 Запускаю http-server на http://localhost:$(PORT)"; \
		(sleep 2 && $(call open_browser)) & \
		npx http-server ./$(PUBLIC_DIR) -p $(PORT); \
	else \
		echo "❌ Помилка: не знайдено серверу!"; \
		exit 1; \
	fi

define open_browser
	if command -v xdg-open &> /dev/null; then \
		xdg-open http://localhost:$(PORT); \
	elif command -v open &> /dev/null; then \
		open http://localhost:$(PORT); \
	elif command -v start &> /dev/null; then \
		start http://localhost:$(PORT); \
	fi
endef
