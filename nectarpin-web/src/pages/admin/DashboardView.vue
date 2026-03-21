<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { FileText, PenSquare } from 'lucide-vue-next'

import { Button } from '@/components/ui/button'
</script>

<template>
  <section class="min-h-svh bg-muted/20">
    <div class="flex w-full flex-col gap-6 px-4 py-6 sm:px-6 lg:px-8">
      <header class="rounded-2xl border bg-background p-6 shadow-sm">
        <p class="text-sm font-medium text-muted-foreground">后台概览</p>
        <h1 class="mt-1 text-3xl font-semibold tracking-tight">欢迎进入 NectarPin 管理后台</h1>
        <p class="mt-2 max-w-2xl text-sm leading-6 text-muted-foreground">
          当前后台已支持文章列表、新增文章与编辑文章，你可以从这里快速进入内容管理。
        </p>
      </header>

      <div class="grid gap-4 md:grid-cols-2">
        <section class="rounded-2xl border bg-background p-6 shadow-sm">
          <div class="mb-4 inline-flex rounded-xl bg-primary/8 p-3 text-primary">
            <FileText class="size-5" />
          </div>
          <h2 class="text-lg font-semibold">查看文章列表</h2>
          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            浏览当前账号下的全部文章，并按状态筛选、编辑或删除已有内容。
          </p>
          <Button as-child class="mt-5">
            <RouterLink :to="{ name: 'admin-articles' }">进入文章列表</RouterLink>
          </Button>
        </section>

        <section class="rounded-2xl border bg-background p-6 shadow-sm">
          <div class="mb-4 inline-flex rounded-xl bg-primary/8 p-3 text-primary">
            <PenSquare class="size-5" />
          </div>
          <h2 class="text-lg font-semibold">开始写文章</h2>
          <p class="mt-2 text-sm leading-6 text-muted-foreground">
            直接新建文章并保存为草稿，或者填写完整内容后立即发布。
          </p>
          <Button as-child variant="outline" class="mt-5">
            <RouterLink :to="{ name: 'admin-article-create' }">新增文章</RouterLink>
          </Button>
        </section>
      </div>
    </div>
  </section>
</template>