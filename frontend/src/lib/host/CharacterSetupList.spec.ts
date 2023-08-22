import { act, render, screen } from '@testing-library/svelte'
import CharacterSetupList from './CharacterSetupList.svelte'
import { characterStore } from './characterStore'
import userEvent from '@testing-library/user-event'

describe('CharacterSetupList', () => {
    const user = userEvent.setup()

    test('render', async () => {
        render(CharacterSetupList)

        await act(() => characterStore.set([
            {id: 'b2f28af8-3458-4b7d-a8bd-9644cc47fd43', name: 'char #2', initiative: 13},
            {id: '005dfc25-e8c3-41a0-875b-71f3940d8607', name: 'char #1', initiative: 14},
            {id: '5f21bdc4-363f-441f-a080-62784eb831f3', name: 'char #3', initiative: 9},
        ]))

        const items = screen.getAllByRole('listitem')
        const names = items.map(i => i.textContent)
        expect(names).toEqual([
            'char #1',
            'char #2',
            'char #3',
        ])

        const deleteButtons = screen.getAllByRole('button')

        deleteButtons.forEach((b) => {
            expect(b).toBeTruthy()
        })
    })

    describe('delete button', () => {
        test('should remove entry from list', async () => {
            render(CharacterSetupList)

            await act(() => characterStore.set([
                {id: 'b2f28af8-3458-4b7d-a8bd-9644cc47fd43', name: 'char #2', initiative: 13},
                {id: '005dfc25-e8c3-41a0-875b-71f3940d8607', name: 'char #1', initiative: 14},
                {id: '5f21bdc4-363f-441f-a080-62784eb831f3', name: 'char #3', initiative: 9},
            ]))

            const deleteButtons = screen.getAllByRole('button')

            await act(() => user.click(deleteButtons[1]))

            const items = screen.getAllByRole('listitem')
            const names = items.map(i => i.textContent)
            expect(names).toEqual([
                'char #1',
                'char #3',
            ])
        })
    })
})
