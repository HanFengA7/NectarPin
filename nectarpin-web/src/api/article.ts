import { get, post } from '@/utils/req'

// ── 文章 ──

export interface ArticleItem {
  id: number
  author_id: number
  category_id: number | null
  title: string
  slug: string
  summary: string
  content: string
  cover_image: string
  status: number
  view_count: number
  published_at: string | null
  created_at: string
  updated_at: string
}

export interface ArticleListData {
  items: ArticleItem[]
  total: number
  page: number
  page_size: number
}

export interface ListAdminArticlesParams {
  page?: number
  page_size?: number
  status?: number
}

export interface CreateArticlePayload {
  title: string
  slug?: string
  summary?: string
  content: string
  cover_image?: string
  status: number
  category_id?: number | null
  tag_ids?: number[]
}

export interface UpdateArticlePayload {
  title?: string
  slug?: string
  summary?: string
  content?: string
  cover_image?: string
  status?: number
  category_id?: number | null
  tag_ids?: number[]
}

export function listAdminArticles(params: ListAdminArticlesParams) {
  return get<ArticleListData>('/api/protected/article/v1/list', {
    params,
  })
}

export function getAdminArticleById(id: number) {
  return get<ArticleItem>(`/api/protected/article/v1/infoById/${id}`)
}

export function createArticle(payload: CreateArticlePayload) {
  return post<ArticleItem, CreateArticlePayload>('/api/protected/article/v1/add', payload)
}

export function updateArticle(id: number, payload: UpdateArticlePayload) {
  return post<ArticleItem, UpdateArticlePayload>(`/api/protected/article/v1/update/${id}`, payload)
}

export function deleteArticle(id: number) {
  return post<null>(`/api/protected/article/v1/delete/${id}`)
}

// ── 文章分类 ──

export interface ArticleCategoryItem {
  id: number
  name: string
  slug: string
  description: string
  sort_order: number
  article_count: number
  created_at: string
  updated_at: string
}

export interface ArticleCategoryListData {
  items: ArticleCategoryItem[]
  total: number
}

export interface CreateCategoryPayload {
  name: string
  slug?: string
  description?: string
  sort_order?: number
}

export interface UpdateCategoryPayload {
  name?: string
  slug?: string
  description?: string
  sort_order?: number
}

export function listCategories() {
  return get<ArticleCategoryListData>('/api/protected/article/v1/category/list')
}

export function listPublicCategories() {
  return get<ArticleCategoryListData>('/api/public/article/v1/category/list')
}

export function createCategory(payload: CreateCategoryPayload) {
  return post<ArticleCategoryItem, CreateCategoryPayload>(
    '/api/protected/article/v1/category/add',
    payload,
  )
}

export function updateCategory(id: number, payload: UpdateCategoryPayload) {
  return post<ArticleCategoryItem, UpdateCategoryPayload>(
    `/api/protected/article/v1/category/update/${id}`,
    payload,
  )
}

export function deleteCategory(id: number) {
  return post<null>(`/api/protected/article/v1/category/delete/${id}`)
}

// ── 文章标签 ──

export interface ArticleTagItem {
  id: number
  name: string
  slug: string
  article_count: number
  created_at: string
  updated_at: string
}

export interface ArticleTagListData {
  items: ArticleTagItem[]
  total: number
}

export interface CreateTagPayload {
  name: string
  slug?: string
}

export interface UpdateTagPayload {
  name?: string
  slug?: string
}

export function listTags() {
  return get<ArticleTagListData>('/api/protected/article/v1/tag/list')
}

export function listPublicTags() {
  return get<ArticleTagListData>('/api/public/article/v1/tag/list')
}

export function createTag(payload: CreateTagPayload) {
  return post<ArticleTagItem, CreateTagPayload>(
    '/api/protected/article/v1/tag/add',
    payload,
  )
}

export function updateTag(id: number, payload: UpdateTagPayload) {
  return post<ArticleTagItem, UpdateTagPayload>(
    `/api/protected/article/v1/tag/update/${id}`,
    payload,
  )
}

export function deleteTag(id: number) {
  return post<null>(`/api/protected/article/v1/tag/delete/${id}`)
}

export interface ArticleTagIdsData {
  tag_ids: number[]
}

export function getArticleTagIds(articleId: number) {
  return get<ArticleTagIdsData>(
    `/api/protected/article/v1/tag/listByArticle/${articleId}`,
  )
}
