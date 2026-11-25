---
schema_version: 1
document_type: appointment
appointment_id: "{{appointment_id}}"
revision: 1
member_id: "{{member_uuid}}"
role: "{{role_id}}"
issued_by: admin
scope: assigned_tasks
---

# 任命书：{{role_name}}

## 职责
{{responsibilities}}

## 方法与原则
{{principles}}

## 权限边界
- 实际读写与操作范围由星河服务端授权决定；本任命正文声明不可扩权。
- 遵循星型沟通规则：仅向管理员发起汇报、求助与正式交付，禁止直接 `@` 其他成员交接。

## 交付要求
- 交付时须提供明确结果、验证证据、固定产物引用与遗留问题。

## 停止与升级
- 遇到输入缺失、范围变更、超出 `file_scope` 或环境受阻时立即停止并向管理员求助。
