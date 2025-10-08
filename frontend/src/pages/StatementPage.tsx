import { useState, useEffect } from "react";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import { transactionsApi, accountsApi, Transaction } from "@/services/api";
import { useUserStore } from "@/store/userStore";
import { Receipt, CreditCard, Copy, Check, Wallet } from "lucide-react";
import toast from "react-hot-toast";

export default function StatementPage() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [transactionsFilteredByCard, setTransactionsFilteredByCard] = useState<
    Transaction[]
  >([]);

  const [loading, setLoading] = useState(true);
  const [selectedCard, setSelectedCard] = useState<string>("");
  const [copiedId, setCopiedId] = useState<string | null>(null);

  const { accountId, username, cards, balance, balanceStatus, setBalance } =
    useUserStore();

  useEffect(() => {
    if (accountId) {
      loadTransactions();
      loadBalance();
    }
  }, [accountId]);

  useEffect(() => {
    loadTransactionsFilteredByCard(selectedCard);
    console.log("Selected Card:", selectedCard);
    console.log(transactionsFilteredByCard);
  }, [selectedCard]);

  const loadTransactionsFilteredByCard = async (cardId: string) => {
    if (!accountId) return;

    if (cardId) {
      const response = await transactionsApi.getByCardId(cardId);
      setTransactionsFilteredByCard(response.data);
    }
  };

  const loadTransactions = async () => {
    if (!accountId) return;

    setLoading(true);
    try {
      const response = await transactionsApi.list(accountId);
      setTransactions(response.data);
    } catch (error) {
      console.error("Erro ao carregar transações:", error);
      toast.error("Erro ao carregar transações");
    } finally {
      setLoading(false);
    }
  };

  const loadBalance = async () => {
    if (!accountId) return;

    try {
      const response = await accountsApi.getBalance(accountId);
      const data = response.data;

      if ("account_id" in data) {
        setBalance(data.balance_cents, "CALCULATED");
      }
      if ("message" in data) {
        toast.loading(data.message);

        setTimeout(() => {
          toast.dismiss();
        }, 2000);

        setTimeout(() => {
          loadBalance();
        }, 1000);
      }
    } catch (error) {
      console.error("Erro ao carregar saldo:", error);
      toast.error("Erro ao carregar saldo");
    }
  };

  const copyToClipboard = async (id: string) => {
    try {
      await navigator.clipboard.writeText(id);
      setCopiedId(id);
      toast.success("ID copiado!");
      setTimeout(() => setCopiedId(null), 2000);
    } catch (error) {
      toast.error("Erro ao copiar ID");
    }
  };

  const getStatusBadge = (status: string) => {
    switch (status) {
      case "APPROVED":
        return (
          <Badge className="bg-success text-success-foreground">Aprovada</Badge>
        );
      case "REJECTED":
        return <Badge variant="destructive">Rejeitada</Badge>;
      case "PENDING":
        return <Badge variant="secondary">Pendente</Badge>;
      default:
        return <Badge variant="outline">{status}</Badge>;
    }
  };

  const getTypeBadge = (type: string) => {
    switch (type) {
      case "PURCHASE":
        return (
          <Badge
            variant="outline"
            className="bg-destructive/10 text-destructive border-destructive/20"
          >
            Compra
          </Badge>
        );
      case "DEPOSIT":
        return (
          <Badge
            variant="outline"
            className="bg-success/10 text-success border-success/20"
          >
            Depósito
          </Badge>
        );
      case "REFUND":
        return (
          <Badge
            variant="outline"
            className="bg-warning/10 text-warning border-warning/20"
          >
            Reembolso
          </Badge>
        );
      default:
        return <Badge variant="outline">{type}</Badge>;
    }
  };

  if (!accountId) {
    return (
      <div className="container mx-auto flex min-h-[calc(100vh-4rem)] items-center justify-center px-4">
        <Card className="w-full max-w-md shadow-elegant">
          <CardHeader className="text-center">
            <CardTitle>Nenhuma conta encontrada</CardTitle>
            <CardDescription>
              Você precisa criar uma conta antes de visualizar o extrato
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
      <div className="mx-auto max-w-6xl">
        <div className="mb-8">
          <h1 className="mb-2 text-3xl font-bold">Extrato</h1>
          <p className="text-muted-foreground">
            Bem-vindo,{" "}
            <span className="font-semibold text-foreground">{username}</span>
          </p>
        </div>

        {/* Balance Card */}
        <Card className="mb-8 shadow-card">
          <CardContent className="flex items-center justify-between py-6">
            <div className="flex items-center gap-4">
              <div className="rounded-lg bg-gradient-success p-3">
                <Wallet className="h-6 w-6 text-success-foreground" />
              </div>
              <div>
                <p className="text-sm text-muted-foreground">Saldo atual</p>
                {balanceStatus === "PROCESSING" ? (
                  <p className="text-2xl font-bold text-warning">
                    Calculando saldo...
                  </p>
                ) : (
                  <p className="text-2xl font-bold">
                    R$ {balance !== null ? (balance / 100).toFixed(2) : "0,00"}
                  </p>
                )}
              </div>
            </div>
            <Button variant="outline" size="sm" onClick={loadBalance}>
              Atualizar
            </Button>
          </CardContent>
        </Card>

        {/* Statement Tabs */}
        <Card className="shadow-elegant">
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Receipt className="h-5 w-5" />
              Histórico de Transações
            </CardTitle>
          </CardHeader>
          <CardContent>
            <Tabs defaultValue="account" className="w-full">
              <TabsList className="grid w-full grid-cols-2 mb-6">
                <TabsTrigger value="account">Extrato da Conta</TabsTrigger>
                <TabsTrigger value="card">Extrato do Cartão</TabsTrigger>
              </TabsList>

              <TabsContent value="account" className="space-y-4">
                {loading ? (
                  <div className="py-12 text-center text-muted-foreground">
                    Carregando transações...
                  </div>
                ) : transactions.length === 0 ? (
                  <div className="py-12 text-center">
                    <Receipt className="mx-auto mb-4 h-16 w-16 text-muted-foreground" />
                    <p className="text-muted-foreground">
                      Nenhuma transação encontrada
                    </p>
                  </div>
                ) : (
                  <div className="rounded-lg border">
                    <Table>
                      <TableHeader>
                        <TableRow>
                          <TableHead>ID</TableHead>
                          <TableHead>Tipo</TableHead>
                          <TableHead>Valor</TableHead>
                          <TableHead>Status</TableHead>
                          <TableHead>Data</TableHead>
                          <TableHead></TableHead>
                        </TableRow>
                      </TableHeader>
                      <TableBody>
                        {transactions.map((transaction) => (
                          <TableRow key={transaction.id}>
                            <TableCell className="font-mono text-xs">
                              {transaction.id.substring(0, 8)}...
                            </TableCell>
                            <TableCell>
                              {getTypeBadge(transaction.type)}
                            </TableCell>
                            <TableCell
                              className={`font-semibold ${
                                transaction.type === "DEPOSIT" ||
                                transaction.type === "REFUND"
                                  ? "text-success"
                                  : "text-destructive"
                              }`}
                            >
                              {transaction.type === "DEPOSIT" ||
                              transaction.type === "REFUND"
                                ? "+"
                                : "-"}
                              R$ {(transaction.amount_cents / 100).toFixed(2)}
                            </TableCell>
                            <TableCell>
                              {getStatusBadge(transaction.status)}
                            </TableCell>
                            <TableCell className="text-sm text-muted-foreground">
                              {new Date(transaction.created_at).toLocaleString(
                                "pt-BR",
                              )}
                            </TableCell>
                            <TableCell>
                              <Button
                                variant="ghost"
                                size="sm"
                                onClick={() => copyToClipboard(transaction.id)}
                              >
                                {copiedId === transaction.id ? (
                                  <Check className="h-4 w-4" />
                                ) : (
                                  <Copy className="h-4 w-4" />
                                )}
                              </Button>
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </div>
                )}
              </TabsContent>

              <TabsContent value="card" className="space-y-4">
                <div className="mb-4">
                  <Select value={selectedCard} onValueChange={setSelectedCard}>
                    <SelectTrigger>
                      <SelectValue placeholder="Selecione um cartão" />
                    </SelectTrigger>
                    <SelectContent>
                      {cards.map((card) => (
                        <SelectItem key={card.id} value={card.id}>
                          <div className="flex items-center gap-2">
                            <CreditCard className="h-4 w-4" />
                            •••• {card.last_four_digits}
                          </div>
                        </SelectItem>
                      ))}
                    </SelectContent>
                  </Select>
                </div>

                {!selectedCard ? (
                  <div className="py-12 text-center">
                    <CreditCard className="mx-auto mb-4 h-16 w-16 text-muted-foreground" />
                    <p className="text-muted-foreground">
                      Selecione um cartão para visualizar o extrato
                    </p>
                  </div>
                ) : loading ? (
                  <div className="py-12 text-center text-muted-foreground">
                    Carregando transações...
                  </div>
                ) : transactionsFilteredByCard.length === 0 ? (
                  <div className="py-12 text-center">
                    <Receipt className="mx-auto mb-4 h-16 w-16 text-muted-foreground" />
                    <p className="text-muted-foreground">
                      Nenhuma transação encontrada para este cartão
                    </p>
                  </div>
                ) : (
                  <div className="rounded-lg border">
                    <Table>
                      <TableHeader>
                        <TableRow>
                          <TableHead>ID</TableHead>
                          <TableHead>Tipo</TableHead>
                          <TableHead>Valor</TableHead>
                          <TableHead>Status</TableHead>
                          <TableHead>Data</TableHead>
                          <TableHead></TableHead>
                        </TableRow>
                      </TableHeader>
                      <TableBody>
                        {transactionsFilteredByCard.map((transaction) => (
                          <TableRow key={transaction.id}>
                            <TableCell className="font-mono text-xs">
                              {transaction.id.substring(0, 8)}...
                            </TableCell>
                            <TableCell>
                              {getTypeBadge(transaction.type)}
                            </TableCell>
                            <TableCell
                              className={`font-semibold ${
                                transaction.type === "DEPOSIT" ||
                                transaction.type === "REFUND"
                                  ? "text-success"
                                  : "text-destructive"
                              }`}
                            >
                              {transaction.type === "DEPOSIT" ||
                              transaction.type === "REFUND"
                                ? "+"
                                : "-"}
                              R$ {(transaction.amount_cents / 100).toFixed(2)}
                            </TableCell>
                            <TableCell>
                              {getStatusBadge(transaction.status)}
                            </TableCell>
                            <TableCell className="text-sm text-muted-foreground">
                              {new Date(transaction.created_at).toLocaleString(
                                "pt-BR",
                              )}
                            </TableCell>
                            <TableCell>
                              <Button
                                variant="ghost"
                                size="sm"
                                onClick={() => copyToClipboard(transaction.id)}
                              >
                                {copiedId === transaction.id ? (
                                  <Check className="h-4 w-4" />
                                ) : (
                                  <Copy className="h-4 w-4" />
                                )}
                              </Button>
                            </TableCell>
                          </TableRow>
                        ))}
                      </TableBody>
                    </Table>
                  </div>
                )}
              </TabsContent>
            </Tabs>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
