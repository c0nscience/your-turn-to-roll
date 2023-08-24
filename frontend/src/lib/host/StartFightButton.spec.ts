import StartFightButton from './StartFightButton.svelte'
import { act, render } from '@testing-library/svelte'
import { expect } from 'vitest'
import userEvent from '@testing-library/user-event'

describe('StartFightButton', () => {
  const user = userEvent.setup()
  test('render', () => {
    const {getByRole} = render(StartFightButton)

    const component = getByRole('button')

    expect(component).toHaveTextContent('Start')
  })
})
