# リッチコマンド出力 実装プラン

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** ネットワーク機器の対話型コマンドフロー（プロンプト付き）と、パラメータ・背景説明を追加して、検索結果ページを参照せずにコマンドが使えるようにする。

**Architecture:** `CommandItem` 構造体に `Prompt`・`Purpose`・`Params` を追加してシステムプロンプトを更新。フロントエンドは `prompt` の有無で表示を切り替え、`purpose` と `params` を追加表示する。

**Tech Stack:** Go (Gin), Vanilla JS, Ollama (gpt-oss:20b)

**Design Doc:** `docs/plans/2026-03-04-rich-command-output-design.md`

---

### Task 1: CommandItem 構造体とシステムプロンプトの更新

**Files:**
- Modify: `search/summarize.go`

**Step 1: `CommandItem` に3フィールドを追加**

`search/summarize.go` の `CommandItem` を以下に変更:

```go
type CommandItem struct {
	Step        int      `json:"step"`
	Prompt      string   `json:"prompt"`
	Command     string   `json:"command"`
	Description string   `json:"description"`
	Purpose     string   `json:"purpose"`
	Params      []string `json:"params"`
	Options     []string `json:"options"`
}
```

**Step 2: システムプロンプトを更新**

`systemPrompt` 変数を以下に差し替える:

```go
	systemPrompt := `あなたはインフラエンジニア向けのアシスタントです。
与えられた検索結果から、実際にターミナルで使えるコマンドを抽出してください。
必ず以下のJSON形式のみを返してください。説明文や前置きは不要です。

フィールドの説明:
- step: 実行順序（1から連番）
- prompt: CLIプロンプト（Cisco等の対話型コマンドの場合のみ。例: "Router#", "Router(config-if)#"。Linuxコマンドは空文字列）
- command: 実行するコマンド
- description: このコマンドの目的を1行で
- purpose: なぜこのコマンドが必要か1〜2文で説明
- params: コマンド内の具体的なパラメータ値の意味（例: ["192.168.1.1: ルーターのIPアドレス"]）。パラメータがない場合は空配列
- options: 主要なオプションフラグの説明（例: ["-t: 構文チェックのみ実行"]）。ない場合は空配列

{"commands": [{"step": 1, "prompt": "", "command": "コマンド", "description": "目的を1行で", "purpose": "なぜ必要か1〜2文", "params": ["値: 説明"], "options": ["-x: 説明"]}]}`
```

**Step 3: ビルド確認**

```bash
cd C:/Users/kaito/project/configuan/infra-search
C:/Users/kaito/AppData/Local/Temp/goinstall/go/bin/go.exe build ./...
```

Expected: エラーなし

**Step 4: Commit**

```bash
git add search/summarize.go
git commit -m "feat: add prompt, purpose, params to CommandItem"
```

---

### Task 2: UI 表示の更新

**Files:**
- Modify: `templates/index.html`
- Modify: `static/style.css`

**Step 1: コマンド表示部分を更新**

`index.html` の `m.commands.map(cmd => ...)` 部分を以下に差し替える。
`prompt` の有無でプレフィックスを切り替え、`purpose` と `params` を追加表示する:

```javascript
${m.commands.map(cmd => {
  const prefix = cmd.prompt
    ? `<span class="command-prompt">${escapeHtml(cmd.prompt)}</span>`
    : `<span class="command-prefix">$</span>`;
  const purposeHtml = cmd.purpose
    ? `<div class="command-purpose">${escapeHtml(cmd.purpose)}</div>`
    : '';
  const paramsHtml = (cmd.params && cmd.params.length > 0)
    ? `<ul class="command-options">${cmd.params.map(p => `<li>${escapeHtml(p)}</li>`).join('')}</ul>`
    : '';
  const optionsHtml = (cmd.options && cmd.options.length > 0)
    ? `<ul class="command-options">${cmd.options.map(o => `<li>${escapeHtml(o)}</li>`).join('')}</ul>`
    : '';
  return `<div class="command-item">
    <div class="command-line">
      <span class="command-step">${cmd.step}.</span>
      ${prefix}
      <code>${escapeHtml(cmd.command)}</code>
      <button class="copy-btn" data-cmd="${escapeHtml(cmd.command)}">コピー</button>
    </div>
    <div class="command-desc">${escapeHtml(cmd.description)}</div>
    ${purposeHtml}
    ${paramsHtml}
    ${optionsHtml}
  </div>`;
}).join('')}
```

**Step 2: CSS を追加**

`static/style.css` の末尾に以下を追加:

```css
.command-prefix {
  font-family: monospace;
  color: var(--subtext0);
  margin-right: 4px;
}

.command-prompt {
  font-family: monospace;
  color: var(--green, #a6e3a1);
  margin-right: 6px;
  font-size: 0.9em;
}

.command-purpose {
  margin: 4px 0 2px 16px;
  font-size: 0.85em;
  color: var(--subtext0);
  font-style: italic;
}
```

**Step 3: 動作確認**

サーバーを起動してブラウザで `http://localhost:8080` を開く:

```bash
C:/Users/kaito/AppData/Local/Temp/goinstall/go/bin/go.exe run main.go
```

以下の2パターンで検索し表示を確認:
- 「nginx 設定確認」→ `prompt` が空、`$ nginx -t` 形式で表示
- 「Cisco BGP 設定」→ `prompt` が `Router(config)#` 等で表示、`purpose` と `params` が出る

**Step 4: Commit**

```bash
git add templates/index.html static/style.css
git commit -m "feat: display prompt, purpose, params in command output"
```
