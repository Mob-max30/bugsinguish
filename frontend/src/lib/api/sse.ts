import type { SseLogEvent } from '$lib/types';

export function createSseListener(
  url: string, 
  onMessage: (event: SseLogEvent) => void,
  onError?: (err: Event) => void
): () => void {
  if (typeof window === 'undefined') return () => {};

  try {
    const eventSource = new EventSource(url);

    eventSource.onmessage = (e) => {
      try {
        const parsed: SseLogEvent = JSON.parse(e.data);
        onMessage(parsed);
      } catch {
        onMessage({
          step: 'INFO',
          message: e.data,
          timestamp: new Date().toLocaleTimeString(),
          type: 'info'
        });
      }
    };

    if (onError) {
      eventSource.onerror = onError;
    }

    return () => {
      eventSource.close();
    };
  } catch (err) {
    console.warn('[SSE] EventSource not supported or connection error:', err);
    return () => {};
  }
}
