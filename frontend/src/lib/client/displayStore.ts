import type { Writable } from 'svelte/store'
import { writable } from 'svelte/store'
import { apiWsUrl } from '../api/config'

export const namesStore: Writable<string[]> = writable([])

let socket: WebSocket

interface Message {
  type: 'update'
  payload: string[]
}

export const connect = (id: number) => {
  if (import.meta.env.PROD) {
    socket = new WebSocket(`ws://${window.location.host}${apiWsUrl}/fight/${id}/ws`)
  }else {
    socket = new WebSocket(`${apiWsUrl}/fight/${id}/ws`)
  }
  const promise = new Promise<void>((resolve, reject) => {
    socket.onerror = ev => {
      reject()
    }

    socket.onopen = ev => {
      resolve()
    }
  })
  socket.onmessage = evt => {
    const msg = JSON.parse(evt.data) as Message
    switch (msg.type) {
      case 'update':
        if (msg.payload) {
          namesStore.set(msg.payload)
        }
    }
  }
  socket.onerror = ev => {
    console.log('socket error', ev)
  }

  return promise
}
