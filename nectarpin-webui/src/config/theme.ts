import type { GlobalThemeOverrides } from 'naive-ui'

// 浅色主题配置
export const lightThemeOverrides: GlobalThemeOverrides = {
  common: {
    // 主色调 - 可以修改为你想要的颜色
    primaryColor: '#2080f0', // 默认绿色，可以改为其他颜色，如 '#2080f0' (蓝色)
    primaryColorHover: '#4098fc',
    primaryColorPressed: '#1060c9',
    primaryColorSuppl: '#4098fc',
    
    // 成功色
    successColor: '#18a058',
    successColorHover: '#36ad6a',
    successColorPressed: '#0c7a43',
    
    // 警告色
    warningColor: '#f0a020',
    warningColorHover: '#fcb040',
    warningColorPressed: '#c97c10',
    
    // 错误色
    errorColor: '#d03050',
    errorColorHover: '#de576d',
    errorColorPressed: '#ab1f3f',
    
    // 信息色
    infoColor: '#2080f0',
    infoColorHover: '#4098fc',
    infoColorPressed: '#1060c9',
  },
  Button: {
    // 按钮相关配置
    borderRadius: '4px',
  },
  Input: {
    // 输入框相关配置
    borderRadius: '4px',
  },
  Card: {
    // 卡片相关配置
    borderRadius: '8px',
  },
}

// 暗色主题配置
export const darkThemeOverrides: GlobalThemeOverrides = {
  common: {
    // 暗色主题的主色调
    primaryColor: '#63e2b7',
    primaryColorHover: '#7fe7c4',
    primaryColorPressed: '#5acea7',
    primaryColorSuppl: '#7fe7c4',
  },
}
