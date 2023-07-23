# testディレクトリ内にあるインテグレーションテストを実行
test-integration:
	docker-compose exec \
	-e GO_ENV=test \
	api richgo test -v ./test/...

# GodenTestのファイルを更新したいときに使用するコマンド
# Ex) make test-integration-update TARGET=<更新したいテスト関数名>
test-integration-update:
	docker-compose exec \
	-e GO_ENV=test \
	api richgo test -v -run ${TARGET} ./test/... -update
