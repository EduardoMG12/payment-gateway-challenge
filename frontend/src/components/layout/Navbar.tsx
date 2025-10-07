import { Link, useLocation } from "react-router-dom";
import { CreditCard, Wallet } from "lucide-react";
import { Button } from "@/components/ui/button";
import { useUserStore } from "@/store/userStore";

export const Navbar = () => {
  const location = useLocation();
  const { username, clearAccount } = useUserStore();

  const isActive = (path: string) => location.pathname === path;

  const navLinkClass = (path: string) =>
    `transition-colors ${
      isActive(path)
        ? "text-primary font-medium"
        : "text-muted-foreground hover:text-foreground"
    }`;

  return (
    <nav className="border-b bg-card shadow-sm">
      <div className="container mx-auto px-4">
        <div className="flex h-16 items-center justify-between">
          <Link to="/" className="flex items-center gap-2">
            <div className="rounded-lg bg-gradient-primary p-2">
              <Wallet className="h-5 w-5 text-primary-foreground" />
            </div>
            <span className="text-xl font-bold">PayGateway</span>
          </Link>

          <div className="flex items-center gap-6">
            <Link to="/" className={navLinkClass("/")}>
              Home
            </Link>
            {!username && (
              <Link
                to="/create-account"
                className={navLinkClass("/create-account")}
              >
                Criar Conta
              </Link>
            )}
            {username && (
              <>
                <Link
                  to="/create-card"
                  className={navLinkClass("/create-card")}
                >
                  Criar Cartão
                </Link>
                <Link
                  to="/transactions"
                  className={navLinkClass("/transactions")}
                >
                  Transações
                </Link>
                <Link to="/statement" className={navLinkClass("/statement")}>
                  Extrato
                </Link>
              </>
            )}
            {username && (
              <div className="flex items-center gap-3 border-l pl-6">
                <div className="flex items-center gap-2 rounded-lg bg-secondary px-3 py-1.5">
                  <CreditCard className="h-4 w-4 text-primary" />
                  <span className="text-sm font-medium">{username}</span>
                </div>
                <Button variant="ghost" size="sm" onClick={clearAccount}>
                  Sair
                </Button>
              </div>
            )}
          </div>
        </div>
      </div>
    </nav>
  );
};
