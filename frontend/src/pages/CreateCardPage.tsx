import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { cardsApi } from "@/services/api";
import { useUserStore } from "@/store/userStore";
import toast from "react-hot-toast";
import { CreditCard, Plus, Copy, Check } from "lucide-react";

export default function CreateCardPage() {
  const [loading, setLoading] = useState(false);
  const [copiedToken, setCopiedToken] = useState<string | null>(null);
  const { accountId, username, cards, setCards, addCard } = useUserStore();

  useEffect(() => {
    if (accountId) {
      loadCards();
    }
  }, [accountId]);

  const loadCards = async () => {
    if (!accountId) return;

    try {
      const response = await cardsApi.list(accountId);
      setCards(response.data);
    } catch (error) {
      console.error("Erro ao carregar cartões:", error);
    }
  };

  const handleCreateCard = async () => {
    if (!accountId) {
      toast.error("Você precisa criar uma conta primeiro");
      return;
    }

    setLoading(true);

    try {
      const response = await cardsApi.create({ account_id: accountId });
      addCard(response.data);
      toast.success("Cartão criado com sucesso!");
    } catch (error) {
      console.error("Erro ao criar cartão:", error);
      toast.error("Erro ao criar cartão. Tente novamente.");
    } finally {
      setLoading(false);
    }
  };

  const copyToClipboard = async (token: string) => {
    try {
      await navigator.clipboard.writeText(token);
      setCopiedToken(token);
      toast.success("Token copiado!");
      setTimeout(() => setCopiedToken(null), 2000);
    } catch (error) {
      toast.error("Erro ao copiar token");
    }
  };

  if (!accountId) {
    return (
      <div className="container mx-auto flex min-h-[calc(100vh-4rem)] items-center justify-center px-4">
        <Card className="w-full max-w-md shadow-elegant">
          <CardHeader className="text-center">
            <CardTitle>Nenhuma conta encontrada</CardTitle>
            <CardDescription>
              Você precisa criar uma conta antes de criar um cartão
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Button asChild className="w-full">
              <a href="/create-account">Criar Conta</a>
            </Button>
          </CardContent>
        </Card>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="mx-auto max-w-4xl">
        <div className="mb-8">
          <h1 className="mb-2 text-3xl font-bold">Meus Cartões</h1>
          <p className="text-muted-foreground">
            Bem-vindo,{" "}
            <span className="font-semibold text-foreground">{username}</span>
          </p>
        </div>

        <Card className="mb-8 shadow-card">
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Plus className="h-5 w-5" />
              Criar Novo Cartão
            </CardTitle>
            <CardDescription>
              Gere um cartão de crédito virtual para realizar transações
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Button
              onClick={handleCreateCard}
              disabled={loading}
              size="lg"
              className="w-full"
            >
              <CreditCard className="mr-2 h-5 w-5" />
              {loading ? "Gerando cartão..." : "Gerar Novo Cartão de Crédito"}
            </Button>
          </CardContent>
        </Card>

        <div>
          <h2 className="mb-4 text-xl font-semibold">Cartões existentes</h2>
          {cards.length === 0 ? (
            <Card className="shadow-card">
              <CardContent className="flex flex-col items-center justify-center py-12">
                <CreditCard className="mb-4 h-16 w-16 text-muted-foreground" />
                <p className="text-muted-foreground">
                  Você ainda não possui cartões. Crie seu primeiro cartão acima!
                </p>
              </CardContent>
            </Card>
          ) : (
            <div className="grid gap-4 md:grid-cols-2">
              {cards.map((card) => (
                <Card key={card.id} className="shadow-card">
                  <CardContent className="pt-6">
                    <div className="mb-4 flex items-start justify-between">
                      <div>
                        <p className="mb-1 text-sm text-muted-foreground">
                          Últimos 4 dígitos
                        </p>
                        <p className="text-2xl font-bold">
                          •••• {card.last_four_digits}
                        </p>
                      </div>
                      <div className="rounded-lg bg-gradient-primary p-3">
                        <CreditCard className="h-6 w-6 text-primary-foreground" />
                      </div>
                    </div>

                    <div className="space-y-2">
                      <p className="text-xs text-muted-foreground">
                        Token do cartão
                      </p>
                      <div className="flex items-center gap-2">
                        <code className="flex-1 rounded bg-secondary px-3 py-2 text-xs font-mono">
                          {card.card_token}
                        </code>
                        <Button
                          size="sm"
                          variant="outline"
                          onClick={() => copyToClipboard(card.card_token)}
                        >
                          {copiedToken === card.card_token ? (
                            <Check className="h-4 w-4" />
                          ) : (
                            <Copy className="h-4 w-4" />
                          )}
                        </Button>
                      </div>
                    </div>

                    <p className="mt-4 text-xs text-muted-foreground">
                      Criado em{" "}
                      {new Date(card.created_at).toLocaleDateString("pt-BR")}
                    </p>
                  </CardContent>
                </Card>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
