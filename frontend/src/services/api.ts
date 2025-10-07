import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
    "Accept-Language": "pt-BR",
  },
});

export interface CreateAccountRequest {
  username: string;
}

export interface CreateAccountResponse {
  id: string;
  username: string;
  created_at: string;
  updated_at: string;
}

export interface CreateCardRequest {
  account_id: string;
}

export interface CreateCardResponse {
  id: string;
  account_id: string;
  card_token: string;
  last_four_digits: string;
  created_at: string;
}

export interface CreateTransactionRequest {
  account_id: string;
  amount_cents: number;
  type: "PURCHASE" | "DEPOSIT" | "REFUND";
  card_token?: string;
  refund_transaction_id?: string;
}

export interface Transaction {
  id: string;
  account_id: string;
  amount_cents: number;
  type: string;
  status: "PENDING" | "APPROVED" | "REJECTED" | "ERROR";
  card_id?: string;
  refund_transaction_id?: string;
  created_at: string;
}

export interface BalanceCalculated {
  account_id: string;
  balance_cents: number;
}

export interface BalanceProcessing {
  message: string;
  status: string;
}

export type BalanceResponse = BalanceCalculated | BalanceProcessing;

export const accountsApi = {
  create: (data: CreateAccountRequest) =>
    api.post<CreateAccountResponse>("/accounts", data),
  getBalance: (accountId: string) =>
    api.get<BalanceResponse>(`/accounts/${accountId}/balance`),
};

export const cardsApi = {
  create: (data: CreateCardRequest) =>
    api.post<CreateCardResponse>("/cards", data),
  list: (accountId: string) =>
    api.get<CreateCardResponse[]>(`/cards?account_id=${accountId}`),
};

export const transactionsApi = {
  create: (data: CreateTransactionRequest) =>
    api.post<Transaction>("/transactions", data),
  list: (accountId: string) =>
    api.get<Transaction[]>(`/transactions/${accountId}`),
  get: (transactionId: string) =>
    api.get<Transaction>(`/transactions/${transactionId}`),
};
