# workspace

docker コマンドを実行するラッパーです

# 開発

	$ ./run.sh [サブコマンド]

# ビルド

	$ make

# config.yaml 例

	docker:
	  image: 'mamemomonga/workspaces:debian'
	  container: 'ws-debian_1'
	
	volume:
	  name:  'ws-debian_1'
	  # 先頭に/があると絶対パス
	  # なければカレントディレクトリからの相対パス
	  mount: 'home/app'
	
	# 公開ポート番号（省略可能）
	ports:
	  - '80:80'
	  - '443:443'

