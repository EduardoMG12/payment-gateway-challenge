import { create } from "zustand";
import { persist } from "zustand/middleware";

interface Card {
  id: string;
  account_id: string;
  card_token: string;
  last_four_digits: string;
  created_at: string;
}

interface UserState {
  accountId: string | null;
  username: string | null;
  cards: Card[];
  balance: number | null;
  balanceStatus: "CALCULATED" | "PROCESSING" | null;
  setAccount: (accountId: string, username: string) => void;
  setCards: (cards: Card[]) => void;
  addCard: (card: Card) => void;
  setBalance: (balance: number, status: "CALCULATED" | "PROCESSING") => void;
  clearAccount: () => void;
}

export const useUserStore = create<UserState>()(
  persist(
    (set) => ({
      accountId: null,
      username: null,
      cards: [],
      balance: null,
      balanceStatus: null,
      setAccount: (accountId, username) => set({ accountId, username }),
      setCards: (cards) => set({ cards }),
      addCard: (card) => set((state) => ({ cards: [...state.cards, card] })),
      setBalance: (balance, status) => set({ balance, balanceStatus: status }),
      clearAccount: () =>
        set({
          accountId: null,
          username: null,
          cards: [],
          balance: null,
          balanceStatus: null,
        }),
    }),
    {
      name: "user-storage",
    },
  ),
);
