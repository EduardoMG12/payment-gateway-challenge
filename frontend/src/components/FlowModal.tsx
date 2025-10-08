import { useEffect, useState } from "react";
import { motion, AnimatePresence } from "framer-motion";
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Checkbox } from "@/components/ui/checkbox";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";
import {
  Loader2,
  CheckCircle2,
  XCircle,
  ArrowRight,
  Server,
  Database,
  MessageSquare,
  Cpu,
  Shield,
  Coins,
} from "lucide-react";
import { transactionsApi, Transaction } from "@/services/api";

interface FlowModalProps {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  transactionId: string;
  onComplete?: (transaction: Transaction) => void;
}

interface FlowStep {
  id: string;
  label: string;
  icon: React.ComponentType<React.SVGProps<SVGSVGElement>>;
  status: "pending" | "processing" | "completed" | "failed";
}

const beginFlowHardcoded: FlowStep[] = [
  {
    id: "frontend",
    label: "Frontend enviando requisição para API...",
    icon: ArrowRight,
    status: "processing",
  },
  {
    id: "api-receive",
    label: "API Go recebeu requisição, validando dados...",
    icon: Server,
    status: "pending",
  },
  {
    id: "api-publish",
    label: "API Go publicando mensagem na fila...",
    icon: MessageSquare,
    status: "pending",
  },
  {
    id: "rabbitmq",
    label: "RabbitMQ - Mensagem enfileirada...",
    icon: MessageSquare,
    status: "pending",
  },
  {
    id: "rust-consume",
    label: "Processador Rust consumiu mensagem...",
    icon: Cpu,
    status: "pending",
  },
  {
    id: "rust-business",
    label: "Aplicando regras de negócio (saldo, idempotência)...",
    icon: Shield,
    status: "pending",
  },
  {
    id: "postgres",
    label: "PostgreSQL persistindo transação...",
    icon: Database,
    status: "pending",
  },
  {
    id: "redis",
    label: "Redis atualizando cache de saldo...",
    icon: Coins,
    status: "pending",
  },
  {
    id: "complete",
    label: "Transação concluída!",
    icon: CheckCircle2,
    status: "pending",
  },
];

export function FlowModal({
  open,
  onOpenChange,
  transactionId,
  onComplete,
}: FlowModalProps) {
  const [steps, setSteps] = useState<FlowStep[]>(beginFlowHardcoded);
  const [currentTransaction, setCurrentTransaction] =
    useState<Transaction | null>(null);
  const [disableAnimation, setDisableAnimation] = useState(() => {
    return localStorage.getItem("disableFlowAnimation") === "true";
  });

  const resetFlow = () => {
    setSteps(beginFlowHardcoded);
    setCurrentTransaction(null);
  };

  useEffect(() => {
    if (open && transactionId) {
      resetFlow();
      simulateFlow();
    }
  }, [open, transactionId]);

  const simulateFlow = async () => {
    const delays = [500, 800, 600, 700, 900, 1200, 800, 600, 500];

    for (let i = 0; i < steps.length; i++) {
      await new Promise((resolve) =>
        setTimeout(resolve, disableAnimation ? 50 : delays[i]),
      );

      setSteps((prev) =>
        prev.map((step, index) => {
          if (index < i) return { ...step, status: "completed" };
          if (index === i) return { ...step, status: "processing" };
          return step;
        }),
      );

      // Check transaction status at key points
      if (i === 4 || i === 7) {
        try {
          const response =
            await transactionsApi.getByTransactionId(transactionId);
          setCurrentTransaction(response.data);

          if (response.data.status === "REJECTED") {
            setSteps((prev) =>
              prev.map((step, index) =>
                index <= i ? { ...step, status: "failed" } : step,
              ),
            );
            break;
          }
        } catch (error) {
          console.error("Error checking transaction:", error);
        }
      }
    }

    // Final status update
    setSteps((prev) => prev.map((step) => ({ ...step, status: "completed" })));

    // Final transaction fetch
    try {
      const response = await transactionsApi.getByTransactionId(transactionId);
      setCurrentTransaction(response.data);
      onComplete?.(response.data);
    } catch (error) {
      console.error("Error fetching final transaction:", error);
    }
  };

  const handleDisableAnimationChange = (checked: boolean) => {
    setDisableAnimation(checked);
    localStorage.setItem("disableFlowAnimation", String(checked));
  };

  const getStatusIcon = (status: FlowStep["status"]) => {
    switch (status) {
      case "completed":
        return <CheckCircle2 className="h-5 w-5 text-success" />;
      case "processing":
        return <Loader2 className="h-5 w-5 animate-spin text-primary" />;
      case "failed":
        return <XCircle className="h-5 w-5 text-destructive" />;
      default:
        return <div className="h-5 w-5 rounded-full border-2 border-muted" />;
    }
  };

  const handleOpenChange = (isOpen: boolean) => {
    if (!isOpen) {
      resetFlow();
    }
    onOpenChange(isOpen);
  };

  return (
    <Dialog open={open} onOpenChange={handleOpenChange}>
      <DialogContent className="max-w-2xl max-h-[80vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle className="text-2xl">Fluxo da Transação</DialogTitle>
        </DialogHeader>

        <div className="space-y-4 py-4">
          <AnimatePresence mode="wait">
            {steps.map((step, index) => (
              <motion.div
                key={step.id}
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ delay: disableAnimation ? 0 : index * 0.1 }}
                className={`flex items-start gap-4 rounded-lg border p-4 transition-colors ${
                  step.status === "processing"
                    ? "border-primary bg-primary/5"
                    : step.status === "completed"
                      ? "border-success/30 bg-success/5"
                      : step.status === "failed"
                        ? "border-destructive/30 bg-destructive/5"
                        : "border-border"
                }`}
              >
                <div className="flex-shrink-0 mt-0.5">
                  {getStatusIcon(step.status)}
                </div>
                <div className="flex-1">
                  <p
                    className={`font-medium ${
                      step.status === "processing"
                        ? "text-primary"
                        : step.status === "completed"
                          ? "text-success"
                          : step.status === "failed"
                            ? "text-destructive"
                            : "text-muted-foreground"
                    }`}
                  >
                    {step.label}
                  </p>
                </div>
                <step.icon
                  className={`h-5 w-5 ${
                    step.status === "completed"
                      ? "text-success"
                      : step.status === "processing"
                        ? "text-primary"
                        : step.status === "failed"
                          ? "text-destructive"
                          : "text-muted-foreground"
                  }`}
                />
              </motion.div>
            ))}
          </AnimatePresence>

          {currentTransaction && (
            <motion.div
              initial={{ opacity: 0, y: 10 }}
              animate={{ opacity: 1, y: 0 }}
              className={`mt-6 rounded-lg border p-4 ${
                currentTransaction.status === "APPROVED"
                  ? "border-success bg-success/5"
                  : currentTransaction.status === "REJECTED"
                    ? "border-destructive bg-destructive/5"
                    : "border-warning bg-warning/5"
              }`}
            >
              <h3 className="mb-2 font-semibold">Status Final</h3>
              <div className="space-y-1 text-sm">
                <p>
                  <span className="text-muted-foreground">ID:</span>{" "}
                  <span className="font-mono">{currentTransaction.id}</span>
                </p>
                <p>
                  <span className="text-muted-foreground">Status:</span>{" "}
                  <span
                    className={`font-semibold ${
                      currentTransaction.status === "APPROVED"
                        ? "text-success"
                        : currentTransaction.status === "REJECTED"
                          ? "text-destructive"
                          : "text-warning"
                    }`}
                  >
                    {currentTransaction.status}
                  </span>
                </p>
                <p>
                  <span className="text-muted-foreground">Valor:</span> R${" "}
                  {(currentTransaction.amount_cents / 100).toFixed(2)}
                </p>
              </div>
            </motion.div>
          )}
        </div>

        <div className="flex items-center justify-between border-t pt-4">
          <div className="flex items-center gap-2">
            <Checkbox
              id="disable-animation"
              checked={disableAnimation}
              onCheckedChange={handleDisableAnimationChange}
            />
            <Label
              htmlFor="disable-animation"
              className="text-sm cursor-pointer"
            >
              Desativar animação para futuras transações
            </Label>
          </div>
          <Button onClick={() => onOpenChange(false)}>Fechar</Button>
        </div>
      </DialogContent>
    </Dialog>
  );
}
