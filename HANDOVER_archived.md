# HANDOVER アーカイブ

---

## 2026-03-01 セッション① — Ollama コマンド整形機能の実装

### セッション概要

`infra-search` アプリに Ollama 連携によるコマンド整形機能を追加した。
DuckDuckGo の検索結果スニペットをローカルの Ollama モデル（`gpt-oss:20b`）に渡し、
ターミナルで使えるコマンド一覧を生成してフロントのコマンドセクションに表示する機能を実装・動作確認まで完了。

### 試行錯誤したポイント

#### モデル名のミス
- 最初 `oss-gpt-20b` → 次に `gpt-oss-20b` と指定したが、実際は `gpt-oss:20b`（コロン区切り）
- `ollama list` で正確な名前を確認して解決

```
NAME           ID              SIZE     MODIFIED
gpt-oss:20b    17052f91a42e    13 GB    3 months ago
```

#### ポート 8080 の競合
- 前回セッションの `main.exe` が残っていて起動できなかった
- エラー: `[GIN-debug] [ERROR] listen tcp :8080: bind: Only one usage of each socket address`
- `taskkill /IM main.exe /F` で解決

#### Go バイナリパスの問題
- プランファイルに `/tmp/goinstall/go/bin/go` と書かれていたが Windows 環境では動作しない
- 実パス: `C:/Users/kaito/AppData/Local/Temp/goinstall/go/bin/go.exe`
- プランファイルをすべて書き換えて対応

#### Ollama `/api/chat` の 404
- モデル名が間違っていたため 404 が返り続けた
- `curl` で直接 Ollama に叩いて応答を確認し、モデル名ミスと特定

### 学んだ教訓

- **Ollama のモデル名は `ollama list` で必ず確認する**（`name:tag` 形式）
- **Windows 環境では `/tmp/` パスは Git Bash 上でしか使えない**。プランに Windows パスを明記する
- **サーバー再起動忘れに注意**。コード変更後は `taskkill /IM main.exe /F` してから `go run main.go`
- `go run` はビルド済みバイナリを `main.exe` として起動するため、プロセス名は常に `main.exe`

---
