# [Xray](https://xtls.github.io/)

[文档](https://xtls.github.io/document/)

```bash
# 安装
bash -c "$(curl -L https://github.com/XTLS/Xray-install/raw/main/install-release.sh)" @ install
# 重启服务
systemctl restart xray
# 查看服务状态
systemctl status xray
# 验证
curl -x socks5://user:pass@localhost:2023 http://ipecho.net/plain
```
