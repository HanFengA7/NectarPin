/**
 * 工具函数模块
 * @module utils/index
 */

import md5 from 'blueimp-md5'

/**
 * 对字符串进行 MD5 加密
 * @param str - 需要加密的字符串
 * @returns MD5 加密后的 32 位小写字符串
 */
export function md5Encrypt(str: string): string {
  return md5(str)
}

/**
 * 对密码进行 MD5 加密
 * @param password - 原始密码
 * @returns MD5 加密后的 32 位小写字符串
 */
export function encryptPassword(password: string): string {
  return md5(password)
}
