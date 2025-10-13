import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import {
  transactionsApi,
  accountsApi,
  Transaction,
  BalanceResponse,
} from "@/services/api";
import { useUserStore } from "@/store/userStore";
import { FlowModal } from "@/components/FlowModal";
import { useSEO } from "@/hooks/useSEO";
import toast from "react-hot-toast";
import {
  ArrowUpCircle,
  ArrowDownCircle,
  RotateCcw,
  Wallet,
} from "lucide-react";
import { set } from "react-hook-form";

type TransactionType = "PURCHASE" | "DEPOSIT" | "REFUND";

export default function TransactionsPage() {
  const seoData = useSEO({
    title: "Transações - PayGateway",
    description:
      "Realize transações seguras, processamento de pagamentos, estornos e gerencie suas operações financeiras no PayGateway. Interface completa para operações bancárias.",
    keywords:
      "transações financeiras, processar pagamento, estorno, operações bancárias, transferência, pagamento online, transação segura",
    ogType: "webapp",
    schema: {
      "@context": "https://schema.org",
      "@type": "WebPage",
      name: "Transações - PayGateway",
      description:
        "Interface para processamento de transações e operações financeiras",
      url: "https://paygateway.exemplo.com/transactions",
      mainEntity: {
        "@type": "FinancialService",
        name: "Processamento de Transações",
        description:
          "Serviço para processamento seguro de transações financeiras",
        serviceType: "Payment Processing",
      },
    },
  });

  const [transactionType, setTransactionType] =
    useState<TransactionType>("PURCHASE");
  const [amountCents, setAmountCents] = useState("");
  const [cardToken, setCardToken] = useState("");
  const [refundTransactionId, setRefundTransactionId] = useState("");
  const [loading, setLoading] = useState(false);
  const [showFlowModal, setShowFlowModal] = useState(false);
  const [currentTransactionId, setCurrentTransactionId] = useState("");

  const { accountId, username, cards, balance, balanceStatus, setBalance } =
    useUserStore();

  useEffect(() => {
    if (accountId) {
      loadBalance();
    }
  }, [accountId]);

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

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (transactionType === "REFUND") {
      setAmountCents("10");
    }

    if (!accountId) {
      toast.error("Você precisa criar uma conta primeiro");
      return;
    }

    if (!amountCents || parseInt(amountCents) <= 0) {
      toast.error("Por favor, insira um valor válido");
      return;
    }

    if (transactionType === "PURCHASE" && !cardToken) {
      toast.error("Por favor, selecione um cartão");
      return;
    }

    if (transactionType === "REFUND" && !refundTransactionId) {
      toast.error("Por favor, insira o ID da transação original");
      return;
    }

    setLoading(true);

    try {
      const response = await transactionsApi.create({
        account_id: accountId,
        amount_cents: parseInt(amountCents),
        type: transactionType,
        card_token: transactionType === "PURCHASE" ? cardToken : undefined,
        refund_transaction_id:
          transactionType === "REFUND" ? refundTransactionId : undefined,
      });

      setCurrentTransactionId(response.data.id);
      setShowFlowModal(true);

      // Reset form
      setAmountCents("");
      setCardToken("");
      setRefundTransactionId("");
    } catch (error) {
      console.error("Erro ao criar transação:", error);
      toast.error(error.response?.data?.message || "Erro ao criar transação");
    } finally {
      setLoading(false);
    }
  };

  const handleFlowComplete = (transaction: Transaction) => {
    if (transaction.status === "APPROVED") {
      toast.success("Transação aprovada com sucesso!");
      loadBalance();
    } else if (transaction.status === "REJECTED") {
      toast.error("Transação rejeitada");
    }
  };

  if (!accountId) {
    return (
      <div className="container mx-auto flex min-h-[calc(100vh-4rem)] items-center justify-center px-4">
        <Card className="w-full max-w-md shadow-elegant">
          <CardHeader className="text-center">
            <CardTitle>Nenhuma conta encontrada</CardTitle>
            <CardDescription>
              Você precisa criar uma conta antes de realizar transações
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
      {seoData}
      <div className="mx-auto max-w-4xl">
        <div className="mb-8">
          <h1 className="mb-2 text-3xl font-bold">Realizar Transação</h1>
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
                    Calculando...
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

        {/* Transaction Form */}
        <Card className="shadow-elegant">
          <CardHeader>
            <CardTitle>Nova Transação</CardTitle>
            <CardDescription>
              Escolha o tipo de transação e preencha os dados
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Tabs
              value={transactionType}
              onValueChange={(v) => setTransactionType(v as TransactionType)}
            >
              <TabsList className="grid w-full grid-cols-3 mb-6">
                <TabsTrigger value="PURCHASE" className="gap-2">
                  <ArrowDownCircle className="h-4 w-4" />
                  Compra
                </TabsTrigger>
                <TabsTrigger value="DEPOSIT" className="gap-2">
                  <ArrowUpCircle className="h-4 w-4" />
                  Depósito
                </TabsTrigger>
                <TabsTrigger value="REFUND" className="gap-2">
                  <RotateCcw className="h-4 w-4" />
                  Reembolso
                </TabsTrigger>
              </TabsList>

              <form onSubmit={handleSubmit} className="space-y-6">
                <TabsContent value="PURCHASE" className="space-y-4 mt-0">
                  <div className="space-y-2">
                    <Label htmlFor="amount-purchase">Valor (em centavos)</Label>
                    <Input
                      id="amount-purchase"
                      type="number"
                      placeholder="Ex: 5000 (R$ 50,00)"
                      value={amountCents}
                      onChange={(e) => setAmountCents(e.target.value)}
                      disabled={loading}
                    />
                    {amountCents && (
                      <p className="text-sm text-muted-foreground">
                        = R$ {(parseInt(amountCents) / 100).toFixed(2)}
                      </p>
                    )}
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="card-token">Cartão</Label>
                    <Select
                      value={cardToken}
                      onValueChange={setCardToken}
                      disabled={loading}
                    >
                      <SelectTrigger id="card-token">
                        <SelectValue placeholder="Selecione um cartão" />
                      </SelectTrigger>
                      <SelectContent>
                        {cards.map((card) => (
                          <SelectItem key={card.id} value={card.card_token}>
                            •••• {card.last_four_digits} - {card.card_token}
                          </SelectItem>
                        ))}
                      </SelectContent>
                    </Select>
                  </div>
                </TabsContent>

                <TabsContent value="DEPOSIT" className="space-y-4 mt-0">
                  <div className="space-y-2">
                    <Label htmlFor="amount-deposit">Valor (em centavos)</Label>
                    <Input
                      id="amount-deposit"
                      type="number"
                      placeholder="Ex: 10000 (R$ 100,00)"
                      value={amountCents}
                      onChange={(e) => setAmountCents(e.target.value)}
                      disabled={loading}
                    />
                    {amountCents && (
                      <p className="text-sm text-muted-foreground">
                        = R$ {(parseInt(amountCents) / 100).toFixed(2)}
                      </p>
                    )}
                  </div>
                </TabsContent>

                <TabsContent value="REFUND" className="space-y-4 mt-0">
                  <div className="space-y-2">
                    <Label htmlFor="refund-id">ID da Transação Original</Label>
                    <Input
                      id="refund-id"
                      type="text"
                      placeholder="Cole o ID da transação"
                      value={refundTransactionId}
                      onChange={(e) => setRefundTransactionId(e.target.value)}
                      disabled={loading}
                    />
                  </div>
                </TabsContent>

                <Button
                  type="submit"
                  className="w-full h-12"
                  disabled={loading}
                >
                  {loading ? "Processando..." : "Realizar Transação"}
                </Button>
              </form>
            </Tabs>
          </CardContent>
        </Card>
      </div>

      <FlowModal
        open={showFlowModal}
        onOpenChange={setShowFlowModal}
        transactionId={currentTransactionId}
        onComplete={handleFlowComplete}
      />
    </div>
  );
}
