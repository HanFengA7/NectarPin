import { defineStore } from 'pinia'
import { ref } from 'vue'

import { getPublicSiteHome, type SiteHomePayload } from '@/api/site'

const DEFAULT_SITE = 'NectarPin'

export const useSiteStore = defineStore('site', () => {
  const siteName = ref(DEFAULT_SITE)
  let inflight: Promise<void> | null = null

  function applyFromHomePayload(data: Pick<SiteHomePayload, 'site_name'>) {
    const n = data.site_name?.trim()
    siteName.value = n || DEFAULT_SITE
  }

  async function hydrateFromApi() {
    if (inflight)
      return inflight
    inflight = (async () => {
      try {
        const { data } = await getPublicSiteHome()
        applyFromHomePayload(data)
      }
      finally {
        inflight = null
      }
    })()
    return inflight
  }

  return { siteName, hydrateFromApi, applyFromHomePayload }
})
