import axios from "axios";

const API_BASE_URL = "http://localhost:8080/api/v1";

export const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

export interface CreateAccountRequest {
  username: string;
}

export interface CreateAccountResponse {
  id: string;
  username: string;
  created_at: string;
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
  status: "PENDING" | "APPROVED" | "REJECTED";
  card_token?: string;
  refund_transaction_id?: string;
  created_at: string;
}

export interface BalanceResponse {
  account_id: string;
  balance_cents: number;
  status: "CALCULATED" | "PROCESSING";
}

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
    api.get<Transaction[]>(`/transactions/test/${accountId}`),
  get: (transactionId: string) =>
    api.get<Transaction>(`/transactions/${transactionId}`),
};
