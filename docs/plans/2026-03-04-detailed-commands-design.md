# 詳細コマンドサマリー 設計ドキュメント

> 作成日: 2026-03-04

---

## 概要

コマンドサマリーに実行順序（`step`）とオプション説明（`options`）を追加し、より実用的な出力にする。

---

## 変更ファイル

| ファイル | 変更内容 |
|---|---|
| `search/summarize.go` | `CommandItem` 構造体に `Step`, `Options` を追加、システムプロンプト更新 |
| `templates/index.html` | ステップ番号・オプション箇条書きの表示追加 |

---

## データ構造

### 変更前

```go
type CommandItem struct {
    Command     string `json:"command"`
    Description string `json:"description"`
}
```

### 変更後

```go
type CommandItem struct {
    Step        int      `json:"step"`
    Command     string   `json:"command"`
    Description string   `json:"description"`
    Options     []string `json:"options"`
}
```

---

## システムプロンプト

```
あなたはインフラエンジニア向けのアシスタントです。
与えられた検索結果から、実際にターミナルで使えるコマンドを抽出してください。
必ず以下のJSON形式のみを返してください。説明文や前置きは不要です。
step は実行順序（1から連番）、options は主要なオプションの説明を配列で記載してください。

{"commands": [{"step": 1, "command": "実際のコマンド", "description": "このコマンドの目的を1行で", "options": ["-x: オプションの説明"]}]}
```

---

## UI 表示イメージ

```
┌─────────────────────────────────┐
│ コマンド                         │
├─────────────────────────────────┤
│ 1. $ nginx -t          [コピー] │
│    設定ファイルの構文チェック      │
│    • -t: 構文チェックのみ実行     │
│    • -c <path>: ファイルを指定    │
│                                 │
│ 2. $ nginx -s reload   [コピー] │
│    設定を無停止で反映             │
│    • -s reload: ワーカー再起動なし│
└─────────────────────────────────┘
```

---

## 採用しなかった案

| 案 | 理由 |
|---|---|
| ステップをグループ化 (`steps[]`) | レスポンス型変更が大きく LLM も出力しにくい |
| オプションを1文字列 | 表示の柔軟性が低い |
