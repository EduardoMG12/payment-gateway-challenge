import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import {
  ArrowRight,
  Database,
  Lock,
  Zap,
  Code2,
  Download,
  ExternalLink,
  FileText,
} from "lucide-react";
import { useSEO } from "@/hooks/useSEO";
import postmanCollection from "../../public/payment-gateway.postman_collection.json";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export default function HomePage() {
  const seoData = useSEO({
    title: "PayGateway - Gateway de Pagamento Fictício",
    description:
      "Sistema completo de gateway de pagamento desenvolvido com arquitetura moderna usando Rust, Go e React. Processamento seguro de transações financeiras com API RESTful completa.",
    keywords:
      "gateway de pagamento, processamento de transações, fintech, rust, go, react, sistema financeiro, api de pagamentos, transações seguras, sistema distribuído",
    ogType: "website",
    schema: {
      "@context": "https://schema.org",
      "@type": "WebApplication",
      name: "PayGateway",
      description:
        "Sistema completo de gateway de pagamento com arquitetura moderna",
      url: "https://paygateway.exemplo.com",
      applicationCategory: "FinanceApplication",
      operatingSystem: "Web",
      offers: {
        "@type": "Offer",
        price: "0",
        priceCurrency: "BRL",
      },
      creator: {
        "@type": "Person",
        name: "Eduardo MG",
        url: "https://github.com/EduardoMG12",
      },
      features: [
        "Processamento de transações",
        "Gerenciamento de contas",
        "Criação de cartões virtuais",
        "Extratos detalhados",
        "API RESTful",
        "Arquitetura distribuída",
      ],
    },
  });

  const downloadPostmanCollection = () => {
    const dataStr = JSON.stringify(postmanCollection, null, 2);
    const dataUri = `data:application/json;charset=utf-8,${encodeURIComponent(dataStr)}`;

    const exportFileDefaultName = "Payment_Gateway_API_Collection.json";

    const linkElement = document.createElement("a");
    linkElement.setAttribute("href", dataUri);
    linkElement.setAttribute("download", exportFileDefaultName);
    linkElement.click();
  };

  return (
    <div className="min-h-screen">
      {seoData}
      {/* Hero Section */}
      <section className="relative overflow-hidden bg-gradient-hero py-20 text-primary-foreground">
        <div className="container mx-auto px-4">
          <div className="mx-auto max-w-3xl text-center">
            <h1 className="mb-6 text-5xl font-bold leading-tight">
              Gateway de Pagamento Fictício
            </h1>
            <p className="mb-8 text-xl opacity-90">
              Um projeto de aprendizado em Rust e Go para explorar conceitos de
              arquitetura de sistemas distribuídos
            </p>
            <div className="flex flex-col gap-4 sm:flex-row sm:justify-center">
              <Button asChild size="lg" variant="secondary">
                <Link to="/create-account">
                  Comece agora <ArrowRight className="ml-2 h-5 w-5" />
                </Link>
              </Button>
              <Button
                asChild
                size="lg"
                variant="outline"
                className="border-primary-foreground/30 bg-primary-foreground/10 text-primary-foreground hover:bg-primary-foreground/20"
              >
                <a
                  href="https://github.com/EduardoMG12/payment-gateway-challenge"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  <Code2 className="mr-2 h-5 w-5" />
                  Ver no GitHub
                </a>
              </Button>
            </div>
          </div>
        </div>
      </section>

      {/* About Section */}
      <section className="py-16">
        <div className="container mx-auto px-4">
          <div className="mx-auto max-w-3xl">
            <h2 className="mb-6 text-3xl font-bold">Sobre o Projeto</h2>
            <div className="space-y-4 text-lg text-muted-foreground">
              <p>
                Este projeto foi criado com o objetivo de aprender e explorar
                tecnologias modernas de backend, especificamente{" "}
                <strong className="text-foreground">Rust</strong> e{" "}
                <strong className="text-foreground">Go</strong>.
              </p>
              <p>
                Ele simula um gateway de pagamento real, implementando conceitos
                fundamentais como filas de mensagens, processamento assíncrono,
                e persistência de dados. O projeto{" "}
                <strong className="text-foreground">não é comercial</strong> e
                serve apenas para fins educacionais.
              </p>
              <p>
                A arquitetura é organizada em um{" "}
                <strong className="text-foreground">monorepo</strong>, contendo
                o frontend (React + TypeScript), a API (Go), e o processador de
                transações (Rust).
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Architecture Section */}
      <section className="bg-secondary/30 py-16">
        <div className="container mx-auto px-4">
          <div className="mx-auto max-w-4xl">
            <h2 className="mb-8 text-center text-3xl font-bold">
              Arquitetura do Sistema
            </h2>
            <Card className="shadow-elegant">
              <CardContent className="p-8">
                <div className="grid gap-6 md:grid-cols-3">
                  <div className="flex flex-col items-center text-center">
                    <div className="mb-4 rounded-lg bg-gradient-primary p-4">
                      <Database className="h-8 w-8 text-primary-foreground" />
                    </div>
                    <h3 className="mb-2 font-semibold">Frontend (React)</h3>
                    <p className="text-sm text-muted-foreground">
                      Interface moderna que se comunica com a API Go
                    </p>
                  </div>
                  <div className="flex flex-col items-center text-center">
                    <div className="mb-4 rounded-lg bg-gradient-primary p-4">
                      <Zap className="h-8 w-8 text-primary-foreground" />
                    </div>
                    <h3 className="mb-2 font-semibold">API Go + RabbitMQ</h3>
                    <p className="text-sm text-muted-foreground">
                      Recebe requisições e enfileira no RabbitMQ
                    </p>
                  </div>
                  <div className="flex flex-col items-center text-center">
                    <div className="mb-4 rounded-lg bg-gradient-primary p-4">
                      <Lock className="h-8 w-8 text-primary-foreground" />
                    </div>
                    <h3 className="mb-2 font-semibold">Processador Rust</h3>
                    <p className="text-sm text-muted-foreground">
                      Consome fila, aplica lógica e persiste no PostgreSQL
                    </p>
                  </div>
                </div>
                <div className="mt-8 rounded-lg bg-muted p-4">
                  <p className="text-center text-sm text-muted-foreground">
                    Frontend → API Go → RabbitMQ → Processador Rust → PostgreSQL
                    + Redis
                  </p>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </section>

      {/* API Documentation Section */}
      <section className="py-16 bg-secondary/30">
        <div className="container mx-auto px-4">
          <div className="mx-auto max-w-4xl">
            <h2 className="mb-8 text-center text-3xl font-bold">
              Documentação da API
            </h2>
            <div className="grid gap-6 md:grid-cols-2">
              {/* Swagger Card */}
              <Card className="shadow-card hover:shadow-elegant transition-shadow">
                <CardHeader>
                  <div className="mb-2 flex items-center gap-2">
                    <div className="rounded-lg bg-gradient-primary p-2">
                      <FileText className="h-5 w-5 text-primary-foreground" />
                    </div>
                    <CardTitle>Swagger API Docs</CardTitle>
                  </div>
                  <CardDescription>
                    Documentação interativa da API com todos os endpoints
                    disponíveis
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <Button asChild className="w-full gap-2">
                    <a
                      href={`${API_BASE_URL}/swagger/index.html`}
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      <ExternalLink className="h-4 w-4" />
                      Abrir Swagger
                    </a>
                  </Button>
                </CardContent>
              </Card>

              {/* Postman Card */}
              <Card className="shadow-card hover:shadow-elegant transition-shadow">
                <CardHeader>
                  <div className="mb-2 flex items-center gap-2">
                    <div className="rounded-lg bg-gradient-success p-2">
                      <Download className="h-5 w-5 text-white" />
                    </div>
                    <CardTitle>Postman Collection</CardTitle>
                  </div>
                  <CardDescription>
                    Baixe a coleção do Postman com exemplos de todas as
                    requisições
                  </CardDescription>
                </CardHeader>
                <CardContent>
                  <Button
                    asChild
                    variant="secondary"
                    className="w-full gap-2"
                    onClick={downloadPostmanCollection}
                  >
                    <div>
                      <Download className="h-4 w-4" />
                      Baixar Collection
                    </div>
                  </Button>
                </CardContent>
              </Card>
            </div>
          </div>
        </div>
      </section>

      {/* Key Concepts Section */}
      <section className="py-16">
        <div className="container mx-auto px-4">
          <div className="mx-auto max-w-3xl">
            <h2 className="mb-8 text-center text-3xl font-bold">
              Conceitos Chave
            </h2>
            <Accordion type="single" collapsible className="w-full">
              <AccordionItem value="ledger">
                <AccordionTrigger className="text-lg font-semibold">
                  Ledger (Razão Contábil)
                </AccordionTrigger>
                <AccordionContent className="text-muted-foreground">
                  O Ledger é a "fonte da verdade" para todos os saldos no
                  sistema. Cada transação é registrada de forma imutável,
                  garantindo que o histórico financeiro seja sempre preciso e
                  auditável. Nenhuma transação pode ser alterada após ser
                  registrada, apenas revertida por outra transação.
                </AccordionContent>
              </AccordionItem>

              <AccordionItem value="idempotency">
                <AccordionTrigger className="text-lg font-semibold">
                  Idempotência
                </AccordionTrigger>
                <AccordionContent className="text-muted-foreground">
                  Garantir que uma mesma operação, repetida várias vezes,
                  produza sempre o mesmo resultado. Isso evita problemas como
                  cobranças duplicadas caso uma requisição seja enviada mais de
                  uma vez (por exemplo, por falha de rede ou clique duplo do
                  usuário).
                </AccordionContent>
              </AccordionItem>

              <AccordionItem value="queue">
                <AccordionTrigger className="text-lg font-semibold">
                  Fila de Mensagens (RabbitMQ)
                </AccordionTrigger>
                <AccordionContent className="text-muted-foreground">
                  O RabbitMQ desacopla a API do processamento das transações.
                  Isso significa que mesmo sob alta carga, nenhuma transação é
                  perdida. As mensagens ficam na fila até serem processadas pelo
                  serviço Rust, garantindo resiliência e escalabilidade.
                </AccordionContent>
              </AccordionItem>

              <AccordionItem value="tech">
                <AccordionTrigger className="text-lg font-semibold">
                  Por que Go e Rust?
                </AccordionTrigger>
                <AccordionContent className="text-muted-foreground">
                  <strong>Go</strong> foi escolhido pela sua simplicidade, alta
                  performance em I/O e excelente suporte para concorrência,
                  ideal para APIs web. <strong>Rust</strong> foi escolhido pela
                  sua segurança de memória e performance, perfeito para
                  processamento crítico de transações onde cada milissegundo e
                  bug importa.
                </AccordionContent>
              </AccordionItem>
            </Accordion>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="bg-gradient-primary py-16 text-primary-foreground mt-0">
        <div className="container mx-auto px-4 text-center">
          <h2 className="mb-4 text-3xl font-bold">Pronto para começar?</h2>
          <p className="mb-8 text-lg opacity-90">
            Crie sua conta e explore o sistema de transações
          </p>
          <Button asChild size="lg" variant="secondary">
            <Link to="/create-account">
              Criar minha conta <ArrowRight className="ml-2 h-5 w-5" />
            </Link>
          </Button>
        </div>
      </section>
    </div>
  );
}
