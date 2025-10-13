// Configurações globais de SEO para o PayGateway
export const SEO_CONFIG = {
  SITE_NAME: "PayGateway",
  SITE_URL: "https://paygateway.exemplo.com",
  DEFAULT_TITLE: "PayGateway - Gateway de Pagamento Fictício",
  DEFAULT_DESCRIPTION:
    "Sistema completo de gateway de pagamento desenvolvido com arquitetura moderna usando Rust, Go e React. Processamento seguro de transações financeiras.",
  DEFAULT_KEYWORDS:
    "gateway de pagamento, processamento de transações, fintech, rust, go, react, sistema financeiro, api de pagamentos",
  TWITTER_HANDLE: "@paygateway",
  CREATOR_TWITTER: "@eduardomg12",
  AUTHOR: "Eduardo MG",
  GITHUB_URL: "https://github.com/EduardoMG12/payment-gateway-challenge",

  // Imagens para Open Graph e Twitter
  OG_IMAGE: "https://paygateway.exemplo.com/og-image.png",
  TWITTER_IMAGE: "https://paygateway.exemplo.com/twitter-image.png",

  // Configurações de Schema.org
  ORGANIZATION_SCHEMA: {
    "@context": "https://schema.org",
    "@type": "SoftwareApplication",
    name: "PayGateway",
    description: "Sistema de gateway de pagamento com arquitetura distribuída",
    url: "https://paygateway.exemplo.com",
    applicationCategory: "FinanceApplication",
    operatingSystem: "Web",
    creator: {
      "@type": "Person",
      name: "Eduardo MG",
      url: "https://github.com/EduardoMG12",
    },
    offers: {
      "@type": "Offer",
      price: "0",
      priceCurrency: "BRL",
    },
    features: [
      "Processamento de transações seguras",
      "Gerenciamento de contas",
      "Criação de cartões virtuais",
      "Extratos e histórico detalhado",
      "API RESTful completa",
      "Arquitetura de microsserviços",
      "Tecnologias: Rust, Go, React",
    ],
  },
};

// Páginas e suas configurações específicas de SEO
export const PAGE_SEO_CONFIG = {
  HOME: {
    title: "PayGateway - Gateway de Pagamento Fictício",
    description:
      "Sistema completo de gateway de pagamento desenvolvido com arquitetura moderna usando Rust, Go e React. Processamento seguro de transações financeiras com API RESTful completa.",
    keywords:
      "gateway de pagamento, processamento de transações, fintech, rust, go, react, sistema financeiro, api de pagamentos, transações seguras, sistema distribuído",
    path: "/",
  },

  CREATE_ACCOUNT: {
    title: "Criar Conta - PayGateway",
    description:
      "Crie sua conta gratuita no PayGateway e comece a processar transações de forma segura. Cadastro rápido e fácil para acessar todas as funcionalidades do sistema.",
    keywords:
      "criar conta, cadastro gratuito, registro paygateway, nova conta, conta financeira, cadastro sistema pagamento",
    path: "/create-account",
  },

  CREATE_CARD: {
    title: "Criar Cartão Virtual - PayGateway",
    description:
      "Crie cartões virtuais seguros para suas transações no PayGateway. Gere tokens únicos e gerencie seus métodos de pagamento de forma segura e eficiente.",
    keywords:
      "criar cartão virtual, cartão digital, token seguro, método de pagamento, cartão online, pagamento seguro",
    path: "/create-card",
  },

  TRANSACTIONS: {
    title: "Transações - PayGateway",
    description:
      "Realize transações seguras, processamento de pagamentos, estornos e gerencie suas operações financeiras no PayGateway. Interface completa para operações bancárias.",
    keywords:
      "transações financeiras, processar pagamento, estorno, operações bancárias, transferência, pagamento online, transação segura",
    path: "/transactions",
  },

  STATEMENT: {
    title: "Extrato e Histórico - PayGateway",
    description:
      "Visualize seu extrato completo, histórico de transações, saldo atual e filtre operações por cartão no PayGateway. Controle total das suas finanças.",
    keywords:
      "extrato bancário, histórico transações, saldo conta, relatório financeiro, extrato cartão, consulta saldo, movimentação financeira",
    path: "/statement",
  },

  NOT_FOUND: {
    title: "Página não encontrada - PayGateway",
    description:
      "A página que você está procurando não foi encontrada. Retorne à página inicial do PayGateway ou explore nossas funcionalidades de gateway de pagamento.",
    keywords: "página não encontrada, erro 404, paygateway, voltar início",
    path: "/404",
  },
};

// Funções utilitárias para SEO
export const generateCanonicalUrl = (path: string): string => {
  return `${SEO_CONFIG.SITE_URL}${path}`;
};

export const generateFullTitle = (pageTitle?: string): string => {
  if (!pageTitle || pageTitle === SEO_CONFIG.DEFAULT_TITLE) {
    return SEO_CONFIG.DEFAULT_TITLE;
  }
  return `${pageTitle} | ${SEO_CONFIG.SITE_NAME}`;
};

export const generateBreadcrumbSchema = (
  pages: Array<{ name: string; url: string }>,
) => {
  return {
    "@context": "https://schema.org",
    "@type": "BreadcrumbList",
    itemListElement: pages.map((page, index) => ({
      "@type": "ListItem",
      position: index + 1,
      name: page.name,
      item: page.url,
    })),
  };
};
