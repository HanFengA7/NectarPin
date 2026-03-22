<script setup lang="ts">
import md5 from 'crypto-js/md5'
import { computed, onMounted, reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { KeyRound, LoaderCircle, Lock, MonitorSmartphone, UserRound } from 'lucide-vue-next'

import {
  changeUserPassword,
  getCurrentUserProfile,
  listUserSessions,
  revokeOtherUserSessions,
  revokeUserSession,
  updateCurrentUserProfile,
  type UserSessionItem,
} from '@/api/user'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Separator } from '@/components/ui/separator'
import { clearAdminSession, getRefreshToken, patchAdminSessionUser } from '@/lib/admin-auth'
import { RequestError } from '@/utils/req'

const router = useRouter()

const form = reactive({
  username: '',
  nickname: '',
  email: '',
  avatar: '',
})

const pwd = reactive({
  current: '',
  next: '',
  confirm: '',
})

const loading = ref(true)
const loadError = ref('')
const savingProfile = ref(false)
const savingPassword = ref(false)
const profileError = ref('')
const profileSuccess = ref('')
const passwordError = ref('')
const passwordSuccess = ref('')

const sessions = ref<UserSessionItem[]>([])
const sessionsLoading = ref(false)
const sessionsError = ref('')
const sessionsSuccess = ref('')
const revokingId = ref<number | null>(null)
const revokingOthers = ref(false)

const sessionDateFmt = new Intl.DateTimeFormat('zh-CN', {
  dateStyle: 'short',
  timeStyle: 'short',
})

const hasOtherSessions = computed(
  () => sessions.value.some(s => !s.is_current),
)

function formatSessionTime(iso: string) {
  const d = new Date(iso)
  return Number.isNaN(d.getTime()) ? iso : sessionDateFmt.format(d)
}

async function loadSessions() {
  sessionsError.value = ''
  sessionsSuccess.value = ''
  sessionsLoading.value = true
  try {
    const rt = getRefreshToken()
    const res = await listUserSessions(rt || undefined)
    sessions.value = res.data.items ?? []
  }
  catch (e) {
    sessions.value = []
    sessionsError.value =
      e instanceof RequestError ? e.message : '无法加载登录设备列表'
  }
  finally {
    sessionsLoading.value = false
  }
}

onMounted(async () => {
  loadError.value = ''
  loading.value = true
  try {
    const res = await getCurrentUserProfile()
    const u = res.data
    form.username = u.username
    form.nickname = u.nickname ?? ''
    form.email = u.email
    form.avatar = u.avatar ?? ''
  }
  catch (e) {
    loadError.value =
      e instanceof RequestError ? e.message : '无法加载个人资料'
  }
  finally {
    loading.value = false
  }

  if (!loadError.value) {
    await loadSessions()
  }
})

async function handleRevokeSession(item: UserSessionItem) {
  sessionsError.value = ''
  sessionsSuccess.value = ''
  revokingId.value = item.id
  try {
    await revokeUserSession(item.id)
    if (item.is_current) {
      clearAdminSession()
      await router.replace({ name: 'admin-login' })
      return
    }
    sessionsSuccess.value = '已踢下线该设备'
    await loadSessions()
  }
  catch (e) {
    sessionsError.value =
      e instanceof RequestError ? e.message : '操作失败'
  }
  finally {
    revokingId.value = null
  }
}

async function handleRevokeOthers() {
  sessionsError.value = ''
  sessionsSuccess.value = ''
  const rt = getRefreshToken()
  if (!rt) {
    sessionsError.value = '无法读取当前会话，请重新登录。'
    return
  }
  revokingOthers.value = true
  try {
    await revokeOtherUserSessions(rt)
    sessionsSuccess.value = '已踢下线其他设备'
    await loadSessions()
  }
  catch (e) {
    sessionsError.value =
      e instanceof RequestError ? e.message : '操作失败'
  }
  finally {
    revokingOthers.value = false
  }
}

async function handleSaveProfile() {
  profileSuccess.value = ''
  profileError.value = ''
  const email = form.email.trim()
  if (!email) {
    profileError.value = '请填写邮箱。'
    return
  }

  savingProfile.value = true
  try {
    const res = await updateCurrentUserProfile({
      email,
      nickname: form.nickname.trim(),
      avatar: form.avatar.trim(),
    })
    const u = res.data
    patchAdminSessionUser({
      email: u.email,
      nickname: u.nickname,
      avatar: u.avatar,
    })
    form.nickname = u.nickname ?? ''
    form.avatar = u.avatar ?? ''
    profileSuccess.value = '资料已保存'
  }
  catch (e) {
    profileError.value =
      e instanceof RequestError ? e.message : '保存失败，请稍后重试。'
  }
  finally {
    savingProfile.value = false
  }
}

function md5Hex(plain: string) {
  return md5(plain).toString()
}

async function handleChangePassword() {
  passwordSuccess.value = ''
  passwordError.value = ''

  if (!pwd.current || !pwd.next || !pwd.confirm) {
    passwordError.value = '请填写完整密码字段。'
    return
  }
  if (pwd.next.length < 6) {
    passwordError.value = '新密码至少 6 位。'
    return
  }
  if (pwd.next !== pwd.confirm) {
    passwordError.value = '两次输入的新密码不一致。'
    return
  }

  savingPassword.value = true
  try {
    await changeUserPassword({
      old_password: md5Hex(pwd.current),
      new_password: md5Hex(pwd.next),
    })
    passwordSuccess.value = '密码已更新，请牢记新密码。'
    pwd.current = ''
    pwd.next = ''
    pwd.confirm = ''
  }
  catch (e) {
    passwordError.value =
      e instanceof RequestError ? e.message : '修改密码失败，请稍后重试。'
  }
  finally {
    savingPassword.value = false
  }
}
</script>

<template>
  <section class="min-h-svh bg-muted/30">
    <div class="mx-auto w-full max-w-6xl px-4 py-6 sm:px-6 lg:px-8 lg:py-8">
      <!-- 页头 -->
      <div
        class="mb-8 flex flex-col gap-4 rounded-2xl border bg-background/80 p-6 shadow-sm backdrop-blur-sm sm:flex-row sm:items-center sm:justify-between sm:p-8"
      >
        <div class="flex items-start gap-4">
          <div class="flex size-12 shrink-0 items-center justify-center rounded-2xl bg-primary/10 text-primary">
            <UserRound class="size-6" />
          </div>
          <div>
            <h1 class="text-2xl font-semibold tracking-tight">
              个人资料与账号安全
            </h1>
            <p class="mt-1 max-w-2xl text-sm leading-relaxed text-muted-foreground">
              更新资料与密码；下方可查看登录设备并踢下线其他客户端。用户名不可更改。
            </p>
          </div>
        </div>
      </div>

      <div
        v-if="loading"
        class="flex items-center justify-center gap-2 rounded-2xl border bg-background py-20 text-sm text-muted-foreground shadow-sm"
      >
        <LoaderCircle class="size-4 animate-spin" />
        加载中…
      </div>

      <div
        v-else-if="loadError"
        class="rounded-2xl border border-destructive/30 bg-destructive/5 px-6 py-12 text-center text-sm text-destructive shadow-sm"
      >
        {{ loadError }}
      </div>

      <div
        v-else
        class="space-y-6 lg:space-y-8"
      >
        <div
          class="grid gap-6 lg:grid-cols-2 lg:items-start lg:gap-8"
        >
        <!-- 基本资料 -->
        <div class="rounded-2xl border bg-background p-6 shadow-sm sm:p-8">
          <div class="mb-6 flex items-center gap-2 border-b border-border/80 pb-4">
            <UserRound class="size-5 text-muted-foreground" />
            <h2 class="text-lg font-semibold">
              基本资料
            </h2>
          </div>

          <form class="space-y-5" @submit.prevent="handleSaveProfile">
            <p
              v-if="profileError"
              class="rounded-lg border border-destructive/30 bg-destructive/10 px-3 py-2 text-sm text-destructive"
            >
              {{ profileError }}
            </p>
            <p
              v-if="profileSuccess"
              class="rounded-lg border border-emerald-200 bg-emerald-50 px-3 py-2 text-sm text-emerald-800 dark:border-emerald-900 dark:bg-emerald-950/40 dark:text-emerald-200"
            >
              {{ profileSuccess }}
            </p>

            <div class="space-y-2">
              <Label for="profile-username">用户名</Label>
              <div
                id="profile-username"
                class="flex h-10 items-center rounded-md border border-input bg-muted/40 px-3 text-sm text-muted-foreground"
              >
                {{ form.username || '—' }}
              </div>
            </div>

            <div class="space-y-2">
              <Label for="profile-nickname">昵称</Label>
              <Input
                id="profile-nickname"
                v-model="form.nickname"
                type="text"
                autocomplete="nickname"
                maxlength="100"
                class="h-10"
                placeholder="显示名称"
              />
            </div>

            <div class="space-y-2">
              <Label for="profile-email">邮箱</Label>
              <Input
                id="profile-email"
                v-model="form.email"
                type="email"
                autocomplete="email"
                required
                class="h-10"
              />
            </div>

            <div class="space-y-2">
              <Label for="profile-avatar">头像 URL</Label>
              <Input
                id="profile-avatar"
                v-model="form.avatar"
                type="url"
                maxlength="500"
                class="h-10"
                placeholder="https://"
              />
            </div>

            <Button type="submit" class="w-full sm:w-auto" :disabled="savingProfile">
              <LoaderCircle
                v-if="savingProfile"
                class="size-4 animate-spin"
              />
              {{ savingProfile ? '保存中…' : '保存资料' }}
            </Button>
          </form>
        </div>

        <!-- 修改密码 -->
        <div class="rounded-2xl border bg-background p-6 shadow-sm sm:p-8">
          <div class="mb-6 flex items-center gap-2 border-b border-border/80 pb-4">
            <Lock class="size-5 text-muted-foreground" />
            <h2 class="text-lg font-semibold">
              修改密码
            </h2>
          </div>

          <p class="mb-5 text-sm text-muted-foreground">
            新密码至少 6 位字符。
          </p>

          <form class="space-y-5" @submit.prevent="handleChangePassword">
            <p
              v-if="passwordError"
              class="rounded-lg border border-destructive/30 bg-destructive/10 px-3 py-2 text-sm text-destructive"
            >
              {{ passwordError }}
            </p>
            <p
              v-if="passwordSuccess"
              class="rounded-lg border border-emerald-200 bg-emerald-50 px-3 py-2 text-sm text-emerald-800 dark:border-emerald-900 dark:bg-emerald-950/40 dark:text-emerald-200"
            >
              {{ passwordSuccess }}
            </p>

            <div class="space-y-2">
              <Label for="pwd-current">当前密码</Label>
              <Input
                id="pwd-current"
                v-model="pwd.current"
                type="password"
                autocomplete="current-password"
                class="h-10"
              />
            </div>

            <Separator class="bg-border/60" />

            <div class="space-y-2">
              <Label for="pwd-next">新密码</Label>
              <Input
                id="pwd-next"
                v-model="pwd.next"
                type="password"
                autocomplete="new-password"
                class="h-10"
                placeholder="至少 6 位"
              />
            </div>

            <div class="space-y-2">
              <Label for="pwd-confirm">确认新密码</Label>
              <Input
                id="pwd-confirm"
                v-model="pwd.confirm"
                type="password"
                autocomplete="new-password"
                class="h-10"
              />
            </div>

            <Button
              type="submit"
              variant="secondary"
              class="w-full sm:w-auto"
              :disabled="savingPassword"
            >
              <KeyRound v-if="!savingPassword" class="size-4" />
              <LoaderCircle
                v-else
                class="size-4 animate-spin"
              />
              {{ savingPassword ? '提交中…' : '更新密码' }}
            </Button>
          </form>
        </div>
        </div>

        <!-- 登录设备 / 踢下线 -->
        <div class="rounded-2xl border bg-background p-6 shadow-sm sm:p-8">
          <div class="mb-4 flex flex-col gap-4 border-b border-border/80 pb-4 sm:flex-row sm:items-center sm:justify-between">
            <div class="flex items-center gap-2">
              <MonitorSmartphone class="size-5 text-muted-foreground" />
              <h2 class="text-lg font-semibold">
                登录设备
              </h2>
            </div>
            <Button
              v-if="hasOtherSessions"
              type="button"
              variant="outline"
              size="sm"
              :disabled="revokingOthers || sessionsLoading"
              @click="handleRevokeOthers"
            >
              <LoaderCircle
                v-if="revokingOthers"
                class="size-4 animate-spin"
              />
              {{ revokingOthers ? '处理中…' : '踢掉其他设备' }}
            </Button>
          </div>

          <p class="mb-4 text-sm text-muted-foreground">
            每个浏览器登录会产生一条会话。撤销后该端无法再刷新令牌，需重新登录。
          </p>

          <p
            v-if="sessionsError"
            class="mb-4 rounded-lg border border-destructive/30 bg-destructive/10 px-3 py-2 text-sm text-destructive"
          >
            {{ sessionsError }}
          </p>
          <p
            v-if="sessionsSuccess"
            class="mb-4 rounded-lg border border-emerald-200 bg-emerald-50 px-3 py-2 text-sm text-emerald-800 dark:border-emerald-900 dark:bg-emerald-950/40 dark:text-emerald-200"
          >
            {{ sessionsSuccess }}
          </p>

          <div
            v-if="sessionsLoading"
            class="flex items-center gap-2 py-8 text-sm text-muted-foreground"
          >
            <LoaderCircle class="size-4 animate-spin" />
            加载设备列表…
          </div>

          <div
            v-else-if="sessions.length === 0"
            class="py-8 text-center text-sm text-muted-foreground"
          >
            暂无会话记录（需重新登录后才会出现）
          </div>

          <div
            v-else
            class="overflow-x-auto rounded-lg border border-border/60"
          >
            <table class="w-full min-w-[520px] text-left text-sm">
              <thead class="border-b bg-muted/40 text-muted-foreground">
                <tr>
                  <th class="px-4 py-3 font-medium">
                    设备
                  </th>
                  <th class="px-4 py-3 font-medium">
                    登录时间
                  </th>
                  <th class="px-4 py-3 font-medium">
                    过期时间
                  </th>
                  <th class="w-32 px-4 py-3 font-medium">
                    操作
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="s in sessions"
                  :key="s.id"
                  class="border-b border-border/50 last:border-0"
                >
                  <td class="px-4 py-3">
                    <span
                      v-if="s.is_current"
                      class="inline-flex items-center rounded-full bg-primary/10 px-2 py-0.5 text-xs font-medium text-primary"
                    >本设备</span>
                    <span v-else class="text-muted-foreground">其他设备</span>
                  </td>
                  <td class="px-4 py-3 text-muted-foreground">
                    {{ formatSessionTime(s.created_at) }}
                  </td>
                  <td class="px-4 py-3 text-muted-foreground">
                    {{ formatSessionTime(s.expires_at) }}
                  </td>
                  <td class="px-4 py-3">
                    <Button
                      type="button"
                      variant="ghost"
                      size="sm"
                      class="h-8 text-destructive hover:bg-destructive/10 hover:text-destructive"
                      :disabled="revokingId === s.id"
                      @click="handleRevokeSession(s)"
                    >
                      <LoaderCircle
                        v-if="revokingId === s.id"
                        class="size-4 animate-spin"
                      />
                      <span v-else>{{ s.is_current ? '退出' : '踢下线' }}</span>
                    </Button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
