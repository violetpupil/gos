# [rancher](https://www.rancher.com/)

[文档](https://ranchermanager.docs.rancher.com/)

```batch
docker run -d --restart=unless-stopped ^
  -p 80:80 -p 443:443 ^
  --privileged ^
  --name rancher ^
  rancher/rancher:latest
```
