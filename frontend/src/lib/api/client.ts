import type { Character } from '../host/character.type';
import { charactersToSend } from '../host/characterStore';
import { apiUrl } from './config';

export const save = async (characters: Character[], sessionId: number) => {
  await fetch(`${apiUrl}/session/${sessionId}`, {
    method: 'POST',
    mode: 'cors',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      characters,
    }),
  });
};


export const sendUpdateToClients = async (characters: Character[], currId: number, sessionId: number) => {
  const payload = charactersToSend(characters, currId);

  await fetch(`${apiUrl}/fight/${sessionId}/update`, {
    method: 'POST',
    mode: 'cors',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ payload }),
  });

};
