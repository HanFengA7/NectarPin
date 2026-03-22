<script setup lang="ts">
import { computed, onMounted, reactive, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { LoaderCircle, LockKeyhole, ShieldCheck, UserRound } from 'lucide-vue-next'

import { loginUser } from '@/api/user'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  buildAdminSessionFromLogin,
  isAdminAuthenticated,
  saveAdminSession,
} from '@/lib/admin-auth'
import { useSiteStore } from '@/stores/site'
import { RequestError } from '@/utils/req'

const router = useRouter()
const route = useRoute()
const siteStore = useSiteStore()

const adminPortalTitle = computed(() => `${siteStore.siteName} 管理后台`)

const form = reactive({
  account: '',
  password: '',
})

const isSubmitting = ref(false)
const errorMessage = ref('')

const redirectTarget = computed(() => {
  const redirect = route.query.redirect

  if (typeof redirect === 'string' && redirect.startsWith('/admin')) {
    return redirect
  }

  return '/admin/dashboard'
})

onMounted(() => {
  void siteStore.hydrateFromApi()
  if (isAdminAuthenticated()) {
    router.replace(redirectTarget.value)
  }
})

async function handleSubmit() {
  errorMessage.value = ''
  isSubmitting.value = true

  if (!form.account.trim() || !form.password.trim()) {
    errorMessage.value = '请输入账号和密码。'
    isSubmitting.value = false
    return
  }

  try {
    const response = await loginUser({
      account: form.account,
      password: form.password,
    })

    saveAdminSession(buildAdminSessionFromLogin(response.data))
    await router.replace(redirectTarget.value)
  } catch (error) {
    errorMessage.value = error instanceof RequestError ? error.message : '登录失败，请稍后重试。'
  } finally {
    isSubmitting.value = false
  }
}
</script>

<template>
  <section
    class="relative flex min-h-screen items-center justify-center overflow-hidden bg-[radial-gradient(circle_at_top,_rgba(24,24,27,0.05),_transparent_35%),linear-gradient(135deg,_rgb(250,250,250),_rgb(244,244,245))] px-4 py-10"
  >
    <div
      class="absolute inset-0 opacity-40"
      style="
        background-image:
          linear-gradient(rgba(24, 24, 27, 0.05) 1px, transparent 1px),
          linear-gradient(90deg, rgba(24, 24, 27, 0.05) 1px, transparent 1px);
        background-size: 32px 32px;
      "
    />

    <div
      class="relative grid w-full max-w-5xl overflow-hidden rounded-3xl border bg-background shadow-2xl lg:grid-cols-[1.1fr_0.9fr]"
    >
      <div class="hidden bg-primary px-10 py-12 text-primary-foreground lg:flex lg:items-center">
        <div class="space-y-4">
          <div
            class="inline-flex size-12 items-center justify-center rounded-2xl bg-primary-foreground/10"
          >
            <ShieldCheck class="size-6" />
          </div>
          <div class="space-y-2">
            <p class="text-sm tracking-[0.3em] text-primary-foreground/70 uppercase">
              Admin Portal
            </p>
            <h1 class="text-3xl font-semibold leading-tight">
              {{ adminPortalTitle }}
            </h1>
            <p class="max-w-sm text-sm leading-6 text-primary-foreground/80">
              登录后即可进入后台概览、内容管理与后续扩展的运营功能模块。
            </p>
          </div>
        </div>
      </div>

      <div class="flex items-center bg-background px-6 py-8 sm:px-10 sm:py-12">
        <div class="mx-auto w-full max-w-md space-y-8">
          <div class="space-y-2">
            <p class="text-sm font-medium text-muted-foreground">后台登录</p>
            <h2 class="text-3xl font-semibold tracking-tight">欢迎回来</h2>
            <p class="text-sm leading-6 text-muted-foreground">
              输入用户名或邮箱与密码，进入 {{ siteStore.siteName }} 管理后台。
            </p>
          </div>

          <form class="space-y-5" @submit.prevent="handleSubmit">
            <div class="space-y-2">
              <label for="account" class="text-sm font-medium">账号</label>
              <div class="relative">
                <UserRound
                  class="pointer-events-none absolute top-1/2 left-3 size-4 -translate-y-1/2 text-muted-foreground"
                />
                <Input
                  id="account"
                  v-model="form.account"
                  type="text"
                  autocomplete="username"
                  placeholder="请输入用户名或邮箱"
                  class="pl-9"
                />
              </div>
            </div>

            <div class="space-y-2">
              <label for="password" class="text-sm font-medium">登录密码</label>
              <div class="relative">
                <LockKeyhole
                  class="pointer-events-none absolute top-1/2 left-3 size-4 -translate-y-1/2 text-muted-foreground"
                />
                <Input
                  id="password"
                  v-model="form.password"
                  type="password"
                  autocomplete="current-password"
                  placeholder="请输入登录密码"
                  class="pl-9"
                />
              </div>
            </div>

            <p
              v-if="errorMessage"
              class="rounded-md border border-destructive/20 bg-destructive/8 px-3 py-2 text-sm text-destructive"
            >
              {{ errorMessage }}
            </p>

            <Button type="submit" class="w-full" size="lg" :disabled="isSubmitting">
              <LoaderCircle v-if="isSubmitting" class="size-4 animate-spin" />
              {{ isSubmitting ? '登录中...' : '登录后台' }}
            </Button>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>
