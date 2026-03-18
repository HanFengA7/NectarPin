import request from '@/utils/request'
import type { ApiResponse } from '@/utils/request'

export interface Article {
  id?: number
  title: string
  slug?: string
  summary?: string
  content: string
  cover_image?: string
  status: number
  author_id?: number
  view_count?: number
  created_at?: string
  updated_at?: string
}

export interface ArticleListItem {
  items: Article[]
  total: number
  page: number
  page_size: number
}

export interface ArticleQueryParams {
  page?: number
  page_size?: number
  author_id?: number
  status?: number
}

export const articleApi = {
  list: (params?: ArticleQueryParams): Promise<ApiResponse<ArticleListItem>> => {
    return request.get('/public/article/v1/list', { params })
  },

  listForAdmin: (params?: ArticleQueryParams): Promise<ApiResponse<ArticleListItem>> => {
    return request.get('/protected/article/v1/list', { params })
  },

  getById: (id: number): Promise<ApiResponse<Article>> => {
    return request.get(`/protected/article/v1/infoById/${id}`)
  },

  getBySlug: (slug: string): Promise<ApiResponse<Article>> => {
    return request.get(`/public/article/v1/infoBySlug/${slug}`)
  },

  create: (data: Partial<Article>): Promise<ApiResponse<Article>> => {
    return request.post('/protected/article/v1/add', data)
  },

  update: (id: number, data: Partial<Article>): Promise<ApiResponse<Article>> => {
    return request.post(`/protected/article/v1/update/${id}`, data)
  },

  delete: (id: number): Promise<ApiResponse<void>> => {
    return request.post(`/protected/article/v1/delete/${id}`)
  },
}
