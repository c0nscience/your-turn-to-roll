import { act, render, screen, within } from '@testing-library/svelte'
import ConfirmationDialog from './ConfirmationDialog.svelte'
import { expect, vi } from 'vitest'
import userEvent from '@testing-library/user-event'

describe('ConfirmationDialog', () => {
    const user = userEvent.setup()
    beforeAll(() => {
        HTMLDialogElement.prototype.show = vi.fn();
        HTMLDialogElement.prototype.showModal = vi.fn();
        HTMLDialogElement.prototype.close = vi.fn();
    });
    test('render', () => {
        render(ConfirmationDialog, {id: 'confirm-dialog', msg: 'dialog text'})

        const modal = screen.getByTestId('confirm-dialog')
        const text = within(modal).getByTestId('modal-text')
        expect(text).toBeTruthy()
    })

    test('confirming the dialog calls component event handler', async () => {
        const {component} = render(ConfirmationDialog, {id: 'confirm-dialog', msg: 'dialog text'})
        const confirmFn = vi.fn()
        component.$on('confirm', confirmFn)

        const modal = screen.getByTestId('confirm-dialog')
        const btn = within(modal).getByTestId('confirm-btn')

        await act(() => user.click(btn))

        expect(confirmFn).toHaveBeenCalledOnce()
    })

    test('closing the dialog does not trigger the confirm event', async () => {
        const {component} = render(ConfirmationDialog, {id: 'confirm-dialog', msg: 'dialog text'})
        const confirmFn = vi.fn()
        component.$on('confirm', confirmFn)

        const modal = screen.getByTestId('confirm-dialog')
        const btn = within(modal).getByTestId('cancel-btn')

        await act(() => user.click(btn))

        expect(confirmFn).not.toHaveBeenCalled()
    })
})
