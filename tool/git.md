# git

```bash
# 配置用户名
git config --global user.name "instafever"
# 配置邮箱
git config --global user.email "495140158@qq.com"
# 删除远程分支
git push origin --delete <branch>
# 取消远程分支绑定
git branch --unset-upstream
# push提交时，如果有未拉取的提交，自动merge远程分支
git config --global pull.rebase false
```
