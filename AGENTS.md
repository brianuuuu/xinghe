# AGENTS.md — 星河（Xinghe）项目工作规范

本文件为所有参与星河项目的开发者与 AI Agent 提供仓库级工程与协作约定。

---

## 1. 本项目架构唯一事实真相文件

- **（暂空，尚未编写）**
- 正式架构唯一事实真相文档尚未定稿并建立；在此之前，请勿将其他临时说明、草稿或讨论记录自行假定为唯一架构真相。后续创建后将在此处补充精确路径。

---

## 2. Git 提交格式

采用精简 Conventional Commits 格式，使用简洁中文祈使句描述：

```text
<type>(<scope>): <中文简明描述>

[可选 body：说明变更原因、影响范围与验证方式]
```

### 2.1 示例

- `chore(repo): 删除 old 目录并清理相关配置与文档引用`
- `feat(star): 实现管理员中心化消息路由拦截`
- `fix(runtime): 修复 Windows 下 claude.exe 进程树清理遗漏`
- `docs(design): 补充任务交付清单校验说明`

### 2.2 `type` 与 `scope` 约定

- **常用 `type`**：`feat`（新功能）、`fix`（缺陷修复）、`docs`（文档/规范）、`refactor`（重构）、`test`（测试/校验）、`style`（样式/格式）、`perf`（性能优化）、`chore`（构建/目录/配置清理）。
- **常用 `scope`**（可选，对应模块或目录）：`repo`、`domain`、`star`、`assistant`、`member`、`docflow`、`store`、`runtime`、`server`、`web`、`templates`、`docs`、`prototype`、`scripts`。

### 2.3 提交纪律

- 标题使用简洁中文祈使句，单次提交聚焦单一意图，末尾不加句号。
- 提交前必须通过对应构建与检查；严禁将 `.xinghe/`、`*.sqlite*`、构建产物或密钥混入提交。
- **未经用户明确指令不执行 `git push`**，严禁 `git push --force` 或改写已公开分支历史。

---

## 3. Web 界面规范

- **强制兼容浅色（Light）与深色（Dark）双模式**：所有背景、面板、文字、边框、状态徽标、代码/Diff 预览与连接线必须同时适配浅色与深色主题，保证两种模式下的对比度与可读性。
- **语义化主题 Token**：严禁在组件或样式中硬编码仅适用单一主题的颜色值；统一通过语义化 CSS 变量定义浅/深两套主题变量，支持跟随系统 `prefers-color-scheme` 及手动切换主题。
- **零啰嗦提示文案（自解释交互）**：界面上严禁堆砌额外的提示性、说明性文字（如大段用法引导、功能解释条、防呆说明句）；如果某个功能需要靠啰嗦的提示文字才能让用户看懂，说明界面信息架构或交互设计本身存在问题，必须优先重构界面与交互而非打文字补丁。
- **保持轻量与布局基准**：前端（`web/`）采用 React + TypeScript（开启 `strict: true`）与原生 CSS 变量/模块化样式，不引入 Ant Design、MUI 等重型 UI 组件库；桌面三栏布局参考 `prototype/workbench.html`。

---

## 4. 数据库操作审批红线

- **任何数据库删除操作必须事先审批**：包括但不限于 `DELETE`、`TRUNCATE`、`DROP TABLE`、`DROP INDEX`、破坏性表结构迁移，以及删除、覆盖或清空任何 `.sqlite` / `.db` 文件（含开发运行库与可重建投影表）。
- 执行任何上述操作前，必须先向用户明确说明目标对象、操作原因与影响范围，获得明确批准后方可执行；严禁在脚本、迁移或业务代码中暗含未获批的自动删库、删表或清空数据逻辑。

---

## 5. Changelog 位置

- **固定路径**：仓库根目录 `CHANGELOG.md`（当前尚未创建，首次发布版本或重要里程碑汇总时于根目录创建）。
- 维护方式：按版本/日期倒序记录重要变更，建议分类为 `新增`、`变更`、`修复`、`移除`。

---

## 6. 一级目录介绍

```text
xinghe/
├── cmd/            # Go 主程序入口（cmd/xinghe）
├── internal/       # Go 核心业务模块（当前 go 1.20：domain / star / assistant / member / docflow / store / runtime / server）
├── web/            # React + TypeScript 桌面 Web 工作台前端（Vite）
├── templates/      # 角色、任命、任务与评审默认 Markdown 模板
├── scripts/        # 辅助脚本（如环境检测脚本 check-env.ps1）
├── docs/           # 需求文档、概要设计等项目文档
├── prototype/      # 交互式 HTML 工作台原型
├── .xinghe/        # [gitignored] 本地运行期数据（文档 Git 库、runtime.sqlite、草稿、产物与日志），严禁提交
├── AGENTS.md       # 仓库协作与工程规范（本文件）
├── CHANGELOG.md    # [待建] 变更日志
├── README.md       # 项目概览与快速说明
└── go.mod          # Go 模块定义
```

---

## 7. 启停脚本

- **目前尚未创建启停脚本**。
- 后续统一放置于 `scripts/` 目录下，补齐后须同步更新本节说明；在此之前，请勿在文档、注释或回复中虚构不存在的启停脚本路径或命令。
