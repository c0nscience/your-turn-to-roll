import { writable } from 'svelte/store'
import type { Character } from './character.type'

export const characterStore = writable(new Array<Character>())
