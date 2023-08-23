import type { Writable } from 'svelte/store'
import { writable } from 'svelte/store'

export const modeStore: Writable<'host' | 'fight' | 'client' | null> = writable(null)
