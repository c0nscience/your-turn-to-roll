import { act, render, screen } from '@testing-library/svelte'
import CharacterForm from './CharacterForm.svelte'
import userEvent from '@testing-library/user-event'
import { characterStore } from './characterStore'
import { get } from 'svelte/store'

describe('CharacterForm', () => {
  const user = userEvent.setup()

  beforeEach(() => {
    characterStore.set([])
  })

  test('render', () => {
    render(CharacterForm)

    const name = screen.getByLabelText('Name')
    expect(name).toBeTruthy()
    expect(name).toBeEmptyDOMElement()

    const initiative = screen.getByLabelText('Initiative')
    expect(initiative).toBeTruthy()
    expect(initiative).toBeEmptyDOMElement()

    const addBtn = screen.getByRole('button')
    expect(addBtn).toHaveTextContent('Add')
  })

  describe('add button', () => {
    test('disabled if both fields are empty', () => {
      render(CharacterForm)
      const addBtn = screen.getByRole('button')
      expect(addBtn).toBeDisabled()
    })

    test('disabled if initiative field is empty', async () => {
      render(CharacterForm)
      const addBtn = screen.getByRole('button')
      const name = screen.getByLabelText('Name')

      await act(() => user.type(name, 'char #1'))

      expect(addBtn).toBeDisabled()
    })

    test('disabled if name field is empty', async () => {
      render(CharacterForm)
      const addBtn = screen.getByRole('button')
      const initiative = screen.getByLabelText('Initiative')

      await act(() => user.type(initiative, '2'))

      expect(addBtn).toBeDisabled()
    })

    test('enabled if both fields are filled', async () => {
      render(CharacterForm)
      const addBtn = screen.getByRole('button')

      const name = screen.getByLabelText('Name')
      const initiative = screen.getByLabelText('Initiative')
      await act(() => user.type(name, 'char #1'))
      await act(() => user.type(initiative, '2'))

      expect(addBtn).toBeEnabled()
    })

    test('adds new character to store after click', async () => {
      render(CharacterForm)
      const addBtn = screen.getByRole('button')
      const name = screen.getByLabelText('Name')
      const initiative = screen.getByLabelText('Initiative')
      await act(() => user.type(name, 'char #1'))
      await act(() => user.type(initiative, '2'))

      expect(get(characterStore)).toHaveLength(0)

      await act(() => user.click(addBtn))

      expect(get(characterStore)).toHaveLength(1)
    })

    test('clears the form afterwards', async () => {
      render(CharacterForm)
      const addBtn = screen.getByRole('button')
      const name = screen.getByLabelText('Name')
      const initiative = screen.getByLabelText('Initiative')
      await act(() => user.type(name, 'char #1'))
      await act(() => user.type(initiative, '2'))

      expect(get(characterStore)).toHaveLength(0)

      await act(() => user.click(addBtn))

      expect(get(characterStore)).toHaveLength(1)

      expect(name).toHaveValue('')
      expect(initiative).toHaveValue(null)
    })
  })

  describe('initiative field', () => {
    test('only accepts numbers', async () => {
      render(CharacterForm)
      const initiative = screen.getByLabelText('Initiative')
      expect(initiative).toHaveAttribute('type', 'number')
    })
  })
})
