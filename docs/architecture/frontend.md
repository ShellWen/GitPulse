# GitPulse 前端架构

GitPulse 的前端部分采用 Vite + React.js 构建，使用 TypeScript 编写。

## 技术栈

- 构建工具：Vite \
  Vite 是久经考验的高性能打包器。Vite 通过利用现代浏览器的原生 ES 模块导入功能，实现了快速的冷启动和热模块重载。Vite 的快速构建速度和热更新功能，使得我们可以更快地开发和调试前端代码。\
  值得注意的是，我们没有使用 SSR，而是直接使用了 Vite 打包 SPA 应用。这有利于减少复杂度，并且减少了服务端渲染带来的复杂性和维护成本。

- 开发语言：TypeScript \
  类型安全是大型项目稳定性的基石，在 GitPulse 的早期开发阶段，我们就决定使用 TypeScript 作为主要开发语言。TypeScript 为我们提供了强大的类型检查功能，使得我们可以在编译阶段就发现潜在的问题，提高了代码的可维护性和可读性。

- 路由库：TanStack Router \
  我们使用 TanStack Router 作为路由库的原因无它。市面上有很多基于 React 的路由库，对 SPA 而言，大部分的用户会选择使用 React Router。诚然，React Router 是一个十分优秀的路由库，开发迭代速度快，社区活跃，文档齐全。但 TanStack Router 的路由绑定使用 TypeScript 确保类型安全，并且提供了编译时代码生成器以在编译时从文件生成路由配置，无需手写路由，这节约了我们的时间，并提供了更高的可维护性。

- 网络请求：React Query \
  React Query 是一个基于 React Hooks 的数据获取库，它提供了一种简单的方式来管理数据获取和缓存。React Query 通过提供了一种简单的方式来管理数据获取和缓存，使得我们可以更加方便地处理数据请求和缓存。React Query 通过提供了一种简单的方式来管理数据获取和缓存，使得我们可以更加方便地处理数据请求和缓存。

- 样式：Tailwind CSS + DaisyUI \
  Tailwind CSS 是一个高度可定制的 CSS 框架，它提供了一系列的 utility classes，使得我们可以更加方便地编写样式。DaisyUI 是一个基于 Tailwind CSS 的组件库，它提供了一系列的组件，使得我们可以更加方便地构建页面。我们所有的样式都是基于 Tailwind CSS 和 DaisyUI 来编写的。

- 开发中 Mock 数据：MSW \
  MSW 是一个用于模拟网络请求的库，它可以拦截浏览器发出的请求，并返回我们预先定义好的数据。在开发阶段，我们使用 MSW 来模拟后端接口，以便在没有后端接口的情况下进行开发。

- 其它：ESLint、Prettier 等 \
  我们使用 ESLint 和 Prettier 来保证基本的代码质量。ESLint 用于检查代码中的潜在问题，并提供来帮助我们编写更加规范的代码的一些规则。Prettier 则用于格式化代码，使得代码风格保持一致。

## 项目结构

在此我们仅罗列出项目的主要目录结构，具体的目录结构和文件结构请参考项目源码。

- `frontend/`: 前端源码目录
  - `src/`:
    - `components/`: 存放 React 组件，这里的组件都是无状态组件
    - `lib/`: 存放工具库和自定义 Hook
      - `api/`: 存放网络请求相关的代码，这些代码是框架无关的，便于后续扩展维护
      - `query/`: 存放 React Query 相关的代码，使用上面的 `api/` 目录定义的请求函数与后端通信
    - `mocks/`: 存放 Mock 数据和对应的 handlers
    - `routes/`: 存放页面文件，这些文件会被 TanStack Router 的编译时代码生成器生成路由配置
    - `index.css`: 全局样式文件
    - `main.tsx`: 应用入口文件

## 主要功能

项目包含以下路由：

- `/`: 首页
- `/rank`: 用户排行榜，在这里可以查看用户的贡献排名，还可通过使用的编程语言于所处的地区进行筛选
- `/u/:username`: 用户详情页，展示用户的详细信息，包括用户主要使用的编程语言、PulsePoint 与 猜测的所处地区的信息 等

## 开发指南

### 安装依赖

```bash
corepnpm prepare pnpm
pnpm install
```

### 修改环境变量

在项目根目录下创建 `.env` 文件，参考 `.env.example` 文件填写相应的环境变量。

### 启动开发服务器

```bash
pnpm run dev
```
