import { defineStore } from 'pinia'
import { ref } from 'vue'

import { getPublicSiteHome, type SiteHomePayload } from '@/api/site'

const DEFAULT_SITE = 'NectarPin'

export const useSiteStore = defineStore('site', () => {
  const siteName = ref(DEFAULT_SITE)
  const avatarUrl = ref('')
  const avatarInitials = ref('')
  const heroName = ref(DEFAULT_SITE)
  const footerIcpText = ref('')
  const footerIcpHref = ref('')
  const footerSince = ref('')
  const footerPsbText = ref('')
  const footerPsbHref = ref('')
  let inflight: Promise<void> | null = null

  function applyFromHomePayload(data: SiteHomePayload) {
    const n = data.site_name?.trim()
    siteName.value = n || DEFAULT_SITE
    avatarUrl.value = (data.avatar_url ?? '').trim()
    avatarInitials.value = (data.avatar_initials ?? '').trim()
    heroName.value = (data.hero_name ?? '').trim() || siteName.value
    footerIcpText.value = (data.footer_icp_text ?? '').trim()
    footerIcpHref.value = (data.footer_icp_href ?? '').trim()
    footerSince.value = (data.footer_since ?? '').trim()
    footerPsbText.value = (data.footer_psb_text ?? '').trim()
    footerPsbHref.value = (data.footer_psb_href ?? '').trim()
  }

  async function hydrateFromApi() {
    if (inflight) return inflight
    inflight = (async () => {
      try {
        const { data } = await getPublicSiteHome()
        applyFromHomePayload(data)
      } finally {
        inflight = null
      }
    })()
    return inflight
  }

  return {
    siteName,
    avatarUrl,
    avatarInitials,
    heroName,
    footerIcpText,
    footerIcpHref,
    footerSince,
    footerPsbText,
    footerPsbHref,
    hydrateFromApi,
    applyFromHomePayload,
  }
})
