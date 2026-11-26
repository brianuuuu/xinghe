# 星河（Xinghe）· 多 Agent 协作工作台

**跨 Harness 协同，以文档组织分工、交接与验收。**

面向需要同时组织多个自主 Agent（如 Codex CLI、Claude Code）的个人开发者与创作者，围绕**能力分工**、**共享额度池**、**星型沟通（禁止成员互 `@`）**与**有版本的正式文件流转**构建可控协作秩序。

## 星河为什么存在

不同 Agent 擅长的事情不同，使用成本也不同。有的适合研究和规划，有的擅长实现和验证；对于同一项工作，也可能有多个能胜任、但成本不同的选择。

一个复杂任务往往包含多个环节，每个环节需要的能力和投入都不一样。因此，需要根据任务要求、Agent 的能力和成本来分工，让合适的 Agent 负责合适的部分，再把各自的成果衔接起来。

**星河就是为这种分工协作而存在的：把不同 Agent 的长处用起来，以合理的成本完成复杂任务。** 星河通过任命书、任务文件和报告，帮助用户明确分工、交接成果和验收结果，让各个 Agent 能够独立工作，也能共同完成任务。

## 一级目录结构

```text
xinghe/
├── cmd/          # Go 主程序入口（xinghe 服务，首版内嵌本地代理）
├── internal/     # Go 核心业务模块（domain / star / assistant / member / docflow / store / runtime / server）
├── web/          # React + TypeScript 桌面 Web 工作台前端
├── templates/    # 内置角色、任命、任务与评审默认 Markdown 模板
├── scripts/      # 本地开发与 P0 双 CLI 环境验证脚本
├── docs/         # 需求文档（v0.5）与系统概要设计（v0.4）
├── prototype/    # 交互式 HTML 工作台原型
├── AGENTS.md     # 开发者与多 Agent 协作、代码、Web UI 与提交规范
└── .xinghe/      # [gitignored] 本地运行期数据（Git 正式文档库、SQLite、草稿、产物与日志）
```

## 快速开始

```powershell
# 1. 检查本机 Go / Node / Git / Codex CLI / Claude Code
.\scripts\check-env.ps1

# 2. 启动 Go 服务骨架（默认 127.0.0.1:8787）
go run ./cmd/xinghe

# 3. 安装并启动前端工作台骨架（可选）
cd web && npm install && npm run dev
```

## 设计文档与原型

- 需求文档：[docs/星河-初版需求文档-v0.1.md](docs/星河-初版需求文档-v0.1.md)（正文 v0.5）
- 概要设计：[docs/星河-系统概要设计-v0.1.md](docs/星河-系统概要设计-v0.1.md)（正文 v0.4）
- 工作台原型：[prototype/workbench.html](prototype/workbench.html)
- 协作与开发规范：[AGENTS.md](AGENTS.md)
