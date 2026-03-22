/// <reference types="vite/client" />

import 'vue'

declare module 'vue' {
  interface OptgroupHTMLAttributes {
    'data-slot'?: string
  }
  interface OptionHTMLAttributes {
    'data-slot'?: string
  }
}

import 'vue-router'

declare module 'vue-router' {
  interface RouteMeta {
    /** 浏览器标签页标题片段（不含站点名后缀） */
    title?: string
    /** 由页面内根据数据设置 `document.title`，路由层不覆盖 */
    titleFromPage?: boolean
  }
}
