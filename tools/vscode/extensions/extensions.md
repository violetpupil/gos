# Extensions

## [ESLint](https://marketplace.visualstudio.com/items?itemName=dbaeumer.vscode-eslint)

javascript代码检查

## [Prettier](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)

代码格式化，可设置为所有或特定语言的格式器

```json
{
  // 默认使用prettier插件格式化
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "[javascript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[typescript]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  // tsx
  "[typescriptreact]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  "[json]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  },
  // 可以有注释的json
  "[jsonc]": {
    "editor.defaultFormatter": "esbenp.prettier-vscode"
  }
}
```

## [Black Formatter](https://marketplace.visualstudio.com/items?itemName=ms-python.black-formatter)

python格式化

```json
{
  "[python]": {
    "editor.defaultFormatter": "ms-python.black-formatter"
  }
}
```
