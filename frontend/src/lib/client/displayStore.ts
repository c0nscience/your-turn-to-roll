import type { Writable } from 'svelte/store'
import { writable } from 'svelte/store'

export const namesStore: Writable<string[]> = writable([])

let socket: WebSocket

interface Message {
    type: 'update'
    payload: string[]
}

export const connect = (id: number) => {
    socket = new WebSocket(`ws://localhost:8081/api/fight/${id}/ws`)
    socket.onmessage = evt => {
        console.log('message received', evt.data)
        const msg = JSON.parse(evt.data) as Message
        console.log('message', typeof msg)
        console.log('message', msg.type)
        switch (msg.type) {
            case 'update':
                console.log('update message: ', msg)
                if (msg.payload) {
                    namesStore.set(msg.payload)
                }
        }
    }
}
