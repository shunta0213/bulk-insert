DATABASE_HOST ?= mysql://dev:dev@localhost:3306
DATABASE_NAME ?= dev

.PHONY: mycli
mycli:
	@mycli -h localhost -u dev -p dev -P 3306 -D dev


.PHONY: migrate-new
MIGRATION_COMMENT ?= $(shell bash -c 'read -p "Comments: " pwd; echo $$pwd')
migrate-new: ## マイグレーションファイル作成
	DATABASE_URL=$(DATABASE_HOST)/$(DATABASE_NAME) dbmate -d migrations -s schema.sql new $(MIGRATION_COMMENT)

.PHONY: migrate-status
migrate-status: ## マイグレーションのステータス確認
	DATABASE_URL=$(DATABASE_HOST)/$(DATABASE_NAME) dbmate -d migrations -s schema.sql status

.PHONY: migrate-up
migrate-up: ## マイグレーションを実行
	DATABASE_URL=$(DATABASE_HOST)/$(DATABASE_NAME) dbmate -d migrations -s schema.sql up

.PHONY: migrate-down
migrate-down: ## マイグレーションをロールバック
	DATABASE_URL=$(DATABASE_HOST)/$(DATABASE_NAME) dbmate -d migrations -s schema.sql down

.PHONY: migrate-drop
migrate-drop: ## マイグレーションを削除
	DATABASE_URL=$(DATABASE_HOST)/$(DATABASE_NAME) dbmate -d migrations -s schema.sql drop


.PHONY: migrate-seed
migrate-seed:
	@go run cmd/seed/main.go
