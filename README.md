<p align="center">
  <img src="web/public/logo.svg" alt="星河 Logo" width="72" height="72" />
</p>

<h1 align="center">星河 Xinghe</h1>

<p align="center"><strong>多 Agent 协作工作台</strong></p>

<p align="center">
  跨 Harness 协同，以文档组织分工、交接与验收。
</p>

<p align="center">
  <img src="https://img.shields.io/badge/stage-prototype-4d78cc?style=flat" alt="阶段：原型开发" />
  <img src="https://img.shields.io/badge/Go-00ADD8?style=flat&amp;logo=go&amp;logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/React%20%2B%20TypeScript-3178C6?style=flat&amp;logo=react&amp;logoColor=white" alt="React + TypeScript" />
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-GPL--2.0-64748B?style=flat" alt="License: GPL-2.0" /></a>
</p>

<p align="center">
  <a href="#产品简介">产品简介</a> ·
  <a href="#为什么需要星河">为什么需要星河</a> ·
  <a href="#快速开始">快速开始</a> ·
  <a href="#文档与原型">文档与原型</a>
</p>

<p align="center">
  <sub>当前为原型与工程骨架阶段 · <a href="prototype/workbench.html">查看工作台原型（HTML）</a></sub>
</p>

---

## 产品简介

星河面向需要同时组织 Codex CLI、Claude Code 等自主 Agent 的个人开发者与创作者，围绕以下方向设计：

<table>
  <tr>
    <td width="50%" valign="top">
      <h3>能力分工</h3>
      <p>根据任务要求、Agent 的能力和成本安排分工，让各个 Agent 独立完成自己负责的部分。</p>
    </td>
    <td width="50%" valign="top">
      <h3>共享额度池</h3>
      <p>统一管理多个 Agent 共用的账号或订阅额度，让任务安排兼顾能力与可用额度。</p>
    </td>
  </tr>
  <tr>
    <td width="50%" valign="top">
      <h3>集中协调</h3>
      <p>由管理员分配任务、处理分歧和验收成果，成员之间不直接互 @。</p>
    </td>
    <td width="50%" valign="top">
      <h3>文档交接</h3>
      <p>以有版本的任命书、任务文件和报告承接工作，明确输入、交付与验收依据。</p>
    </td>
  </tr>
</table>

## 为什么需要星河

不同 Agent 擅长的事情不同，使用成本也不同。有的适合研究和规划，有的擅长实现和验证；对于同一项工作，也可能有多个能胜任、但成本不同的选择。

一个复杂任务往往包含多个环节，每个环节需要的能力和投入都不一样。因此，需要根据任务要求、Agent 的能力和成本来分工，让合适的 Agent 负责合适的部分，再把各自的成果衔接起来。

**星河就是为这种分工协作而存在的：把不同 Agent 的长处用起来，以合理的成本完成复杂任务。** 星河通过任命书、任务文件和报告，帮助用户明确分工、交接成果和验收结果，让各个 Agent 能够独立工作，也能共同完成任务。

## 快速开始

```powershell
# 1. 检查本机 Go / Node / Git / Codex CLI / Claude Code
.\scripts\check-env.ps1

# 2. 启动 Go 服务骨架（默认 127.0.0.1:8787）
go run ./cmd/xinghe

# 3. 安装并启动前端工作台骨架（可选）
cd web && npm install && npm run dev
```

## 目录结构

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

## 文档与原型

| 入口 | 内容 |
| --- | --- |
| [需求文档](docs/星河-初版需求文档-v0.1.md) | 产品定位、业务概念与功能范围 · 正文 v0.5 |
| [概要设计](docs/星河-系统概要设计-v0.1.md) | 模块划分、数据存储与运行机制 · 正文 v0.4 |
| [工作台原型](prototype/workbench.html) | 下载仓库后用浏览器打开，体验交互与主题切换 |
| [协作与开发规范](AGENTS.md) | 工程约定、界面规范与提交规则 |

<p align="center"><sub>开源许可证：<a href="LICENSE">GPL-2.0</a></sub></p>
