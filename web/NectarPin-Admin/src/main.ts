import './assets/main.css'

import {createApp} from 'vue'
import {createPinia} from 'pinia'

import ECharts from 'vue-echarts' // 引入ECharts
import 'echarts';

//编辑器[Start]
import VMdEditor from '@kangc/v-md-editor/lib/codemirror-editor';
import '@kangc/v-md-editor/lib/style/codemirror-editor.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';
import hljs from 'highlight.js';
// codemirror 编辑器的相关资源
import Codemirror from 'codemirror';
// mode
import 'codemirror/mode/markdown/markdown';
import 'codemirror/mode/javascript/javascript';
import 'codemirror/mode/css/css';
import 'codemirror/mode/htmlmixed/htmlmixed';
import 'codemirror/mode/vue/vue';
// edit
import 'codemirror/addon/edit/closebrackets';
import 'codemirror/addon/edit/closetag';
import 'codemirror/addon/edit/matchbrackets';
// placeholder
import 'codemirror/addon/display/placeholder';
// active-line
import 'codemirror/addon/selection/active-line';
// scrollbar
import 'codemirror/addon/scroll/simplescrollbars';
import 'codemirror/addon/scroll/simplescrollbars.css';
// [Plugin] [emoji]
import createEmojiPlugin from '@kangc/v-md-editor/lib/plugins/emoji/index';
import '@kangc/v-md-editor/lib/plugins/emoji/emoji.css';
// [Plugin] [Katex]
import createKatexPlugin from '@kangc/v-md-editor/lib/plugins/katex/cdn';
// [Plugin] [Mermaid]
import createMermaidPlugin from '@kangc/v-md-editor/lib/plugins/mermaid/cdn';
import '@kangc/v-md-editor/lib/plugins/mermaid/mermaid.css';
// [Plugin] [TodoList]
import createTodoListPlugin from '@kangc/v-md-editor/lib/plugins/todo-list/index';
import '@kangc/v-md-editor/lib/plugins/todo-list/todo-list.css';
// [Plugin] [LineNumber]
import createLineNumberPlugin from '@kangc/v-md-editor/lib/plugins/line-number/index';
// [Plugin] [CopyCode]
import createCopyCodePlugin from '@kangc/v-md-editor/lib/plugins/copy-code/index';
import '@kangc/v-md-editor/lib/plugins/copy-code/copy-code.css';
// [Plugin] [Align]
import createAlignPlugin from '@kangc/v-md-editor/lib/plugins/align';

// style
import 'codemirror/lib/codemirror.css';
//VMdEditor

VMdEditor.Codemirror = Codemirror;
VMdEditor.use(githubTheme, {
    Hljs: hljs,
});
//插件列表
VMdEditor.use(createEmojiPlugin())
VMdEditor.use(createKatexPlugin())
VMdEditor.use(createMermaidPlugin())
VMdEditor.use(createTodoListPlugin())
VMdEditor.use(createLineNumberPlugin())
VMdEditor.use(createCopyCodePlugin())
VMdEditor.use(createAlignPlugin())
//编辑器[End]

import ArcoVueIcon from '@arco-design/web-vue/es/icon';

import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(ArcoVueIcon);
app.use(createPinia())
app.use(router)
app.component('ECharts', ECharts)
app.use(VMdEditor);
app.mount('#app')
