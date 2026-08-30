export type TicketStatus = 'new' | 'triaging' | 'sandbox_running' | 'diagnosed' | 'resolved';
export type TicketSeverity = 'low' | 'medium' | 'high' | 'critical';

export interface Diagnosis {
  root_cause: string;
  explanation: string;
  file: string;
  diff: string;
}

export interface Ticket {
  id: string;
  title: string;
  description: string;
  stack_trace: string;
  repo_branch_url: string;
  severity: TicketSeverity;
  status: TicketStatus;
  diagnosis?: Diagnosis;
  logs?: string[];
  created_at: string;
  updated_at: string;
}

export interface SseLogEvent {
  step: string;
  message: string;
  timestamp: string;
  type: 'info' | 'warn' | 'error' | 'success';
}
