import { useState } from "react";
import { useNavigate } from "react-router-dom";
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
import { accountsApi } from "@/services/api";
import { useUserStore } from "@/store/userStore";
import { useSEO } from "@/hooks/useSEO";
import toast from "react-hot-toast";
import { UserPlus } from "lucide-react";

export default function CreateAccountPage() {
  const seoData = useSEO({
    title: "Criar Conta - PayGateway",
    description:
      "Crie sua conta gratuita no PayGateway e comece a processar transações de forma segura. Cadastro rápido e fácil para acessar todas as funcionalidades do sistema.",
    keywords:
      "criar conta, cadastro gratuito, registro paygateway, nova conta, conta financeira, cadastro sistema pagamento",
    ogType: "website",
    schema: {
      "@context": "https://schema.org",
      "@type": "WebPage",
      name: "Criar Conta - PayGateway",
      description:
        "Página de cadastro para criar nova conta no sistema PayGateway",
      url: "https://paygateway.exemplo.com/create-account",
      mainEntity: {
        "@type": "CreateAction",
        name: "Criar Conta",
        description: "Formulário para criação de nova conta de usuário",
        target: {
          "@type": "EntryPoint",
          urlTemplate: "https://paygateway.exemplo.com/create-account",
          actionApplication: {
            "@type": "WebApplication",
            name: "PayGateway",
          },
        },
      },
    },
  });

  const [username, setUsername] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  const setAccount = useUserStore((state) => state.setAccount);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!username.trim()) {
      toast.error("Por favor, insira um nome de usuário");
      return;
    }

    setLoading(true);

    try {
      const response = await accountsApi.create({ username: username.trim() });
      setAccount(response.data.id, response.data.username);
      toast.success("Conta criada com sucesso!");
      navigate("/create-card");
    } catch (error) {
      console.error("Erro ao criar conta:", error);
      toast.error("Erro ao criar conta. Tente novamente.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto flex min-h-[calc(100vh-4rem)] items-center justify-center px-4 py-8">
      {seoData}
      <Card className="w-full max-w-md shadow-elegant">
        <CardHeader className="text-center">
          <div className="mx-auto mb-4 flex h-16 w-16 items-center justify-center rounded-full bg-gradient-primary">
            <UserPlus className="h-8 w-8 text-primary-foreground" />
          </div>
          <CardTitle className="text-2xl">Criar Nova Conta</CardTitle>
          <CardDescription>
            Comece criando sua conta no gateway de pagamento
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="space-y-2">
              <Label htmlFor="username">Nome de usuário</Label>
              <Input
                id="username"
                type="text"
                placeholder="Digite seu nome de usuário"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                disabled={loading}
                className="h-12"
              />
            </div>

            <Button type="submit" className="w-full h-12" disabled={loading}>
              {loading ? "Criando..." : "Criar Conta"}
            </Button>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
