# KurokoWeb

## 如何编译

```bash
git submodule update --init --recursive
cd ./web/app
yarn install && yarn build
```

托管 `./web/app/dist` 即可。

```bash
caddy file-server --listen :26743 --root ./web/app/dist
```
