import { get, post } from '@/utils/req'

// ── 前台友链页 ──

export interface PublicFriendLinkItem {
  title: string
  url: string
  description: string
  avatar_url: string
  sort_order: number
}

export interface PublicFriendSection {
  category_id: number | null
  category_name: string
  category_description: string
  sort_order: number
  links: PublicFriendLinkItem[]
}

export interface PublicFriendPageData {
  intro_html: string
  sections: PublicFriendSection[]
}

export function getPublicFriendPage() {
  return get<PublicFriendPageData>('/api/public/link/v1/page')
}

// ── 友链页设置（后台） ──

export interface FriendLinkPageSettings {
  id: number
  intro_html: string
  created_at: string
  updated_at: string
}

export function getFriendLinkPageSettings() {
  return get<FriendLinkPageSettings>('/api/protected/link/v1/page/settings')
}

export function saveFriendLinkPageSettings(body: { intro_html: string }) {
  return post<FriendLinkPageSettings>('/api/protected/link/v1/page/settings', body)
}

// ── 友链分组 ──

export interface FriendLinkCategoryItem {
  id: number
  name: string
  slug: string
  description: string
  sort_order: number
  created_at: string
  updated_at: string
}

export interface FriendLinkCategoryListData {
  items: FriendLinkCategoryItem[]
  total: number
}

export interface CreateFriendLinkCategoryPayload {
  name: string
  slug?: string
  description?: string
  sort_order?: number
}

export interface UpdateFriendLinkCategoryPayload {
  name?: string
  slug?: string
  description?: string
  sort_order?: number
}

export function listFriendLinkCategories() {
  return get<FriendLinkCategoryListData>('/api/protected/link/v1/category/list')
}

export function createFriendLinkCategory(payload: CreateFriendLinkCategoryPayload) {
  return post<FriendLinkCategoryItem, CreateFriendLinkCategoryPayload>(
    '/api/protected/link/v1/category/add',
    payload,
  )
}

export function updateFriendLinkCategory(id: number, payload: UpdateFriendLinkCategoryPayload) {
  return post<FriendLinkCategoryItem, UpdateFriendLinkCategoryPayload>(
    `/api/protected/link/v1/category/update/${id}`,
    payload,
  )
}

export function deleteFriendLinkCategory(id: number) {
  return post<null>(`/api/protected/link/v1/category/delete/${id}`)
}

// ── 友链条目（后台） ──

export interface FriendLinkAdminItem {
  id: number
  category_id: number | null
  title: string
  url: string
  description: string
  avatar_url: string
  sort_order: number
  is_enabled: boolean
  created_at: string
  updated_at: string
}

export interface FriendLinkAdminListData {
  items: FriendLinkAdminItem[]
  total: number
}

export interface CreateFriendLinkPayload {
  category_id?: number | null
  title: string
  url: string
  description?: string
  avatar_url?: string
  sort_order?: number
  is_enabled?: boolean
}

export interface UpdateFriendLinkPayload {
  category_id?: number | null
  title?: string
  url?: string
  description?: string
  avatar_url?: string
  sort_order?: number
  is_enabled?: boolean
}

export function listFriendLinksAdmin() {
  return get<FriendLinkAdminListData>('/api/protected/link/v1/list')
}

export function createFriendLink(payload: CreateFriendLinkPayload) {
  return post<FriendLinkAdminItem, CreateFriendLinkPayload>('/api/protected/link/v1/add', payload)
}

export function updateFriendLink(id: number, payload: UpdateFriendLinkPayload) {
  return post<FriendLinkAdminItem, UpdateFriendLinkPayload>(
    `/api/protected/link/v1/update/${id}`,
    payload,
  )
}

export function deleteFriendLink(id: number) {
  return post<null>(`/api/protected/link/v1/delete/${id}`)
}
