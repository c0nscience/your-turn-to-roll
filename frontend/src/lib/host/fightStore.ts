import type { Writable } from 'svelte/store'
import { writable } from 'svelte/store'

export const currentIdxStore: Writable<number> = writable(0)
