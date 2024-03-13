// vite.config.ts
import { defineConfig } from "file:///C:/Users/123/Desktop/code/lucy-demo/%E5%9C%9F%E6%8B%A8%E9%BC%A0%E5%85%85%E7%94%B5%E7%B3%BB%E7%BB%9F/groundhog-charging-system/front-manager/node_modules/.pnpm/registry.npmmirror.com+vite@4.3.9_@types+node@20.4.2_sass@1.63.6/node_modules/vite/dist/node/index.js";
import { createSvgIconsPlugin } from "file:///C:/Users/123/Desktop/code/lucy-demo/%E5%9C%9F%E6%8B%A8%E9%BC%A0%E5%85%85%E7%94%B5%E7%B3%BB%E7%BB%9F/groundhog-charging-system/front-manager/node_modules/.pnpm/registry.npmmirror.com+vite-plugin-svg-icons@2.0.1_vite@4.3.9/node_modules/vite-plugin-svg-icons/dist/index.mjs";
import { viteMockServe } from "file:///C:/Users/123/Desktop/code/lucy-demo/%E5%9C%9F%E6%8B%A8%E9%BC%A0%E5%85%85%E7%94%B5%E7%B3%BB%E7%BB%9F/groundhog-charging-system/front-manager/node_modules/.pnpm/registry.npmmirror.com+vite-plugin-mock@2.9.6_mockjs@1.1.0_rollup@2.79.1_vite@4.3.9/node_modules/vite-plugin-mock/dist/index.js";
import vue from "file:///C:/Users/123/Desktop/code/lucy-demo/%E5%9C%9F%E6%8B%A8%E9%BC%A0%E5%85%85%E7%94%B5%E7%B3%BB%E7%BB%9F/groundhog-charging-system/front-manager/node_modules/.pnpm/registry.npmmirror.com+@vitejs+plugin-vue@4.1.0_vite@4.3.9_vue@3.2.47/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import path from "path";
var vite_config_default = defineConfig(({ command }) => {
  const isProduction = command === "build";
  return {
    base: isProduction ? "https://static.tuboshulucy.top/" : "/",
    plugins: [
      vue(),
      createSvgIconsPlugin({
        iconDirs: [path.resolve(process.cwd(), "src/assets/icons")],
        symbolId: "icon-[dir]-[name]"
      }),
      viteMockServe({
        localEnabled: command === "serve"
      })
    ],
    resolve: {
      alias: {
        "@": path.resolve("./src")
        // 相对路径别名配置，使用 @ 代替 src
      }
    },
    css: {
      preprocessorOptions: {
        scss: {
          javascriptEnabled: true,
          additionalData: '@import "./src/styles/variable.scss";'
        }
      }
    },
    //代理跨域
    server: {
      proxy: {
        "/saas-api": {
          target: "http://127.0.0.1:8000/",
          changeOrigin: true
        }
      }
    }
  };
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJDOlxcXFxVc2Vyc1xcXFwxMjNcXFxcRGVza3RvcFxcXFxjb2RlXFxcXGx1Y3ktZGVtb1xcXFxcdTU3MUZcdTYyRThcdTlGMjBcdTUxNDVcdTc1MzVcdTdDRkJcdTdFREZcXFxcZ3JvdW5kaG9nLWNoYXJnaW5nLXN5c3RlbVxcXFxmcm9udC1tYW5hZ2VyXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ZpbGVuYW1lID0gXCJDOlxcXFxVc2Vyc1xcXFwxMjNcXFxcRGVza3RvcFxcXFxjb2RlXFxcXGx1Y3ktZGVtb1xcXFxcdTU3MUZcdTYyRThcdTlGMjBcdTUxNDVcdTc1MzVcdTdDRkJcdTdFREZcXFxcZ3JvdW5kaG9nLWNoYXJnaW5nLXN5c3RlbVxcXFxmcm9udC1tYW5hZ2VyXFxcXHZpdGUuY29uZmlnLnRzXCI7Y29uc3QgX192aXRlX2luamVjdGVkX29yaWdpbmFsX2ltcG9ydF9tZXRhX3VybCA9IFwiZmlsZTovLy9DOi9Vc2Vycy8xMjMvRGVza3RvcC9jb2RlL2x1Y3ktZGVtby8lRTUlOUMlOUYlRTYlOEIlQTglRTklQkMlQTAlRTUlODUlODUlRTclOTQlQjUlRTclQjMlQkIlRTclQkIlOUYvZ3JvdW5kaG9nLWNoYXJnaW5nLXN5c3RlbS9mcm9udC1tYW5hZ2VyL3ZpdGUuY29uZmlnLnRzXCI7aW1wb3J0IHsgZGVmaW5lQ29uZmlnIH0gZnJvbSAndml0ZSdcclxuaW1wb3J0IHsgY3JlYXRlU3ZnSWNvbnNQbHVnaW4gfSBmcm9tICd2aXRlLXBsdWdpbi1zdmctaWNvbnMnXHJcbmltcG9ydCB7IHZpdGVNb2NrU2VydmUgfSBmcm9tICd2aXRlLXBsdWdpbi1tb2NrJ1xyXG5pbXBvcnQgdnVlIGZyb20gJ0B2aXRlanMvcGx1Z2luLXZ1ZSdcclxuaW1wb3J0IHBhdGggZnJvbSAncGF0aCdcclxuXHJcbi8vIGh0dHBzOi8vdml0ZWpzLmRldi9jb25maWcvXHJcbmV4cG9ydCBkZWZhdWx0IGRlZmluZUNvbmZpZygoe2NvbW1hbmR9KT0+e1xyXG4gIGNvbnN0IGlzUHJvZHVjdGlvbiA9IGNvbW1hbmQgPT09ICdidWlsZCc7XHJcbiAgcmV0dXJuIHtcclxuICAgIGJhc2U6IGlzUHJvZHVjdGlvbiA/ICdodHRwczovL3N0YXRpYy50dWJvc2h1bHVjeS50b3AvJyA6ICcvJyxcclxuICAgIHBsdWdpbnM6IFt2dWUoKSxcclxuICAgICAgY3JlYXRlU3ZnSWNvbnNQbHVnaW4oe1xyXG4gICAgICAgIGljb25EaXJzOiBbcGF0aC5yZXNvbHZlKHByb2Nlc3MuY3dkKCksICdzcmMvYXNzZXRzL2ljb25zJyldLFxyXG4gICAgICAgIHN5bWJvbElkOiAnaWNvbi1bZGlyXS1bbmFtZV0nXHJcbiAgICAgIH0pLFxyXG4gICAgICB2aXRlTW9ja1NlcnZlKHtcclxuICAgICAgICBsb2NhbEVuYWJsZWQ6IGNvbW1hbmQgPT09ICdzZXJ2ZSdcclxuICAgICAgfSlcclxuICAgIF0sXHJcbiAgICByZXNvbHZlOiB7XHJcbiAgICAgIGFsaWFzOiB7XHJcbiAgICAgICAgJ0AnOiBwYXRoLnJlc29sdmUoJy4vc3JjJykgLC8vIFx1NzZGOFx1NUJGOVx1OERFRlx1NUY4NFx1NTIyQlx1NTQwRFx1OTE0RFx1N0Y2RVx1RkYwQ1x1NEY3Rlx1NzUyOCBAIFx1NEVFM1x1NjZGRiBzcmNcclxuICAgICAgfVxyXG4gICAgfSxcclxuICAgIGNzczoge1xyXG4gICAgICBwcmVwcm9jZXNzb3JPcHRpb25zOiB7XHJcbiAgICAgICAgc2Nzczoge1xyXG4gICAgICAgICAgamF2YXNjcmlwdEVuYWJsZWQ6IHRydWUsXHJcbiAgICAgICAgICBhZGRpdGlvbmFsRGF0YTogJ0BpbXBvcnQgXCIuL3NyYy9zdHlsZXMvdmFyaWFibGUuc2Nzc1wiOycsXHJcbiAgICAgICAgfSxcclxuICAgICAgfSxcclxuICAgIH0sXHJcbiAgICAvL1x1NEVFM1x1NzQwNlx1OERFOFx1NTdERlxyXG4gICAgc2VydmVyOntcclxuICAgICAgcHJveHk6e1xyXG4gICAgICAgICcvc2Fhcy1hcGknOiB7XHJcbiAgICAgICAgICB0YXJnZXQ6ICdodHRwOi8vMTI3LjAuMC4xOjgwMDAvJyxcclxuICAgICAgICAgIGNoYW5nZU9yaWdpbjogdHJ1ZVxyXG4gICAgICAgIH0sXHJcbiAgICAgIH1cclxuICAgIH1cclxuICB9XHJcbn0pXHJcbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBcWYsU0FBUyxvQkFBb0I7QUFDbGhCLFNBQVMsNEJBQTRCO0FBQ3JDLFNBQVMscUJBQXFCO0FBQzlCLE9BQU8sU0FBUztBQUNoQixPQUFPLFVBQVU7QUFHakIsSUFBTyxzQkFBUSxhQUFhLENBQUMsRUFBQyxRQUFPLE1BQUk7QUFDdkMsUUFBTSxlQUFlLFlBQVk7QUFDakMsU0FBTztBQUFBLElBQ0wsTUFBTSxlQUFlLG9DQUFvQztBQUFBLElBQ3pELFNBQVM7QUFBQSxNQUFDLElBQUk7QUFBQSxNQUNaLHFCQUFxQjtBQUFBLFFBQ25CLFVBQVUsQ0FBQyxLQUFLLFFBQVEsUUFBUSxJQUFJLEdBQUcsa0JBQWtCLENBQUM7QUFBQSxRQUMxRCxVQUFVO0FBQUEsTUFDWixDQUFDO0FBQUEsTUFDRCxjQUFjO0FBQUEsUUFDWixjQUFjLFlBQVk7QUFBQSxNQUM1QixDQUFDO0FBQUEsSUFDSDtBQUFBLElBQ0EsU0FBUztBQUFBLE1BQ1AsT0FBTztBQUFBLFFBQ0wsS0FBSyxLQUFLLFFBQVEsT0FBTztBQUFBO0FBQUEsTUFDM0I7QUFBQSxJQUNGO0FBQUEsSUFDQSxLQUFLO0FBQUEsTUFDSCxxQkFBcUI7QUFBQSxRQUNuQixNQUFNO0FBQUEsVUFDSixtQkFBbUI7QUFBQSxVQUNuQixnQkFBZ0I7QUFBQSxRQUNsQjtBQUFBLE1BQ0Y7QUFBQSxJQUNGO0FBQUE7QUFBQSxJQUVBLFFBQU87QUFBQSxNQUNMLE9BQU07QUFBQSxRQUNKLGFBQWE7QUFBQSxVQUNYLFFBQVE7QUFBQSxVQUNSLGNBQWM7QUFBQSxRQUNoQjtBQUFBLE1BQ0Y7QUFBQSxJQUNGO0FBQUEsRUFDRjtBQUNGLENBQUM7IiwKICAibmFtZXMiOiBbXQp9Cg==
