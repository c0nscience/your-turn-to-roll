import { writable } from 'svelte/store'
import type { Character } from './character.type'

export const characterStore = writable(new Array<Character>())

const numberOfCharactersToSend: number = 3

export const charactersToSend = (c: Character[], from: number = 0) => {
  const sorted = [...c]
    .sort((a, b) => b.initiative - a.initiative)

  const selectedCharacters = sorted
    .slice(from, from + numberOfCharactersToSend)

  let overflow: Character[] = []
  if ((from + numberOfCharactersToSend) > sorted.length) {
    const remaining = (from + numberOfCharactersToSend) - sorted.length
    overflow = sorted.slice(0, remaining)
  }

  return [...selectedCharacters, ...overflow]
    .map(c => c.name)
}
