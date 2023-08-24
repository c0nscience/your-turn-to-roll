import type { Writable } from 'svelte/store'
import { writable } from 'svelte/store'

export const modeStore: Writable<'host' | 'fight' | 'join' | 'client-fight' | null> = writable(null)
