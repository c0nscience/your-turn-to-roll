import { act, render, screen } from '@testing-library/svelte'
import CharacterSetupListEntry from './CharacterSetupListEntry.svelte'

describe('CharacterSetupListEntry', () => {
  test('render', () => {
    render(CharacterSetupListEntry, {name: 'char #2', initiative: 13})

    const name = screen.getByText('char #2')
    expect(name).toBeTruthy()

    const initiative = screen.getByText('13')
    expect(initiative).toBeTruthy()

  })
})
