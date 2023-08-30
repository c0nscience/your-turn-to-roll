import { writable } from 'svelte/store'
import type { Character } from './character.type'

export const characterStore = writable(new Array<Character>())

const numberOfCharactersToSend: number = 3

export const byInitiativeDesc = (a: Character, b: Character) => b.initiative - a.initiative

export const charactersToSend = (characters: Character[], from: number = 0) => {
  const sorted = [...characters]
      .sort(byInitiativeDesc)

  const selChar = new Array<Character>()
  let i = from
  while (selChar.length < numberOfCharactersToSend) {
    const c = sorted[i]

    if (i < sorted.length - 1) {
      i = i + 1
    } else {
      i = 0
    }

    if (c.hidden) {
      continue
    }

    selChar.push(c)
  }

  return selChar
    .map(c => c.name)
}
