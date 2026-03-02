# infra-search

インフラエンジニア向けのコマンド検索ツール。自然言語で質問すると、DuckDuckGo で検索した結果をローカル LLM (Ollama) が解析し、実際に使えるコマンドを抽出して提示します。

---

## 特徴

- **自然言語検索** — 「Cisco の OSPF 設定確認コマンド」のように日本語で質問できる
- **ベンダー認識** — Cisco / Juniper / Arista / AWS / Azure などを自動検出し、公式ドキュメントを優先検索
- **コマンド抽出** — Ollama (ローカル LLM) が検索結果スニペットからターミナルで使えるコマンドを整形して返す
- **コピーボタン付き UI** — Catppuccin テーマのチャット UI でコマンドをワンクリックでコピー

---

## プロジェクト構成

```
infra-search/
├── main.go                   # Gin サーバーエントリポイント
├── handlers/
│   └── search.go             # POST /api/search ハンドラー
├── search/
│   ├── query.go              # キーワードベースのクエリ生成
│   ├── fetch.go              # DuckDuckGo HTML スクレイピング
│   └── summarize.go          # Ollama 連携・コマンド抽出
├── templates/
│   └── index.html            # チャット UI
└── static/
    └── style.css             # スタイル（Catppuccin カラーテーマ）
```

---

## セットアップ

### 前提条件

- Go 1.21 以上
- [Ollama](https://ollama.com/) がローカルで起動していること

### インストール

```bash
git clone https://github.com/your-org/infra-search.git
cd infra-search
go mod download
```

### Ollama モデルの準備

```bash
ollama pull gpt-oss:20b
```

### サーバー起動

```bash
go run main.go
```

ブラウザで `http://localhost:8080` にアクセスしてください。

---

## 環境変数

| 変数 | 必須 | デフォルト | 説明 |
|---|---|---|---|
| `OLLAMA_URL` | - | `http://localhost:11434` | Ollama エンドポイント |
| `OLLAMA_MODEL` | - | `gpt-oss:20b` | 使用モデル名 |

---

## 使い方

1. テキストボックスに質問を入力する（例: `Cisco のルーティングテーブル確認コマンド`）
2. 送信すると DuckDuckGo での検索結果と、LLM が抽出したコマンド一覧が表示される
3. コマンドの横のコピーボタンでクリップボードにコピーできる

---

## 開発予定

- [ ] GitHub OAuth 認証・レート制限・セキュリティヘッダー
- [ ] Ollama レスポンスのストリーミング対応
- [ ] 会話履歴の保持、検索履歴、お気に入り保存
- [ ] テストの追加
