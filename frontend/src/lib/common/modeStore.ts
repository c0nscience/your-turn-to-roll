import type { Writable } from 'svelte/store'
import { writable } from 'svelte/store'

export const modeStore: Writable<'setup' | 'join-setup' | 'fight' | 'join' | 'client-fight' | null> = writable(null)
