import type { Ticket } from '$lib/types';

export const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 
  (typeof window !== 'undefined' && window.location.hostname !== 'localhost' && window.location.hostname !== '127.0.0.1'
    ? 'https://bugsinguish-api.onrender.com'
    : 'http://localhost:8080');

export async function fetchTickets(): Promise<Ticket[]> {
  const res = await fetch(`${API_BASE_URL}/tickets`);
  if (!res.ok) {
    throw new Error(`HTTP ${res.status}: ${res.statusText}`);
  }
  return await res.json();
}

export async function createTicket(payload: {
  title: string;
  description: string;
  stack_trace: string;
  repo_branch_url: string;
  severity: string;
}): Promise<Ticket> {
  const res = await fetch(`${API_BASE_URL}/tickets`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  });
  if (!res.ok) {
    const errText = await res.text();
    throw new Error(`Failed to create ticket: ${errText}`);
  }
  return await res.json();
}

export async function updateTicketStatus(id: string, status: string): Promise<void> {
  const res = await fetch(`${API_BASE_URL}/tickets/${id}`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ status })
  });
  if (!res.ok) {
    throw new Error(`Failed to update status: ${res.statusText}`);
  }
}
