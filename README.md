# Leet Book

数据结构与算法（Data Structures & Algorithms）— 个人学习笔记，包含 Hugo 驱动的在线文档和 Go 实现的数据结构。

在线地址：**[www.subond.com/leetcode](https://www.subond.com/leetcode)**

## 项目结构

```
book/          # Markdown 书稿（Hugo content 源文件）
  docs/        #   章节：part1（基础）、part2（进阶）、part3（高级）
  menu/        #   侧边栏菜单定义
  assets/      #   图片资源（PNG/SVG）
website/       # Hugo 站点 — 构建并本地预览
  make.go      #   脚本：将 book/ 文档复制到 content/，修正资源路径
  Makefile     #   make all 运行 make.go；make s 启动 Hugo 开发服务器
  hugo.toml    #   Hugo 配置（主题: hugo-book, baseURL: subond.com/leetcode）
  content/     #   由 make.go 生成 — 不提交到 git
  static/      #   由 make.go 生成 — 不提交到 git
template/      # Go 语言实现的数据结构（package template）
  *.go         #   linkedlist, stack, lru/lfu cache, sort, trie, segment_tree 等
  *_test.go    #   对应测试文件
```

## 本地构建

### 环境要求

- **Go**（运行 make.go 和 template 测试）
- **Hugo extended**（>= v0.158.0，主题需要 SCSS 支持）

```bash
# 安装 Hugo extended
brew install hugo
hugo version  # 应显示 +extended
```

### 克隆仓库

```bash
git clone --recurse-submodules git@github.com:yusubond/leetcode.git
cd leetcode
```

### 启动网站

```bash
cd website
make all    # 将 book/ 的 Markdown 和图片复制到 Hugo content/ 和 static/
make s      # 启动 Hugo 开发服务器 http://localhost:1313/leetcode/
```

`make.go` 的核心工作：
1. 遍历 `../book/`，将 `.md` 文件复制到 `content/`（保持目录结构），并修正 `./assets` 图片路径
2. 将 `.png`/`.svg` 图片复制到 `static/img/`

### 运行 Go 测试

```bash
cd template
go test ./...
```

## 提交与发布

站点通过 GitHub Actions 自动构建并发布到 GitHub Pages（`gh-pages` 分支）。

### 提交流程

直接提交到 `master` 分支：

```bash
git add .
git commit -m "描述你的改动"
git push
```

### 自动发布

推送后 GitHub Actions 自动执行（`.github/workflows/deploy.yml`）：

1. 检出仓库及子模块
2. 运行 `make all` 将 `book/` 转换为 Hugo content
3. 运行 `hugo --minify` 构建站点到 `public/`
4. 将 `public/` 推送到 `gh-pages` 分支

GitHub Pages 从 `gh-pages` 分支提供站点，几分钟内更新。

## 写作约定

- LeetCode 题目文件命名：`No.XXX_题目名.md`
- 评分系统：⭐ 1–5 表示题目设计质量
- 章节图片放在各自 `assets/` 子目录下，make.go 会将其展平到 `static/img/`
