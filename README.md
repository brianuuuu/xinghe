# 星河（Xinghe）· 多 Agent 协作工作台

**跨 Harness 协同，以文档组织分工、交接与验收。**

面向需要同时组织多个自主 Agent（如 Codex CLI、Claude Code）的个人开发者与创作者，围绕**能力分工**、**共享额度池**、**星型沟通（禁止成员互 `@`）**与**有版本的正式文件流转**构建可控协作秩序。

## 星河为什么存在

**星河存在的前提，是 Agent 之间有值得协作的差异。** 这种差异可以来自模型、Harness、工具、领域知识，以及角色的方法论。同一 Harness 下的不同角色也可以形成互补。星河面向的是需要保留这些差异，让多个自主 Agent 共同完成一项工作的场景。

当用户在多个 Agent 之间切换时，协调成本也随之出现：重复交代背景、搬运阶段成果、核对当前进度、处理意见分歧。仅靠分散的会话记录，很难持续说清楚谁负责什么、以哪个版本为依据、交付是否完成、结果是否通过验收。

星河以正式文档建立协作边界：任命书明确职责，任务文件约定目标、输入与验收标准，报告和版本化产物承接交付。成员在各自的 Harness 中自主执行，通过明确版本的文件共享必要上下文；任务分配、交接与验收由管理员统一协调。星河关注职责、进度和正式交付，成员保留自己的内部执行过程与会话。

**自主执行、职责清晰、文档交接、任务流转可追溯，是星河的核心理念。** 如果单个 Agent 已能可靠完成任务，就没有必要额外组织团队。星河的必要性，应由能力互补、接力交付和独立评审带来的实际收益来验证。

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
