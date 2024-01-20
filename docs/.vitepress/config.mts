import {defineConfig} from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
    title: "NectarPin Docs",
    description: "A moment like nailing nectar!",


    head: [
        ['link', {rel: 'icon', href: '/nectarpin.png'}],
        ['link', {rel: 'stylesheet', href: '/styles/custom.css'}]
    ],

    locales: {
        root: {
            label: '简体中文',
            lang: 'zh',
        },
        fr: {
            label: 'English',
            lang: 'en', // 可选，将作为 `lang` 属性添加到 `html` 标签中
            link: '/en' // 默认 /fr/ -- 显示在导航栏翻译菜单上，可以是外部的

            // 其余 locale 特定属性...
        }
    },

    themeConfig: {
        // https://vitepress.dev/reference/default-theme-config
        //logo: '/nectarpin.png',
        search: {
            provider: 'local'
        },
        nav: [
            {text: '首页', link: '/'},
            {text: '使用指南', link: '/guide/'},
            {text: '开发文档', link: '/markdown-examples'},
            {
                text: '关于',
                items: [
                    {text: 'Markdown Examples', link: '/markdown-examples'},
                    {text: 'Runtime API Examples', link: '/api-examples'}
                ]
            },
            {
                text: 'V1.0.0-alpha1',
                items: [
                    {text: 'Markdown Examples', link: '/markdown-examples'},
                    {text: 'Runtime API Examples', link: '/api-examples'}
                ]
            }
        ],

        sidebar: {
            // 当用户位于 `guide` 目录时，会显示此侧边栏
            '/guide/': [
                {
                    text: '介绍',
                    collapsed: true,
                    items: [
                        { text: '什么是 NectarPin？', link: '/guide/' },
                        {   text: '安装部署',
                            collapsed:false,
                            items:[
                                { text: 'Docker 部署', link: '/guide/install-docker' },
                                { text: '二进制编译 部署', link: '/guide/install-2code' },
                                { text: '源码 部署', link: '/guide/install-code' },
                            ]
                        }
                    ]
                }
            ],

            // 当用户位于 `config` 目录时，会显示此侧边栏
            '/config/': [
                {
                    text: 'Config',
                    items: [
                        { text: 'Index', link: '/config/' },
                        { text: 'Three', link: '/config/three' },
                        { text: 'Four', link: '/config/four' }
                    ]
                }
            ]
        },

        socialLinks: [
            {icon: 'github', link: 'https://github.com/HanFengA7/NectarPin'}
        ],

        footer: {
            copyright: 'Copyright © 2015 - 2024 Lychape'
        }

    }
})
