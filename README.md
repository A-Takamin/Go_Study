go build -o chat
./chat

# ポートを指定したいなら。
./chat -addr=":3000"

# テスト
go test