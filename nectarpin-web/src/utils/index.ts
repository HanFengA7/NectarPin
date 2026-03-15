/**
 * 工具函数模块
 * @module utils/index
 */

import md5 from 'blueimp-md5'

/**
 * 对字符串进行 MD5 加密
 * @param str - 需要加密的字符串
 * @returns MD5 加密后的 32 位小写字符串
 * @description
 * 前端使用 MD5 对密码进行第一次加密，主要目的：
 * 1. 避免明文密码在网络传输中暴露
 * 2. 配合 HTTPS 双重保护传输安全
 * 3. 后端会对 MD5 哈希值进行 bcrypt 二次加密存储
 *
 * 安全说明：
 * - MD5 仅用于传输层保护，不用于最终存储
 * - 必须配合 HTTPS 使用确保传输安全
 * - 后端使用 bcrypt (cost=12) 进行安全存储
 */
export function md5Encrypt(str: string): string {
  return md5(str)
}

/**
 * 对密码进行 MD5 加密（前端第一层加密）
 * @param password - 原始密码
 * @returns MD5 加密后的 32 位小写字符串
 * @description
 * 密码加密流程：
 * 1. 前端：原始密码 → MD5 哈希 (32位)
 * 2. 传输：HTTPS 加密传输 MD5 哈希值
 * 3. 后端：MD5 哈希 → bcrypt 哈希 (60位) → 数据库存储
 *
 * 注意事项：
 * - 输入密码长度建议 6-50 个字符
 * - 输出固定为 32 位小写十六进制字符串
 * - 后端会验证输入长度必须为 32 位
 */
export function encryptPassword(password: string): string {
  return md5(password)
}

/**
 * 验证密码格式
 * @param password - 原始密码
 * @returns 验证结果对象
 * @description
 * 密码规则：
 * - 长度：6-50 个字符
 * - 必须包含至少一个字母
 * - 可选包含数字和特殊字符
 */
export function validatePassword(password: string): { valid: boolean; message: string } {
  if (!password) {
    return { valid: false, message: '请输入密码' }
  }
  if (password.length < 6) {
    return { valid: false, message: '密码长度至少 6 个字符' }
  }
  if (password.length > 50) {
    return { valid: false, message: '密码长度不能超过 50 个字符' }
  }
  if (!/[a-zA-Z]/.test(password)) {
    return { valid: false, message: '密码必须包含至少一个字母' }
  }
  return { valid: true, message: '' }
}
