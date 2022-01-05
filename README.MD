# 功能实现进度

```text
ok 管理员登录 
ok 展示主页 
ok 修改密码	 
ok 管理员退出 
ok 菜单
ok api
ok 管理员  Tue 04 Jan 2022	 03:48:23 PM	
ok 角色 Tue 04 Jan 2022	 08:59:04 PM	
ok 字典 Tue 04 Jan 2022	 08:59:04 PM	
ok 文件  Tue 04 Jan 2022	 08:59:04 PM	

ok 文件生成 Wed 05 Jan 2022	09:02:32 PM	
    - service : 文件名，实体名称
    - api  : 文件名 
    - handler  文件名
    - cmd 文件名
    - data 文件名
    - index 文件名 
```

## 自动代码生成步骤

1. 创建数据表
2. 使用 gf gen dao 生成实体文件
3. 转到文件生成方法 ciel-admin/interface/utility/utils/xfile/xfile_test.go
4. 配置 `TemplateConfig`  具体说明见 TemplateConfig   `/interface/utility/utils/xfile/xfile.go`
5. 运行 Gen方法生成文件
6. 到菜单页面创建对应菜单
7. 到角色页面分配角色菜单权限
8. 退出重新登录