---
schema_version: 1
document_type: task
task_id: "{{task_id}}"
parent_task_id: "{{parent_task_id}}"
revision: 1
kind: review
owner_member_id: "{{reviewer_member_uuid}}"
assign_reason: "与实现方异构，独立交叉评审"
appointment_ref:
  document_id: "{{appointment_id}}"
  repo_id: xinghe-docs
  commit_id: "{{appointment_commit}}"
  path: "appointments/{{appointment_id}}.md"
file_scope:
  - "仅写评审报告，不修改被审代码或原任务文件"
spec_revision: 1
business_state: published
target_submission_id: "{{target_submission_id}}"
inputs: []
acceptance_owner: admin
---

# 评审：{{target_task_title}}

## 评审对象与标准
- 目标交付：`{{target_submission_id}}`（固定版本）
- 评审标准：功能正确性、边界与异常、是否遵守 `file_scope`、验证证据充分性。

## 必须修改项
- 待填写。

## 可选建议
- 待填写。

## 总结论与建议动作
- 向管理员交付结论（建议验收 / 建议退回返工）。
