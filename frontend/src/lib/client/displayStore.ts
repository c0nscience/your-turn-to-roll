import type { Writable } from 'svelte/store'
import { writable } from 'svelte/store'
import { apiUrl, apiWsUrl } from '../api/config'

export const namesStore: Writable<string[]> = writable([])

let socket: WebSocket

interface Message {
  type: 'update'
  payload: string[]
}

export const connect = (id: number) => {
  socket = new WebSocket(`${apiWsUrl}/fight/${id}/ws`)
  const promise = new Promise<void>((resolve, reject) => {
    socket.onerror = ev => {
      reject()
    }

    socket.onopen = ev => {
      resolve()
    }
  })
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

  return promise
}
