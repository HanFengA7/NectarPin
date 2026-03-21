import { get, post } from '@/utils/req'

export interface ArticleItem {
  id: number
  author_id: number
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
}

export interface UpdateArticlePayload {
  title?: string
  slug?: string
  summary?: string
  content?: string
  cover_image?: string
  status?: number
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
