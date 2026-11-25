---
schema_version: 1
document_type: task
task_id: "{{task_id}}"
parent_task_id: "{{parent_task_id}}"
revision: 1
kind: execute
owner_member_id: "{{owner_member_uuid}}"
assign_reason: "{{assign_reason}}"
appointment_ref:
  document_id: "{{appointment_id}}"
  repo_id: xinghe-docs
  commit_id: "{{appointment_commit}}"
  path: "appointments/{{appointment_id}}.md"
file_scope:
  - "{{file_or_dir_pattern}}"
spec_revision: 1
business_state: published
inputs: []
acceptance_owner: admin
---

# {{title}}

## 目标与范围
{{goal_and_scope}}

## 输入与验收要求
- 固定输入：见上方 `inputs` 引用。
- 验收标准：{{acceptance_criteria}}

## 流程、约束与决策
- 严格限定修改在 `file_scope` 范围内；范围变更须由管理员发布新约定版本。

## 进度
- 待领取执行。

## 结果与验证
- 尚未交付。

## 产物
- 待提交交付清单后自动关联。

## 遗留与下一步
- —
